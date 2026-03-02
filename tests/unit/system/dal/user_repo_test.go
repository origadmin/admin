/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/runtime/log"
	_ "github.com/sqlite3ent/sqlite3"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/dal"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/tests/tools"
)

// setupUserTest creates a test database and initializes repos
func setupUserTest(t *testing.T) (context.Context, *ent.Database, func()) {
	t.Helper()
	ctx := context.Background()

	// Use tools.SetupTestDatabase for consistent test setup
	client := tools.SetupTestDatabase(t)
	db := ent.NewDatabaseWithClient(client)

	cleanup := func() {
		client.Close()
	}

	return ctx, db, cleanup
}

// TestUserSoftDeleteAndRestore tests soft delete and restore functionality for User entity
func TestUserSoftDeleteAndRestore(t *testing.T) {
	ctx, db, cleanup := setupUserTest(t)
	defer cleanup()

	// Initialize repos
	userRepo := dal.NewUserRepo(db, log.DefaultLogger)
	roleRepo := dal.NewRoleRepo(db, log.DefaultLogger)

	// Step 1: Create test role
	testRole, err := roleRepo.Create(ctx, &types.Role{
		Name:    "Test Role",
		Keyword: "test_role_" + time.Now().Format("20060102150405"),
		Status:  int32(enums.StatusActive),
	})
	require.NoError(t, err)
	require.NotNil(t, testRole)
	t.Logf("�?Created test role: ID=%d, Name=%s", testRole.Id, testRole.Name)

	// Step 2: Create a test user with role
	testUser, err := userRepo.Create(ctx, &types.User{
		Username: "testuser_" + time.Now().Format("20060102150405"),
		Nickname: "Test User",
		Status:   int32(enums.StatusActive),
	}, "")
	require.NoError(t, err)
	require.NotNil(t, testUser)
	t.Logf("�?Created test user: ID=%d, Username=%s", testUser.Id, testUser.Username)

	// Add role to user
	_, err = userRepo.AddRoleIDs(ctx, testUser.Id, []int64{testRole.Id})
	require.NoError(t, err)
	require.NotNil(t, testUser)
	t.Logf("�?Created test user: ID=%d, Username=%s", testUser.Id, testUser.Username)

	// Step 3: Verify user has role association
	userWithRoles, err := userRepo.Get(ctx, testUser.Id, &dto.UserQueryOption{WithRoles: true})
	require.NoError(t, err)
	require.NotNil(t, userWithRoles)
	require.Len(t, userWithRoles.Roles, 1, "User should have 1 role before deletion")
	t.Logf("�?User has %d role(s) before deletion", len(userWithRoles.Roles))

	// Step 4: Soft delete the user
	err = userRepo.Delete(ctx, testUser.Id)
	require.NoError(t, err)
	t.Logf("�?Soft deleted user: ID=%d", testUser.Id)

	// Step 5: Verify user is soft deleted (delete_time is set)
	// Note: User should not be found in normal query after soft delete
	_, err = userRepo.Get(ctx, testUser.Id, &dto.UserQueryOption{WithRoles: true})
	assert.Error(t, err, "Should get error when querying soft-deleted user normally")
	t.Logf("�?User is not found in normal query (as expected for soft-deleted records)")

	// Step 6: Restore the user
	err = userRepo.Restore(ctx, testUser.Id)
	require.NoError(t, err)
	t.Logf("�?Restored user: ID=%d", testUser.Id)

	// Step 7: Verify user is restored and still has role association
	restoredUser, err := userRepo.Get(ctx, testUser.Id, &dto.UserQueryOption{WithRoles: true})
	require.NoError(t, err)
	require.NotNil(t, restoredUser)
	require.Len(t, restoredUser.Roles, 1, "User should have 1 role after restoration")
	require.Equal(t, testRole.Id, restoredUser.Roles[0].Id, "Role association should be preserved")
	t.Logf("�?User restored with %d role(s) intact", len(restoredUser.Roles))
}

// TestUserSoftDeleteDoesNotClearAssociations verifies that soft-delete does not clear associations
func TestUserSoftDeleteDoesNotClearAssociations(t *testing.T) {
	ctx, db, cleanup := setupUserTest(t)
	defer cleanup()

	userRepo := dal.NewUserRepo(db, log.DefaultLogger)
	roleRepo := dal.NewRoleRepo(db, log.DefaultLogger)

	// Create two roles
	role1, err := roleRepo.Create(ctx, &types.Role{
		Name:    "Role 1",
		Keyword: "role_1_" + time.Now().Format("20060102150405"),
		Status:  int32(enums.StatusActive),
	})
	require.NoError(t, err)

	role2, err := roleRepo.Create(ctx, &types.Role{
		Name:    "Role 2",
		Keyword: "role_2_" + time.Now().Format("20060102150405"),
		Status:  int32(enums.StatusActive),
	})
	require.NoError(t, err)

	// Create user with multiple roles
	user, err := userRepo.Create(ctx, &types.User{
		Username: "user_multi_" + time.Now().Format("20060102150405"),
		Nickname: "Multi Role User",
		Status:   int32(enums.StatusActive),
	}, "")
	require.NoError(t, err)

	// Add both roles
	_, err = userRepo.AddRoleIDs(ctx, user.Id, []int64{role1.Id, role2.Id})
	require.NoError(t, err)

	// Verify user has 2 roles
	userWithRoles, err := userRepo.Get(ctx, user.Id, &dto.UserQueryOption{WithRoles: true})
	require.NoError(t, err)
	require.Len(t, userWithRoles.Roles, 2)
	t.Logf("�?User initially has %d roles", len(userWithRoles.Roles))

	// Soft delete user
	err = userRepo.Delete(ctx, user.Id)
	require.NoError(t, err)
	t.Logf("�?User soft deleted")

	// Restore user
	err = userRepo.Restore(ctx, user.Id)
	require.NoError(t, err)
	t.Logf("�?User restored")

	// Verify user still has both roles
	restoredUser, err := userRepo.Get(ctx, user.Id, &dto.UserQueryOption{WithRoles: true})
	require.NoError(t, err)
	require.Len(t, restoredUser.Roles, 2, "User should still have 2 roles after restoration")
	t.Logf("�?User still has %d roles after deletion and restoration", len(restoredUser.Roles))

	// Verify the roles are the same ones
	roleIDs := make(map[int64]bool)
	for _, r := range restoredUser.Roles {
		roleIDs[r.Id] = true
	}
	assert.True(t, roleIDs[role1.Id], "Role1 should be preserved")
	assert.True(t, roleIDs[role2.Id], "Role2 should be preserved")
	t.Logf("�?All role associations preserved correctly")
}
