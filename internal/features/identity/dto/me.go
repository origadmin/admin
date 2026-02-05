package dto

import (
	"context"
	"origadmin/application/admin/api/v1/services/types"
)

// MeRepo defines the data access methods for the currently identityenticated user's profile and context.
type MeRepo interface {
	// GetProfile retrieves the basic profile information for the user.
	GetProfile(ctx context.Context, userID int64) (*types.User, error)

	// ListActiveViews retrieves all active views from the database.
	ListActiveViews(ctx context.Context) ([]*types.View, error)

	// GetPermissionKeywordsByUserID retrieves all permission keywords associated with a specific user ID.
	GetPermissionKeywordsByUserID(ctx context.Context, userID string) ([]string, error)

	// HasSystemRole checks if the user is assigned any role of type 'system'.
	HasSystemRole(ctx context.Context, userID int64) (bool, error)

	// UpdateProfile updates the user's profile information.
	UpdateProfile(ctx context.Context, userID int64, user *types.User) error

	// ChangePassword changes the user's password.
	ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error

	// UpdatePreferences updates user preferences (P2).
	UpdatePreferences(ctx context.Context, userID int64, preferences map[string]string) error

	// GetUserSettings retrieves user settings (P2).
	GetUserSettings(ctx context.Context, userID int64) (map[string]string, error)
}
