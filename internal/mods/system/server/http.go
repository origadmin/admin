/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/internal/configs"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bootstrap *configs.Bootstrap, l log.KLogger, ss ...service.ServerOption) *service.HTTPServer {
	//options := settings.ApplyZero(ss)
	//for i, config := range bootstrap.GetServices() {
	//	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetServices(), options.ToHTTP())
	//	if err != nil {
	//		panic(err)
	//	}
	//	return srv
	//}
	return nil
}
