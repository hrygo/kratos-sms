package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/config"

	pb "kratos-sms/api/sms/v1"
	"kratos-sms/internal/biz"
)

type SmsService struct {
	pb.UnimplementedSmsServer
	conf config.Config
	uc   *biz.SmsUseCase
}

func NewSmsService(conf config.Config, uc *biz.SmsUseCase) *SmsService {
	return &SmsService{conf: conf, uc: uc}
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
