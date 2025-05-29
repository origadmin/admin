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

// UserServiceBridge is a menu service.
type UserServiceBridge struct {
	resp.Response

	client pb.UserServiceClient
}

func (s UserServiceBridge) ListUserResources(ctx context.Context, request *pb.ListUserResourcesRequest) (*pb.ListUserResourcesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListUserResources(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.DataArray{
		Success: true,
		Data:    resp.Proto2AnyPBArray(response.Resources...),
	})
	return nil, nil
}

func (s UserServiceBridge) UpdateUserRoles(ctx context.Context, request *pb.UpdateUserRolesRequest) (*pb.UpdateUserRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceBridge) ResetUserPassword(ctx context.Context, request *pb.ResetUserPasswordRequest) (*pb.ResetUserPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceBridge) UpdateUserStatus(ctx context.Context, request *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceBridge) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.User),
	})
	return nil, nil
}

func (s UserServiceBridge) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
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

func (s UserServiceBridge) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.User),
	})
	return nil, nil
}

func (s UserServiceBridge) ListUsers(ctx context.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
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

func (s UserServiceBridge) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateUser(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.User),
	})
	return nil, nil
}

// NewUserServiceBridge new a menu service.
func NewUserServiceBridge(client pb.UserServiceClient) *UserServiceBridge {
	return &UserServiceBridge{client: client}
}

// NewUserServiceBridgePB new a menu service.
func NewUserServiceBridgePB(client pb.UserServiceClient) pb.UserServiceBridge {
	return &UserServiceBridge{client: client}
}
func NewUserServiceBridgeClient(client *service.GRPCClient) pb.UserServiceBridge {
	c := pb.NewUserServiceClient(client)
	return NewUserServiceBridge(c)
}

var _ pb.UserServiceBridge = (*UserServiceBridge)(nil)
