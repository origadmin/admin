/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/google/wire"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service/transport"
	configs "github.com/origadmin/runtime/test/integration/app/proto"
	"github.com/origadmin/toolkits/errors"
	systemservice "origadmin/application/admin/internal/features/system/service"
)

const (
	// ServiceName is service name.
	ServiceName = "system"
)

var (
	// ProviderSet is server providers.
	ProviderSet = wire.NewSet(
		NewSystemClient,
		NewSystemServer,
	)
)

func init() {
}

func NewSystemServer(app *runtime.App, bootstrap *configs.Bootstrap,
	svc systemservice.SystemService) []transport.
Server {
	var servers []transport.Server

	return servers
}

func NewSystemClient(app *runtime.App, bootstrap *configs.Bootstrap) (*transport.GRPCClient, error) {
	discovery := bootstrap.GetDiscovery()
	if discovery == nil {
		return nil, errors.New("no discovery")
	}
	serviceConfig := &configv1.Service{
		Name: ServiceName,
		//Grpc: entry.GetGrpc(),
		//Http: entry.GetHttp(),
		Selector: &configv1.Service_Selector{
			Version: "v1.0.0",
			Builder: "bbr",
		},
	}
	//if v, ok := bootstrap.GetServices()[ServiceName]; ok {
	//	discovery.ServiceName = ServiceName
	//}
	helper := log.NewHelper(r.Logger())
	//discovery.ServiceName = ServiceName
	helper.Infof("service name: %s", discovery.ServiceName)
	discover, err := runtime.NewDiscovery(discovery)
	if err != nil {
		return nil, errors.Wrap(err, "create discovery")
	}
	var ms []middleware.KMiddleware
	options := []servicegrpc.Option{
		servicegrpc.WithDiscovery(discovery.ServiceName, discover),
	}
	ms = append(ms, middleware.NewClient(bootstrap.GetMiddleware())...)
	ms = append(ms, MiddlewareServer())
	if len(ms) > 0 {
		options = append(options, servicegrpc.WithMiddlewares(ms...))
	}
	client, err := runtime.NewGRPCServiceClient(context.Background(), serviceConfig, options...)
	if err != nil {
		return nil, errors.Wrap(err, "create menu grpc client")
	}
	return client, nil
}

func MiddlewareServer() middleware.KMiddleware {
	return func(handler middleware.KHandler) middleware.KHandler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if md, ok := metadata.FromClientContext(ctx); ok {
				log.Debugf("MiddlewareServer: found client context metadata: %+v", md)
			} else {
				log.Debugf("MiddlewareServer: no client context metadata found")
			}
			if md, ok := metadata.FromServerContext(ctx); ok {
				log.Debugf("MiddlewareServer: found server context metadata: %+v", md)
			} else {
				log.Debugf("MiddlewareServer: no server context metadata found")
			}
			reply, err = handler(ctx, req)
			return
		}
	}
}
