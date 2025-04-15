/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// CasbinServiceHTTPServer is a login service.
type CasbinServiceHTTPServer struct {
	pb.UnimplementedCasbinSourceServiceServer

	client pb.CasbinSourceServiceHTTPClient
}

func (c CasbinServiceHTTPServer) ListGroupings(ctx context.Context, request *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CasbinServiceHTTPServer) ListPolicies(ctx context.Context, request *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CasbinServiceHTTPServer) WatchUpdate(ctx context.Context, request *pb.WatchUpdateRequest) (*pb.WatchUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewCasbinServiceHTTPServer new a login service.
func NewCasbinServiceHTTPServer(client pb.CasbinSourceServiceHTTPClient) *CasbinServiceHTTPServer {
	return &CasbinServiceHTTPServer{client: client}
}

// NewCasbinServiceHTTPServerPB new a login service.
func NewCasbinServiceHTTPServerPB(client pb.CasbinSourceServiceHTTPClient) pb.CasbinSourceServiceServer {
	return &CasbinServiceHTTPServer{client: client}
}

var _ pb.CasbinSourceServiceHTTPServer = (*CasbinServiceHTTPServer)(nil)
