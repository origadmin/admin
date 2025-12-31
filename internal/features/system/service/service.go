/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

type SystemService struct {
	Resource   *ResourceService
	Role       *RoleService
	User       *UserService
	Permission *PermissionService
	View       *ViewService
}

func NewSystemService(
	resource *ResourceService,
	role *RoleService,
	user *UserService,
	permission *PermissionService,
	view *ViewService,
) *SystemService {
	return &SystemService{
		Resource:   resource,
		Role:       role,
		User:       user,
		Permission: permission,
		View:       view,
	}
}
