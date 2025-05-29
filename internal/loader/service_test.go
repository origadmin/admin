/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"testing"

	configv1 "github.com/origadmin/runtime/api/gen/go/config/v1"
	v11 "github.com/origadmin/runtime/api/gen/go/middleware/v1"
	jwtv1 "github.com/origadmin/runtime/api/gen/go/middleware/v1/jwt"
	metricsv1 "github.com/origadmin/runtime/api/gen/go/middleware/v1/metrics"
	ratelimitv1 "github.com/origadmin/runtime/api/gen/go/middleware/v1/ratelimit"
	selectorv1 "github.com/origadmin/runtime/api/gen/go/middleware/v1/selector"
	validatorv1 "github.com/origadmin/runtime/api/gen/go/middleware/v1/validator"

	"github.com/stretchr/testify/assert"

	"origadmin/application/admin/internal/configs"
)

func TestServiceDefaultOutput(t *testing.T) {
	// Test the default service configuration initialization
	ss := make([]*configs.ServiceServer, 0)

	// Verify empty slice initialization
	assert.Empty(t, ss, "默认服务列表应为空")

	// 添加测试服务实例
	testService := &configs.ServiceServer{
		Core: &configs.ServiceCore{
			Name:    "test-service",
			Version: "v1.0.0",
			Discovery: &configv1.Discovery{
				Type:        "",
				ServiceName: "",
				Debug:       false,
				Consul: &configv1.Discovery_Consul{
					Address:                        "",
					Scheme:                         "",
					Token:                          "",
					HeartBeat:                      false,
					HealthCheck:                    false,
					Datacenter:                     "",
					HealthCheckInterval:            0,
					Timeout:                        0,
					DeregisterCriticalServiceAfter: 0,
				},
				Etcd: &configv1.Discovery_ETCD{
					Endpoints: nil,
				},
			},
			Storages: nil,
		},
		Services: []*configv1.Service{
			&configv1.Service{
				Name:            "",
				DynamicEndpoint: false,
				Grpc: &configv1.Service_GRPC{
					Network: "",
					Addr:    "",
					UseTls:  false,
					TlsConfig: &configv1.TLSConfig{
						File: &configv1.TLSConfig_File{
							Cert: "",
							Key:  "",
							Ca:   "",
						},
						Pem: &configv1.TLSConfig_PEM{
							Cert: nil,
							Key:  nil,
							Ca:   nil,
						},
					},
					Timeout:         0,
					ShutdownTimeout: 0,
					ReadTimeout:     0,
					WriteTimeout:    0,
					IdleTimeout:     0,
					Endpoint:        "",
				},
				Http: &configv1.Service_HTTP{
					Network: "",
					Addr:    "",
					UseTls:  false,
					TlsConfig: &configv1.TLSConfig{
						File: &configv1.TLSConfig_File{
							Cert: "",
							Key:  "",
							Ca:   "",
						},
						Pem: &configv1.TLSConfig_PEM{
							Cert: nil,
							Key:  nil,
							Ca:   nil,
						},
					},
					Timeout:         0,
					ShutdownTimeout: 0,
					ReadTimeout:     0,
					WriteTimeout:    0,
					IdleTimeout:     0,
					Endpoint:        "",
				},
				Websocket: &configv1.WebSocket{
					Network: "",
					Addr:    "",
					Path:    "",
					Codec:   "",
					Timeout: 0,
				},
				Message: &configv1.Message{
					Type: "",
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
				},
				Task: &configv1.Task{
					Type: "",
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
				},
				Middleware: &v11.Middleware{
					//Logging:        false,
					//Recovery:       false,
					//Tracing:        false,
					//CircuitBreaker: false,
					Metadata: &v11.Middleware_Metadata{
						Enabled: false,
						Prefix:  "",
						Data:    nil,
					},
					RateLimiter: &ratelimitv1.RateLimiter{
						Enabled:             false,
						Name:                "",
						Period:              0,
						XRatelimitLimit:     0,
						XRatelimitRemaining: 0,
						XRatelimitReset:     0,
						RetryAfter:          0,
						Memory: &ratelimitv1.RateLimiter_Memory{
							Expiration:      0,
							CleanupInterval: 0,
						},
						Redis: &ratelimitv1.RateLimiter_Redis{
							Addr:     "",
							Username: "",
							Password: "",
							Db:       0,
						},
					},
					Metrics: &metricsv1.Metrics{
						Enabled:          false,
						SupportedMetrics: nil,
						UserMetrics:      nil,
					},
					Validator: &validatorv1.Validator{
						Enabled:  false,
						Version:  0,
						FailFast: false,
					},
					Jwt: &jwtv1.JWT{
						Enabled:     false,
						Subject:     "",
						ClaimType:   "",
						TokenHeader: nil,
						//Config: &jwtv1.Config{
						//	SigningMethod:        "",
						//	Key:                  "",
						//	Key2:                 "",
						//	AccessTokenLifetime:  0,
						//	RefreshTokenLifetime: 0,
						//	Issuer:               "",
						//	Audience:             nil,
						//	TokenType:            "",
						//},
					},
					Selector: &selectorv1.Selector{
						Enabled:  false,
						Names:    nil,
						Paths:    nil,
						Regex:    "",
						Prefixes: nil,
					},
				},
				Selector: &configv1.Service_Selector{
					Version: "",
					Builder: "",
				},
			},
		},
		Middleware: &v11.Middleware{
			//Logging:        false,
			//Recovery:       false,
			//Tracing:        false,
			//CircuitBreaker: false,
			Metadata: &v11.Middleware_Metadata{
				Enabled: false,
				Prefix:  "",
				Data:    nil,
			},
			RateLimiter: &ratelimitv1.RateLimiter{
				Enabled:             false,
				Name:                "",
				Period:              0,
				XRatelimitLimit:     0,
				XRatelimitRemaining: 0,
				XRatelimitReset:     0,
				RetryAfter:          0,
				Memory: &ratelimitv1.RateLimiter_Memory{
					Expiration:      0,
					CleanupInterval: 0,
				},
				Redis: &ratelimitv1.RateLimiter_Redis{
					Addr:     "",
					Username: "",
					Password: "",
					Db:       0,
				},
			},
			Metrics: &metricsv1.Metrics{
				Enabled:          false,
				SupportedMetrics: nil,
				UserMetrics:      nil,
			},
			Validator: &validatorv1.Validator{
				Enabled:  false,
				Version:  0,
				FailFast: false,
			},
			Jwt: &jwtv1.JWT{
				Enabled:     false,
				Subject:     "",
				ClaimType:   "",
				TokenHeader: nil,
				//Config: &jwtv1.Config{
				//	SigningMethod:        "",
				//	Key:                  "",
				//	Key2:                 "",
				//	AccessTokenLifetime:  0,
				//	RefreshTokenLifetime: 0,
				//	Issuer:               "",
				//	Audience:             nil,
				//	TokenType:            "",
				//},
			},
			Selector: &selectorv1.Selector{
				Enabled:  false,
				Names:    nil,
				Paths:    nil,
				Regex:    "",
				Prefixes: nil,
			},
		},
	}
	ss = append(ss, testService)

	// 验证服务添加后的数量和内容
	assert.Len(t, ss, 1, "添加服务后应包含一个元素")
	assert.Equal(t, "test-service", ss[0].Services[0].Name, "服务名称不匹配")
}
