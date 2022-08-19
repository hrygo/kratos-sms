//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"kratos-sms/internal/biz"
	"kratos-sms/internal/conf"
	"kratos-sms/internal/data"
	"kratos-sms/internal/server"
	"kratos-sms/internal/service"
)

// wireApp init kratos application.
func wireApp(config.Config, *conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
