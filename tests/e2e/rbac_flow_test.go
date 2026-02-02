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
	var userEditorID, userViewerID, userNoRoleID, userEditorTestID, userViewerTestID int64

	const (
		// waitFor is the maximum time to wait for policy propagation.
		waitFor = 30 * time.Second
		// Use a smarter polling strategy: initial delay + longer tick to reduce API calls
		// This reduces log spam and unnecessary backend load
		tick = 2 * time.Second
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
		// The order is important: delete users first, then roles, then permissions, then resources.
		if userEditorID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/users", userEditorID, true)
		}
		if userViewerID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/users", userViewerID, true)
		}
		if userNoRoleID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/users", userNoRoleID, true)
		}
		if userEditorTestID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/users", userEditorTestID, true)
		}
		if userViewerTestID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/users", userViewerTestID, true)
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
		// Resources are shared, so we don't delete them, but ensure they have correct gRPC operations.
		// Revert resources to their original state if necessary, or ensure they are idempotent.
		// Since Step2 finds existing resources, we don't delete them.
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

	// 2b. Ensure Resources have correct gRPC Operation
	// This step is crucial because the AuthZ middleware uses gRPC method names (Operation) for enforcement,
	// while the initial data might only have HTTP paths. We update them to ensure consistency.
	t.Run("Step2b_UpdateResourcesForGRPC", func(t *testing.T) {
		updateResource(t, adminToken, resUserListID, "system:user:list_users", "/api/v1/sys/users", "GET", "/api.v1.services.system.UserService/ListUsers")
		updateResource(t, adminToken, resUserCreateID, "system:user:create_user", "/api/v1/sys/users", "POST", "/api.v1.services.system.UserService/CreateUser")
		updateResource(t, adminToken, resUserDeleteID, "system:user:delete_user", "/api/v1/sys/users/{id}", "DELETE", "/api.v1.services.system.UserService/DeleteUser")
		t.Log("Updated resources with correct gRPC operations.")
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
	t.Run("Step4_CreateUsersAndAssignRoles", func(t *testing.T) {
		userEditorID = createUser(t, adminToken, editorUser, "password123", []int64{roleEditorID})
		t.Logf("Created Editor User (ID: %d) and assigned to Editor Role", userEditorID)

		userViewerID = createUser(t, adminToken, viewerUser, "password123", []int64{roleViewerID})
		t.Logf("Created Viewer User (ID: %d) and assigned to Viewer Role", userViewerID)

		userNoRoleID = createUser(t, adminToken, noRoleUser, "password123", []int64{})
		t.Logf("Created No-Role User (ID: %d) with no roles assigned", userNoRoleID)
	})

	// 4a. Verify Policy Sync by Polling
	t.Run("Step4a_VerifyPolicySync", func(t *testing.T) {
		t.Log("Waiting for policy propagation (async event processing)...")
		// Add initial delay to allow event processing before first check
		time.Sleep(15 * time.Second)

		require.Eventually(t, func() bool {
			token := login(t, editorUser, "password123")
			if token == "" {
				// Only log on first failure to reduce spam
				return false
			}
			resp := doRequest(t, "GET", "/api/v1/sys/users", nil, token)
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return false
			}
			t.Logf("Policy sync verified for user %s (ID: %d)", editorUser, userEditorID)
			return true
		}, waitFor, tick, "Policy did not sync: editor user %s (ID: %d) could not list users within the time limit.", editorUser, userEditorID)
		t.Logf("Verified: Policy is fully synced for the editor user %s (ID: %d).", editorUser, userEditorID)
	})

	// 5. Verify Initial Permissions
	t.Run("Step5_VerifyInitialPermissions", func(t *testing.T) {
		editorToken := login(t, editorUser, "password123")
		require.NotEmpty(t, editorToken, "Editor user %s (ID: %d) login failed", editorUser, userEditorID)

		t.Logf("Checking create permission for editor user %s (ID: %d)", editorUser, userEditorID)
		respCreate := doRequest(t, "POST", "/api/v1/sys/users", &systemv1.CreateUserRequest{User: &typesv1.User{Username: editorTestUser}, Password: "password123"}, editorToken)
		bodyCreate, _ := io.ReadAll(respCreate.Body)
		respCreate.Body.Close()
		require.Equal(t, http.StatusOK, respCreate.StatusCode, "Editor user should be able to create a user. Response: %s", string(bodyCreate))
		var createResp systemv1.CreateUserResponse
		require.NoError(t, protojson.Unmarshal(bodyCreate, &createResp))
		userEditorTestID = createResp.GetUser().GetId()

		t.Logf("Checking delete permission (forbidden) for editor user %s (ID: %d)", editorUser, userEditorID)
		assert.Equal(t, http.StatusForbidden, doRequest(t, "DELETE", "/api/v1/sys/users/"+strconv.FormatInt(userNoRoleID, 10), nil, editorToken).StatusCode, "Editor user should NOT be able to delete a user")

		var viewerToken string
		time.Sleep(2 * time.Second) // Initial delay for event processing
		require.Eventually(t, func() bool {
			token := login(t, viewerUser, "password123")
			if token == "" {
				return false
			}
			viewerToken = token
			resp := doRequest(t, "GET", "/api/v1/sys/users", nil, viewerToken)
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				t.Logf("Policy sync verified for viewer user %s (ID: %d)", viewerUser, userViewerID)
			}
			return resp.StatusCode == http.StatusOK
		}, waitFor, tick, "Policy did not sync: viewer user %s (ID: %d) could not list users.", viewerUser, userViewerID)
		t.Logf("Verified: Policy is fully synced for the viewer user %s (ID: %d).", viewerUser, userViewerID)

		t.Logf("Checking create permission (forbidden) for viewer user %s (ID: %d)", viewerUser, userViewerID)
		assert.Equal(t, http.StatusForbidden, doRequest(t, "POST", "/api/v1/sys/users", &systemv1.CreateUserRequest{User: &typesv1.User{Username: viewerTestUser}, Password: "password123"}, viewerToken).StatusCode, "Viewer user should NOT be able to create a user")

		noRoleToken := login(t, noRoleUser, "password123")
		require.NotEmpty(t, noRoleToken, "No-role user login failed")
		t.Logf("Checking list permission (forbidden) for no-role user %s (ID: %d)", noRoleUser, userNoRoleID)
		assert.Equal(t, http.StatusForbidden, doRequest(t, "GET", "/api/v1/sys/users", nil, noRoleToken).StatusCode, "No-role user should NOT be able to list users")
	})

	// 6. Dynamically Update Role and Verify
	t.Run("Step6_UpdateRoleAndVerify", func(t *testing.T) {
		t.Log("Updating Viewer Role to include Create permission...")
		updateRole(t, adminToken, roleViewerID, viewerRoleName, viewerRoleKeyword, []int64{permUserListID, permUserCreateID})

		// Verify Viewer can now create a user. Re-login inside Eventually to get a fresh token.

		time.Sleep(15 * time.Second) // Initial delay for event processing
		require.Eventually(t, func() bool {
			token := login(t, viewerUser, "password123")
			if token == "" {
				return false
			}
			resp := doRequest(t, "POST", "/api/v1/sys/users", &systemv1.CreateUserRequest{User: &typesv1.User{Username: viewerTestUser}, Password: "password123"}, token)
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return false
			}
			body, _ := io.ReadAll(resp.Body)
			var createResp systemv1.CreateUserResponse
			if err := protojson.Unmarshal(body, &createResp); err != nil {
				return false
			}
			userViewerTestID = createResp.GetUser().GetId()
			t.Logf("Verified viewer user %s (ID: %d) can now create users", viewerUser, userViewerID)
			return true
		}, waitFor, tick, "Viewer user %s (ID: %d) should be able to create a user after role update", viewerUser, userViewerID)
		t.Logf("Verified: Viewer user %s (ID: %d) can now create users.", viewerUser, userViewerID)
	})

	// 7. Revoke Role and Verify
	t.Run("Step7_RevokeRoleAndVerify", func(t *testing.T) {
		t.Logf("Revoking Editor Role from user %s (ID: %d)...", editorUser, userEditorID)
		updateUser(t, adminToken, userEditorID, editorUser, []int64{})

		time.Sleep(15 * time.Second) // Initial delay for event processing
		require.Eventually(t, func() bool {
			token := login(t, editorUser, "password123")
			if token == "" {
				return false
			}
			resp := doRequest(t, "GET", "/api/v1/sys/users", nil, token)
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusForbidden {
				t.Logf("Verified role revocation for user %s (ID: %d)", editorUser, userEditorID)
			}
			return resp.StatusCode == http.StatusForbidden
		}, waitFor, tick, "Former editor user %s (ID: %d) should NOT be able to list users after role revocation", editorUser, userEditorID)
		t.Logf("Verified: Former editor's permissions for user %s (ID: %d) have been revoked.", editorUser, userEditorID)
	})
}
