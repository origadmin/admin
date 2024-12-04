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

	pb "origadmin/application/admin/api/v1/services/common"
	"origadmin/application/admin/internal/mods/system/dto"
)

// LoginBiz is a Login use case.
type LoginBiz struct {
	dao     dto.LoginRepo
	limiter pagination.PageLimiter
	log     *log.Helper
}

func (biz LoginBiz) CaptchaID(ctx context.Context, in *pb.CaptchaIDRequest, opts ...grpc.CallOption) (*pb.CaptchaIDResponse, error) {
	log.Info("CaptchaID")
	return biz.dao.CaptchaID(ctx, in)
}

func (biz LoginBiz) CaptchaImage(ctx context.Context, in *pb.CaptchaImageRequest, opts ...grpc.CallOption) (*pb.CaptchaImageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz LoginBiz) Login(ctx context.Context, in *pb.LoginRequest, opts ...grpc.CallOption) (*pb.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz LoginBiz) Logout(ctx context.Context, in *pb.LogoutRequest, opts ...grpc.CallOption) (*pb.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz LoginBiz) CurrentUser(ctx context.Context, in *pb.CurrentUserRequest, opts ...grpc.CallOption) (*pb.CurrentUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz LoginBiz) CurrentMenus(ctx context.Context, in *pb.CurrentMenusRequest, opts ...grpc.CallOption) (*pb.CurrentMenusResponse, error) {
	//TODO implement me
	panic("implement me")
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
