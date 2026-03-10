// Copyright 2024 OrigAdmin. All rights reserved.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
)

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
