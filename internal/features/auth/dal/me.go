/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/entity/ent/view"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/auth/dto"
	systemDto "origadmin/application/admin/internal/features/system/dto"
)

type MeRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewMeRepo .
func NewMeRepo(db *ent.Database, logger log.Logger) dto.MeRepo {
	return &MeRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

func (r *MeRepo) GetProfile(ctx context.Context, userID int64) (*types.User, error) {
	u, err := r.db.User(ctx).Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return systemDto.ConvertUserToUserPB(u), nil
}

// ListActiveViews retrieves all active views, ordered by sequence.
func (r *MeRepo) ListActiveViews(ctx context.Context) ([]*types.View, error) {
	views, err := r.db.View(ctx).Query().
		Where(view.StatusEQ(enums.StatusActive)).
		Order(ent.Asc(view.FieldSequence)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return systemDto.ConvertViewsToViewsPB(views), nil
}

// GetPermissionKeywordsByUserID retrieves all permission keywords for a user.
func (r *MeRepo) GetPermissionKeywordsByUserID(ctx context.Context, userID string) ([]string, error) {
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, err
	}

	permissions, err := r.db.User(ctx).Query().
		Where(user.ID(id)).
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
func (r *MeRepo) HasSystemRole(ctx context.Context, userID int64) (bool, error) {
	return r.db.User(ctx).
		Query().
		Where(user.ID(userID)).
		QueryRoles().
		Where(role.TypeEQ(enums.RoleTypeSystem)).
		Exist(ctx)
}
