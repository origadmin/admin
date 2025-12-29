package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	v1 "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/auth/biz"
)

// MeService is a service for the currently authenticated user.
type MeService struct {
	v1.UnimplementedMeServer
	uc *biz.MeUseCase
}

// NewMeService creates a new Me service.
func NewMeService(uc *biz.MeUseCase) *MeService {
	return &MeService{uc: uc}
}

// GetProfile retrieves the profile of the currently authenticated user.
func (s *MeService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*types.User, error) {
	// TODO: Get userID from context
	userID := int64(1) // Placeholder
	return s.uc.GetProfile(ctx, userID)
}

// UpdateProfile updates the profile of the currently authenticated user.
func (s *MeService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// UpdatePassword changes the password for the currently authenticated user.
func (s *MeService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// GetUserResources retrieves the menu/resource list for the current user.
func (s *MeService) GetUserResources(ctx context.Context, req *v1.GetUserResourcesRequest) (*v1.GetUserResourcesResponse, error) {
	return &v1.GetUserResourcesResponse{}, nil
}

// GetUserRoles retrieves the role list for the current user.
func (s *MeService) GetUserRoles(ctx context.Context, req *v1.GetUserRolesRequest) (*v1.GetUserRolesResponse, error) {
	return &v1.GetUserRolesResponse{}, nil
}
