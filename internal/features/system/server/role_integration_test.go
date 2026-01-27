/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/sqlite3ent/sqlite3"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	entrole "origadmin/application/admin/internal/data/entity/ent/role"
)

// TestCreateRoleAndAssignPermissionsIntegration tests the creation of a role and assignment of permissions
// using the standard UpdateRole endpoint.
func TestCreateRoleAndAssignPermissionsIntegration(t *testing.T) {
	components := setupTestServer(t)
	server := components.HTTPServer

	// 1. Create necessary permissions
	permView, err := createPermissionAPI(t, server, &types.Permission{Name: "View", Keyword: "perm:view"})
	require.NoError(t, err)
	permEdit, err := createPermissionAPI(t, server, &types.Permission{Name: "Edit", Keyword: "perm:edit"})
	require.NoError(t, err)

	// 2. Create a role without any permissions initially
	role, err := createRoleAPI(t, server, &types.Role{Name: "Contributor", Keyword: "contributor"})
	require.NoError(t, err)

	// 3. Assign permissions using the general-purpose UpdateRole endpoint
	initialPermIDs := []int64{permView.Id, permEdit.Id}
	updatedRoleReq := &systemv1.UpdateRoleRequest{
		Role: &types.Role{
			Id:   role.Id,
			Name: "Contributor", // Name can be updated or stay the same
		},
		PermissionIds: initialPermIDs,
	}
	updatedRole, err := updateRoleAPI(t, server, role.Id, updatedRoleReq)
	require.NoError(t, err)
	require.NotNil(t, updatedRole)

	// 4. Verification
	t.Run("Verify role permissions via API", func(t *testing.T) {
		retrievedRole, err := getRoleAPI(t, server, role.Id, true)
		require.NoError(t, err)
		require.NotNil(t, retrievedRole)
		assert.Len(t, retrievedRole.Permissions, 2)

		retrievedPermIDs := make([]int64, len(retrievedRole.Permissions))
		for i, p := range retrievedRole.Permissions {
			retrievedPermIDs[i] = p.Id
		}
		assert.ElementsMatch(t, initialPermIDs, retrievedPermIDs)
	})
}

// TestUpdateRoleAndPermissionsIntegration tests that updating a role's permissions
// correctly replaces the old set of permissions with the new one.
func TestUpdateRoleAndPermissionsIntegration(t *testing.T) {
	components := setupTestServer(t)
	server := components.HTTPServer
	client := components.DBClient

	// 1. Create permissions
	permA, _ := createPermissionAPI(t, server, &types.Permission{Name: "Perm A", Keyword: "perm:a"})
	permB, _ := createPermissionAPI(t, server, &types.Permission{Name: "Perm B", Keyword: "perm:b"})
	permC, _ := createPermissionAPI(t, server, &types.Permission{Name: "Perm C", Keyword: "perm:c"})

	// 2. Create a role and assign initial permissions (A, B)
	role, _ := createRoleAPI(t, server, &types.Role{Name: "Operator", Keyword: "operator"})
	initialPerms := &systemv1.UpdateRoleRequest{
		Role:          &types.Role{Id: role.Id},
		PermissionIds: []int64{permA.Id, permB.Id},
	}
	_, err := updateRoleAPI(t, server, role.Id, initialPerms)
	require.NoError(t, err)

	// Verify initial state
	dbRole, _ := client.Role.Query().Where(entrole.ID(role.Id)).WithPermissions().Only(components.Ctx)
	require.Len(t, dbRole.Edges.Permissions, 2)

	// 3. Update the role with a new set of permissions (B, C), replacing the old set
	newPerms := &systemv1.UpdateRoleRequest{
		Role:          &types.Role{Id: role.Id, Name: "Senior Operator"}, // Also update the name
		PermissionIds: []int64{permB.Id, permC.Id},
	}
	_, err = updateRoleAPI(t, server, role.Id, newPerms)
	require.NoError(t, err)

	// 4. Verification
	t.Run("Verify permissions were replaced via direct DB query", func(t *testing.T) {
		finalRole, err := client.Role.Query().Where(entrole.ID(role.Id)).WithPermissions().Only(components.Ctx)
		require.NoError(t, err)
		require.NotNil(t, finalRole)

		// Check that the name was updated
		assert.Equal(t, "Senior Operator", finalRole.Name)

		// Check that the permissions were correctly replaced
		assert.Len(t, finalRole.Edges.Permissions, 2)
		finalPermIDs := make([]int64, len(finalRole.Edges.Permissions))
		for i, p := range finalRole.Edges.Permissions {
			finalPermIDs[i] = p.ID
		}
		assert.ElementsMatch(t, []int64{permB.Id, permC.Id}, finalPermIDs, "Permissions should be B and C")
		assert.NotContains(t, finalPermIDs, permA.Id, "Permission A should have been removed")
	})
}

// --- API Helper Functions ---

func createPermissionAPI(t *testing.T, server *httptest.Server, perm *types.Permission) (*types.Permission, error) {
	t.Helper()
	url := server.URL + "/v1/system/permissions"
	body, _ := json.Marshal(&systemv1.CreatePermissionRequest{Permission: perm})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create permission, status: %s, body: %s", resp.Status, body)
	}

	var createResp systemv1.CreatePermissionResponse
	err = json.NewDecoder(resp.Body).Decode(&createResp)
	return createResp.Permission, err
}

func createRoleAPI(t *testing.T, server *httptest.Server, role *types.Role) (*types.Role, error) {
	t.Helper()
	url := server.URL + "/v1/system/roles"
	body, _ := json.Marshal(&systemv1.CreateRoleRequest{Role: role})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create role, status: %s, body: %s", resp.Status, body)
	}

	var createResp systemv1.CreateRoleResponse
	err = json.NewDecoder(resp.Body).Decode(&createResp)
	return createResp.Role, err
}

func updateRoleAPI(t *testing.T, server *httptest.Server, roleID int64, req *systemv1.UpdateRoleRequest) (*types.Role, error) {
	t.Helper()
	url := fmt.Sprintf("%s/v1/system/roles/%d", server.URL, roleID)
	body, _ := json.Marshal(req)

	httpReq, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update role, status: %s, body: %s", resp.Status, body)
	}

	var updateResp systemv1.UpdateRoleResponse
	err = json.NewDecoder(resp.Body).Decode(&updateResp)
	return updateResp.Role, err
}

func getRoleAPI(t *testing.T, server *httptest.Server, roleID int64, withPermissions bool) (*types.Role, error) {
	t.Helper()
	url := fmt.Sprintf("%s/v1/system/roles/%d", server.URL, roleID)
	if withPermissions {
		url += "?with_permissions=true"
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get role, status: %s", resp.Status)
	}

	var getResp systemv1.GetRoleResponse
	err = json.NewDecoder(resp.Body).Decode(&getResp)
	return getResp.Role, err
}
