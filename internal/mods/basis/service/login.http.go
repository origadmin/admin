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

func (l LoginAPIHTTPService) CaptchaResource(ctx context.Context, request *pb.CaptchaResourceRequest) (*pb.CaptchaResourceResponse, error) {
	return l.client.CaptchaResource(ctx, request)
}

func (l LoginAPIHTTPService) CaptchaResources(ctx context.Context, request *pb.CaptchaResourcesRequest) (*pb.CaptchaResourcesResponse, error) {
	return l.client.CaptchaResources(ctx, request)
}

func (l LoginAPIHTTPService) Refresh(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	return l.client.Refresh(ctx, request)
}

func (l LoginAPIHTTPService) CaptchaID(ctx context.Context, request *pb.CaptchaIDRequest) (*pb.CaptchaIDResponse, error) {
	return l.client.CaptchaID(ctx, request)
}

func (l LoginAPIHTTPService) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
	return l.client.CaptchaImage(ctx, request)
}

func (l LoginAPIHTTPService) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return l.client.Login(ctx, request)
}

func (l LoginAPIHTTPService) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return l.client.Logout(ctx, request)
}

func (l LoginAPIHTTPService) CurrentUser(ctx context.Context, request *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
	return l.client.CurrentUser(ctx, request)
}

func (l LoginAPIHTTPService) CurrentMenus(ctx context.Context, request *pb.CurrentMenusRequest) (*pb.CurrentMenusResponse, error) {
	return l.client.CurrentMenus(ctx, request)
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
