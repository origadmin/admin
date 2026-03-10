// Copyright 2024 OrigAdmin. All rights reserved.

package casbin

import (
	"testing"

	"github.com/casbin/casbin/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/contrib/security/authz/casbin/adapter"
)

func TestCasbinModel(t *testing.T) {
	// 加载模型
	modelPath := "../../../resources/casbin_model.conf"

	// 创建内存适配器
	memAdapter := adapter.NewMemory()
	enforcer, err := casbin.NewEnforcer(modelPath, memAdapter)
	require.NoError(t, err, "Failed to create enforcer")

	t.Run("TestBasicPermission", func(t *testing.T) {
		// 添加策略: role1 可以访问 resource1 的 read 操作
		_, err := enforcer.AddPolicy("role1", "domain1", "system:user:list", "read")
		require.NoError(t, err)

		// 验证权限
		allowed, err := enforcer.Enforce("role1", "domain1", "system:user:list", "read")
		require.NoError(t, err)
		assert.True(t, allowed, "role1 should have read permission")

		// 验证无权限的操作
		allowed, err = enforcer.Enforce("role1", "domain1", "system:user:delete", "write")
		require.NoError(t, err)
		assert.False(t, allowed, "role1 should not have delete permission")
	})

	t.Run("TestWildcardMatching", func(t *testing.T) {
		// 添加通配符策略: domain使用通配符匹配所有domain
		_, err := enforcer.AddPolicy("admin", "*", "system:*", "ANY")
		require.NoError(t, err)

		// 验证通配符权限: 测试不同domain都能匹配
		tests := []struct {
			sub  string
			dom  string
			obj  string
			act  string
			want bool
		}{
			{"admin", "domain1", "system:user:list", "read", true},
			{"admin", "domain2", "system:user:create", "write", true},
			{"admin", "domain3", "system:role:delete", "delete", true},
			{"admin", "any_domain", "system:user:list", "read", true},
			{"user", "domain1", "system:user:list", "read", false},
		}

		for _, tt := range tests {
			t.Run(tt.sub+"_"+tt.dom+"_"+tt.obj, func(t *testing.T) {
				allowed, err := enforcer.Enforce(tt.sub, tt.dom, tt.obj, tt.act)
				require.NoError(t, err)
				assert.Equal(t, tt.want, allowed)
			})
		}
	})

	t.Run("TestRoleInheritance", func(t *testing.T) {
		// 角色继承: user1 属于 role1, role1 继承 admin
		_, err := enforcer.AddGroupingPolicy("user1", "role1", "domain1")
		require.NoError(t, err)

		_, err = enforcer.AddGroupingPolicy("role1", "admin", "domain1")
		require.NoError(t, err)

		// 添加admin权限
		_, err = enforcer.AddPolicy("admin", "domain1", "system:user:list", "read")
		require.NoError(t, err)

		// 验证继承权限
		allowed, err := enforcer.Enforce("user1", "domain1", "system:user:list", "read")
		require.NoError(t, err)
		assert.True(t, allowed, "user1 should inherit permission from admin through role1")
	})

	t.Run("TestDomainIsolation", func(t *testing.T) {
		// 域隔离测试
		_, err := enforcer.AddPolicy("user1", "domain1", "resource1", "read")
		require.NoError(t, err)

		_, err = enforcer.AddPolicy("user1", "domain2", "resource1", "read")
		require.NoError(t, err)

		// user1 对 domain1 有权限
		allowed, err := enforcer.Enforce("user1", "domain1", "resource1", "read")
		require.NoError(t, err)
		assert.True(t, allowed)

		// user1 对 domain2 也有权限
		allowed, err = enforcer.Enforce("user1", "domain2", "resource1", "read")
		require.NoError(t, err)
		assert.True(t, allowed)

		// user1 对 domain3 没有权限
		allowed, err = enforcer.Enforce("user1", "domain3", "resource1", "read")
		require.NoError(t, err)
		assert.False(t, allowed)
	})
}
