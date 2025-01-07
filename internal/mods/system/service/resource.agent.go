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

// ResourceServiceAgent is a menu service.
type ResourceServiceAgent struct {
	resp.Response

	client pb.ResourceServiceClient
}

func (s ResourceServiceAgent) CreateResource(ctx context.Context, request *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
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

func (s ResourceServiceAgent) DeleteResource(ctx context.Context, request *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
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

func (s ResourceServiceAgent) GetResource(ctx context.Context, request *pb.GetResourceRequest) (*pb.GetResourceResponse, error) {
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

func (s ResourceServiceAgent) ListResources(ctx context.Context, request *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
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

func (s ResourceServiceAgent) UpdateResource(ctx context.Context, request *pb.UpdateResourceRequest) (*pb.UpdateResourceResponse, error) {
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

// NewResourceServiceAgent new a menu service.
func NewResourceServiceAgent(client pb.ResourceServiceClient) *ResourceServiceAgent {
	return &ResourceServiceAgent{client: client}
}

// NewResourceServiceAgentPB new a menu service.
func NewResourceServiceAgentPB(client pb.ResourceServiceClient) pb.ResourceServiceAgent {
	return &ResourceServiceAgent{client: client}
}
func NewResourceServiceAgentClient(client *service.GRPCClient) pb.ResourceServiceAgent {
	cli := pb.NewResourceServiceClient(client)
	return NewResourceServiceAgent(cli)
}

var _ pb.ResourceServiceAgent = (*ResourceServiceAgent)(nil)
