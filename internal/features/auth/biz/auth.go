package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"

	"origadmin/application/admin/internal/features/auth/dto"
)

// AuthUseCase is a authentication use case.
type AuthUseCase struct {
	repo dto.AuthRepo
	log  *log.Helper
}

// NewAuthUseCase new a authentication use case.
func NewAuthUseCase(repo dto.AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// VerifyUser verifies the user's credentials and returns the user ID if successful.
func (uc *AuthUseCase) VerifyUser(ctx context.Context, username, password string) (int64, error) {
	user, err := uc.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return 0, err
	}

	// Compare the provided password with the stored hash.
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		// If the passwords don't match, return a generic error.
		return 0, errors.New("invalid username or password")
	}

	return user.ID, nil
}
