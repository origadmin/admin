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

// RoleServiceAgent is a menu service.
type RoleServiceAgent struct {
	resp.Response

	client pb.RoleServiceClient
}

func (s RoleServiceAgent) CreateRole(ctx context.Context, request *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
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

func (s RoleServiceAgent) DeleteRole(ctx context.Context, request *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
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

func (s RoleServiceAgent) GetRole(ctx context.Context, request *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
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

func (s RoleServiceAgent) ListRoles(ctx context.Context, request *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
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

func (s RoleServiceAgent) UpdateRole(ctx context.Context, request *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateRole(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response.Role),
	})
	return nil, nil
}

// NewRoleServiceAgent new a menu service.
func NewRoleServiceAgent(client pb.RoleServiceClient) *RoleServiceAgent {
	return &RoleServiceAgent{client: client}
}

// NewRoleServiceAgentPB new a menu service.
func NewRoleServiceAgentPB(client pb.RoleServiceClient) pb.RoleServiceAgent {
	return &RoleServiceAgent{client: client}
}
func NewRoleServiceAgentClient(client *service.GRPCClient) pb.RoleServiceAgent {
	c := pb.NewRoleServiceClient(client)
	return NewRoleServiceAgent(c)
}

var _ pb.RoleServiceAgent = (*RoleServiceAgent)(nil)
