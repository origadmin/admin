/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/configs"
)

var (
	// ProviderServer is server providers.
	ProviderServer = wire.NewSet(NewMenuServers)
	// ProviderAgent is agent providers.
	ProviderAgent = wire.NewSet(NewGinHTTPServer)
)

func NewMenuServers(bs *configs.Bootstrap, menus pb.MenuAPIServer, l log.Logger) []transport.Server {
	var servers []transport.Server
	if serv := NewGRPCServer(bs, menus, l); serv != nil {
		servers = append(servers, serv)
	}
	if serv := NewGINSServer(bs, menus, l); serv != nil {
		servers = append(servers, serv)
	}
	if serv := NewHTTPServer(bs, menus, l); serv != nil {
		servers = append(servers, serv)
	}
	return servers
}
