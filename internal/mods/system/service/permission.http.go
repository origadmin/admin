/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// PermissionServiceHTTPServer is a menu service.
type PermissionServiceHTTPServer struct {
	pb.UnimplementedPermissionServiceServer

	client pb.PermissionServiceHTTPClient
}

func (s PermissionServiceHTTPServer) CreatePermission(ctx context.Context, request *pb.CreatePermissionRequest) (*pb.CreatePermissionResponse, error) {
	return s.client.CreatePermission(ctx, request)
}

func (s PermissionServiceHTTPServer) DeletePermission(ctx context.Context, request *pb.DeletePermissionRequest) (*pb.DeletePermissionResponse, error) {
	return s.client.DeletePermission(ctx, request)
}

func (s PermissionServiceHTTPServer) GetPermission(ctx context.Context, request *pb.GetPermissionRequest) (*pb.GetPermissionResponse, error) {
	return s.client.GetPermission(ctx, request)
}

func (s PermissionServiceHTTPServer) ListPermissions(ctx context.Context, request *pb.ListPermissionsRequest) (*pb.ListPermissionsResponse, error) {
	return s.client.ListPermissions(ctx, request)
}

func (s PermissionServiceHTTPServer) UpdatePermission(ctx context.Context, request *pb.UpdatePermissionRequest) (*pb.UpdatePermissionResponse, error) {
	return s.client.UpdatePermission(ctx, request)
}

//func (m PermissionServiceHTTPServer) mustEmbedUnimplementedPermissionServiceHTTPServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewPermissionServiceHTTPServer new a menu service.
func NewPermissionServiceHTTPServer(client pb.PermissionServiceHTTPClient) *PermissionServiceHTTPServer {
	return &PermissionServiceHTTPServer{client: client}
}

// NewPermissionServiceHTTPServerPB new a menu service.
func NewPermissionServiceHTTPServerPB(client pb.PermissionServiceHTTPClient) pb.PermissionServiceHTTPServer {
	return &PermissionServiceHTTPServer{client: client}
}

var _ pb.PermissionServiceServer = (*PermissionServiceHTTPServer)(nil)
