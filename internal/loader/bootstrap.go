/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/selector/random"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/settings"
	"github.com/google/wire"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/bootstrap"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/codec"
	"github.com/origadmin/toolkits/errors"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/configs"
	systemservice "origadmin/application/admin/internal/mods/system/service"
)

var (
	ProviderSet = wire.NewSet(
		NewRegistrar,
		NewDiscovery,
		wire.Struct(new(InjectorServer), "*"),
		wire.Struct(new(InjectorClient), "*"),
	)
)

var (
	_ *gins.Server
	_ *http.Server
	_ *grpc.Server
)

type InjectorClient struct {
	Logger    log.Logger
	Bootstrap *configs.Bootstrap
	Discovery registry.Discovery
	Server    *gin.Engine
	Servers   []transport.Server
}

type InjectorServer struct {
	Logger    log.Logger
	Bootstrap *configs.Bootstrap
	Registrar registry.Registrar
	Servers   []transport.Server
}

func InjectorGinServer(injector *InjectorClient) error {
	if err := newHelloWorldServer(injector); err != nil {
		return err
	}
	if err := newSecondWorldServer(injector); err != nil {
		return err
	}
	return nil
}

func newHelloWorldServer(injector *InjectorClient) error {
	// Create route Filter: Filter instances whose version number is "2.0.0"
	filter := filter.Version("v1.0.0")
	// Create the Selector for the P2C load balancing algorithm and inject the route Filter
	selector.SetGlobalSelector(random.NewBuilder())
	//selector.SetGlobalSelector(wrr.NewBuilder())

	serviceName := "origadmin.service.v1.helloworld"
	discovery := injector.Discovery
	if discovery == nil {
		return errors.String("discovery is nil")
	}
	//if discovery, ok := injector.Discoveries[serviceName]; ok {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
		),
		grpc.WithEndpoint("discovery:///"+serviceName),
		grpc.WithDiscovery(discovery),
		grpc.WithNodeFilter(filter),
	)
	if err != nil {
		return err
	}
	gClient := pb.NewMenuAPIClient(conn)
	// new http client
	hConn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
		),
		http.WithEndpoint("discovery:///"+serviceName),
		http.WithDiscovery(discovery),
		http.WithNodeFilter(filter),
	)
	if err != nil {
		return err
	}
	hClient := pb.NewMenuAPIHTTPClient(hConn)

	var client pb.MenuAPIServer
	if entry := injector.Bootstrap.GetEntry(); entry != nil && entry.Scheme == "http" {
		client = systemservice.NewMenuAPIHTTPServer(hClient)
	} else {
		client = systemservice.NewMenuAPIServer(gClient)
	}
	//service, err := runtime.NewGRPCServiceServer(injector.Bootstrap.GetService())
	//grpcClient := service.NewGreeterServer(gClient)
	//httpClient := service.NewGreeterHTTPServer(hClient)
	//// add _ to avoid unused
	//_ = grpcClient
	//_ = httpClient
	//pb.RegisterMenuAPIGINSServer(injector.ServerGINS, client)
	//pb.RegisterMenuAPIHTTPServer(injector.Agents, client)
	//}
	_ = client
	return nil
}

func newSecondWorldServer(injector *InjectorClient) error {
	// Create route Filter: Filter instances whose version number is "2.0.0"
	filter := filter.Version("v1.0.0")
	// Create the Selector for the P2C load balancing algorithm and inject the route Filter
	selector.SetGlobalSelector(random.NewBuilder())
	//selector.SetGlobalSelector(wrr.NewBuilder())

	serviceName := "origadmin.service.v1.pb"
	discovery := injector.Discovery
	if discovery == nil {
		return errors.String("discovery is nil")
	}
	//if discovery, ok := injector.Discoveries[serviceName]; ok {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
		),
		grpc.WithEndpoint("discovery:///"+serviceName),
		grpc.WithDiscovery(discovery),
		grpc.WithNodeFilter(filter),
	)
	if err != nil {
		return err
	}
	gClient := pb.NewMenuAPIClient(conn)
	// new http client
	hConn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
		),
		http.WithEndpoint("discovery:///"+serviceName),
		http.WithDiscovery(discovery),
		http.WithNodeFilter(filter),
	)
	if err != nil {
		return err
	}
	hClient := pb.NewMenuAPIHTTPClient(hConn)

	var client pb.MenuAPIServer
	if entry := injector.Bootstrap.GetEntry(); entry != nil && entry.Scheme == "http" {
		client = systemservice.NewMenuAPIHTTPServer(hClient)
	} else {
		client = systemservice.NewMenuAPIServer(gClient)
	}
	//grpcClient := service.NewGreeterServer(gClient)
	//httpClient := service.NewGreeterHTTPServer(hClient)
	//// add _ to avoid unused
	//_ = grpcClient
	//_ = httpClient
	//pb.RegisterMenuAPIGINSServer(injector.ServerGINS, client)
	//pb.RegisterMenuAPIHTTPServer(injector.Server, client)
	//}
	_ = client

	return nil
}

// LoadFileBootstrap load config from file
func LoadFileBootstrap(path string) (*configs.Bootstrap, error) {
	typo := codec.TypeFromPath(path)
	if typo == codec.UNKNOWN {
		return nil, fmt.Errorf("unknown file type: %s", path)
	}
	var cfg configs.Bootstrap
	err := codec.DecodeFromFile(path, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadLocalBootstrap(source *configv1.SourceConfig) (*configs.Bootstrap, error) {
	if source.File == nil {
		return nil, errors.String("file config is nil")
	}
	path := WorkPath("", source.File.Path)
	log.Infof("loading config from %s", path)
	stat, err := os.Stat(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to state file %s", path)
	}
	if stat.IsDir() {
		return nil, errors.String("file config is a directory")
	}
	return LoadFileBootstrap(path)
}

// LoadRemoteBootstrap load config from source
func LoadRemoteBootstrap(source *configv1.SourceConfig) (*configs.Bootstrap, error) {
	config, err := runtime.NewConfig(source)
	if err != nil {
		return nil, err
	}
	if err := config.Load(); err != nil {
		return nil, err
	}
	var cfg configs.Bootstrap
	if err := config.Scan(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// LoadBootstrap load config from file
func LoadBootstrap(flags *Bootstrap) (*configs.Bootstrap, error) {
	fmt.Println("load config from: ", flags.WorkPath())
	sourceConfig, err := bootstrap.LoadSourceConfig(flags)
	if err != nil {
		return nil, err
	}
	log.Infof("load soure config: %+v", sourceConfig)
	switch sourceConfig.GetType() {
	case "file":
		return LoadLocalBootstrap(sourceConfig)
	default:
		return LoadRemoteBootstrap(sourceConfig)
	}
}

// LoadRemoteServiceBootstrap get the configuration from the remote Configuration Center
func LoadRemoteServiceBootstrap(flags *Bootstrap, name string, ss ...ConfigSetting) (*configs.Bootstrap, error) {
	var cfg configs.Bootstrap
	sourceConfig, err := bootstrap.LoadSourceConfig(flags)
	if err != nil {
		return nil, err
	}

	sourceConfig.Name = name
	sourceConfig = settings.Apply(sourceConfig, ss)
	config, err := runtime.NewConfig(sourceConfig)
	if err != nil {
		return nil, err
	}
	if err := config.Load(); err != nil {
		return nil, err
	}
	if err := config.Scan(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
