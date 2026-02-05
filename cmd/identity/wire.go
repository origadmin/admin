//go:build wireinject
// +build wireinject

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
		// General backend providers
		providers.ProviderBackendSet,

		// Data layer provider
		data.ProviderSet,

		// Client provider for gRPC calls
		//identityclient.ProviderSet,

		// Auth feature module providers
		identitydal.ProviderSet,
		identitybiz.ProviderSet,
		identityservice.ProviderSet,
		identityserver.ProviderSet,

		// Bootstrap options provider
		NewBootstrapOptions,

		NewApp,
	))
}
