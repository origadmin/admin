/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/system/biz" // Corrected import path
)

// PermissionServiceServer is a menu service.
type PermissionServiceServer struct {
	pb.UnimplementedPermissionServiceServer

	client *biz.PermissionServiceBiz
	log    *log.KHelper
}

func (s PermissionServiceServer) ListPermissions(ctx context.Context, request *pb.ListPermissionsRequest) (*pb.ListPermissionsResponse, error) {
	return s.client.ListPermissions(ctx, request)
}

func (s PermissionServiceServer) GetPermission(ctx context.Context, request *pb.GetPermissionRequest) (*pb.GetPermissionResponse, error) {
	return s.client.GetPermission(ctx, request)
}

func (s PermissionServiceServer) CreatePermission(ctx context.Context, request *pb.CreatePermissionRequest) (*pb.CreatePermissionResponse, error) {
	return s.client.CreatePermission(ctx, request)
}

func (s PermissionServiceServer) UpdatePermission(ctx context.Context, request *pb.UpdatePermissionRequest) (*pb.UpdatePermissionResponse, error) {
	return s.client.UpdatePermission(ctx, request)
}

func (s PermissionServiceServer) DeletePermission(ctx context.Context, request *pb.DeletePermissionRequest) (*pb.DeletePermissionResponse, error) {
	return s.client.DeletePermission(ctx, request)
}

//func (m PermissionServiceServer) mustEmbedUnimplementedPermissionServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewPermissionServiceServer new a menu service.
func NewPermissionServiceServer(r runtime.Runtime, client *biz.PermissionServiceBiz) *PermissionServiceServer {
	return &PermissionServiceServer{
		log:    log.NewHelper(r.WithLogger("module", "service/system")),
		client: client,
	}
}

// NewPermissionServiceServerPB new a menu service.
func NewPermissionServiceServerPB(r runtime.Runtime, client *biz.PermissionServiceBiz) pb.PermissionServiceServer {
	return NewPermissionServiceServer(r, client)
}

var _ pb.PermissionServiceServer = (*PermissionServiceServer)(nil)
