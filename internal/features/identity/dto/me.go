package dto

import (
	"context"
)

// MeRepo defines the interface for the current authenticated user's data access.
type MeRepo interface {
	// GetUser retrieves a user by their ID.
	GetUser(ctx context.Context, userID int64) (*User, error)

	// GetProfile retrieves the current user's profile information.
	GetProfile(ctx context.Context, userID int64) (*UserProfilePB, error)

	// GetSettings retrieves the current user's settings.
	GetSettings(ctx context.Context, userID int64) (*UserSettingPB, error)

	// UpdateProfile updates the current user's profile information.
	UpdateProfile(ctx context.Context, userID int64, profileData *UserProfilePB) error

	// UpdateSettings updates the current user's settings.
	UpdateSettings(ctx context.Context, userID int64, settingsData *UserSettingPB) error

	// UpdatePreferences updates the current user's preferences.
	UpdatePreferences(ctx context.Context, userID int64, preferences map[string]string) error

	// UpdatePassword updates the user's password in the database.
	UpdatePassword(ctx context.Context, userID int64, hashedPassword string) error
}
