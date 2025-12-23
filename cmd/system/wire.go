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
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/data"
	"origadmin/application/admin/internal/features/system/server"
	"origadmin/application/admin/internal/features/system/service"
)

func provideLogger(r *runtime.App) log.Logger {
	return r.Logger()
}

func provideHasher() (hash.Crypto, error) {
	return hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(10))
}

// wireApp init kratos application.
func wireApp(r *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		provideLogger,
		provideHasher,
		wire.FieldsOf(new(*runtime.App), "AppInfo"),
		wire.FieldsOf(new(*conf.Config), "Bootstrap"),
		wire.FieldsOf(new(*confpb.Bootstrap), "Servers"),
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		NewApp,
	))
}
