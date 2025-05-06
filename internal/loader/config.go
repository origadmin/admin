/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/service"
)

// BootstrapConfig 包含启动配置
type BootstrapConfig struct {
	Service service.ServiceBuilder
	Source  bootstrap.SourceConfig
}

// LoadBootstrap 加载基础配置
func LoadBootstrap(cfg BootstrapConfig) (*configs.Bootstrap, error) {
	var bs *configs.Bootstrap
	var err error
	
	switch cfg.Source.GetType() {
	case "file":
		bs, err = LoadLocalBootstrap(&cfg.Source)
	default:
		bs, err = LoadRemoteBootstrap(&cfg.Source)
	}
	
	if err != nil {
		return nil, fmt.Errorf("load bootstrap error: %v", err)
	}
	
	return bs, nil
}

// 新增服务发现配置加载
func LoadRemoteBootstrap(cfg *bootstrap.SourceConfig) (*configs.Bootstrap, error) {
	discoveryConfig := &registry.ConsulConfig{ /*...*/ }
	registrar, _ := registry.NewConsulRegistrar(discoveryConfig)
	
	return &configs.Bootstrap{
		Registry: registrar, // 远程配置携带服务注册能力
		Discovery: discoveryConfig.Client,
	}, nil
}
