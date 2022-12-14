// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-sms/internal/biz"
	"kratos-sms/internal/conf"
	"kratos-sms/internal/data"
	"kratos-sms/internal/server"
	"kratos-sms/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(configConfig config.Config, bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	smsRepo := data.NewSmsRepo(configConfig, bootstrap, dataData, logger)
	smsUseCase := biz.NewSmsUseCase(configConfig, bootstrap, smsRepo, logger)
	smsService := service.NewSmsService(configConfig, bootstrap, smsUseCase, logger)
	grpcServer := server.NewGRPCServer(bootstrap, smsService, logger)
	httpServer := server.NewHTTPServer(bootstrap, smsService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
