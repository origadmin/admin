/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"net/http"

	"github.com/origadmin/runtime/agent"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// RoleServiceBridge is a menu service.
type RoleServiceBridge struct {
	resp.Response

	client pb.RoleServiceClient
}

func (s RoleServiceBridge) CreateRole(ctx context.Context, request *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CreateRole(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Role),
	})
	return nil, nil
}

func (s RoleServiceBridge) DeleteRole(ctx context.Context, request *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.DeleteRole(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Empty),
	})
	return nil, nil
}

func (s RoleServiceBridge) GetRole(ctx context.Context, request *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.GetRole(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Role),
	})
	return nil, nil
}

func (s RoleServiceBridge) ListRoles(ctx context.Context, request *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListRoles(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		Total:   response.TotalSize,
		Data:    resp.Proto2AnyPBArray(response.Roles...),
	})
	return nil, nil
}

func (s RoleServiceBridge) UpdateRole(ctx context.Context, request *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateRole(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Role),
	})
	return nil, nil
}

// NewRoleServiceBridge new a menu service.
func NewRoleServiceBridge(client pb.RoleServiceClient) *RoleServiceBridge {
	return &RoleServiceBridge{client: client}
}

// NewRoleServiceBridgePB new a menu service.
func NewRoleServiceBridgePB(client pb.RoleServiceClient) pb.RoleServiceBridge {
	return &RoleServiceBridge{client: client}
}
func NewRoleServiceBridgeClient(client *service.GRPCClient) pb.RoleServiceBridge {
	c := pb.NewRoleServiceClient(client)
	return NewRoleServiceBridge(c)
}

var _ pb.RoleServiceBridge = (*RoleServiceBridge)(nil)
