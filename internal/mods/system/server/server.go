/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/config"
	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"
	"github.com/origadmin/toolkits/errors"

	commonpb "origadmin/application/admin/api/v1/services/common"
	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/configs"
	commonservice "origadmin/application/admin/internal/mods/common/service"
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

func NewSystemServer(register *systemservice.RegisterServer, bootstrap *configs.Bootstrap, l log.Logger) []transport.Server {
	var servers []transport.Server
	service := bootstrap.GetService()
	if service == nil {
		return servers
	}
	if service.Name == "" {
		service.Name = ServiceName
	}
	if serv := NewGRPCServer(bootstrap, l); serv != nil {
		servers = append(servers, register.GRPC(serv))
	}
	if serv := NewGINSServer(bootstrap, l); serv != nil {
		servers = append(servers, register.GINS(serv))
	}
	if serv := NewHTTPServer(bootstrap, l); serv != nil {
		servers = append(servers, register.HTTP(serv))
	}
	return servers
}

type RegisterAgent struct {
	Menu  pb.MenuAPIGINRPCAgent
	Role  pb.RoleAPIGINRPCAgent
	User  pb.UserAPIGINRPCAgent
	Login commonpb.LoginAPIGINRPCAgent
}

func (s RegisterAgent) GIN(server gins.IRouter) {
	log.Info("gin server init")
	pb.RegisterMenuAPIGINRPCAgent(server, s.Menu)
	pb.RegisterRoleAPIGINRPCAgent(server, s.Role)
	pb.RegisterUserAPIGINRPCAgent(server, s.User)
	commonpb.RegisterLoginAPIGINRPCAgent(server, s.Login)
}

func NewSystemServerAgent(bootstrap *configs.Bootstrap, l log.Logger) (*RegisterAgent, error) {
	client, err := NewSystemClient(bootstrap, l)
	if err != nil {
		return nil, err
	}
	register := RegisterAgent{
		Menu:  NewMenuServerAgent(client),
		Role:  NewRoleServerAgent(client),
		User:  NewUserServerAgent(client),
		Login: NewLoginServerAgent(client),
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
	service := &configv1.Service{
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

	client, err := runtime.NewGRPCServiceClient(context.Background(), service, config.WithServiceOption(config.WithServiceDiscovery(registry.ServiceName, discovery)))
	if err != nil {
		return nil, errors.Wrap(err, "create menu grpc client")
	}
	return client, nil
}

func NewLoginServerAgent(client *service.GRPCClient) commonpb.LoginAPIGINRPCAgent {
	cli := commonpb.NewLoginAPIClient(client)
	return commonservice.NewLoginAPIGINRPCAgent(cli)
}

func NewUserServerAgent(client *service.GRPCClient) pb.UserAPIGINRPCAgent {
	c := pb.NewUserAPIClient(client)
	return systemservice.NewUserAPIGINRPCAgent(c)
}

func NewRoleServerAgent(client *service.GRPCClient) pb.RoleAPIGINRPCAgent {
	c := pb.NewRoleAPIClient(client)
	return systemservice.NewRoleAPIGINRPCAgent(c)
}

func NewMenuServerAgent(client *service.GRPCClient) pb.MenuAPIGINRPCAgent {
	cli := pb.NewMenuAPIClient(client)
	return systemservice.NewMenuAPIGINRPCAgent(cli)
}

//func NewMenuHTTPServerAgent(service *configv1.Service) (pb.MenuAPIGINRPCAgent, error) {
//	client, err := runtime.NewHTTPServiceClient(context.Background(), service)
//	if err != nil {
//		return nil, errors.Wrap(err, "create menu http client")
//	}
//	cli := pb.NewMenuAPIHTTPClient(client)
//	return systemservice.NewMenuAPIGINRPCAgent(cli), nil
//}
