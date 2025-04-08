/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"

	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// LoginServiceBiz is a Login use case.
type LoginServiceBiz struct {
	dao     dto.LoginRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz LoginServiceBiz) CaptchaId(ctx context.Context, in *pb.CaptchaIdRequest) (*pb.CaptchaIdResponse, error) {
	log.Info("CaptchaId")
	return biz.dao.CaptchaID(ctx, in)
}

func (biz LoginServiceBiz) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Info("Register")
	return biz.dao.Register(ctx, in)
}

func (biz LoginServiceBiz) Captcha(ctx context.Context, in *pb.CaptchaRequest) (*pb.CaptchaResponse, error) {
	log.Info("Captcha")
	return biz.dao.Captcha(ctx, in)
}

func (biz LoginServiceBiz) CaptchaImage(ctx context.Context, in *dto.CaptchaImageRequest) (*dto.CaptchaImageResponse, error) {
	log.Info("CaptchaImage")
	return biz.dao.CaptchaImage(ctx, in.Id, in.Reload == "1" || in.Reload == "true")
}

func (biz LoginServiceBiz) CaptchaAudio(ctx context.Context, in *pb.CaptchaAudioRequest) (*pb.CaptchaAudioResponse, error) {
	log.Info("CaptchaAudio")
	return biz.dao.CaptchaAudio(ctx, in.Id, in.Reload == "1" || in.Reload == "true")
}

func (biz LoginServiceBiz) Login(ctx context.Context, in *dto.LoginRequest) (*dto.LoginResponse, error) {
	log.Info("Login")
	return biz.dao.Login(ctx, in)
}

func (biz LoginServiceBiz) Logout(ctx context.Context, in *dto.LogoutRequest) (*dto.LogoutResponse, error) {
	log.Info("Logout")
	return biz.dao.Logout(ctx, in)
}

func (biz LoginServiceBiz) CurrentUser(ctx context.Context, in *dto.CurrentUserRequest) (*dto.CurrentUserResponse, error) {
	log.Info("CurrentUser")
	return biz.dao.CurrentUser(ctx, in)
}

func (biz LoginServiceBiz) TokenRefresh(ctx context.Context, in *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	log.Info("TokenRefresh")
	return biz.dao.TokenRefresh(ctx, in)
}

// NewLoginServiceBiz new a Login use case.
func NewLoginServiceBiz(repo dto.LoginRepo, logger log.KLogger) *LoginServiceBiz {
	return &LoginServiceBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}
