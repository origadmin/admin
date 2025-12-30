package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/origadmin/toolkits/crypto/hash"

	"origadmin/application/admin/internal/features/auth/dto"
)

// AuthUseCase is a authentication use case.
type AuthUseCase struct {
	repo   dto.AuthRepo
	hasher hash.Crypto
	log    *log.Helper
}

// NewAuthUseCase new a authentication use case.
func NewAuthUseCase(repo dto.AuthRepo, hasher hash.Crypto, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		repo:   repo,
		hasher: hasher,
		log:    log.NewHelper(logger),
	}
}

// VerifyUser verifies the user's credentials and returns the user ID if successful.
func (uc *AuthUseCase) VerifyUser(ctx context.Context, username, password string) (int64, error) {
	user, err := uc.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return 0, err
	}

	// Compare the provided password with the stored hash.
	if err := uc.hasher.Verify(user.EncryptedPassword, password); err != nil {
		return 0, errors.New("invalid username or password")
	}

	return user.ID, nil
}
