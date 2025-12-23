/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"

	"github.com/origadmin/runtime/service"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/system/biz"
)

type SystemService struct {
	system.UnimplementedResourceServiceServer
	system.UnimplementedRoleServiceServer
	system.UnimplementedUserServiceServer
	system.UnimplementedPermissionServiceServer

	resource   *biz.ResourceUseCase
	role       *biz.RoleUseCase
	user       *biz.UserUseCase
	permission *biz.PermissionUseCase
}

func New(
	resource *biz.ResourceUseCase,
	role *biz.RoleUseCase,
	user *biz.UserUseCase,
	permission *biz.PermissionUseCase,
) *SystemService {
	return &SystemService{
		resource:   resource,
		role:       role,
		user:       user,
		permission: permission,
	}
}

func (s *SystemService) Register(ctx context.Context, srv any) {
	switch srv.(type) {
	case *transport.Server:
	case *service.Server:
	}
	system.RegisterResourceServiceServer(srv.GRPC, s)
	system.RegisterRoleServiceServer(srv.GRPC, s)
	system.RegisterUserServiceServer(srv.GRPC, s)
	system.RegisterPermissionServiceServer(srv.GRPC, s)

	system.RegisterResourceServiceHTTPServer(srv.HTTP, s)
	system.RegisterRoleServiceHTTPServer(srv.HTTP, s)
	system.RegisterUserServiceHTTPServer(srv.HTTP, s)
	system.RegisterPermissionServiceHTTPServer(srv.HTTP, s)
}
