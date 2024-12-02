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
	"github.com/origadmin/runtime/service/grpc"
	http "github.com/origadmin/runtime/service/http"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/configs"
	systemservice "origadmin/application/admin/internal/mods/system/service"
)

var (
	// ProviderServer is server providers.
	ProviderServer = wire.NewSet(NewSystemServers, wire.Struct(new(Server), "*"))
	// ProviderAgent is agent providers.
	ProviderAgent = wire.NewSet(NewGinHTTPServer)
)

func init() {
	runtime.RegisterService("system", &serverBuilder{})
}

type Server struct {
	Menu pb.MenuAPIServer
	Role pb.RoleAPIServer
	User pb.UserAPIServer
}

func (s Server) GRPC(server *service.GRPCServer) {
	log.Info("grpc server init")
	pb.RegisterMenuAPIServer(server, s.Menu)
	pb.RegisterRoleAPIServer(server, s.Role)
	pb.RegisterUserAPIServer(server, s.User)
}

func (s Server) GINS(server *gins.Server) {
	log.Info("gins server init")
	pb.RegisterMenuAPIGINSServer(server, s.Menu)
	pb.RegisterRoleAPIGINSServer(server, s.Role)
	pb.RegisterUserAPIGINSServer(server, s.User)
}

func (s Server) HTTP(server *service.HTTPServer) {
	log.Info("http server init")
	pb.RegisterMenuAPIHTTPServer(server, s.Menu)
	pb.RegisterRoleAPIHTTPServer(server, s.Role)
	pb.RegisterUserAPIHTTPServer(server, s.User)
}

type serverBuilder struct {
}

func (s serverBuilder) NewGRPCServer(cfg *configv1.Service, opts ...config.ServiceSetting) (*service.GRPCServer, error) {
	return grpc.NewServer(cfg, opts...), nil
}

func (s serverBuilder) NewHTTPServer(cfg *configv1.Service, opts ...config.ServiceSetting) (*service.HTTPServer, error) {
	return http.NewServer(cfg, opts...), nil
}

func (s serverBuilder) NewGRPCClient(ctx context.Context, cfg *configv1.Service, opts ...config.ServiceSetting) (*service.GRPCClient, error) {
	return grpc.NewClient(ctx, cfg, opts...)
}

func (s serverBuilder) NewHTTPClient(ctx context.Context, cfg *configv1.Service, opts ...config.ServiceSetting) (*service.HTTPClient, error) {
	return http.NewClient(ctx, cfg, opts...)
}

func NewSystemServers(bootstrap *configs.Bootstrap, s *Server, l log.Logger) []transport.Server {
	var servers []transport.Server
	service := bootstrap.GetService()
	if service == nil {
		return servers
	}
	if service.Name == "" {
		service.Name = "system"
	}
	if serv := NewGRPCServer(bootstrap, s.GRPC, l); serv != nil {
		servers = append(servers, serv)
	}
	if serv := NewGINSServer(bootstrap, s.GINS, l); serv != nil {
		servers = append(servers, serv)
	}
	if serv := NewHTTPServer(bootstrap, s.HTTP, l); serv != nil {
		servers = append(servers, serv)
	}
	return servers
}

func NewGinMenuAgent(bootstrap *configs.Bootstrap, server *gins.Server, l log.Logger) error {
	entry := bootstrap.GetEntry()
	if entry == nil {
		return nil
	}
	if entry.Scheme == "http" {
		if client, err := NewMenuHTTPClient(&configv1.Service{
			Grpc: entry.GetGrpc(),
			Http: entry.GetHttp(),
			Gins: entry.GetGins(),
		}); err == nil {
			pb.RegisterMenuAPIGINSServer(server, client)
		}
		return nil
	}
	if client, err := NewMenuClient(&configv1.Service{
		Grpc: entry.GetGrpc(),
		Http: entry.GetHttp(),
		Gins: entry.GetGins(),
	}); err == nil {
		pb.RegisterMenuAPIGINSServer(server, client)
	}
	return nil
}

func NewMenuClient(service *configv1.Service) (pb.MenuAPIServer, error) {
	client, err := runtime.NewGRPCServiceClient(context.Background(), service)
	if err != nil {
		return nil, err
	}
	cli := pb.NewMenuAPIClient(client)
	return systemservice.NewMenuAPIServer(cli), nil
}

func NewMenuHTTPClient(service *configv1.Service) (pb.MenuAPIServer, error) {
	client, err := runtime.NewHTTPServiceClient(context.Background(), service)
	if err != nil {
		return nil, err
	}
	cli := pb.NewMenuAPIHTTPClient(client)
	return systemservice.NewMenuAPIHTTPServer(cli), nil
}
