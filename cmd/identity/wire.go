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
	systemclient "origadmin/application/admin/api/v1/services/system"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data"
	identitybiz "origadmin/application/admin/internal/features/identity/biz"
	identitydal "origadmin/application/admin/internal/features/identity/dal"
	identityserver "origadmin/application/admin/internal/features/identity/server"
	identityservice "origadmin/application/admin/internal/features/identity/service"
	"origadmin/application/admin/internal/helpers/grpcclient"
	"origadmin/application/admin/internal/helpers/providers"
)

// NewSystemServiceClient creates a new SystemService client using service discovery.
func NewSystemServiceClient(
	app *runtime.App,
	bootstrap *confpb.Bootstrap,
) (systemclient.UserServiceClient, func(), error) {
	conn, err := grpcclient.NewConn(app, bootstrap, "system")
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		conn.Close()
	}
	return systemclient.NewUserServiceClient(conn), cleanup, nil
}

// wireApp init kratos application.
func wireApp(app *runtime.App, b *confpb.Bootstrap) (*kratos.App, func(), error) {
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
