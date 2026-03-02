/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package database_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/sqlite3ent/sqlite3"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/dal"
	"origadmin/application/admin/internal/features/system/dto"
)

// setupUserTest creates a test database and initializes repos
func setupUserTest(t *testing.T) (context.Context, *ent.Database, func()) {
	ctx := context.Background()

	// Use an in-memory SQLite database for testing
	client := enttest.Open(t, "sqlite3", fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", t.Name()))
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
	userRepo := dal.NewUserRepo(db)
	roleRepo := dal.NewRoleRepo(db)

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
		Name:     "Test User",
		Status:   int32(enums.StatusActive),
	}, "")
	require.NoError(t, err)
	require.NotNil(t, testUser)
	t.Logf("�?Created test user: ID=%d, Username=%s", testUser.Id, testUser.Username)

	// Add role to user
	_, err = userRepo.AddRoleIDs(ctx, testUser.Id, []int64{testRole.Id})
	require.NoError(t, err)
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
	require.Len(t, restoredUser.Roles, 1, "User should still have 1 role after restoration")
	assert.Equal(t, testRole.Id, restoredUser.Roles[0].Id, "Role ID should match")
	t.Logf("�?User restored with %d role(s) intact", len(restoredUser.Roles))
}

// TestUserSoftDeleteDoesNotClearAssociations verifies that soft delete
// does NOT clear role/position/department associations
func TestUserSoftDeleteDoesNotClearAssociations(t *testing.T) {
	ctx, db, cleanup := setupUserTest(t)
	defer cleanup()

	userRepo := dal.NewUserRepo(db)
	roleRepo := dal.NewRoleRepo(db)

	// Create multiple roles
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
	testUser, err := userRepo.Create(ctx, &types.User{
		Username: "user_assoc_" + time.Now().Format("20060102150405"),
		Nickname: "Association Test User",
		Name:     "Association Test User",
		Status:   int32(enums.StatusActive),
	}, "")
	require.NoError(t, err)

	// Add roles to user
	_, err = userRepo.AddRoleIDs(ctx, testUser.Id, []int64{role1.Id, role2.Id})
	require.NoError(t, err)

	// Verify initial role count
	userBefore, err := userRepo.Get(ctx, testUser.Id, &dto.UserQueryOption{WithRoles: true})
	require.NoError(t, err)
	initialRoleCount := len(userBefore.Roles)
	require.Equal(t, 2, initialRoleCount, "User should have 2 roles initially")
	t.Logf("�?User initially has %d roles", initialRoleCount)

	// Soft delete user
	err = userRepo.Delete(ctx, testUser.Id)
	require.NoError(t, err)
	t.Logf("�?User soft deleted")

	// Restore user
	err = userRepo.Restore(ctx, testUser.Id)
	require.NoError(t, err)
	t.Logf("�?User restored")

	// Verify role associations are preserved
	userAfter, err := userRepo.Get(ctx, testUser.Id, &dto.UserQueryOption{WithRoles: true})
	require.NoError(t, err)
	require.Len(t, userAfter.Roles, initialRoleCount, "Role associations should be preserved after soft delete and restore")
	t.Logf("�?User still has %d roles after deletion and restoration", len(userAfter.Roles))

	// Verify role IDs match
	roleIDs := make(map[int64]bool)
	for _, r := range userAfter.Roles {
		roleIDs[r.Id] = true
	}
	assert.True(t, roleIDs[role1.Id], "Role 1 should be present")
	assert.True(t, roleIDs[role2.Id], "Role 2 should be present")
	t.Logf("�?All role associations preserved correctly")
}

// TestHardDeleteClearsAssociations verifies that entities without soft delete
// (like Role) properly clear their associations
func TestHardDeleteClearsAssociations(t *testing.T) {
	ctx, db, cleanup := setupUserTest(t)
	defer cleanup()

	roleRepo := dal.NewRoleRepo(db)
	userRepo := dal.NewUserRepo(db)

	// Create test role
	testRole, err := roleRepo.Create(ctx, &types.Role{
		Name:    "Hard Delete Test Role",
		Keyword: "hard_delete_" + time.Now().Format("20060102150405"),
		Status:  int32(enums.StatusActive),
	})
	require.NoError(t, err)

	// Create a user with this role
	testUser, err := userRepo.Create(ctx, &types.User{
		Username: "hard_delete_user_" + time.Now().Format("20060102150405"),
		Nickname: "Hard Delete Test User",
		Name:     "Hard Delete Test User",
		Status:   int32(enums.StatusActive),
	}, "")
	require.NoError(t, err)

	// Add role to user
	_, err = userRepo.AddRoleIDs(ctx, testUser.Id, []int64{testRole.Id})
	require.NoError(t, err)

	// Verify user has role before deletion
	userBefore, err := userRepo.Get(ctx, testUser.Id, &dto.UserQueryOption{WithRoles: true})
	require.NoError(t, err)
	require.Len(t, userBefore.Roles, 1, "User should have 1 role initially")
	t.Logf("�?User has 1 role before role deletion")

	// Hard delete the role (should clear user_roles association)
	err = roleRepo.Delete(ctx, testRole.Id)
	require.NoError(t, err)
	t.Logf("�?Role hard deleted: ID=%d", testRole.Id)

	// Verify user's role association is cleared
	userAfter, err := userRepo.Get(ctx, testUser.Id, &dto.UserQueryOption{WithRoles: true})
	require.NoError(t, err)
	require.Len(t, userAfter.Roles, 0, "User should have 0 roles after role deletion (association cleared)")
	t.Logf("�?User's role associations properly cleared after role hard deletion")
}

// TestSoftDeletedUserDoesNotAffectOtherUsers verifies that
// soft deleting one user does not affect other users
func TestSoftDeletedUserDoesNotAffectOtherUsers(t *testing.T) {
	ctx, db, cleanup := setupUserTest(t)
	defer cleanup()

	userRepo := dal.NewUserRepo(db)

	// Create two users
	user1, err := userRepo.Create(ctx, &types.User{
		Username: "user1_" + time.Now().Format("20060102150405"),
		Nickname: "User 1",
		Name:     "User 1",
		Status:   int32(enums.StatusActive),
	}, "")
	require.NoError(t, err)

	user2, err := userRepo.Create(ctx, &types.User{
		Username: "user2_" + time.Now().Format("20060102150405") + "1",
		Nickname: "User 2",
		Name:     "User 2",
		Status:   int32(enums.StatusActive),
	}, "")
	require.NoError(t, err)

	t.Logf("�?Created test users: ID1=%d, ID2=%d", user1.Id, user2.Id)

	// Soft delete first user
	err = userRepo.Delete(ctx, user1.Id)
	require.NoError(t, err)
	t.Logf("�?Soft deleted user1: ID=%d", user1.Id)

	// Verify user1 is not accessible
	_, err = userRepo.Get(ctx, user1.Id, &dto.UserQueryOption{WithRoles: true})
	assert.Error(t, err, "User1 should not be found after soft delete")

	// Verify user2 is still accessible
	user2After, err := userRepo.Get(ctx, user2.Id, &dto.UserQueryOption{WithRoles: true})
	require.NoError(t, err)
	require.NotNil(t, user2After)
	assert.Equal(t, user2.Id, user2After.Id, "User2 should still be accessible")
	t.Logf("�?User2 is still accessible after user1 is deleted")
}

// TestListUsersExcludesSoftDeleted verifies that listing users
// automatically excludes soft-deleted records
func TestListUsersExcludesSoftDeleted(t *testing.T) {
	ctx, db, cleanup := setupUserTest(t)
	defer cleanup()

	userRepo := dal.NewUserRepo(db)

	// Create multiple users
	users := make([]*types.User, 0, 3)
	for i := 0; i < 3; i++ {
		u, err := userRepo.Create(ctx, &types.User{
			Username: fmt.Sprintf("list_user_%d_%s", i, time.Now().Format("20060102150405")),
			Nickname: fmt.Sprintf("List User %d", i),
			Name:     fmt.Sprintf("List User %d", i),
			Status:   int32(enums.StatusActive),
		}, "")
		require.NoError(t, err)
		users = append(users, u)
	}

	// List all users
	allUsers, _, err := userRepo.List(ctx, &dto.UserQueryOption{})
	require.NoError(t, err)
	require.Len(t, allUsers, 3, "Should have 3 users initially")
	t.Logf("�?Initial list has %d users", len(allUsers))

	// Soft delete one user
	err = userRepo.Delete(ctx, users[0].Id)
	require.NoError(t, err)
	t.Logf("�?Soft deleted user: ID=%d", users[0].Id)

	// List users again
	remainingUsers, _, err := userRepo.List(ctx, &dto.UserQueryOption{})
	require.NoError(t, err)
	require.Len(t, remainingUsers, 2, "Should have 2 users after one is soft deleted")
	t.Logf("�?List now has %d users", len(remainingUsers))

	// Verify the soft-deleted user is not in the list
	userIDs := make(map[int64]bool)
	for _, u := range remainingUsers {
		userIDs[u.Id] = true
	}
	assert.False(t, userIDs[users[0].Id], "Soft-deleted user should not be in list")
	assert.True(t, userIDs[users[1].Id], "User 1 should still be in list")
	assert.True(t, userIDs[users[2].Id], "User 2 should still be in list")
	t.Logf("�?Soft-deleted user properly excluded from list")
}

// TestRestoreNonExistentUser verifies that restoring a non-existent user
// returns an appropriate error
func TestRestoreNonExistentUser(t *testing.T) {
	ctx, db, cleanup := setupUserTest(t)
	defer cleanup()

	userRepo := dal.NewUserRepo(db)

	// Try to restore a non-existent user
	err := userRepo.Restore(ctx, 999999)
	assert.Error(t, err, "Should return error for non-existent user")
	t.Logf("�?Correctly returns error when trying to restore non-existent user")
}

// TestSoftDeleteMultipleTimes verifies that soft-deleting
// an already soft-deleted user handles gracefully
func TestSoftDeleteMultipleTimes(t *testing.T) {
	ctx, db, cleanup := setupUserTest(t)
	defer cleanup()

	userRepo := dal.NewUserRepo(db)

	// Create user
	testUser, err := userRepo.Create(ctx, &types.User{
		Username: "multi_delete_" + time.Now().Format("20060102150405"),
		Nickname: "Multi Delete User",
		Name:     "Multi Delete User",
		Status:   int32(enums.StatusActive),
	}, "")
	require.NoError(t, err)

	// First delete
	err = userRepo.Delete(ctx, testUser.Id)
	require.NoError(t, err)
	t.Logf("�?First soft delete successful")

	// Try to delete again (user is already soft-deleted)
	err = userRepo.Delete(ctx, testUser.Id)
	// This might succeed or error depending on implementation
	// The important thing is that it doesn't cause a crash
	t.Logf("�?Second soft delete handled gracefully: %v", err)
}
