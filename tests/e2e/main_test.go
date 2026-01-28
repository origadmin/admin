package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	systemv1 "origadmin/application/admin/api/v1/services/system"
	typesv1 "origadmin/application/admin/api/v1/services/types"
)

// TestMain is the entry point for E2E tests.
func TestMain(m *testing.M) {
	// Here you can add global setup and teardown logic,
	// such as starting a server or connecting to a database.
	exitCode := m.Run()
	os.Exit(exitCode)
}

// --- Resource Management Helpers ---

func createRole(t *testing.T, token, name, keyword string, resourceIDs []int64) int64 {
	t.Helper()
	reqBody := map[string]interface{}{
		"role": map[string]interface{}{
			"name":    name,
			"keyword": keyword,
			"status":  1,
		},
		"resource_ids": resourceIDs,
	}

	resp := doRequest(t, "POST", "/api/v1/sys/roles", reqBody, token)
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode, "Failed to create role: %s", string(bodyBytes))

	var resMap map[string]interface{}
	err := json.Unmarshal(bodyBytes, &resMap)
	require.NoError(t, err)

	roleMap, ok := resMap["role"].(map[string]interface{})
	require.True(t, ok)
	idVal, ok := roleMap["id"]
	require.True(t, ok)

	var id int64
	switch v := idVal.(type) {
	case float64:
		id = int64(v)
	case string:
		_, err := fmt.Sscanf(v, "%d", &id)
		require.NoError(t, err)
	}
	return id
}

func createUser(t *testing.T, token, username, password string, roleIDs []int64) int64 {
	t.Helper()
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

func deleteResource(t *testing.T, token, path string, id int64, force bool) {
	t.Helper()
	if id == 0 {
		return
	}
	fullPath := fmt.Sprintf("%s/%d", path, id)
	if force {
		fullPath += "?force=true"
	}
	resp := doRequest(t, "DELETE", fullPath, nil, token)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Failed to delete resource at %s", fullPath)
}
