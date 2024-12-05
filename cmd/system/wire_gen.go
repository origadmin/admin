// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/loader"
	biz2 "origadmin/application/admin/internal/mods/common/biz"
	dal2 "origadmin/application/admin/internal/mods/common/dal"
	service2 "origadmin/application/admin/internal/mods/common/service"
	"origadmin/application/admin/internal/mods/system/biz"
	"origadmin/application/admin/internal/mods/system/dal"
	"origadmin/application/admin/internal/mods/system/server"
	"origadmin/application/admin/internal/mods/system/service"
)

import (
	_ "github.com/origadmin/contrib/consul/config"
	_ "github.com/origadmin/contrib/consul/registry"
	_ "github.com/origadmin/contrib/database"
)

// Injectors from wire.go:

// buildInjectors init kratos application.
func buildInjectors(contextContext context.Context, bootstrap *configs.Bootstrap, arg log.Logger) (*kratos.App, func(), error) {
	v, err := loader.NewRegistrar(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	data, cleanup, err := dal.NewData(bootstrap, arg)
	if err != nil {
		return nil, nil, err
	}
	menuRepo := dal.NewMenuRepo(data, arg)
	menuAPIClient := biz.NewMenusClient(menuRepo, arg)
	menuAPIServer := service.NewMenuAPIServer(menuAPIClient)
	roleRepo := dal.NewRoleRepo(data, arg)
	roleAPIClient := biz.NewRolesClient(roleRepo, arg)
	roleAPIServer := service.NewRoleAPIServer(roleAPIClient)
	userRepo := dal.NewUserRepo(data, arg)
	userAPIClient := biz.NewUsersClient(userRepo, arg)
	userAPIServer := service.NewUserAPIServer(userAPIClient)
	registerServer := &service.RegisterServer{
		Menu: menuAPIServer,
		Role: roleAPIServer,
		User: userAPIServer,
	}
	captcha := loader.NewCaptcha(bootstrap)
	loginRepo := dal2.NewLoginRepo(captcha, menuRepo, arg)
	loginAPIClient := biz2.NewLoginClient(loginRepo, arg)
	loginAPIServer := service2.NewLoginAPIServer(loginAPIClient)
	serviceRegisterServer := service2.RegisterServer{
		Login: loginAPIServer,
	}
	v2 := server.NewSystemServer(registerServer, serviceRegisterServer, bootstrap, arg)
	injectorServer := &loader.InjectorServer{
		Logger:    arg,
		Bootstrap: bootstrap,
		Registrar: v,
		Servers:   v2,
	}
	app := NewApp(contextContext, injectorServer)
	return app, func() {
		cleanup()
	}, nil
}
