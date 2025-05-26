/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package registry implements the functions, types, and interfaces for the module.
package registry

import (
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
)

const (
	SingleDatacenter = consul.SingleDatacenter
	MultiDatacenter  = consul.MultiDatacenter
	Type             = "consul"
)

type (
	Datacenter      = consul.Datacenter
	Client          = consul.Client
	ServiceResolver = consul.ServiceResolver
	Option          = consul.Option
	Config          = consul.Config
	Registry        = consul.Registry
)

// WithHealthCheck is a wrapper for consul.WithHealthCheck
func WithHealthCheck(check bool) Option {
	return consul.WithHealthCheck(check)
}

// WithTimeout is a wrapper for consul.WithTimeout
func WithTimeout(timeout time.Duration) Option {
	return consul.WithTimeout(timeout)
}

// WithDatacenter is a wrapper for consul.WithDatacenter
func WithDatacenter(datacenter Datacenter) Option {
	return consul.WithDatacenter(datacenter)
}

// WithHeartbeat is a wrapper for consul.WithHeartbeat
func WithHeartbeat(heartbeat bool) Option {
	return consul.WithHeartbeat(heartbeat)
}

// WithServiceResolver is a wrapper for consul.WithServiceResolver
func WithServiceResolver(resolver ServiceResolver) Option {
	return consul.WithServiceResolver(resolver)
}

// WithHealthCheckInterval is a wrapper for consul.WithHealthCheckInterval
func WithHealthCheckInterval(interval int) Option {
	return consul.WithHealthCheckInterval(interval)
}

// WithDeregisterCriticalServiceAfter is a wrapper for consul.WithDeregisterCriticalServiceAfter
func WithDeregisterCriticalServiceAfter(duration int) Option {
	return consul.WithDeregisterCriticalServiceAfter(duration)
}

// WithServiceCheck is a wrapper for consul.WithServiceCheck
func WithServiceCheck(check *api.AgentServiceCheck) Option {
	return consul.WithServiceCheck(check)
}

// New is a wrapper for consul.New
func New(client *api.Client, opts ...Option) *Registry {
	return consul.New(client, opts...)
}
