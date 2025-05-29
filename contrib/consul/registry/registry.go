/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package registry

import (
	"time"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/origadmin/runtime"
	configv1 "github.com/origadmin/runtime/api/gen/go/config/v1"
	"github.com/origadmin/runtime/registry"
	"github.com/origadmin/toolkits/errors"
)

type consulBuilder struct {
}

func init() {
	runtime.RegisterRegistry(Type, &consulBuilder{})
}

func configFromConfig(registry *configv1.Discovery) *consulapi.Config {
	apiconfig := consulapi.DefaultConfig()
	cfg := registry.GetConsul()
	if cfg == nil {
		return apiconfig
	}
	if cfg.Address != "" {
		apiconfig.Address = cfg.Address
	}
	if cfg.Scheme != "" {
		apiconfig.Scheme = cfg.Scheme
	}
	if cfg.Datacenter != "" {
		apiconfig.Datacenter = cfg.Datacenter
	}
	if cfg.Token != "" {
		apiconfig.Token = cfg.Token
	}
	return apiconfig
}

func optionsFromConfig(discovery *configv1.Discovery) []Option {
	var opts []Option

	cfg := discovery.GetConsul()
	if cfg == nil {
		return opts
	}

	if cfg.HealthCheck {
		opts = append(opts, WithHealthCheck(cfg.HealthCheck))
	}
	if cfg.HeartBeat {
		opts = append(opts, WithHeartbeat(cfg.HeartBeat))
	}
	if cfg.Timeout != 0 {
		opts = append(opts, WithTimeout(time.Duration(cfg.Timeout)))
	}
	if cfg.Datacenter != "" {
		opts = append(opts, WithDatacenter(Datacenter(cfg.Datacenter)))
	}
	if cfg.HealthCheckInterval > 0 {
		opts = append(opts, WithHealthCheckInterval(int(cfg.HealthCheckInterval)))
	}
	if cfg.DeregisterCriticalServiceAfter > 0 {
		opts = append(opts, WithDeregisterCriticalServiceAfter(int(cfg.DeregisterCriticalServiceAfter)))
	}
	return opts
}

func (c *consulBuilder) NewDiscovery(cfg *configv1.Discovery, opts ...registry.Option) (registry.KDiscovery, error) {
	return c.Create(cfg, opts...)
}

func (c *consulBuilder) NewRegistrar(cfg *configv1.Discovery, opts ...registry.Option) (registry.KRegistrar, error) {
	return c.Create(cfg, opts...)
}

func (c *consulBuilder) Create(cfg *configv1.Discovery, _ ...registry.Option) (registry.Registry, error) {
	if cfg == nil || cfg.Consul == nil {
		return nil, errors.New("configuration: consul config is required")
	}
	apiConfig := configFromConfig(cfg)
	apiClient, err := consulapi.NewClient(apiConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create consul client")
	}
	r := New(apiClient, optionsFromConfig(cfg)...)
	return r, nil
}
