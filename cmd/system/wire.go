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
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dal"
	"origadmin/application/admin/internal/features/system/server"
	"origadmin/application/admin/internal/features/system/service"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// The injector function's parameter `app` is an implicit provider for *runtime.App.
		infraProviderSet,
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
