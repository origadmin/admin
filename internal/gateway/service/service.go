/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"errors"

	"github.com/google/wire"

	"origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/api/v1/services/gateway"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGatewayService)

// GatewayService implements the GatewayAPI service.
type GatewayService struct {
	gateway.UnimplementedGatewayServiceServer

	authClient   auth.AuthServiceClient
	systemClient system.UserServiceClient
}

// NewGatewayService creates a new gateway service.
func NewGatewayService(authClient auth.AuthServiceClient, systemClient system.UserServiceClient) *GatewayService {
	return &GatewayService{
		authClient:   authClient,
		systemClient: systemClient,
	}
}

func (g *GatewayService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	return g.authClient.Login(ctx, request)
}

func (g *GatewayService) GetCaptcha(ctx context.Context, request *auth.GetCaptchaRequest) (*auth.GetCaptchaResponse, error) {
	return g.authClient.GetCaptcha(ctx, request)
}

func (g *GatewayService) GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error) {
	// Assuming MeService is part of AuthService, if not, a new client for MeService is needed
	// This is a placeholder, as MeService is not directly available on AuthServiceClient
	// You might need to create a MeServiceClient
	return nil, errors.New("GetProfile not implemented on auth client")
}

func (g *GatewayService) ListUsers(ctx context.Context, request *system.ListUsersRequest) (*system.ListUsersResponse, error) {
	return g.systemClient.ListUsers(ctx, request)
}

func (g *GatewayService) GetUser(ctx context.Context, request *system.GetUserRequest) (*types.User, error) {
	res, err := g.systemClient.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.User, nil
}

func (g *GatewayService) CreateUser(ctx context.Context, request *system.CreateUserRequest) (*types.User, error) {
	res, err := g.systemClient.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.User, nil
}

func (g *GatewayService) UpdateUser(ctx context.Context, request *system.UpdateUserRequest) (*types.User, error) {
	res, err := g.systemClient.UpdateUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.User, nil
}

func (g *GatewayService) DeleteUser(ctx context.Context, request *system.DeleteUserRequest) (*system.DeleteUserResponse, error) {
	return g.systemClient.DeleteUser(ctx, request)
}
