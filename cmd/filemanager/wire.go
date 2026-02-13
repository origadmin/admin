//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	objclient "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/filemanager/biz"
	"origadmin/application/admin/internal/features/filemanager/dal"
	"origadmin/application/admin/internal/features/filemanager/server"
	"origadmin/application/admin/internal/features/filemanager/service"
	"origadmin/application/admin/internal/helpers/grpcclient"
	"origadmin/application/admin/internal/helpers/providers"
)

// NewObjectStoreServiceClient creates a new ObjectStoreService client using service discovery.
func NewObjectStoreServiceClient(
	app *runtime.App,
	bootstrap *conf.Config,
	middlewareProvider container.ClientMiddlewareProvider,
) (objclient.ObjectStoreServiceClient, func(), error) {
	conn, err := grpcclient.NewConn(app, bootstrap, "objectstore", middlewareProvider)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		conn.Close()
	}
	return objclient.NewObjectStoreServiceClient(conn), cleanup, nil
}

// wireApp init kratos application.
func wireApp(app *runtime.App, c *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// Shared infrastructure providers
		providers.ProviderBackendSet,
		data.ProviderSet, // Provides database connection

		// Clients
		NewObjectStoreServiceClient,

		// Server
		server.ProviderSet,

		// FileManager Feature
		biz.ProviderSet,
		dal.ProviderSet,
		service.ProviderSet,

		NewApp,
	))
}
