package bootstrap

import (
	"syscall"

	"github.com/go-kratos/kratos/v2"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	logger "github.com/origadmin/slog-karatos"
	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/loge"

	"origadmin/backend/internal/config"
)

func Run(ctx context.Context, cfg Config) error {
	config, err := config.Load(cfg.Dir, cfg.Configs...)
	if err != nil {
		return err
	}

	opts := []kratos.Option{
		kratos.Name(config.Settings.ServiceName),
		kratos.Context(ctx),
		kratos.Signal(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
		kratos.Logger(logger.NewLogger(loge.New(loge.WithUseDefault()))),
	}

	srv := transhttp.NewServer(transhttp.Address(config.Settings.HTTP.Addr))
	r := srv.Route("/")
	r.GET("/healthz", func(c transhttp.Context) error {
		return c.Result(200, "ok")
	})

	opts = append(opts, kratos.Server(
		srv,
	))

	app := kratos.New(opts...)

	return app.Run()
}
