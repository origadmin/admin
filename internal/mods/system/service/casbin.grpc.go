/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the moduls.enforcer.
package service

import (
	"context"
	"io"

	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "origadmin/application/admin/api/v1/services/system"
)

type CasbinSourceServiceServer struct {
	pb.UnimplementedCasbinSourceServiceServer

	client   pb.CasbinSourceServiceClient
	enforcer *casbin.Enforcer
}

func (c *CasbinSourceServiceServer) ListPolicies(ctx context.Context,
	request *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	return c.client.ListPolicies(ctx, request)
}

func (c *CasbinSourceServiceServer) ListGroupings(ctx context.Context,
	request *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	return c.client.ListGroupings(ctx, request)
}

func (c *CasbinSourceServiceServer) StreamRules(request *pb.StreamRulesRequest,
	g grpc.ServerStreamingServer[pb.StreamRulesResponse]) error {
	stream, err := c.client.StreamRules(g.Context(), request)
	if err != nil {
		return status.Errorf(codes.Unavailable, "connect server failed: %v", err)
	}

	for {
		rule, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			st, _ := status.FromError(err)
			return status.Errorf(st.Code(), "recvied error: %v", st.Message())
		}
		if err := g.Send(rule); err != nil {
			if status.Code(err) == codes.Canceled {
				_ = stream.CloseSend()
				return nil
			}
			return status.Errorf(codes.Internal, "send data error: %v", err)
		}
		if c.enforcer != nil {
			switch v := rule.RuleType.(type) {
			case *pb.StreamRulesResponse_Policy:
				_, _ = c.enforcer.AddPolicy(v.Policy.Params)
			case *pb.StreamRulesResponse_Grouping:
				_, _ = c.enforcer.AddGroupingPolicy(v.Grouping.Params)
			}
		}
	}
	return nil
}

// NewCasbinSourceServiceServer new a menu service.
func NewCasbinSourceServiceServer(client pb.CasbinSourceServiceClient) *CasbinSourceServiceServer {
	return &CasbinSourceServiceServer{client: client}
}

// NewCasbinSourceServiceServerPB new a menu service.
func NewCasbinSourceServiceServerPB(client pb.CasbinSourceServiceClient) pb.CasbinSourceServiceServer {
	return &CasbinSourceServiceServer{client: client}
}

var _ pb.CasbinSourceServiceServer = (*CasbinSourceServiceServer)(nil)
