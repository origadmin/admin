//go:build wireinject
// +build wireinject

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dal"
	"origadmin/application/admin/internal/helpers/providers"
	"origadmin/application/admin/internal/tasks/seeder"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *conf.Config) (*seeder.Seeder, func(), error) {
	panic(wire.Build(
		// Shared infrastructure providers
		providers.ProviderSet,

		// Data and Biz layers are needed for seeding
		data.ProviderSet,
		dal.ProviderSet,
		biz.ProviderSet,

		// The Seeder itself
		seeder.ProviderSet,
	))
}
