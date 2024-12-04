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
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/loader"
	commonbiz "origadmin/application/admin/internal/mods/common/biz"
	//commondal "origadmin/application/admin/internal/mods/common/dal"
	//commonserver "origadmin/application/admin/internal/mods/common/server"
	commonservice "origadmin/application/admin/internal/mods/common/service"
	systembiz "origadmin/application/admin/internal/mods/system/biz"
	systemdal "origadmin/application/admin/internal/mods/system/dal"
	systemserver "origadmin/application/admin/internal/mods/system/server"
	systemservice "origadmin/application/admin/internal/mods/system/service"
)

// buildInjectors init kratos application.
func buildInjectors(context.Context, *configs.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		loader.ProviderSet,
		//commondal.ProviderSet,
		commonbiz.ProviderSet,
		commonservice.ProviderSet,
		//commonserver.ProviderSet,
		systemdal.ProviderSet,
		systembiz.ProviderSet,
		systemservice.ProviderSet,
		systemserver.ProviderSet,
		/* add your providers here */
		NewApp))
}
