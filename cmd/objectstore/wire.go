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
	"origadmin/application/admin/internal/features/objectstore/biz"
	"origadmin/application/admin/internal/features/objectstore/dal"
	"origadmin/application/admin/internal/features/objectstore/server"
	"origadmin/application/admin/internal/features/objectstore/service"
	"origadmin/application/admin/internal/helpers/providers"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, b *confpb.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(
		// Common providers
		providers.ProviderBackendSet,

		// Feature-specific providers
		dal.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,

		NewApp,
	))
}
