//go:build mysql

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package drivers is the database client wrapper
package drivers

import (
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct{}
