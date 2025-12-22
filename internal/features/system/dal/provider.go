/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dal implements the functions, types, and interfaces for the module.
package dal

import (
	"github.com/google/wire"

	"origadmin/application/admin/internal/features/system/dto" // Corrected import path
)

// Repositories is a collection of all repositories.
type Repositories struct {
	MenuRepo      dto.MenuRepo
	ResourceRepo  dto.ResourceRepo
	RoleRepo      dto.RoleRepo
	UserRepo      dto.UserRepo
	PermissionRepo dto.PermissionRepo
}

// NewRepositories creates a new Repositories instance.
func NewRepositories(
	menuRepo dto.MenuRepo,
	resourceRepo dto.ResourceRepo,
	roleRepo dto.RoleRepo,
	userRepo dto.UserRepo,
	permissionRepo dto.PermissionRepo,
) *Repositories {
	return &Repositories{
		MenuRepo:      menuRepo,
		ResourceRepo:  resourceRepo,
		RoleRepo:      roleRepo,
		UserRepo:      userRepo,
		PermissionRepo: permissionRepo,
	}
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewMenuRepo,
	NewResourceRepo,
	NewRoleRepo,
	NewUserRepo,
	NewPermissionRepo,
	NewRepositories, // Provide the aggregated Repositories struct
)
