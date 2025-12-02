/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// ResourceServiceHTTPServer is a menu service.
type ResourceServiceHTTPServer struct {
	pb.UnimplementedResourceServiceServer

	client pb.ResourceServiceHTTPClient
}

func (s ResourceServiceHTTPServer) CreateResource(ctx context.Context, request *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
	return s.client.CreateResource(ctx, request)
}

func (s ResourceServiceHTTPServer) DeleteResource(ctx context.Context, request *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
	return s.client.DeleteResource(ctx, request)
}

func (s ResourceServiceHTTPServer) GetResource(ctx context.Context, request *pb.GetResourceRequest) (*pb.GetResourceResponse, error) {
	return s.client.GetResource(ctx, request)
}

func (s ResourceServiceHTTPServer) ListResources(ctx context.Context, request *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	return s.client.ListResources(ctx, request)
}

func (s ResourceServiceHTTPServer) UpdateResource(ctx context.Context, request *pb.UpdateResourceRequest) (*pb.UpdateResourceResponse, error) {
	return s.client.UpdateResource(ctx, request)
}

//func (m ResourceServiceHTTPServer) mustEmbedUnimplementedResourceServiceHTTPServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewResourceServiceHTTPServer new a menu service.
func NewResourceServiceHTTPServer(client pb.ResourceServiceHTTPClient) *ResourceServiceHTTPServer {
	return &ResourceServiceHTTPServer{client: client}
}

// NewResourceServiceHTTPServerPB new a menu service.
func NewResourceServiceHTTPServerPB(client pb.ResourceServiceHTTPClient) pb.ResourceServiceHTTPServer {
	return &ResourceServiceHTTPServer{client: client}
}

var _ pb.ResourceServiceServer = (*ResourceServiceHTTPServer)(nil)
