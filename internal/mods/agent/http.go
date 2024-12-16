/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"github.com/goexts/generic/types"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/internal/configs"
)

type HTTPRegistrar interface {
	HTTP(server service.HTTPServer)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bootstrap *configs.Bootstrap, l log.Logger) *service.HTTPServer {
	serviceConfig := bootstrap.GetService()
	if serviceConfig == nil {
		panic("no serviceConfig")
	}
	serviceConfig.Name = types.ZeroOr(serviceConfig.Name, "ORIGADMIN_SERVICE")
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetService())
	if err != nil {
		panic(err)
	}

	return srv
}

//func RegisterHTTPServer(srv *http.Register, menus pb.MenuAPIServer, login common.LoginAPIServer) {
//	common.RegisterLoginAPIHTTPServer(srv, login)
//	pb.RegisterMenuAPIHTTPServer(srv, menus)
//}
