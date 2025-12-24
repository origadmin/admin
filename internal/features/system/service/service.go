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
	system.UnimplementedViewServiceServer

	Resource   *biz.ResourceUseCase
	Role       *biz.RoleUseCase
	User       *biz.UserUseCase
	Permission *biz.PermissionUseCase
	View       *biz.ViewUseCase
}

func New(
	resource *biz.ResourceUseCase,
	role *biz.RoleUseCase,
	user *biz.UserUseCase,
	permission *biz.PermissionUseCase,
	view *biz.ViewUseCase,
) *SystemService {
	return &SystemService{
		Resource:   resource,
		Role:       role,
		User:       user,
		Permission: permission,
		View:       view,
	}
}
