/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "origadmin/application/admin/api/v1/services/identity"
	"origadmin/application/admin/internal/features/identity/biz"
)

// MeService is the service for the /me endpoints.
type MeService struct {
	pb.UnimplementedMeServiceServer
	meUseCase *biz.MeUseCase
	log       *log.Helper
}

// NewMeService creates a new MeService.
func NewMeService(meUseCase *biz.MeUseCase, logger log.Logger) *MeService {
	return &MeService{
		meUseCase: meUseCase,
		log:       log.NewHelper(log.With(logger, "module", "service.me")),
	}
}

// GetProfile gets the current user's profile.
func (s *MeService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	profile, err := s.meUseCase.GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileResponse{Profile: profile}, nil
}

// UpdateProfile updates the current user's profile.
func (s *MeService) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	err := s.meUseCase.UpdateProfile(ctx, req.GetProfile())
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProfileResponse{}, nil
}

// UpdatePassword changes the current user's password.
func (s *MeService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	err := s.meUseCase.ChangePassword(ctx, req.GetOldPassword(), req.GetNewPassword())
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePasswordResponse{}, nil
}

// UpdatePreferences updates the current user's preferences (P2).
func (s *MeService) UpdatePreferences(ctx context.Context, req *pb.UpdatePreferencesRequest) (*pb.UpdatePreferencesResponse, error) {
	err := s.meUseCase.UpdatePreferences(ctx, req.GetPreferences())
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePreferencesResponse{}, nil
}

// GetUserSettings retrieves the current user's settings (P2).
func (s *MeService) GetUserSettings(ctx context.Context, req *pb.GetUserSettingRequest) (*pb.GetUserSettingResponse, error) {
	settings, err := s.meUseCase.GetSettings(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserSettingResponse{Setting: settings}, nil
}
