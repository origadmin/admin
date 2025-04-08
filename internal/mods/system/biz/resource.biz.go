/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dto"
)

// ResourceServiceBiz is a ResourcePB use case.
type ResourceServiceBiz struct {
	dao     dto.ResourceRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz ResourceServiceBiz) ListResources(ctx context.Context, in *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
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

func (biz ResourceServiceBiz) GetResource(ctx context.Context, in *pb.GetResourceRequest) (*pb.GetResourceResponse, error) {
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

func (biz ResourceServiceBiz) CreateResource(ctx context.Context, in *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
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

var updateFields = []string{
	resource.FieldIcon, resource.FieldType, resource.FieldStatus,
	resource.FieldName, resource.FieldPath, resource.FieldKeyword,
	resource.FieldSequence, resource.FieldProperties, resource.FieldDescription,
}

func (biz ResourceServiceBiz) UpdateResource(ctx context.Context, in *pb.UpdateResourceRequest) (*pb.UpdateResourceResponse, error) {
	var option dto.ResourceQueryOption

	log.Info("UpdateResource")
	option.Fields = resource.SelectColumns(updateFields)
	result, err := biz.dao.Update(ctx, in.Resource, option)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResourceResponse{
		Resource: result,
	}, nil
}

func (biz ResourceServiceBiz) DeleteResource(ctx context.Context, in *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
	log.Info("DeleteResource")
	if err := biz.dao.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteResourceResponse{}, nil
}

// NewResourceServiceBiz new a ResourcePB use case.
func NewResourceServiceBiz(repo dto.ResourceRepo, logger log.KLogger) *ResourceServiceBiz {
	return &ResourceServiceBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}
