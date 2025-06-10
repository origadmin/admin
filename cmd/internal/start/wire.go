//go:build wireinject
// +build wireinject

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// The build tag makes sure the stub is not built in the final build.
package start

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/origadmin/runtime"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/loader"
	authbiz "origadmin/application/admin/internal/mods/auth/biz"
	authdal "origadmin/application/admin/internal/mods/auth/dal"
	authservice "origadmin/application/admin/internal/mods/auth/service"
	systembiz "origadmin/application/admin/internal/mods/system/biz"
	systemdal "origadmin/application/admin/internal/mods/system/dal"
	systemservice "origadmin/application/admin/internal/mods/system/service"
)

// buildInjectors init kratos application.
func buildLocalInjectors(r runtime.Runtime, bootstrap *configs.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(
		loader.ProviderSet,
		//agent.ProviderSet,
		data.ProviderSet,
		systemdal.ProviderSet,
		systembiz.ProviderSet,
		systemservice.ProviderSet,
		//systemserver.ProviderSet,
		authdal.ProviderSet,
		authbiz.ProviderSet,
		authservice.ProviderSet,
		NewApp))
}

func buildRemoteInjectors(r runtime.Runtime, bootstrap *configs.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(
		loader.ProviderSet,
		systemservice.RemoteProviderSet,
		authservice.RemoteProviderSet,
		NewApp,
	))
}
