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

func (l LoginAPIService) CaptchaResource(ctx context.Context, request *pb.CaptchaResourceRequest) (*pb.CaptchaResourceResponse, error) {
	return l.client.CaptchaResource(ctx, request)
}

func (l LoginAPIService) CaptchaResources(ctx context.Context, request *pb.CaptchaResourcesRequest) (*pb.CaptchaResourcesResponse, error) {
	return l.client.CaptchaResources(ctx, request)
}

func (l LoginAPIService) Refresh(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	return l.client.Refresh(ctx, request)
}

func (l LoginAPIService) CaptchaID(ctx context.Context, request *pb.CaptchaIDRequest) (*pb.CaptchaIDResponse, error) {
	return l.client.CaptchaID(ctx, request)
}

func (l LoginAPIService) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
	return l.client.CaptchaImage(ctx, request)
}

func (l LoginAPIService) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return l.client.Login(ctx, request)
}

func (l LoginAPIService) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return l.client.Logout(ctx, request)
}

func (l LoginAPIService) CurrentUser(ctx context.Context, request *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
	return l.client.CurrentUser(ctx, request)
}

func (l LoginAPIService) CurrentMenus(ctx context.Context, request *pb.CurrentMenusRequest) (*pb.CurrentMenusResponse, error) {
	return l.client.CurrentMenus(ctx, request)
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
