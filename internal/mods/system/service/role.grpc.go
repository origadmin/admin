/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type RoleServiceServer struct {
	pb.UnimplementedRoleServiceServer

	client pb.RoleServiceClient
}

func (s RoleServiceServer) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	return s.client.ListRoles(ctx, req)
}
func (s RoleServiceServer) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	return s.client.GetRole(ctx, req)
}
func (s RoleServiceServer) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return s.client.CreateRole(ctx, req)
}
func (s RoleServiceServer) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	return s.client.UpdateRole(ctx, req)
}
func (s RoleServiceServer) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	return s.client.DeleteRole(ctx, req)
}

// NewRoleServiceServer new a user service.
func NewRoleServiceServer(client pb.RoleServiceClient) *RoleServiceServer {
	return &RoleServiceServer{
		client: client,
	}
}

// NewRoleServiceServerPB new a user service.
func NewRoleServiceServerPB(client pb.RoleServiceClient) pb.RoleServiceServer {
	return &RoleServiceServer{
		client: client,
	}
}

var _ pb.RoleServiceServer = (*RoleServiceServer)(nil)
