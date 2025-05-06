/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"fmt"

	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/registry"

	"origadmin/application/admin/internal/configs"
)

// ConfigManager 实现配置获取和转换能力
type ConfigManager struct {
	discover registry.KDiscovery
}

func NewConfigManager(discover registry.KDiscovery) *ConfigManager {
	return &ConfigManager{discover: discover}
}

// 从 Consul KV 获取特定配置
func (cm *ConfigManager) GetConfig(configName string) ([]byte, error) {
	client, ok := cm.discover.(*registry.ConsulClient)
	if !ok {
		return nil, fmt.Errorf("discovery client is not Consul")
	}

	// 从 Consul KV 获取配置
	return client.GetKV(configName)
}

// 从发现服务获取配置
func (cm *ConfigManager) GetServiceConfig(ctx context.Context, name string) (*configs.Bootstrap, error) {
	instances, err := cm.discover.GetService(ctx, name)
	if err != nil {
		return nil, err
	}

	// 实现配置转换逻辑
	return convertInstancesToConfig(instances)
}

// 实现配置转换逻辑
func convertInstancesToConfig(instances []*registry.KServiceInstance) (*configs.Bootstrap, error) {
	if len(instances) == 0 {
		return nil, fmt.Errorf("no instances found")
	}

	// 示例实现：从第一个实例提取基础配置
	firstInstance := instances[0]
	return &configs.Bootstrap{
		ServiceName: firstInstance.Name,
		Discovery:   nil, // 需要重新初始化发现客户端
		// 其他字段根据需要映射...

	}, nil
}
