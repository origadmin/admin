//go:build postgres && !pgx

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package drivers is the database client wrapper
package drivers

import (
	_ "github.com/lib/pq"
)

type Postgres struct{}
