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
	"google.golang.org/grpc"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	systemclient "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/data"
	identitybiz "origadmin/application/admin/internal/features/identity/biz"
	//identityclient "origadmin/application/admin/internal/features/identity/client"
	identitydal "origadmin/application/admin/internal/features/identity/dal"
	identityserver "origadmin/application/admin/internal/features/identity/server"
	identityservice "origadmin/application/admin/internal/features/identity/service"
	"origadmin/application/admin/internal/helpers/grpcclient"
	"origadmin/application/admin/internal/helpers/providers"
)

// NewSystemServiceClient creates a new SystemService client using service discovery.
func NewSystemServiceClient(
	app *runtime.App,
	bootstrap *conf.Config,
	middlewareProvider container.ClientMiddlewareProvider,
) (systemclient.UserServiceClient, func(), error) {
	conn, err := grpcclient.NewConn(app, bootstrap, "system", middlewareProvider)
	if err != nil {
		return nil, nil, err
	}
	grpcConn := conn.(*grpc.ClientConn)
	cleanup := func() {
		grpcConn.Close()
	}
	return systemclient.NewUserServiceClient(grpcConn), cleanup, nil
}

// wireApp init kratos application.
func wireApp(app *runtime.App, bootstrap *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		// Shared infrastructure providers
		providers.ProviderBackendSet,
		// Service-specific providers
		data.ProviderSet,

		// Clients
		NewSystemServiceClient,

		// Auth feature module providers
		identitydal.ProviderSet,
		identitybiz.ProviderSet,
		identityservice.ProviderSet,
		identityserver.ProviderSet,

		NewApp,
	))
}
