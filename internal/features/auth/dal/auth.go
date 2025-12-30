package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/features/auth/dto"
)

type authRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(db *ent.Database, logger log.Logger) dto.AuthRepo {
	return &authRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

// GetUserByUsername retrieves a user by their username.
func (r *authRepo) GetUserByUsername(ctx context.Context, username string) (*dto.User, error) {
	u, err := r.db.User(ctx).Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		ID:                u.ID,
		Username:          u.Username,
		EncryptedPassword: u.EncryptedPassword,
	}, nil
}
