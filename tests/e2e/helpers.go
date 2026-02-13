// Copyright 2024 OrigAdmin. All rights reserved.

package e2e

import (
	"net/http"
	"testing"

	typesv1 "origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/tests/tools"
)

const (
	baseURL = "http://localhost:8000"
)

// SystemTestClient provides system-specific API testing functionality for e2e tests
type SystemTestClient struct {
	*tools.SystemTestClient
}

// NewSystemTestClient creates a new system test client for e2e tests
func NewSystemTestClient() *SystemTestClient {
	return &SystemTestClient{
		SystemTestClient: tools.NewSystemTestClient(baseURL),
	}
}

// NewSystemTestClientWithPrefix creates a new system test client with custom prefix
func NewSystemTestClientWithPrefix(prefix string) *SystemTestClient {
	return &SystemTestClient{
		SystemTestClient: tools.NewSystemTestClientWithPrefix(baseURL, prefix),
	}
}

// Global system test client instance for backward compatibility
var defaultSystemClient = NewSystemTestClientWithPrefix("/api/v1")

// LoginAndGetToken performs admin login and returns the access token (exported version)
func LoginAndGetToken(t *testing.T) string {
	return loginAndGetToken(t)
}

// Login performs login and returns the access token (exported version)
func Login(t *testing.T, username, password string) string {
	return login(t, username, password)
}

// DoRequest sends an HTTP request with JSON body (exported version)
func DoRequest(t *testing.T, method, path string, body interface{}, token string) *http.Response {
	return doRequest(t, method, path, body, token)
}

// DeleteResource deletes a resource by ID (exported version)
func DeleteResource(t *testing.T, token, path string, id int64, isUser bool) {
	deleteResource(t, token, path, id, isUser)
}

// loginAndGetToken performs admin login and returns the access token
func loginAndGetToken(t *testing.T) string {
	return login(t, "admin", "admin123")
}

// login performs login and returns the access token
func login(t *testing.T, username, password string) string {
	return defaultSystemClient.Login(t, username, password)
}

// doRequest sends an HTTP request with JSON body (backward compatibility)
func doRequest(t *testing.T, method, path string, body interface{}, token string) *http.Response {
	client := tools.NewTestHTTPClient(baseURL)
	client.SetPrefix("/api/v1")
	return client.Request(t, method, path, body, token)
}

// CreateUser creates a new user via API and returns its ID (exported version)
func CreateUser(t *testing.T, token, username, password string, roleIDs []int64) int64 {
	return createUser(t, token, username, password, roleIDs)
}

// UpdateRole updates an existing role via API (exported version)
func UpdateRole(t *testing.T, token string, roleID int64, name, keyword string, permissionIDs []int64) {
	updateRole(t, token, roleID, name, keyword, permissionIDs)
}

// UpdateUser updates an existing user via API (exported version)
func UpdateUser(t *testing.T, token string, userID int64, user interface{}, updateMaskPaths []string, roleIDs []int64) {
	updateUser(t, token, userID, user, updateMaskPaths, roleIDs)
}

// UpdateResource updates an existing resource via API (exported version)
func UpdateResource(t *testing.T, token string, resID int64, keyword, path, method, operation string) {
	updateResource(t, token, resID, keyword, path, method, operation)
}

// CreatePermission creates a new permission via API and returns its ID (exported version)
func CreatePermission(t *testing.T, token, name, keyword string, resourceIDs []int64) int64 {
	return createPermission(t, token, name, keyword, resourceIDs)
}

// CreateRole creates a new role via API and returns its ID (exported version)
func CreateRole(t *testing.T, token, name, keyword string, permissionIDs []int64) int64 {
	return createRole(t, token, name, keyword, permissionIDs)
}

// createRole creates a new role via API and returns its ID (backward compatibility)
func createRole(t *testing.T, token, name, keyword string, permissionIDs []int64) int64 {
	return defaultSystemClient.CreateRole(t, token, name, keyword, permissionIDs)
}

// updateRole updates an existing role via API (backward compatibility)
func updateRole(t *testing.T, token string, roleID int64, name, keyword string, permissionIDs []int64) {
	defaultSystemClient.UpdateRole(t, token, roleID, name, keyword, permissionIDs)
}

// createUser creates a new user via API and returns its ID (backward compatibility)
func createUser(t *testing.T, token, username, password string, roleIDs []int64) int64 {
	return defaultSystemClient.CreateUser(t, token, username, password, roleIDs)
}

// updateUser updates an existing user via API (backward compatibility)
func updateUser(t *testing.T, token string, userID int64, user interface{}, updateMaskPaths []string, roleIDs []int64) {
	// Convert user interface to *typesv1.User if needed
	if userObj, ok := user.(*typesv1.User); ok {
		defaultSystemClient.UpdateUser(t, token, userID, userObj, updateMaskPaths, roleIDs)
	}
}

// createPermission creates a new permission via API and returns its ID (backward compatibility)
func createPermission(t *testing.T, token, name, keyword string, resourceIDs []int64) int64 {
	return defaultSystemClient.CreatePermission(t, token, name, keyword, resourceIDs)
}

// updateResource updates an existing resource via API (backward compatibility)
func updateResource(t *testing.T, token string, resID int64, keyword, path, method, operation string) {
	defaultSystemClient.UpdateResource(t, token, resID, keyword, path, method, operation)
}

// deleteResource deletes a resource by ID (used for cleanup, backward compatibility)
func deleteResource(t *testing.T, token, path string, id int64, isUser bool) {
	defaultSystemClient.DeleteResource(t, token, path, id, isUser)
}
