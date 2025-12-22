/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package main is the main entry point for the gateway application.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/goexts/generic/cmp"

	"github.com/origadmin/runtime"
	middlewarev1 "github.com/origadmin/runtime/api/gen/go/config/middleware/v1"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/config"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/codec/toml"

	"origadmin/application/admin/internal/conf" // Updated import
	// _ "origadmin/application/admin/contrib/consul/config" // Removed
	// _ "origadmin/application/admin/contrib/consul/registry" // Removed
	// _ "origadmin/application/admin/contrib/database/drivers" // Removed
	_ "origadmin/application/admin/internal/data/entity/ent/runtime" // Updated import
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

func init() {
	encoding.RegisterCodec(toml.Codec)
	flags.SetServiceInfo(Name, Version)
}

// ResolvedBootstrap implements config.Resolver for the application's bootstrap configuration.
type ResolvedBootstrap struct {
	bootstrap *conf.Bootstrap
}

// FillServiceInfo populates service information into the bootstrap flags.
func (r *ResolvedBootstrap) FillServiceInfo(flags *bootstrap.Bootstrap) {
	core := r.bootstrap.GetServer().GetCore()
	name := cmp.Or(flags.ServiceName(), core.GetName())
	version := cmp.Or(flags.Version(), core.GetVersion())
	flags.SetServiceInfo(name, version)
}

// Discovery returns the discovery configuration.
func (r *ResolvedBootstrap) Discovery() *configv1.Discovery {
	log.NewHelper(log.GetLogger()).Infow("msg", "discovery config", "value", r.bootstrap.GetDiscovery())
	return r.bootstrap.GetDiscovery()
}

// Resolve scans the configuration into the bootstrap structure.
func (r *ResolvedBootstrap) Resolve(cfg config.KConfig) (config.Resolved, error) {
	if err := cfg.Scan(r.bootstrap); err != nil {
		return nil, err
	}
	return r, nil
}

// WithDecode is not implemented for ResolvedBootstrap.
func (r *ResolvedBootstrap) WithDecode(name string, v any, decode func([]byte, any) error) error {
	if decode == nil {
		return fmt.Errorf("decode function is nil")
	}
	return nil
}

// Value is not implemented for ResolvedBootstrap.
func (r *ResolvedBootstrap) Value(name string) (any, error) {
	return nil, fmt.Errorf("unknown config name: %s", name)
}

// Middleware returns the middleware configuration.
func (r *ResolvedBootstrap) Middleware() *middlewarev1.Middleware {
	return r.bootstrap.GetMiddleware()
}

// Services returns the service configurations.
func (r *ResolvedBootstrap) Services() []*configv1.Service {
	return r.bootstrap.GetServer().GetServices()
}

// Logger returns the logger configuration.
func (r *ResolvedBootstrap) Logger() *configv1.Logger {
	return r.bootstrap.GetLogger()
}

func main() {
	// Simplified flag parsing for demonstration, replace with actual flag parsing if needed
	// For now, hardcode debug mode for testing, or use os.Args to parse
	debug := false // Default to false
	for _, arg := range os.Args {
		if arg == "--debug" || arg == "-d" {
			debug = true
			break
		}
	}

	if debug {
		flags.SetEnv("debug")
		flags.SetConfigPath("resources/configs/bootstrap.toml") // Updated path
		flags.SetWorkDir(".")
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	ll := log.NewHelper(log.GetLogger())
	ll.Infof("bootstrap flags: %+v", flags)

	// Replicate loader.Bootstrap logic
	rb := &ResolvedBootstrap{
		bootstrap: &conf.Bootstrap{}, // Initialize with new conf.Bootstrap
	}
	r, err := runtime.Load(flags, runtime.WithResolver(rb), runtime.WithContext(context.Background())) // Use context.Background() as no cobra.Command context
	if err != nil {
		ll.Errorf("failed to load runtime: %v", err)
		os.Exit(1)
	}
	rb.FillServiceInfo(flags)
	r = r.WithLoggerAttrs(
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", flags.ServiceID(),
		"service.name", flags.ServiceName(),
		"service.version", flags.Version(),
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	app, clean, err := buildInjectors(r, rb.bootstrap) // Use rb.bootstrap
	if err != nil {
		ll.Errorf("failed to build injectors: %v", err)
		os.Exit(1)
	}
	defer clean()
	if err := app.Run(); err != nil {
		ll.Errorf("application run failed: %v", err)
		os.Exit(1)
	}
}

func NewApp(r runtime.Runtime, servers []transport.Server) *kratos.App {
	r = r.Client()
	return r.CreateApp(servers...)
}

func buildInjectors(r runtime.Runtime, bootstrap *conf.Bootstrap) (*kratos.App, func(), error) { // Updated type
	ll := log.NewHelper(r.Logger())
	if bootstrap.GetMode() == "cluster" {
		ll.Infof("start cluster mode")
		return buildRemoteInjectors(r, bootstrap)
	} else {
		ll.Infof("start local mode")
		return buildLocalInjectors(r, bootstrap)
	}
}
