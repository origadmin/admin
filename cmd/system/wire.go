//go:build wireinject
// +build wireinject

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/origadmin/runtime/log"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/loader"
	basisbiz "origadmin/application/admin/internal/mods/basis/biz"
	basisdal "origadmin/application/admin/internal/mods/basis/dal"
	basisservice "origadmin/application/admin/internal/mods/basis/service"
	systembiz "origadmin/application/admin/internal/mods/system/biz"
	systemdal "origadmin/application/admin/internal/mods/system/dal"
	systemserver "origadmin/application/admin/internal/mods/system/server"
	systemservice "origadmin/application/admin/internal/mods/system/service"
)

// buildInjectors init kratos application.
func buildInjectors(context.Context, *configs.Bootstrap, log.KLogger) (*kratos.App, func(), error) {
	panic(wire.Build(
		loader.ProviderSet,
		basisdal.ProviderSet,
		basisbiz.ProviderSet,
		basisservice.ProviderSet,
		//basisserver.ProviderSet,
		systemdal.ProviderSet,
		systembiz.ProviderSet,
		systemservice.ProviderSet,
		systemserver.ProviderSet,
		/* add your providers here */
		NewApp))
}
