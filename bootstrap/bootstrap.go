package bootstrap

import (
	"syscall"

	"github.com/go-kratos/kratos/v2"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	logger "github.com/origadmin/slog-karatos"
	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/loge"
)

func Run(ctx context.Context, cfg Config) error {
	//config.Load(cfg.ConfigPath, cfg.ConfigType)
	opts := []kratos.Option{
		kratos.Name(cfg.Name),
		kratos.Context(ctx),
		kratos.Signal(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
		kratos.Logger(logger.NewLogger(loge.New(loge.WithUseDefault()))),
	}

	srv := transhttp.NewServer(transhttp.Address(":8000"))
	r := srv.Route("/")
	r.GET("/healthz", func(c transhttp.Context) error {
		return c.Result(200, "ok")
	})

	opts = append(opts, kratos.Server(
		srv,
	))

	app := kratos.New(opts...)

	return app.Run()
	//b := redis.NewBroker(
	//	broker.WithCodec("json"),
	//	broker.WithAddress(localBroker),
	//)
	//
	//_ = b.Init()
	//
	//if err := b.Connect(); err != nil {
	//	fmt.Println(err)
	//}
	//defer func(b broker.Broker) {
	//	err := b.Disconnect()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}(b)
	//
	//_, _ = b.Subscribe(testTopic,
	//	api.RegisterHygrothermographJsonHandler(handleHygrothermograph),
	//	api.HygrothermographCreator,
	//)

	return nil
}
