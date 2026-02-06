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
	"origadmin/application/admin/internal/data"
	identitybiz "origadmin/application/admin/internal/features/identity/biz"
	//identityclient "origadmin/application/admin/internal/features/identity/client"
	identitydal "origadmin/application/admin/internal/features/identity/dal"
	identityserver "origadmin/application/admin/internal/features/identity/server"
	identityservice "origadmin/application/admin/internal/features/identity/service"
	"origadmin/application/admin/internal/helpers/providers"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// Shared infrastructure providers
		providers.ProviderBackendSet,
		// Service-specific providers
		data.ProviderSet,

		// Auth feature module providers
		identitydal.ProviderSet,
		identitybiz.ProviderSet,
		identityservice.ProviderSet,
		identityserver.ProviderSet,

		NewApp,
	))
}
