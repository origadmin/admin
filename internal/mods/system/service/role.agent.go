/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// RoleAPIGINRPCService is a menu service.
type RoleAPIGINRPCService struct {
	resp.Response

	client pb.RoleAPIClient
}

func (s RoleAPIGINRPCService) CreateRole(context transhttp.Context, request *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	response, err := s.client.CreateRole(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Role),
	})
	return nil, nil
}

func (s RoleAPIGINRPCService) DeleteRole(context transhttp.Context, request *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	response, err := s.client.DeleteRole(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Empty),
	})
	return nil, nil
}

func (s RoleAPIGINRPCService) GetRole(context transhttp.Context, request *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	response, err := s.client.GetRole(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Role),
	})
	return nil, nil
}

func (s RoleAPIGINRPCService) ListRoles(context transhttp.Context, request *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	response, err := s.client.ListRoles(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Page{
		Success: true,
		Total:   response.TotalSize,
		Data:    resp.Proto2AnyPBArray(response.Roles...),
	})
	return nil, nil
}

func (s RoleAPIGINRPCService) UpdateRole(context transhttp.Context, request *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	response, err := s.client.UpdateRole(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response.Role),
	})
	return nil, nil
}

// NewRoleAPIGINRPCService new a menu service.
func NewRoleAPIGINRPCService(client pb.RoleAPIClient) *RoleAPIGINRPCService {
	return &RoleAPIGINRPCService{client: client}
}

// NewRoleAPIAgent new a menu service.
func NewRoleAPIAgent(client pb.RoleAPIClient) pb.RoleAPIAgent {
	return &RoleAPIGINRPCService{client: client}
}
func NewRoleServerAgent(client *service.GRPCClient) pb.RoleAPIAgent {
	c := pb.NewRoleAPIClient(client)
	return NewRoleAPIAgent(c)
}

var _ pb.RoleAPIAgent = (*RoleAPIGINRPCService)(nil)
