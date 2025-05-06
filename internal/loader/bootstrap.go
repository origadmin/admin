/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"github.com/origadmin/runtime/bootstrap"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/registry"
)

// LoadRemoteBootstrap 从 Consul KV 获取配置
func LoadRemoteBootstrap(cfg *bootstrap.SourceConfig) (*configs.Bootstrap, error) {
	// 创建 Consul 发现客户端（同时作为 KV 客户端）
	discover, err := registry.NewConsulDiscover(&registry.ConsulConfig{
		Address: os.Getenv("CONSUL_ADDR"),
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	// 使用 ConfigManager 从 KV 获取配置
	configMgr := NewConfigManager(discover)
	kvData, err := configMgr.GetConfig("config/" + cfg.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get KV config: %w", err)
	}

	// 解析配置数据（示例实现）
	var bs configs.Bootstrap
	if err := parseConfigData(kvData, &bs); err != nil {
		return nil, err
	}

	return &bs, nil
}

// parseConfigData 实现配置数据解析逻辑
func parseConfigData(data []byte, out *configs.Bootstrap) error {
	// 根据实际格式实现解析（如 JSON/TOML）
	// 示例：json.Unmarshal(data, out)
	return nil
}
