package bootstrap

import (
	"crypto/tls"
	"syscall"
	"time"

	"github.com/go-kratos/kratos/contrib/opensergo/v2"
	"github.com/go-kratos/kratos/v2"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	logger "github.com/origadmin/slog-kratos"
	"github.com/origadmin/toolkits/context"

	"origadmin/backend/internal/config"
)

func Run(ctx context.Context, cfg Config) error {
	config, err := config.Load(cfg.WorkDir, cfg.Configs...)
	if err != nil {
		return err
	}

	opts := []kratos.Option{
		kratos.Name(config.Settings.ServiceName),
		kratos.Context(ctx),
		kratos.Signal(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
		kratos.Logger(logger.NewLogger()),
	}

	var httpOpts []transhttp.ServerOption
	if config.Settings.HTTP.Addr != "" {
		httpOpts = append(httpOpts, transhttp.Address(config.Settings.HTTP.Addr))
	} else {
		httpOpts = append(httpOpts, transhttp.Address(":28080"))
	}

	if config.Settings.HTTP.UseTLS {
		httpOpts = append(httpOpts, transhttp.TLSConfig(&tls.Config{
			// TODO: load cert from file
		}))
	}
	if config.Settings.HTTP.ReadTimeout > 0 {
		httpOpts = append(httpOpts, transhttp.Timeout(time.Duration(config.Settings.HTTP.ReadTimeout)*time.Second))
	}

	middlewares, err := LoadMiddlewares(config.Settings.ServiceName, config.Middleware)
	if err != nil {
		return err
	}
	if len(middlewares) > 0 {
		httpOpts = append(httpOpts, transhttp.Middleware(middlewares...))
	}

	if config.Middleware.Cors.Enabled {
		cors := transhttp.Filter(handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		),
		)
		httpOpts = append(httpOpts, cors)
	}

	srv := transhttp.NewServer(httpOpts...)

	r := srv.Route("/")
	r.GET("/healthz", func(c transhttp.Context) error {
		return c.Result(200, "ok")
	})

	opts = append(opts, kratos.Server(
		srv,
	))

	app := kratos.New(opts...)
	osg, err := opensergo.New(opensergo.WithEndpoint("locahost:9090"))
	if err != nil {
		return err
	}
	if err = osg.ReportMetadata(ctx, app); err != nil {
		return err
	}
	return app.Run()
}
