/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type RoleAPIService struct {
	pb.UnimplementedRoleAPIServer

	client pb.RoleAPIClient
}

func (s *RoleAPIService) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	return &pb.ListRolesResponse{}, nil
}
func (s *RoleAPIService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	return &pb.GetRoleResponse{}, nil
}
func (s *RoleAPIService) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return &pb.CreateRoleResponse{}, nil
}
func (s *RoleAPIService) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	return &pb.UpdateRoleResponse{}, nil
}
func (s *RoleAPIService) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	return &pb.DeleteRoleResponse{}, nil
}

// NewRoleAPIService new a user service.
func NewRoleAPIService(client pb.RoleAPIClient) *RoleAPIService {
	return &RoleAPIService{
		client: client,
	}
}

// NewRoleAPIServer new a user service.
func NewRoleAPIServer(client pb.RoleAPIClient) pb.RoleAPIServer {
	return &RoleAPIService{
		client: client,
	}
}

var _ pb.RoleAPIServer = (*RoleAPIService)(nil)
