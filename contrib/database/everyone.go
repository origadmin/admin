//go:build !sqlite3 && !mysql && !postgres && !sqlserver && !mssql && !pgx

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package database is the database client wrapper
package database

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/sqlite3ent/sqlite3"
)

// EveryOne ...
type EveryOne struct{}
