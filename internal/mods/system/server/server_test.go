/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package server implements the functions, types, and interfaces for the module.
package server

import (
	"reflect"
	"testing"

	_ "github.com/origadmin/contrib/consul/config"
	_ "github.com/origadmin/contrib/consul/registry"
	_ "github.com/origadmin/contrib/database"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/service"
	servicegrpc "github.com/origadmin/runtime/service/grpc"

	systemservice "origadmin/application/admin/internal/mods/system/service"
)

func TestNewLoginServerAgent(t *testing.T) {
	type args struct {
		client *service.GRPCClient
	}

	serviceConfig := &configv1.Service{
		Name: ServiceName,
		//Grpc: entry.GetGrpc(),
		//Http: entry.GetHttp(),
		//Gins: entry.GetGins(),
		Selector: &configv1.Service_Selector{
			Version: "v1.0.0",
			Builder: "bbr",
		},
	}
	discovery, err := runtime.NewDiscovery(&configv1.Registry{
		Type:   "consul",
		Consul: &configv1.Registry_Consul{},
	})
	if err != nil {
		t.Fatalf("create discovery: %v", err)
	}

	client, err := runtime.NewGRPCServiceClient(context.Background(), serviceConfig, service.WithGRPC(func(o *servicegrpc.Option) {
		o.Discovery = discovery
		o.ServiceName = ServiceName
	}))
	if err != nil {
		t.Fatalf("create login grpc client: %v", err)
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				client: client,
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//agent := NewLoginServerAgent(tt.args.client)
			//agent.CaptchaID(nil, &dto.CaptchaIDRequest{Id: "123"})
			if got := systemservice.NewLoginServerAgentGINRegister(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginServerAgent() = %v, want %v", got, tt.want)
			}
		})
	}
}
