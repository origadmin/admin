package e2e

import (
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	systemv1 "origadmin/application/admin/api/v1/services/system"
	typesv1 "origadmin/application/admin/api/v1/services/types"
)

// TestRBACFlow executes the complete Role-Based Access Control workflow,
// including dynamic permission updates and role revocation.
func TestRBACFlow(t *testing.T) {
	// Generate a unique suffix for this test run to ensure data isolation.
	uniqueSuffix := strconv.FormatInt(time.Now().UnixNano(), 36)
	editorUser := "eeu_" + uniqueSuffix
	viewerUser := "evu_" + uniqueSuffix
	noRoleUser := "nru_" + uniqueSuffix
	editorRoleName := "EER_" + uniqueSuffix
	editorRoleKeyword := "eek_" + uniqueSuffix
	viewerRoleName := "EVR_" + uniqueSuffix
	viewerRoleKeyword := "evk_" + uniqueSuffix
	editorTestUser := "tbe_" + uniqueSuffix
	viewerTestUser := "tbv_" + uniqueSuffix

	var adminToken string
	var resUserListID, resUserCreateID, resUserDeleteID int64
	var permUserListID, permUserCreateID, permUserDeleteID int64
	var roleEditorID, roleViewerID int64
	var userEditorID, userViewerID, userNoRoleID int64

	const (
		waitFor = 5 * time.Second
		tick    = 200 * time.Millisecond
	)

	// 1. Admin Login
	t.Run("Step1_AdminLogin", func(t *testing.T) {
		adminToken = loginAndGetToken(t)
		require.NotEmpty(t, adminToken, "Admin login failed")
		t.Log("Admin logged in successfully")
	})

	// Defer cleanup to ensure it runs even if the test fails.
	t.Cleanup(func() {
		t.Log("Performing post-test cleanup...")
		// Use assert to ensure all cleanup attempts are made.
		if userEditorID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/users", userEditorID, true)
		}
		if userViewerID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/users", userViewerID, true)
		}
		if userNoRoleID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/users", userNoRoleID, true)
		}
		if roleEditorID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/roles", roleEditorID, false)
		}
		if roleViewerID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/roles", roleViewerID, false)
		}
		if permUserListID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/permissions", permUserListID, false)
		}
		if permUserCreateID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/permissions", permUserCreateID, false)
		}
		if permUserDeleteID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/permissions", permUserDeleteID, false)
		}
		t.Log("Post-test cleanup complete.")
	})

	// 2. Find Core Resources (APIs)
	t.Run("Step2_FindResources", func(t *testing.T) {
		resp := doRequest(t, "GET", "/api/v1/sys/resources?page_size=1000", nil, adminToken)
		defer resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)
		require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to find resources. Response: %s", string(bodyBytes))

		var listResp systemv1.ListResourcesResponse
		err := protojson.Unmarshal(bodyBytes, &listResp)
		require.NoError(t, err)

		for _, res := range listResp.Resources {
			switch res.Keyword {
			case "system:user:list_users":
				resUserListID = res.Id
			case "system:user:create_user":
				resUserCreateID = res.Id
			case "system:user:delete_user":
				resUserDeleteID = res.Id
			}
		}
		require.NotZero(t, resUserListID, "Resource 'system:user:list_users' should be found")
		require.NotZero(t, resUserCreateID, "Resource 'system:user:create_user' should be found")
		require.NotZero(t, resUserDeleteID, "Resource 'system:user:delete_user' should be found")
		t.Logf("Found resource IDs: UserList=%d, UserCreate=%d, UserDelete=%d", resUserListID, resUserCreateID, resUserDeleteID)
	})

	// 2a. Create Permissions for Resources
	t.Run("Step2a_CreatePermissions", func(t *testing.T) {
		permUserListID = createPermission(t, adminToken, "Perm_UserList_"+uniqueSuffix, "system:user:list_users:"+uniqueSuffix, []int64{resUserListID})
		t.Logf("Created Permission for UserList (ID: %d)", permUserListID)

		permUserCreateID = createPermission(t, adminToken, "Perm_UserCreate_"+uniqueSuffix, "system:user:create_user:"+uniqueSuffix, []int64{resUserCreateID})
		t.Logf("Created Permission for UserCreate (ID: %d)", permUserCreateID)

		permUserDeleteID = createPermission(t, adminToken, "Perm_UserDelete_"+uniqueSuffix, "system:user:delete_user:"+uniqueSuffix, []int64{resUserDeleteID})
		t.Logf("Created Permission for UserDelete (ID: %d)", permUserDeleteID)
	})

	// 3. Create Roles and Assign Initial Permissions
	t.Run("Step3_CreateRoles", func(t *testing.T) {
		// Editor Role: Can list and create users.
		roleEditorID = createRole(t, adminToken, editorRoleName, editorRoleKeyword, []int64{permUserListID, permUserCreateID})
		t.Logf("Created Editor Role (ID: %d) with List and Create permissions", roleEditorID)

		// Viewer Role: Can only list users.
		roleViewerID = createRole(t, adminToken, viewerRoleName, viewerRoleKeyword, []int64{permUserListID})
		t.Logf("Created Viewer Role (ID: %d) with List-only permission", roleViewerID)
	})

	// 4. Create Users and Assign Roles
	t.Run("Step4_CreateUsers", func(t *testing.T) {
		userEditorID = createUser(t, adminToken, editorUser, "password123", []int64{roleEditorID})
		t.Logf("Created Editor User (ID: %d) and assigned to Editor Role", userEditorID)

		userViewerID = createUser(t, adminToken, viewerUser, "password123", []int64{roleViewerID})
		t.Logf("Created Viewer User (ID: %d) and assigned to Viewer Role", userViewerID)

		userNoRoleID = createUser(t, adminToken, noRoleUser, "password123", []int64{})
		t.Logf("Created No-Role User (ID: %d) with no roles assigned", userNoRoleID)
	})

	// 4a. Debug Step: Verify Associations
	t.Run("Step4a_DebugVerifyAssociations", func(t *testing.T) {
		// Verify user-role association
		userResp := doRequest(t, "GET", "/api/v1/sys/users/"+strconv.FormatInt(userEditorID, 10), nil, adminToken)
		defer userResp.Body.Close()
		userBody, _ := io.ReadAll(userResp.Body)
		require.Equal(t, http.StatusOK, userResp.StatusCode, "Failed to get user details. Response: %s", string(userBody))
		var userDetails systemv1.GetUserResponse
		require.NoError(t, protojson.Unmarshal(userBody, &userDetails))
		require.Len(t, userDetails.GetUser().GetRoles(), 1, "Editor user should have 1 role")
		assert.Equal(t, roleEditorID, userDetails.GetUser().GetRoles()[0].Id, "Editor user should be associated with the editor role")
		t.Logf("Verified: User %d is correctly associated with Role %d", userEditorID, roleEditorID)

		// Verify role-permission association
		roleResp := doRequest(t, "GET", "/api/v1/sys/roles/"+strconv.FormatInt(roleEditorID, 10), nil, adminToken)
		defer roleResp.Body.Close()
		roleBody, _ := io.ReadAll(roleResp.Body)
		require.Equal(t, http.StatusOK, roleResp.StatusCode, "Failed to get role details. Response: %s", string(roleBody))
		var roleDetails systemv1.GetRoleResponse
		require.NoError(t, protojson.Unmarshal(roleBody, &roleDetails))
		require.Len(t, roleDetails.GetRole().GetPermissions(), 2, "Editor role should have 2 permissions")

		permIDs := make(map[int64]bool)
		for _, p := range roleDetails.GetRole().GetPermissions() {
			permIDs[p.Id] = true
		}
		assert.True(t, permIDs[permUserListID], "Editor role should have list permission")
		assert.True(t, permIDs[permUserCreateID], "Editor role should have create permission")
		t.Logf("Verified: Role %d is correctly associated with its permissions", roleEditorID)
	})

	// 5. Verify Initial Permissions
	t.Run("Step5_VerifyInitialPermissions", func(t *testing.T) {
		// Editor
		editorToken := login(t, editorUser, "password123")
		require.NotEmpty(t, editorToken, "Editor user login failed")
		require.Eventually(t, func() bool {
			return doRequest(t, "GET", "/api/v1/sys/users", nil, editorToken).StatusCode == http.StatusOK
		}, waitFor, tick, "Editor should be able to list users")
		require.Eventually(t, func() bool {
			return doRequest(t, "POST", "/api/v1/sys/users", &systemv1.CreateUserRequest{User: &typesv1.User{Username: editorTestUser}, Password: "password123"}, editorToken).StatusCode == http.StatusOK
		}, waitFor, tick, "Editor should be able to create a user")
		// This permission was never granted, so it should be forbidden immediately.
		assert.Equal(t, http.StatusForbidden, doRequest(t, "DELETE", "/api/v1/sys/users/"+strconv.FormatInt(userNoRoleID, 10), nil, editorToken).StatusCode, "Editor should NOT be able to delete a user")

		// Viewer
		viewerToken := login(t, viewerUser, "password123")
		require.NotEmpty(t, viewerToken, "Viewer user login failed")
		require.Eventually(t, func() bool {
			return doRequest(t, "GET", "/api/v1/sys/users", nil, viewerToken).StatusCode == http.StatusOK
		}, waitFor, tick, "Viewer should be able to list users")
		// This permission was never granted, so it should be forbidden immediately.
		assert.Equal(t, http.StatusForbidden, doRequest(t, "POST", "/api/v1/sys/users", &systemv1.CreateUserRequest{User: &typesv1.User{Username: viewerTestUser}, Password: "password123"}, viewerToken).StatusCode, "Viewer should NOT be able to create a user")

		// No-Role User
		noRoleToken := login(t, noRoleUser, "password123")
		require.NotEmpty(t, noRoleToken, "No-role user login failed")
		assert.Equal(t, http.StatusForbidden, doRequest(t, "GET", "/api/v1/sys/users", nil, noRoleToken).StatusCode, "No-role user should NOT be able to list users")
	})

	// 6. Dynamically Update Role and Verify
	t.Run("Step6_UpdateRoleAndVerify", func(t *testing.T) {
		t.Log("Updating Viewer Role to include Create permission...")
		updateRole(t, adminToken, roleViewerID, viewerRoleName, viewerRoleKeyword, []int64{permUserListID, permUserCreateID})

		viewerToken := login(t, viewerUser, "password123")
		require.NotEmpty(t, viewerToken)

		// Verify Viewer can now create a user
		require.Eventually(t, func() bool {
			resp := doRequest(t, "POST", "/api/v1/sys/users", &systemv1.CreateUserRequest{User: &typesv1.User{Username: viewerTestUser}, Password: "password123"}, viewerToken)
			return resp.StatusCode == http.StatusOK
		}, waitFor, tick, "Viewer should be able to create a user after role update")
		t.Log("Verified: Viewer can now create users.")
	})

	// 7. Revoke Role and Verify
	t.Run("Step7_RevokeRoleAndVerify", func(t *testing.T) {
		t.Logf("Revoking Editor Role from user %s...", editorUser)
		updateUser(t, adminToken, userEditorID, editorUser, []int64{}) // Update user with empty role list

		editorToken := login(t, editorUser, "password123")
		require.NotEmpty(t, editorToken)

		// Verify Editor can no longer list users
		require.Eventually(t, func() bool {
			resp := doRequest(t, "GET", "/api/v1/sys/users", nil, editorToken)
			return resp.StatusCode == http.StatusForbidden
		}, waitFor, tick, "Former editor should NOT be able to list users after role revocation")
		t.Log("Verified: Former editor's permissions have been revoked.")
	})
}

func createPermission(t *testing.T, token, name, keyword string, resourceIDs []int64) int64 {
	t.Helper()
	permPayload := &typesv1.Permission{
		Name:    name,
		Keyword: keyword,
	}
	req := &systemv1.CreatePermissionRequest{
		Permission:  permPayload,
		ResourceIds: resourceIDs,
	}
	resp := doRequest(t, "POST", "/api/v1/sys/permissions", req, token)
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to create permission. Response: %s", string(bodyBytes))

	var createResp systemv1.CreatePermissionResponse
	err := protojson.Unmarshal(bodyBytes, &createResp)
	require.NoError(t, err)
	require.NotZero(t, createResp.GetPermission().GetId())
	return createResp.GetPermission().GetId()
}

func updateRole(t *testing.T, token string, roleID int64, name, keyword string, permissionIDs []int64) {
	t.Helper()
	rolePayload := &typesv1.Role{
		Id:      roleID,
		Name:    name,
		Keyword: keyword,
	}
	req := &systemv1.UpdateRoleRequest{
		Role:          rolePayload,
		PermissionIds: permissionIDs,
	}
	url := "/api/v1/sys/roles/" + strconv.FormatInt(roleID, 10)
	resp := doRequest(t, "PUT", url, req, token)
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to update role. Response: %s", string(bodyBytes))
}

func updateUser(t *testing.T, token string, userID int64, username string, roleIDs []int64) {
	t.Helper()
	userPayload := &typesv1.User{
		Id:       userID,
		Username: username,
	}
	req := &systemv1.UpdateUserRequest{
		User:    userPayload,
		RoleIds: roleIDs,
	}
	url := "/api/v1/sys/users/" + strconv.FormatInt(userID, 10)
	resp := doRequest(t, "PUT", url, req, token)
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to update user. Response: %s", string(bodyBytes))
}
