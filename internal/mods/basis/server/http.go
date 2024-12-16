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

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bootstrap *configs.Bootstrap, l log.Logger) *service.HTTPServer {
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetService())
	if err != nil {
		panic(err)
	}
	return srv
}

//func RegisterHTTPServer(srv *http.Register, menus pb.MenuAPIServer, login basis.LoginAPIServer) {
//	basis.RegisterLoginAPIHTTPServer(srv, login)
//	pb.RegisterMenuAPIHTTPServer(srv, menus)
//}