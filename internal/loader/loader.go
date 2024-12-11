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
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/registry"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/agent"
	systemserver "origadmin/application/admin/internal/mods/system/server"
)

var (
	ProviderSet = wire.NewSet(
		NewBasisConfig,
		NewRegistrar,
		NewAgentGINRegistrar,
		wire.Struct(new(InjectorServer), "*"),
		wire.Struct(new(InjectorClient), "*"),
	)
)

var (
	_ *gins.Server
	_ *http.Server
	_ *grpc.Server
)

type InjectorClient struct {
	Logger     log.Logger
	Bootstrap  *configs.Bootstrap
	Registrars []agent.GINRegistrar
	Server     *http.Server
	//Router      *gin.Engine
	//Servers []transport.Server
}

type InjectorServer struct {
	Logger    log.Logger
	Bootstrap *configs.Bootstrap
	Registrar registry.Registrar
	Servers   []transport.Server
}

func init() {
	runtime.RegisterConfigFunc("file", NewFileConfig)
	runtime.RegisterService("admin", service.DefaultServiceBuilder)
}

func NewBasisConfig(bootstrap *configs.Bootstrap) *configs.BasisConfig {
	//c := DefaultCaptcha()
	// todo Read from the configuration file
	return DefaultBasisConfig()
}

func NewAgentGINRegistrar(systemagent *systemserver.RegisterAgent) []agent.GINRegistrar {
	return []agent.GINRegistrar{
		systemagent,
	}
}
