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
	"origadmin/application/admin/internal/pkg/token"
)

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// The injector function's parameter `app` is an implicit provider for *runtime.App.
		infraProviderSet,
		wire.FieldsOf(new(*conf.Config), "Bootstrap"),
		wire.FieldsOf(new(*confpb.Bootstrap), "Servers"),
		data.ProviderSet, // This was the missing piece
		dal.ProviderSet,
		biz.ProviderSet,
		token.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		NewApp,
	))
}
