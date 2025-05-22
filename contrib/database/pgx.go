//go:build !postgres && pgx

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package database implements the functions, types, and interfaces for the module.
package database

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Pgx struct{}
