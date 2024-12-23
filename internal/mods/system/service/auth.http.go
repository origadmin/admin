/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// AuthAPIHTTPService is a menu service.
type AuthAPIHTTPService struct {
	pb.UnimplementedAuthAPIServer

	client pb.AuthAPIHTTPClient
}

func (s AuthAPIHTTPService) ListResources(ctx context.Context, request *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	return s.client.ListResources(ctx, request)
}

//func (m AuthAPIHTTPService) mustEmbedUnimplementedAuthAPIHTTPServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewAuthAPIHTTPService new a menu service.
func NewAuthAPIHTTPService(client pb.AuthAPIHTTPClient) *AuthAPIHTTPService {
	return &AuthAPIHTTPService{client: client}
}

// NewAuthAPIHTTPServer new a menu service.
func NewAuthAPIHTTPServer(client pb.AuthAPIHTTPClient) pb.AuthAPIServer {
	return &AuthAPIHTTPService{client: client}
}

var _ pb.AuthAPIServer = (*AuthAPIHTTPService)(nil)
