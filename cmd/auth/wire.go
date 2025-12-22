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

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/data"
	authbiz "origadmin/application/admin/internal/features/auth/biz"       // Corrected import path
	authdal "origadmin/application/admin/internal/features/auth/dal"       // Corrected import path
	authserver "origadmin/application/admin/internal/features/auth/server" // Corrected import path
	authservice "origadmin/application/admin/internal/features/auth/service" // Corrected import path
)

// buildInjectors init kratos application.
func buildInjectors(r runtime.Runtime, bootstrap *configs.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(
		//loader.ProviderSet,
		data.ProviderSet,
		//authdal.ProviderSet,
		//basisbiz.ProviderSet,
		//basisservice.ProviderSet,
		//basisserver.ProviderSet,
		authdal.ProviderSet,
		authbiz.ProviderSet,
		authservice.ProviderSet,
		authserver.ProviderSet,
		/* add your providers here */
		NewApp,
	))
}
