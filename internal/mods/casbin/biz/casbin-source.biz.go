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

	pb "origadmin/application/admin/api/v1/services/casbin"
	"origadmin/application/admin/internal/mods/casbin/dto"
)

// CasbinSourceServiceClientBiz is a CasbinSource use case.
type CasbinSourceServiceClientBiz struct {
	dao     dto.CasbinSourceRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (c CasbinSourceServiceClientBiz) ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest, opts ...grpc.CallOption) (*pb.ListPoliciesResponse, error) {
	return c.dao.ListPolicies(ctx, in)
}

func (c CasbinSourceServiceClientBiz) ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest, opts ...grpc.CallOption) (*pb.ListGroupingsResponse, error) {
	return c.dao.ListGroupings(ctx, in)
}

func (c CasbinSourceServiceClientBiz) StreamRules(ctx context.Context, in *pb.StreamRulesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.StreamRulesResponse], error) {
	return c.dao.StreamRules(ctx, in)
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
