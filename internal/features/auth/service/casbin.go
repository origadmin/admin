/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	pb "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/contrib/security/authz/casbin"
	"origadmin/application/admin/internal/features/auth/biz" // Corrected import path
)

// CasbinSourceBiz is a Casbin rule source service.
type CasbinSourceBiz struct {
	client *biz.CasbinSourceServiceBiz
	log    *log.KHelper
}

func (c CasbinSourceBiz) ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	return c.client.ListPolicies(ctx, in)
}

func (c CasbinSourceBiz) ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	return c.client.ListGroupings(ctx, in)
}

func (c CasbinSourceBiz) WatchUpdate(ctx context.Context, in *pb.WatchUpdateRequest) (*pb.WatchUpdateResponse, error) {
	return c.client.WatchUpdate(ctx, in)
}

func (c CasbinSourceBiz) StreamRules(ctx context.Context, in *pb.StreamRulesRequest) (grpc.ServerStreamingClient[pb.StreamRulesResponse], error) {
	stream := biz.NewCasbinRuleStream(ctx, c.client)
	go func() {
		err := stream.Start(in)
		if err != nil {
			c.log.Error("stream error", "error", err)
		}
	}()
	return stream, nil
}

func NewCasbinSourceBiz(r runtime.Runtime, client *biz.CasbinSourceServiceBiz) casbin.RuleSource {
	return &CasbinSourceBiz{
		client: client,
		log:    log.NewHelper(r.WithLogger("module", "service/casbin")),
	}
}

// CasbinSourceClient is a Casbin rule source service.
type CasbinSourceClient struct {
	client pb.CasbinSourceServiceClient
	log    *log.KHelper
}

func (c CasbinSourceClient) ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	return c.client.ListPolicies(ctx, in)
}

func (c CasbinSourceClient) ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	return c.client.ListGroupings(ctx, in)
}

func (c CasbinSourceClient) WatchUpdate(ctx context.Context, in *pb.WatchUpdateRequest) (*pb.WatchUpdateResponse, error) {
	return c.client.WatchUpdate(ctx, in)
}

func (c CasbinSourceClient) StreamRules(ctx context.Context, in *pb.StreamRulesRequest) (grpc.ServerStreamingClient[pb.StreamRulesResponse], error) {
	return c.client.StreamRules(ctx, in)
}

// NewCasbinSourceClient new a menu service.
func NewCasbinSourceClient(r runtime.Runtime, clients map[string]*service.GRPCClient) casbin.RuleSource {
	ll := log.NewHelper(r.WithLogger("module", "service/casbin"))
	client, ok := clients["auth"]
	c := NewUnimplementedCasbinSource(r)
	if ok {
		c = pb.NewCasbinSourceServiceClient(client)
	}
	return &CasbinSourceClient{
		client: c,
		log:    ll,
	}
}

func NewUnimplementedCasbinSource(r runtime.Runtime) pb.CasbinSourceServiceClient {
	return UnimplementedCasbinSource{
		log: log.NewHelper(r.WithLogger("module", "service/casbin")),
	}
}

type UnimplementedCasbinSource struct {
	log *log.KHelper
}

func (u UnimplementedCasbinSource) ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest, opts ...grpc.CallOption) (*pb.ListPoliciesResponse, error) {
	u.log.Error("ListPolicies not implemented")
	return &pb.ListPoliciesResponse{}, nil
}

func (u UnimplementedCasbinSource) ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest, opts ...grpc.CallOption) (*pb.ListGroupingsResponse, error) {
	u.log.Error("ListGroupings not implemented")
	return &pb.ListGroupingsResponse{}, nil
}

func (u UnimplementedCasbinSource) WatchUpdate(ctx context.Context, in *pb.WatchUpdateRequest, opts ...grpc.CallOption) (*pb.WatchUpdateResponse, error) {
	u.log.Error("WatchUpdate not implemented")
	return &pb.WatchUpdateResponse{}, nil
}

func (u UnimplementedCasbinSource) StreamRules(ctx context.Context, in *pb.StreamRulesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pb.StreamRulesResponse], error) {
	u.log.Error("StreamRules not implemented")
	return nil, status.Error(400, "not implemented")
}
