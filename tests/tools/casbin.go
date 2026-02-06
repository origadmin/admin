// Copyright 2024 OrigAdmin. All rights reserved.

package tools

import (
	"context"
	"testing"

	"github.com/casbin/casbin/v3"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/data/entity/ent"
)

// SetupTestCasbin creates a test Casbin Enforcer
func SetupTestCasbin(t *testing.T, adapter *data.Adapter, modelPath string) *casbin.Enforcer {
	t.Helper()
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	require.NoError(t, err, "Failed to create enforcer")

	return enforcer
}

// NewTestEnforcer creates a test enforcer with client and model path
func NewTestEnforcer(t *testing.T, client *ent.Client, modelPath string) *casbin.Enforcer {
	t.Helper()
	db := ent.NewDatabaseWithClient(client)
	adapter, err := data.NewAdapter(context.Background(), db, log.DefaultLogger)
	require.NoError(t, err, "Failed to create casbin adapter")

	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	require.NoError(t, err, "Failed to create enforcer")
	return enforcer
}

// AddTestPolicies adds test policies
func AddTestPolicies(t *testing.T, enforcer *casbin.Enforcer, policies [][]string) {
	t.Helper()
	for _, policy := range policies {
		params := make([]interface{}, len(policy))
		for i, v := range policy {
			params[i] = v
		}
		_, err := enforcer.AddPolicy(params...)
		require.NoError(t, err, "Failed to add policy: %v", policy)
	}
}

// AddTestGroupingPolicies adds test grouping policies
func AddTestGroupingPolicies(t *testing.T, enforcer *casbin.Enforcer, groupingPolicies [][]string) {
	t.Helper()
	for _, gp := range groupingPolicies {
		params := make([]interface{}, len(gp))
		for i, v := range gp {
			params[i] = v
		}
		_, err := enforcer.AddGroupingPolicy(params...)
		require.NoError(t, err, "Failed to add grouping policy: %v", gp)
	}
}

// AssertPermission asserts permission
func AssertPermission(t *testing.T, enforcer *casbin.Enforcer, sub, obj, act, dom string, allowed bool) {
	t.Helper()
	result, err := enforcer.Enforce(sub, obj, act, dom)
	require.NoError(t, err)
	require.Equal(t, allowed, result,
		"Permission check failed for sub=%s, obj=%s, act=%s, dom=%s. Expected %v, got %v",
		sub, obj, act, dom, allowed, result)
}

// ClearTestPolicies clears all policies from enforcer
func ClearTestPolicies(t *testing.T, enforcer *casbin.Enforcer) {
	t.Helper()
	enforcer.ClearPolicy()
}
