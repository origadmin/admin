//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/auth/biz"
	"origadmin/application/admin/internal/features/auth/dal"
	"origadmin/application/admin/internal/features/auth/server"
	"origadmin/application/admin/internal/features/auth/service"
	"origadmin/application/admin/internal/helpers/providers"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// Shared infrastructure providers
		providers.ProviderSet,

		// Instructions for wire to extract nested configs
		wire.FieldsOf(new(*conf.Config), "Bootstrap"),
		wire.FieldsOf(new(*confpb.Bootstrap), "Servers"),
		wire.FieldsOf(new(*confpb.Bootstrap), "Captcha"),

		// Service-specific providers
		data.ProviderSet,
		dal.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		NewApp,
	))
}
