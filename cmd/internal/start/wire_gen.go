// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package start

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
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
func buildInjectors(contextContext context.Context, bootstrap *configs.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	registerAgent, err := server.NewSystemServerAgent(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	httpServer := agent.NewHTTPServer(bootstrap, logger)
	injectorClient := &loader.InjectorClient{
		Logger:      logger,
		Bootstrap:   bootstrap,
		SystemAgent: registerAgent,
		Server:      httpServer,
	}
	app := NewApp(contextContext, injectorClient)
	return app, func() {
	}, nil
}
