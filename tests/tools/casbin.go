// Copyright 2024 OrigAdmin. All rights reserved.

package tools

import (
	"testing"

	"github.com/casbin/casbin/v3"
	"github.com/stretchr/testify/require"

	"origadmin/application/admin/internal/data"
)

// SetupTestCasbin 创建测试用的Casbin Enforcer
func SetupTestCasbin(t *testing.T, adapter *data.CasbinAdapter, modelPath string) *casbin.Enforcer {
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	require.NoError(t, err, "Failed to create enforcer")

	return enforcer
}

// AddTestPolicies 添加测试策略
func AddTestPolicies(t *testing.T, enforcer *casbin.Enforcer, policies [][]string) {
	for _, policy := range policies {
		// 转换 []string 为 []interface{}
		params := make([]interface{}, len(policy))
		for i, v := range policy {
			params[i] = v
		}
		_, err := enforcer.AddPolicy(params...)
		require.NoError(t, err, "Failed to add policy: %v", policy)
	}
}

// AddTestGroupingPolicies 添加测试角色分组
func AddTestGroupingPolicies(t *testing.T, enforcer *casbin.Enforcer, groupingPolicies [][]string) {
	for _, gp := range groupingPolicies {
		// 转换 []string 为 []interface{}
		params := make([]interface{}, len(gp))
		for i, v := range gp {
			params[i] = v
		}
		_, err := enforcer.AddGroupingPolicy(params...)
		require.NoError(t, err, "Failed to add grouping policy: %v", gp)
	}
}

// AssertPermission 断言权限
func AssertPermission(t *testing.T, enforcer *casbin.Enforcer, sub, obj, act, dom string, allowed bool) {
	result, err := enforcer.Enforce(sub, obj, act, dom)
	require.NoError(t, err)
	require.Equal(t, allowed, result,
		"Permission check failed for sub=%s, obj=%s, act=%s, dom=%s. Expected %v, got %v",
		sub, obj, act, dom, allowed, result)
}
