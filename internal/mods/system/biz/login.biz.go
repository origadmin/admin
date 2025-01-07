/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"

	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// LoginServiceClientBiz is a Login use case.
type LoginServiceClientBiz struct {
	dao     dto.LoginRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz LoginServiceClientBiz) CaptchaId(ctx context.Context, in *pb.CaptchaIdRequest, opts ...grpc.CallOption) (*pb.CaptchaIdResponse, error) {
	log.Info("CaptchaId")
	return biz.dao.CaptchaID(ctx, in)
}

func (biz LoginServiceClientBiz) Register(ctx context.Context, in *pb.RegisterRequest, opts ...grpc.CallOption) (*pb.RegisterResponse, error) {
	log.Info("Register")
	return biz.dao.Register(ctx, in)
}

func (biz LoginServiceClientBiz) CaptchaResource(ctx context.Context, in *pb.CaptchaResourceRequest, opts ...grpc.CallOption) (*pb.CaptchaResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz LoginServiceClientBiz) CaptchaResources(ctx context.Context, in *pb.CaptchaResourcesRequest, opts ...grpc.CallOption) (*pb.CaptchaResourcesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz LoginServiceClientBiz) CaptchaImage(ctx context.Context, in *dto.CaptchaImageRequest, opts ...grpc.CallOption) (*dto.CaptchaImageResponse, error) {
	log.Info("CaptchaImage")
	return biz.dao.CaptchaImage(ctx, in.Id, in.Reload == "1" || in.Reload == "true")
}

func (biz LoginServiceClientBiz) Login(ctx context.Context, in *dto.LoginRequest, opts ...grpc.CallOption) (*dto.LoginResponse, error) {
	log.Info("Login")
	return biz.dao.Login(ctx, in)
}

func (biz LoginServiceClientBiz) Logout(ctx context.Context, in *dto.LogoutRequest, opts ...grpc.CallOption) (*dto.LogoutResponse, error) {
	log.Info("Logout")
	return biz.dao.Logout(ctx, in)
}

func (biz LoginServiceClientBiz) CurrentUser(ctx context.Context, in *dto.CurrentUserRequest, opts ...grpc.CallOption) (*dto.CurrentUserResponse, error) {
	log.Info("CurrentUser")
	return biz.dao.CurrentUser(ctx, in)
}

func (biz LoginServiceClientBiz) TokenRefresh(ctx context.Context, in *pb.TokenRefreshRequest, opts ...grpc.CallOption) (*pb.TokenRefreshResponse, error) {
	log.Info("TokenRefresh")
	return biz.dao.TokenRefresh(ctx, in)
}

// NewLoginServiceClientBiz new a Login use case.
func NewLoginServiceClientBiz(repo dto.LoginRepo, logger log.KLogger) *LoginServiceClientBiz {
	return &LoginServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewLoginServiceClient new a Login use case.
func NewLoginServiceClient(repo dto.LoginRepo, logger log.KLogger) pb.LoginServiceClient {
	return &LoginServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.LoginServiceClient = (*LoginServiceClientBiz)(nil)
