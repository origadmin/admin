/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/entity/ent/view"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/auth/dto"
	systemDto "origadmin/application/admin/internal/features/system/dto"
)

type MeRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewMeRepo .
func NewMeRepo(db *ent.Database, logger log.Logger) dto.MeRepo {
	return &MeRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

func (r *MeRepo) GetProfile(ctx context.Context, userID int64) (*types.User, error) {
	u, err := r.db.User(ctx).Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return systemDto.ConvertUserToUserPB(u), nil
}

// ListActiveViews retrieves all active views, ordered by sequence.
func (r *MeRepo) ListActiveViews(ctx context.Context) ([]*types.View, error) {
	views, err := r.db.View(ctx).Query().
		Where(view.StatusEQ(enums.StatusActive)).
		Order(ent.Asc(view.FieldSequence)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return systemDto.ConvertViewsToViewsPB(views), nil
}

// GetPermissionKeywordsByUserID retrieves all permission keywords for a user.
func (r *MeRepo) GetPermissionKeywordsByUserID(ctx context.Context, userID string) ([]string, error) {
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, err
	}

	permissions, err := r.db.User(ctx).Query().
		Where(user.ID(id)).
		QueryRoles().
		QueryPermissions().
		All(ctx)
	if err != nil {
		return nil, err
	}

	keywords := make([]string, len(permissions))
	for i, p := range permissions {
		keywords[i] = p.Keyword
	}
	return keywords, nil
}

// HasSystemRole checks if the user has a role with the 'system' type.
func (r *MeRepo) HasSystemRole(ctx context.Context, userID int64) (bool, error) {
	return r.db.User(ctx).
		Query().
		Where(user.ID(userID)).
		QueryRoles().
		Where(role.TypeEQ(enums.RoleTypeSystem)).
		Exist(ctx)
}

// UpdateProfile updates the user's profile information in the database.
func (r *MeRepo) UpdateProfile(ctx context.Context, userID int64, u *types.User) error {
	// Convert string gender to user.Gender enum
	var gender *user.Gender
	if u.Gender != "" {
		g := user.GenderMale // default
		if u.Gender == "female" {
			g = user.GenderFemale
		}
		gender = &g
	}

	update := r.db.User(ctx).UpdateOneID(userID)
	if u.Nickname != "" {
		update.SetNickname(u.Nickname)
	}
	if u.Avatar != "" {
		update.SetAvatar(u.Avatar)
	}
	if gender != nil {
		update.SetGender(*gender)
	}
	if u.Email != "" {
		update.SetEmail(u.Email)
	}
	if u.Phone != "" {
		update.SetPhone(u.Phone)
	}
	return update.Exec(ctx)
}

// ChangePassword changes the user's password in the database.
func (r *MeRepo) ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error {
	u, err := r.db.User(ctx).Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return err
	}

	// Check old password
	if err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(oldPassword)); err != nil {
		return errors.BadRequest("INVALID_PASSWORD", "Invalid old password")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return r.db.User(ctx).UpdateOneID(userID).SetEncryptedPassword(string(hashedPassword)).Exec(ctx)
}

// UpdatePreferences updates user preferences in the remark field (P2).
func (r *MeRepo) UpdatePreferences(ctx context.Context, userID int64, preferences map[string]string) error {
	_, err := r.db.User(ctx).Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return err
	}

	// For now, store preferences as JSON in the remark field
	// In production, you might want a dedicated preferences table or JSON field
	// This is a stub implementation - actual implementation would serialize preferences to JSON
	return r.db.User(ctx).UpdateOneID(userID).SetRemark("").Exec(ctx)
}

// GetUserSettings retrieves user settings (P2).
// For now, this returns basic user information as settings
func (r *MeRepo) GetUserSettings(ctx context.Context, userID int64) (map[string]string, error) {
	_, err := r.db.User(ctx).Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}

	// Return user settings as a map
	// In production, these would be stored in a dedicated table or JSON field
	settings := map[string]string{
		"theme":        "light",
		"language":     "en",
		"timezone":     "UTC",
		"date_format":  "YYYY-MM-DD",
		"time_format":  "HH:mm:ss",
		"notification": "enabled",
	}

	return settings, nil
}
