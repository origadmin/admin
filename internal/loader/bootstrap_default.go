/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"github.com/origadmin/runtime/gen/go/config/v1"
	"google.golang.org/protobuf/types/known/durationpb"

	"origadmin/application/admin/internal/configs"
)

func DefaultBootstrap() *configs.Bootstrap {
	return &configs.Bootstrap{
		Name:       "origadmin.service.v1.demo",
		Version:    "v1.0.0",
		CryptoType: "argon2",

		//Server: &Server{
		//	Entry: &Server_Entry{
		//		Network: "tcp",
		//		Addr:    "0.0.0.0:8000",
		//		Timeout: durationpb.New(3 * time.Minute),
		//	},
		//	Gins: &Server_GINS{
		//		Network:         "tcp",
		//		Addr:            "${gins_address=0.0.0.0:8100}",
		//		UseTls:          false,
		//		CertFile:        "",
		//		KeyFile:         "",
		//		Timeout:         durationpb.New(3 * time.Minute),
		//		ShutdownTimeout: durationpb.New(3 * time.Minute),
		//		ReadTimeout:     durationpb.New(3 * time.Minute),
		//		WriteTimeout:    durationpb.New(3 * time.Minute),
		//		IdleTimeout:     durationpb.New(3 * time.Minute),
		//	},
		//	Http: &Server_HTTP{
		//		Network:         "tcp",
		//		Addr:            "${http_address=0.0.0.0:8200}",
		//		UseTls:          false,
		//		CertFile:        "",
		//		KeyFile:         "",
		//		Timeout:         durationpb.New(3 * time.Minute),
		//		ShutdownTimeout: durationpb.New(3 * time.Minute),
		//		ReadTimeout:     durationpb.New(3 * time.Minute),
		//		WriteTimeout:    durationpb.New(3 * time.Minute),
		//		IdleTimeout:     durationpb.New(3 * time.Minute),
		//	},
		//	Grpc: &Server_GRPC{
		//		Network:         "tcp",
		//		Addr:            "${grpc_address=0.0.0.0:8300}",
		//		UseTls:          false,
		//		CertFile:        "",
		//		KeyFile:         "",
		//		Timeout:         durationpb.New(3 * time.Minute),
		//		ShutdownTimeout: durationpb.New(3 * time.Minute),
		//		ReadTimeout:     durationpb.New(3 * time.Minute),
		//		WriteTimeout:    durationpb.New(3 * time.Minute),
		//		IdleTimeout:     durationpb.New(3 * time.Minute),
		//	},
		//	Middleware: &Server_Middleware{
		//		Cors: &Server_Middleware_Cors{
		//			Enabled:                false,
		//			AllowAllOrigins:        false,
		//			AllowOrigins:           nil,
		//			AllowMethods:           nil,
		//			AllowHeaders:           nil,
		//			AllowCredentials:       false,
		//			ExposeHeaders:          nil,
		//			MaxAge:                 0,
		//			AllowWildcard:          false,
		//			AllowBrowserExtensions: false,
		//			AllowWebSockets:        false,
		//			AllowFiles:             false,
		//		},
		//		Metrics: &Server_Middleware_Metrics{
		//			Enabled: false,
		//			Name:    "metrics",
		//		},
		//		Traces: &Server_Middleware_Traces{
		//			Enabled: false,
		//			Name:    "traces",
		//		},
		//		Logger: &Server_Middleware_Logger{
		//			Enabled: false,
		//			Name:    "logger",
		//		},
		//	},
		//	Host: "${host=127.0.0.1}",
		//},
		//Data: &Data{
		//	Database: &Data_Database{
		//		Driver: "mysql",
		//		Source: "dsn",
		//	},
		//	Redis: &Data_Redis{
		//		Network:      "tcp",
		//		Addr:         "${redis_address=127.0.0.1:6379}",
		//		ReadTimeout:  durationpb.New(3 * time.Minute),
		//		WriteTimeout: durationpb.New(3 * time.Minute),
		//	},
		//},
		//Config: &Config{
		//	Type: "file",
		//	Consul: &Consul{
		//		Address: "${consul_address=127.0.0.1:8500}",
		//		Scheme:  "http",
		//	},
		//},
		//Discovery: &Discovery{
		//	Type: "${discovery_type:consul}",
		//	Consul: &Consul{
		//		Address: "${consul_address=127.0.0.1:8500}",
		//		//Scheme:              "",
		//		//SessionToken:               "",
		//		//Datacenter:          "",
		//		//Tag:                 "",
		//		//HealthCheckInterval: "",
		//		//HealthCheckTimeout:  "",
		//		HealthCheck: true,
		//		HeartBeat:   true,
		//	},
		//	Etcd: &Etcd{
		//		Endpoints: "${etcd_address=127.0.0.1:2379}",
		//	},
		//},
		//Settings: &Settings{
		//	CryptoType: "argon2",
		//},
	}

}

func DefaultServiceGins() *configv1.Service_GINS {
	return &configv1.Service_GINS{
		Network: "",
		Addr:    "",
		//UseTls:   false,
		//CertFile: "",
		//KeyFile:  "",
		Timeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		ShutdownTimeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		ReadTimeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		WriteTimeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		IdleTimeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		Endpoint: "",
	}
}

func DefaultServiceGrpc() *configv1.Service_GRPC {
	return &configv1.Service_GRPC{
		Network:  "",
		Addr:     "",
		UseTls:   false,
		CertFile: "",
		KeyFile:  "",
		Timeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		ShutdownTimeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		ReadTimeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		WriteTimeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		IdleTimeout: &durationpb.Duration{
			Seconds: 0,
			Nanos:   0,
		},
		Endpoint: "",
	}
}

func DefaultEntry() *configs.Bootstrap_Entry {
	return &configs.Bootstrap_Entry{
		Name: "application",
	}
}
