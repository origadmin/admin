package dto

import (
	"context"
)

// MeRepo defines the interface for the current authenticated user's data access.
type MeRepo interface {
	// GetByID retrieves a user by their ID.
	GetByID(ctx context.Context, userID int64) (*UserPB, error)

	// GetProfile retrieves the current user's profile information.
	GetProfile(ctx context.Context, userID int64) (*UserProfilePB, error)

	// GetSetting retrieves the current user's settings.
	GetSetting(ctx context.Context, userID int64) (*UserSettingPB, error)

	// UpdateProfile updates the current user's profile information.
	UpdateProfile(ctx context.Context, userID int64, profileData *UserProfilePB) error

	// UpdateSetting updates the current user's settings.
	UpdateSetting(ctx context.Context, userID int64, settingsData *UserSettingPB) error

	// UpdatePreferences updates the current user's preferences.
	UpdatePreferences(ctx context.Context, userID int64, preferences map[string]string) error

	// UpdatePassword updates the user's password.
	// Expects plain text passwords (old and new), verification happens in the implementation.
	UpdatePassword(ctx context.Context, userID int64, hashedPassword string) error
}
