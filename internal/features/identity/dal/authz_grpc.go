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

// GetPermissions retrieves all permissions for the current user via gRPC.
func (r *authzGRPCRepo) GetPermissions(ctx context.Context, userID int64) (dto.PermissionsPB, error) {
	r.log.WithContext(ctx).Debugf("Fetching permissions for user ID %d via gRPC", userID)

	// Get user with roles
	resp, err := r.userClient.GetUser(ctx, &systemv1.GetUserRequest{Id: userID, WithRoles: true})
	if err != nil {
		return nil, err
	}

	user := resp.GetUser()
	if user == nil || len(user.Roles) == 0 {
		return nil, nil
	}

	// Get each role with permissions
	permissionSet := make(map[string]*dto.PermissionPB)
	for _, role := range user.Roles {
		roleResp, err := r.roleClient.GetRole(ctx, &systemv1.GetRoleRequest{Id: role.Id})
		if err != nil {
			return nil, err
		}
		fullRole := roleResp.GetRole()
		if fullRole == nil {
			continue
		}
		// Collect permissions from role
		for _, permission := range fullRole.Permissions {
			if permission != nil && permission.Keyword != "" {
				permissionSet[permission.Keyword] = permission
			}
		}
	}

	permissions := make(dto.PermissionsPB, 0, len(permissionSet))
	for _, p := range permissionSet {
		permissions = append(permissions, p)
	}

	return permissions, nil
}

// GetRoles retrieves all roles for the current user via gRPC.
func (r *authzGRPCRepo) GetRoles(ctx context.Context, userID int64) (dto.RolesPB, error) {
	r.log.WithContext(ctx).Debugf("Fetching roles for user ID %d via gRPC", userID)

	resp, err := r.userClient.GetUser(ctx, &systemv1.GetUserRequest{Id: userID, WithRoles: true})
	if err != nil {
		return nil, err
	}

	user := resp.GetUser()
	if user == nil {
		return nil, nil
	}

	return user.Roles, nil
}

// GetViews retrieves all views for the current user via gRPC.
func (r *authzGRPCRepo) GetViews(ctx context.Context, userID int64) (dto.ViewsPB, error) {
	r.log.WithContext(ctx).Debugf("Fetching views for user ID %d via gRPC", userID)

	// Get permissions and extract views
	permissions, err := r.GetPermissions(ctx, userID)
	if err != nil {
		return nil, err
	}

	viewSet := make(map[string]*dto.ViewPB)
	for _, permission := range permissions {
		for _, view := range permission.Views {
			if view != nil && view.Keyword != "" {
				viewSet[view.Keyword] = view
			}
		}
	}

	views := make(dto.ViewsPB, 0, len(viewSet))
	for _, v := range viewSet {
		views = append(views, v)
	}

	return views, nil
}

// GetResources retrieves all resources for the current user via gRPC.
func (r *authzGRPCRepo) GetResources(ctx context.Context, userID int64) (dto.ResourcesPB, error) {
	r.log.WithContext(ctx).Debugf("Fetching resources for user ID %d via gRPC", userID)

	// Get permissions and extract resources
	permissions, err := r.GetPermissions(ctx, userID)
	if err != nil {
		return nil, err
	}

	resourceSet := make(map[string]*dto.ResourcePB)
	for _, permission := range permissions {
		for _, resource := range permission.Resources {
			if resource != nil && resource.Keyword != "" {
				resourceSet[resource.Keyword] = resource
			}
		}
	}

	resources := make(dto.ResourcesPB, 0, len(resourceSet))
	for _, r := range resourceSet {
		resources = append(resources, r)
	}

	return resources, nil
}

// GetPermissionKeywordsByUserID retrieves all permission keywords for the current user via gRPC.
func (r *authzGRPCRepo) GetPermissionKeywordsByUserID(ctx context.Context, userID int64) ([]string, error) {
	r.log.WithContext(ctx).Debugf("Fetching permission keywords for user ID %d via gRPC", userID)

	permissions, err := r.GetPermissions(ctx, userID)
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
