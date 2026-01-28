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

// TestRBACFlow executes the complete Role-Based Access Control workflow.
func TestRBACFlow(t *testing.T) {
	// Generate a unique suffix for this test run to ensure data isolation.
	uniqueSuffix := strconv.FormatInt(time.Now().Unix(), 36)
	editorUser := "eeu_" + uniqueSuffix
	viewerUser := "evu_" + uniqueSuffix
	editorRoleName := "EER_" + uniqueSuffix
	editorRoleKeyword := "eek_" + uniqueSuffix
	viewerRoleName := "EVR_" + uniqueSuffix
	viewerRoleKeyword := "evk_" + uniqueSuffix
	editorTestUser := "tbe_" + uniqueSuffix
	viewerTestUser := "tbv_" + uniqueSuffix

	var adminToken string
	var resUserListID, resUserCreateID int64
	var roleEditorID, roleViewerID int64
	var userEditorID, userViewerID int64

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
		if roleEditorID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/roles", roleEditorID, false)
		}
		if roleViewerID != 0 {
			deleteResource(t, adminToken, "/api/v1/sys/roles", roleViewerID, false)
		}
		t.Log("Post-test cleanup complete.")
	})

	// 2. Find Resources (APIs)
	t.Run("Step2_FindResources", func(t *testing.T) {
		resp := doRequest(t, "GET", "/api/v1/sys/resources?page_size=1000", nil, adminToken)
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)

		var listResp systemv1.ListResourcesResponse
		bodyBytes, _ := io.ReadAll(resp.Body)
		err := protojson.Unmarshal(bodyBytes, &listResp)
		require.NoError(t, err)

		for _, res := range listResp.Resources {
			if res.Keyword == "system:user:list_users" {
				resUserListID = res.Id
			}
			if res.Keyword == "system:user:create_user" {
				resUserCreateID = res.Id
			}
		}
		require.NotZero(t, resUserListID, "Resource 'system:user:list_users' should be found")
		require.NotZero(t, resUserCreateID, "Resource 'system:user:create_user' should be found")
		t.Logf("Found resource IDs: UserList=%d, UserCreate=%d", resUserListID, resUserCreateID)
	})

	// 3. Create Roles and Assign Permissions
	t.Run("Step3_CreateRoles", func(t *testing.T) {
		// Create Editor Role with Read & Write permissions
		roleEditorID = createRole(t, adminToken, editorRoleName, editorRoleKeyword, []int64{resUserListID, resUserCreateID})
		t.Logf("Created Editor Role ID: %d", roleEditorID)

		// Create Viewer Role with Read-only permissions
		roleViewerID = createRole(t, adminToken, viewerRoleName, viewerRoleKeyword, []int64{resUserListID})
		t.Logf("Created Viewer Role ID: %d", roleViewerID)
	})

	// 4. Create Users and Assign Roles
	t.Run("Step4_CreateUsers", func(t *testing.T) {
		userEditorID = createUser(t, adminToken, editorUser, "password123", []int64{roleEditorID})
		t.Logf("Created Editor User ID: %d", userEditorID)

		userViewerID = createUser(t, adminToken, viewerUser, "password123", []int64{roleViewerID})
		t.Logf("Created Viewer User ID: %d", userViewerID)
	})

	// 5. Verify Editor Permissions
	t.Run("Step5_VerifyEditor", func(t *testing.T) {
		editorToken := login(t, editorUser, "password123")
		require.NotEmpty(t, editorToken)

		// Try to list users (Should Succeed)
		respList := doRequest(t, "GET", "/api/v1/sys/users", nil, editorToken)
		assert.Equal(t, http.StatusOK, respList.StatusCode, "Editor should be able to list users")
		respList.Body.Close()

		// Try to create a user (Should Succeed)
		userPayload := &typesv1.User{Username: editorTestUser, Nickname: "TestByEditor"}
		req := &systemv1.CreateUserRequest{User: userPayload, Password: "password123"}
		respCreate := doRequest(t, "POST", "/api/v1/sys/users", req, editorToken)

		// Cleanup the user created by the editor immediately
		if respCreate.StatusCode == http.StatusOK {
			var createResp systemv1.CreateUserResponse
			bodyBytes, _ := io.ReadAll(respCreate.Body)
			err := protojson.Unmarshal(bodyBytes, &createResp)
			if err == nil {
				deleteResource(t, adminToken, "/api/v1/sys/users", createResp.User.Id, true)
			}
		}
		assert.Equal(t, http.StatusOK, respCreate.StatusCode, "Editor should be able to create a user")
		respCreate.Body.Close()
	})

	// 6. Verify Viewer Permissions
	t.Run("Step6_VerifyViewer", func(t *testing.T) {
		viewerToken := login(t, viewerUser, "password123")
		require.NotEmpty(t, viewerToken)

		// Try to list users (Should Succeed)
		respList := doRequest(t, "GET", "/api/v1/sys/users", nil, viewerToken)
		assert.Equal(t, http.StatusOK, respList.StatusCode, "Viewer should be able to list users")
		respList.Body.Close()

		// Try to create a user (Should Fail)
		userPayload := &typesv1.User{Username: viewerTestUser, Nickname: "TestByViewer"}
		req := &systemv1.CreateUserRequest{User: userPayload, Password: "password123"}
		respCreate := doRequest(t, "POST", "/api/v1/sys/users", req, viewerToken)
		assert.Equal(t, http.StatusForbidden, respCreate.StatusCode, "Viewer should NOT be able to create a user")
		respCreate.Body.Close()
	})
}
