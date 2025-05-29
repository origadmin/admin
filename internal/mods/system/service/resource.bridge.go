/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	"github.com/origadmin/runtime/agent"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// ResourceServiceBridge is a menu service.
type ResourceServiceBridge struct {
	resp.Response

	client pb.ResourceServiceClient
}

func (s ResourceServiceBridge) CreateResource(ctx context.Context, request *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CreateResource(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Resource),
	})
	return nil, nil
}

func (s ResourceServiceBridge) DeleteResource(ctx context.Context, request *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	_, err := s.client.DeleteResource(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    nil,
	})
	return nil, nil
}

func (s ResourceServiceBridge) GetResource(ctx context.Context, request *pb.GetResourceRequest) (*pb.GetResourceResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.GetResource(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Resource),
	})
	return nil, nil
}

func (s ResourceServiceBridge) ListResources(ctx context.Context, request *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListResources(ctx, request)
	if err != nil {
		return nil, err
	}

	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		Total:   response.TotalSize,
		Data:    resp.Proto2AnyPBArray(response.Resources...),
	})
	return nil, nil
}

func (s ResourceServiceBridge) UpdateResource(ctx context.Context, request *pb.UpdateResourceRequest) (*pb.UpdateResourceResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateResource(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Resource),
	})
	return nil, nil
}

// NewResourceServiceBridge new a menu service.
func NewResourceServiceBridge(client pb.ResourceServiceClient) *ResourceServiceBridge {
	return &ResourceServiceBridge{client: client}
}

// NewResourceServiceBridgePB new a menu service.
func NewResourceServiceBridgePB(client pb.ResourceServiceClient) pb.ResourceServiceBridge {
	return &ResourceServiceBridge{client: client}
}
func NewResourceServiceBridgeClient(client *service.GRPCClient) pb.ResourceServiceBridge {
	cli := pb.NewResourceServiceClient(client)
	return NewResourceServiceBridge(cli)
}

var _ pb.ResourceServiceBridge = (*ResourceServiceBridge)(nil)
