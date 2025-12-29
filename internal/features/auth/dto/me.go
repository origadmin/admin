package dto

import (
	"context"
	"origadmin/application/admin/api/v1/services/types"
)

// MeRepo defines the data access methods for user profile.
type MeRepo interface {
	GetProfile(ctx context.Context, userID int64) (*types.User, error)
}
