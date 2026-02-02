/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"context"
	"fmt"
	"testing"

	"github.com/casbin/casbin/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
	"origadmin/application/admin/internal/features/auth/dal"
)

// setupTest now provides a simpler setup without a watcher.
// The responsibility of reloading the policy is moved into each test case,
// simulating the service layer's role.
func setupTest(t *testing.T) (context.Context, authz.PolicyModifier, *casbin.Enforcer, func()) {
	ctx := context.Background()

	// Use an in-memory SQLite database for testing.
	client := enttest.Open(t, "sqlite3", fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", t.Name()))
	db := ent.NewDatabaseWithClient(client)

	// Create the adapter without a watcher.
	adapter, err := data.NewAdapter(context.Background(), db, log.DefaultLogger)
	require.NoError(t, err)

	// Create the modifier.
	modifier, err := dal.NewCasbinModifier(adapter, log.DefaultLogger)
	require.NoError(t, err)

	// Create the enforcer.
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	require.NoError(t, err)

	// Initial policy load.
	err = enforcer.LoadPolicy()
	require.NoError(t, err, "Initial policy load should succeed")

	cleanup := func() {
		client.Close()
	}

	return ctx, modifier, enforcer, cleanup
}

func TestPolicyModifier(t *testing.T) {
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
		require.NoError(t, enforcer.LoadPolicy(), "Load after adding roles")

		// Assertions
		roles := enforcer.GetRolesForUserInDomain(subject, domain1)
		require.Len(t, roles, 2, "Should have 2 roles in domain1")
		roles = enforcer.GetRolesForUserInDomain(subject, domain2)
		require.Len(t, roles, 1, "Should have 1 role in domain2")

		// Remove specific role
		_, err = modifier.RemoveRoles(ctx, subject, role1)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy(), "Load after removing a role")
		has, _ := enforcer.HasGroupingPolicy(subject, role1.Role, role1.Domain)
		require.False(t, has, "Should remove specific role")

		// Remove all roles in a domain
		_, err = modifier.RemoveRoles(ctx, subject, authz.RoleSpec{Domain: domain1})
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy(), "Load after removing roles from domain")
		roles = enforcer.GetRolesForUserInDomain(subject, domain1)
		require.Empty(t, roles, "Should remove all roles in domain1")

		// Remove a role across all domains
		_, err = modifier.RemoveRoles(ctx, subject, authz.RoleSpec{Role: "admin"})
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy(), "Load after removing role across domains")
		has, _ = enforcer.HasGroupingPolicy(subject, role3.Role, role3.Domain)
		require.False(t, has, "Should remove admin role from domain2")

		// Final check: remove all remaining roles for the user
		_, _ = modifier.AddRoles(ctx, subject, role1) // re-add for final test
		require.NoError(t, enforcer.LoadPolicy())
		_, err = modifier.RemoveRoles(ctx, subject)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy(), "Load after removing all roles")
		implicitRoles, _ := enforcer.GetImplicitRolesForUser(subject)
		require.Empty(t, implicitRoles, "Should remove all roles for user")
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
		require.NoError(t, enforcer.LoadPolicy())
		has, _ := enforcer.HasGroupingPolicy(subject, oldRole.Role, oldRole.Domain)
		require.True(t, has)

		// Update the role
		updated, err := modifier.UpdateRole(ctx, subject, oldRole, newRole)
		require.NoError(t, err)
		require.True(t, updated)
		require.NoError(t, enforcer.LoadPolicy())

		oldHas, _ := enforcer.HasGroupingPolicy(subject, oldRole.Role, oldRole.Domain)
		newHas, _ := enforcer.HasGroupingPolicy(subject, newRole.Role, newRole.Domain)
		assert.False(t, oldHas, "Old role should be gone")
		assert.True(t, newHas, "New role should be present")
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
		require.NoError(t, enforcer.LoadPolicy())
		perms := enforcer.GetPermissionsForUserInDomain(subject, domain1)
		require.Len(t, perms, 2)

		// Remove specific permission
		_, err = modifier.RemovePermissions(ctx, subject, perm1)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy())
		has, _ := enforcer.Enforce(subject, perm1.Domain, perm1.Resource, perm1.Action)
		require.False(t, has)

		// Remove all permissions for a resource in a domain
		_, err = modifier.RemovePermissions(ctx, subject, authz.RuleSpec{Resource: "orders", Domain: domain1})
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy())
		perms = enforcer.GetPermissionsForUserInDomain(subject, domain1)
		require.Empty(t, perms)

		// Remove all permissions for the subject
		_, err = modifier.RemovePermissions(ctx, subject)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy())
		implicitPerms, _ := enforcer.GetImplicitPermissionsForUser(subject)
		require.Empty(t, implicitPerms)
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
		require.NoError(t, enforcer.LoadPolicy())
		has, _ := enforcer.Enforce(subject, oldSpec.Domain, oldSpec.Resource, oldSpec.Action)
		require.True(t, has)

		// Update the permission
		updated, err := modifier.UpdatePermission(ctx, subject, oldSpec, newSpec)
		require.NoError(t, err)
		require.True(t, updated)
		require.NoError(t, enforcer.LoadPolicy())

		oldHas, _ := enforcer.Enforce(subject, oldSpec.Domain, oldSpec.Resource, oldSpec.Action)
		newHas, _ := enforcer.Enforce(subject, newSpec.Domain, newSpec.Resource, newSpec.Action)
		assert.False(t, oldHas)
		assert.True(t, newHas)
	})

	t.Run("OperationsInDefaultDomain", func(t *testing.T) {
		ctx, modifier, enforcer, cleanup := setupTest(t)
		defer cleanup()

		subject := "user_no_domain"
		role := authz.RoleSpec{Role: "global_admin", Domain: defaultDomain}
		perm := authz.RuleSpec{Resource: "global_settings", Action: "write", Domain: defaultDomain}

		// Add role
		_, err := modifier.AddRoles(ctx, subject, role)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy())
		has, _ := enforcer.HasGroupingPolicy(subject, role.Role, defaultDomain)
		require.True(t, has, "Should add role in default domain")

		// Add permission
		_, err = modifier.AddPermissions(ctx, subject, perm)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy())
		has, _ = enforcer.Enforce(subject, defaultDomain, perm.Resource, perm.Action)
		require.True(t, has, "Should add permission in default domain")

		// Remove role
		_, err = modifier.RemoveRoles(ctx, subject, role)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy())
		has, _ = enforcer.HasGroupingPolicy(subject, role.Role, defaultDomain)
		require.False(t, has, "Should remove role from default domain")

		// Remove permission
		_, err = modifier.RemovePermissions(ctx, subject, perm)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy())
		has, _ = enforcer.Enforce(subject, defaultDomain, perm.Resource, perm.Action)
		require.False(t, has, "Should remove permission from default domain")
	})

	t.Run("ClearPolicies", func(t *testing.T) {
		ctx, modifier, enforcer, cleanup := setupTest(t)
		defer cleanup()

		subject := "user_clear"
		role1 := authz.RoleSpec{Role: "admin", Domain: domain1}
		perm1 := authz.RuleSpec{Resource: "orders", Action: "read", Domain: domain1}

		// Add policies
		_, err := modifier.AddRoles(ctx, subject, role1)
		require.NoError(t, err)
		_, err = modifier.AddPermissions(ctx, subject, perm1)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy())

		// Verify policies exist
		hasRole, _ := enforcer.HasGroupingPolicy(subject, role1.Role, role1.Domain)
		require.True(t, hasRole)
		hasPerm, _ := enforcer.Enforce(subject, perm1.Domain, perm1.Resource, perm1.Action)
		require.True(t, hasPerm)

		// Clear all policies
		_, err = modifier.ClearPolicies(ctx, subject)
		require.NoError(t, err)
		require.NoError(t, enforcer.LoadPolicy())

		// Verify policies are gone
		roles, _ := enforcer.GetImplicitRolesForUser(subject)
		require.Empty(t, roles, "Should have no roles after clear")
		perms, _ := enforcer.GetImplicitPermissionsForUser(subject)
		require.Empty(t, perms, "Should have no permissions after clear")
	})
}
