/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package db

import (
	"strings"

	"entgo.io/ent/dialect/sql"
)

// OrderBy dynamically builds a list of order functions from a slice of strings.
// Each string can be in the format "field_name" (for ascending) or "field_name,desc" (for descending).
// This function is designed to be perfectly compatible with Ent's `order()` method.
func OrderBy[T selectable](fields []string, orders ...T) []T {
	for _, field := range fields {
		parts := strings.Split(field, ",")
		fieldName := parts[0]
		var orderOpt sql.OrderTermOption

		if len(parts) > 1 {
			switch strings.ToLower(parts[1]) {
			case "desc":
				orderOpt = sql.OrderDesc()
			default:
				orderOpt = sql.OrderAsc()
			}
		} else {
			orderOpt = sql.OrderAsc()
		}

		// This conversion is specific to Ent's sql.OrderFunc and is intentionally kept
		// to maintain compatibility.
		orders = append(orders, sql.OrderByField(fieldName, orderOpt).ToFunc())
	}
	return orders
}

func OrderByField[T selectable](fieldName string, desc bool) T {
	var orderOpt sql.OrderTermOption
	if desc {
		orderOpt = sql.OrderDesc()
	} else {
		orderOpt = sql.OrderAsc()
	}
	return sql.OrderByField(fieldName, orderOpt).ToFunc()
}
