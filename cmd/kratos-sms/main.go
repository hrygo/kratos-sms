package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"

	"kratos-sms/internal/conf"
	mylog "kratos-sms/internal/log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// confPath is the config file path.
	confPath string

	id, _ = os.Hostname()
)

func init() {
	// json.MarshalOptions = protojson.MarshalOptions{
	//   EmitUnpopulated: true,
	//   UseProtoNames:   true,
	// }
	Name = "kratos-sms"
	Version = "v0.0.1"
	flag.StringVar(&confPath, "conf", "./configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	// c, bc, closeFunc := fileConfig()
	c, bc, closeFunc := consulConfig()
	defer closeFunc()
	if bc == nil || bc.Server == nil || bc.Log == nil || bc.Data == nil {
		os.Exit(-1)
	}

	logger := customLogger(c, bc.Log)
	defer mylog.Sync()

	app, cleanup, err := wireApp(c, bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func customLogger(c config.Config, lc *conf.Log) log.Logger {
	resetLogger(lc)
	logger := log.With(mylog.Default(),
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	_ = c.Watch("log", func(key string, value config.Value) {
		err := value.Scan(lc)
		if err != nil {
			mylog.Errorf(err.Error())
			return
		}
		resetLogger(lc)
		// TODO logger的私有成员(.logger) 如何改变？如不能改变，则日志重置其实是失效的
		mylog.Warnf("logger was reset!")
	})
	return logger
}

func resetLogger(lc *conf.Log) {
	mylog.ProductionDefault(lc,
		zap.AddStacktrace(mylog.ErrorLevel),
		zap.Hooks(),
	)
}

// 从文件读取配置
func fileConfig() (config.Config, *conf.Bootstrap, func()) {
	c := config.New(
		config.WithSource(
			file.NewSource(confPath),
		),
	)
	defer func() { _ = c.Close() }()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	return c, &bc, func() { _ = c.Close() }
}

// 从Consul读取配置
func consulConfig() (config.Config, *conf.Bootstrap, func()) {
	// 注意consul里面配置的名称，必须有后缀，否则无法正确解析
	var path = "kratos/sms.yaml"
	consulClient, err := api.NewClient(&api.Config{
		Address: "127.0.0.1:8500",
	})
	if err != nil {
		panic(err)
	}
	cs, err := consul.New(
		consulClient,
		consul.WithPath(path),
	)
	if err != nil {
		panic(err)
	}
	c := config.New(config.WithSource(cs))

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	return c, &bc, func() { _ = c.Close() }
}
