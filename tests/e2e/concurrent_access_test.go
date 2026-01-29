// Copyright 2024 OrigAdmin. All rights reserved.

package e2e

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
)

func TestConcurrentPermissionAccess(t *testing.T) {
	// 准备测试环境
	adminToken := loginAndGetToken(t)
	require.NotEmpty(t, adminToken, "Admin login failed")

	var resUserListID, resUserCreateID int64

	// Find Core Resources (APIs)
	t.Run("FindResources", func(t *testing.T) {
		resp := doRequest(t, "GET", "/api/v1/sys/resources?page_size=1000", nil, adminToken)
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

	// 创建测试角色
	uniqueSuffix := fmt.Sprintf("%d", time.Now().Unix())
	roleName := "ConcurrentRole_" + uniqueSuffix
	roleID := createRole(t, adminToken, roleName, "cr_"+uniqueSuffix, []int64{resUserListID})
	defer deleteResource(t, adminToken, "/api/v1/sys/roles", roleID, false)

	// 创建多个测试用户
	userCount := 10
	userTokens := make([]string, userCount)
	userIDs := make([]int64, userCount)

	for i := 0; i < userCount; i++ {
		username := fmt.Sprintf("cu_%s_%d", uniqueSuffix, i)
		userIDs[i] = createUser(t, adminToken, username, "password123", []int64{roleID})
		defer deleteResource(t, adminToken, "/api/v1/sys/users", userIDs[i], true)

		// 登录获取token
		userTokens[i] = login(t, username, "password123")
	}

	// 并发测试
	t.Run("ConcurrentReadAccess", func(t *testing.T) {
		var wg sync.WaitGroup
		errors := make(chan error, userCount)
		successCount := make(chan int, userCount)

		for i, token := range userTokens {
			wg.Add(1)
			go func(idx int, userToken string) {
				defer wg.Done()

				// 并发访问
				resp := doRequest(t, "GET", "/api/v1/sys/users", nil, userToken)
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

		// 验证结果
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
		// 创建有写权限的角色
		writeRoleName := "ConcurrentWriteRole_" + uniqueSuffix
		writeRoleID := createRole(t, adminToken, writeRoleName, "cwr_"+uniqueSuffix,
			[]int64{resUserListID, resUserCreateID})
		defer deleteResource(t, adminToken, "/api/v1/sys/roles", writeRoleID, false)

		// 创建有写权限的用户
		writeUsername := "cu_write_" + uniqueSuffix
		writeUserID := createUser(t, adminToken, writeUsername, "password123", []int64{writeRoleID})
		defer deleteResource(t, adminToken, "/api/v1/sys/users", writeUserID, true)

		writeToken := login(t, writeUsername, "password123")

		// 并发创建用户
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

				resp := doRequest(t, "POST", "/api/v1/sys/users", req, writeToken)
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					errors <- fmt.Errorf("create user %d failed with status %d", idx, resp.StatusCode)
				}
			}(i)
		}

		wg.Wait()
		close(errors)

		// 验证结果
		errCount := 0
		for range errors {
			errCount++
		}
		assert.Equal(t, 0, errCount, "All concurrent create operations should succeed")
	})
}
