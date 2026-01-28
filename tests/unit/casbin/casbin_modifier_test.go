/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/casbin/casbin/v3"
	"github.com/origadmin/casbin-watcher/v3"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/origadmin/casbin-watcher/v3/drivers/mem"

	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/features/auth/dal"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
)

// TestCasbinModifier_WriteAndEnforce 测试写入Adapter后Enforcer能否正确验证
// 这是文档中指出的缺失测试: CasbinAdapter操作测试
// 测试流程: Modifier写入Adapter → Watcher通知 → Enforcer重新加载 → 验证权限
func TestCasbinModifier_WriteAndEnforce(t *testing.T) {
	ctx := context.Background()

	// 1. 创建测试数据库
	client := enttest.Open(t, "sqlite3", "file:casbin_modifier?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// 2. 创建 Adapter (数据写入层)
	adapter := &data.CasbinAdapter{
		Ctx: ctx,
		DB:  client,
	}

	// 3. 创建 GoChannel Watcher (通知层) - 使用内存 channel 用于测试
	endpointURL := "mem://casbin_updates?shared=true"
	watcher, err := watcher.NewWatcher(ctx, endpointURL)
	require.NoError(t, err)
	defer watcher.Close()

	// 4. 创建 PolicyModifier (业务层)
	modifier, err := dal.NewCasbinModifier(adapter, watcher, log.DefaultLogger)
	require.NoError(t, err)

	// 5. 创建 Enforcer (验证层) - 模拟另一个实例
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	require.NoError(t, err)

	// 设置 watcher 的回调,让 enforcer 自动重新加载
	updateCh := make(chan string, 1)
	err = watcher.SetUpdateCallback(func(msg string) {
		updateCh <- msg
	})
	require.NoError(t, err)

	t.Run("AddUserRole_ThenVerifyWithEnforcer", func(t *testing.T) {
		// 步骤1: 通过 Modifier 写入用户-角色关系
		userID := "user123"
		roleID := "role456"

		added, err := modifier.AddUserRole(ctx, userID, roleID)
		require.NoError(t, err)
		assert.True(t, added, "AddUserRole should succeed")

		// 步骤2: 等待 Watcher 通知
		select {
		case msg := <-updateCh:
			assert.NotEmpty(t, msg, "Watcher should send notification")
		case <-time.After(5 * time.Second):
			t.Fatal("Watcher notification timeout")
		}

		// 步骤3: Enforcer 重新加载策略 (模拟收到通知后)
		err = enforcer.LoadPolicy()
		require.NoError(t, err, "Enforcer should reload policy")

		// 步骤4: 使用 Enforcer 验证用户是否有角色
		hasRole, err := enforcer.HasGroupingPolicy(userID, roleID)
		require.NoError(t, err)
		assert.True(t, hasRole, "Enforcer should see the user-role link")
	})

	t.Run("AddRolePermission_ThenVerifyWithEnforcer", func(t *testing.T) {
		// 步骤1: 添加角色-权限关系
		roleID := "admin"
		spec := authz.RuleSpec{
			Resource: "system:user:create",
			Action:   "write",
		}

		added, err := modifier.AddRolePermission(ctx, roleID, spec)
		require.NoError(t, err)
		assert.True(t, added, "AddRolePermission should succeed")

		// 步骤2: 等待 Watcher 通知
		select {
		case msg := <-updateCh:
			assert.NotEmpty(t, msg, "Watcher should send notification")
		case <-time.After(5 * time.Second):
			t.Fatal("Watcher notification timeout")
		}

		// 步骤3: Enforcer 重新加载策略
		err = enforcer.LoadPolicy()
		require.NoError(t, err)

		// 步骤4: 验证角色是否有权限
		allowed, err := enforcer.Enforce(roleID, spec.Resource, spec.Action, "*")
		require.NoError(t, err)
		assert.True(t, allowed, "Enforcer should see the role-permission link")
	})

	t.Run("RemoveUserRole_ThenVerifyWithEnforcer", func(t *testing.T) {
		userID := "user789"
		roleID := "role999"

		// 先添加
		_, _ = modifier.AddUserRole(ctx, userID, roleID)

		// 等待通知并重新加载
		select {
		case <-updateCh:
		case <-time.After(5 * time.Second):
		}
		_ = enforcer.LoadPolicy()

		// 验证存在
		hasRole, _ := enforcer.HasGroupingPolicy(userID, roleID)
		assert.True(t, hasRole, "User should have role before removal")

		// 移除
		removed, err := modifier.RemoveUserRole(ctx, userID, roleID)
		require.NoError(t, err)
		assert.True(t, removed, "RemoveUserRole should succeed")

		// 等待通知并重新加载
		select {
		case <-updateCh:
		case <-time.After(5 * time.Second):
		}
		_ = enforcer.LoadPolicy()

		// 验证已移除
		hasRole, err = enforcer.HasGroupingPolicy(userID, roleID)
		require.NoError(t, err)
		assert.False(t, hasRole, "Enforcer should not see the user-role link after removal")
	})

	t.Run("RemoveAllUserRoles_ThenVerifyWithEnforcer", func(t *testing.T) {
		userID := "user_multi"

		// 添加多个角色
		_, _ = modifier.AddUserRole(ctx, userID, "role1")
		_, _ = modifier.AddUserRole(ctx, userID, "role2")
		_, _ = modifier.AddUserRole(ctx, userID, "role3")

		// 等待通知并重新加载
		for i := 0; i < 3; i++ {
			select {
			case <-updateCh:
			case <-time.After(5 * time.Second):
			}
		}
		_ = enforcer.LoadPolicy()

		// 验证有3个角色
		roles, _ := enforcer.GetRolesForUser(userID)
		assert.Len(t, roles, 3, "User should have 3 roles")

		// 移除所有角色
		removed, err := modifier.RemoveAllUserRoles(ctx, userID)
		require.NoError(t, err)
		assert.True(t, removed, "RemoveAllUserRoles should succeed")

		// 等待通知并重新加载
		select {
		case <-updateCh:
		case <-time.After(5 * time.Second):
		}
		_ = enforcer.LoadPolicy()

		// 验证已移除所有角色
		roles, _ = enforcer.GetRolesForUser(userID)
		assert.Len(t, roles, 0, "User should have no roles after removal")
	})

	t.Run("DirectUserPermission_ThenVerifyWithEnforcer", func(t *testing.T) {
		userID := "user_direct"
		spec := authz.RuleSpec{
			Resource: "system:config:read",
			Action:   "read",
		}

		// 添加用户直接权限
		added, err := modifier.AddUserPermission(ctx, userID, spec)
		require.NoError(t, err)
		assert.True(t, added)

		// 等待通知并重新加载
		select {
		case <-updateCh:
		case <-time.After(5 * time.Second):
		}
		_ = enforcer.LoadPolicy()

		// 验证用户直接权限
		allowed, err := enforcer.Enforce(userID, spec.Resource, spec.Action, "*")
		require.NoError(t, err)
		assert.True(t, allowed, "Enforcer should see the user-permission link")
	})
}

// TestCasbinModifier_ConcurrentWrites 测试并发写入场景
// 验证 Adapter 在并发情况下的正确性
func TestCasbinModifier_ConcurrentWrites(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:casbin_concurrent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	adapter := &data.CasbinAdapter{
		Ctx: ctx,
		DB:  client,
	}

	endpointURL := "mem://casbin_updates?shared=true"
	watcher, err := watcher.NewWatcher(ctx, endpointURL)
	require.NoError(t, err)
	defer watcher.Close()

	modifier, err := dal.NewCasbinModifier(adapter, watcher, log.DefaultLogger)
	require.NoError(t, err)

	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	require.NoError(t, err)

	t.Run("ConcurrentAddUserRole", func(t *testing.T) {
		const concurrency = 10
		userID := "user_concurrent"

		var wg sync.WaitGroup
		errChan := make(chan error, concurrency)

		for i := 0; i < concurrency; i++ {
			wg.Add(1)
			go func(roleIndex int) {
				defer wg.Done()
				roleID := "role_" + string(rune('0'+roleIndex))

				_, err := modifier.AddUserRole(ctx, userID, roleID)
				if err != nil {
					errChan <- err
				}
			}(i)
		}

		wg.Wait()
		close(errChan)

		// 检查是否有错误
		for err := range errChan {
			t.Errorf("Concurrent add failed: %v", err)
		}

		// 重新加载并验证
		_ = enforcer.LoadPolicy()
		roles, err := enforcer.GetRolesForUser(userID)
		require.NoError(t, err)
		assert.Len(t, roles, concurrency, "All concurrent adds should succeed")
	})

	t.Run("ConcurrentAddRolePermission", func(t *testing.T) {
		const concurrency = 10
		roleID := "role_perm_concurrent"

		var wg sync.WaitGroup
		errChan := make(chan error, concurrency)

		for i := 0; i < concurrency; i++ {
			wg.Add(1)
			go func(permIndex int) {
				defer wg.Done()
				spec := authz.RuleSpec{
					Resource: string(rune('a' + permIndex)),
					Action:   "read",
				}

				_, err := modifier.AddRolePermission(ctx, roleID, spec)
				if err != nil {
					errChan <- err
				}
			}(i)
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			t.Errorf("Concurrent add permission failed: %v", err)
		}

		// 验证每个权限都能正确检查
		_ = enforcer.LoadPolicy()

		for i := 0; i < concurrency; i++ {
			resource := string(rune('a' + i))
			allowed, err := enforcer.Enforce(roleID, resource, "read", "*")
			require.NoError(t, err)
			assert.True(t, allowed, "Permission %s should be allowed", resource)
		}
	})
}

// TestCasbinModifier_WatcherNotification 测试 Watcher 通知机制
func TestCasbinModifier_WatcherNotification(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:casbin_watcher?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	adapter := &data.CasbinAdapter{
		Ctx: ctx,
		DB:  client,
	}

	endpointURL := "mem://casbin_updates?shared=true"
	watcher, err := watcher.NewWatcher(ctx, endpointURL)
	require.NoError(t, err)
	defer watcher.Close()

	updateCount := 0
	var mu sync.Mutex

	err = watcher.SetUpdateCallback(func(msg string) {
		mu.Lock()
		defer mu.Unlock()
		updateCount++
	})
	require.NoError(t, err)

	modifier, err := dal.NewCasbinModifier(adapter, watcher, log.DefaultLogger)
	require.NoError(t, err)

	t.Run("WatcherUpdateCount", func(t *testing.T) {
		mu.Lock()
		initialCount := updateCount
		mu.Unlock()

		// 执行多次修改操作
		_, _ = modifier.AddUserRole(ctx, "user1", "role1")
		_, _ = modifier.AddUserRole(ctx, "user2", "role2")
		_, _ = modifier.AddRolePermission(ctx, "role1", authz.RuleSpec{Resource: "r1", Action: "read"})

		// 等待所有通知
		time.Sleep(100 * time.Millisecond)

		mu.Lock()
		count := updateCount
		mu.Unlock()

		assert.Equal(t, initialCount+3, count,
			"Watcher should be notified for each successful modification")
	})

	t.Run("NoNotificationOnDuplicate", func(t *testing.T) {
		// 添加重复的角色 (第二次应该不触发通知)
		_, _ = modifier.AddUserRole(ctx, "user_dup", "role_dup")

		// 等待通知
		time.Sleep(100 * time.Millisecond)

		mu.Lock()
		countAfterFirstAdd := updateCount
		mu.Unlock()

		_, _ = modifier.AddUserRole(ctx, "user_dup", "role_dup") // 重复添加

		// 等待(不应该有新通知)
		time.Sleep(100 * time.Millisecond)

		mu.Lock()
		count := updateCount
		mu.Unlock()

		assert.Equal(t, countAfterFirstAdd, count,
			"Watcher should not be notified for duplicate/no-op operations")
	})

	t.Run("NoWatcher_SilentFailure", func(t *testing.T) {
		// 创建没有 watcher 的 modifier
		modifierNoWatcher, err := dal.NewCasbinModifier(adapter, nil, log.DefaultLogger)
		require.NoError(t, err)

		// 应该能正常工作,只是没有通知
		added, err := modifierNoWatcher.AddUserRole(ctx, "user_nowatcher", "role_nowatcher")
		require.NoError(t, err)
		assert.True(t, added)

		mu.Lock()
		count := updateCount
		mu.Unlock()

		assert.Equal(t, countAfterFirstAdd, count, "Watcher count should not change")
	})
}

// TestCasbinModifier_AdapterDirectAccess 测试直接访问 Adapter 并用 Enforcer 验证
// 这正是你提到的测试场景:写入Adapter → Enforcer检测
func TestCasbinModifier_AdapterDirectAccess(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:adapter_direct?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	adapter := &data.CasbinAdapter{
		Ctx: ctx,
		DB:  client,
	}

	endpointURL := "mem://casbin_updates?shared=true"
	watcher, err := watcher.NewWatcher(ctx, endpointURL)
	require.NoError(t, err)
	defer watcher.Close()

	modifier, err := dal.NewCasbinModifier(adapter, watcher, log.DefaultLogger)
	require.NoError(t, err)

	updateCh := make(chan string, 10)
	err = watcher.SetUpdateCallback(func(msg string) {
		updateCh <- msg
	})
	require.NoError(t, err)

	modelPath := "../../../resources/casbin_model.conf"

	t.Run("SimulateDataUpdateNotification", func(t *testing.T) {
		// 场景: 模拟接收到数据更新通知后的完整流程

		// 1. 模拟外部写入 (通过 Modifier)
		userID := "user_notify"
		roleID := "role_notify"

		added, err := modifier.AddUserRole(ctx, userID, roleID)
		require.NoError(t, err)
		require.True(t, added)

		// 2. 等待 Watcher 通知
		select {
		case msg := <-updateCh:
			assert.NotEmpty(t, msg, "Watcher should send notification")
		case <-time.After(5 * time.Second):
			t.Fatal("Watcher notification timeout")
		}

		// 3. 模拟另一个实例收到 Watcher 通知,创建新的 Enforcer
		// 这模拟了分布式环境下,其他实例收到通知后的行为
		enforcer2, err := casbin.NewEnforcer(modelPath, adapter)
		require.NoError(t, err)

		// 4. Enforcer 重新加载策略 (模拟收到通知后的动作)
		err = enforcer2.LoadPolicy()
		require.NoError(t, err)

		// 5. 验证 Enforcer 能看到最新数据
		hasRole, err := enforcer2.HasGroupingPolicy(userID, roleID)
		require.NoError(t, err)
		assert.True(t, hasRole, "Enforcer should see the latest policy after reload")

		// 6. 进一步验证: 权限检查也能正确工作
		_, _ = modifier.AddRolePermission(ctx, roleID, authz.RuleSpec{
			Resource: "test:resource",
			Action:   "read",
		})

		// 等待通知
		select {
		case <-updateCh:
		case <-time.After(5 * time.Second):
			t.Fatal("Watcher notification timeout")
		}

		err = enforcer2.LoadPolicy()
		require.NoError(t, err)

		allowed, err := enforcer2.Enforce(roleID, "test:resource", "read", "*")
		require.NoError(t, err)
		assert.True(t, allowed, "Permission check should work after policy update")
	})

	t.Run("UpdatePropagationDelay", func(t *testing.T) {
		// 测试更新传播和时序问题

		// 1. 初始状态
		enforcer1, _ := casbin.NewEnforcer(modelPath, adapter)
		enforcer1.LoadPolicy()

		userID := "user_delay"
		roleID := "role_delay"

		// 2. 验证初始无角色
		hasRole, _ := enforcer1.HasGroupingPolicy(userID, roleID)
		assert.False(t, hasRole)

		// 3. 添加角色
		_, _ = modifier.AddUserRole(ctx, userID, roleID)

		// 等待通知
		select {
		case <-updateCh:
		case <-time.After(5 * time.Second):
		}

		// 4. 不重新加载,Enforcer 应该看不到新数据
		hasRole, _ = enforcer1.HasGroupingPolicy(userID, roleID)
		assert.False(t, hasRole, "Enforcer should not see update without reload")

		// 5. 重新加载后应该看到
		err := enforcer1.LoadPolicy()
		require.NoError(t, err)

		hasRole, _ = enforcer1.HasGroupingPolicy(userID, roleID)
		assert.True(t, hasRole, "Enforcer should see update after reload")
	})

	t.Run("TwoInstanceConsistency", func(t *testing.T) {
		// 测试两个实例的一致性

		// 创建两个 watcher 模拟两个服务实例
		watcher1, err := watcher.NewWatcher(ctx, endpointURL)
		require.NoError(t, err)
		defer watcher1.Close()

		watcher2, err := watcher.NewWatcher(ctx, endpointURL)
		require.NoError(t, err)
		defer watcher2.Close()

		// 创建两个 enforcer
		enf1, err := casbin.NewEnforcer(modelPath, adapter)
		require.NoError(t, err)
		enf1.SetWatcher(watcher1)

		enf2, err := casbin.NewEnforcer(modelPath, adapter)
		require.NoError(t, err)
		enf2.SetWatcher(watcher2)

		// 设置回调自动重新加载
		enf1.SetWatcher(watcher1)
		enf2.SetWatcher(watcher2)

		// 通过 enforcer1 添加策略
		_, _ = enf1.AddPolicy("role1", "resource1", "read", "*")

		// 等待传播
		time.Sleep(200 * time.Millisecond)

		// 验证 enforcer2 也能看到
		allowed, err := enf2.Enforce("role1", "resource1", "read", "*")
		require.NoError(t, err)
		assert.True(t, allowed, "Enforcer2 should see policy added by Enforcer1")
	})
}
