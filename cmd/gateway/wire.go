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
	"origadmin/application/admin/internal/loader"
	authbiz "origadmin/application/admin/internal/features/auth/biz"       // Corrected import path
	authdal "origadmin/application/admin/internal/features/auth/dal"       // Corrected import path
	authservice "origadmin/application/admin/internal/features/auth/service" // Corrected import path
	"origadmin/application/admin/internal/features/gateway"                // Corrected import path
	systembiz "origadmin/application/admin/internal/features/system/biz"     // Corrected import path
	systemdal "origadmin/application/admin/internal/features/system/dal"     // Corrected import path
	systemservice "origadmin/application/admin/internal/features/system/service" // Corrected import path

	"origadmin/application/admin/internal/data"
)

// buildInjectors init kratos application.
func buildLocalInjectors(r runtime.Runtime, bootstrap *configs.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(
		loader.ProviderSet,
		//agent.ProviderSet,
		data.ProviderSet,
		systemdal.ProviderSet,
		systembiz.ProviderSet,
		systemservice.LocalProviderSet,
		//systemserver.ProviderSet,
		authdal.ProviderSet,
		authbiz.ProviderSet,
		authservice.LocalProviderSet,
		gateway.ProviderSet,
		NewApp))
}

func buildRemoteInjectors(r runtime.Runtime, bootstrap *configs.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(
		loader.ProviderSet,
		systemservice.RemoteProviderSet,
		authservice.RemoteProviderSet,
		gateway.ProviderSet,
		NewApp,
	))
}
