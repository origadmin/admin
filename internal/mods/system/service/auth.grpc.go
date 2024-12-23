/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// AuthAPIService is a menu service.
type AuthAPIService struct {
	pb.UnimplementedAuthAPIServer

	client pb.AuthAPIClient
}

func (s AuthAPIService) ListResources(ctx context.Context, request *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	return s.client.ListResources(ctx, request)
}

//func (m AuthAPIService) mustEmbedUnimplementedAuthAPIServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewAuthAPIService new a menu service.
func NewAuthAPIService(client pb.AuthAPIClient) *AuthAPIService {
	return &AuthAPIService{client: client}
}

// NewAuthAPIServer new a menu service.
func NewAuthAPIServer(client pb.AuthAPIClient) pb.AuthAPIServer {
	return &AuthAPIService{client: client}
}

var _ pb.AuthAPIServer = (*AuthAPIService)(nil)
