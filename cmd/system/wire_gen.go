// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/loader"
	"origadmin/application/admin/internal/mods/system/biz"
	"origadmin/application/admin/internal/mods/system/dal"
	"origadmin/application/admin/internal/mods/system/server"
	"origadmin/application/admin/internal/mods/system/service"
)

import (
	_ "github.com/origadmin/contrib/consul/config"
	_ "github.com/origadmin/contrib/consul/registry"
)

// Injectors from wire.go:

// buildInjectors init kratos application.
func buildInjectors(contextContext context.Context, bootstrap *configs.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	registrar, err := loader.NewRegistrar(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	data, cleanup, err := dal.NewData(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	menuRepo := dal.NewMenuDal(data, logger)
	menuAPIClient := biz.NewMenusClient(menuRepo, logger)
	menuAPIServer := service.NewMenuAPIServer(menuAPIClient)
	grpcServer := server.NewGRPCServer(bootstrap, menuAPIServer, logger)
	httpServer := server.NewHTTPServer(bootstrap, menuAPIServer, logger)
	injectorServer := &loader.InjectorServer{
		Bootstrap:  bootstrap,
		Logger:     logger,
		Registrar:  registrar,
		ServerGRPC: grpcServer,
		ServerHTTP: httpServer,
	}
	app := NewApp(contextContext, injectorServer)
	return app, func() {
		cleanup()
	}, nil
}
