/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"golang.org/x/net/context"

	pb "origadmin/application/admin/api/v1/services/basis"
)

// LoginAPIService is a login service.
type LoginAPIService struct {
	pb.UnimplementedLoginAPIServer

	client pb.LoginAPIClient
}

func (obj LoginAPIService) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return obj.client.Register(ctx, request)
}

func (obj LoginAPIService) CaptchaResource(ctx context.Context, request *pb.CaptchaResourceRequest) (*pb.CaptchaResourceResponse, error) {
	return obj.client.CaptchaResource(ctx, request)
}

func (obj LoginAPIService) CaptchaResources(ctx context.Context, request *pb.CaptchaResourcesRequest) (*pb.CaptchaResourcesResponse, error) {
	return obj.client.CaptchaResources(ctx, request)
}

func (obj LoginAPIService) Refresh(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	return obj.client.Refresh(ctx, request)
}

func (obj LoginAPIService) CaptchaID(ctx context.Context, request *pb.CaptchaIDRequest) (*pb.CaptchaIDResponse, error) {
	return obj.client.CaptchaID(ctx, request)
}

func (obj LoginAPIService) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
	return obj.client.CaptchaImage(ctx, request)
}

func (obj LoginAPIService) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return obj.client.Login(ctx, request)
}

func (obj LoginAPIService) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return obj.client.Logout(ctx, request)
}

func (obj LoginAPIService) CurrentUser(ctx context.Context, request *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
	return obj.client.CurrentUser(ctx, request)
}

func (obj LoginAPIService) CurrentMenus(ctx context.Context, request *pb.CurrentMenusRequest) (*pb.CurrentMenusResponse, error) {
	return obj.client.CurrentMenus(ctx, request)
}

//func (l LoginAPIService) mustEmbedUnimplementedLoginAPIServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewLoginAPIService new a login service.
func NewLoginAPIService(client pb.LoginAPIClient) *LoginAPIService {
	return &LoginAPIService{client: client}
}

// NewLoginAPIServer new a login service.
func NewLoginAPIServer(client pb.LoginAPIClient) pb.LoginAPIServer {
	return &LoginAPIService{client: client}
}

var _ pb.LoginAPIServer = (*LoginAPIService)(nil)
