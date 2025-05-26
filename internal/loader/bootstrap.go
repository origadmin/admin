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
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/config"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	middlewarev1 "github.com/origadmin/runtime/gen/go/middleware/v1"
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

func (r *ResolvedBootstrap) Resolve(config config.KConfig) (config.Resolved, error) {
	var unknown map[string]any
	if err := config.Scan(&unknown); err != nil {
		return nil, err
	}
	log.NewHelper(log.DefaultLogger).Infof("bootstrap: %+v", unknown)
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

func (r *ResolvedBootstrap) Registry() *configv1.Registry {
	return r.bootstrap.GetRegistry()
}

func (r *ResolvedBootstrap) Middleware() *middlewarev1.Middleware {
	return r.bootstrap.GetMiddleware()
}

func (r *ResolvedBootstrap) Service() *configv1.Service {
	panic("unimplemented")
	//return r.bootstrap.GetService()
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
	r = r.WithLoggerAttrs(
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", flags.ServiceID(),
		"service.name", flags.ServiceName(),
		"service.version", flags.Version(),
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	help := log.NewHelper(r.Logger())
	help.Infof("bootstrap: %+v", &rb.bootstrap)
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
