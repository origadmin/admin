/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"syscall"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	_ "github.com/origadmin/contrib/consul/config"
	_ "github.com/origadmin/contrib/consul/registry"
	_ "github.com/origadmin/contrib/database"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/log"
	kslog "github.com/origadmin/slog-kratos"

	"origadmin/application/admin/internal/loader"
)

// go build -ldflags "-X main.Version=vx.y.z -X main.Name=origadmin.service.system.v1"
var (
	// Name is the Name of the compiled software.
	Name = "origadmin.service.system.v1"
	// Version is the Version of the compiled software.
	Version = "v1.0.0"
	// boot are the bootstrap boot.
	flags = bootstrap.DefaultBootstrap()
	// debug mode
	debug = false
)

func init() {
	flags.Env = "release"
	flags.SetFlags(Name, Version)
	flag.BoolVar(&debug, "debug", false, "set environment, eg: -debug")
	flag.StringVar(&flags.ConfigPath, "c", "config.toml", "config path, eg: -c config.toml")
}

func main() {
	flag.Parse()

	// the release mode, work dir sets to empty, use config path as work dir
	if debug {
		fmt.Println("debug mode")
		flags.Env = "debug"
		flags.WorkDir = "resources/configs"
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	l := log.With(kslog.NewLogger(),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", flags.ID(),
		"service.name", flags.ServiceName(),
		"service.version", flags.Version(),
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(l)
	log.Infof("bootstrap flags: %+v\n", flags)
	bs, err := loader.LoadBootstrap(flags)
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
		return
	}
	//v, err := validate.NewValidateV2()
	//if err != nil {
	//	log.Fatalf("failed to new validate: %s", err.Error())
	//}
	// v1 used method Validate to check the config
	if err := bs.Validate(); err != nil {
		log.Fatalf("failed to validate config: %s", err.Error())
	}

	if err := loader.InitSetup(bs); err != nil {
		log.Fatalf("failed to init setup: %s", err.Error())
		return
	}

	//log.Infof("bootstrap config: %+v\n", loader.PrintString(bs))
	ctx := context.Background()
	//info to ctx
	app, cleanup, err := buildInjectors(ctx, bs, l)
	if err != nil {
		log.Fatalf("failed to build injector: %s", err.Error())
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Fatalf("app stopped with error: %s", err.Error())
	}
}

func NewApp(ctx context.Context, injector *loader.InjectorServer) *kratos.App {
	opts := []kratos.Option{
		kratos.ID(flags.ServiceID()),
		kratos.Name(flags.ServiceName()),
		kratos.Version(flags.Version()),
		kratos.Metadata(flags.Metadata()),
		kratos.Context(ctx),
		kratos.Signal(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
		kratos.Logger(injector.Logger),
		kratos.Server(injector.Servers...),
	}
	if injector.Registrar != nil {
		opts = append(opts, kratos.Registrar(injector.Registrar))
	}

	return kratos.New(opts...)
}
