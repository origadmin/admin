/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/auth/dto"
	"github.com/origadmin/contrib/security"
)

// MeUseCase is the use case for managing the current user's profile.
type MeUseCase struct {
	repo dto.MeRepo
	log  *log.Helper
}

// NewMeUseCase creates a new MeUseCase.
func NewMeUseCase(repo dto.MeRepo, logger log.Logger) *MeUseCase {
	return &MeUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz.me")),
	}
}

// GetProfile retrieves the current user's profile.
func (uc *MeUseCase) GetProfile(ctx context.Context) (*types.User, error) {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return nil, errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return nil, err
	}
	return uc.repo.GetProfile(ctx, userID)
}

// UpdateProfile updates the current user's profile.
func (uc *MeUseCase) UpdateProfile(ctx context.Context, user *types.User) error {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return err
	}
	return uc.repo.UpdateProfile(ctx, userID, user)
}

// ChangePassword changes the current user's password.
func (uc *MeUseCase) ChangePassword(ctx context.Context, oldPassword, newPassword string) error {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return err
	}
	return uc.repo.ChangePassword(ctx, userID, oldPassword, newPassword)
}

// UpdatePreferences updates the current user's preferences (P2).
func (uc *MeUseCase) UpdatePreferences(ctx context.Context, preferences map[string]string) error {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return err
	}
	return uc.repo.UpdatePreferences(ctx, userID, preferences)
}

// GetUserSettings retrieves the current user's settings (P2).
func (uc *MeUseCase) GetUserSettings(ctx context.Context) (map[string]string, error) {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return nil, errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return nil, err
	}
	return uc.repo.GetUserSettings(ctx, userID)
}
