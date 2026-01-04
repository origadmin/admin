package service

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"

	"github.com/origadmin/contrib/security/principal"
	v1 "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/internal/features/auth/biz"
)

// MeService is a service for the currently authenticated user.
type MeService struct {
	v1.UnimplementedMeServiceServer
	uc *biz.MeUseCase
}

// NewMeService creates a new Me service.
func NewMeService(uc *biz.MeUseCase) *MeService {
	return &MeService{uc: uc}
}

// GetProfile retrieves the profile of the currently authenticated user.
func (s *MeService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.GetProfileResponse, error) {
	// Get the principal from the context, which is populated by the auth middleware.
	p, ok := principal.FromContext(ctx)
	if !ok {
		return nil, errors.Unauthorized("UNAUTHORIZED", "Missing user principal in context")
	}

	// The principal's ID is a string, so it needs to be converted to an integer.
	userID, err := strconv.ParseInt(p.GetID(), 10, 64)
	if err != nil {
		return nil, errors.InternalServer("INVALID_PRINCIPAL_ID", "User ID in principal is not a valid integer")
	}

	user, err := s.uc.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &v1.GetProfileResponse{User: user}, nil
}

// UpdateProfile updates the profile of the currently authenticated user.
func (s *MeService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) (*v1.UpdateProfileResponse, error) {
	return &v1.UpdateProfileResponse{}, nil
}

// UpdatePassword changes the password for the currently authenticated user.
func (s *MeService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*v1.UpdatePasswordResponse, error) {
	return &v1.UpdatePasswordResponse{}, nil
}

// GetUserResources retrieves the menu/resource list for the current user.
func (s *MeService) GetUserResources(ctx context.Context, req *v1.GetUserResourcesRequest) (*v1.GetUserResourcesResponse, error) {
	return &v1.GetUserResourcesResponse{}, nil
}

// GetUserRoles retrieves the role list for the current user.
func (s *MeService) GetUserRoles(ctx context.Context, req *v1.GetUserRolesRequest) (*v1.GetUserRolesResponse, error) {
	return &v1.GetUserRolesResponse{}, nil
}
