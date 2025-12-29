package dto

import (
	"context"
)

// AuthRepo defines the data access methods for authentication.
type AuthRepo interface {
	// GetUserByUsername retrieves a user by their username.
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}
