package biz

import (
	"context"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/auth/dto"

	"github.com/go-kratos/kratos/v2/log"
)

// MeUseCase is a user profile use case.
type MeUseCase struct {
	repo dto.MeRepo
	log  *log.Helper
}

// NewMeUseCase new a user profile use case.
func NewMeUseCase(repo dto.MeRepo, logger log.Logger) *MeUseCase {
	return &MeUseCase{repo: repo, log: log.NewHelper(logger)}
}

// GetProfile gets the user profile.
func (uc *MeUseCase) GetProfile(ctx context.Context, userID int64) (*types.User, error) {
	return uc.repo.GetProfile(ctx, userID)
}
