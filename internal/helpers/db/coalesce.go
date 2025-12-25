/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package db provides common helpers for database operations,
// currently focusing on query construction for Ent.
package db

import (
	"entgo.io/ent/dialect/sql"
)

// CoalesceMax returns the max value of the given field.
func CoalesceMax(field string) func(selector *sql.Selector) string {
	return func(selector *sql.Selector) string {
		fn := sql.Func{}
		fn.Append(func(builder *sql.Builder) {
			builder.WriteString("COALESCE")
			builder.Wrap(func(b *sql.Builder) {
				b.Ident(sql.Max(selector.C(field))).Comma().WriteByte('0')
			})
		})
		return fn.String()
	}
}
