//go:build mssql || sqlserver

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package drivers is the database client wrapper
package drivers

import (
	_ "github.com/denisenkom/go-mssqldb"
)

type MSSQL struct{}
