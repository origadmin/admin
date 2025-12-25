/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package db provides common helpers for database operations,
// currently focusing on query construction for Ent.
package db

import (
	"entgo.io/ent/dialect/sql"
)

// This file is intentionally left blank after refactoring.
// Common interfaces or future helpers can be added here.

type selectable interface {
	~func(*sql.Selector)
}
