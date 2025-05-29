/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/goexts/generic/cmp"
	"github.com/origadmin/runtime"
	configv1 "github.com/origadmin/runtime/api/gen/go/config/v1"
	middlewarev1 "github.com/origadmin/runtime/api/gen/go/middleware/v1"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/config"
	"github.com/origadmin/runtime/log"

	"origadmin/application/admin/internal/configs"
)

type NewApp func(runtime.Runtime, *configs.Bootstrap) (*kratos.App, func(), error)

func Resolve(config config.KConfig) (config.Resolved, error) {
	var rb ResolvedBootstrap
	if err := config.Load(); err != nil {
		return nil, err
	}
	if err := config.Scan(&rb.bootstrap); err != nil {
		return nil, err
	}
	return &rb, nil
}

type BootstrapConfig func(config config.KConfig) (config.Resolved, error)

func (b BootstrapConfig) Resolve(config config.KConfig) (config.Resolved, error) {
	return b(config)
}

type ResolvedBootstrap struct {
	bootstrap configs.Bootstrap
}

func (r *ResolvedBootstrap) FillServiceInfo(flags *bootstrap.Bootstrap) {
	core := r.bootstrap.GetServer().GetCore()
	name := cmp.Or(flags.ServiceName(), core.GetName())
	version := cmp.Or(flags.Version(), core.GetVersion())
	flags.SetServiceInfo(name, version)
}

func (r *ResolvedBootstrap) Discovery() *configv1.Discovery {
	log.NewHelper(log.GetLogger()).Infow("msg", "discovery config", "value", r.bootstrap.GetDiscovery())
	return r.bootstrap.GetDiscovery()
}

func (r *ResolvedBootstrap) Resolve(config config.KConfig) (config.Resolved, error) {
	if err := config.Scan(&r.bootstrap); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *ResolvedBootstrap) WithDecode(name string, v any, decode func([]byte, any) error) error {
	if decode == nil {
		return fmt.Errorf("decode function is nil")
	}
	return nil
}

func (r *ResolvedBootstrap) Value(name string) (any, error) {
	switch name {

	default:
		return nil, fmt.Errorf("unknown config name: %s", name)
	}

}

func (r *ResolvedBootstrap) Middleware() *middlewarev1.Middleware {
	return r.bootstrap.GetMiddleware()
}

func (r *ResolvedBootstrap) Services() []*configv1.Service {
	return r.bootstrap.GetServer().GetServices()
}

func (r *ResolvedBootstrap) Logger() *configv1.Logger {
	return r.bootstrap.GetLogger()
}

func Bootstrap(ctx context.Context, flags *bootstrap.Bootstrap, newApp NewApp) error {
	var rb ResolvedBootstrap
	r, err := runtime.Load(flags, runtime.WithResolver(&rb), runtime.WithContext(ctx))
	if err != nil {
		return err
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
	app, clean, err := newApp(r, &rb.bootstrap)
	if err != nil {
		return err
	}
	defer clean()
	if err := app.Run(); err != nil {
		return err
	}
	return nil
}
