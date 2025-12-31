/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime/context"
	"origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/gateway/client"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGatewayService)

// GatewayService is a gateway service.
type GatewayService struct {
	Auth   *client.AuthClientSet
	System *client.SystemClientSet
}

func (g GatewayService) GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) GetUserResources(ctx context.Context, request *auth.GetUserResourcesRequest) (*auth.GetUserResourcesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) GetUserRoles(ctx context.Context, request *auth.GetUserRolesRequest) (*auth.GetUserRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) UpdatePassword(ctx context.Context, request *auth.UpdatePasswordRequest) (*auth.UpdatePasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) UpdateProfile(ctx context.Context, request *auth.UpdateProfileRequest) (*auth.UpdateProfileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) GetCaptcha(ctx context.Context, request *auth.GetCaptchaRequest) (*auth.GetCaptchaResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) RefreshToken(ctx context.Context, request *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) CreateView(ctx context.Context, request *system.CreateViewRequest) (*system.CreateViewResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) DeleteView(ctx context.Context, request *system.DeleteViewRequest) (*system.DeleteViewResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) GetView(ctx context.Context, request *system.GetViewRequest) (*system.GetViewResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) ListViews(ctx context.Context, request *system.ListViewsRequest) (*system.ListViewsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) UpdateView(ctx context.Context, request *system.UpdateViewRequest) (*system.UpdateViewResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) CreateResource(ctx context.Context, request *system.CreateResourceRequest) (*system.CreateResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) DeleteResource(ctx context.Context, request *system.DeleteResourceRequest) (*system.DeleteResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) GetResource(ctx context.Context, request *system.GetResourceRequest) (*system.GetResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) ListResources(ctx context.Context, request *system.ListResourcesRequest) (*system.ListResourcesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) UpdateResource(ctx context.Context, request *system.UpdateResourceRequest) (*system.UpdateResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) CreatePermission(ctx context.Context, request *system.CreatePermissionRequest) (*system.CreatePermissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) DeletePermission(ctx context.Context, request *system.DeletePermissionRequest) (*system.DeletePermissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) GetPermission(ctx context.Context, request *system.GetPermissionRequest) (*system.GetPermissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) ListPermissions(ctx context.Context, request *system.ListPermissionsRequest) (*system.ListPermissionsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) UpdatePermission(ctx context.Context, request *system.UpdatePermissionRequest) (*system.UpdatePermissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) CreateRole(ctx context.Context, request *system.CreateRoleRequest) (*system.CreateRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) DeleteRole(ctx context.Context, request *system.DeleteRoleRequest) (*system.DeleteRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) GetRole(ctx context.Context, request *system.GetRoleRequest) (*system.GetRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) ListRoles(ctx context.Context, request *system.ListRolesRequest) (*system.ListRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) UpdateRole(ctx context.Context, request *system.UpdateRoleRequest) (*system.UpdateRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) CreateUser(ctx context.Context, request *system.CreateUserRequest) (*system.CreateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) DeleteUser(ctx context.Context, request *system.DeleteUserRequest) (*system.DeleteUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) GetUser(ctx context.Context, request *system.GetUserRequest) (*system.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) ListUserResources(ctx context.Context, request *system.ListUserResourcesRequest) (*system.ListUserResourcesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) ListUsers(ctx context.Context, request *system.ListUsersRequest) (*system.ListUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) ResetUserPassword(ctx context.Context, request *system.ResetUserPasswordRequest) (*system.ResetUserPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) UpdateUser(ctx context.Context, request *system.UpdateUserRequest) (*system.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) UpdateUserRoles(ctx context.Context, request *system.UpdateUserRolesRequest) (*system.UpdateUserRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GatewayService) UpdateUserStatus(ctx context.Context, request *system.UpdateUserStatusRequest) (*system.UpdateUserStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewGatewayService new a gateway service.
func NewGatewayService(authClient *client.AuthClientSet, systemClient *client.SystemClientSet) (*GatewayService, error) {
	return &GatewayService{
		Auth:   authClient,
		System: systemClient,
	}, nil
}
