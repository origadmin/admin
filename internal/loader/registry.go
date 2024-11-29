/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"errors"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/registry"

	"origadmin/application/admin/internal/configs"
)

func NewRegistrar(bootstrap *configs.Bootstrap) (registry.Registrar, error) {
	cfg := bootstrap.GetRegistry()
	if cfg == nil {
		return nil, errors.New("registry config is nil")
	}
	registrar, err := runtime.NewRegistrar(cfg)
	if err != nil {
		return nil, err
	}
	return registrar, nil
}

func NewDiscovery(bootstrap *configs.Bootstrap) (registry.Discovery, error) {
	cfg := bootstrap.GetRegistry()
	if cfg == nil {
		return nil, errors.New("registry config is nil")
	}
	discovery, err := runtime.NewDiscovery(cfg)
	if err != nil {
		return nil, err
	}
	return discovery, nil
}
