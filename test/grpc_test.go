package test

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	md "github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	v1 "kratos-sms/api/sms/v1"
)

var conn, _ = grpc.DialInsecure(
	context.Background(),
	grpc.WithEndpoint("127.0.0.1:9000"),
	grpc.WithMiddleware(
		metadata.Client(),
	))

var smsClient = v1.NewSmsClient(conn)

func TestSendMsg(t *testing.T) {
	ctx := md.AppendToClientContext(context.Background(),
		"x-md-global-app-id", "0123456789012345",
		"x-md-global-token", "0123456789012345012345678901234501234567890123450123456789012345",
	)
	reply, err := smsClient.TextMessageSend(ctx, &v1.TextMessageRequest{
		Content: "hello world",
		Phones:  []string{"17600460000"},
	})
	if err != nil {
		return
	}
	log.Info(reply)
}
