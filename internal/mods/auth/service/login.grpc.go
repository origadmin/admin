/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"golang.org/x/net/context"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/biz"
)

// LoginServiceServer is a login service.
type LoginServiceServer struct {
	pb.UnimplementedLoginServiceServer

	client *biz.LoginServiceBiz
}

func (obj LoginServiceServer) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return obj.client.Logout(ctx, request)
}

func (obj LoginServiceServer) TokenRefresh(ctx context.Context, request *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	return obj.client.TokenRefresh(ctx, request)
}

func (obj LoginServiceServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return obj.client.Register(ctx, request)
}

func (obj LoginServiceServer) Captcha(ctx context.Context, request *pb.CaptchaRequest) (*pb.CaptchaResponse, error) {
	return obj.client.Captcha(ctx, request)
}

func (obj LoginServiceServer) CaptchaId(ctx context.Context, request *pb.CaptchaIdRequest) (*pb.CaptchaIdResponse, error) {
	return obj.client.CaptchaId(ctx, request)
}

func (obj LoginServiceServer) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
	return obj.client.CaptchaImage(ctx, request)
}

func (obj LoginServiceServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return obj.client.Login(ctx, request)
}

//func (l LoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewLoginServiceServer new a login service.
func NewLoginServiceServer(client *biz.LoginServiceBiz) *LoginServiceServer {
	return &LoginServiceServer{client: client}
}

// NewLoginServiceServerPB new a login service.
func NewLoginServiceServerPB(client *biz.LoginServiceBiz) pb.LoginServiceServer {
	return &LoginServiceServer{client: client}
}

var _ pb.LoginServiceServer = (*LoginServiceServer)(nil)
