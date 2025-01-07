/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// AuthServiceServer is a menu service.
type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer

	client pb.AuthServiceClient
}

func (s AuthServiceServer) ListAuthResources(ctx context.Context, request *pb.ListAuthResourcesRequest) (*pb.ListAuthResourcesResponse, error) {
	return s.client.ListAuthResources(ctx, request)
}

//func (m AuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewAuthServiceServer new a menu service.
func NewAuthServiceServer(client pb.AuthServiceClient) *AuthServiceServer {
	return &AuthServiceServer{client: client}
}

// NewAuthServiceServerPB new a menu service.
func NewAuthServiceServerPB(client pb.AuthServiceClient) pb.AuthServiceServer {
	return &AuthServiceServer{client: client}
}

var _ pb.AuthServiceServer = (*AuthServiceServer)(nil)
