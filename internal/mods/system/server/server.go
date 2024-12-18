/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service"
	servicegrpc "github.com/origadmin/runtime/service/grpc"
	servicehttp "github.com/origadmin/runtime/service/http"
	"github.com/origadmin/toolkits/errors"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/configs"
	basisservice "origadmin/application/admin/internal/mods/basis/service"
	systemservice "origadmin/application/admin/internal/mods/system/service"
)

const (
	// ServiceName is service name.
	ServiceName = "system"
)

var (
	// ProviderSet is server providers.
	ProviderSet = wire.NewSet(NewSystemServer, NewSystemServerAgent)
)

func init() {
	runtime.RegisterService(ServiceName, service.DefaultServiceBuilder)
}

func NewSystemServer(register *systemservice.RegisterServer, register2 basisservice.RegisterServer, bootstrap *configs.Bootstrap, l log.Logger) []transport.Server {
	var servers []transport.Server
	middlewares := middleware.NewServer(bootstrap.GetMiddlewares())
	serviceConfig := bootstrap.GetService()
	if serviceConfig == nil {
		return servers
	}
	if serviceConfig.Name == "" {
		serviceConfig.Name = ServiceName
	}
	if serv := NewGRPCServer(bootstrap, l, service.WithGRPC(
		servicegrpc.WithMiddlewares(middlewares...)),
	); serv != nil {
		serv = register.GRPC(serv)
		serv = register2.GRPC(serv)
		servers = append(servers, serv)
	}
	if serv := NewGINSServer(bootstrap, l, service.WithHTTP(
		servicehttp.WithMiddlewares(middlewares...)),
	); serv != nil {
		serv = register.GINS(serv)
		serv = register2.GINS(serv)
		servers = append(servers, serv)
	}
	if serv := NewHTTPServer(bootstrap, l, service.WithHTTP(
		servicehttp.WithMiddlewares(middlewares...)),
	); serv != nil {
		serv = register.HTTP(serv)
		serv = register2.HTTP(serv)
		servers = append(servers, serv)
	}
	return servers
}

type RegisterAgent struct {
	Menu           pb.MenuAPIGINRPCAgent
	Role           pb.RoleAPIGINRPCAgent
	User           pb.UserAPIGINRPCAgent
	basisRegisters []func(server gins.IRouter)
}

func (s RegisterAgent) GIN(server gins.IRouter) {
	log.Info("gin server system init")
	for _, register := range s.basisRegisters {
		register(server)
	}
	pb.RegisterMenuAPIGINRPCAgent(server, s.Menu)
	pb.RegisterRoleAPIGINRPCAgent(server, s.Role)
	pb.RegisterUserAPIGINRPCAgent(server, s.User)
}

func NewSystemServerAgent(bootstrap *configs.Bootstrap, l log.Logger) (*RegisterAgent, error) {
	client, err := NewSystemClient(bootstrap, l)
	if err != nil {
		return nil, err
	}
	register := RegisterAgent{
		Menu: systemservice.NewMenuServerAgent(client),
		Role: systemservice.NewRoleServerAgent(client),
		User: systemservice.NewUserServerAgent(client),
		basisRegisters: []func(server gins.IRouter){
			basisservice.NewLoginServerAgentGINRegister(client),
		},
	}
	return &register, nil
}

func NewSystemClient(bootstrap *configs.Bootstrap, l log.Logger) (*service.GRPCClient, error) {
	entry := bootstrap.GetEntry()
	if entry == nil {
		return nil, errors.New("no entry")
	}

	servers := bootstrap.GetServers()
	if servers == nil {
		return nil, errors.New("no servers")
	}
	registry := bootstrap.GetRegistry()
	if registry == nil {
		return nil, errors.New("no registry")
	}
	serviceConfig := &configv1.Service{
		Name: ServiceName,
		Grpc: entry.GetGrpc(),
		Http: entry.GetHttp(),
		Gins: entry.GetGins(),
		Selector: &configv1.Service_Selector{
			Version: "v1.0.0",
			Builder: "bbr",
		},
	}
	if v, ok := bootstrap.GetServers()[ServiceName]; ok {
		registry.ServiceName = v
	}
	log.Infof("service name: %s", registry.ServiceName)
	discovery, err := runtime.NewDiscovery(registry)
	if err != nil {
		return nil, errors.Wrap(err, "create discovery")
	}

	client, err := runtime.NewGRPCServiceClient(context.Background(), serviceConfig,
		service.WithGRPC(func(o *servicegrpc.Option) {
			o.Discovery = discovery
			o.ServiceName = registry.ServiceName
		}))
	if err != nil {
		return nil, errors.Wrap(err, "create menu grpc client")
	}
	return client, nil
}

//func NewMenuHTTPServerAgent(service *configv1.Service) (pb.MenuAPIGINRPCAgent, error) {
//	client, err := runtime.NewHTTPServiceClient(context.Background(), service)
//	if err != nil {
//		return nil, errors.Wrap(err, "create menu http client")
//	}
//	cli := pb.NewMenuAPIHTTPClient(client)
//	return systemservice.NewMenuAPIGINRPCAgent(cli), nil
//}
