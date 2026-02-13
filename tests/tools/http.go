// Copyright 2024 OrigAdmin. All rights reserved.

package tools

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

	identityv1 "origadmin/application/admin/api/v1/services/identity"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	typesv1 "origadmin/application/admin/api/v1/services/types"
)

// BaseHTTPClient provides basic HTTP client functionality for testing
type BaseHTTPClient struct {
	client  *http.Client
	baseURL string
}

// NewBaseHTTPClient creates a new base HTTP client
func NewBaseHTTPClient(baseURL string) *BaseHTTPClient {
	return &BaseHTTPClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
	}
}

// NewBaseHTTPClientWithTimeout creates a new base HTTP client with custom timeout
func NewBaseHTTPClientWithTimeout(baseURL string, timeout time.Duration) *BaseHTTPClient {
	return &BaseHTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

// Request sends HTTP request with automatic detection of protobuf messages
func (c *BaseHTTPClient) Request(t *testing.T, method, path string, body interface{}, token string) *http.Response {
	t.Helper()

	var bodyReader io.Reader
	if body != nil {
		// Check if body is a protobuf message
		if protoMsg, ok := body.(proto.Message); ok {
			jsonBytes, err := protojson.Marshal(protoMsg)
			require.NoError(t, err, "Failed to marshal protobuf to JSON")
			bodyReader = bytes.NewReader(jsonBytes)
		} else {
			// Regular struct, use standard JSON marshaling
			jsonBody, err := json.Marshal(body)
			require.NoError(t, err)
			bodyReader = bytes.NewReader(jsonBody)
		}
	} else {
		bodyReader = http.NoBody
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := c.client.Do(req)
	require.NoError(t, err)

	return resp
}

// Get sends GET request
func (c *BaseHTTPClient) Get(t *testing.T, path, token string) *http.Response {
	t.Helper()
	return c.Request(t, "GET", path, nil, token)
}

// Post sends POST request
func (c *BaseHTTPClient) Post(t *testing.T, path string, body interface{}, token string) *http.Response {
	t.Helper()
	return c.Request(t, "POST", path, body, token)
}

// Put sends PUT request
func (c *BaseHTTPClient) Put(t *testing.T, path string, body interface{}, token string) *http.Response {
	t.Helper()
	return c.Request(t, "PUT", path, body, token)
}

// Delete sends DELETE request
func (c *BaseHTTPClient) Delete(t *testing.T, path, token string) *http.Response {
	t.Helper()
	return c.Request(t, "DELETE", path, nil, token)
}

// SetBaseURL updates the base URL for the client
func (c *BaseHTTPClient) SetBaseURL(url string) {
	c.baseURL = url
}

// GetBaseURL returns the current base URL
func (c *BaseHTTPClient) GetBaseURL() string {
	return c.baseURL
}

// TestHTTPClient HTTP client for testing with automatic prefix handling.
type TestHTTPClient struct {
	*BaseHTTPClient
	prefix string
}

// NewTestHTTPClient creates a new test HTTP client with default prefix
func NewTestHTTPClient(baseURL string) *TestHTTPClient {
	return &TestHTTPClient{
		BaseHTTPClient: NewBaseHTTPClient(baseURL),
		prefix:         "/api/v1",
	}
}

// NewTestHTTPClientWithPrefix creates a new test HTTP client with custom prefix
func NewTestHTTPClientWithPrefix(baseURL, prefix string) *TestHTTPClient {
	return &TestHTTPClient{
		BaseHTTPClient: NewBaseHTTPClient(baseURL),
		prefix:         prefix,
	}
}

// NewTestHTTPClientWithTimeout creates a new test HTTP client with custom timeout and default prefix
func NewTestHTTPClientWithTimeout(baseURL string, timeout time.Duration) *TestHTTPClient {
	return &TestHTTPClient{
		BaseHTTPClient: NewBaseHTTPClientWithTimeout(baseURL, timeout),
		prefix:         "/api/v1",
	}
}

// Request overrides the embedded BaseHTTPClient's Request method to automatically prepend the prefix.
func (c *TestHTTPClient) Request(t *testing.T, method, path string, body interface{}, token string) *http.Response {
	t.Helper()
	fullPath := c.prefix + path
	return c.BaseHTTPClient.Request(t, method, fullPath, body, token)
}

// Get overrides BaseHTTPClient.Get to use TestHTTPClient.Request
func (c *TestHTTPClient) Get(t *testing.T, path, token string) *http.Response {
	t.Helper()
	return c.Request(t, "GET", path, nil, token)
}

// Post overrides BaseHTTPClient.Post to use TestHTTPClient.Request
func (c *TestHTTPClient) Post(t *testing.T, path string, body interface{}, token string) *http.Response {
	t.Helper()
	return c.Request(t, "POST", path, body, token)
}

// Put overrides BaseHTTPClient.Put to use TestHTTPClient.Request
func (c *TestHTTPClient) Put(t *testing.T, path string, body interface{}, token string) *http.Response {
	t.Helper()
	return c.Request(t, "PUT", path, body, token)
}

// Delete overrides BaseHTTPClient.Delete to use TestHTTPClient.Request
func (c *TestHTTPClient) Delete(t *testing.T, path, token string) *http.Response {
	t.Helper()
	return c.Request(t, "DELETE", path, nil, token)
}

// SetPrefix updates the API prefix for the client
func (c *TestHTTPClient) SetPrefix(prefix string) {
	c.prefix = prefix
}

// GetPrefix returns the current API prefix
func (c *TestHTTPClient) GetPrefix() string {
	return c.prefix
}

// Login performs login request and returns token.
// This method now uses the overridden Request method, so the prefix is handled automatically.
func (c *TestHTTPClient) Login(t *testing.T, username, password string) string {
	t.Helper()
	reqBody := &identityv1.LoginRequest{
		Username: username,
		Password: password,
	}

	// The prefix is now automatically added by the overridden Request method.
	resp := c.Request(t, "POST", "/auth/login", reqBody, "")
	defer resp.Body.Close()

	AssertHTTPStatusCode(t, resp, http.StatusOK)

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read login response body")

	var loginResp identityv1.LoginResponse
	err = protojson.Unmarshal(bodyBytes, &loginResp)
	require.NoError(t, err, "Failed to unmarshal login response")
	require.NotNil(t, loginResp, "Token object should not be nil")
	require.NotEmpty(t, loginResp.AccessToken, "Access token should not be empty")

	return loginResp.AccessToken
}

// ============== System Test Client ==============

// SystemTestClient provides system-specific API testing functionality
type SystemTestClient struct {
	*TestHTTPClient
}

// NewSystemTestClient creates a new system test client
func NewSystemTestClient(baseURL string) *SystemTestClient {
	return &SystemTestClient{
		TestHTTPClient: NewTestHTTPClient(baseURL),
	}
}

// NewSystemTestClientWithPrefix creates a new system test client with custom prefix
func NewSystemTestClientWithPrefix(baseURL, prefix string) *SystemTestClient {
	return &SystemTestClient{
		TestHTTPClient: NewTestHTTPClientWithPrefix(baseURL, prefix),
	}
}

// ============== E2E Test Helpers for System API ==============

// NOTE: System-specific helpers below (CreateRole, CreateUser, etc.) are NOT generic.
// Consider moving to tests/tools/system/ or tests/e2e/helpers/ for better organization.

// CreateRole creates a new role via API (System-specific helper)
func (c *SystemTestClient) CreateRole(t *testing.T, token, name, keyword string, permissionIDs []int64) int64 {
	t.Helper()
	rolePayload := &typesv1.Role{
		Name:    name,
		Keyword: keyword,
	}
	req := &systemv1.CreateRoleRequest{
		Role:          rolePayload,
		PermissionIds: permissionIDs,
	}
	resp := c.Post(t, "/sys/roles", req, token)
	defer resp.Body.Close()

	AssertHTTPStatusCode(t, resp, http.StatusOK)

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	var createResp systemv1.CreateRoleResponse
	err = protojson.Unmarshal(bodyBytes, &createResp)
	require.NoError(t, err)
	require.NotNil(t, createResp.Role)
	return createResp.Role.Id
}

// CreateUser creates a new user via API (System-specific helper)
func (c *SystemTestClient) CreateUser(t *testing.T, token, username, password string, roleIDs []int64) int64 {
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
	resp := c.Post(t, "/sys/users", req, token)
	defer resp.Body.Close()

	AssertHTTPStatusCode(t, resp, http.StatusOK)

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	var createResp systemv1.CreateUserResponse
	err = protojson.Unmarshal(bodyBytes, &createResp)
	require.NoError(t, err)
	require.NotNil(t, createResp.User)
	return createResp.User.Id
}

// CreateResource creates a new resource via API (System-specific helper)
func (c *SystemTestClient) CreateResource(t *testing.T, token, name, keyword, path, method, operation string) int64 {
	t.Helper()
	resPayload := &typesv1.Resource{
		Name:      name,
		Keyword:   keyword,
		Path:      path,
		Method:    method,
		Operation: operation,
	}
	req := &systemv1.CreateResourceRequest{
		Resource: resPayload,
	}
	resp := c.Post(t, "/sys/resources", req, token)
	defer resp.Body.Close()

	AssertHTTPStatusCode(t, resp, http.StatusOK)

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	var createResp systemv1.CreateResourceResponse
	err = protojson.Unmarshal(bodyBytes, &createResp)
	require.NoError(t, err)
	require.NotZero(t, createResp.GetResource().GetId())
	return createResp.GetResource().GetId()
}

// CreatePermission creates a new permission via API (System-specific helper)
func (c *SystemTestClient) CreatePermission(t *testing.T, token, name, keyword string, resourceIDs []int64) int64 {
	t.Helper()
	permPayload := &typesv1.Permission{
		Name:    name,
		Keyword: keyword,
	}
	req := &systemv1.CreatePermissionRequest{
		Permission:  permPayload,
		ResourceIds: resourceIDs,
	}
	resp := c.Post(t, "/sys/permissions", req, token)
	defer resp.Body.Close()

	AssertHTTPStatusCode(t, resp, http.StatusOK)

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	var createResp systemv1.CreatePermissionResponse
	err = protojson.Unmarshal(bodyBytes, &createResp)
	require.NoError(t, err)
	require.NotZero(t, createResp.GetPermission().GetId())
	return createResp.GetPermission().GetId()
}

// UpdateRole updates an existing role via API (System-specific helper)
func (c *SystemTestClient) UpdateRole(t *testing.T, token string, roleID int64, name, keyword string, permissionIDs []int64) {
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
	path := "/sys/roles/" + strconv.FormatInt(roleID, 10)
	resp := c.Put(t, path, req, token)
	defer resp.Body.Close()

	AssertHTTPStatusCode(t, resp, http.StatusOK)
	_, _ = io.ReadAll(resp.Body)
}

// UpdateUser updates an existing user via API (System-specific helper)
func (c *SystemTestClient) UpdateUser(t *testing.T, token string, userID int64, user *typesv1.User, updateMaskPaths []string, roleIDs []int64) {
	t.Helper()
	req := &systemv1.UpdateUserRequest{
		User:    user,
		RoleIds: roleIDs,
	}
	path := "/sys/users/" + strconv.FormatInt(userID, 10)
	resp := c.Put(t, path, req, token)
	defer resp.Body.Close()

	AssertHTTPStatusCode(t, resp, http.StatusOK)
	_, _ = io.ReadAll(resp.Body)
}

// UpdateResource updates an existing resource via API (System-specific helper)
func (c *SystemTestClient) UpdateResource(t *testing.T, token string, resID int64, keyword, path, method, operation string) {
	t.Helper()
	resPayload := &typesv1.Resource{
		Id:        resID,
		Keyword:   keyword,
		Path:      path,
		Method:    method,
		Operation: operation,
	}
	req := &systemv1.UpdateResourceRequest{
		Resource: resPayload,
	}
	urlPath := "/sys/resources/" + strconv.FormatInt(resID, 10)
	resp := c.Put(t, urlPath, req, token)
	defer resp.Body.Close()

	AssertHTTPStatusCode(t, resp, http.StatusOK)
	_, _ = io.ReadAll(resp.Body)
}

// DeleteResource deletes a resource by ID (System-specific helper)
func (c *SystemTestClient) DeleteResource(t *testing.T, token, path string, id int64, isUser bool) {
	t.Helper()
	urlPath := path + "/" + strconv.FormatInt(id, 10)
	if isUser {
		urlPath += "?force=true"
	}
	resp := c.Delete(t, urlPath, token)
	defer resp.Body.Close()
	// Best-effort deletion for cleanup, ignore body read
}
