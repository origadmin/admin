/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"

	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/entity/ent/userprofile"
	"origadmin/application/admin/internal/data/entity/ent/usersetting"
	"origadmin/application/admin/internal/features/identity/dto"
)

// MeRepo implements the dto.MeRepo interface for handling "me" (current user) related data operations.
type MeRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewMeRepo creates a new MeRepo with the given database client and logger.
func NewMeRepo(db *ent.Database, logger log.Logger) dto.MeRepo {
	return &MeRepo{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "dal/me")),
	}
}

// GetUserWithRelations retrieves a user by ID along with their profile and settings.
func (r *MeRepo) GetUserWithRelations(ctx context.Context, userID int64) (*ent.User, error) {
	u, err := r.db.User(ctx).Query().
		Where(user.ID(userID)).
		WithProfile().
		WithSetting().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		r.log.WithContext(ctx).Errorf("failed to get user with relations for user_id %d: %v", userID, err)
		return nil, errors.InternalServer("DATABASE_ERROR", "failed to retrieve user data")
	}
	return u, nil
}

// UpdateProfile updates the user's profile information within a transaction.
func (r *MeRepo) UpdateProfile(ctx context.Context, userID int64, profileData *dto.UpdateProfileRequest) (*ent.User, error) {
	// Use the project's custom transaction wrapper for consistency with other DAL implementations like user.go.
	err := r.db.Tx(ctx, func(tx context.Context) error {
		// The query must start from the transactional client, which is obtained by passing the tx context.
		profile, err := r.db.User(tx).Query().Where(user.ID(userID)).QueryProfile().Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return errors.NotFound("USER_PROFILE_NOT_FOUND", "User profile not found")
			}
			r.log.WithContext(ctx).Errorf("failed to query profile for user_id %d: %v", userID, err)
			return errors.InternalServer("DATABASE_ERROR", "failed to query user profile")
		}
		profileUpdate := profile.Update()
		if profileData.Nickname != nil {
			profileUpdate.SetNickname(*profileData.Nickname)
		}
		if profileData.Avatar != nil {
			profileUpdate.SetAvatar(*profileData.Avatar)
		}
		if profileData.Gender != nil {
			profileUpdate.SetGender(userprofile.Gender(*profileData.Gender))
		}
		// The update builder is derived from `profile` which was queried in the transaction, so it's transactional.
		// The context passed to Exec is for cancellation/deadlines.
		if err := profileUpdate.Exec(ctx); err != nil {
			r.log.WithContext(ctx).Errorf("failed to update profile for user_id %d: %v", userID, err)
			return errors.InternalServer("PROFILE_UPDATE_FAILED", "failed to update profile")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return r.GetUserWithRelations(ctx, userID)
}

// ChangePassword verifies the old password and updates it to the new one.
func (r *MeRepo) ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) (err error) {
	u, err := r.db.User(ctx).Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		r.log.WithContext(ctx).Errorf("failed to find user for password change, user_id %d: %v", userID, err)
		return errors.InternalServer("DATABASE_ERROR", "failed to retrieve user")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(oldPassword)); err != nil {
		return errors.BadRequest("INVALID_PASSWORD", "Invalid old password")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to hash new password for user_id %d: %v", userID, err)
		return errors.InternalServer("PASSWORD_HASH_FAILED", "Failed to process new password")
	}
	err = r.db.User(ctx).UpdateOneID(userID).SetEncryptedPassword(string(hashedPassword)).Exec(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to update password for user_id %d: %v", userID, err)
		return errors.InternalServer("DATABASE_ERROR", "failed to update password")
	}
	return nil
}

// UpdateSettings updates the user's application settings.
func (r *MeRepo) UpdateSettings(ctx context.Context, userID int64, settingsData *dto.UpdateSettingsRequest) (*ent.UserSetting, error) {
	// Query the setting directly for clarity and efficiency.
	setting, err := r.db.UserSetting(ctx).Query().
		Where(usersetting.HasUserWith(user.ID(userID))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_SETTINGS_NOT_FOUND", "User settings not found")
		}
		r.log.WithContext(ctx).Errorf("failed to retrieve settings for user_id %d: %v", userID, err)
		return nil, errors.InternalServer("DATABASE_ERROR", "failed to retrieve user settings")
	}
	update := setting.Update()
	if settingsData.Theme != nil {
		update.SetTheme(*settingsData.Theme)
	}
	if settingsData.Language != nil {
		update.SetLanguage(*settingsData.Language)
	}
	if settingsData.Timezone != nil {
		update.SetTimezone(*settingsData.Timezone)
	}
	updatedSetting, err := update.Save(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to update settings for user_id %d: %v", userID, err)
		return nil, errors.InternalServer("DATABASE_ERROR", "failed to update settings")
	}
	return updatedSetting, nil
}
