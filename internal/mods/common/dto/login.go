/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/common"
)

type (
	CaptchaIDRequest     = pb.CaptchaIDRequest
	CaptchaIDResponse    = pb.CaptchaIDResponse
	CaptchaImageRequest  = pb.CaptchaImageRequest
	CaptchaImageResponse = pb.CaptchaImageResponse
	CaptchaData          = pb.CaptchaData
	LoginRequest         = pb.LoginRequest
	LoginResponse        = pb.LoginResponse
	LogoutRequest        = pb.LogoutRequest
	LogoutResponse       = pb.LogoutResponse
	CurrentMenusRequest  = pb.CurrentMenusRequest
	CurrentMenusResponse = pb.CurrentMenusResponse
	CurrentUserRequest   = pb.CurrentUserRequest
	CurrentUserResponse  = pb.CurrentUserResponse
)

type LoginRepo interface {
	Login(ctx context.Context, username, password string) (*LoginResponse, error)
	CaptchaID(ctx context.Context, in *CaptchaIDRequest) (*CaptchaIDResponse, error)
	CaptchaImage(ctx context.Context, in *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error)
	CurrentMenus(ctx context.Context, in *pb.CurrentMenusRequest) (*pb.CurrentMenusResponse, error)
	CurrentUser(ctx context.Context, in *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error)
	Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error)
}
