/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/internal/configs"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(r runtime.Runtime, bootstrap *configs.Bootstrap) *service.HTTPServer {
	services := bootstrap.GetServer().GetServices()
	for _, serviceConfig := range services {
		if serviceConfig.GetType() == "http" {
			httpServer, err := r.Builder().NewHTTPServer(serviceConfig)
			if err != nil {
				return nil
			}
			return httpServer
		}
	}
	return nil
}
