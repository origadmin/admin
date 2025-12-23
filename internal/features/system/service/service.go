/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
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
