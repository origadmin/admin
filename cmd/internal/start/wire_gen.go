// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package start

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/loader"
	"origadmin/application/admin/internal/mods/agent"
	"origadmin/application/admin/internal/mods/system/server"
)

import (
	_ "github.com/origadmin/contrib/consul/config"
	_ "github.com/origadmin/contrib/consul/registry"
	_ "github.com/origadmin/contrib/database"
)

// Injectors from wire.go:

// buildInjectors init kratos application.
func buildInjectors(contextContext context.Context, bootstrap *configs.Bootstrap, arg log.KLogger) (*kratos.App, func(), error) {
	registerAgent, err := server.NewSystemServerAgent(bootstrap, arg)
	if err != nil {
		return nil, nil, err
	}
	v := loader.NewAgentHTTPRegistrar(registerAgent)
	authenticator, err := loader.NewAuthenticator(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	authorizer, err := loader.NewAuthorizer(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	httpServer := agent.NewHTTPServer(bootstrap, v, authenticator, authorizer, arg)
	injectorClient := &loader.InjectorClient{
		Logger:     arg,
		Bootstrap:  bootstrap,
		Registrars: v,
		Server:     httpServer,
	}
	app := NewApp(contextContext, injectorClient)
	return app, func() {
	}, nil
}
