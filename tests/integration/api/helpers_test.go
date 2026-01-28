// Copyright 2024 OrigAdmin. All rights reserved.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
	"origadmin/application/admin/internal/features/system"
	"origadmin/application/admin/internal/features/system/server"
)

// TestComponents holds all the necessary components for running integration tests.
type TestComponents struct {
	Ctx        context.Context
	DBClient   *ent.Client
	HTTPServer *httptest.Server
	Router     *gin.Engine
}

// setupTestServer initializes a test server with an in-memory database and returns the components.
func setupTestServer(t *testing.T) *TestComponents {
	t.Helper()

	// Setup in-memory SQLite database
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	t.Cleanup(func() { client.Close() })

	// Initialize router and server
	gin.SetMode(gin.TestMode)
	router := gin.New()
	httpServer := httptest.NewServer(router)
	t.Cleanup(httpServer.Close)

	// Register system routes
	system.RegisterRoutes(router, client)

	return &TestComponents{
		Ctx:        context.Background(),
		DBClient:   client,
		HTTPServer: httpServer,
		Router:     router,
	}
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
