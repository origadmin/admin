/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dto

import (
	"context"

	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
)

// AuthorizationRepo defines the data access interface for authorization policies.
// It provides methods to read policies from casbin_rule table and convert them to PolicySpec format.
type AuthorizationRepo interface {
	// ListPolicies queries policies from casbin_rule table with filter criteria.
	// Returns policies in authzv1.PolicySpec format converted from casbin_rule.
	ListPolicies(ctx context.Context, req *system.ListPoliciesRequest) ([]*authzv1.PolicySpec, int32, error)

	ListRolePermissions(ctx context.Context) ([]*types.RolePermission, error)
	ListRolePermissionsByRoleKeywords(ctx context.Context, roleKeywords ...string) ([]*types.RolePermission, error)
	ListPermissions(ctx context.Context) ([]*types.Permission, error)
	ListRolesByIDs(ctx context.Context, ids ...int64) ([]*types.Role, error)
	ListPermissionsByIDs(ctx context.Context, ids ...int64) ([]*types.Permission, error)
	ListUserRoles(ctx context.Context) ([]*types.UserRole, error)
	ListUserRolePermissions(ctx context.Context) ([]*types.RolePermission, []*types.UserRole, error)
}
