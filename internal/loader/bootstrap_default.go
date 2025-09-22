package loader

import (
	"time"

	// configv1 "github.com/origadmin/runtime/api/gen/go/config/v1" // REMOVE
	middlewarev1 "github.com/origadmin/runtime/api/gen/go/middleware/v1"
	jwtv1 "github.com/origadmin/runtime/api/gen/go/middleware/v1/jwt"
	"github.com/origadmin/runtime/api/gen/go/middleware/v1/metrics"
	"github.com/origadmin/runtime/api/gen/go/middleware/v1/ratelimit"
	"github.com/origadmin/runtime/api/gen/go/middleware/v1/selector"
	"github.com/origadmin/runtime/api/gen/go/middleware/v1/validator"
	sjwtv1 "github.com/origadmin/runtime/api/gen/go/security/jwt/v1"
	securityv1 "github.com/origadmin/runtime/api/gen/go/security/transport/v1" // ADD for TLSConfig
	transportv1 "github.com/origadmin/runtime/api/gen/go/transport/v1" // ADD

	"origadmin/application/admin/internal/configs"
	"google.golang.org/protobuf/types/known/durationpb"
)

const (
	SigningKey = "%VH_C!Vpa$_aK2kOynB&q+x=4$27&Ios"
)

func DefaultBootstrap() *configs.Bootstrap {
	return &configs.Bootstrap{
		Name:       "origadmin.service.admin.v1",
		Mode:       "singleton",
		Version:    "v1.0.0",
		CryptoType: "argon2",
		//Servers: map[string]string{
		//	systemserver.ServiceName: "origadmin.service.system.v1",
		//},
		Id: "",
		Entry: &configs.Bootstrap_Entry{
			Scheme:   "http",
			Services: DefaultServices(), // This will cause a type mismatch after this change
			Cors:     DefaultEntryCors(),
		},
		Server: &configs.ServiceServer{
			Services:   DefaultServices(), // This will cause a type mismatch after this change
			Middleware: DefaultServiceMiddleware(),
		},
		Clients:    DefaultServiceClients(),
		Logger:     DefaultLogger(),
		Storage:    DefaultStorage(),
		Discovery:  DefaultDiscovery(),
		Middleware: DefaultServiceMiddleware(),
		Security: &configs.SecurityConfig{
			RootUser: DefaultRootUser(),
			Captcha:  DefaultCaptcha(),
			Security: &configs.Security{
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
					"/api.v1.services.system.LoginAPI/CaptchaId",
					"/api.v1.services.system.LoginAPI/CaptchaImage",
					"/api.v1.services.system.LoginAPI/CaptchaResource",
					"/api.v1.services.system.LoginAPI/CaptchaResources",
					"/api.v1.services.system.LoginAPI/Login",
					"/api.v1.services.system.LoginAPI/Register",
					"/api.v1.services.system.LoginAPI/Refresh",
					//"/api.v1.services.basis.LoginAPI/Logout",
					//"/api.v1.services.basis.LoginAPI/CurrentUser",
					//"/api.v1.services.basis.LoginAPI/CurrentMenus",
				},
				Authz: &configs.AuthZConfig{
					Disabled:    false,
					PublicPaths: nil,
					Type:        "casbin",
					Casbin: &configs.AuthZConfig_CasbinConfig{
						PolicyFile: "",
						ModelFile:  "",
					},
					Opa:      nil,
					Zanzibar: nil,
				},
				Authn: &configs.AuthNConfig{
					Disabled: false,
					Type:     "jwt",
					Jwt: &configs.AuthNConfig_JWTConfig{
						Algorithm:     "HS512",
						SigningKey:    SigningKey,
						OldSigningKey: "",
						ExpireTime:    0, // use default
						RefreshTime:   0, // use default
						CacheName:     "",
					},
				},
			},
		},
	}
}

func DefaultEntryCors() *configs.Cors {
	return &configs.Cors{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           0,
	}
}

func DefaultServices() []*configs.Service {
	return []*configs.Service{
		{
			Name:            "",
			DynamicEndpoint: true,
			Type:            "grpc",
			Grpc:            DefaultServiceGrpc(), // This will cause a type mismatch
			Websocket:       DefaultServiceWebsocket(),
			Message:         DefaultServiceMessage(),
			Task:            DefaultServiceTask(),
			//Middleware:      DefaultServiceMiddleware(),
			Selector: &configs.Service_Selector{
				Version: "v1.0.0",
				Builder: "bbr",
			},
		},
		{
			Name:            "",
			DynamicEndpoint: true,
			Type:            "http",
			Http:            DefaultServiceHttp(), // This will cause a type mismatch
			Websocket:       DefaultServiceWebsocket(),
			Message:         DefaultServiceMessage(),
			Task:            DefaultServiceTask(),
			//Middleware:      DefaultServiceMiddleware(),
			Selector: &configs.Service_Selector{
				Version: "v1.0.0",
				Builder: "bbr",
			},
		},
	}
}

func DefaultServiceClients() []*configs.ServiceClient {
	serviceNames := map[string]string{
		"system": "origadmin.service.system.v1",
		"auth":   "origadmin.service.auth.v1",
	}
	clients := make([]*configs.ServiceClient, 0, len(serviceNames))
	for name, serviceName := range serviceNames {
		core := &configs.ServiceCore{
			Name:      name,
			Discovery: DefaultDiscovery(),
		}
		core.Discovery.ServiceName = serviceName
		clients = append(clients, &configs.ServiceClient{
			Core:       core,
			Services:   DefaultServices(), // This will cause a type mismatch after this change
			Middleware: DefaultServiceMiddleware(),
		})
	}
	return clients
}

func DefaultLogger() *configs.Logger {
	return &configs.Logger{
		Disabled:      false,
		Develop:       true,
		Default:       true,
		Name:          "output.log",
		Format:        "json",
		Level:         "info",
		Stdout:        true,
		DisableCaller: false,
		CallerSkip:    0,
		TimeFormat:    "",
		File: &configs.Logger_File{
			Path:       "logs",
			Lumberjack: true,
			Compress:   false,
			LocalTime:  false,
			MaxSize:    0,
			MaxAge:     0,
			MaxBackups: 0,
		},
		DevLogger: nil,
	}
}

func DefaultServiceWebsocket() *configs.WebSocket {
	return &configs.WebSocket{
		Addr: "",
		Path: "",
	}
}

func DefaultStorage() *configs.Storage {
	return &configs.Storage{
		Name: "",
		Type: "",
		Database: &configs.Database{
			Debug:   false,
			Dialect: "sqlite3",
			Source:  "data/admin.db",
			Migration: &configs.Migration{
				Enabled: false,
				Path:    "",
				Names:   nil,
				Version: "",
				Mode:    "",
			},
			EnableTrace:           false,
			EnableMetrics:         false,
			MaxIdleConnections:    0,
			MaxOpenConnections:    0,
			ConnectionMaxLifetime: 0,
			ConnectionMaxIdleTime: 0,
		},
		Cache: &configs.Cache{
			Driver: "memory", //["none", "redis", "memcached", "memory"] [string.in]
			Memcached: &configs.Memcached{
				Addr:     "",
				Username: "",
				Password: "",
				MaxIdle:  0,
				Timeout:  0,
			},
			Memory: &configs.Memory{
				Size:            0,
				Capacity:        0,
				Expiration:      0,
				CleanupInterval: 0,
			},
			Redis: &configs.Redis{
				Network:      "",
				Addr:         "",
				Password:     "",
				Db:           0,
				DialTimeout:  0,
				ReadTimeout:  0,
				WriteTimeout: 0,
			},
			Badger: &configs.BadgerDS{
				Path:             "",
				SyncWrites:       false,
				ValueLogFileSize: 0,
				LogLevel:         0,
			},
		},
		File:   nil,
		Redis:  nil,
		Badger: nil,
		Mongo:  nil,
		Oss:    nil,
	}
}

func DefaultServiceTask() *configs.Task {
	return &configs.Task{
		Type: "none", //["none", "asynq", "machinery", "cron"] [string.in]
		Name: "",
		Asynq: &configs.Task_Asynq{
			Endpoint: "",
			Password: "",
			Db:       0,
			Location: "",
		},
		Machinery: &configs.Task_Machinery{
			Brokers:  nil,
			Backends: nil,
		},
		Cron: &configs.Task_Cron{
			Addr: "",
		},
	}
}

func DefaultServiceMessage() *configs.Message {
	return &configs.Message{
		Type: "none", //["none", "mqtt", "kafka", "rabbitmq", "activemq", "nats", "nsq", "pulsar", "redis", "rocketmq"]
		Name: "",
		Mqtt: &configs.Message_MQTT{
			Endpoint: "",
			Codec:    "",
		},
		Kafka: &configs.Message_Kafka{
			Endpoint: "",
			Codec:    "",
		},
		Rabbitmq: &configs.Message_RabbitMQ{
			Endpoint: "",			Codec:    "",
		},
		Activemq: &configs.Message_ActiveMQ{
			Endpoint: "",
			Codec:    "",
		},
		Nats: &configs.Message_NATS{
			Endpoint: "",
			Codec:    "",
		},
		Nsq: &configs.Message_NSQ{
			Endpoint: "",
			Codec:    "",
		},
		Pulsar: &configs.Message_Pulsar{
			Endpoint: "",
			Codec:    "",
		},
		Redis: &configs.Message_Redis{
			Endpoint: "",
			Codec:    "",
		},
		Rocketmq: &configs.Message_RocketMQ{
			Endpoint:         "",
			Codec:            "",
			EnableTrace:      false,
			NameServers:      nil,
			NameServerDomain: "",
			AccessKey:        "",
			SecretKey:        "",			SecurityToken:    "",
			Namespace:        "",
			InstanceName:     "",
			GroupName:        "",
		},
	}
}

func DefaultDiscovery() *configs.Discovery {
	return &configs.Discovery{
		Debug: false,
		Type:  "consul",
		Consul: &configs.Discovery_Consul{
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
		EnabledMiddlewares: []string{
			"logging",
			"recovery",
			"tracing",
			"circuit_breaker",
			"metadata",
			"rate_limiter",
			"metrics",
			"validator",
			"jwt",
			"selector",
		},
		//Logging:        true,
		//Recovery:       true,
		//Tracing:        true,
		//CircuitBreaker: true,
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

func DefaultServiceGrpc() *transportv1.GRPCServer {
	return &transportv1.GRPCServer{
		Network: "tcp",
		Addr:    "${grpc_address:0.0.0.0:18000}",
		Tls:     &securityv1.TLSConfig{Enabled: false}, // Replaced UseTls
		Timeout: &durationpb.Duration{}, // Use durationpb.Duration
		ShutdownTimeout: &durationpb.Duration{}, // Use durationpb.Duration
		Endpoint:        "",
	}
}

func DefaultServiceHttp() *transportv1.HTTPServer {
	return &transportv1.HTTPServer{
		Network: "tcp",
		Addr:    "${http_address:0.0.0.0:18100}",
		Tls:     &securityv1.TLSConfig{Enabled: false}, // Replaced UseTls
		Timeout: &durationpb.Duration{}, // Use durationpb.Duration
		ShutdownTimeout: &durationpb.Duration{}, // Use durationpb.Duration
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
		Length:      4,
		Width:       400,
		Height:      160,
		StorageName: "captcha",
		Storage:     &configs.Storage{},
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
