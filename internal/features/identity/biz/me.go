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
	"github.com/origadmin/toolkits/crypto/hash"
	"origadmin/application/admin/internal/features/identity/dto"
)

// MeUseCase is the use case for managing the current user's profile.
type MeUseCase struct {
	meRepo    dto.MeRepo
	authzRepo dto.AuthzRepo
	hasher    hash.Crypto
	log       *log.Helper
}

// NewMeUseCase creates a new MeUseCase.
func NewMeUseCase(meRepo dto.MeRepo, authzRepo dto.AuthzRepo, hasher hash.Crypto, logger log.Logger) *MeUseCase {
	return &MeUseCase{
		meRepo:    meRepo,
		authzRepo: authzRepo,
		hasher:    hasher,
		log:       log.NewHelper(log.With(logger, "module", "biz.me")),
	}
}

// getUserID extracts and validates the user ID from the security context.
func (uc *MeUseCase) getUserID(ctx context.Context) (int64, error) {
	principal, ok := security.FromContext(ctx)
	if !ok {
		return 0, errors.Unauthorized("UNAUTHORIZED", "User not authenticated")
	}
	userID, err := strconv.ParseInt(principal.GetID(), 10, 64)
	if err != nil {
		return 0, errors.InternalServer("INVALID_USER_ID", "Invalid user ID in token")
	}
	return userID, nil
}

// GetProfile retrieves the current user's profile.
func (uc *MeUseCase) GetProfile(ctx context.Context) (*dto.UserProfilePB, error) {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return uc.meRepo.GetProfile(ctx, userID)
}

// GetSettings retrieves the current user's settings.
func (uc *MeUseCase) GetSettings(ctx context.Context) (*dto.UserSettingPB, error) {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return uc.meRepo.GetSettings(ctx, userID)
}

// UpdateProfile updates the current user's profile.
func (uc *MeUseCase) UpdateProfile(ctx context.Context, req *dto.UserProfilePB) error {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return err
	}

	return uc.meRepo.UpdateProfile(ctx, userID, req)
}

// UpdateSettings updates the current user's settings.
func (uc *MeUseCase) UpdateSettings(ctx context.Context, req *dto.UserSettingPB) error {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return err
	}

	return uc.meRepo.UpdateSettings(ctx, userID, req)
}

// UpdatePreferences updates the current user's preferences.
func (uc *MeUseCase) UpdatePreferences(ctx context.Context, preferences map[string]string) error {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return err
	}

	return uc.meRepo.UpdatePreferences(ctx, userID, preferences)
}

// ChangePassword changes the current user's password after verifying the old one.
func (uc *MeUseCase) ChangePassword(ctx context.Context, oldPassword, newPassword string) error {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return err
	}

	// Get the user to verify the old password
	user, err := uc.meRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// Verify the old password
	if err := uc.hasher.Verify(user.EncryptedPassword, oldPassword); err != nil {
		return errors.BadRequest("INVALID_PASSWORD", "Invalid old password")
	}

	// Hash the new password
	hashedPassword, err := uc.hasher.Hash(newPassword)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("failed to hash new password for user_id %d: %v", userID, err)
		return errors.InternalServer("PASSWORD_HASH_FAILED", "Failed to process new password")
	}

	// Update the password in the repository
	return uc.meRepo.UpdatePassword(ctx, userID, hashedPassword)
}

// ListMyPermissions retrieves the current user's permissions.
func (uc *MeUseCase) ListMyPermissions(ctx context.Context) (dto.PermissionsPB, error) {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return uc.authzRepo.ListMyPermissions(ctx, userID)
}

// ListMyRoles retrieves the current user's roles.
func (uc *MeUseCase) ListMyRoles(ctx context.Context) (dto.RolesPB, error) {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return uc.authzRepo.ListMyRoles(ctx, userID)
}

// ListMyViews retrieves the current user's views.
func (uc *MeUseCase) ListMyViews(ctx context.Context) (dto.ViewsPB, error) {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return uc.authzRepo.ListMyViews(ctx, userID)
}

// ListMyResources retrieves the current user's resources.
func (uc *MeUseCase) ListMyResources(ctx context.Context) (dto.ResourcesPB, error) {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return uc.authzRepo.ListMyResources(ctx, userID)
}

// GetPermissionKeywords retrieves the current user's permission keywords.
func (uc *MeUseCase) GetPermissionKeywords(ctx context.Context) ([]string, error) {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return uc.authzRepo.GetPermissionKeywordsByUserID(ctx, userID)
}

// HasSystemRole checks if the current user has a system role.
func (uc *MeUseCase) HasSystemRole(ctx context.Context) (bool, error) {
	userID, err := uc.getUserID(ctx)
	if err != nil {
		return false, err
	}

	return uc.authzRepo.HasSystemRole(ctx, userID)
}
