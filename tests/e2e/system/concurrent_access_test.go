// Copyright 2024 OrigAdmin. All rights reserved.

package system

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	systemv1 "origadmin/application/admin/api/v1/services/system"
	typesv1 "origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/tests/e2e"
)

func TestConcurrentPermissionAccess(t *testing.T) {
	// Prepare the test environment
	adminToken := e2e.LoginAndGetToken(t)
	require.NotEmpty(t, adminToken, "Admin login failed")

	var resUserListID, resUserCreateID int64

	// Find Core Resources (APIs)
	t.Run("FindResources", func(t *testing.T) {
		resp := e2e.DoRequest(t, "GET", "/api/v1/sys/resources?page_size=1000", nil, adminToken)
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)

		var listResp systemv1.ListResourcesResponse
		bodyBytes, _ := io.ReadAll(resp.Body)
		err := protojson.Unmarshal(bodyBytes, &listResp)
		require.NoError(t, err)

		for _, res := range listResp.Resources {
			switch res.Keyword {
			case "system:user:list_users":
				resUserListID = res.Id
			case "system:user:create_user":
				resUserCreateID = res.Id
			}
		}
		require.NotZero(t, resUserListID, "Resource 'system:user:list_users' should be found")
		require.NotZero(t, resUserCreateID, "Resource 'system:user:create_user' should be found")
		t.Logf("Found resource IDs: UserList=%d, UserCreate=%d", resUserListID, resUserCreateID)
	})

	// Create a test role
	uniqueSuffix := fmt.Sprintf("%d", time.Now().Unix())

	// Create permissions first
	permUserListID := e2e.CreatePermission(t, adminToken, "Perm_UserList_"+uniqueSuffix, "system:user:list_users:"+uniqueSuffix, []int64{resUserListID})
	permUserCreateID := e2e.CreatePermission(t, adminToken, "Perm_UserCreate_"+uniqueSuffix, "system:user:create_user:"+uniqueSuffix, []int64{resUserCreateID})
	defer e2e.DeleteResource(t, adminToken, "/api/v1/sys/permissions", permUserListID, false)
	defer e2e.DeleteResource(t, adminToken, "/api/v1/sys/permissions", permUserCreateID, false)

	// Create role with permissions
	roleName := "ConcurrentRole_" + uniqueSuffix
	roleID := e2e.CreateRole(t, adminToken, roleName, "cr_"+uniqueSuffix, []int64{permUserListID})
	defer e2e.DeleteResource(t, adminToken, "/api/v1/sys/roles", roleID, false)

	// Create multiple test users
	userCount := 10
	userTokens := make([]string, userCount)
	userIDs := make([]int64, userCount)

	for i := 0; i < userCount; i++ {
		username := fmt.Sprintf("cu_%s_%d", uniqueSuffix, i)
		userIDs[i] = e2e.CreateUser(t, adminToken, username, "password123", []int64{roleID})
		defer e2e.DeleteResource(t, adminToken, "/api/v1/sys/users", userIDs[i], true)

		// Log in to get the token
		userTokens[i] = e2e.Login(t, username, "password123")
	}

	// Concurrency testing
	t.Run("ConcurrentReadAccess", func(t *testing.T) {
		// Wait for policy synchronization for ALL users
		t.Log("Waiting for policy propagation for all users...")
		for i, token := range userTokens {
			require.Eventually(t, func() bool {
				resp := e2e.DoRequest(t, "GET", "/api/v1/sys/users", nil, token)
				defer resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
					t.Logf("User %d permissions synced", i)
					return true
				}
				return false
			}, 10*time.Second, 1*time.Second, "Policy synchronization timeout for user %d", i)
		}

		var wg sync.WaitGroup
		errors := make(chan error, userCount)
		successCount := make(chan int, userCount)

		for i, token := range userTokens {
			wg.Add(1)
			go func(idx int, userToken string) {
				defer wg.Done()

				// Concurrent access
				resp := e2e.DoRequest(t, "GET", "/api/v1/sys/users", nil, userToken)
				defer resp.Body.Close()

				if resp.StatusCode == http.StatusOK {
					successCount <- 1
				} else {
					errors <- fmt.Errorf("user %d got status %d", idx, resp.StatusCode)
				}
			}(i, token)
		}

		wg.Wait()
		close(successCount)
		close(errors)

		// Validate the results
		totalSuccess := 0
		for range successCount {
			totalSuccess++
		}

		for err := range errors {
			t.Errorf("Concurrent test failed: %v", err)
		}

		assert.Equal(t, userCount, totalSuccess, "All users should have read permission")
	})

	t.Run("ConcurrentWriteAccess", func(t *testing.T) {
		// Create write permissions
		writePermUserListID := e2e.CreatePermission(t, adminToken, "Write_Perm_UserList_"+uniqueSuffix, "system:user:list_users_write:"+uniqueSuffix, []int64{resUserListID})
		writePermUserCreateID := e2e.CreatePermission(t, adminToken, "Write_Perm_UserCreate_"+uniqueSuffix, "system:user:create_user_write:"+uniqueSuffix, []int64{resUserCreateID})
		defer e2e.DeleteResource(t, adminToken, "/api/v1/sys/permissions", writePermUserListID, false)
		defer e2e.DeleteResource(t, adminToken, "/api/v1/sys/permissions", writePermUserCreateID, false)

		// Create a role with write permissions
		writeRoleName := "ConcurrentWriteRole_" + uniqueSuffix
		writeRoleID := e2e.CreateRole(t, adminToken, writeRoleName, "cwr_"+uniqueSuffix,
			[]int64{writePermUserListID, writePermUserCreateID})
		defer e2e.DeleteResource(t, adminToken, "/api/v1/sys/roles", writeRoleID, false)

		// Create a user with write access
		writeUsername := "cu_write_" + uniqueSuffix
		writeUserID := e2e.CreateUser(t, adminToken, writeUsername, "password123", []int64{writeRoleID})
		defer e2e.DeleteResource(t, adminToken, "/api/v1/sys/users", writeUserID, true)

		writeToken := e2e.Login(t, writeUsername, "password123")

		// Wait for write permissions to sync for the write user
		t.Log("Waiting for write permission propagation...")
		require.Eventually(t, func() bool {
			resp := e2e.DoRequest(t, "GET", "/api/v1/sys/users", nil, writeToken)
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				t.Log("Write user permissions synced")
				return true
			}
			return false
		}, 15*time.Second, 1*time.Second, "Write permission synchronization timeout")

		// Concurrent user creation
		var wg sync.WaitGroup
		errors := make(chan error, 5)

		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()

				username := fmt.Sprintf("concurrent_%s_%d", uniqueSuffix, idx)
				userPayload := &typesv1.User{
					Username: username,
					Nickname: username,
				}
				req := &systemv1.CreateUserRequest{
					User:     userPayload,
					Password: "password123",
				}

				resp := e2e.DoRequest(t, "POST", "/api/v1/sys/users", req, writeToken)
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					errors <- fmt.Errorf("create user %d failed with status %d", idx, resp.StatusCode)
				}
			}(i)
		}

		wg.Wait()
		close(errors)

		// Validate the results
		errCount := 0
		for range errors {
			errCount++
		}
		assert.Equal(t, 0, errCount, "All concurrent create operations should succeed")
	})
}
