/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the auth module of OrigAdmin.
package biz

import (
	"context"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"

	"github.com/origadmin/runtime/log"

	"origadmin/application/admin/internal/helpers/repo"

	"github.com/origadmin/runtime"
	pb "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/internal/features/auth/dto"
)

// CasbinServiceBiz is a Casbin use case.
type CasbinServiceBiz struct {
	dao          dto.CasbinRepo
	limiter      repo.PageLimiter
	log          *log.Helper
	lastModified *atomic.Int64
}

func (c CasbinServiceBiz) StreamRules(request *pb.StreamRulesRequest, stream grpc.ServerStreamingServer[pb.StreamRulesResponse]) error {
	c.log.Debug("StreamRules")
	ctx := stream.Context()
	if request.WithPolicies {
		if err := c.streamPolicies(ctx, stream); err != nil {
			return err
		}
	}

	if request.WithGroupings {
		if err := c.streamGroupings(ctx, stream); err != nil {
			return err
		}
	}
	return nil
}

func (c CasbinServiceBiz) ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	c.log.Debug("ListPolicies")
	return c.dao.ListPolicies(ctx, in)
}

func (c CasbinServiceBiz) ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	c.log.Debug("ListGroupings")
	return c.dao.ListGroupings(ctx, in)
}

func (c CasbinServiceBiz) WatchUpdate(_ context.Context,
	in *pb.WatchUpdateRequest) (*pb.WatchUpdateResponse, error) {
	c.log.Debug("WatchUpdate")
	return &pb.WatchUpdateResponse{ModifiedDate: c.lastModified.Load()}, nil
}

func (c CasbinServiceBiz) UpdateRules() {
	// todo: load from db
	c.lastModified.Store(time.Now().Unix())
}

func (c CasbinServiceBiz) streamPolicies(ctx context.Context, stream grpc.ServerStreamingServer[pb.StreamRulesResponse]) error {
	policies, err := c.ListPolicies(ctx, &pb.ListPoliciesRequest{})
	if err != nil {
		return err
	}
	for _, rule := range policies.Rules {
		if err := stream.Send(newPolicyResponse(rule)); err != nil {
			return err
		}
	}
	return nil
}

func (c CasbinServiceBiz) streamGroupings(ctx context.Context, stream grpc.ServerStreamingServer[pb.StreamRulesResponse]) error {
	groupings, err := c.ListGroupings(ctx, &pb.ListGroupingsRequest{})
	if err != nil {
		return err
	}
	for _, rule := range groupings.Rules {
		if err := stream.Send(newGroupingResponse(rule)); err != nil {
			return err
		}
	}
	return nil
}

func newPolicyResponse(rule *pb.PolicyRule) *pb.StreamRulesResponse {
	return &pb.StreamRulesResponse{
		RuleType: &pb.StreamRulesResponse_Policy{Policy: rule},
	}
}

func newGroupingResponse(rule *pb.GroupingRule) *pb.StreamRulesResponse {
	return &pb.StreamRulesResponse{
		RuleType: &pb.StreamRulesResponse_Grouping{Grouping: rule},
	}
}

// NewCasbinServiceBiz new a Casbin use case.
func NewCasbinServiceBiz(r *runtime.App, repo dto.CasbinRepo) *CasbinServiceBiz {
	return &CasbinServiceBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(log.With(r.Logger(), "module", "biz/casbin")),
		lastModified: &atomic.Int64{}}
}
