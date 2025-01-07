/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type (
	CaptchaIDRequest     = pb.CaptchaIdRequest
	CaptchaIDResponse    = pb.CaptchaIdResponse
	CaptchaImageRequest  = pb.CaptchaImageRequest
	CaptchaImageResponse = pb.CaptchaImageResponse
	CaptchaData          = pb.CaptchaData
	LoginRequest         = pb.LoginRequest
	LoginResponse        = pb.LoginResponse
	LogoutRequest        = pb.LogoutRequest
	LogoutResponse       = pb.LogoutResponse
	TokenRefreshRequest  = pb.TokenRefreshRequest
	TokenRefreshResponse = pb.TokenRefreshResponse
	RegisterRequest      = pb.RegisterRequest
	RegisterResponse     = pb.RegisterResponse
	CurrentUserRequest   = pb.CurrentUserRequest
	CurrentUserResponse  = pb.CurrentUserResponse
)

type LoginRepo interface {
	CaptchaID(ctx context.Context, in *CaptchaIDRequest) (*CaptchaIDResponse, error)
	CaptchaImage(ctx context.Context, id string, reload bool) (*CaptchaImageResponse, error)
	CurrentUser(ctx context.Context, in *CurrentUserRequest) (*CurrentUserResponse, error)
	Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest) (*LogoutResponse, error)
	Register(ctx context.Context, in *RegisterRequest) (*RegisterResponse, error)
	TokenRefresh(ctx context.Context, in *TokenRefreshRequest) (*TokenRefreshResponse, error)
}
