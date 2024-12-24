/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/internal/configs"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(bootstrap *configs.Bootstrap, l log.KLogger, ss ...service.OptionSetting) *service.GRPCServer {
	srv, err := runtime.NewGRPCServiceServer(bootstrap.GetService(), ss...)
	if err != nil {
		panic(err)
	}
	return srv
}

//func RegisterGRPCServer(srv *grpc.Server, menus pb.MenuAPIServer, login basis.LoginAPIServer) {
//	basis.RegisterLoginAPIServer(srv, login)
//	pb.RegisterMenuAPIServer(srv, menus)
//}
