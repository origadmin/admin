package dto

import (
	"context"

	"origadmin/application/admin/api/v1/services/types"
)

// AuthedUser is a data transfer object that encapsulates authentication-specific user data.
// It includes the publicly safe User object and the sensitive encrypted password,
// which should never be exposed outside the authentication business logic.
type AuthedUser struct {
	*types.User
	EncryptedPassword string
}

// AuthRepo defines the repository interface for authentication-related operations.
type AuthRepo interface {
	// GetUserByUsername retrieves a user's essential authentication data,
	// including the encrypted password, by their username.
	GetUserByUsername(ctx context.Context, username string) (*AuthedUser, error)

	// UpdateLoginInfo updates the user's last login timestamp and IP address.
	UpdateLoginInfo(ctx context.Context, userID int64, loginIP string) error
}
