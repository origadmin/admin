//go:build mssql || sqlserver

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package database implements the functions, types, and interfaces for the module.
package database

import (
	_ "github.com/denisenkom/go-mssqldb"
)

type MSSQL struct{}
