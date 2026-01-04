/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/features/auth/dto"
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
		// CORRECTED: Check for a "not found" error and return a specific, application-level error.
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		// For all other errors, return them as is.
		return nil, err
	}
	return dto.ConvertUserToUserPB(u), nil
}
