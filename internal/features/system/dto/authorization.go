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
	ListPermissions(ctx context.Context) ([]*types.Permission, error)
	ListUserRoles(ctx context.Context) ([]*types.UserRole, error)
}
