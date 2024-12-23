/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/agent"
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

func NewSystemServer(register *systemservice.RegisterServer, register2 basisservice.RegisterServer, bootstrap *configs.Bootstrap, l log.KLogger) []transport.Server {
	var servers []transport.Server
	middlewares := middleware.NewServer(bootstrap.GetMiddleware())
	serviceConfig := bootstrap.GetService()
	if serviceConfig == nil {
		return servers
	}
	if serviceConfig.Name == "" {
		serviceConfig.Name = ServiceName
	}
	if serv := NewGRPCServer(bootstrap, l, service.WithGRPC(
		servicegrpc.WithMiddlewares(middlewares...),
		servicegrpc.WithPrefix(runtime.DefaultEnvPrefix),
	)); serv != nil {
		serv = register.GRPC(serv)
		serv = register2.GRPC(serv)
		servers = append(servers, serv)
	}
	//if serv := NewGINSServer(bootstrap, l, service.WithHTTP(
	//	servicehttp.WithMiddlewares(middlewares...)),
	//); serv != nil {
	//	serv = register.GINS(serv)
	//	serv = register2.GINS(serv)
	//	servers = append(servers, serv)
	//}
	if serv := NewHTTPServer(bootstrap, l, service.WithHTTP(
		servicehttp.WithMiddlewares(middlewares...),
		servicehttp.WithPrefix(runtime.DefaultEnvPrefix),
	)); serv != nil {
		serv = register.HTTP(serv)
		serv = register2.HTTP(serv)
		servers = append(servers, serv)
	}
	return servers
}

type RegisterAgent struct {
	Menu           pb.MenuAPIAgent
	Role           pb.RoleAPIAgent
	User           pb.UserAPIAgent
	basisRegisters []func(server *http.Server)
}

func (s RegisterAgent) HTTP(server *http.Server) {
	log.Info("http server system init")
	log.Info("gin server system init")
	for _, register := range s.basisRegisters {
		register(server)
	}
	ag := agent.New(server)
	pb.RegisterMenuAPIAgent(ag, s.Menu)
	pb.RegisterRoleAPIAgent(ag, s.Role)
	pb.RegisterUserAPIAgent(ag, s.User)
}

func (s RegisterAgent) GIN(server gins.IRouter) {
	log.Info("gin server system init")
	//for _, register := range s.basisRegisters {
	//	register(server)
	//}
	//pb.RegisterMenuAPIAgent(server, s.Menu)
	//pb.RegisterRoleAPIAgent(server, s.Role)
	//pb.RegisterUserAPIAgent(server, s.User)
}

func NewSystemServerAgent(bootstrap *configs.Bootstrap, l log.KLogger) (*RegisterAgent, error) {
	client, err := NewSystemClient(bootstrap, l)
	if err != nil {
		return nil, err
	}
	register := RegisterAgent{
		Menu: systemservice.NewMenuServerAgent(client),
		Role: systemservice.NewRoleServerAgent(client),
		User: systemservice.NewUserServerAgent(client),
		basisRegisters: []func(server *http.Server){
			basisservice.NewLoginServerAgentGINRegister(client),
		},
	}
	return &register, nil
}

func NewSystemClient(bootstrap *configs.Bootstrap, l log.KLogger) (*service.GRPCClient, error) {
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
	var ms []middleware.KMiddleware
	options := []servicegrpc.OptionSetting{
		servicegrpc.WithDiscovery(registry.ServiceName, discovery),
	}
	ms = middleware.NewClient(bootstrap.GetMiddleware())
	if len(ms) > 0 {
		options = append(options, servicegrpc.WithMiddlewares(ms...))
	}
	client, err := runtime.NewGRPCServiceClient(context.Background(), serviceConfig, service.WithGRPC(options...))
	if err != nil {
		return nil, errors.Wrap(err, "create menu grpc client")
	}
	return client, nil
}

//func NewMenuHTTPServerAgent(service *configv1.Service) (pb.MenuAPIAgent, error) {
//	client, err := runtime.NewHTTPServiceClient(context.Background(), service)
//	if err != nil {
//		return nil, errors.Wrap(err, "create menu http client")
//	}
//	cli := pb.NewMenuAPIHTTPClient(client)
//	return systemservice.NewMenuAPIAgent(cli), nil
//}
