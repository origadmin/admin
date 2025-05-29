/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"github.com/origadmin/runtime"
	configv1 "github.com/origadmin/runtime/api/gen/go/config/v1"

	"origadmin/application/admin/internal/configs"
)

func LoadBootstrap(cfg *configv1.SourceConfig) (*configs.Bootstrap, error) {
	source, err := runtime.NewConfig(cfg)
	if err != nil {
		return nil, err
	}
	if err := source.Load(); err != nil {
		return nil, err
	}
	var bs configs.Bootstrap
	if err := source.Scan(&bs); err != nil {
		return nil, err
	}
	return &bs, nil
}

func LoadLocalBootstrap(path string) (*configs.Bootstrap, error) {
	source := configv1.SourceConfig{
		Types: []string{"file"},
		File: &configv1.SourceConfig_File{
			Path: path,
		},
	}
	return LoadBootstrap(&source)
}
