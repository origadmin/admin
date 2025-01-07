/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"google.golang.org/grpc"

	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// ResourceServiceClientBiz is a ResourcePB use case.
type ResourceServiceClientBiz struct {
	dao     dto.ResourceRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz ResourceServiceClientBiz) ListResources(ctx context.Context, in *pb.ListResourcesRequest, opts ...grpc.CallOption) (*pb.ListResourcesResponse, error) {
	var option dto.ResourceQueryOption
	if err := option.FromListRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("ListResources")
	result, total, err := biz.dao.List(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ToListResourcesResponse(result, in, total)
}

func (biz ResourceServiceClientBiz) GetResource(ctx context.Context, in *pb.GetResourceRequest, opts ...grpc.CallOption) (*pb.GetResourceResponse, error) {
	var option dto.ResourceQueryOption
	if err := option.FromGetRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("GetResource")
	result, err := biz.dao.Get(ctx, in.GetId(), option)
	if err != nil {
		return nil, err
	}
	return &pb.GetResourceResponse{
		Resource: result,
	}, nil
}

func (biz ResourceServiceClientBiz) CreateResource(ctx context.Context, in *pb.CreateResourceRequest, opts ...grpc.CallOption) (*pb.CreateResourceResponse, error) {
	var option dto.ResourceQueryOption
	if err := option.FromCreateRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("CreateResource")
	result, err := biz.dao.Create(ctx, in.Resource, option)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResourceResponse{
		Resource: result,
	}, nil
}

func (biz ResourceServiceClientBiz) UpdateResource(ctx context.Context, in *pb.UpdateResourceRequest, opts ...grpc.CallOption) (*pb.UpdateResourceResponse, error) {
	//var option dto.UpdateResourceOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	log.Info("UpdateResource")
	result, err := biz.dao.Update(ctx, in.Resource)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResourceResponse{
		Resource: result,
	}, nil
}

func (biz ResourceServiceClientBiz) DeleteResource(ctx context.Context, in *pb.DeleteResourceRequest, opts ...grpc.CallOption) (*pb.DeleteResourceResponse, error) {
	//var option dto.DeleteResourceOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	//_, err := biz.dao.Get(ctx, in.GetId())
	//if err != nil {
	//	return nil, err
	//}
	log.Info("DeleteResource")
	if err := biz.dao.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteResourceResponse{}, nil
}

// NewResourceServiceClientBiz new a ResourcePB use case.
func NewResourceServiceClientBiz(repo dto.ResourceRepo, logger log.KLogger) *ResourceServiceClientBiz {
	return &ResourceServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewResourceServiceClient new a ResourcePB use case.
func NewResourceServiceClient(repo dto.ResourceRepo, logger log.KLogger) pb.ResourceServiceClient {
	return &ResourceServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.ResourceServiceClient = (*ResourceServiceClientBiz)(nil)
