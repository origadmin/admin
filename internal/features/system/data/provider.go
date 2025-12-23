/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data implements the functions, types, and interfaces for the module.
package data

import (
	"origadmin/application/admin/internal/features/system/dto"
)

// Repositories is a collection of all repositories.
type Repositories struct {
	ResourceRepo   dto.ResourceRepo
	RoleRepo       dto.RoleRepo
	UserRepo       dto.UserRepo
	PermissionRepo dto.PermissionRepo
}

// NewRepositories creates a new Repositories instance.
func NewRepositories(
	resourceRepo dto.ResourceRepo,
	roleRepo dto.RoleRepo,
	userRepo dto.UserRepo,
	permissionRepo dto.PermissionRepo,
) (*Repositories, error) {
	return &Repositories{
		ResourceRepo:   resourceRepo,
		RoleRepo:       roleRepo,
		UserRepo:       userRepo,
		PermissionRepo: permissionRepo,
	}, nil
}
