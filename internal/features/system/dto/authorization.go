/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dto

import (
	"context"

	"origadmin/application/admin/api/v1/services/types"
)

// AuthorizationRepo defines the data access interface for authorization policies.
// It returns Protobuf types directly, as conversion is handled within the DAL.
type AuthorizationRepo interface {
	ListRolePermissions(ctx context.Context) ([]*types.RolePermission, error)
	ListRolePermissionsByRoleKeywords(ctx context.Context, roleKeywords ...string) ([]*types.RolePermission, error)
	ListPermissions(ctx context.Context) ([]*types.Permission, error)
	ListRolesByIDs(ctx context.Context, ids ...int64) ([]*types.Role, error)
	ListPermissionsByIDs(ctx context.Context, ids ...int64) ([]*types.Permission, error)
	ListUserRoles(ctx context.Context) ([]*types.UserRole, error)

	// ListAllPolicies retrieves all role permissions and user roles in a single, atomic operation.
	// This ensures data consistency for full policy synchronization.
	ListAllPolicies(ctx context.Context) ([]*types.RolePermission, []*types.UserRole, error)
}
