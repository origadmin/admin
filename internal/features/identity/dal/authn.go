/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/features/identity/dto"
	systemDto "origadmin/application/admin/internal/features/system/dto"
)

// AuthnRepo implements the dto.AuthnRepo interface.
type AuthnRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewAuthnRepo creates a new AuthnRepo.
func NewAuthnRepo(db *ent.Database, logger log.Logger) dto.AuthnRepo {
	return &AuthnRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

// GetUserByUsername retrieves a user's identity-specific data by their username.
func (r *AuthnRepo) GetUserByUsername(ctx context.Context, username string) (*dto.AuthedUser, error) {
	u, err := r.db.User(ctx).Query().Where(user.UsernameEQ(username)).WithRoles().Only(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.AuthedUser{
		User:              systemDto.ConvertUserToUserPB(u),
		EncryptedPassword: u.EncryptedPassword,
	}, nil
}

// UpdateLoginInfo updates the last login time, current login time, and last login IP for a user.
func (r *AuthnRepo) UpdateLoginInfo(ctx context.Context, userID int64, loginIP string) error {
	currentUser, err := r.db.User(ctx).Get(ctx, userID)
	if err != nil {
		return err
	}
	return r.db.User(ctx).
		UpdateOneID(userID).
		SetLastLoginTime(currentUser.LoginTime).
		SetLoginTime(time.Now()).
		SetLastLoginIP(currentUser.LoginIP).
		SetLoginIP(loginIP).
		Exec(ctx)
}
