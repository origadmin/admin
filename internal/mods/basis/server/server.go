/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"
	servicegrpc "github.com/origadmin/runtime/service/grpc"
	"github.com/origadmin/toolkits/errors"

	pb "origadmin/application/admin/api/v1/services/basis"
	"origadmin/application/admin/internal/configs"
	basisservice "origadmin/application/admin/internal/mods/basis/service"
)

const (
	// ServiceName is service name.
	ServiceName = "basis"
)

var (
// ProviderSet is server providers.
// ProviderSet = wire.NewSet(NewBasisServer, NewBasisServerAgent)
)

func init() {
	runtime.RegisterService("basis", service.DefaultServiceBuilder)
}

type RegisterAgent struct {
	Login pb.LoginAPIAgent
}

func (s RegisterAgent) GIN(server gins.IRouter) {
	log.Info("gin server basis init")
	//pb.RegisterLoginAPIAgent(server, s.Login)
}

func (s RegisterAgent) GRPC(server *service.GRPCServer) {
	log.Info("grpc server basis init")
	//pb.RegisterLoginAPIServer(server, s.Login)
}

func (s RegisterAgent) HTTP(server *service.HTTPServer) {
	log.Info("http server basis init")
	pb.RegisterLoginAPIAgent(server, s.Login)
}

func NewBasisServerAgent(bootstrap *configs.Bootstrap, l log.KLogger) (*RegisterAgent, error) {
	client, err := NewBasisClient(bootstrap, l)
	if err != nil {
		return nil, err
	}
	register := RegisterAgent{
		Login: NewLoginServerAgent(client),
	}
	return &register, nil
}

func NewBasisClient(bootstrap *configs.Bootstrap, l log.KLogger) (*service.GRPCClient, error) {
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

	client, err := runtime.NewGRPCServiceClient(context.Background(), serviceConfig, service.WithGRPC(func(o *servicegrpc.Option) {
		// todo: add your grpc options here.
		o.Discovery = discovery
		o.ServiceName = registry.ServiceName
	}))
	if err != nil {
		return nil, errors.Wrap(err, "create menu grpc client")
	}
	return client, nil
}

func NewLoginServerAgent(client *service.GRPCClient) pb.LoginAPIAgent {
	cli := pb.NewLoginAPIClient(client)
	return basisservice.NewLoginAPIAgent(cli)
}

func NewLoginClient(service *configv1.Service) (pb.LoginAPIServer, error) {
	client, err := runtime.NewGRPCServiceClient(context.Background(), service)
	if err != nil {
		return nil, errors.Wrap(err, "create menu grpc client")
	}
	cli := pb.NewLoginAPIClient(client)
	return basisservice.NewLoginAPIServer(cli), nil
}

//func NewBasisHTTPClient(service *configv1.Service) (pb.LoginAPIServer, error) {
//	client, err := runtime.NewHTTPServiceClient(context.Background(), service)
//	if err != nil {
//		return nil, errors.Wrap(err, "create menu http client")
//	}
//	cli := pb.NewBasisAPIHTTPClient(client)
//	return basisservice.NewBasisAPIHTTPServer(cli), nil
//}
