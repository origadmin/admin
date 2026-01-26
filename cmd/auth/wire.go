//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/data"
	authbiz "origadmin/application/admin/internal/features/auth/biz"
	authdal "origadmin/application/admin/internal/features/auth/dal"
	authserver "origadmin/application/admin/internal/features/auth/server"
	authservice "origadmin/application/admin/internal/features/auth/service"
	"origadmin/application/admin/internal/gateway/client"
	"origadmin/application/admin/internal/helpers/providers"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// General backend providers
		providers.ProviderBackendSet,

		// Data layer provider
		data.ProviderSet,

		// Gateway clients (for auth to call system service via gRPC)
		client.ProviderSet,

		// Auth feature module providers
		authdal.ProviderSet,
		authbiz.ProviderSet,
		authservice.ProviderSet,
		authserver.ProviderSet,

		NewApp,
	))
}
