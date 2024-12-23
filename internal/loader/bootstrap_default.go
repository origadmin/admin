/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"time"

	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	jwtv1 "github.com/origadmin/runtime/gen/go/middleware/jwt/v1"
	"github.com/origadmin/runtime/gen/go/middleware/metrics/v1"
	"github.com/origadmin/runtime/gen/go/middleware/ratelimit/v1"
	"github.com/origadmin/runtime/gen/go/middleware/selector/v1"
	middlewarev1 "github.com/origadmin/runtime/gen/go/middleware/v1"
	"github.com/origadmin/runtime/gen/go/middleware/validator/v1"
	sjwtv1 "github.com/origadmin/runtime/gen/go/security/jwt/v1"

	"origadmin/application/admin/internal/configs"
	basisserver "origadmin/application/admin/internal/mods/basis/server"
	systemserver "origadmin/application/admin/internal/mods/system/server"
)

const (
	SigningKey = "%VH_C!Vpa$_aK2kOynB&q+x=4$27&Ios"
)

func DefaultBootstrap() *configs.Bootstrap {

	return &configs.Bootstrap{
		Name:       "origadmin.agent.service.admin.v1",
		Mode:       "singleton",
		Version:    "v1.0.0",
		CryptoType: "argon2",
		Servers: map[string]string{
			basisserver.ServiceName:  "origadmin.service.basis.v1",
			systemserver.ServiceName: "origadmin.service.system.v1",
		},
		Id: "",
		Entry: &configs.Bootstrap_Entry{
			Scheme: "http",
		},
		Service: &configv1.Service{
			Name:            "",
			DynamicEndpoint: true,
			Grpc:            DefaultServiceGrpc(),
			Http:            DefaultServiceHttp(),
			Websocket:       DefaultServiceWebsocket(),
			Message:         DefaultServiceMessage(),
			Task:            DefaultServiceTask(),
			Middleware:      DefaultServiceMiddleware(),
			Selector: &configv1.Service_Selector{
				Version: "v1.0.0",
				Builder: "bbr",
			},
		},
		Data:       DefaultData(),
		Registry:   DefaultRegistry(),
		Middleware: DefaultServiceMiddleware(),
		Security: &configv1.Security{
			PublicPaths: []string{
				"/swagger/*",
				"/api/v1/health",
				"/api/v1/health/*",
				"/api/v1/captcha",
				"/api/v1/captcha/*",
				"/api/v1/login",
				"/api/v1/register",
				"/api/v1/current/logout",
				"/api/v1/refresh_token",
				"/api.v1.services.basis.LoginAPI/CaptchaID",
				"/api.v1.services.basis.LoginAPI/CaptchaImage",
				"/api.v1.services.basis.LoginAPI/CaptchaResource",
				"/api.v1.services.basis.LoginAPI/CaptchaResources",
				"/api.v1.services.basis.LoginAPI/Login",
				"/api.v1.services.basis.LoginAPI/Register",
				"/api.v1.services.basis.LoginAPI/Refresh",
				//"/api.v1.services.basis.LoginAPI/Logout",
				//"/api.v1.services.basis.LoginAPI/CurrentUser",
				//"/api.v1.services.basis.LoginAPI/CurrentMenus",
			},
			Authz: &configv1.AuthZConfig{
				Disabled:    false,
				PublicPaths: nil,
				Type:        "casbin",
				Casbin: &configv1.AuthZConfig_CasbinConfig{
					PolicyFile: "",
					ModelFile:  "",
				},
				Opa:      nil,
				Zanzibar: nil,
			},
			Authn: &configv1.AuthNConfig{
				Disabled: false,
				Type:     "jwt",
				Jwt: &configv1.AuthNConfig_JWTConfig{
					Algorithm:     "HS512",
					SigningKey:    SigningKey,
					OldSigningKey: "",
					ExpireTime:    0, // use default
					RefreshTime:   0, // use default
					CacheName:     "",
				},
			},
		},
	}
}

func DefaultServiceWebsocket() *configv1.WebSocket {
	return &configv1.WebSocket{
		Addr: "",
		Path: "",
	}
}

func DefaultData() *configv1.Data {
	return &configv1.Data{
		Database: &configv1.Data_Database{
			Debug:                 false,
			Driver:                "sqlite3",
			Source:                "data/admin.db",
			Migrate:               false,
			EnableTrace:           false,
			EnableMetrics:         false,
			MaxIdleConnections:    0,
			MaxOpenConnections:    0,
			ConnectionMaxLifetime: 0,
			ConnectionMaxIdleTime: 0,
		},
		Cache: &configv1.Data_Cache{
			Driver: "memory", //["none", "redis", "memcached", "memory"] [string.in]
			Memcached: &configv1.Data_Memcached{
				Addr:     "",
				Username: "",
				Password: "",
				MaxIdle:  0,
				Timeout:  0,
			},
			Memory: &configv1.Data_Memory{
				Size:            0,
				Capacity:        0,
				Expiration:      0,
				CleanupInterval: 0,
			},
			Redis: &configv1.Data_Redis{
				Network:      "",
				Addr:         "",
				Password:     "",
				Db:           0,
				DialTimeout:  0,
				ReadTimeout:  0,
				WriteTimeout: 0,
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
				Network:      "",
				Addr:         "",
				Password:     "",
				Db:           0,
				DialTimeout:  0,
				ReadTimeout:  0,
				WriteTimeout: 0,
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
			Timeout:                        0,
			DeregisterCriticalServiceAfter: 0,
		},
		Etcd: nil,
	}
}

func DefaultServiceMiddleware() *middlewarev1.Middleware {
	return &middlewarev1.Middleware{
		Logging:        true,
		Recovery:       true,
		Tracing:        true,
		CircuitBreaker: true,
		Metadata: &middlewarev1.Middleware_Metadata{
			Enabled: true,
		},
		RateLimiter: &ratelimitv1.RateLimiter{
			Enabled: true,
			Name:    "bbr",
		},
		Metrics: &metricsv1.Metrics{
			Enabled: true,
		},
		Validator: &validatorv1.Validator{
			Enabled:  true,
			Version:  1,
			FailFast: true,
		},
		//It is not recommended to use integrated JWT components
		Jwt: &jwtv1.JWT{
			Enabled: false,
			Config: &sjwtv1.Config{
				SigningMethod:        "HS512",
				Key:                  SigningKey,
				Key2:                 "can empty next version fixed",
				AccessTokenLifetime:  int64(15 * time.Minute),
				RefreshTokenLifetime: int64(3 * 24 * time.Hour),
				Issuer:               "localhost",
				Audience:             nil,
				TokenType:            "Bearer",
			},
		},
		// Middleware filters
		Selector: &selectorv1.Selector{
			Enabled: false,
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
		Timeout:         0,
		ShutdownTimeout: 0,
		ReadTimeout:     0,
		WriteTimeout:    0,
		IdleTimeout:     0,
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
		Timeout:         0,
		ShutdownTimeout: 0,
		ReadTimeout:     0,
		WriteTimeout:    0,
		IdleTimeout:     0,
		Endpoint:        "",
	}
}

//func DefaultServiceGins() *configv1.Service_GINS {
//	return &configv1.Service_GINS{
//		Network:         "tcp",
//		Addr:            "${gins_address:0.0.0.0:18200}",
//		UseTls:          false,
//		CertFile:        "",
//		KeyFile:         "",
//		Timeout:         durationpb.New(time.Minute * 3),
//		ShutdownTimeout: durationpb.New(time.Minute * 3),
//		ReadTimeout:     durationpb.New(time.Minute * 3),
//		WriteTimeout:    durationpb.New(time.Minute * 3),
//		IdleTimeout:     durationpb.New(time.Minute * 3),
//		Endpoint:        "",
//	}
//}

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

func DefaultRootUser() *configs.RootUser {
	return &configs.RootUser{
		Enabled:        true,
		Username:       "admin",
		Password:       "admin",
		RandomPassword: true,
		Name:           "admin",
		Nickname:       "admin",
		Email:          "admin@admin.com",
		Mobile:         "1380000000",
	}
}

func DefaultBasisConfig() *configs.BasisConfig {
	return &configs.BasisConfig{
		RootUser: DefaultRootUser(),
		Captcha:  DefaultCaptcha(),
	}
}
