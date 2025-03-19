/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// AuthServiceBiz is a Auth use case.
type AuthServiceBiz struct {
	dao     dto.AuthRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz AuthServiceBiz) AuthLogout(ctx context.Context, in *pb.AuthLogoutRequest) (*pb.AuthLogoutResponse, error) {
	return biz.dao.AuthLogout(ctx, in)
}

func (biz AuthServiceBiz) CreateToken(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	return biz.dao.CreateToken(ctx, in)
}

func (biz AuthServiceBiz) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	return biz.dao.ValidateToken(ctx, in)
}

func (biz AuthServiceBiz) DestroyToken(ctx context.Context, in *pb.DestroyTokenRequest) (*pb.DestroyTokenResponse, error) {
	return biz.dao.DestroyToken(ctx, in)
}

func (biz AuthServiceBiz) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	return biz.dao.Authenticate(ctx, in)
}

func (biz AuthServiceBiz) ListAuthResources(ctx context.Context, in *pb.ListAuthResourcesRequest) (*pb.ListAuthResourcesResponse, error) {
	var option dto.AuthResourceQueryOption
	if err := option.FromListRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("ListAuths")
	result, total, err := biz.dao.ListAuthResources(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ToListAuthResourcesResponse(result, in, total)
}

// NewAuthServiceBiz new a Auth use case.
func NewAuthServiceBiz(repo dto.AuthRepo, logger log.KLogger) *AuthServiceBiz {
	return &AuthServiceBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}
