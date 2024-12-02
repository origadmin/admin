//go:build wireinject
// +build wireinject

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// The build tag makes sure the stub is not built in the final build.
package start

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/system/server"
)

// buildInjectors init kratos application.
func buildInjectors(context.Context, *loader.Config, *configs.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		loader.ProviderSet,
		server.ProviderAgent,
		//helloworlddal.ProviderSet,
		//helloworldbiz.ProviderSet,
		//helloworldservice.ProviderSet,
		//secondworlddal.ProviderSet,
		//secondworldbiz.ProviderSet,
		//secondworldservice.ProviderSet,
		NewApp))
}
