/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/internal/configs"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bootstrap *configs.Bootstrap, register func(*service.HTTPServer), l log.Logger) *service.HTTPServer {
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetService())
	if err != nil {
		panic(err)
	}
	if register != nil {
		register(srv)
	}
	return srv
}

//func RegisterHTTPServer(srv *http.Server, menus pb.MenuAPIServer, login common.LoginAPIServer) {
//	common.RegisterLoginAPIHTTPServer(srv, login)
//	pb.RegisterMenuAPIHTTPServer(srv, menus)
//}
