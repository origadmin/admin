/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type RoleAPIHTTPService struct {
	pb.UnimplementedRoleAPIServer

	client pb.RoleAPIHTTPClient
}

func (s *RoleAPIHTTPService) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	return &pb.ListRolesResponse{}, nil
}
func (s *RoleAPIHTTPService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	return &pb.GetRoleResponse{}, nil
}
func (s *RoleAPIHTTPService) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return &pb.CreateRoleResponse{}, nil
}
func (s *RoleAPIHTTPService) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	return &pb.UpdateRoleResponse{}, nil
}
func (s *RoleAPIHTTPService) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	return &pb.DeleteRoleResponse{}, nil
}

// NewRoleAPIHTTPService new a role service.
func NewRoleAPIHTTPService(client pb.RoleAPIHTTPClient) *RoleAPIHTTPService {
	return &RoleAPIHTTPService{
		client: client,
	}
}

// NewRoleAPIHTTPServer new a role service.
func NewRoleAPIHTTPServer(client pb.RoleAPIHTTPClient) pb.RoleAPIHTTPServer {
	return &RoleAPIHTTPService{
		client: client,
	}
}

var _ pb.RoleAPIServer = (*RoleAPIHTTPService)(nil)
