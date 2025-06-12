/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
	"github.com/origadmin/runtime"
	configv1 "github.com/origadmin/runtime/api/gen/go/config/v1"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service"
	servicegrpc "github.com/origadmin/runtime/service/grpc"
	servicehttp "github.com/origadmin/runtime/service/http"
	"github.com/origadmin/toolkits/errors"

	"origadmin/application/admin/internal/configs"
	systemservice "origadmin/application/admin/internal/mods/system/service"
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
	runtime.RegisterService(ServiceName, service.DefaultServiceFactory)
}

func NewSystemServer(r runtime.Runtime, bootstrap *configs.Bootstrap, svc systemservice.SystemServerRegistrar) []transport.
Server {
	var servers []transport.Server
	serverConfig := bootstrap.GetServer()
	if serverConfig == nil {
		return servers
	}

	ll := log.NewHelper(r.WithLogger("module", "system/server"))
	middlewares := r.Builder().Middleware().BuildServer(bootstrap.GetServer().GetMiddleware())
	services := bootstrap.GetServer().GetServices()
	coreinfo := bootstrap.GetServer().GetCore()
	for _, serviceConfig := range services {
		ll.Infow("msg", "service init", "name", serviceConfig.GetName(), "type", serviceConfig.GetType())
		var option service.ServerOption
		switch serviceConfig.GetType() {
		case "grpc":
			options := []servicegrpc.Option{
				servicegrpc.WithMiddlewares(middlewares...),
				servicegrpc.WithPrefix(runtime.DefaultEnvPrefix),
			}
			option = service.WithGRPC(options...)
		case "http":
			options := []servicehttp.Option{
				servicehttp.WithMiddlewares(middlewares...),
				servicehttp.WithPrefix(runtime.DefaultEnvPrefix),
			}
			//httpServer, err := r.Builder().NewServer(serviceConfig, options...)
			//if err != nil {
			//	continue
			//}
			//ll.Infow("msg", "http server init", "name", coreinfo.GetName(), "version",
			//	coreinfo.GetVersion())
			//svc.Register(r.Context(), httpServer)
			//servers = append(servers, httpServer)
			option = service.WithHTTP(options...)
		default:
			ll.Warnw("msg", "service type not support", "name", serviceConfig.GetName(), "type", serviceConfig.GetType())
			continue
		}
		grpcServer, err := r.Builder().NewServer("system", serviceConfig, option)
		if err != nil {
			continue
		}
		ll.Infow("msg", "system server init", "name", coreinfo.GetName(), "version",
			coreinfo.GetVersion())
		svc.Register(r.Context(), grpcServer)
		servers = append(servers, grpcServer)
	}
	return servers
}

func NewSystemClient(r runtime.Runtime, bootstrap *configs.Bootstrap) (*service.GRPCClient, error) {
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
