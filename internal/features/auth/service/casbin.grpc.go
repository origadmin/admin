/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the moduls.enforcer.
package service

import (
	"context"

	"github.com/origadmin/runtime/service"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/internal/features/auth/biz" // Corrected import path
)

type CasbinSourceServiceServer struct {
	pb.UnimplementedCasbinSourceServiceServer
	client *biz.CasbinSourceServiceBiz
}

func (c *CasbinSourceServiceServer) WatchUpdate(ctx context.Context, request *pb.WatchUpdateRequest) (*pb.WatchUpdateResponse, error) {
	return c.client.WatchUpdate(ctx, request)
}

func (c *CasbinSourceServiceServer) mustEmbedUnimplementedCasbinSourceServiceServer() {

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
	stream grpc.ServerStreamingServer[pb.StreamRulesResponse]) error {
	return c.client.StreamRules(request, stream)
}

// NewCasbinSourceServiceServerPB new a menu service.
func NewCasbinSourceServiceServerPB(client *biz.CasbinSourceServiceBiz) pb.CasbinSourceServiceServer {
	return &CasbinSourceServiceServer{client: client}
}

func NewCasbinSourceServiceClient(client *service.GRPCClient) pb.CasbinSourceServiceClient {
	return pb.NewCasbinSourceServiceClient(client)
}

var _ pb.CasbinSourceServiceServer = (*CasbinSourceServiceServer)(nil)
