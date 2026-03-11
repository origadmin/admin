//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dal"
	"origadmin/application/admin/internal/helpers/providers"
	"origadmin/application/admin/internal/jobs/initializer"
)

// wireApp init the initializer service.
func wireApp(rt *runtime.App, b *confpb.Bootstrap) (initializer.Initializer, func(), error) {
	panic(wire.Build(
		// Shared providers needed for data access and other common components.
		providers.ProviderCommonSet,
		data.ProviderSet,

		// Add providers from the 'system' feature, which are required by the seeder.
		dal.ProviderSet,
		biz.ProviderSet,

		// The main initializer provider set which includes all sub-initializers (NATS, Seeder, etc.)
		initializer.ProviderSet,

		// Note: rt (*runtime.App) and b (*confpb.Bootstrap) are automatically available
		// as providers because they are arguments to this function.
	))
}
