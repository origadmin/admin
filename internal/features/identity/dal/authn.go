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
	identitydto "origadmin/application/admin/internal/features/identity/dto"
)

// AuthnRepo implements the dto.AuthnRepo interface.
type AuthnRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewAuthnRepo creates a new AuthnRepo.
func NewAuthnRepo(db *ent.Database, logger log.Logger) identitydto.AuthnRepo {
	return &AuthnRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

// GetUserByCredential retrieves a user's identity-specific data by username, phone, or email.
func (r *AuthnRepo) GetUserByCredential(ctx context.Context, credential string) (*identitydto.AuthedUser, error) {
	u, err := r.db.User(ctx).Query().
		Where(
			user.Or(
				user.Username(credential),
				user.Phone(credential),
				user.Email(credential),
			),
		).
		WithRoles().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &identitydto.AuthedUser{
		User:              identitydto.ConvertUserToUserPB(u),
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
