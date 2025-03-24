/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// CasbinSourceServiceBiz is a CasbinSource use case.
type CasbinSourceServiceBiz struct {
	dao          dto.CasbinSourceRepo
	limiter      pagination.PageLimiter
	log          *log.KHelper
	lastModified *atomic.Int64
}

func (c CasbinSourceServiceBiz) StreamRules(request *pb.StreamRulesRequest, stream grpc.ServerStreamingServer[pb.StreamRulesResponse]) error {
	log.Info("StreamRules")
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

func (c CasbinSourceServiceBiz) ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	log.Info("ListPolicies")
	return c.dao.ListPolicies(ctx, in)
}

func (c CasbinSourceServiceBiz) ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	log.Info("ListGroupings")
	return c.dao.ListGroupings(ctx, in)
}

func (c CasbinSourceServiceBiz) WatchUpdate(_ context.Context,
	request *pb.WatchUpdateRequest) (*pb.WatchUpdateResponse, error) {
	log.Info("WatchUpdate")
	ModifiedDate := request.LastModified
	if c.lastModified.Load() > ModifiedDate {
		ModifiedDate = c.lastModified.Load()
	}
	return &pb.WatchUpdateResponse{ModifiedDate: ModifiedDate}, nil
}

func (c CasbinSourceServiceBiz) UpdateRules() {
	c.lastModified.Store(time.Now().Unix())
}

func (c CasbinSourceServiceBiz) streamPolicies(ctx context.Context, stream grpc.ServerStreamingServer[pb.StreamRulesResponse]) error {
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

func (c CasbinSourceServiceBiz) streamGroupings(ctx context.Context, stream grpc.ServerStreamingServer[pb.StreamRulesResponse]) error {
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

// NewCasbinSourceServiceBiz new a CasbinSource use case.
func NewCasbinSourceServiceBiz(repo dto.CasbinSourceRepo, logger log.KLogger) *CasbinSourceServiceBiz {
	return &CasbinSourceServiceBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger),
		lastModified: &atomic.Int64{}}
}
