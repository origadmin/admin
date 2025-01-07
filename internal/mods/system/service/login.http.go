/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// LoginServiceHTTPServer is a login service.
type LoginServiceHTTPServer struct {
	pb.UnimplementedLoginServiceServer

	client pb.LoginServiceHTTPClient
}

func (obj LoginServiceHTTPServer) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return obj.client.Logout(ctx, request)
}

func (obj LoginServiceHTTPServer) TokenRefresh(ctx context.Context, request *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	return obj.client.TokenRefresh(ctx, request)
}

func (obj LoginServiceHTTPServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return obj.client.Register(ctx, request)
}

func (obj LoginServiceHTTPServer) CaptchaResource(ctx context.Context, request *pb.CaptchaResourceRequest) (*pb.CaptchaResourceResponse, error) {
	return obj.client.CaptchaResource(ctx, request)
}

func (obj LoginServiceHTTPServer) CaptchaResources(ctx context.Context, request *pb.CaptchaResourcesRequest) (*pb.CaptchaResourcesResponse, error) {
	return obj.client.CaptchaResources(ctx, request)
}

func (obj LoginServiceHTTPServer) CaptchaId(ctx context.Context, request *pb.CaptchaIdRequest) (*pb.CaptchaIdResponse, error) {
	return obj.client.CaptchaId(ctx, request)
}

func (obj LoginServiceHTTPServer) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
	return obj.client.CaptchaImage(ctx, request)
}

func (obj LoginServiceHTTPServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return obj.client.Login(ctx, request)
}

//func (l LoginServiceHTTPServer) mustEmbedUnimplementedLoginServiceHTTPServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewLoginServiceHTTPServer new a login service.
func NewLoginServiceHTTPServer(client pb.LoginServiceHTTPClient) *LoginServiceHTTPServer {
	return &LoginServiceHTTPServer{client: client}
}

// NewLoginServiceHTTPServerPB new a login service.
func NewLoginServiceHTTPServerPB(client pb.LoginServiceHTTPClient) pb.LoginServiceHTTPServer {
	return &LoginServiceHTTPServer{client: client}
}

var _ pb.LoginServiceHTTPServer = (*LoginServiceHTTPServer)(nil)
