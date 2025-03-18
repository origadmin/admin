/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dal implements the functions, types, and interfaces for the module.
package dal

import (
	"context"

	"origadmin/application/admin/internal/mods/casbin/dal/entity/ent"
	"origadmin/application/admin/internal/mods/casbin/dto"
)

type casbinSourceRepo struct {
	ctx  context.Context
	data *Data
}

// NewCasbinSourceRepo returns a new CasbinSourceRepo
func NewCasbinSourceRepo(data *Data) (dto.CasbinSourceRepo, error) {
	a := &casbinSourceRepo{
		data: data,
	}
	return a, nil
}

// NewCasbinSourceWithClient create a new CasbinSourceRepo with given client.
// This method does not ensure the existence of database, user should create database manually.
func NewCasbinSourceWithClient(client *ent.Client) (dto.CasbinSourceRepo, error) {
	a := &casbinSourceRepo{
		data: NewDataWithClient(client),
	}
	return a, nil
}
