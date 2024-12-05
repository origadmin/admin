/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/config"
	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"
	"github.com/origadmin/toolkits/errors"

	pb "origadmin/application/admin/api/v1/services/common"
	"origadmin/application/admin/internal/configs"
	commonservice "origadmin/application/admin/internal/mods/common/service"
)

const (
	// ServiceName is service name.
	ServiceName = "common"
)

var (
// ProviderSet is server providers.
// ProviderSet = wire.NewSet(NewCommonServer, NewCommonServerAgent)
)

func init() {
	runtime.RegisterService("common", service.DefaultServiceBuilder)
}

type RegisterAgent struct {
	Login pb.LoginAPIGINRPCAgent
}

func (s RegisterAgent) GIN(server gins.IRouter) {
	log.Info("gin server common init")
	pb.RegisterLoginAPIGINRPCAgent(server, s.Login)
}

func NewCommonServerAgent(bootstrap *configs.Bootstrap, l log.Logger) (*RegisterAgent, error) {
	client, err := NewCommonClient(bootstrap, l)
	if err != nil {
		return nil, err
	}
	register := RegisterAgent{
		Login: NewLoginServerAgent(client),
	}
	return &register, nil
}

func NewCommonClient(bootstrap *configs.Bootstrap, l log.Logger) (*service.GRPCClient, error) {
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

func NewLoginServerAgent(client *service.GRPCClient) pb.LoginAPIGINRPCAgent {
	cli := pb.NewLoginAPIClient(client)
	return commonservice.NewLoginAPIGINRPCAgent(cli)
}

func NewLoginClient(service *configv1.Service) (pb.LoginAPIServer, error) {
	client, err := runtime.NewGRPCServiceClient(context.Background(), service)
	if err != nil {
		return nil, errors.Wrap(err, "create menu grpc client")
	}
	cli := pb.NewLoginAPIClient(client)
	return commonservice.NewLoginAPIServer(cli), nil
}

//func NewCommonHTTPClient(service *configv1.Service) (pb.LoginAPIServer, error) {
//	client, err := runtime.NewHTTPServiceClient(context.Background(), service)
//	if err != nil {
//		return nil, errors.Wrap(err, "create menu http client")
//	}
//	cli := pb.NewCommonAPIHTTPClient(client)
//	return commonservice.NewCommonAPIHTTPServer(cli), nil
//}
