/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/enums"
	identitydto "origadmin/application/admin/internal/features/identity/dto"
)

// AuthzRepo implements the dto.AuthzRepo interface.
type AuthzRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewAuthzRepo creates a new AuthzRepo.
func NewAuthzRepo(db *ent.Database, logger log.Logger) identitydto.AuthzRepo {
	return &AuthzRepo{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "dal.authz")),
	}
}

// ListMyPermissions retrieves all permissions for the current user.
func (r *AuthzRepo) ListMyPermissions(ctx context.Context, userID int64) (identitydto.PermissionsPB, error) {
	if userID != 0 && userID == data.GetSystemUserID() {
		permissions, err := r.db.Permission(ctx).Query().All(ctx)
		if err != nil {
			return nil, err
		}
		return identitydto.ConvertPermissionsToPermissionsPB(permissions), nil
	}
	permissions, err := r.db.User(ctx).Query().
		Where(user.ID(userID)).
		QueryRoles().
		QueryPermissions().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return identitydto.ConvertPermissionsToPermissionsPB(permissions), nil
}

// ListMyRoles retrieves all roles for the current user.
func (r *AuthzRepo) ListMyRoles(ctx context.Context, userID int64) (identitydto.RolesPB, error) {
	if userID != 0 && userID == data.GetSystemUserID() {
		roles, err := r.db.Role(ctx).Query().All(ctx)
		if err != nil {
			return nil, err
		}
		return identitydto.ConvertRolesToRolesPB(roles), nil
	}
	roles, err := r.db.User(ctx).Query().
		Where(user.ID(userID)).
		QueryRoles().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return identitydto.ConvertRolesToRolesPB(roles), nil
}

// ListMyViews retrieves all views for the current user.
func (r *AuthzRepo) ListMyViews(ctx context.Context, userID int64) (identitydto.ViewsPB, error) {
	if userID != 0 && userID == data.GetSystemUserID() {
		views, err := r.db.View(ctx).Query().All(ctx)
		if err != nil {
			return nil, err
		}
		return identitydto.ConvertViewsToViewsPB(views), nil
	}
	// Query: User -> Roles -> Permissions -> Views
	views, err := r.db.User(ctx).Query().
		Where(user.ID(userID)).
		QueryRoles().
		QueryPermissions().
		QueryViews().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return identitydto.ConvertViewsToViewsPB(views), nil
}

// ListMyResources retrieves all resources for the current user.
func (r *AuthzRepo) ListMyResources(ctx context.Context, userID int64) (identitydto.ResourcesPB, error) {
	if userID != 0 && userID == data.GetSystemUserID() {
		resources, err := r.db.Resource(ctx).Query().All(ctx)
		if err != nil {
			return nil, err
		}
		return identitydto.ConvertResourcesToResourcesPB(resources), nil
	}
	// Query: User -> Roles -> Permissions -> Resources
	resources, err := r.db.User(ctx).Query().
		Where(user.ID(userID)).
		QueryRoles().
		QueryPermissions().
		QueryResources().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return identitydto.ConvertResourcesToResourcesPB(resources), nil
}

// GetPermissionKeywordsByUserID retrieves all permission keywords for the current user.
func (r *AuthzRepo) GetPermissionKeywordsByUserID(ctx context.Context, userID int64) ([]string, error) {
	var permissions []*ent.Permission
	var err error
	if userID != 0 && userID == data.GetSystemUserID() {
		permissions, err = r.db.Permission(ctx).Query().All(ctx)
	} else {
		permissions, err = r.db.User(ctx).Query().
			Where(user.ID(userID)).
			QueryRoles().
			QueryPermissions().
			All(ctx)
	}
	if err != nil {
		return nil, err
	}
	keywords := make([]string, len(permissions))
	for i, p := range permissions {
		keywords[i] = p.Keyword
	}
	return keywords, nil
}

// HasSystemRole checks if the current user has a role with the 'system' type.
func (r *AuthzRepo) HasSystemRole(ctx context.Context, userID int64) (bool, error) {
	return r.db.User(ctx).
		Query().
		Where(user.ID(userID)).
		QueryRoles().
		Where(role.TypeEQ(enums.RoleTypeSystem)).
		Exist(ctx)
}
