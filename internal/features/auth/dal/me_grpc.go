/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"sort"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/auth/dto"
)

// meGRPCRepo implements the MeRepo interface using gRPC calls to the system service.
type meGRPCRepo struct {
	userClient systemv1.UserServiceClient
	viewClient systemv1.ViewServiceClient
	log        *log.Helper
}

// NewMeGRPCRepo creates a new MeRepo that communicates with the system service via gRPC.
func NewMeGRPCRepo(
	userClient systemv1.UserServiceClient,
	viewClient systemv1.ViewServiceClient,
	logger log.Logger,
) dto.MeRepo {
	return &meGRPCRepo{
		userClient: userClient,
		viewClient: viewClient,
		log:        log.NewHelper(log.With(logger, "module", "dal.me_grpc")),
	}
}

// GetProfile retrieves user profile information via a gRPC call to the UserService.
func (r *meGRPCRepo) GetProfile(ctx context.Context, userID int64) (*types.User, error) {
	r.log.WithContext(ctx).Debugf("Fetching profile for user ID %d via gRPC", userID)
	resp, err := r.userClient.GetUser(ctx, &systemv1.GetUserRequest{Id: userID})
	if err != nil {
		return nil, err
	}
	return resp.GetUser(), nil
}

// ListActiveViews retrieves all active views via a gRPC call to the ViewService.
// It performs client-side filtering and sorting to match the behavior of the database implementation.
func (r *meGRPCRepo) ListActiveViews(ctx context.Context) ([]*types.View, error) {
	r.log.WithContext(ctx).Debug("Listing all views via gRPC and filtering for active ones")

	// 1. Fetch all views, as the gRPC endpoint does not support filtering by status.
	resp, err := r.viewClient.ListViews(ctx, &systemv1.ListViewsRequest{
		// Request all items to ensure we can filter and sort correctly.
		PageSize: -1, // Assuming -1 or a large number fetches all items.
	})
	if err != nil {
		return nil, err
	}

	allViews := resp.GetViews()
	activeViews := make([]*types.View, 0, len(allViews))

	// 2. Filter for active views on the client-side.
	for _, view := range allViews {
		if view.Status == int32(enums.StatusActive) {
			activeViews = append(activeViews, view)
		}
	}

	// 3. Sort the active views by sequence to match the database implementation's behavior.
	sort.Slice(activeViews, func(i, j int) bool {
		return activeViews[i].Sequence < activeViews[j].Sequence
	})

	return activeViews, nil
}

// GetPermissionKeywordsByUserID retrieves user permission keywords via a gRPC call to the UserService.
func (r *meGRPCRepo) GetPermissionKeywordsByUserID(ctx context.Context, userID string) ([]string, error) {
	r.log.WithContext(ctx).Debugf("Fetching permission keywords for user ID %s via gRPC", userID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, err
	}

	// The ListUserResources method is the correct one to get permissions.
	resp, err := r.userClient.ListUserResources(ctx, &systemv1.ListUserResourcesRequest{Id: id})
	if err != nil {
		return nil, err
	}

	// Extract unique permission keywords from the resources.
	keywordSet := make(map[string]struct{})
	for _, resource := range resp.GetResources() {
		// A resource might not have an associated permission.
		if len(resource.Permissions) > 0 {
			for _, permission := range resource.Permissions {
				if permission.Keyword != "" {
					keywordSet[permission.Keyword] = struct{}{}
				}
			}
		}
	}

	keywords := make([]string, 0, len(keywordSet))
	for k := range keywordSet {
		keywords = append(keywords, k)
	}

	return keywords, nil
}

// HasSystemRole checks for system role membership via a gRPC call to the UserService.
func (r *meGRPCRepo) HasSystemRole(ctx context.Context, userID int64) (bool, error) {
	r.log.WithContext(ctx).Debugf("Checking system role for user ID %d via gRPC", userID)
	resp, err := r.userClient.GetUser(ctx, &systemv1.GetUserRequest{Id: userID, WithRoles: true})
	if err != nil {
		return false, err
	}

	user := resp.GetUser()
	if user == nil {
		return false, nil
	}

	for _, role := range user.GetRoles() {
		// The type of role.Type is int32, so we compare it with the int32 value of the enum.
		if role.Type == int32(enums.RoleTypeSystem) {
			return true, nil
		}
	}

	return false, nil
}

// UpdateProfile updates the user's profile information via a gRPC call to the UserService.
func (r *meGRPCRepo) UpdateProfile(ctx context.Context, userID int64, user *types.User) error {
	// The UpdateUser RPC in systemv1.UserServiceClient can be used to update user profiles.
	// We need to construct a FieldMask to specify which fields are being updated.
	_, err := r.userClient.UpdateUser(ctx, &systemv1.UpdateUserRequest{
		User: &types.User{
			Id:       userID,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Gender:   user.Gender,
			Email:    user.Email,
			Phone:    user.Phone,
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"nickname", "avatar", "gender", "email", "phone"},
		},
	})
	return err
}

// ChangePassword changes the user's password.
// This method cannot be implemented securely via gRPC with the current UserService API,
// as there's no RPC to verify the old password or update the password directly for the current user.
// ResetUserPassword is an admin-level operation.
func (r *meGRPCRepo) ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error {
	r.log.WithContext(ctx).Warnf("ChangePassword for user ID %d is not implemented via gRPC due to API limitations.", userID)
	return errors.ServiceUnavailable("CHANGE_PASSWORD_UNIMPLEMENTED", "Change password is not implemented via gRPC for current user due to UserService API limitations.")
}

// UpdatePreferences updates user preferences via gRPC (P2).
func (r *meGRPCRepo) UpdatePreferences(ctx context.Context, userID int64, preferences map[string]string) error {
	r.log.WithContext(ctx).Debugf("UpdatePreferences for user ID %d via gRPC is not implemented", userID)
	// This would require a dedicated preferences service or UserService extension
	return errors.ServiceUnavailable("UPDATE_PREFERENCES_UNIMPLEMENTED", "Update preferences is not implemented via gRPC yet.")
}

// GetUserSettings retrieves user settings via gRPC (P2).
func (r *meGRPCRepo) GetUserSettings(ctx context.Context, userID int64) (map[string]string, error) {
	r.log.WithContext(ctx).Debugf("GetUserSettings for user ID %d via gRPC is not implemented", userID)
	// This would require a dedicated settings service or UserService extension
	return nil, errors.ServiceUnavailable("GET_USER_SETTINGS_UNIMPLEMENTED", "Get user settings is not implemented via gRPC yet.")
}

var _ dto.MeRepo = (*meGRPCRepo)(nil)
