/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/origadmin/contrib/security"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/identity/dto"
)

// MeUseCase is the use case for managing the current user's profile.
type MeUseCase struct {
	meRepo    dto.MeRepo
	authzRepo dto.AuthzRepo
	log       *log.Helper
}

// NewMeUseCase creates a new MeUseCase.
func NewMeUseCase(meRepo dto.MeRepo, authzRepo dto.AuthzRepo, logger log.Logger) *MeUseCase {
	return &MeUseCase{
		meRepo:    meRepo,
		authzRepo: authzRepo,
		log:       log.NewHelper(log.With(logger, "module", "biz.me")),
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
		return nil, errors.InternalServer("INVALID_USER_ID", "Invalid user ID in token")
	}

	userEntity, err := uc.meRepo.GetUserWithRelations(ctx, userID)
	if err != nil {
		return nil, err
	}

	return dto.ConvertUserToUserPB(userEntity), nil
}

// UpdateProfile updates the current user's profile.
func (uc *MeUseCase) UpdateProfile(ctx context.Context, req *types.UpdateProfileRequest) (*types.User, error) {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return nil, errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return nil, errors.Internal("INVALID_USER_ID", "Invalid user ID in token")
	}

	// Convert from PB request to DTO request
	dtoReq := &dto.UpdateProfileRequest{
		Nickname: &req.Nickname,
		Avatar:   &req.Avatar,
		Gender:   &req.Gender,
	}

	updatedUser, err := uc.meRepo.UpdateProfile(ctx, userID, dtoReq)
	if err != nil {
		return nil, err
	}

	return dto.ConvertUserEntityToPB(updatedUser), nil
}

// ChangePassword changes the current user's password.
func (uc *MeUseCase) ChangePassword(ctx context.Context, oldPassword, newPassword string) error {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return errors.Internal("INVALID_USER_ID", "Invalid user ID in token")
	}
	return uc.meRepo.ChangePassword(ctx, userID, oldPassword, newPassword)
}

// UpdateSettings updates the current user's settings.
func (uc *MeUseCase) UpdateSettings(ctx context.Context, req *types.UpdateSettingsRequest) error {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return errors.Internal("INVALID_USER_ID", "Invalid user ID in token")
	}

	dtoReq := &dto.UpdateSettingsRequest{
		Theme:    &req.Theme,
		Language: &req.Language,
		Timezone: &req.Timezone,
	}

	_, err = uc.meRepo.UpdateSettings(ctx, userID, dtoReq)
	return err
}

// GetPermissions retrieves the current user's permission keywords.
func (uc *MeUseCase) GetPermissions(ctx context.Context) ([]string, error) {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return nil, errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return nil, errors.Internal("INVALID_USER_ID", "Invalid user ID in token")
	}
	return uc.authzRepo.GetPermissionKeywordsByUserID(ctx, userID)
}
