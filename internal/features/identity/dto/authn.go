package dto

import (
	"context"

	"origadmin/application/admin/api/v1/services/types"
)

// AuthedUser is a data transfer object that encapsulates authentication-specific user data.
// It includes a User object and sensitive encrypted password,
// which should never be exposed outside of authentication business logic.
type AuthedUser struct {
	User              *types.User
	EncryptedPassword string
}

// AuthnRepo defines the repository interface for authentication-related operations.
type AuthnRepo interface {
	// GetUserByCredential retrieves a user's essential authentication data,
	// including the encrypted password, by their username, phone, or email.
	GetUserByCredential(ctx context.Context, credential string) (*AuthedUser, error)

	// UpdateLoginInfo updates the user's last login timestamp and IP address.
	UpdateLoginInfo(ctx context.Context, userID int64, loginIP string) error
}
