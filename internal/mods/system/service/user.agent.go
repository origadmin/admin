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

// UserAPIGINRPCService is a menu service.
type UserAPIGINRPCService struct {
	resp.Response

	client pb.UserAPIClient
}

func (s UserAPIGINRPCService) UpdateUserRoles(ctx context.Context, request *pb.UpdateUserRolesRequest) (*pb.UpdateUserRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserAPIGINRPCService) ResetUserPassword(ctx context.Context, request *pb.ResetUserPasswordRequest) (*pb.ResetUserPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserAPIGINRPCService) UpdateUserStatus(ctx context.Context, request *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserAPIGINRPCService) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response.User),
	})
	return nil, nil
}

func (s UserAPIGINRPCService) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	_, err := s.client.DeleteUser(ctx, request)
	if err != nil {
		return nil, err
	}

	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    nil,
	})
	return nil, nil
}

func (s UserAPIGINRPCService) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response.User),
	})
	return nil, nil
}

func (s UserAPIGINRPCService) ListUsers(ctx context.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListUsers(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		Total:   response.TotalSize,
		Data:    resp.Proto2AnyPBArray(response.Users...),
	})
	return nil, nil
}

func (s UserAPIGINRPCService) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateUser(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response.User),
	})
	return nil, nil
}

// NewUserAPIGINRPCService new a menu service.
func NewUserAPIGINRPCService(client pb.UserAPIClient) *UserAPIGINRPCService {
	return &UserAPIGINRPCService{client: client}
}

// NewUserAPIAgent new a menu service.
func NewUserAPIAgent(client pb.UserAPIClient) pb.UserAPIAgent {
	return &UserAPIGINRPCService{client: client}
}
func NewUserServerAgent(client *service.GRPCClient) pb.UserAPIAgent {
	c := pb.NewUserAPIClient(client)
	return NewUserAPIAgent(c)
}

var _ pb.UserAPIAgent = (*UserAPIGINRPCService)(nil)
