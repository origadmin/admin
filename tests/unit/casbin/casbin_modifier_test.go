/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/casbin/casbin/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/casbin-watcher/v3"
	_ "github.com/origadmin/casbin-watcher/v3/drivers/mem"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"

	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
	"origadmin/application/admin/internal/features/auth/dal"
)

func setupTest(t *testing.T) (context.Context, authz.PolicyModifier, *casbin.SyncedEnforcer, func()) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", t.Name()))
	db := ent.NewDatabaseWithClient(client)

	adapter := &data.CasbinAdapter{
		Ctx: ctx,
		DB:  db,
	}

	w, err := watcher.NewWatcher(ctx, fmt.Sprintf("mem://%s?shared=true", t.Name()))
	require.NoError(t, err)

	modifier, err := dal.NewCasbinModifier(adapter, w, log.DefaultLogger)
	require.NoError(t, err)

	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewSyncedEnforcer(modelPath, adapter)
	require.NoError(t, err)

	err = w.SetUpdateCallback(func(s string) {
		t.Logf("Watcher callback triggered, reloading policy. Message: %s", s)
		if err := enforcer.LoadPolicy(); err != nil {
			t.Logf("Error reloading policy in watcher callback: %v", err)
		}
	})
	require.NoError(t, err)

	err = enforcer.LoadPolicy()
	require.NoError(t, err, "Initial policy load should succeed")

	cleanup := func() {
		w.Close()
		client.Close()
	}

	return ctx, modifier, enforcer, cleanup
}

func TestPolicyModifier(t *testing.T) {
	const waitFor = 5 * time.Second
	const tick = 100 * time.Millisecond
	const domain1 = "domain1"
	const domain2 = "domain2"
	const defaultDomain = ""

	t.Run("AddAndRemoveRoles", func(t *testing.T) {
		ctx, modifier, enforcer, cleanup := setupTest(t)
		defer cleanup()

		subject := "user1"
		role1 := authz.RoleSpec{Role: "admin", Domain: domain1}
		role2 := authz.RoleSpec{Role: "viewer", Domain: domain1}
		role3 := authz.RoleSpec{Role: "admin", Domain: domain2}

		// Add multiple roles
		_, err := modifier.AddRoles(ctx, subject, role1, role2, role3)
		require.NoError(t, err)

		require.Eventually(t, func() bool {
			r := enforcer.GetRolesForUserInDomain(subject, domain1)
			return len(r) == 2
		}, waitFor, tick, "Should have 2 roles in domain1")
		require.Eventually(t, func() bool {
			r := enforcer.GetRolesForUserInDomain(subject, domain2)
			return len(r) == 1
		}, waitFor, tick, "Should have 1 role in domain2")

		// Remove specific role
		_, err = modifier.RemoveRoles(ctx, subject, role1)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			has, _ := enforcer.HasGroupingPolicy(subject, role1.Role, role1.Domain)
			return !has
		}, waitFor, tick, "Should remove specific role")

		// Remove all roles in a domain
		_, err = modifier.RemoveRoles(ctx, subject, authz.RoleSpec{Domain: domain1})
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			r := enforcer.GetRolesForUserInDomain(subject, domain1)
			return len(r) == 0
		}, waitFor, tick, "Should remove all roles in domain1")

		// Remove a role across all domains
		_, err = modifier.RemoveRoles(ctx, subject, authz.RoleSpec{Role: "admin"})
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			has, _ := enforcer.HasGroupingPolicy(subject, role3.Role, role3.Domain)
			return !has
		}, waitFor, tick, "Should remove admin role from domain2")

		// Final check: remove all remaining roles for the user
		_, _ = modifier.AddRoles(ctx, subject, role1) // re-add for final test
		time.Sleep(tick * 2)
		_, err = modifier.RemoveRoles(ctx, subject)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			r, _ := enforcer.GetImplicitRolesForUser(subject)
			return len(r) == 0
		}, waitFor, tick, "Should remove all roles for user")
	})

	t.Run("UpdateRole", func(t *testing.T) {
		ctx, modifier, enforcer, cleanup := setupTest(t)
		defer cleanup()

		subject := "user_updater"
		oldRole := authz.RoleSpec{Role: "editor", Domain: domain1}
		newRole := authz.RoleSpec{Role: "publisher", Domain: domain1}

		// Add initial role
		_, err := modifier.AddRoles(ctx, subject, oldRole)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			has, _ := enforcer.HasGroupingPolicy(subject, oldRole.Role, oldRole.Domain)
			return has
		}, waitFor, tick)

		// Update the role
		updated, err := modifier.UpdateRole(ctx, subject, oldRole, newRole)
		require.NoError(t, err)
		assert.True(t, updated)

		require.Eventually(t, func() bool {
			oldHas, _ := enforcer.HasGroupingPolicy(subject, oldRole.Role, oldRole.Domain)
			newHas, _ := enforcer.HasGroupingPolicy(subject, newRole.Role, newRole.Domain)
			return !oldHas && newHas
		}, waitFor, tick, "Enforcer should see updated role but not the old one")
	})

	t.Run("AddAndRemovePermissions", func(t *testing.T) {
		ctx, modifier, enforcer, cleanup := setupTest(t)
		defer cleanup()

		subject := "role_manager"
		perm1 := authz.RuleSpec{Resource: "orders", Action: "read", Domain: domain1}
		perm2 := authz.RuleSpec{Resource: "orders", Action: "write", Domain: domain1}
		perm3 := authz.RuleSpec{Resource: "products", Action: "read", Domain: domain2}

		// Add permissions
		_, err := modifier.AddPermissions(ctx, subject, perm1, perm2, perm3)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			p := enforcer.GetPermissionsForUserInDomain(subject, domain1)
			return len(p) == 2
		}, waitFor, tick)

		// Remove specific permission
		_, err = modifier.RemovePermissions(ctx, subject, perm1)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			has, _ := enforcer.Enforce(subject, perm1.Resource, perm1.Action, perm1.Domain)
			return !has
		}, waitFor, tick)

		// Remove all permissions for a resource in a domain
		_, err = modifier.RemovePermissions(ctx, subject, authz.RuleSpec{Resource: "orders", Domain: domain1})
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			p := enforcer.GetPermissionsForUserInDomain(subject, domain1)
			return len(p) == 0
		}, waitFor, tick)

		// Remove all permissions for the subject
		_, err = modifier.RemovePermissions(ctx, subject)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			p, _ := enforcer.GetImplicitPermissionsForUser(subject)
			return len(p) == 0
		}, waitFor, tick)
	})

	t.Run("UpdatePermission", func(t *testing.T) {
		ctx, modifier, enforcer, cleanup := setupTest(t)
		defer cleanup()

		subject := "user_updater"
		oldSpec := authz.RuleSpec{Resource: "profile", Action: "read", Domain: domain1}
		newSpec := authz.RuleSpec{Resource: "profile", Action: "write", Domain: domain1}

		// Add initial permission
		_, err := modifier.AddPermissions(ctx, subject, oldSpec)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			has, _ := enforcer.Enforce(subject, oldSpec.Resource, oldSpec.Action, oldSpec.Domain)
			return has
		}, waitFor, tick)

		// Update the permission
		updated, err := modifier.UpdatePermission(ctx, subject, oldSpec, newSpec)
		require.NoError(t, err)
		assert.True(t, updated)

		require.Eventually(t, func() bool {
			oldHas, _ := enforcer.Enforce(subject, oldSpec.Resource, oldSpec.Action, oldSpec.Domain)
			newHas, _ := enforcer.Enforce(subject, newSpec.Resource, newSpec.Action, newSpec.Domain)
			return !oldHas && newHas
		}, waitFor, tick, "Enforcer should see updated permission but not the old one")
	})

	t.Run("OperationsInDefaultDomain", func(t *testing.T) {
		ctx, modifier, enforcer, cleanup := setupTest(t)
		defer cleanup()

		subject := "user_no_domain"
		role := authz.RoleSpec{Role: "global_admin", Domain: defaultDomain}
		perm := authz.RuleSpec{Resource: "global_settings", Action: "write", Domain: defaultDomain}

		// Add role without domain
		_, err := modifier.AddRoles(ctx, subject, role)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			has, _ := enforcer.HasGroupingPolicy(subject, role.Role, defaultDomain)
			return has
		}, waitFor, tick, "Should add role in default domain")

		// Add permission without domain
		_, err = modifier.AddPermissions(ctx, subject, perm)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			has, _ := enforcer.Enforce(subject, perm.Resource, perm.Action, defaultDomain)
			return has
		}, waitFor, tick, "Should add permission in default domain")

		// Remove role without domain
		_, err = modifier.RemoveRoles(ctx, subject, role)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			has, _ := enforcer.HasGroupingPolicy(subject, role.Role, defaultDomain)
			return !has
		}, waitFor, tick, "Should remove role from default domain")

		// Remove permission without domain
		_, err = modifier.RemovePermissions(ctx, subject, perm)
		require.NoError(t, err)
		require.Eventually(t, func() bool {
			has, _ := enforcer.Enforce(subject, perm.Resource, perm.Action, defaultDomain)
			return !has
		}, waitFor, tick, "Should remove permission from default domain")
	})
}
