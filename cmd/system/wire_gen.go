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
func buildInjectors(contextContext context.Context, bootstrap *configs.Bootstrap, arg log.KLogger) (*kratos.App, func(), error) {
	v, err := loader.NewRegistrar(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	data, cleanup, err := dal.NewData(bootstrap, arg)
	if err != nil {
		return nil, nil, err
	}
	resourceRepo := dal.NewResourceRepo(data, arg)
	resourceServiceClient := biz.NewResourceServiceClient(resourceRepo, arg)
	resourceServiceServer := service.NewResourceServiceServerPB(resourceServiceClient)
	roleRepo := dal.NewRoleRepo(data, arg)
	roleServiceClient := biz.NewRoleServiceClient(roleRepo, arg)
	roleServiceServer := service.NewRoleServiceServerPB(roleServiceClient)
	userRepo := dal.NewUserRepo(data, arg)
	userServiceClient := biz.NewUserServiceClient(userRepo, arg)
	userServiceServer := service.NewUserServiceServerPB(userServiceClient)
	authRepo := dal.NewAuthRepo(data, arg)
	authServiceClient := biz.NewAuthServiceClient(authRepo, arg)
	authServiceServer := service.NewAuthServiceServerPB(authServiceClient)
	basisConfig := loader.NewBasisConfig(bootstrap)
	tokenizer, err := loader.NewTokenizer(bootstrap)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	refreshTokenizer := dal.RefreshTokenizer(tokenizer)
	loginData := &dal.LoginData{
		BasisConfig: basisConfig,
		Tokenizer:   refreshTokenizer,
		Resource:    resourceRepo,
		Role:        roleRepo,
		User:        userRepo,
	}
	loginRepo := dal.NewLoginRepo(loginData, arg)
	loginServiceClient := biz.NewLoginServiceClient(loginRepo, arg)
	loginServiceServer := service.NewLoginServiceServerPB(loginServiceClient)
	personalRepo := dal.NewPersonalRepo(data, arg)
	personalServiceClient := biz.NewPersonalServiceClient(personalRepo, arg)
	personalServiceServer := service.NewPersonalServiceServerPB(personalServiceClient)
	registerServer := &service.RegisterServer{
		Resource: resourceServiceServer,
		Role:     roleServiceServer,
		User:     userServiceServer,
		Auth:     authServiceServer,
		Login:    loginServiceServer,
		Personal: personalServiceServer,
	}
	v2 := server.NewRegisterServer(registerServer)
	v3 := server.NewSystemServer(bootstrap, v2, arg)
	injectorServer := &loader.InjectorServer{
		Logger:    arg,
		Bootstrap: bootstrap,
		Registrar: v,
		Servers:   v3,
	}
	app := NewApp(contextContext, injectorServer)
	return app, func() {
		cleanup()
	}, nil
}
