/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package conf implements the functions, types, and contracts for the module.
package conf

import (
	"fmt"

	"github.com/origadmin/runtime/config"
	"github.com/origadmin/runtime/engine/bootstrap"
	confpb "origadmin/application/admin/internal/conf/pb"
)

const (
	// APIPrefix is the prefix for all API routes.
	APIPrefix = "/api/v1"
)

// Config is the business configuration.
// It embeds confpb.Bootstrap to provide direct access to all configuration fields
// and satisfy the component.Config interfaces automatically.
type Config struct {
	confpb.Bootstrap
}

// --- Runtime Interface Adapters ---

// Transform scans the configuration from the source and performs any necessary transformations.
func (c *Config) Transform(cfg config.KConfig) (any, error) {
	if err := cfg.Scan(&c.Bootstrap); err != nil {
		return nil, fmt.Errorf("failed to scan config: %w", err)
	}
	return &c.Bootstrap, nil
}

// New creates a new Config transformer.
func New() bootstrap.ConfigTransformer {
	return &Config{}
}

func transformer(cfg config.KConfig) (any, error) {
	var b confpb.Bootstrap
	if err := cfg.Scan(&b); err != nil {
		return nil, fmt.Errorf("failed to scan config: %w", err)
	}
	return &b, nil
}

var Transformer bootstrap.ConfigTransformFunc = transformer
