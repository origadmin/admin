/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// AuthServiceHTTPServer is a menu service.
type AuthServiceHTTPServer struct {
	pb.UnimplementedAuthServiceServer

	client pb.AuthServiceHTTPClient
}

func (s AuthServiceHTTPServer) ListAuthResources(ctx context.Context, request *pb.ListAuthResourcesRequest) (*pb.ListAuthResourcesResponse, error) {
	return s.client.ListAuthResources(ctx, request)
}

//func (m AuthServiceHTTPServer) mustEmbedUnimplementedAuthServiceHTTPServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewAuthServiceHTTPServer new a menu service.
func NewAuthServiceHTTPServer(client pb.AuthServiceHTTPClient) *AuthServiceHTTPServer {
	return &AuthServiceHTTPServer{client: client}
}

// NewAuthServiceHTTPServerPB new a menu service.
func NewAuthServiceHTTPServerPB(client pb.AuthServiceHTTPClient) pb.AuthServiceHTTPServer {
	return &AuthServiceHTTPServer{client: client}
}

var _ pb.AuthServiceServer = (*AuthServiceHTTPServer)(nil)
