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
	"google.golang.org/protobuf/proto"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// UserServiceAgent is a menu service.
type UserServiceAgent struct {
	resp.Response

	client pb.UserServiceClient
}

func (s UserServiceAgent) ListUserResources(ctx context.Context, request *pb.ListUserResourcesRequest) (*pb.ListUserResourcesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListUserResources(ctx, request)
	if err != nil {
		return nil, err
	}
	var data []proto.Message
	for _, v := range response.Resources {
		data = append(data, v)
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.ProtoArray2AnyPB(data...),
	})
	return nil, nil
}

func (s UserServiceAgent) UpdateUserRoles(ctx context.Context, request *pb.UpdateUserRolesRequest) (*pb.UpdateUserRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceAgent) ResetUserPassword(ctx context.Context, request *pb.ResetUserPasswordRequest) (*pb.ResetUserPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceAgent) UpdateUserStatus(ctx context.Context, request *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceAgent) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
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

func (s UserServiceAgent) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
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

func (s UserServiceAgent) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
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

func (s UserServiceAgent) ListUsers(ctx context.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
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

func (s UserServiceAgent) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
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

// NewUserServiceAgent new a menu service.
func NewUserServiceAgent(client pb.UserServiceClient) *UserServiceAgent {
	return &UserServiceAgent{client: client}
}

// NewUserServiceAgentPB new a menu service.
func NewUserServiceAgentPB(client pb.UserServiceClient) pb.UserServiceAgent {
	return &UserServiceAgent{client: client}
}
func NewUserServiceAgentClient(client *service.GRPCClient) pb.UserServiceAgent {
	c := pb.NewUserServiceClient(client)
	return NewUserServiceAgent(c)
}

var _ pb.UserServiceAgent = (*UserServiceAgent)(nil)
