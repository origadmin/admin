/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/identity/dto"
)

// meGRPCRepo implements dto.MeRepo interface using gRPC calls to system service.
type meGRPCRepo struct {
	userClient systemv1.UserServiceClient
	log        *log.Helper
}

// NewMeGRPCRepo creates a new MeRepo that communicates with system service via gRPC.
func NewMeGRPCRepo(
	userClient systemv1.UserServiceClient,
	logger log.Logger,
) dto.MeRepo {
	return &meGRPCRepo{
		userClient: userClient,
		log:        log.NewHelper(log.With(logger, "module", "dal.me_grpc")),
	}
}

func (r *meGRPCRepo) GetByID(ctx context.Context, userID int64) (*dto.UserPB, error) {
	//TODO implement me
	panic("implement me")
}

// GetProfile retrieves user profile information via a gRPC call to UserService.
func (r *meGRPCRepo) GetProfile(ctx context.Context, userID int64) (*dto.UserProfilePB, error) {
	r.log.WithContext(ctx).Debugf("Fetching profile for user ID %d via gRPC", userID)
	resp, err := r.userClient.GetUser(ctx, &systemv1.GetUserRequest{Id: userID, WithProfile: true})
	if err != nil {
		return nil, err
	}
	user := resp.GetUser()
	if user == nil || user.Profile == nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "User profile not found")
	}
	return user.Profile, nil
}

// GetSetting retrieves user settings via gRPC.
func (r *meGRPCRepo) GetSetting(ctx context.Context, userID int64) (*dto.UserSettingPB, error) {
	r.log.WithContext(ctx).Debugf("Fetching settings for user ID %d via gRPC", userID)
	resp, err := r.userClient.GetUser(ctx, &systemv1.GetUserRequest{Id: userID, WithSetting: true})
	if err != nil {
		return nil, err
	}
	user := resp.GetUser()
	if user == nil || user.Setting == nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "User settings not found")
	}
	return user.Setting, nil
}

// UpdateProfile updates user's profile information via a gRPC call to UserService.
func (r *meGRPCRepo) UpdateProfile(ctx context.Context, userID int64, profileData *dto.UserProfilePB) error {
	r.log.WithContext(ctx).Debugf("Updating profile for user ID %d via gRPC", userID)
	_, err := r.userClient.UpdateUser(ctx, &systemv1.UpdateUserRequest{
		User: &dto.UserPB{
			Id:      userID,
			Profile: profileData,
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"profile.nickname", "profile.avatar", "profile.gender", "profile.name", "profile.department", "profile.remark"},
		},
	})
	return err
}

// UpdateSetting updates user's settings via gRPC.
func (r *meGRPCRepo) UpdateSetting(ctx context.Context, userID int64, settingsData *dto.UserSettingPB) error {
	r.log.WithContext(ctx).Debugf("Updating settings for user ID %d via gRPC", userID)
	_, err := r.userClient.UpdateUser(ctx, &systemv1.UpdateUserRequest{
		User: &dto.UserPB{
			Id:      userID,
			Setting: settingsData,
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"setting.theme", "setting.language", "setting.timezone", "setting.preferences"},
		},
	})
	return err
}

// UpdatePreferences updates user preferences via gRPC.
func (r *meGRPCRepo) UpdatePreferences(ctx context.Context, userID int64, preferences map[string]string) error {
	r.log.WithContext(ctx).Debugf("Updating preferences for user ID %d via gRPC", userID)

	// Get current settings first
	resp, err := r.userClient.GetUser(ctx, &systemv1.GetUserRequest{Id: userID, WithSetting: true})
	if err != nil {
		return err
	}
	user := resp.GetUser()
	if user == nil || user.Setting == nil {
		return errors.NotFound("USER_NOT_FOUND", "User not found")
	}

	// Merge preferences
	currentPreferences := user.Setting.Preferences
	if currentPreferences == nil {
		currentPreferences = make(map[string]string)
	}
	for k, v := range preferences {
		currentPreferences[k] = v
	}

	// Update settings
	_, err = r.userClient.UpdateUser(ctx, &systemv1.UpdateUserRequest{
		User: &dto.UserPB{
			Id: userID,
			Setting: &dto.UserSettingPB{
				Theme:       user.Setting.Theme,
				Language:    user.Setting.Language,
				Timezone:    user.Setting.Timezone,
				Preferences: currentPreferences,
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"setting.preferences"},
		},
	})
	return err
}

// UpdatePassword changes the current user's password via gRPC.
// Passes plain text passwords to system's ChangeUserPassword which handles verification and hashing.
func (r *meGRPCRepo) UpdatePassword(ctx context.Context, userID int64, hashedPassword string) error {
	r.log.WithContext(ctx).Debugf("Changing password for user ID %d via gRPC", userID)
	_, err := r.userClient.ChangeUserPassword(ctx, &systemv1.ChangeUserPasswordRequest{
		Id:       userID,
		Password: hashedPassword,
	})
	return err
}

var _ dto.MeRepo = (*meGRPCRepo)(nil)
