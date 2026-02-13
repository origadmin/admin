//go:build wireinject
// +build wireinject

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/objectstore/biz"
	"origadmin/application/admin/internal/features/objectstore/dal"
	"origadmin/application/admin/internal/features/objectstore/server"
	"origadmin/application/admin/internal/features/objectstore/service"
	"origadmin/application/admin/internal/helpers/providers"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, c *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// Shared infrastructure providers
		providers.ProviderBackendSet,
		data.ProviderSet, // Provides database connection

		// Server
		server.ProviderSet,

		// ObjectStore Feature
		biz.ProviderSet,
		dal.ProviderSet,
		service.ProviderSet,

		NewApp,
	))
}
