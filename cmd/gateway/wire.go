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
	"origadmin/application/admin/internal/gateway/client"
	"origadmin/application/admin/internal/gateway/server"
	"origadmin/application/admin/internal/gateway/service"
	"origadmin/application/admin/internal/helpers/providers"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// Gateway-specific providers, which includes common providers.
		providers.ProviderGatewaySet,

		// Service-specific providers
		server.ProviderSet,
		service.ProviderSet,
		client.ProviderSet,
		NewApp,
	))
}
