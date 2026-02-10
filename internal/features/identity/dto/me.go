package dto

import (
	"context"

	"origadmin/application/admin/internal/data/entity/ent"
)

// MeRepo defines the interface for the current authenticated user's data access.
type MeRepo interface {
	// GetUserWithRelations retrieves the current user and their profile and settings by ID.
	GetUserWithRelations(ctx context.Context, userID int64) (*ent.User, error)

	// UpdateProfile updates the current user's profile information.
	UpdateProfile(ctx context.Context, userID int64, profileData *UpdateProfileRequest) (*ent.User, error)

	// ChangePassword changes the current user's password.
	ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error

	// UpdateSettings updates the current user's settings.
	UpdateSettings(ctx context.Context, userID int64, settingsData *UpdateSettingsRequest) (*ent.UserSetting, error)
}

// UpdateProfileRequest represents the data for updating a user's profile.
type UpdateProfileRequest struct {
	Nickname *string
	Avatar   *string
	Gender   *string
}

// UpdateSettingsRequest represents the data for updating user settings.
type UpdateSettingsRequest struct {
	Theme    *string
	Language *string
	Timezone *string
}
