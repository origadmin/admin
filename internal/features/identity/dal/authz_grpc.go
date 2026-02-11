/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/identity/dto"
)

// authzGRPCRepo implements dto.AuthzRepo interface using gRPC calls to system service.
type authzGRPCRepo struct {
	userClient     systemv1.UserServiceClient
	roleClient     systemv1.RoleServiceClient
	permClient     systemv1.PermissionServiceClient
	viewClient     systemv1.ViewServiceClient
	resourceClient systemv1.ResourceServiceClient
	log            *log.Helper
}

// NewAuthzGRPCRepo creates a new AuthzRepo that communicates with system service via gRPC.
func NewAuthzGRPCRepo(
	userClient systemv1.UserServiceClient,
	roleClient systemv1.RoleServiceClient,
	permClient systemv1.PermissionServiceClient,
	viewClient systemv1.ViewServiceClient,
	resourceClient systemv1.ResourceServiceClient,
	logger log.Logger,
) dto.AuthzRepo {
	return &authzGRPCRepo{
		userClient:     userClient,
		roleClient:     roleClient,
		permClient:     permClient,
		viewClient:     viewClient,
		resourceClient: resourceClient,
		log:            log.NewHelper(log.With(logger, "module", "dal.authz_grpc")),
	}
}

func (r *authzGRPCRepo) ListMyPermissions(ctx context.Context, userID int64) (dto.PermissionsPB, error) {
	resp, err := r.userClient.ListUserPermissions(ctx, &systemv1.ListUserPermissionsRequest{Id: userID})
	if err != nil {
		return nil, err
	}
	return resp.GetPermissions(), nil
}

func (r *authzGRPCRepo) ListMyRoles(ctx context.Context, userID int64) (dto.RolesPB, error) {
	resp, err := r.userClient.ListUserRoles(ctx, &systemv1.ListUserRolesRequest{Id: userID})
	if err != nil {
		return nil, err
	}
	return resp.GetRoles(), nil
}

func (r *authzGRPCRepo) ListMyViews(ctx context.Context, userID int64) (dto.ViewsPB, error) {
	resp, err := r.userClient.ListUserViews(ctx, &systemv1.ListUserViewsRequest{Id: userID})
	if err != nil {
		return nil, err
	}
	return resp.GetViews(), nil
}

func (r *authzGRPCRepo) ListMyResources(ctx context.Context, userID int64) (dto.ResourcesPB, error) {
	resp, err := r.userClient.ListUserResources(ctx, &systemv1.ListUserResourcesRequest{Id: userID})
	if err != nil {
		return nil, err
	}
	return resp.GetResources(), nil
}

// GetPermissionKeywordsByUserID retrieves all permission keywords for the current user via gRPC.
func (r *authzGRPCRepo) GetPermissionKeywordsByUserID(ctx context.Context, userID int64) ([]string, error) {
	r.log.WithContext(ctx).Debugf("Fetching permission keywords for user ID %d via gRPC", userID)

	permissions, err := r.ListMyPermissions(ctx, userID)
	if err != nil {
		return nil, err
	}

	keywords := make([]string, 0, len(permissions))
	for _, p := range permissions {
		if p.Keyword != "" {
			keywords = append(keywords, p.Keyword)
		}
	}

	return keywords, nil
}

// HasSystemRole checks if the current user has a role with the 'system' type via gRPC.
func (r *authzGRPCRepo) HasSystemRole(ctx context.Context, userID int64) (bool, error) {
	r.log.WithContext(ctx).Debugf("Checking system role for user ID %d via gRPC", userID)

	resp, err := r.userClient.GetUser(ctx, &systemv1.GetUserRequest{Id: userID, WithRoles: true})
	if err != nil {
		return false, err
	}

	user := resp.GetUser()
	if user == nil {
		return false, nil
	}

	for _, role := range user.Roles {
		if role.Type == 1 { // RoleTypeSystem = 1
			return true, nil
		}
	}

	return false, nil
}

var _ dto.AuthzRepo = (*authzGRPCRepo)(nil)
