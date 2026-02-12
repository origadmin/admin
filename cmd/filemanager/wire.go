//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	objclient "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/filemanager/biz"
	"origadmin/application/admin/internal/features/filemanager/dal"
	"origadmin/application/admin/internal/features/filemanager/server"
	"origadmin/application/admin/internal/features/filemanager/service"
	"origadmin/application/admin/internal/helpers/providers"
)

// NewObjectStoreServiceClient creates a new ObjectStoreService client.
func NewObjectStoreServiceClient(c *conf.Config) (objclient.ObjectStoreServiceClient, func(), error) {
	// TODO: Get address from config 'c.Clients.Objectstore.Grpc.Addr'
	// For now, we hardcode the address.
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(":9001"), // Assuming objectstore runs on port 9001
	)
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
