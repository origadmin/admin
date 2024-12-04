/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/api/v1/services/common"
	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/configs"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(bootstrap *configs.Bootstrap, l log.Logger) *service.GRPCServer {
	srv, err := runtime.NewGRPCServiceServer(bootstrap.GetService())
	if err != nil {
		panic(err)
	}
	return srv
}

func RegisterGRPCServer(srv *grpc.Server, menus pb.MenuAPIServer, login common.LoginAPIServer) {
	common.RegisterLoginAPIServer(srv, login)
	pb.RegisterMenuAPIServer(srv, menus)
}
