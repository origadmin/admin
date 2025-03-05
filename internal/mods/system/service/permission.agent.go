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

// PermissionServiceAgent is a menu service.
type PermissionServiceAgent struct {
	resp.Response

	client pb.PermissionServiceClient
}

func (s PermissionServiceAgent) CreatePermission(ctx context.Context, request *pb.CreatePermissionRequest) (*pb.CreatePermissionResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CreatePermission(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Permission),
	})
	return nil, nil
}

func (s PermissionServiceAgent) DeletePermission(ctx context.Context, request *pb.DeletePermissionRequest) (*pb.DeletePermissionResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	_, err := s.client.DeletePermission(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    nil,
	})
	return nil, nil
}

func (s PermissionServiceAgent) GetPermission(ctx context.Context, request *pb.GetPermissionRequest) (*pb.GetPermissionResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.GetPermission(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Permission),
	})
	return nil, nil
}

func (s PermissionServiceAgent) ListPermissions(ctx context.Context, request *pb.ListPermissionsRequest) (*pb.ListPermissionsResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListPermissions(ctx, request)
	if err != nil {
		return nil, err
	}

	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		Total:   response.TotalSize,
		Data:    resp.Proto2AnyPBArray(response.Permissions...),
	})
	return nil, nil
}

func (s PermissionServiceAgent) UpdatePermission(ctx context.Context, request *pb.UpdatePermissionRequest) (*pb.UpdatePermissionResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdatePermission(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Permission),
	})
	return nil, nil
}

// NewPermissionServiceAgent new a menu service.
func NewPermissionServiceAgent(client pb.PermissionServiceClient) *PermissionServiceAgent {
	return &PermissionServiceAgent{client: client}
}

// NewPermissionServiceAgentPB new a menu service.
func NewPermissionServiceAgentPB(client pb.PermissionServiceClient) pb.PermissionServiceAgent {
	return &PermissionServiceAgent{client: client}
}
func NewPermissionServiceAgentClient(client *service.GRPCClient) pb.PermissionServiceAgent {
	cli := pb.NewPermissionServiceClient(client)
	return NewPermissionServiceAgent(cli)
}

var _ pb.PermissionServiceAgent = (*PermissionServiceAgent)(nil)
