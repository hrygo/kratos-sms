package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"

	pb "kratos-sms/api/sms/v1"
	"kratos-sms/internal/biz"
	"kratos-sms/internal/conf"
	mylog "kratos-sms/internal/log"
)

type SmsService struct {
	pb.UnimplementedSmsServer
	conf config.Config
	bs   *conf.Bootstrap
	uc   *biz.SmsUseCase
	log  *log.Helper
}

func (s *SmsService) Replace(logger *log.Helper) {
	s.log = logger
}

func NewSmsService(cc config.Config, bs *conf.Bootstrap, uc *biz.SmsUseCase, logger log.Logger) *SmsService {
	if bs.Log.Filter == nil {
		log.Fatal("Configuration `bs.log.filter` has not set.")
	}

	logger = log.With(logger, "caller", log.Caller(6))
	srv := &SmsService{
		conf: cc,
		bs:   bs,
		uc:   uc,
		log: log.NewHelper(log.NewFilter(logger,
			mylog.KeyFilter,
			mylog.FilterLevel(bs.Log.Filter.ServiceLogLevel),
		)),
	}
	// 根据配置变更动态修改过滤器级别
	mylog.FilterChangeWatch(cc, srv, mylog.ServiceLogLvlConfKey, logger)

	return srv
}

func (s *SmsService) TextMessageSend(ctx context.Context, req *pb.TextMessageRequest) (*pb.SendMessageReply, error) {
	return s.uc.SendSmsWithJournal(ctx, req)
}
func (s *SmsService) TemplateMessageSend(ctx context.Context, req *pb.TemplateMessageRequest) (*pb.SendMessageReply, error) {
	// 1. *pb.TemplateMessageRequest ===> *pb.TextMessageRequest
	textReq := &pb.TextMessageRequest{
		Phones: req.Phones,
		AtTime: req.AtTime,
	}
	return s.uc.SendSmsWithJournal(ctx, textReq)
}
func (s *SmsService) AsyncResultQuery(ctx context.Context, req *pb.AsyncResultQueryRequest) (*pb.AsyncResultQueryReply, error) {
	return s.uc.QueryAsyncResults(ctx, req.QueryId)
}
