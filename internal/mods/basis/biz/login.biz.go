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

	pb "origadmin/application/admin/api/v1/services/basis"
	"origadmin/application/admin/internal/mods/basis/dto"
)

// LoginBiz is a Login use case.
type LoginBiz struct {
	dao     dto.LoginRepo
	limiter pagination.PageLimiter
	log     *log.Helper
}

func (biz LoginBiz) Refresh(ctx context.Context, in *pb.RefreshRequest, opts ...grpc.CallOption) (*pb.RefreshResponse, error) {
	return biz.dao.Refresh(ctx, in)
}

func (biz LoginBiz) CaptchaResource(ctx context.Context, in *pb.CaptchaResourceRequest, opts ...grpc.CallOption) (*pb.CaptchaResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz LoginBiz) CaptchaResources(ctx context.Context, in *pb.CaptchaResourcesRequest, opts ...grpc.CallOption) (*pb.CaptchaResourcesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz LoginBiz) CaptchaID(ctx context.Context, in *dto.CaptchaIDRequest, opts ...grpc.CallOption) (*dto.CaptchaIDResponse, error) {
	log.Info("CaptchaID")
	return biz.dao.CaptchaID(ctx, in)
}

func (biz LoginBiz) CaptchaImage(ctx context.Context, in *dto.CaptchaImageRequest, opts ...grpc.CallOption) (*dto.CaptchaImageResponse, error) {
	log.Info("CaptchaImage")
	return biz.dao.CaptchaImage(ctx, in.Id, in.Reload == "1" || in.Reload == "true")
}

func (biz LoginBiz) Login(ctx context.Context, in *dto.LoginRequest, opts ...grpc.CallOption) (*dto.LoginResponse, error) {
	log.Info("Login")
	return biz.dao.Login(ctx, in)
}

func (biz LoginBiz) Logout(ctx context.Context, in *dto.LogoutRequest, opts ...grpc.CallOption) (*dto.LogoutResponse, error) {
	log.Info("Logout")
	return biz.dao.Logout(ctx, in)
}

func (biz LoginBiz) CurrentUser(ctx context.Context, in *dto.CurrentUserRequest, opts ...grpc.CallOption) (*dto.CurrentUserResponse, error) {
	log.Info("CurrentUser")
	return biz.dao.CurrentUser(ctx, in)
}

func (biz LoginBiz) CurrentMenus(ctx context.Context, in *dto.CurrentMenusRequest, opts ...grpc.CallOption) (*dto.CurrentMenusResponse, error) {
	log.Info("CurrentMenus")
	return biz.dao.CurrentMenus(ctx, in)
}

// NewLoginBiz new a Login use case.
func NewLoginBiz(repo dto.LoginRepo, logger log.Logger) *LoginBiz {
	return &LoginBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewLoginClient new a Login use case.
func NewLoginClient(repo dto.LoginRepo, logger log.Logger) pb.LoginAPIClient {
	return &LoginBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.LoginAPIClient = (*LoginBiz)(nil)
