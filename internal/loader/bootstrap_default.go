/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"time"

	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"google.golang.org/protobuf/types/known/durationpb"

	"origadmin/application/admin/internal/configs"
	basisserver "origadmin/application/admin/internal/mods/basis/server"
	systemserver "origadmin/application/admin/internal/mods/system/server"
)

func DefaultBootstrap() *configs.Bootstrap {
	return &configs.Bootstrap{
		Name:       "origadmin.service.v1.admin",
		Mode:       "singleton",
		Version:    "v1.0.0",
		CryptoType: "argon2",
		Servers: map[string]string{
			basisserver.ServiceName:  "origadmin.service.v1.basis",
			systemserver.ServiceName: "origadmin.service.v1.system",
		},
		Entry: &configs.Bootstrap_Entry{
			Scheme: "http",
		},
		Service: &configv1.Service{
			Name: "",
			Grpc: DefaultServiceGrpc(),
			Http: DefaultServiceHttp(),
			Gins: DefaultServiceGins(),
			Websocket: &configv1.WebSocket{
				Network: "",
				Addr:    "",
				Path:    "",
				Codec:   "",
				Timeout: &durationpb.Duration{
					Seconds: 0,
					Nanos:   0,
				},
			},
			Message:    DefaultServiceMessage(),
			Task:       DefaultServiceTask(),
			Middleware: DefaultServiceMiddleware(),
			Selector: &configv1.Service_Selector{
				Version: "v1.0.0",
				Builder: "bbr",
			},
			Host: "ORIGADMIN_SERVICE_HOST",
		},
		Data:     DefaultData(),
		Registry: DefaultRegistry(),
		Middlewares: &configs.Middlewares{
			RegisterAsGlobal: false,
			Middleware:       DefaultServiceMiddleware(),
		},
		Id: "",
	}

}

func DefaultData() *configv1.Data {
	return &configv1.Data{
		Database: &configv1.Data_Database{
			Debug:              false,
			Driver:             "sqlite3",
			Source:             "data/admin.db",
			Migrate:            false,
			EnableTrace:        false,
			EnableMetrics:      false,
			MaxIdleConnections: 0,
			MaxOpenConnections: 0,
			ConnectionMaxLifetime: &durationpb.Duration{
				Seconds: 0,
				Nanos:   0,
			},
			ConnectionMaxIdleTime: &durationpb.Duration{
				Seconds: 0,
				Nanos:   0,
			},
		},
		Cache: &configv1.Data_Cache{
			Driver: "memory", //["none", "redis", "memcached", "memory"] [string.in]
			Memcached: &configv1.Data_Memcached{
				Addr:     "",
				Username: "",
				Password: "",
				MaxIdle:  0,
				Timeout: &durationpb.Duration{
					Seconds: 0,
					Nanos:   0,
				},
			},
			Memory: &configv1.Data_Memory{
				Size:     0,
				Capacity: 0,
				Expiration: &durationpb.Duration{
					Seconds: 0,
					Nanos:   0,
				},
				CleanupInterval: &durationpb.Duration{
					Seconds: 0,
					Nanos:   0,
				},
			},
			Redis: &configv1.Data_Redis{
				Network:  "",
				Addr:     "",
				Password: "",
				Db:       0,
				DialTimeout: &durationpb.Duration{
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
			},
			Badger: &configv1.Data_BadgerDS{
				Path:             "",
				SyncWrites:       false,
				ValueLogFileSize: 0,
				LogLevel:         0,
			},
		},
		Storage: &configv1.Data_Storage{
			Type: "none", //["none", "file", "redis", "mongo", "oss"] [string.in]
			File: &configv1.Data_File{
				Root: "",
			},
			Redis: &configv1.Data_Redis{
				Network:  "",
				Addr:     "",
				Password: "",
				Db:       0,
				DialTimeout: &durationpb.Duration{
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
			},
			Badger: &configv1.Data_BadgerDS{
				Path:             "",
				SyncWrites:       false,
				ValueLogFileSize: 0,
				LogLevel:         0,
			},
			Mongo: &configv1.Data_Mongo{},
			Oss:   &configv1.Data_Oss{},
		},
	}
}

func DefaultServiceTask() *configv1.Task {
	return &configv1.Task{
		Type: "none", //["none", "asynq", "machinery", "cron"] [string.in]
		Name: "",
		Asynq: &configv1.Task_Asynq{
			Endpoint: "",
			Password: "",
			Db:       0,
			Location: "",
		},
		Machinery: &configv1.Task_Machinery{
			Brokers:  nil,
			Backends: nil,
		},
		Cron: &configv1.Task_Cron{
			Addr: "",
		},
	}
}

func DefaultServiceMessage() *configv1.Message {
	return &configv1.Message{
		Type: "none", //["none", "mqtt", "kafka", "rabbitmq", "activemq", "nats", "nsq", "pulsar", "redis", "rocketmq"]
		Name: "",
		Mqtt: &configv1.Message_MQTT{
			Endpoint: "",
			Codec:    "",
		},
		Kafka: &configv1.Message_Kafka{
			Endpoint: "",
			Codec:    "",
		},
		Rabbitmq: &configv1.Message_RabbitMQ{
			Endpoint: "",
			Codec:    "",
		},
		Activemq: &configv1.Message_ActiveMQ{
			Endpoint: "",
			Codec:    "",
		},
		Nats: &configv1.Message_NATS{
			Endpoint: "",
			Codec:    "",
		},
		Nsq: &configv1.Message_NSQ{
			Endpoint: "",
			Codec:    "",
		},
		Pulsar: &configv1.Message_Pulsar{
			Endpoint: "",
			Codec:    "",
		},
		Redis: &configv1.Message_Redis{
			Endpoint: "",
			Codec:    "",
		},
		Rocketmq: &configv1.Message_RocketMQ{
			Endpoint:         "",
			Codec:            "",
			EnableTrace:      false,
			NameServers:      nil,
			NameServerDomain: "",
			AccessKey:        "",
			SecretKey:        "",
			SecurityToken:    "",
			Namespace:        "",
			InstanceName:     "",
			GroupName:        "",
		},
	}
}

func DefaultRegistry() *configv1.Registry {
	return &configv1.Registry{
		Debug: false,
		Type:  "consul",
		Consul: &configv1.Registry_Consul{
			Address:                        "${consul_address:127.0.0.1:8500}",
			Scheme:                         "http",
			Token:                          "",
			HeartBeat:                      true,
			HealthCheck:                    true,
			Datacenter:                     "",
			HealthCheckInterval:            30,
			Timeout:                        nil,
			DeregisterCriticalServiceAfter: 0,
		},
		Etcd: nil,
	}
}

func DefaultServiceMiddleware() *configv1.Middleware {
	return &configv1.Middleware{
		EnableLogging:        true,
		EnableRecovery:       true,
		EnableTracing:        true,
		EnableValidate:       true,
		EnableCircuitBreaker: true,
		EnableMetadata:       true,
		RateLimiter: &configv1.Middleware_RateLimiter{
			Name:                "bbr",
			Period:              0,
			XRatelimitLimit:     0,
			XRatelimitRemaining: 0,
			XRatelimitReset:     0,
			RetryAfter:          0,
			Memory: &configv1.Middleware_RateLimiter_Memory{
				Expiration: &durationpb.Duration{
					Seconds: 0,
					Nanos:   0,
				},
				CleanupInterval: &durationpb.Duration{
					Seconds: 0,
					Nanos:   0,
				},
			},
			Redis: &configv1.Middleware_RateLimiter_Redis{
				Addr:     "",
				Username: "",
				Password: "",
				Db:       0,
			},
		},
		Metadata: &configv1.Middleware_Metadata{
			Prefix: "",
			Data:   nil,
		},
		Metrics: &configv1.Middleware_Metrics{
			SupportedMetrics: nil,
			UserMetrics:      nil,
		},
		Validator: &configv1.Middleware_Validator{
			Version:  2,
			FailFast: true,
		},
	}
}

func DefaultServiceGrpc() *configv1.Service_GRPC {
	return &configv1.Service_GRPC{
		Network:         "tcp",
		Addr:            "${grpc_address:0.0.0.0:18000}",
		UseTls:          false,
		CertFile:        "",
		KeyFile:         "",
		Timeout:         durationpb.New(time.Minute * 3),
		ShutdownTimeout: durationpb.New(time.Minute * 3),
		ReadTimeout:     durationpb.New(time.Minute * 3),
		WriteTimeout:    durationpb.New(time.Minute * 3),
		IdleTimeout:     durationpb.New(time.Minute * 3),
		Endpoint:        "",
	}
}

func DefaultServiceHttp() *configv1.Service_HTTP {
	return &configv1.Service_HTTP{
		Network:         "tcp",
		Addr:            "${http_address:0.0.0.0:18100}",
		UseTls:          false,
		CertFile:        "",
		KeyFile:         "",
		Timeout:         durationpb.New(time.Minute * 3),
		ShutdownTimeout: durationpb.New(time.Minute * 3),
		ReadTimeout:     durationpb.New(time.Minute * 3),
		WriteTimeout:    durationpb.New(time.Minute * 3),
		IdleTimeout:     durationpb.New(time.Minute * 3),
		Endpoint:        "",
	}
}

func DefaultServiceGins() *configv1.Service_GINS {
	return &configv1.Service_GINS{
		Network:         "tcp",
		Addr:            "${gins_address:0.0.0.0:18200}",
		UseTls:          false,
		CertFile:        "",
		KeyFile:         "",
		Timeout:         durationpb.New(time.Minute * 3),
		ShutdownTimeout: durationpb.New(time.Minute * 3),
		ReadTimeout:     durationpb.New(time.Minute * 3),
		WriteTimeout:    durationpb.New(time.Minute * 3),
		IdleTimeout:     durationpb.New(time.Minute * 3),
		Endpoint:        "",
	}
}

func DefaultEntry() *configs.Bootstrap_Entry {
	return &configs.Bootstrap_Entry{
		Scheme: "http",
	}
}

func DefaultCaptcha() *configs.Captcha {
	return &configs.Captcha{
		CacheType: "memory",
		Width:     400,
		Height:    160,
		Length:    4,
		Redis: &configs.Captcha_Redis{
			Addr:      "${captcha_redis_address:127.0.0.1:6379}",
			Db:        0,
			KeyPrefix: "captcha",
		},
	}
}
