/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/auth"
)

// CasbinSourceServiceHTTPServer is a login service.
type CasbinSourceServiceHTTPServer struct {
	pb.UnimplementedCasbinSourceServiceServer
	client pb.CasbinSourceServiceHTTPClient
}

func (c CasbinSourceServiceHTTPServer) ListGroupings(ctx context.Context, request *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CasbinSourceServiceHTTPServer) ListPolicies(ctx context.Context, request *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CasbinSourceServiceHTTPServer) WatchUpdate(ctx context.Context, request *pb.WatchUpdateRequest) (*pb.WatchUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewCasbinServiceHTTPServer new a login service.
func NewCasbinServiceHTTPServer(client pb.CasbinSourceServiceHTTPClient) *CasbinSourceServiceHTTPServer {
	return &CasbinSourceServiceHTTPServer{client: client}
}

// NewCasbinSourceServiceHTTPServerPB new a login service.
func NewCasbinSourceServiceHTTPServerPB(client pb.CasbinSourceServiceHTTPClient) pb.CasbinSourceServiceHTTPServer {
	return &CasbinSourceServiceHTTPServer{client: client}
}

var _ pb.CasbinSourceServiceHTTPServer = (*CasbinSourceServiceHTTPServer)(nil)
