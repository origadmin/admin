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

// AuthsBiz is a Auth use case.
type AuthsBiz struct {
	dao     dto.AuthRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz AuthsBiz) CreateToken(ctx context.Context, in *pb.CreateTokenRequest, opts ...grpc.CallOption) (*pb.CreateTokenResponse, error) {
	return biz.dao.CreateToken(ctx, in)
}

func (biz AuthsBiz) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest, opts ...grpc.CallOption) (*pb.ValidateTokenResponse, error) {
	return biz.dao.ValidateToken(ctx, in)
}

func (biz AuthsBiz) DestroyToken(ctx context.Context, in *pb.DestroyTokenRequest, opts ...grpc.CallOption) (*pb.DestroyTokenResponse, error) {
	return biz.dao.DestroyToken(ctx, in)
}

func (biz AuthsBiz) Authenticate(ctx context.Context, in *pb.AuthenticateRequest, opts ...grpc.CallOption) (*pb.AuthenticateResponse, error) {
	return biz.dao.Authenticate(ctx, in)
}

func (biz AuthsBiz) ListAuthResources(ctx context.Context, in *pb.ListAuthResourcesRequest, opts ...grpc.CallOption) (*pb.ListAuthResourcesResponse, error) {
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

// NewAuthsBiz new a Auth use case.
func NewAuthsBiz(repo dto.AuthRepo, logger log.KLogger) *AuthsBiz {
	return &AuthsBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewAuthsClient new a Auth use case.
func NewAuthsClient(repo dto.AuthRepo, logger log.KLogger) pb.AuthAPIClient {
	return &AuthsBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.AuthAPIClient = (*AuthsBiz)(nil)
