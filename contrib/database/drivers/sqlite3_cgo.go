//go:build sqlite3 && cgo

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package drivers is the database client wrapper
package drivers

import (
	_ "github.com/mattn/go-sqlite3"
)

type SQLite3Cgo struct{}
