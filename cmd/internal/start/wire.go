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
	"github.com/google/wire"
	"github.com/origadmin/runtime/log"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/loader"
	"origadmin/application/admin/internal/mods/agent"
	"origadmin/application/admin/internal/mods/system/server"
)

// buildInjectors init kratos application.
func buildInjectors(context.Context, *configs.Bootstrap, log.KLogger) (*kratos.App, func(), error) {
	panic(wire.Build(
		loader.ProviderSet,
		agent.ProviderSet,
		server.ProviderSet,
		//basisserver.ProviderSet,
		NewApp))
}
