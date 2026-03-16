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
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dal"
	"origadmin/application/admin/internal/features/system/server"
	"origadmin/application/admin/internal/features/system/service"
	"origadmin/application/admin/internal/helpers/providers"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *confpb.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(
		// Shared infrastructure providers
		providers.ProviderBackendSet,

		dal.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,

		// Bootstrap options provider
		NewBootstrapOptions,

		NewApp,
	))
}
