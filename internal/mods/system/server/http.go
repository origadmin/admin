/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	servicehttp "github.com/origadmin/runtime/service/http"

	"origadmin/application/admin/api/v1/services/common"
	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/configs"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bs *configs.Bootstrap, menus pb.MenuAPIServer, l log.Logger) *http.Server {
	srv := servicehttp.NewServer(bs.GetService())
	pb.RegisterMenuAPIHTTPServer(srv, menus)
	return srv
}

func RegisterHTTPServer(srv *http.Server, menus pb.MenuAPIServer, login common.LoginAPIServer) {
	common.RegisterLoginAPIHTTPServer(srv, login)
	pb.RegisterMenuAPIHTTPServer(srv, menus)
}
