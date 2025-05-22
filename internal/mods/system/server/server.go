/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/agent"
	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service"
	servicegrpc "github.com/origadmin/runtime/service/grpc"
	"github.com/origadmin/toolkits/errors"

	pb "origadmin/application/admin/api/v1/services/system"
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
		NewRegisterServer,
		NewSystemClient,
		NewSystemServer,
		NewSystemServiceAgentClient,
		NewCasbinServiceClient,
	)
)

func init() {
	runtime.RegisterService(ServiceName, service.DefaultServiceFactory)
}

func NewSystemServer(r runtime.Runtime, bootstrap *configs.Bootstrap, registers []service.ServerRegister) []transport.
Server {
	var servers []transport.Server
	serviceConfig := bootstrap.GetServices()
	if serviceConfig == nil {
		return servers
	}
	//if serviceConfig.Name == "" {
	//	serviceConfig.Name = ServiceName
	//}
	//ctx := context.Background()
	//middlewares := middleware.NewServer(bootstrap.GetMiddleware())

	//if serv, _ := runtime.NewGRPCServiceServer(bootstrap, l, service.WithGRPC(
	//	servicegrpc.WithMiddlewares(middlewares...),
	//	servicegrpc.WithPrefix(runtime.DefaultEnvPrefix),
	//)); serv != nil {
	//	for i := range registers {
	//		registers[i].GRPCServer(ctx, serv)
	//	}
	//	servers = append(servers, serv)
	//}
	//if serv := NewHTTPServer(bootstrap, l, service.WithHTTP(
	//	servicehttp.WithMiddlewares(middlewares...),
	//	servicehttp.WithPrefix(runtime.DefaultEnvPrefix),
	//)); serv != nil {
	//	for i := range registers {
	//		registers[i].HTTPServer(ctx, serv)
	//	}
	//	servers = append(servers, serv)
	//}
	return servers
}

type RegisterAgent struct {
	Auth       pb.AuthServiceAgent
	Login      pb.LoginServiceAgent
	Personal   pb.PersonalServiceAgent
	Resource   pb.ResourceServiceAgent
	Role       pb.RoleServiceAgent
	User       pb.UserServiceAgent
	Permission pb.PermissionServiceAgent
}

func (s RegisterAgent) GRPCServer(ctx context.Context, server *service.GRPCServer) {
	log.Info("grpc server system init")
}

func (s RegisterAgent) HTTPServer(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server system init")
	ag := agent.NewHTTP(server)
	pb.RegisterAuthServiceAgent(ag, s.Auth)
	pb.RegisterLoginServiceAgent(ag, s.Login)
	pb.RegisterPersonalServiceAgent(ag, s.Personal)
	pb.RegisterResourceServiceAgent(ag, s.Resource)
	pb.RegisterRoleServiceAgent(ag, s.Role)
	pb.RegisterUserServiceAgent(ag, s.User)
	pb.RegisterPermissionServiceAgent(ag, s.Permission)
}

func (s RegisterAgent) Server(ctx context.Context, grpcServer *service.GRPCServer, httpServer *service.HTTPServer) {
	s.HTTPServer(ctx, httpServer)
	s.GRPCServer(ctx, grpcServer)
}

func NewSystemServiceAgentClient(client *service.GRPCClient, l log.KLogger) (*RegisterAgent, error) {
	register := RegisterAgent{
		Auth:       systemservice.NewAuthServiceAgentClient(client),
		Login:      systemservice.NewLoginServiceAgentClient(client),
		Personal:   systemservice.NewPersonalServiceAgentClient(client),
		Resource:   systemservice.NewResourceServiceAgentClient(client),
		Role:       systemservice.NewRoleServiceAgentClient(client),
		User:       systemservice.NewUserServiceAgentClient(client),
		Permission: systemservice.NewPermissionServiceAgentClient(client),
	}
	return &register, nil
}

func NewCasbinServiceClient(client *service.GRPCClient, l log.KLogger) pb.CasbinSourceServiceClient {
	return systemservice.NewCasbinSourceServiceClient(client)
}

func NewSystemClient(r runtime.Runtime, bootstrap *configs.Bootstrap) (*service.GRPCClient, error) {
	entry := bootstrap.GetEntry()
	if entry == nil {
		return nil, errors.New("no entry")
	}

	//servers := bootstrap.GetServers()
	//if servers == nil {
	//	return nil, errors.New("no servers")
	//}
	registry := bootstrap.GetRegistry()
	if registry == nil {
		return nil, errors.New("no registry")
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
	//	registry.ServiceName = ServiceName
	//}
	helper := log.NewHelper(r.Logger())
	//registry.ServiceName = ServiceName
	helper.Infof("service name: %s", registry.ServiceName)
	discovery, err := runtime.NewDiscovery(registry)
	if err != nil {
		return nil, errors.Wrap(err, "create discovery")
	}
	var ms []middleware.KMiddleware
	options := []servicegrpc.Option{
		servicegrpc.WithDiscovery(registry.ServiceName, discovery),
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

func NewRegisterServer(
	Resource pb.ResourceServiceServer,
	Role pb.RoleServiceServer,
	User pb.UserServiceServer,
//Auth pb.AuthServiceServer,
//Login pb.LoginServiceServer,
//Personal pb.PersonalServiceServer,
	Permission pb.PermissionServiceServer,
//Casbin pb.CasbinSourceServiceServer,
) []service.ServerRegister {
	return []service.ServerRegister{
		&systemservice.RegisterServer{
			Resource: Resource,
			Role:     Role,
			User:     User,
			//Auth:       Auth,
			//Login:      Login,
			//Personal:   Personal,
			Permission: Permission,
			//Casbin:     Casbin,
		},
	}
}

var _ service.ServerRegister = (*RegisterAgent)(nil)
