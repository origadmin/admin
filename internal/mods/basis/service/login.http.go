/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/basis"
)

// LoginAPIHTTPService is a login service.
type LoginAPIHTTPService struct {
	pb.UnimplementedLoginAPIServer

	client pb.LoginAPIHTTPClient
}

func (obj LoginAPIHTTPService) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return obj.client.Register(ctx, request)
}

func (obj LoginAPIHTTPService) mustEmbedUnimplementedLoginAPIServer() {
	//TODO implement me
	panic("implement me")
}

func (obj LoginAPIHTTPService) CaptchaResource(ctx context.Context, request *pb.CaptchaResourceRequest) (*pb.CaptchaResourceResponse, error) {
	return obj.client.CaptchaResource(ctx, request)
}

func (obj LoginAPIHTTPService) CaptchaResources(ctx context.Context, request *pb.CaptchaResourcesRequest) (*pb.CaptchaResourcesResponse, error) {
	return obj.client.CaptchaResources(ctx, request)
}

func (obj LoginAPIHTTPService) Refresh(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	return obj.client.Refresh(ctx, request)
}

func (obj LoginAPIHTTPService) CaptchaID(ctx context.Context, request *pb.CaptchaIDRequest) (*pb.CaptchaIDResponse, error) {
	return obj.client.CaptchaID(ctx, request)
}

func (obj LoginAPIHTTPService) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
	return obj.client.CaptchaImage(ctx, request)
}

func (obj LoginAPIHTTPService) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return obj.client.Login(ctx, request)
}

func (obj LoginAPIHTTPService) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return obj.client.Logout(ctx, request)
}

func (obj LoginAPIHTTPService) CurrentUser(ctx context.Context, request *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
	return obj.client.CurrentUser(ctx, request)
}

func (obj LoginAPIHTTPService) CurrentMenus(ctx context.Context, request *pb.CurrentMenusRequest) (*pb.CurrentMenusResponse, error) {
	return obj.client.CurrentMenus(ctx, request)
}

//func (l LoginAPIHTTPService) mustEmbedUnimplementedLoginAPIServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewLoginAPIHTTPService new a login service.
func NewLoginAPIHTTPService(client pb.LoginAPIHTTPClient) *LoginAPIHTTPService {
	return &LoginAPIHTTPService{client: client}
}

// NewLoginHTTPServer new a login service.
func NewLoginHTTPServer(client pb.LoginAPIHTTPClient) pb.LoginAPIServer {
	return &LoginAPIHTTPService{client: client}
}

var _ pb.LoginAPIServer = (*LoginAPIHTTPService)(nil)
