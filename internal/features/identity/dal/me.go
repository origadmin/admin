/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/entity/ent/userprofile"
	"origadmin/application/admin/internal/data/entity/ent/usersetting"
	identitydto "origadmin/application/admin/internal/features/identity/dto"
)

// MeRepo implements the dto.MeRepo interface for handling "me" (current user) related data operations.
type MeRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewMeRepo creates a new MeRepo with the given database client and logger.
func NewMeRepo(db *ent.Database, logger log.Logger) identitydto.MeRepo {
	return &MeRepo{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "dal.me")),
	}
}

// GetByID retrieves a user by their ID.
func (r *MeRepo) GetByID(ctx context.Context, userID int64) (*identitydto.UserPB, error) {
	u, err := r.db.User(ctx).Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		r.log.WithContext(ctx).Errorf("failed to find user, user_id %d: %v", userID, err)
		return nil, errors.InternalServer("DATABASE_ERROR", "failed to retrieve user")
	}
	return identitydto.ConvertUserToUserPB(u), nil
}

// GetSetting retrieves the user's settings by ID.
func (r *MeRepo) GetSetting(ctx context.Context, userID int64) (*identitydto.UserSettingPB, error) {
	setting, err := r.db.User(ctx).Query().
		Where(user.ID(userID)).
		QuerySetting().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_SETTINGS_NOT_FOUND", "User settings not found")
		}
		r.log.WithContext(ctx).Errorf("failed to get settings for user_id %d: %v", userID, err)
		return nil, errors.InternalServer("DATABASE_ERROR", "failed to retrieve user settings")
	}
	return identitydto.ConvertUserSettingToUserSettingPB(setting), nil
}

// UpdateSetting updates the user's application settings.
func (r *MeRepo) UpdateSetting(ctx context.Context, userID int64, settingsData *identitydto.UserSettingPB) error {
	setting, err := r.db.UserSetting(ctx).Query().
		Where(usersetting.HasUserWith(user.ID(userID))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.NotFound("USER_SETTINGS_NOT_FOUND", "User settings not found")
		}
		r.log.WithContext(ctx).Errorf("failed to retrieve settings for user_id %d: %v", userID, err)
		return errors.InternalServer("DATABASE_ERROR", "failed to retrieve user settings")
	}
	update := setting.Update()
	if settingsData.Theme != "" {
		update.SetTheme(settingsData.Theme)
	}
	if settingsData.Language != "" {
		update.SetLanguage(settingsData.Language)
	}
	if settingsData.Timezone != "" {
		update.SetTimezone(settingsData.Timezone)
	}
	if settingsData.Preferences != nil && len(settingsData.Preferences) > 0 {
		// Convert map[string]string to map[string]interface{} for JSON field
		preferencesMap := make(map[string]interface{}, len(settingsData.Preferences))
		for k, v := range settingsData.Preferences {
			preferencesMap[k] = v
		}
		update.SetPreferences(preferencesMap)
	}
	return update.Exec(ctx)
}

// GetProfile retrieves the user's profile by ID.
func (r *MeRepo) GetProfile(ctx context.Context, userID int64) (*identitydto.UserProfilePB, error) {
	profile, err := r.db.User(ctx).Query().
		Where(user.ID(userID)).
		QueryProfile().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_PROFILE_NOT_FOUND", "User profile not found")
		}
		r.log.WithContext(ctx).Errorf("failed to get profile for user_id %d: %v", userID, err)
		return nil, errors.InternalServer("DATABASE_ERROR", "failed to retrieve user profile")
	}
	return identitydto.ConvertUserProfileToUserProfilePB(profile), nil
}

// UpdateProfile updates the user's profile information within a transaction.
func (r *MeRepo) UpdateProfile(ctx context.Context, userID int64, profileData *identitydto.UserProfilePB) error {
	return r.db.Tx(ctx, func(tx context.Context) error {
		profile, err := r.db.User(tx).Query().Where(user.ID(userID)).QueryProfile().Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return errors.NotFound("USER_PROFILE_NOT_FOUND", "User profile not found")
			}
			r.log.WithContext(ctx).Errorf("failed to query profile for user_id %d: %v", userID, err)
			return errors.InternalServer("DATABASE_ERROR", "failed to query user profile")
		}
		profileUpdate := profile.Update()
		if profileData.Avatar != "" {
			profileUpdate.SetAvatar(profileData.Avatar)
		}
		if profileData.Gender != "" {
			profileUpdate.SetGender(userprofile.Gender(profileData.Gender))
		}
		if profileData.Name != "" {
			profileUpdate.SetName(profileData.Name)
		}
		if profileData.Department != "" {
			profileUpdate.SetDepartment(profileData.Department)
		}
		if profileData.Remark != "" {
			profileUpdate.SetRemark(profileData.Remark)
		}
		if err := profileUpdate.Exec(ctx); err != nil {
			r.log.WithContext(ctx).Errorf("failed to update profile for user_id %d: %v", userID, err)
			return errors.InternalServer("PROFILE_UPDATE_FAILED", "failed to update profile")
		}
		return nil
	})
}

// UpdatePassword updates the user's encrypted password in the database.
func (r *MeRepo) UpdatePassword(ctx context.Context, userID int64, hashedPassword string) error {
	err := r.db.User(ctx).UpdateOneID(userID).SetEncryptedPassword(hashedPassword).Exec(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to update password for user_id %d: %v", userID, err)
		return errors.InternalServer("DATABASE_ERROR", "failed to update password")
	}
	return nil
}

// UpdatePreferences updates the user's preferences.
func (r *MeRepo) UpdatePreferences(ctx context.Context, userID int64, preferences map[string]string) error {
	return r.db.Tx(ctx, func(tx context.Context) error {
		setting, err := r.db.UserSetting(tx).Query().
			Where(usersetting.HasUserWith(user.ID(userID))).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return errors.NotFound("USER_SETTINGS_NOT_FOUND", "User settings not found")
			}
			r.log.WithContext(ctx).Errorf("failed to retrieve settings for user_id %d: %v", userID, err)
			return errors.InternalServer("DATABASE_ERROR", "failed to retrieve user settings")
		}

		// Convert map[string]string to map[string]interface{} for JSON field
		preferencesMap := make(map[string]interface{}, len(preferences))
		for k, v := range preferences {
			preferencesMap[k] = v
		}

		return setting.Update().SetPreferences(preferencesMap).Exec(ctx)
	})
}
