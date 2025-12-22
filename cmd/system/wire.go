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
	"origadmin/application/admin/internal/data" // Added missing import for data package
	systembiz "origadmin/application/admin/internal/features/system/biz"
	systemdal "origadmin/application/admin/internal/features/system/dal"
	systemserver "origadmin/application/admin/internal/features/system/server"
	systemservice "origadmin/application/admin/internal/features/system/service"
)

// wireApp init kratos application.
func wireApp(r *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		//loader.ProviderSet, // Uncomment if loader.ProviderSet is needed
		data.ProviderSet,
		systemdal.ProviderSet,
		systembiz.ProviderSet,
		systemservice.ProviderSet,
		systemserver.ProviderSet,
		/* add your providers here */
		NewApp,
	))
}
