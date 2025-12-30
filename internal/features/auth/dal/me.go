/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/features/auth/dto"
)

type meRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewMeRepo .
func NewMeRepo(db *ent.Database, logger log.Logger) dto.MeRepo {
	return &meRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

func (r *meRepo) GetProfile(ctx context.Context, userID int64) (*types.User, error) {
	u, err := r.db.User(ctx).Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(u), nil
}
