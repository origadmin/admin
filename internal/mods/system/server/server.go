/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
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
	servicehttp "github.com/origadmin/runtime/service/http"
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
	ProviderSet = wire.NewSet(NewSystemClient, NewRegisterServer, NewSystemServer, NewSystemServerAgent)
)

func init() {
	runtime.RegisterService(ServiceName, service.DefaultServiceBuilder)
}

func NewSystemServer(bootstrap *configs.Bootstrap, registers []service.ServerRegister, l log.KLogger) []transport.Server {
	var servers []transport.Server
	middlewares := middleware.NewServer(bootstrap.GetService().GetMiddleware())
	serviceConfig := bootstrap.GetService()
	if serviceConfig == nil {
		return servers
	}
	if serviceConfig.Name == "" {
		serviceConfig.Name = ServiceName
	}
	ctx := context.Background()
	if serv := NewGRPCServer(bootstrap, l, service.WithGRPC(
		servicegrpc.WithMiddlewares(middlewares...),
		servicegrpc.WithPrefix(runtime.DefaultEnvPrefix),
	)); serv != nil {
		for i := range registers {
			registers[i].GRPCServer(ctx, serv)
		}
		servers = append(servers, serv)
	}
	if serv := NewHTTPServer(bootstrap, l, service.WithHTTP(
		servicehttp.WithMiddlewares(middlewares...),
		servicehttp.WithPrefix(runtime.DefaultEnvPrefix),
	)); serv != nil {
		for i := range registers {
			registers[i].HTTPServer(ctx, serv)
		}
		servers = append(servers, serv)
	}
	return servers
}

type RegisterAgent struct {
	Auth    pb.AuthAPIAgent
	Login   pb.LoginAPIAgent
	Current pb.CurrentAPIAgent
	Menu    pb.MenuAPIAgent
	Role    pb.RoleAPIAgent
	User    pb.UserAPIAgent
}

func (s RegisterAgent) GRPCServer(ctx context.Context, server *service.GRPCServer) {
	log.Info("grpc server system init")
}

func (s RegisterAgent) HTTPServer(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server system init")
	ag := agent.NewHTTP(server)
	pb.RegisterAuthAPIAgent(ag, s.Auth)
	pb.RegisterLoginAPIAgent(ag, s.Login)
	pb.RegisterCurrentAPIAgent(ag, s.Current)
	pb.RegisterMenuAPIAgent(ag, s.Menu)
	pb.RegisterRoleAPIAgent(ag, s.Role)
	pb.RegisterUserAPIAgent(ag, s.User)
}

func (s RegisterAgent) Server(ctx context.Context, grpcServer *service.GRPCServer, httpServer *service.HTTPServer) {
	s.HTTPServer(ctx, httpServer)
	s.GRPCServer(ctx, grpcServer)
}

func NewSystemServerAgent(client *service.GRPCClient, l log.KLogger) (*RegisterAgent, error) {
	register := RegisterAgent{
		Auth:    systemservice.NewAuthServerAgent(client),
		Login:   systemservice.NewLoginServerAgent(client),
		Current: systemservice.NewCurrentServerAgent(client),
		Menu:    systemservice.NewMenuServerAgent(client),
		Role:    systemservice.NewRoleServerAgent(client),
		User:    systemservice.NewUserServerAgent(client),
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

func NewRegisterServer(s1 *systemservice.RegisterServer) []service.ServerRegister {
	return []service.ServerRegister{
		s1,
	}
}

var _ service.ServerRegister = (*RegisterAgent)(nil)
