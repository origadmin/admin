/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"
	"github.com/origadmin/toolkits/errors"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/configs"
	systemservice "origadmin/application/admin/internal/mods/system/service"
)

const (
	// ServiceName is service name.
	ServiceName = "common"
)

var (
	// ProviderSet is server providers.
	ProviderSet = wire.NewSet(NewCommonServer, NewCommonServerAgent)
	// ProviderServer is server providers.
	ProviderServer = wire.NewSet(NewCommonServer)
	// ProviderAgent is agent providers.
	ProviderAgent = wire.NewSet(NewCommonServerAgent)
)

func init() {
	runtime.RegisterService("common", service.DefaultServiceBuilder)
}

func NewCommonServer(register *systemservice.RegisterServer, bootstrap *configs.Bootstrap, l log.Logger) []transport.Server {
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

func NewCommonServerAgent(bootstrap *configs.Bootstrap, server *gin.Engine, l log.Logger) error {
	entry := bootstrap.GetEntry()
	if entry == nil {
		return nil
	}
	if entry.Scheme == "http" {
		client, err := NewMenuHTTPClient(&configv1.Service{
			Grpc: entry.GetGrpc(),
			Http: entry.GetHttp(),
			Gins: entry.GetGins(),
		})
		if err != nil {
			return errors.Wrap(err, "create menu http client")
		}
		pb.RegisterMenuAPIGINSServer(server, client)
		return nil
	}
	client, err := NewMenuClient(&configv1.Service{
		Grpc: entry.GetGrpc(),
		Http: entry.GetHttp(),
		Gins: entry.GetGins(),
	})
	if err != nil {
		return errors.Wrap(err, "create menu grpc client")
	}
	pb.RegisterMenuAPIGINSServer(server, client)
	return nil
}

func NewMenuClient(service *configv1.Service) (pb.MenuAPIServer, error) {
	client, err := runtime.NewGRPCServiceClient(context.Background(), service)
	if err != nil {
		return nil, errors.Wrap(err, "create menu grpc client")
	}
	cli := pb.NewMenuAPIClient(client)
	return systemservice.NewMenuAPIServer(cli), nil
}

func NewMenuHTTPClient(service *configv1.Service) (pb.MenuAPIServer, error) {
	client, err := runtime.NewHTTPServiceClient(context.Background(), service)
	if err != nil {
		return nil, errors.Wrap(err, "create menu http client")
	}
	cli := pb.NewMenuAPIHTTPClient(client)
	return systemservice.NewMenuAPIHTTPServer(cli), nil
}
