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
	"origadmin/application/admin/internal/loader"
	authservice "origadmin/application/admin/internal/mods/auth/service"
	systemservice "origadmin/application/admin/internal/mods/system/service"
)

// buildInjectors init kratos application.
func buildInjectors(r runtime.Runtime, bootstrap *configs.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(
		loader.ProviderSet,
		//agent.ProviderSet,
		systemservice.ProviderSet,
		authservice.ProviderSet,
		//basisserver.ProviderSet,
		NewApp))
}
