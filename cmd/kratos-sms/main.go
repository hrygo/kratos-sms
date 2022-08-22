package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
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

	// 1. 从配置文件读取配置
	c, bc, closeFunc := fileConfig()
	defer closeFunc()

	// 2. 如果有配置Consul，则采用Consul配置覆盖本地配置
	if bc.Consul != nil && bc.Consul.Address != "" {
		closeFunc()
		c, bc, closeFunc = consulConfig(bc.Consul)
		defer closeFunc()
	}
	if bc == nil || bc.Server == nil || bc.Log == nil || bc.Data == nil {
		log.Fatal("App will exit, because the necessary configuration information is missing!")
		os.Exit(-1)
	}
	// 为 app_debug 变化增加观察者
	_ = c.Watch("app_debug", observerDebugMode(bc))

	// 3. 设置日志
	logger := customLogger(bc)
	defer mylog.Sync()

	// 4. 装配APP
	app, cleanup, err := wireApp(c, bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 5. Start and wait for stop signal
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

func customLogger(bs *conf.Bootstrap) log.Logger {
	mylog.ProductionDefault(bs,
		zap.AddStacktrace(mylog.ErrorLevel),
		zap.Hooks(), // 此处可以添加回调钩子，用于日志同步ElasticSearch等操作
	)

	logger := log.With(mylog.Default(),
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	return logger
}

func observerDebugMode(bs *conf.Bootstrap) func(key string, value config.Value) {
	return func(key string, value config.Value) {
		mode, err := value.Bool()
		if err != nil {
			log.Error(err)
			return
		}
		bs.AppDebug = mode
		log.Warnf("App debug mode changed to %v!", mode)
		// TODO MORE
	}
}

// 从文件读取配置
func fileConfig() (config.Config, *conf.Bootstrap, func()) {
	c := config.New(
		config.WithSource(
			// 添加前缀为 KRATOS_ 的环境变量
			env.NewSource("KRATOS_"),
			file.NewSource(confPath),
		),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	return c, &bc, func() { _ = c.Close() }
}

// 从Consul读取配置, 注意consul里面配置的名称，必须有后缀，否则无法正确解析
func consulConfig(cc *conf.Consul) (config.Config, *conf.Bootstrap, func()) {
	apic := &api.Config{
		Address: cc.GetAddress(),
		Scheme:  cc.GetScheme(),
	}
	if cc.WaitTime != nil {
		apic.WaitTime = cc.WaitTime.AsDuration()
	}
	consulClient, err := api.NewClient(apic)
	if err != nil {
		panic(err)
	}
	var path = "kratos/application.yaml"
	if cc.Path != "" {
		path = cc.Path
	}
	cs, err := consul.New(
		consulClient,
		consul.WithPath(path),
	)
	if err != nil {
		panic(err)
	}
	c := config.New(config.WithSource(
		// 添加前缀为 KRATOS_ 的环境变量
		env.NewSource("KRATOS_"),
		cs,
	))

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	return c, &bc, func() { _ = c.Close() }
}
