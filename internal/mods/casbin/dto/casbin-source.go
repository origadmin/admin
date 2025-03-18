/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

import (
	"context"

	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/casbin"
)

type CasbinSourceRepo interface {
	ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error)
	ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error)
	StreamRules(ctx context.Context, in *pb.StreamRulesRequest) (grpc.ServerStreamingClient[pb.StreamRulesResponse], error)
}
