/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCasbinModelRules verifies the core logic of the Casbin model,
// including role inheritance, resource matching, and action matching.
func TestCasbinModelRules(t *testing.T) {
	// Load the Casbin model from the project's resource file.
	// The path is relative to the project root.
	modelPath := "../../../resources/casbin_model.conf"
	enforcer, err := casbin.NewEnforcer(modelPath)
	require.NoError(t, err, "Failed to create Casbin enforcer from model file")

	// --- Test Data Setup ---
	// Define policies (p) and role assignments (g).
	// These represent the rules that would be loaded from the database.
	policies := [][]string{
		// p, subject, domain, object, action
		{"admin", "domain1", "/api/v1/users/:id", "GET"},
		{"admin", "domain1", "/api/v1/users", "POST"},
		{"editor", "domain1", "/api/v1/posts/:id", "PUT"},
		{"viewer", "domain1", "/api/v1/posts/:id", "GET"},
		{"viewer", "domain1", "/api/v1/posts", "GET"},
		// Policy for a different domain
		{"domain2_admin", "domain2", "/api/v1/settings", "ANY"},
	}

	groupingPolicies := [][]string{
		// g, user, role, domain
		{"alice", "admin", "domain1"},
		{"bob", "editor", "domain1"},
		{"carol", "viewer", "domain1"},
		// User with multiple roles in the same domain
		{"dave", "editor", "domain1"},
		{"dave", "viewer", "domain1"},
		// User in a different domain
		{"eve", "domain2_admin", "domain2"},
	}

	// Add policies and groupings to the enforcer
	_, err = enforcer.AddPolicies(policies)
	require.NoError(t, err)
	_, err = enforcer.AddGroupingPolicies(groupingPolicies)
	require.NoError(t, err)

	// --- Test Cases ---
	testCases := []struct {
		name     string
		sub      string
		dom      string // Moved dom before obj
		obj      string
		act      string
		expected bool
	}{
		// Admin tests (domain1)
		{"Admin can get a specific user", "alice", "domain1", "/api/v1/users/123", "GET", true},
		{"Admin can create a user", "alice", "domain1", "/api/v1/users", "POST", true},
		{"Admin cannot edit a post", "alice", "domain1", "/api/v1/posts/456", "PUT", false},

		// Editor tests (domain1)
		{"Editor can update a post", "bob", "domain1", "/api/v1/posts/456", "PUT", true},
		{"Editor cannot get a user", "bob", "domain1", "/api/v1/users/123", "GET", false},

		// Viewer tests (domain1)
		{"Viewer can get a specific post", "carol", "domain1", "/api/v1/posts/789", "GET", true},
		{"Viewer can list posts", "carol", "domain1", "/api/v1/posts", "GET", true},
		{"Viewer cannot update a post", "carol", "domain1", "/api/v1/posts/789", "PUT", false},

		// Multi-role user tests (domain1)
		{"Multi-role user can edit a post", "dave", "domain1", "/api/v1/posts/1", "PUT", true},
		{"Multi-role user can view a post", "dave", "domain1", "/api/v1/posts/2", "GET", true},
		{"Multi-role user cannot access admin resources", "dave", "domain1", "/api/v1/users/3", "GET", false},

		// Domain separation tests
		{"Admin from domain1 cannot access domain2 resources", "alice", "domain2", "/api/v1/settings", "GET", false},
		{"Admin from domain2 can access domain2 resources (ANY action)", "eve", "domain2", "/api/v1/settings", "GET", true},
		{"Admin from domain2 can access domain2 resources (POST action)", "eve", "domain2", "/api/v1/settings", "POST", true},
		{"Admin from domain2 cannot access domain1 resources", "eve", "domain1", "/api/v1/users/1", "GET", false},

		// Negative cases
		{"Unknown user has no access", "frank", "domain1", "/api/v1/users/1", "GET", false},
		{"Access denied for undefined object", "alice", "domain1", "/api/v1/nonexistent", "GET", false},
		{"Access denied for undefined action", "alice", "domain1", "/api/v1/users/1", "DELETE", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform the enforcement check
			hasPermission, err := enforcer.Enforce(tc.sub, tc.dom, tc.obj, tc.act)
			require.NoError(t, err, "Enforce call failed")
			assert.Equal(t, tc.expected, hasPermission, "Permission mismatch")
		})
	}
}
