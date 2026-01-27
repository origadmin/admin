/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package database

import (
	"context"
	"fmt"
	"testing"

	"entgo.io/ent/dialect"
	"github.com/casbin/casbin/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/sqlite3ent/sqlite3"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/data/entity/ent"
	entcasbinrule "origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
)

// setupCasbinTest creates a test environment for the Casbin adapter.
func setupCasbinTest(t *testing.T) (*ent.Client, *data.CasbinAdapter, *casbin.Enforcer) {
	t.Helper()

	// Create an in-memory database
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	t.Cleanup(func() {
		err := client.Close()
		require.NoError(t, err)
	})

	// Create the adapter
	database := ent.NewDatabaseWithClient(client)
	adapter := &data.CasbinAdapter{Ctx: context.Background(), DB: database}

	// Create the Enforcer
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	require.NoError(t, err, "Enforcer should be created successfully")

	// Clear policies
	err = enforcer.LoadPolicy()
	require.NoError(t, err, "Policies should be cleared successfully")

	return client, adapter, enforcer
}

// clearPolicies is a helper function to clear all policies for a sub-test.
func clearPolicies(t *testing.T, client *ent.Client, enforcer *casbin.Enforcer) {
	t.Helper()
	_, err := client.CasbinRule.Delete().Exec(context.Background())
	require.NoError(t, err)
	enforcer.ClearPolicy()
}

// TestCasbinAdapter_AddPolicy tests adding policies.
func TestCasbinAdapter_AddPolicy(t *testing.T) {
	client, _, enforcer := setupCasbinTest(t)

	t.Run("AddSinglePolicy", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		policy := []string{"user1", "system:user:list", "read", "domain1"}

		// Add a policy
		added, err := enforcer.AddPolicy(policy)
		require.NoError(t, err)
		assert.True(t, added, "Policy should be added successfully")

		// Verify the policy in the database
		ctx := context.Background()
		policies, err := client.CasbinRule.Query().Where(
			entcasbinrule.PtypeEQ("p"),
		).All(ctx)
		require.NoError(t, err)
		assert.Len(t, policies, 1, "There should be one policy")

		// Verify policy values
		p := policies[0]
		assert.Equal(t, "user1", p.V0)
		assert.Equal(t, "system:user:list", p.V1)
		assert.Equal(t, "read", p.V2)
		assert.Equal(t, "domain1", p.V3)

		// Verify the Enforcer can read the policy
		hasPolicy, err := enforcer.HasPolicy(policy)
		require.NoError(t, err)
		assert.True(t, hasPolicy, "Enforcer should be able to read the policy")
	})

	t.Run("AddMultiplePolicies", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		policies := [][]string{
			{"user1", "system:user:list", "read", "domain1"},
			{"user1", "system:user:detail", "read", "domain1"},
			{"user2", "system:role:list", "read", "domain2"},
		}

		// Add multiple policies
		added, err := enforcer.AddPolicies(policies)
		require.NoError(t, err)
		assert.True(t, added, "Should add 3 policies")

		// Verify policies in the database
		ctx := context.Background()
		count, err := client.CasbinRule.Query().Count(ctx)
		require.NoError(t, err)
		assert.Equal(t, 3, count, "There should be 3 policies in the database")

		// Verify the Enforcer can read all policies
		for _, policy := range policies {
			hasPolicy, err := enforcer.HasPolicy(policy)
			require.NoError(t, err)
			assert.True(t, hasPolicy, "Enforcer should be able to read the policy %v", policy)
		}
	})

	t.Run("AddDuplicatePolicy", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		policy := []string{"user1", "system:user:list", "read", "domain1"}

		// Add a policy
		added, err := enforcer.AddPolicy(policy)
		require.NoError(t, err)
		assert.True(t, added)

		// Add the same policy again
		added, err = enforcer.AddPolicy(policy)
		require.NoError(t, err)
		assert.False(t, added, "Duplicate policy should not be added")
	})
}

// TestCasbinAdapter_RemovePolicy tests removing policies.
func TestCasbinAdapter_RemovePolicy(t *testing.T) {
	client, _, enforcer := setupCasbinTest(t)

	t.Run("RemoveSinglePolicy", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		policy := []string{"user1", "system:user:list", "read", "domain1"}

		// Add a policy
		_, err := enforcer.AddPolicy(policy)
		require.NoError(t, err)

		// Remove the policy
		removed, err := enforcer.RemovePolicy(policy)
		require.NoError(t, err)
		assert.True(t, removed, "Policy should be removed successfully")

		// Verify the policy does not exist
		hasPolicy, err := enforcer.HasPolicy(policy)
		require.NoError(t, err)
		assert.False(t, hasPolicy, "Policy should not exist")
	})

	t.Run("RemoveMultiplePolicies", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		policies := [][]string{
			{"user1", "system:user:list", "read", "domain1"},
			{"user1", "system:user:detail", "read", "domain1"},
			{"user2", "system:role:list", "read", "domain2"},
		}

		// Add multiple policies
		_, err := enforcer.AddPolicies(policies)
		require.NoError(t, err)

		// Remove multiple policies
		removed, err := enforcer.RemovePolicies(policies)
		require.NoError(t, err)
		assert.True(t, removed, "Should remove 3 policies")

		// Verify all policies do not exist
		for _, policy := range policies {
			hasPolicy, err := enforcer.HasPolicy(policy)
			require.NoError(t, err)
			assert.False(t, hasPolicy, "Policy %v should not exist", policy)
		}
	})

	t.Run("RemoveFilteredPolicy", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		policies := [][]string{
			{"user1", "system:user:list", "read", "domain1"},
			{"user1", "system:user:detail", "read", "domain1"},
			{"user2", "system:role:list", "read", "domain1"},
		}

		// Add multiple policies
		_, err := enforcer.AddPolicies(policies)
		require.NoError(t, err)

		// Remove all policies for user1 (filter by the first field)
		removed, err := enforcer.RemoveFilteredPolicy(0, "user1")
		require.NoError(t, err)
		assert.True(t, removed, "Should remove some policies")

		// Verify all policies for user1 do not exist
		for _, policy := range policies {
			if policy[0] == "user1" {
				hasPolicy, err := enforcer.HasPolicy(policy)
				require.NoError(t, err)
				assert.False(t, hasPolicy, "Policy %v should not exist", policy)
			}
		}
	})
}

// TestCasbinAdapter_UpdatePolicy tests updating policies.
func TestCasbinAdapter_UpdatePolicy(t *testing.T) {
	client, _, enforcer := setupCasbinTest(t)

	t.Run("UpdateSinglePolicy", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		oldPolicy := []string{"user1", "system:user:list", "read", "domain1"}
		newPolicy := []string{"user1", "system:user:detail", "write", "domain1"}

		// Add the old policy
		_, err := enforcer.AddPolicy(oldPolicy)
		require.NoError(t, err)

		// Update the policy
		updated, err := enforcer.UpdatePolicy(oldPolicy, newPolicy)
		require.NoError(t, err)
		assert.True(t, updated, "Policy should be updated successfully")

		// Verify the old policy does not exist
		hasPolicy, err := enforcer.HasPolicy(oldPolicy)
		require.NoError(t, err)
		assert.False(t, hasPolicy, "Old policy should not exist")

		// Verify the new policy exists
		hasPolicy, err = enforcer.HasPolicy(newPolicy)
		require.NoError(t, err)
		assert.True(t, hasPolicy, "New policy should exist")
	})

	t.Run("UpdatePolicies", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		oldPolicies := [][]string{
			{"user1", "system:user:list", "read", "domain1"},
			{"user2", "system:role:list", "read", "domain2"},
		}
		newPolicies := [][]string{
			{"user1", "system:user:detail", "write", "domain1"},
			{"user2", "system:role:detail", "write", "domain2"},
		}

		// Add old policies
		_, err := enforcer.AddPolicies(oldPolicies)
		require.NoError(t, err)

		// Update multiple policies
		updated, err := enforcer.UpdatePolicies(oldPolicies, newPolicies)
		require.NoError(t, err)
		assert.True(t, updated, "Should update 2 policies")

		// Verify old policies do not exist
		for _, policy := range oldPolicies {
			hasPolicy, err := enforcer.HasPolicy(policy)
			require.NoError(t, err)
			assert.False(t, hasPolicy, "Old policy %v should not exist", policy)
		}

		// Verify new policies exist
		for _, policy := range newPolicies {
			hasPolicy, err := enforcer.HasPolicy(policy)
			require.NoError(t, err)
			assert.True(t, hasPolicy, "New policy %v should exist", policy)
		}
	})

	t.Run("UpdateFilteredPolicies", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		policies := [][]string{
			{"user1", "system:user:list", "read", "domain1"},
			{"user1", "system:user:detail", "read", "domain1"},
			{"user2", "system:role:list", "read", "domain2"},
		}

		// Add multiple policies
		_, err := enforcer.AddPolicies(policies)
		require.NoError(t, err)

		// Update all policies for user1
		newPolicies := [][]string{
			{"user1", "system:user:update", "write", "domain1"},
			{"user1", "system:user:delete", "delete", "domain1"},
		}
		updated, err := enforcer.UpdateFilteredPolicies(newPolicies, 0, "user1")
		require.NoError(t, err)
		assert.True(t, updated, "Should return true for updated policies")

		// Verify old policies do not exist
		oldUser1Policies := [][]string{
			{"user1", "system:user:list", "read", "domain1"},
			{"user1", "system:user:detail", "read", "domain1"},
		}
		for _, policy := range oldUser1Policies {
			hasPolicy, err := enforcer.HasPolicy(policy)
			require.NoError(t, err)
			assert.False(t, hasPolicy, "Old policy %v should not exist", policy)
		}

		// Verify new policies exist
		for _, policy := range newPolicies {
			hasPolicy, err := enforcer.HasPolicy(policy)
			require.NoError(t, err)
			assert.True(t, hasPolicy, "New policy %v should exist", policy)
		}
	})
}

// TestCasbinAdapter_LoadPolicy tests loading policies.
func TestCasbinAdapter_LoadPolicy(t *testing.T) {
	client, _, enforcer := setupCasbinTest(t)

	t.Run("LoadPolicyFromDatabase", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		ctx := context.Background()

		// Insert policies directly into the database
		_, err := client.CasbinRule.Create().
			SetPtype("p").
			SetV0("user1").
			SetV1("system:user:list").
			SetV2("read").
			SetV3("domain1").
			Save(ctx)
		require.NoError(t, err)

		_, err = client.CasbinRule.Create().
			SetPtype("p").
			SetV0("user2").
			SetV1("system:role:list").
			SetV2("read").
			SetV3("domain2").
			Save(ctx)
		require.NoError(t, err)

		// Reload policies
		err = enforcer.LoadPolicy()
		require.NoError(t, err, "Policies should be loaded successfully")

		// Verify policies
		hasPolicy, err := enforcer.HasPolicy("user1", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, hasPolicy, "Should be able to load the first policy")

		hasPolicy, err = enforcer.HasPolicy("user2", "system:role:list", "read", "domain2")
		require.NoError(t, err)
		assert.True(t, hasPolicy, "Should be able to load the second policy")
	})
}

// TestCasbinAdapter_SavePolicy tests saving policies.
func TestCasbinAdapter_SavePolicy(t *testing.T) {
	client, _, enforcer := setupCasbinTest(t)

	t.Run("SavePolicyToDatabase", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		ctx := context.Background()

		// Add policies to the Enforcer
		policies := [][]string{
			{"user1", "system:user:list", "read", "domain1"},
			{"user2", "system:role:list", "read", "domain2"},
		}
		_, err := enforcer.AddPolicies(policies)
		require.NoError(t, err)

		// Save policies to the database
		err = enforcer.SavePolicy()
		require.NoError(t, err, "Policies should be saved successfully")

		// Verify the number of policies in the database
		count, err := client.CasbinRule.Query().Count(ctx)
		require.NoError(t, err)
		assert.Equal(t, 2, count, "There should be 2 policies in the database")

		// Verify policy content
		rules, err := client.CasbinRule.Query().All(ctx)
		require.NoError(t, err)
		assert.Len(t, rules, 2)

		// Verify policy values
		for _, rule := range rules {
			assert.Equal(t, "p", rule.Ptype)
			assert.NotEmpty(t, rule.V0)
		}
	})
}

// TestCasbinAdapter_Transaction tests transaction support.
func TestCasbinAdapter_Transaction(t *testing.T) {
	client, _, enforcer := setupCasbinTest(t)

	t.Run("RollbackOnError", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		policies := [][]string{
			{"user1", "system:user:list", "read", "domain1"},
			{"user1", "system:user:detail", "read", "domain1"},
			{"user2", "system:role:list", "read", "domain2"},
		}

		// Add policies
		_, err := enforcer.AddPolicies(policies)
		require.NoError(t, err)

		// Try to remove a non-existent policy (should fail)
		removed, err := enforcer.RemovePolicy("user3", "invalid", "invalid", "domain1")
		require.NoError(t, err)
		assert.False(t, removed)

		// Verify policies still exist
		for _, policy := range policies {
			hasPolicy, err := enforcer.HasPolicy(policy)
			require.NoError(t, err)
			assert.True(t, hasPolicy, "Policy %v should still exist", policy)
		}
	})
}

// TestCasbinAdapter_RoleInheritance tests role inheritance.
func TestCasbinAdapter_RoleInheritance(t *testing.T) {
	client, _, enforcer := setupCasbinTest(t)

	t.Run("AddRoleInheritance", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		// Add a role policy
		_, err := enforcer.AddPolicy("admin", "system:*", "ANY", "domain1")
		require.NoError(t, err)

		// Add a role inheritance relationship
		_, err = enforcer.AddGroupingPolicy("alice", "admin", "domain1")
		require.NoError(t, err)

		// alice should have admin's permissions
		allowed, err := enforcer.Enforce("alice", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed, "alice should inherit permissions through the role")
	})

	t.Run("RemoveRoleInheritance", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		// Add role policy and inheritance
		_, err := enforcer.AddPolicy("admin", "system:*", "ANY", "domain1")
		require.NoError(t, err)
		_, err = enforcer.AddGroupingPolicy("alice", "admin", "domain1")
		require.NoError(t, err)

		// Verify permission
		allowed, err := enforcer.Enforce("alice", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.True(t, allowed)

		// Remove role inheritance
		removed, err := enforcer.RemoveGroupingPolicy("alice", "admin", "domain1")
		require.NoError(t, err)
		assert.True(t, removed)

		// alice should no longer have permission
		allowed, err = enforcer.Enforce("alice", "system:user:list", "read", "domain1")
		require.NoError(t, err)
		assert.False(t, allowed, "alice should no longer have permission")
	})
}

// TestCasbinAdapter_Performance is a performance test (optional).
func TestCasbinAdapter_Performance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test")
	}

	client, _, enforcer := setupCasbinTest(t)

	t.Run("BulkOperations", func(t *testing.T) {
		clearPolicies(t, client, enforcer)
		// Bulk add policies
		policies := make([][]string, 1000)
		for i := 0; i < 1000; i++ {
			policies[i] = []string{
				fmt.Sprintf("user%d", i),
				"system:*",
				"read",
				"domain1",
			}
		}

		_, err := enforcer.AddPolicies(policies)
		require.NoError(t, err, "Bulk adding 1000 policies should succeed")

		// Verify the number of policies
		allPolicies, err := enforcer.GetPolicy()
		require.NoError(t, err)
		assert.Len(t, allPolicies, 1000, "There should be 1000 policies")
	})
}
