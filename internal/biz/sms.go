package biz

import (
	"context"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	pb "kratos-sms/api/sms/v1"
)

type SmsRepo interface {
	SaveJournal(ctx context.Context, jo *SmsJournal) (*SmsJournal, error)
	QueryJournal(ctx context.Context, queryId uint64) (*SmsJournal, error)
	FindTemplate(ctx context.Context, tempId string) (*SmsTemplate, error)
}

type SmsJournal struct {
	Id       uint64             // ID
	AppId    string             // 服务端分配的AppId
	Content  string             // 消息内容
	Priority int32              // 消息优先级
	AtTime   time.Time          // 定时发送时间 （时间戳, 大于当前时间, 24小时以内）
	Phones   []string           // 手机号列表
	QueryId  uint64             // 查询ID
	Code     int32              // 状态码, 200 成功，其他
	Message  string             // 状态描述信息
	Results  []*AsyncResultList // 发送结果
}

type AsyncResultList struct {
	Phone        string    // 手机号
	SequenceId   uint64    // 短信发送流水号
	Result       uint32    // 运营商网关响应
	MsgId        string    // 运营商网关短信编号
	SendTime     time.Time // 消息发送时间
	ResponseTime time.Time // 运营商网关响应时间
	ReportTime   time.Time // 状态报告接收到的时间
	Report       string    // 状态报告内容

}

type SmsTemplate struct {
	Id               uint64   // ID
	TempId           string   // 模板编号
	Template         string   // 模板内容
	Priority         int32    // 优先级
	ProhibitedPeriod []uint32 // 禁止发送时段，数字含义为当前距离零点的秒数，两个数字一对，前者必须小于后者
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

type SmsUseCase struct {
	repo SmsRepo
	log  *log.Helper
}

func NewSmsUseCase(repo SmsRepo, logger log.Logger) *SmsUseCase {
	return &SmsUseCase{repo: repo, log: log.NewHelper(logger)}
}

// SendSmsWithJournal 由Service层调用该方法
func (uc *SmsUseCase) SendSmsWithJournal(ctx context.Context, req *pb.TextMessageRequest) (*pb.SendMessageReply, error) {
	// 1. 调用短信网关的接口发送短信
	uc.log.WithContext(ctx).Debugf("send sms: %v", req)
	// send result to journal
	jo := &SmsJournal{
		QueryId: 1,
	}

	// 2. 发送结果入库（异步）
	journal, err := uc.repo.SaveJournal(ctx, jo)
	if err != nil {
		return nil, err
	}
	uc.log.WithContext(ctx).Debugf("save journal: %v", journal)

	// 3. 返回发送结果给service层
	reply := &pb.SendMessageReply{
		QueryId: 1,
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
