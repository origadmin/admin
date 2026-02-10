package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/origadmin/toolkits/crypto/hash"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/identity/dto"
)

// AuthUseCase is a identityentication use case.
type AuthUseCase struct {
	repo   dto.AuthnRepo
	hasher hash.Crypto
	log    *log.Helper
}

// NewAuthUseCase new a identityentication use case.
func NewAuthUseCase(repo dto.AuthnRepo, hasher hash.Crypto, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{
		repo:   repo,
		hasher: hasher,
		log:    log.NewHelper(logger),
	}
}

// VerifyUser verifies the user's credentials and returns the secure DTO if successful.
func (uc *AuthUseCase) VerifyUser(ctx context.Context, username, password string) (*types.User, error) {
	// 1. Get the internal AuthedUser DTO from the AuthRepo.
	identityedUser, err := uc.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// 2. Compare the provided password with the stored hash.
	if err := uc.hasher.Verify(identityedUser.EncryptedPassword, password); err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 3. On successful verification, return the safe User object from the DTO.
	return identityedUser.User, nil
}

// UpdateLoginInfo delegates the update of login-related information to the repository.
func (uc *AuthUseCase) UpdateLoginInfo(ctx context.Context, userID int64, loginIP string) error {
	return uc.repo.UpdateLoginInfo(ctx, userID, loginIP)
}
