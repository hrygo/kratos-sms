package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	sms "kratos-sms/api/sms/v1"
	"kratos-sms/internal/conf"
	"kratos-sms/internal/service"
)

// NewGRPCServer new gRPC server.
func NewGRPCServer(bs *conf.Bootstrap, service *service.SmsService, logger log.Logger) *grpc.Server {
	c := bs.GetServer()
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			metadata.Server(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	sms.RegisterSmsServer(srv, service)

	log.NewHelper(logger).Debug("Create gRPC Server.")
	return srv
}
