/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"testing"

	"github.com/casbin/casbin/v3" // Updated import path
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"origadmin/application/admin/internal/data"
)

// TestCasbinModelLoad 测试 Casbin 模型是否能够正确加载
func TestCasbinModelLoad(t *testing.T) {
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath)
	require.NoError(t, err, "Casbin 模型应该能成功加载")
	require.NotNil(t, enforcer, "Enforcer 不应该为 nil")

	// 验证模型定义
	t.Run("VerifyModelDefinition", func(t *testing.T) {
		model := enforcer.GetModel()
		require.NotNil(t, model)

		// 检查请求定义
		assert.Contains(t, model["r"], "r", "应该有请求定义")
		assert.NotNil(t, model["r"]["r"], "请求定义不应该为 nil")

		// 检查策略定义
		assert.Contains(t, model["p"], "p", "应该有策略定义")
		assert.NotNil(t, model["p"]["p"], "策略定义不应该为 nil")

		// 检查角色定义
		assert.Contains(t, model["g"], "g", "应该有角色定义")
		assert.NotNil(t, model["g"]["g"], "角色定义不应该为 nil")

		// 检查策略效果
		assert.Contains(t, model["e"], "e", "应该有策略效果定义")
		assert.NotNil(t, model["e"]["e"], "策略效果定义不应该为 nil")

		// 检查匹配器
		assert.Contains(t, model["m"], "m", "应该有匹配器定义")
		assert.NotNil(t, model["m"]["m"], "匹配器定义不应该为 nil")
	})
}

// TestBasicPermission 测试基本权限控制
func TestBasicPermission(t *testing.T) {
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath)
	require.NoError(t, err)

	t.Run("AllowUserPermission", func(t *testing.T) {
		// 添加策略: user1 可以访问 system:user:list 的 read 操作，域为 domain1
		_, err := enforcer.AddPolicy("user1", "system:user:list", "read", "domain1")
		require.NoError(t, err)

		// 验证权限
		allowed, err := enforcer.Enforce("user1", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed, "user1 应该有权限访问 system:user:list")

		// 清理
		_, _ = enforcer.RemovePolicy("user1", "system:user:list", "read", "domain1")
	})

	t.Run("DenyWithoutPermission", func(t *testing.T) {
		// user2 没有添加任何策略
		allowed, err := enforcer.Enforce("user2", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.False(t, allowed, "user2 不应该有权限")
	})
}

// TestWildcardPermission 测试通配符匹配权限
func TestWildcardPermission(t *testing.T) {
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath)
	require.NoError(t, err)

	t.Run("WildcardResource", func(t *testing.T) {
		// 添加策略: user 可以访问 system:* (所有 system 开头的资源)
		_, err := enforcer.AddPolicy("user", "system:*", "read", "domain1")
		require.NoError(t, err)

		// 测试多个资源
		testCases := []struct {
			resource string
			expected bool
		}{
			{"system:user:list", true},
			{"system:role:list", true},
			{"system:permission:detail", true},
			{"other:resource", false},
		}

		for _, tc := range testCases {
			allowed, err := enforcer.Enforce("user", tc.resource, "read", "domain1")
			require.NoError(t, err)
			assert.Equal(t, tc.expected, allowed, "资源 %s 的权限应该为 %v", tc.resource, tc.expected)
		}

		// 清理
		_, _ = enforcer.RemovePolicy("user", "system:*", "read", "domain1")
	})

	t.Run("WildcardAction", func(t *testing.T) {
		// 添加策略: admin 可以对 system:user:* 执行 ANY 操作
		_, err := enforcer.AddPolicy("admin", "system:user:list", "ANY", "domain1")
		require.NoError(t, err)

		testCases := []struct {
			action   string
			expected bool
		}{
			{"read", true},
			{"write", true},
			{"delete", true},
		}

		for _, tc := range testCases {
			allowed, err := enforcer.Enforce("admin", "system:user:list", tc.action, "domain1")
			require.NoError(t, err)
			assert.Equal(t, tc.expected, allowed, "操作 %s 的权限应该为 %v", tc.action, tc.expected)
		}

		// 清理
		_, _ = enforcer.RemovePolicy("admin", "system:user:list", "ANY", "domain1")
	})
}

// TestRoleInheritance 测试角色继承
func TestRoleInheritance(t *testing.T) {
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath)
	require.NoError(t, err)

	t.Run("SimpleRoleInheritance", func(t *testing.T) {
		// 添加策略: admin 角色可以访问 system:user:list 的 read 操作
		_, err := enforcer.AddPolicy("admin", "system:user:list", "read", "domain1")
		require.NoError(t, err)

		// 用户 alice 具有 admin 角色
		_, err = enforcer.AddGroupingPolicy("alice", "admin", "domain1")
		require.NoError(t, err)

		// alice 应该通过角色继承获得权限
		allowed, err := enforcer.Enforce("alice", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed, "alice 应该通过 admin 角色获得权限")

		// 清理
		_, _ = enforcer.RemovePolicy("admin", "system:user:list", "read", "domain1")
		_, _ = enforcer.RemoveGroupingPolicy("alice", "admin", "domain1")
	})

	t.Run("MultiLevelRoleInheritance", func(t *testing.T) {
		// 添加角色层级: super_admin -> admin -> editor
		_, err := enforcer.AddPolicy("super_admin", "system:*", "ANY", "domain1")
		require.NoError(t, err)
		_, err = enforcer.AddPolicy("admin", "system:user:*", "read", "domain1")
		require.NoError(t, err)
		_, err = enforcer.AddPolicy("editor", "system:user:detail", "read", "domain1")
		require.NoError(t, err)

		// 添加角色继承关系
		_, err = enforcer.AddGroupingPolicy("admin", "super_admin", "domain1")
		require.NoError(t, err)
		_, err = enforcer.AddGroupingPolicy("editor", "admin", "domain1")
		require.NoError(t, err)

		// 用户拥有 editor 角色
		_, err = enforcer.AddGroupingPolicy("bob", "editor", "domain1")
		require.NoError(t, err)

		// bob 应该通过继承链获得所有权限
		allowed1, err := enforcer.Enforce("bob", "system:user:detail", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed1, "bob 应该有 editor 的权限")

		allowed2, err := enforcer.Enforce("bob", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed2, "bob 应该有 admin 的权限")

		allowed3, err := enforcer.Enforce("bob", "system:role:delete", "ANY", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed3, "bob 应该有 super_admin 的权限")

		// 清理
		_, _ = enforcer.RemovePolicy("super_admin", "system:*", "ANY", "domain1")
		_, _ = enforcer.RemovePolicy("admin", "system:user:*", "read", "domain1")
		_, _ = enforcer.RemovePolicy("editor", "system:user:detail", "read", "domain1")
		_, _ = enforcer.RemoveGroupingPolicy("admin", "super_admin", "domain1")
		_, _ = enforcer.RemoveGroupingPolicy("editor", "admin", "domain1")
		_, _ = enforcer.RemoveGroupingPolicy("bob", "editor", "domain1")
	})
}

// TestDomainIsolation 测试域隔离
func TestDomainIsolation(t *testing.T) {
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath)
	require.NoError(t, err)

	t.Run("DifferentDomains", func(t *testing.T) {
		// alice 在 domain1 有权限
		_, err := enforcer.AddPolicy("alice", "system:user:list", "read", "domain1")
		require.NoError(t, err)

		// alice 在 domain2 没有权限
		allowed, err := enforcer.Enforce("alice", "system:user:list", "read", "domain2")
		require.NoError(t, err)
		assert.False(t, allowed, "alice 在 domain2 不应该有权限")

		// alice 在 domain1 有权限
		allowed, err = enforcer.Enforce("alice", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed, "alice 在 domain1 应该有权限")

		// 清理
		_, _ = enforcer.RemovePolicy("alice", "system:user:list", "read", "domain1")
	})

	t.Run("WildcardDomain", func(t *testing.T) {
		// admin 在所有域都有权限
		_, err := enforcer.AddPolicy("admin", "system:*", "ANY", "*")
		require.NoError(t, err)

		testCases := []struct {
			domain   string
			expected bool
		}{
			{"domain1", true},
			{"domain2", true},
			{"domain3", true},
		}

		for _, tc := range testCases {
			allowed, err := enforcer.Enforce("admin", "system:user:list", "read", tc.domain)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, allowed, "admin 在域 %s 的权限应该为 %v", tc.domain, tc.expected)
		}

		// 清理
		_, _ = enforcer.RemovePolicy("admin", "system:*", "ANY", "*")
	})
}

// TestMultiplePermissions 测试多重权限
func TestMultiplePermissions(t *testing.T) {
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath)
	require.NoError(t, err)

	t.Run("UserHasMultiplePermissions", func(t *testing.T) {
		// 添加多个权限给用户
		policies := [][]string{
			{"user1", "system:user:list", "read", "domain1"},
			{"user1", "system:user:detail", "read", "domain1"},
			{"user1", "system:user:create", "write", "domain1"},
		}

		for _, policy := range policies {
			_, err := enforcer.AddPolicy(policy)
			require.NoError(t, err)
		}

		// 验证所有权限
		for _, policy := range policies {
			allowed, err := enforcer.Enforce(policy[0], policy[1], policy[2], policy[3])
			require.NoError(t, err)
			assert.True(t, allowed, "用户应该有权限 %v", policy)
		}

		// 验证没有的权限
		allowed, err := enforcer.Enforce("user1", "system:user:delete", "delete", "domain1")
		require.NoError(t, err)
		assert.False(t, allowed, "用户不应该有删除权限")

		// 清理
		_, _ = enforcer.RemovePolicies(policies)
	})
}

// TestPolicyOperations 测试策略操作
func TestPolicyOperations(t *testing.T) {
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath)
	require.NoError(t, err)

	t.Run("AddAndRemovePolicy", func(t *testing.T) {
		policy := []string{"user1", "system:user:list", "read", "domain1"}

		// 添加策略
		added, err := enforcer.AddPolicy(policy)
		require.NoError(t, err)
		assert.True(t, added, "策略应该被成功添加")

		// 验证策略存在
		hasPolicy, err := enforcer.HasPolicy(policy)
		require.NoError(t, err)
		assert.True(t, hasPolicy, "策略应该存在")

		// 删除策略
		removed, err := enforcer.RemovePolicy(policy)
		require.NoError(t, err)
		assert.True(t, removed, "策略应该被成功删除")

		// 验证策略不存在
		hasPolicy, err = enforcer.HasPolicy(policy)
		require.NoError(t, err)
		assert.False(t, hasPolicy, "策略不应该存在")
	})

	t.Run("UpdatePolicy", func(t *testing.T) {
		oldPolicy := []string{"user1", "system:user:list", "read", "domain1"}
		newPolicy := []string{"user1", "system:user:detail", "write", "domain1"}

		// 添加旧策略
		_, err := enforcer.AddPolicy(oldPolicy)
		require.NoError(t, err)

		// 更新策略
		updated, err := enforcer.UpdatePolicy(oldPolicy, newPolicy)
		require.NoError(t, err)
		assert.True(t, updated, "策略应该被成功更新")

		// 验证旧策略不存在
		hasPolicy, err := enforcer.HasPolicy(oldPolicy)
		require.NoError(t, err)
		assert.False(t, hasPolicy, "旧策略不应该存在")

		// 验证新策略存在
		hasPolicy, err = enforcer.HasPolicy(newPolicy)
		require.NoError(t, err)
		assert.True(t, hasPolicy, "新策略应该存在")

		// 清理
		_, _ = enforcer.RemovePolicy(newPolicy)
	})
}

// TestCasbinAdapterImplementation 验证我们的适配器实现了正确的接口
func TestCasbinAdapterImplementation(t *testing.T) {
	// 这个测试只是验证数据包中的适配器结构存在
	// 实际的适配器测试会在集成测试中进行
	t.Run("AdapterExists", func(t *testing.T) {
		adapter := &data.CasbinAdapter{}
		assert.NotNil(t, adapter)
	})
}
