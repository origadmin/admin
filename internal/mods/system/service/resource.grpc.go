/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/biz"
)

// ResourceServiceServer is a menu service.
type ResourceServiceServer struct {
	pb.UnimplementedResourceServiceServer

	client *biz.ResourceServiceBiz
}

func (s ResourceServiceServer) ListResources(ctx context.Context, request *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	return s.client.ListResources(ctx, request)
}

func (s ResourceServiceServer) GetResource(ctx context.Context, request *pb.GetResourceRequest) (*pb.GetResourceResponse, error) {
	return s.client.GetResource(ctx, request)
}

func (s ResourceServiceServer) CreateResource(ctx context.Context, request *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
	return s.client.CreateResource(ctx, request)
}

func (s ResourceServiceServer) UpdateResource(ctx context.Context, request *pb.UpdateResourceRequest) (*pb.UpdateResourceResponse, error) {
	return s.client.UpdateResource(ctx, request)
}

func (s ResourceServiceServer) DeleteResource(ctx context.Context, request *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
	return s.client.DeleteResource(ctx, request)
}

//func (m ResourceServiceServer) mustEmbedUnimplementedResourceServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewResourceServiceServer new a menu service.
func NewResourceServiceServer(client *biz.ResourceServiceBiz) *ResourceServiceServer {
	return &ResourceServiceServer{client: client}
}

// NewResourceServiceServerPB new a menu service.
func NewResourceServiceServerPB(client *biz.ResourceServiceBiz) pb.ResourceServiceServer {
	return &ResourceServiceServer{client: client}
}

var _ pb.ResourceServiceServer = (*ResourceServiceServer)(nil)
