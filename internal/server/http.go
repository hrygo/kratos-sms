package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"

	sms "kratos-sms/api/sms/v1"
	"kratos-sms/internal/conf"
	"kratos-sms/internal/service"
)

// NewHTTPServer new HTTP server.
func NewHTTPServer(bs *conf.Bootstrap, service *service.SmsService, logger log.Logger) *http.Server {
	c := bs.GetServer()
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			metadata.Server(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	sms.RegisterSmsHTTPServer(srv, service)

	log.NewHelper(logger).Debug("Create HTTP Server.")
	return srv
}
