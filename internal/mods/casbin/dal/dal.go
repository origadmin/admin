/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package entity implements the functions, types, and interfaces for the module.
package dal

import (
	"origadmin/application/admin/internal/mods/casbin/dal/entity/ent"
)

type Data struct {
	*ent.Database
}

func NewDataWithClient(client *ent.Client) *Data {
	return &Data{
		Database: ent.NewDatabase(client),
	}
}
