/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package start is the start command for the application.
package start

import (
	"context"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/origadmin/contrib/consul/config"
	_ "github.com/origadmin/contrib/consul/registry"
	_ "github.com/origadmin/contrib/database"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/registry"
	"github.com/spf13/cobra"

	"origadmin/application/admin/internal/loader"
)

const (
	startRandom  = `random`
	startWorkDir = `workdir`
	startConfig  = `config`
	startStatic  = `static`
	startDaemon  = `daemon`
	startDebug   = `debug`
)

var (
	// Name is the name of the compiled software.
	Name = "origadmin.server.v1.admin"
	// Version is the Version of the compiled software.
	Version = "v1.0.0"
	// flags are the bootstrap flags.
	flags = bootstrap.New()
)

var cmd = &cobra.Command{
	Use:   "start",
	Short: "start the server",
	RunE:  startCommandRun,
}

func init() {
	flags.SetServiceInfo(Name, Version)
}

// Cmd The function defines a CLI command to start a server with various flags and options, including the
// ability to run as a daemon.
func Cmd() *cobra.Command {
	cmd.Flags().BoolP(startRandom, "r", false, "start with random password")
	cmd.Flags().StringP(startConfig, "c", "bootstrap.toml",
		"runtime configuration files or directory (relative to workdir, multiple separated by commas)")
	cmd.Flags().StringP(startStatic, "s", "", "static files directory")
	cmd.Flags().BoolP(startDebug, "d", false, "set debug mode, eg: --debug")
	cmd.Flags().Bool(startDaemon, false, "run as a daemon")
	return cmd
}

func startCommandRun(cmd *cobra.Command, args []string) error {
	r, err := runtime.Load(flags, func(options *runtime.Options) {
		// Set your runtime options.
	})
	if err != nil {
		return err
	}

	var registrar registry.KRegistrar
	if flags.IsMainService() {
		registrar, _ = registry.NewConsulRegistrar()
	}

	buildInjectors()

	r.CreateApp(cmd.Context())

	// 组合使用配置和服务
	appInstance := loader.NewApp(cmd.Context(), loader.AppOptions{
		Name:    bs.ServiceName,
		Version: flags.Version(),
		Server:  grpcServer,
	})
}

func NewApp(ctx context.Context, injector *loader.InjectorClient) *kratos.App {
	opts := []kratos.Option{
		kratos.ID(flags.ServiceID()),
		kratos.Name(flags.ServiceName()),
		kratos.Version(flags.Version()),
		kratos.Metadata(map[string]string{}),
		kratos.Context(ctx),
		kratos.Signal(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
		kratos.Logger(injector.Logger),
		kratos.Server(injector.Server),
	}
	mux := gwruntime.NewServeMux()
	srv := transhttp.NewServer()
	srv.Handler = mux
	kratos.Server(srv)

	if flags.Env() == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infow("msg", "GIN route", "method", httpMethod, "path", absolutePath, "operation", handlerName, "handlers", nuHandlers)
	}

	return kratos.New(opts...)
}
