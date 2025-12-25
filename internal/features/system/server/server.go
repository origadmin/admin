/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/system/service"

	grpcv1 "github.com/origadmin/runtime/api/gen/go/config/transport/grpc/v1"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the system service servers (gRPC, HTTP).
func NewServers(cfg *transportv1.Servers, svc *service.SystemService, logger log.Logger) ([]transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("servers config is nil")
	}

	var transportServers []transport.Server
	for _, serverCfg := range cfg.GetConfigs() {
		switch serverCfg.GetProtocol() {
		case "http":
			srv, err := NewHTTPServer(serverCfg.GetHttp(), svc, logger)
			if err != nil {
				return nil, err
			}
			srv.Server
			transportServers = append(transportServers, srv)
		case "grpc":
			srv, err := NewGRPCServer(serverCfg.GetGrpc(), svc, logger)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		default:
			return nil, errors.New("protocol is not supported: " + serverCfg.GetProtocol())
		}
	}
	return transportServers, nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cfg *httpv1.Server, svc *service.SystemService, logger log.Logger) (*http.Server, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	var opts []http.ServerOption
	if cfg.GetAddr() != "" {
		opts = append(opts, http.Address(cfg.GetAddr()))
	}
	if cfg.GetTimeout() != nil {
		opts = append(opts, http.Timeout(cfg.GetTimeout().AsDuration()))
	}
	srv := http.NewServer(opts...)

	// Register HTTP handlers
	systemv1.RegisterUserServiceHTTPServer(srv, svc)
	systemv1.RegisterRoleServiceHTTPServer(srv, svc)
	systemv1.RegisterPermissionServiceHTTPServer(srv, svc)
	systemv1.RegisterResourceServiceHTTPServer(srv, svc)
	srv.WalkHandle(func(method, path string, handler stdhttp.HandlerFunc) {
		log.Infof("HTTP %s %s", method, path)
	})
	return srv, nil
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(cfg *grpcv1.Server, svc *service.SystemService, logger log.Logger) (*grpc.Server, error) {
	if cfg == nil {
		return nil, errors.New("grpc config is nil")
	}

	var opts []grpc.ServerOption
	if cfg.GetAddr() != "" {
		opts = append(opts, grpc.Address(cfg.GetAddr()))
	}
	if cfg.GetTimeout() != nil {
		opts = append(opts, grpc.Timeout(cfg.GetTimeout().AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	// Register gRPC handlers
	systemv1.RegisterUserServiceServer(srv, svc)
	systemv1.RegisterRoleServiceServer(srv, svc)
	systemv1.RegisterPermissionServiceServer(srv, svc)
	systemv1.RegisterResourceServiceServer(srv, svc)

	return srv, nil
}
