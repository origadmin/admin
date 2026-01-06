package dal

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/features/auth/dto"
)

type AuthRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(db *ent.Database, logger log.Logger) dto.AuthRepo {
	return &AuthRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

// GetUserByUsername retrieves a user's auth-specific data by their username.
func (r *AuthRepo) GetUserByUsername(ctx context.Context, username string) (*dto.AuthedUser, error) {
	u, err := r.db.User(ctx).Query().Where(user.UsernameEQ(username)).WithRoles().Only(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.AuthedUser{
		User:              dto.ConvertUserToUserPB(u),
		EncryptedPassword: u.EncryptedPassword,
	}, nil
}

// UpdateLoginInfo updates the last login time, current login time, and last login IP for a user.
func (r *AuthRepo) UpdateLoginInfo(ctx context.Context, userID int64, loginIP string) error {
	// First, get the current user entity to perform the "shift change".
	currentUser, err := r.db.User(ctx).Get(ctx, userID)
	if err != nil {
		return err
	}

	// Perform the "shift change" and update to the new values.
	return r.db.User(ctx).
		UpdateOneID(userID).
		SetLastLoginTime(currentUser.LoginTime). // Previous login time becomes the last login time
		SetLoginTime(time.Now()).               // Set current login time
		SetLastLoginIP(currentUser.LoginIP).    // Previous login IP becomes the last login IP
		SetLoginIP(loginIP).                    // Set current login IP
		Exec(ctx)
}
