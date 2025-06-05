/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime"
	configv1 "github.com/origadmin/runtime/api/gen/go/config/v1"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/registry"

	"origadmin/application/admin/internal/configs"
)

// AppOptions 包含微服务核心配置
type AppOptions struct {
	ID       string
	Name     string
	Version  string
	Metadata map[string]string
	Logger   log.KLogger
	Server   transport.Server
}

var (
	ProviderSet = wire.NewSet(
		NewRegistrar,
		NewProxyServer,
		wire.Struct(new(Injector), "*"),
		wire.Struct(new(InjectorClient), "*"),
	)
)

var (
	_ *gins.Server
	_ *http.Server
	_ *grpc.Server
)

type Loader interface {
	SetupEnv() error
}

type InjectorClient struct {
	Server *http.Server
}

type Injector struct {
	Registrar registry.KRegistrar
	Servers   []transport.Server
}

func init() {
	runtime.RegisterConfigFunc("file", NewFileConfig)
}

type loader struct {
	flags *bootstrap.Bootstrap
	cfg   *configv1.SourceConfig
}

func (l loader) SetupEnv() error {
	if len(l.cfg.EnvPrefixes) > 0 {
		SetupEnv(l.cfg.EnvArgs, l.cfg.EnvPrefixes[0])
	}
	return nil
}

func (l loader) Bootstrap() (*configs.Bootstrap, error) {
	return LoadBootstrap(l.cfg)
}

func NewLoader(bs *bootstrap.Bootstrap) (Loader, error) {
	load := &loader{
		flags: bs,
	}
	return load, nil
}

func MockHttpServer() *http.Server {
	return http.NewServer()
}
