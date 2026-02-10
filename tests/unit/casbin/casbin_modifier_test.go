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

	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/contrib/security/authz/casbin/adapter"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
	"origadmin/application/admin/internal/features/system/dal"
)

// setupTest provides a test setup with an in-memory SQLite database,
// a policy modifier, and a Casbin enforcer.
func setupTest(t *testing.T) (context.Context, authz.PolicyModifier, *casbin.Enforcer, func()) {
	ctx := context.Background()

	// Use an in-memory SQLite database for testing.
	client := enttest.Open(t, "sqlite3", fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", t.Name()))
	db := ent.NewDatabaseWithClient(client)

	// Create the modifier.
	modifier, err := dal.NewCasbinModifier(db, log.DefaultLogger)
	require.NoError(t, err)

	// Create the enforcer with the modifier as the adapter.
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath, adapter.NewAdapter(modifier))
	require.NoError(t, err)

	// Initial policy load.
	err = enforcer.LoadPolicy()
	require.NoError(t, err, "Initial policy load should succeed")

	cleanup := func() {
		client.Close()
	}

	return ctx, modifier, enforcer, cleanup
}

func TestAddPolicies(t *testing.T) {
	ctx, modifier, enforcer, cleanup := setupTest(t)
	defer cleanup()

	domain1 := "domain1"
	pPolicy := &authzv1.PolicySpec{
		Type:      "p",
		Subject:   "admin",
		Domain:    &domain1,
		Resources: []string{"data1"},
		Actions:   []string{"read"},
	}
	gPolicy := &authzv1.PolicySpec{
		Type:    "g",
		Subject: "alice",
		Domain:  &domain1,
		Roles:   []string{"admin"},
	}

	// Add policies
	changed, err := modifier.AddPolicies(ctx, pPolicy, gPolicy)
	require.NoError(t, err)
	require.True(t, changed)

	// Verify with enforcer
	require.NoError(t, enforcer.LoadPolicy())

	has, err := enforcer.HasPolicy("admin", domain1, "data1", "read")
	require.NoError(t, err)
	assert.True(t, has)

	has, err = enforcer.HasGroupingPolicy("alice", "admin", domain1)
	require.NoError(t, err)
	assert.True(t, has)

	// Add duplicate policy - should not error and report no change
	changed, err = modifier.AddPolicies(ctx, pPolicy)
	require.NoError(t, err)
	require.False(t, changed, "Adding a duplicate policy should not result in a change")
}

func TestListPolicies(t *testing.T) {
	ctx, modifier, _, cleanup := setupTest(t)
	defer cleanup()

	domain1, domain2 := "domain1", "domain2"
	policies := []*authzv1.PolicySpec{
		{Type: "p", Subject: "sub1", Domain: &domain1, Resources: []string{"res1"}, Actions: []string{"act1"}},
		{Type: "g", Subject: "sub1", Domain: &domain1, Roles: []string{"role1"}},
		{Type: "p", Subject: "sub2", Domain: &domain2, Resources: []string{"res2"}, Actions: []string{"act2"}},
	}
	_, err := modifier.AddPolicies(ctx, policies...)
	require.NoError(t, err)

	// List all
	listed, err := modifier.ListPolicies(ctx, nil)
	require.NoError(t, err)
	assert.Len(t, listed, 3)

	// Filter by type
	listed, err = modifier.ListPolicies(ctx, &authzv1.PolicySpec{Type: "p"})
	require.NoError(t, err)
	assert.Len(t, listed, 2)

	// Filter by subject
	listed, err = modifier.ListPolicies(ctx, &authzv1.PolicySpec{Subject: "sub1"})
	require.NoError(t, err)
	assert.Len(t, listed, 2)

	// Filter by domain
	listed, err = modifier.ListPolicies(ctx, &authzv1.PolicySpec{Domain: &domain2})
	require.NoError(t, err)
	assert.Len(t, listed, 1)
	assert.Equal(t, "sub2", listed[0].Subject)
}

func TestRemovePolicies(t *testing.T) {
	ctx, modifier, enforcer, cleanup := setupTest(t)
	defer cleanup()

	domain1 := "domain1"
	p1 := &authzv1.PolicySpec{Type: "p", Subject: "sub1", Domain: &domain1, Resources: []string{"res1"}, Actions: []string{"act1"}}
	p2 := &authzv1.PolicySpec{Type: "p", Subject: "sub1", Domain: &domain1, Resources: []string{"res2"}, Actions: []string{"act2"}}
	g1 := &authzv1.PolicySpec{Type: "g", Subject: "sub1", Domain: &domain1, Roles: []string{"role1"}}
	_, err := modifier.AddPolicies(ctx, p1, p2, g1)
	require.NoError(t, err)
	require.NoError(t, enforcer.LoadPolicy())

	// Remove a single specific policy
	changed, err := modifier.RemovePolicies(ctx, p1)
	require.NoError(t, err)
	require.True(t, changed)
	require.NoError(t, enforcer.LoadPolicy())

	has, err := enforcer.HasPolicy("sub1", domain1, "res1", "act1")
	require.NoError(t, err)
	assert.False(t, has)

	has, err = enforcer.HasPolicy("sub1", domain1, "res2", "act2")
	require.NoError(t, err)
	assert.True(t, has) // Ensure others remain

	// Remove all policies for a subject
	changed, err = modifier.RemovePolicies(ctx, &authzv1.PolicySpec{Subject: "sub1"})
	require.NoError(t, err)
	require.True(t, changed)
	require.NoError(t, enforcer.LoadPolicy())

	allPolicies, err := enforcer.GetPolicy()
	require.NoError(t, err)
	assert.Empty(t, allPolicies, "All policies for sub1 should be removed")
}

func TestUpdatePolicies(t *testing.T) {
	ctx, modifier, enforcer, cleanup := setupTest(t)
	defer cleanup()

	domain1 := "domain1"
	oldPolicy := &authzv1.PolicySpec{Type: "p", Subject: "sub1", Domain: &domain1, Resources: []string{"res1"}, Actions: []string{"read"}}
	newPolicy := &authzv1.PolicySpec{Type: "p", Subject: "sub1", Domain: &domain1, Resources: []string{"res1"}, Actions: []string{"write"}}
	_, err := modifier.AddPolicies(ctx, oldPolicy)
	require.NoError(t, err)
	require.NoError(t, enforcer.LoadPolicy())

	// Perform update
	changed, err := modifier.UpdatePolicies(ctx, []*authzv1.PolicySpec{oldPolicy}, []*authzv1.PolicySpec{newPolicy})
	require.NoError(t, err)
	require.True(t, changed)
	require.NoError(t, enforcer.LoadPolicy())

	// Verify
	has, err := enforcer.HasPolicy("sub1", domain1, "res1", "read")
	require.NoError(t, err)
	assert.False(t, has, "Old policy should be gone")

	has, err = enforcer.HasPolicy("sub1", domain1, "res1", "write")
	require.NoError(t, err)
	assert.True(t, has, "New policy should exist")

	// Test update non-existent policy
	nonExistentOld := &authzv1.PolicySpec{Type: "p", Subject: "nosub", Domain: &domain1, Resources: []string{"nores"}, Actions: []string{"noact"}}
	_, err = modifier.UpdatePolicies(ctx, []*authzv1.PolicySpec{nonExistentOld}, []*authzv1.PolicySpec{newPolicy})
	require.Error(t, err, "Updating a non-existent policy should return an error")
}

func TestClearPolicies(t *testing.T) {
	ctx, modifier, enforcer, cleanup := setupTest(t)
	defer cleanup()

	domain1 := "domain1"
	policiesToAdd := []*authzv1.PolicySpec{
		{Type: "p", Subject: "sub1", Domain: &domain1, Resources: []string{"res1"}, Actions: []string{"act1"}},
		{Type: "g", Subject: "sub1", Domain: &domain1, Roles: []string{"role1"}},
	}
	_, err := modifier.AddPolicies(ctx, policiesToAdd...)
	require.NoError(t, err)
	require.NoError(t, enforcer.LoadPolicy())

	p, err := enforcer.GetPolicy()
	require.NoError(t, err)
	require.NotEmpty(t, p)

	// Clear all
	changed, err := modifier.ClearPolicies(ctx)
	require.NoError(t, err)
	require.True(t, changed)
	require.NoError(t, enforcer.LoadPolicy())

	p, err = enforcer.GetPolicy()
	require.NoError(t, err)
	assert.Empty(t, p, "Policy should be empty after Clear")

	gp, err := enforcer.GetGroupingPolicy()
	require.NoError(t, err)
	assert.Empty(t, gp, "Grouping policy should be empty after Clear")
}
