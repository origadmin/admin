package dto

import (
	"context"
	"origadmin/application/admin/api/v1/services/types"
)

// AuthedUser is an internal DTO for authentication, containing sensitive data.
// It should NOT be returned to the service layer or external clients.
type AuthedUser struct {
	User              *types.User
	EncryptedPassword string
}

// AuthRepo defines the data access methods for authentication.
type AuthRepo interface {
	// GetUserByUsername retrieves a user's auth-specific data by their username.
	GetUserByUsername(ctx context.Context, username string) (*AuthedUser, error)
	// UpdateLoginInfo updates the login-related fields for a user.
	UpdateLoginInfo(ctx context.Context, userID int64, loginIP string) error
	// GetAllViewsByScope retrieves all views for a given scope, ordered and structured as a tree.
	GetAllViewsByScope(ctx context.Context, scope string) ([]*types.View, error)
	// GetPermissionKeywordsByUserID retrieves all permission keywords associated with a user.
	GetPermissionKeywordsByUserID(ctx context.Context, userID string) ([]string, error)
}
