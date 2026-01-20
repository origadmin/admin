package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	authv1 "origadmin/application/admin/api/v1/services/auth"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	typesv1 "origadmin/application/admin/api/v1/services/types"
)

// Global variables to store state between test steps
var (
	adminToken  string
	editorToken string
	viewerToken string

	// IDs of created resources
	roleEditorID int64
	roleViewerID int64
	userEditorID int64
	userViewerID int64

	// Resource IDs found from system
	resUserListID   int64
	resUserCreateID int64
)

// TestRBACFlow executes the complete Role-Based Access Control workflow.
func TestRBACFlow(t *testing.T) {
	// 1. Admin Login
	t.Run("Step1_AdminLogin", func(t *testing.T) {
		adminToken = loginAndGetToken(t)
		require.NotEmpty(t, adminToken, "Admin login failed")
		t.Log("Admin logged in successfully")
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
		roleEditorID = createRole(t, adminToken, "E2E_Editor", "e2e_editor", []int64{resUserListID, resUserCreateID})
		t.Logf("Created Editor Role ID: %d", roleEditorID)

		// Create Viewer Role with Read-only permissions
		roleViewerID = createRole(t, adminToken, "E2E_Viewer", "e2e_viewer", []int64{resUserListID})
		t.Logf("Created Viewer Role ID: %d", roleViewerID)
	})

	// 4. Create Users and Assign Roles
	t.Run("Step4_CreateUsers", func(t *testing.T) {
		userEditorID = createUser(t, adminToken, "e2e_editor_user", "password123", []int64{roleEditorID})
		t.Logf("Created Editor User ID: %d", userEditorID)

		userViewerID = createUser(t, adminToken, "e2e_viewer_user", "password123", []int64{roleViewerID})
		t.Logf("Created Viewer User ID: %d", userViewerID)
	})

	// 5. Verify Editor Permissions
	t.Run("Step5_VerifyEditor", func(t *testing.T) {
		editorToken = login(t, "e2e_editor_user", "password123")
		require.NotEmpty(t, editorToken)

		// Try to list users (Should Succeed)
		respList := doRequest(t, "GET", "/api/v1/sys/users", nil, editorToken)
		assert.Equal(t, http.StatusOK, respList.StatusCode, "Editor should be able to list users")
		respList.Body.Close()

		// Try to create a user (Should Succeed)
		userPayload := &typesv1.User{Username: "test_by_editor", Nickname: "TestByEditor"}
		req := &systemv1.CreateUserRequest{User: userPayload, Password: "password123"}
		respCreate := doRequest(t, "POST", "/api/v1/sys/users", req, editorToken)
		assert.Equal(t, http.StatusOK, respCreate.StatusCode, "Editor should be able to create a user")
		respCreate.Body.Close()
	})

	// 6. Verify Viewer Permissions
	t.Run("Step6_VerifyViewer", func(t *testing.T) {
		viewerToken = login(t, "e2e_viewer_user", "password123")
		require.NotEmpty(t, viewerToken)

		// Try to list users (Should Succeed)
		respList := doRequest(t, "GET", "/api/v1/sys/users", nil, viewerToken)
		assert.Equal(t, http.StatusOK, respList.StatusCode, "Viewer should be able to list users")
		respList.Body.Close()

		// Try to create a user (Should Fail)
		userPayload := &typesv1.User{Username: "test_by_viewer", Nickname: "TestByViewer"}
		req := &systemv1.CreateUserRequest{User: userPayload, Password: "password123"}
		respCreate := doRequest(t, "POST", "/api/v1/sys/users", req, viewerToken)
		assert.Equal(t, http.StatusForbidden, respCreate.StatusCode, "Viewer should NOT be able to create a user")
		respCreate.Body.Close()
	})

	// 7. Cleanup
	t.Run("Step7_Cleanup", func(t *testing.T) {
		// Delete users
		deleteResource(t, adminToken, "/api/v1/sys/users", userEditorID)
		deleteResource(t, adminToken, "/api/v1/sys/users", userViewerID)

		// Delete roles
		deleteResource(t, adminToken, "/api/v1/sys/roles", roleEditorID)
		deleteResource(t, adminToken, "/api/v1/sys/roles", roleViewerID)

		t.Log("Cleanup complete")
	})
}

// --- Helper functions ---

func login(t *testing.T, username, password string) string {
	loginReq := &authv1.LoginRequest{
		Username: username,
		Password: password,
	}

	resp := doRequest(t, "POST", "/api/v1/auth/login", loginReq, "")
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		t.Fatalf("Login failed for %s: %d %s", username, resp.StatusCode, string(bodyBytes))
	}

	var loginResp authv1.LoginResponse
	respBody, _ := io.ReadAll(resp.Body)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	err := unmarshaler.Unmarshal(respBody, &loginResp)
	require.NoError(t, err)

	return loginResp.AccessToken
}

func createRole(t *testing.T, token, name, keyword string, resourceIDs []int64) int64 {
	reqBody := map[string]interface{}{
		"role": map[string]interface{}{
			"name":    name,
			"keyword": keyword,
			"status":  1,
		},
		"resource_ids": resourceIDs, // Assign resources on creation
	}

	resp := doRequest(t, "POST", "/api/v1/sys/roles", reqBody, token)
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to create role: %s", string(bodyBytes))

	var resMap map[string]interface{}
	json.Unmarshal(bodyBytes, &resMap)

	roleMap := resMap["role"].(map[string]interface{})
	idVal := roleMap["id"]

	var id int64
	switch v := idVal.(type) {
	case float64:
		id = int64(v)
	case string:
		fmt.Sscanf(v, "%d", &id)
	}

	return id
}

func createUser(t *testing.T, token, username, password string, roleIDs []int64) int64 {
	user := &typesv1.User{
		Username: username,
		Nickname: username,
		Email:    username + "@example.com",
		Status:   1,
	}

	req := &systemv1.CreateUserRequest{
		User:     user,
		Password: password,
		RoleIds:  roleIDs,
	}

	resp := doRequest(t, "POST", "/api/v1/sys/users", req, token)
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to create user: %s", string(bodyBytes))

	var createResp systemv1.CreateUserResponse
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	err := unmarshaler.Unmarshal(bodyBytes, &createResp)
	require.NoError(t, err)

	return createResp.User.Id
}

func deleteResource(t *testing.T, token, path string, id int64) {
	fullPath := fmt.Sprintf("%s/%d", path, id)
	resp := doRequest(t, "DELETE", fullPath, nil, token)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Failed to delete resource at %s", fullPath)
}
