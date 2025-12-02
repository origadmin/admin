/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type RoleServiceHTTPServer struct {
	pb.UnimplementedRoleServiceServer

	client pb.RoleServiceHTTPClient
}

func (s RoleServiceHTTPServer) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	return s.client.ListRoles(ctx, req)
}
func (s RoleServiceHTTPServer) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	return s.client.GetRole(ctx, req)
}
func (s RoleServiceHTTPServer) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return s.client.CreateRole(ctx, req)
}
func (s RoleServiceHTTPServer) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	return s.client.UpdateRole(ctx, req)
}
func (s RoleServiceHTTPServer) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	return s.client.DeleteRole(ctx, req)
}

// NewRoleServiceHTTPServer new a role service.
func NewRoleServiceHTTPServer(client pb.RoleServiceHTTPClient) *RoleServiceHTTPServer {
	return &RoleServiceHTTPServer{
		client: client,
	}
}

// NewRoleServiceHTTPServerPB new a role service.
func NewRoleServiceHTTPServerPB(client pb.RoleServiceHTTPClient) pb.RoleServiceHTTPServer {
	return &RoleServiceHTTPServer{
		client: client,
	}
}

var _ pb.RoleServiceServer = (*RoleServiceHTTPServer)(nil)
