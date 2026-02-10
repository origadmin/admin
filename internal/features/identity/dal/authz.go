/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/identity/dto"
)

// AuthzRepo implements the dto.AuthzRepo interface.
type AuthzRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewAuthzRepo creates a new AuthzRepo.
func NewAuthzRepo(db *ent.Database, logger log.Logger) dto.AuthzRepo {
	return &AuthzRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

// GetPermissionKeywordsByUserID retrieves all permission keywords for a user.
func (r *AuthzRepo) GetPermissionKeywordsByUserID(ctx context.Context, userID int64) ([]string, error) {
	permissions, err := r.db.User(ctx).Query().
		Where(user.ID(userID)).
		QueryRoles().
		QueryPermissions().
		All(ctx)
	if err != nil {
		return nil, err
	}
	keywords := make([]string, len(permissions))
	for i, p := range permissions {
		keywords[i] = p.Keyword
	}
	return keywords, nil
}

// HasSystemRole checks if the user has a role with the 'system' type.
func (r *AuthzRepo) HasSystemRole(ctx context.Context, userID int64) (bool, error) {
	return r.db.User(ctx).
		Query().
		Where(user.ID(userID)).
		QueryRoles().
		Where(role.TypeEQ(enums.RoleTypeSystem)).
		Exist(ctx)
}
