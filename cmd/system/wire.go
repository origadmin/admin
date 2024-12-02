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

	"origadmin/application/admin/internal/loader"
	"origadmin/application/admin/internal/mods/system/biz"
	"origadmin/application/admin/internal/mods/system/dal"
	"origadmin/application/admin/internal/mods/system/server"
	"origadmin/application/admin/internal/mods/system/service"

	"origadmin/application/admin/internal/configs"
)

// buildInjectors init kratos application.
func buildInjectors(context.Context, *configs.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(loader.ProviderSet, server.ProviderServer, dal.ProviderSet, biz.ProviderSet, service.ProviderSet, NewApp))
}
