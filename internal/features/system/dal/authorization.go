/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/permission"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/data/entity/ent/rolepermission"
	"origadmin/application/admin/internal/features/system/dto"
)

// authorizationRepo implements the dto.AuthorizationRepo interface.
type authorizationRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewAuthorizationRepo creates a new authorization repository.
func NewAuthorizationRepo(database *ent.Database, logger log.Logger) dto.AuthorizationRepo {
	return &authorizationRepo{
		db:  database,
		log: log.NewHelper(log.With(logger, "module", "dal.authorization")),
	}
}

// ListRolePermissions queries the role_permissions join table, converts entities to PB types, and returns them.
func (r *authorizationRepo) ListRolePermissions(ctx context.Context) ([]*types.RolePermission, error) {
	// Query the explicit join table entity `RolePermission`.
	// We need to eager-load Role, and Permission with its Resources.
	rolePerms, err := r.db.RolePermission(ctx).Query().
		WithRole().
		WithPermission(func(q *ent.PermissionQuery) {
			q.WithResources()
		}).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query role_permissions: %v", err)
		return nil, err
	}

	r.log.WithContext(ctx).Infof("DAL: Loaded %d role_permissions with edges", len(rolePerms))

	// Manually convert the slice using the generated item converter.
	dtos := make([]*types.RolePermission, len(rolePerms))
	for i, rp := range rolePerms {
		dtos[i] = dto.ConvertRolePermissionToRolePermissionPB(rp)
		// Debug: check if edges are loaded
		if rp.Edges.Role == nil {
			r.log.WithContext(ctx).Warnf("DAL: RolePermission #%d has nil Role edge", rp.ID)
		}
		if rp.Edges.Permission == nil {
			r.log.WithContext(ctx).Warnf("DAL: RolePermission #%d has nil Permission edge", rp.ID)
		}
	}

	return dtos, nil
}

// ListRolePermissionsByRoleKeywords queries role-permission relations for a specific set of role keywords.
func (r *authorizationRepo) ListRolePermissionsByRoleKeywords(ctx context.Context, roleKeywords ...string) ([]*types.RolePermission, error) {
	r.log.WithContext(ctx).Infof("DAL: Querying role_permissions for keywords: %v", roleKeywords)
	rolePerms, err := r.db.RolePermission(ctx).Query().
		Where(rolepermission.HasRoleWith(role.KeywordIn(roleKeywords...))).
		WithRole().
		WithPermission(func(q *ent.PermissionQuery) {
			q.WithResources()
		}).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query role_permissions by role keywords: %v", err)
		return nil, err
	}

	r.log.WithContext(ctx).Infof("DAL: Found %d role_permission records for keywords: %v", len(rolePerms), roleKeywords)
	for i, rp := range rolePerms {
		r.log.WithContext(ctx).Infof("CONFIRM: DAL Result #%d: RoleID=%d, PermissionID=%d", i, rp.Edges.Role.ID, rp.Edges.Permission.ID)
	}

	dtos := make([]*types.RolePermission, len(rolePerms))
	for i, rp := range rolePerms {
		dtos[i] = dto.ConvertRolePermissionToRolePermissionPB(rp)
	}

	return dtos, nil
}

// ListPermissions queries all permissions, converts entities to PB types, and returns them.
func (r *authorizationRepo) ListPermissions(ctx context.Context) ([]*types.Permission, error) {
	// Eager-load resources as they contain the actual service/method info.
	permissions, err := r.db.Permission(ctx).Query().WithResources().All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query permissions: %v", err)
		return nil, err
	}

	// Use the generated slice converter.
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

// ListRolesByIDs queries roles by their IDs.
func (r *authorizationRepo) ListRolesByIDs(ctx context.Context, ids ...int64) ([]*types.Role, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	r.log.WithContext(ctx).Infof("DAL: Querying roles by IDs: %v", ids)
	roles, err := r.db.Role(ctx).Query().
		Where(role.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query roles by IDs %v: %v", ids, err)
		return nil, err
	}
	r.log.WithContext(ctx).Infof("DAL: Found %d roles by IDs %v", len(roles), ids)
	return dto.ConvertRolesToRolesPB(roles), nil
}

// ListPermissionsByIDs queries permissions by their IDs.
func (r *authorizationRepo) ListPermissionsByIDs(ctx context.Context, ids ...int64) ([]*types.Permission, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	permissions, err := r.db.Permission(ctx).Query().
		Where(permission.IDIn(ids...)).
		WithResources().
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query permissions by IDs: %v", err)
		return nil, err
	}
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

// ListUserRoles queries the user_roles join table, converts entities to PB types, and returns them.
func (r *authorizationRepo) ListUserRoles(ctx context.Context) ([]*types.UserRole, error) {
	// Query the explicit join table entity `UserRole`.
	userRoles, err := r.db.UserRole(ctx).Query().
		WithUser(). // Eager-load the associated User.
		WithRole(). // Eager-load the associated Role.
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query user_roles: %v", err)
		return nil, err
	}

	// Manually convert the slice using the generated item converter.
	dtos := make([]*types.UserRole, len(userRoles))
	for i, ur := range userRoles {
		dtos[i] = dto.ConvertUserRoleToUserRolePB(ur)
	}

	return dtos, nil
}
