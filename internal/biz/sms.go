package biz

import (
	"context"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"

	pb "kratos-sms/api/sms/v1"
	"kratos-sms/internal/conf"
	mylog "kratos-sms/internal/log"
	"kratos-sms/pkg/properties"
)

type SmsRepo interface {
	SaveJournal(ctx context.Context, jo *SmsJournal) (*SmsJournal, error)
	QueryJournal(ctx context.Context, queryId uint64) (*SmsJournal, error)
	FindTemplate(ctx context.Context, tempId string) (*SmsTemplate, error)
}

type SmsUseCase struct {
	conf config.Config
	bs   *conf.Bootstrap
	repo SmsRepo
	log  *log.Helper
}

func (uc *SmsUseCase) Replace(logger *log.Helper) {
	uc.log = logger
}

func NewSmsUseCase(cc config.Config, bs *conf.Bootstrap, repo SmsRepo, logger log.Logger) *SmsUseCase {
	if bs.Log.Filter == nil {
		log.Fatal("Configuration `bs.log.filter` has not set.")
	}

	logger = log.With(logger, "caller", log.Caller(6))
	uc := &SmsUseCase{
		conf: cc,
		bs:   bs,
		repo: repo,
		log: log.NewHelper(log.NewFilter(logger,
			mylog.KeyFilter,
			mylog.FilterLevel(bs.Log.Filter.BizLogLevel),
		)),
	}
	// 根据配置变更动态修改过滤器级别
	mylog.FilterChangeWatch(cc, uc, mylog.BizLogLvlConfKey, logger)

	return uc
}

type SmsJournal struct {
	Id       string             `bson:"_id,omitempty"`      // ID
	AppId    string             `bson:"appId,omitempty"`    // 服务端分配的AppId
	Content  string             `bson:"content,omitempty"`  // 消息内容
	Priority int32              `bson:"priority,omitempty"` // 消息优先级
	AtTime   time.Time          `bson:"ts,omitempty"`       // 定时发送时间 （时间戳, 大于当前时间, 24小时以内）
	Phones   []string           `bson:"phones,omitempty"`   // 手机号列表
	QueryId  uint64             `bson:"queryId,omitempty"`  // 查询ID
	Code     int32              `bson:"code,omitempty"`     // 状态码, 200 成功，其他
	Message  string             `bson:"message,omitempty"`  // 状态描述信息
	Results  []*AsyncResultList `bson:"results,omitempty"`  // 发送结果
}

type AsyncResultList struct {
	Phone        string    `bson:"phone,omitempty"`        // 手机号
	SequenceId   uint64    `bson:"sequenceId,omitempty"`   // 短信发送流水号
	Result       uint32    `bson:"result,omitempty"`       // 运营商网关响应
	MsgId        string    `bson:"msgId,omitempty"`        // 运营商网关短信编号
	SendTime     time.Time `bson:"sendTime,omitempty"`     // 消息发送时间
	ResponseTime time.Time `bson:"responseTime,omitempty"` // 运营商网关响应时间
	ReportTime   time.Time `bson:"reportTime,omitempty"`   // 状态报告接收到的时间
	Report       string    `bson:"report,omitempty"`       // 状态报告内容

}

var smsJournalKeys = []string{
	"Id",
	"AppId",
	"Content",
	"Priority",
	"AtTime",
	"Phones",
	"QueryId",
	"Code",
	"Message",
	"Results",
	"Auth",
}

// Keys 返回需要打印到日志中的字段
func (j *SmsJournal) Keys() []string {
	return smsJournalKeys
}

type SmsTemplate struct {
	Id               uint64   `bson:"_id,omitempty"`              // ID
	TempId           string   `bson:"tempId,omitempty"`           // 模板编号
	Template         string   `bson:"template,omitempty"`         // 模板内容
	Priority         int32    `bson:"priority,omitempty"`         // 优先级
	ProhibitedPeriod []uint32 `bson:"prohibitedPeriod,omitempty"` // 禁止发送时段，数字含义为当前距离零点的秒数，两个数字一对，前者必须小于后者
}

func (st *SmsTemplate) Parse(args map[string]string) string {
	s := st.Template
	for key, value := range args {
		s = strings.ReplaceAll(st.Template, "__"+key+"__", value)
	}
	return s
}

func (st *SmsTemplate) Allowed() (allow bool) {
	allow = true
	if len(st.ProhibitedPeriod) < 2 {
		return
	}
	seconds := time.Now().Second()
	for i := 0; i < len(st.ProhibitedPeriod)-1; i += 2 {
		// 小于左边界，大于右边界，说明不在禁止发送的区间内
		allow = allow && (seconds < int(st.ProhibitedPeriod[i]))
		allow = allow && (seconds > int(st.ProhibitedPeriod[i+1]))
	}
	return
}

// SendSmsWithJournal 由Service层调用该方法
func (uc *SmsUseCase) SendSmsWithJournal(ctx context.Context, req *pb.TextMessageRequest) (*pb.SendMessageReply, error) {
	// 1. 调用短信网关的接口发送短信
	jo := &SmsJournal{}
	properties.Copy(req, jo)
	jo.AtTime = req.AtTime.AsTime()
	uc.log.WithContext(ctx).Debugw(properties.KVPairs(jo, "SMS.Jo.")...)

	// 2. 发送结果入库（异步）
	journal, err := uc.repo.SaveJournal(ctx, jo)
	if err != nil {
		return nil, err
	}
	uc.log.WithContext(ctx).Debugw(properties.KVPairs(journal, "SMS.Jo.")...)

	// 3. 返回发送结果给service层
	reply := &pb.SendMessageReply{
		QueryId: "0123456789ABCDEF",
		Status: &pb.ReplyStatus{
			Code:    200,
			Message: "SUCCESS",
		},
	}
	return reply, nil
}

func (uc *SmsUseCase) QueryAsyncResults(ctx context.Context, queryId uint64) (*pb.AsyncResultQueryReply, error) {
	// 1. 查询数据库，看是否已经有返回结果(可增加二级缓存)
	journal, err := uc.repo.QueryJournal(ctx, queryId)
	if err != nil {
		err := pb.ErrorNotFound("Query DB error: %v", err)
		return nil, err
	}
	uc.log.WithContext(ctx).Debugf("query journal: %v", journal)

	// 2. 如果1未查到，从短信网关查询，并更新数据库（可增加二级缓存）

	// 3. 返回查询结果
	reply := &pb.AsyncResultQueryReply{
		QueryId: queryId,
		Results: nil,
	}
	return reply, nil
}
