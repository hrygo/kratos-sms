package service

import (
  "context"

  "github.com/go-kratos/kratos/v2/config"
  "github.com/go-kratos/kratos/v2/log"

  pb "kratos-sms/api/sms/v1"
  "kratos-sms/internal/biz"
  "kratos-sms/internal/conf"
)

type SmsService struct {
  pb.UnimplementedSmsServer
  conf config.Config
  bs   *conf.Bootstrap
  uc   *biz.SmsUseCase
  log  *log.Helper
}

func NewSmsService(conf config.Config, bs *conf.Bootstrap, uc *biz.SmsUseCase, logger log.Logger) *SmsService {
  return &SmsService{
    conf: conf,
    bs:   bs,
    uc:   uc,
    log:  log.NewHelper(logger),
  }
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
