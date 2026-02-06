/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package authorization

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/system/dto"
)

// mockAuthorizationRepo is a mock implementation of the AuthorizationRepo interface for testing.
type mockAuthorizationRepo struct {
	rolePerms []*types.RolePermission
	userRoles []*types.UserRole
	err       error
}

func (m *mockAuthorizationRepo) ListRolePermissions(ctx context.Context) ([]*types.RolePermission, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.rolePerms, nil
}

// ListPermissions is not used by AuthorizationUseCase, so it returns an empty slice.
func (m *mockAuthorizationRepo) ListPermissions(ctx context.Context) ([]*types.Permission, error) {
	return []*types.Permission{}, nil
}

func (m *mockAuthorizationRepo) ListUserRoles(ctx context.Context) ([]*types.UserRole, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.userRoles, nil
}

func TestAuthorizationUseCase_ListAllPolicies(t *testing.T) {
	logger := log.NewStdLogger(log.Writer(nil))
	ctx := context.Background()

	// --- Test Data ---
	mockRolePerms := []*types.RolePermission{
		{
			Id: 1,
			Role: &types.Role{
				Id:      1,
				Keyword: "admin",
			},
			Permission: &types.Permission{
				Id:      101,
				Keyword: "user:manage",
				Resources: []*types.Resource{
					{Path: "/api/users", Method: "GET"},
					{Path: "/api/users", Method: "POST"},
				},
			},
		},
		{
			Id: 2,
			Role: &types.Role{
				Id:      2,
				Keyword: "viewer",
			},
			Permission: &types.Permission{
				Id:      102,
				Keyword: "post:view",
				Resources: []*types.Resource{
					{Path: "/api/posts", Method: "GET"},
				},
			},
		},
		// Incomplete data to test resilience
		{Id: 3, Role: nil, Permission: &types.Permission{Id: 103}},
		{Id: 4, Role: &types.Role{Id: 3}, Permission: nil},
		{Id: 5, Role: &types.Role{Id: 4, Keyword: "no_resource_role"}, Permission: &types.Permission{Id: 104, Resources: []*types.Resource{}}},
	}

	mockUserRoles := []*types.UserRole{
		{
			Id:   1,
			User: &types.User{Id: 1},
			Role: &types.Role{Id: 1, Keyword: "admin"},
		},
		{
			Id:   2,
			User: &types.User{Id: 2},
			Role: &types.Role{Id: 2, Keyword: "viewer"},
		},
		// Incomplete data
		{Id: 3, User: nil, Role: &types.Role{Id: 1}},
		{Id: 4, User: &types.User{Id: 3}, Role: nil},
	}

	testCases := []struct {
		name                  string
		repo                  dto.AuthorizationRepo
		expectedAccessRules   []*systemv1.AccessRule
		expectedGroupingRules []*systemv1.GroupingRule
		expectError           bool
	}{
		{
			name: "Success - policies are converted correctly",
			repo: &mockAuthorizationRepo{
				rolePerms: mockRolePerms,
				userRoles: mockUserRoles,
			},
			expectedAccessRules: []*systemv1.AccessRule{
				{Subject: "admin", Object: "/api/users", Action: "GET"},
				{Subject: "admin", Object: "/api/users", Action: "POST"},
				{Subject: "viewer", Object: "/api/posts", Action: "GET"},
			},
			expectedGroupingRules: []*systemv1.GroupingRule{
				{User: "user:1", Group: "admin"},
				{User: "user:2", Group: "viewer"},
			},
			expectError: false,
		},
		{
			name: "Success - empty repository returns empty policies",
			repo: &mockAuthorizationRepo{
				rolePerms: []*types.RolePermission{},
				userRoles: []*types.UserRole{},
			},
			expectedAccessRules:   []*systemv1.AccessRule{},
			expectedGroupingRules: []*systemv1.GroupingRule{},
			expectError:           false,
		},
		{
			name: "Failure - repository returns an error",
			repo: &mockAuthorizationRepo{
				err: errors.New("database connection failed"),
			},
			expectedAccessRules:   nil,
			expectedGroupingRules: nil,
			expectError:           true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uc := NewAuthorizationUseCase(tc.repo, logger)

			resp, err := uc.ListAllPolicies(ctx, &systemv1.ListAllPoliciesRequest{})

			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				assert.ElementsMatch(t, tc.expectedAccessRules, resp.AccessRules, "AccessRules should match")
				assert.ElementsMatch(t, tc.expectedGroupingRules, resp.GroupingRules, "GroupingRules should match")
			}
		})
	}
}
