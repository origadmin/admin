package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	authv1 "origadmin/application/admin/api/v1/services/auth"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	typesv1 "origadmin/application/admin/api/v1/services/types"
)

const (
	baseURL        = "http://localhost:8000"
	adminUser      = "admin"
	adminPass      = "admin123"
	requestTimeout = 5 * time.Second
)

// doRequest performs an HTTP request and returns the response.
func doRequest(t *testing.T, method, path string, body interface{}, token string) *http.Response {
	t.Helper()

	var bodyReader io.Reader
	if body != nil {
		var jsonBytes []byte
		var err error

		// Check if the body is a protobuf message
		if pb, ok := body.(proto.Message); ok {
			jsonBytes, err = protojson.Marshal(pb)
			require.NoError(t, err)
		} else {
			// Fallback to standard JSON marshaling
			jsonBytes, err = json.Marshal(body)
			require.NoError(t, err)
		}
		bodyReader = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequest(method, baseURL+path, bodyReader)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{Timeout: requestTimeout}
	resp, err := client.Do(req)
	require.NoError(t, err)

	return resp
}

// login performs a login request and returns the auth token.
func login(t *testing.T, username, password string) string {
	t.Helper()
	reqBody := &authv1.LoginRequest{
		Username: username,
		Password: password,
	}
	resp := doRequest(t, "POST", "/api/v1/auth/login", reqBody, "")
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read login response body")

	// Check status code and provide detailed error on failure
	require.Equal(t, http.StatusOK, resp.StatusCode, "Login failed. Response: %s", string(bodyBytes))

	var loginResp authv1.LoginResponse
	err = protojson.Unmarshal(bodyBytes, &loginResp)
	require.NoError(t, err, "Failed to unmarshal login response")
	require.NotEmpty(t, loginResp.AccessToken, "Access token should not be empty")

	return loginResp.AccessToken
}

// loginAndGetToken is a convenience helper to log in as the default admin.
func loginAndGetToken(t *testing.T) string {
	t.Helper()
	return login(t, adminUser, adminPass)
}

// createRole is a helper to create a new role.
func createRole(t *testing.T, token, name, keyword string, permissionIDs []int64) int64 {
	t.Helper()
	rolePayload := &typesv1.Role{
		Name:    name,
		Keyword: keyword,
	}
	req := &systemv1.CreateRoleRequest{
		Role:          rolePayload,
		PermissionIds: permissionIDs,
	}
	resp := doRequest(t, "POST", "/api/v1/sys/roles", req, token)
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to create role. Response: %s", string(bodyBytes))

	var createResp systemv1.CreateRoleResponse
	err := protojson.Unmarshal(bodyBytes, &createResp)
	require.NoError(t, err)
	require.NotNil(t, createResp.Role)
	return createResp.Role.Id
}

// createUser is a helper to create a new user.
func createUser(t *testing.T, token, username, password string, roleIDs []int64) int64 {
	t.Helper()
	userPayload := &typesv1.User{
		Username: username,
		Nickname: username,
	}
	req := &systemv1.CreateUserRequest{
		User:     userPayload,
		Password: password,
		RoleIds:  roleIDs,
	}
	resp := doRequest(t, "POST", "/api/v1/sys/users", req, token)
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to create user. Response: %s", string(bodyBytes))

	var createResp systemv1.CreateUserResponse
	err := protojson.Unmarshal(bodyBytes, &createResp)
	require.NoError(t, err)
	require.NotNil(t, createResp.User)
	return createResp.User.Id
}

// deleteResource is a generic helper to delete a resource by its ID.
func deleteResource(t *testing.T, token, path string, id int64, isUser bool) {
	t.Helper()
	url := path + "/" + strconv.FormatInt(id, 10)
	// Special handling for user deletion if needed
	if isUser {
		url += "?force=true"
	}
	resp := doRequest(t, "DELETE", url, nil, token)
	defer resp.Body.Close()
	// We don't strictly require OK, as cleanup should be best-effort.
	// require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to delete resource %s", url)
}
