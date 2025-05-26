/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/internal/configs"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(r runtime.Runtime, bootstrap *configs.Bootstrap) *service.GRPCServer {
	services := bootstrap.GetServices()
	for _, config := range services {
		serviceConfig := config.GetService()
		if serviceConfig.GetType() == "grpc" {
			grpcServer, err := r.Builder().NewGRPCServer(serviceConfig)
			if err != nil {
				return nil
			}
			return grpcServer
		}
	}
	return nil
}
