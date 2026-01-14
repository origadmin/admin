package dto

import (
	"context"
	"origadmin/application/admin/api/v1/services/types"
)

// MeRepo defines the data access methods for the currently authenticated user's profile and context.
type MeRepo interface {
	// GetProfile retrieves the basic profile information for the user.
	GetProfile(ctx context.Context, userID int64) (*types.User, error)

	// GetAllViewsByScope retrieves all views for a given scope, used to build a complete view tree.
	GetAllViewsByScope(ctx context.Context, scope string) ([]*types.View, error)

	// GetPermissionKeywordsByUserID retrieves all permission keywords associated with a specific user ID.
	GetPermissionKeywordsByUserID(ctx context.Context, userID string) ([]string, error)
}
