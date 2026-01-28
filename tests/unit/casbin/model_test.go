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
	adapter := adapter.NewMemory()
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	require.NoError(t, err, "Failed to create enforcer")

	t.Run("TestBasicPermission", func(t *testing.T) {
		// 添加策略: role1 可以访问 resource1 的 read 操作
		_, err := enforcer.AddPolicy("role1", "system:user:list", "read", "domain1")
		require.NoError(t, err)

		// 验证权限
		allowed, err := enforcer.Enforce("role1", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed, "role1 should have read permission")

		// 验证无权限的操作
		allowed, err = enforcer.Enforce("role1", "system:user:delete", "write", "domain1")
		require.NoError(t, err)
		assert.False(t, allowed, "role1 should not have delete permission")
	})

	t.Run("TestWildcardMatching", func(t *testing.T) {
		// 添加通配符策略: * 表示所有操作
		_, err := enforcer.AddPolicy("admin", "system:*", "*", "*")
		require.NoError(t, err)

		// 验证通配符权限
		tests := []struct {
			sub  string
			obj  string
			act  string
			dom  string
			want bool
		}{
			{"admin", "system:user:list", "read", "domain1", true},
			{"admin", "system:user:create", "write", "domain2", true},
			{"admin", "system:role:delete", "delete", "domain3", true},
			{"user", "system:user:list", "read", "domain1", false},
		}

		for _, tt := range tests {
			t.Run(tt.sub+"_"+tt.obj, func(t *testing.T) {
				allowed, err := enforcer.Enforce(tt.sub, tt.obj, tt.act, tt.dom)
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
		_, err = enforcer.AddPolicy("admin", "system:user:list", "read", "domain1")
		require.NoError(t, err)

		// 验证继承权限
		allowed, err := enforcer.Enforce("user1", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed, "user1 should inherit permission from admin through role1")
	})

	t.Run("TestDomainIsolation", func(t *testing.T) {
		// 域隔离测试
		_, err := enforcer.AddPolicy("user1", "resource1", "read", "domain1")
		require.NoError(t, err)

		_, err = enforcer.AddPolicy("user1", "resource1", "read", "domain2")
		require.NoError(t, err)

		// user1 在 domain1 有权限
		allowed, err := enforcer.Enforce("user1", "resource1", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed)

		// user1 在 domain2 也有权限
		allowed, err = enforcer.Enforce("user1", "resource1", "read", "domain2")
		require.NoError(t, err)
		assert.True(t, allowed)

		// user1 在 domain3 没有权限
		allowed, err = enforcer.Enforce("user1", "resource1", "read", "domain3")
		require.NoError(t, err)
		assert.False(t, allowed)
	})
}
