//go:build wireinject
// +build wireinject

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/types"

	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dal"
	"origadmin/application/admin/internal/features/system/server"
	"origadmin/application/admin/internal/features/system/service"
)

func provideHasher() (hash.Crypto, error) {
	// Using a default cost for bcrypt. In a real application, this might come from config.
	return hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(bcrypt.DefaultCost))
}

func provideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// The injector function's parameter `app` is an implicit provider for *runtime.App.
		provideLogger,
		provideHasher,
		wire.FieldsOf(new(*conf.Config), "Bootstrap"),
		wire.FieldsOf(new(*confpb.Bootstrap), "Servers"),
		data.ProviderSet,
		dal.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		NewApp,
	))
}
