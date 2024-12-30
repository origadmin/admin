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
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/configs"
)

var (
	ProviderSet = wire.NewSet(
		NewBasisConfig,
		NewRegistrar,
		NewTokenizer,
		NewAuthorizer,
		NewAuthenticator,
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
	Logger    log.KLogger
	Bootstrap *configs.Bootstrap
	Server    *http.Server
}

type InjectorServer struct {
	Logger    log.KLogger
	Bootstrap *configs.Bootstrap
	Registrar registry.KRegistrar
	Servers   []transport.Server
}

func init() {
	runtime.RegisterConfigFunc("file", NewFileConfig)
	runtime.RegisterService("ORIGADMIN_SERVICE", service.DefaultServiceBuilder)
}

func NewAuthenticator(bootstrap *configs.Bootstrap) (security.Authenticator, error) {
	return securityx.NewAuthenticator(bootstrap)
}

func NewTokenizer(bootstrap *configs.Bootstrap) (security.Tokenizer, error) {
	authenticator, err := securityx.NewAuthenticator(bootstrap)
	if err != nil {
		return nil, err
	}
	return securityx.NewTokenizer(authenticator), nil
}

func NewAuthorizer(bootstrap *configs.Bootstrap) (security.Authorizer, error) {
	return securityx.NewAuthorizer(bootstrap)
}

func NewBasisConfig(bootstrap *configs.Bootstrap) *configs.BasisConfig {
	//c := DefaultCaptcha()
	// todo Read from the configuration file
	return DefaultBasisConfig()
}
