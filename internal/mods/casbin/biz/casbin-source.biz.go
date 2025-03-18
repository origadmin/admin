/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// CasbinSourceServiceClientBiz is a CasbinSource use case.
type CasbinSourceServiceClientBiz struct {
	dao     dto.CasbinSourceRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz CasbinSourceServiceClientBiz) CasbinSourceLogout(ctx context.Context, in *pb.CasbinSourceLogoutRequest, opts ...grpc.CallOption) (*pb.CasbinSourceLogoutResponse, error) {
	return biz.dao.CasbinSourceLogout(ctx, in)
}

func (biz CasbinSourceServiceClientBiz) CreateToken(ctx context.Context, in *pb.CreateTokenRequest, opts ...grpc.CallOption) (*pb.CreateTokenResponse, error) {
	return biz.dao.CreateToken(ctx, in)
}

func (biz CasbinSourceServiceClientBiz) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest, opts ...grpc.CallOption) (*pb.ValidateTokenResponse, error) {
	return biz.dao.ValidateToken(ctx, in)
}

func (biz CasbinSourceServiceClientBiz) DestroyToken(ctx context.Context, in *pb.DestroyTokenRequest, opts ...grpc.CallOption) (*pb.DestroyTokenResponse, error) {
	return biz.dao.DestroyToken(ctx, in)
}

func (biz CasbinSourceServiceClientBiz) CasbinSourceenticate(ctx context.Context, in *pb.CasbinSourceenticateRequest, opts ...grpc.CallOption) (*pb.CasbinSourceenticateResponse, error) {
	return biz.dao.CasbinSourceenticate(ctx, in)
}

func (biz CasbinSourceServiceClientBiz) ListCasbinSourceResources(ctx context.Context, in *pb.ListCasbinSourceResourcesRequest, opts ...grpc.CallOption) (*pb.ListCasbinSourceResourcesResponse, error) {
	var option dto.CasbinSourceResourceQueryOption
	if err := option.FromListRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("ListCasbinSources")
	result, total, err := biz.dao.ListCasbinSourceResources(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ToListCasbinSourceResourcesResponse(result, in, total)
}

// NewCasbinSourceServiceClientBiz new a CasbinSource use case.
func NewCasbinSourceServiceClientBiz(repo dto.CasbinSourceRepo, logger log.KLogger) *CasbinSourceServiceClientBiz {
	return &CasbinSourceServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewCasbinSourceServiceClient new a CasbinSource use case.
func NewCasbinSourceServiceClient(repo dto.CasbinSourceRepo, logger log.KLogger) pb.CasbinSourceServiceClient {
	return &CasbinSourceServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.CasbinSourceServiceClient = (*CasbinSourceServiceClientBiz)(nil)
