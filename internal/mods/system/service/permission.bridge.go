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

// PermissionServiceBridge is a menu service.
type PermissionServiceBridge struct {
	resp.Response

	client pb.PermissionServiceClient
}

func (s PermissionServiceBridge) CreatePermission(ctx context.Context, request *pb.CreatePermissionRequest) (*pb.CreatePermissionResponse, error) {
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

func (s PermissionServiceBridge) DeletePermission(ctx context.Context, request *pb.DeletePermissionRequest) (*pb.DeletePermissionResponse, error) {
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

func (s PermissionServiceBridge) GetPermission(ctx context.Context, request *pb.GetPermissionRequest) (*pb.GetPermissionResponse, error) {
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

func (s PermissionServiceBridge) ListPermissions(ctx context.Context, request *pb.ListPermissionsRequest) (*pb.ListPermissionsResponse, error) {
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

func (s PermissionServiceBridge) UpdatePermission(ctx context.Context, request *pb.UpdatePermissionRequest) (*pb.UpdatePermissionResponse, error) {
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

// NewPermissionServiceBridge new a menu service.
func NewPermissionServiceBridge(client pb.PermissionServiceClient) *PermissionServiceBridge {
	return &PermissionServiceBridge{client: client}
}

// NewPermissionServiceBridgePB new a menu service.
func NewPermissionServiceBridgePB(client pb.PermissionServiceClient) pb.PermissionServiceBridge {
	return &PermissionServiceBridge{client: client}
}
func NewPermissionServiceBridgeClient(client *service.GRPCClient) pb.PermissionServiceBridge {
	cli := pb.NewPermissionServiceClient(client)
	return NewPermissionServiceBridge(cli)
}

var _ pb.PermissionServiceBridge = (*PermissionServiceBridge)(nil)
