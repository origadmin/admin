/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the moduls.enforcer.
package service

import (
	"context"

	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/biz"
)

type CasbinSourceServiceServer struct {
	pb.UnimplementedCasbinSourceServiceServer

	client   *biz.CasbinSourceServiceBiz
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
	stream grpc.ServerStreamingServer[pb.StreamRulesResponse]) error {
	return c.client.StreamRules(request, stream)
}

// NewCasbinSourceServiceServerPB new a menu service.
func NewCasbinSourceServiceServerPB(client *biz.CasbinSourceServiceBiz) pb.CasbinSourceServiceServer {
	return &CasbinSourceServiceServer{client: client}
}

var _ pb.CasbinSourceServiceServer = (*CasbinSourceServiceServer)(nil)
