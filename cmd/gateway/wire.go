//go:build wireinject
// +build wireinject

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/gateway/client"
	"origadmin/application/admin/internal/gateway/server"
	"origadmin/application/admin/internal/gateway/service"
)

// wireApp init kratos application.
func wireApp(bootstrap *confpb.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, client.ProviderSet, service.ProviderSet, newApp))
}
