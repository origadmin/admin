/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the auth module.
package dto

//go:generate abgen -debug .

//go:abgen:package:path=origadmin/application/admin/internal/data/entity/ent,alias=ent
//go:abgen:package:path=origadmin/application/admin/api/v1/services/types,alias=types
//go:abgen:pair:packages="ent,types"
//go:abgen:convert:direction="both"
//go:abgen:convert:source:suffix=""
//go:abgen:convert:target:suffix="PB"
