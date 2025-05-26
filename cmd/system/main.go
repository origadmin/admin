/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/log"

	"github.com/origadmin/toolkits/codec/toml"

	_ "origadmin/application/admin/contrib/consul/config"
	_ "origadmin/application/admin/contrib/consul/registry"
	_ "origadmin/application/admin/contrib/database"
	"origadmin/application/admin/internal/loader"
)

// go build -ldflags "-X main.Version=vx.y.z -X main.Name=origadmin.service.system.v1"
var (
	// Name is the Name of the compiled software.
	Name = "origadmin.service.system.v1"
	// Version is the Version of the compiled software.
	Version = "v1.0.0"
	// boot are the bootstrap boot.
	flags = bootstrap.New()
	// debug mode
	debug = false
	// configPath is the config path, default is config.toml
	configPath = ""
)

func init() {
	encoding.RegisterCodec(toml.Codec)
	flags.SetServiceInfo(Name, Version)
	flag.BoolVar(&debug, "debug", false, "set environment, eg: -debug")
	flag.StringVar(&configPath, "c", "config.toml", "config path, eg: -c config.toml")
}

func main() {
	flag.Parse()

	// the release mode, work dir sets to empty, use config path as work dir
	if debug {
		fmt.Println("debug mode")
		flags.SetEnv("debug")
		flags.SetConfigPath("resources/configs/config.toml")
		flags.SetWorkDir(".")
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	//r, err := runtime.Load(flags)
	//if err != nil {
	//	return
	//}
	//l := r.Logger(
	//	"ts", log.DefaultTimestamp,
	//	"caller", log.DefaultCaller,
	//	"service.id", flags.ServiceID(),
	//	"service.name", flags.ServiceName(),
	//	"service.version", flags.Version(),
	//	"trace.id", tracing.TraceID(),
	//	"span.id", tracing.SpanID(),
	//)
	//log.SetLogger(l)
	log.Infof("bootstrap flags: %+v\n", flags)
	if err := loader.Bootstrap(context.Background(), flags, buildInjectors); err != nil {
		log.Fatalf("failed to bootstrap: %s", err.Error())
		return
	}
}

// NewApp new app with runtime and injector
func NewApp(r runtime.Runtime, injector *loader.Injector) *kratos.App {
	return r.CreateApp(injector.Servers...)
}
