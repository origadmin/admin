/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types implements the functions, types, and interfaces for the module.
package types

const (
	Invalid  = 0
	Enabled  = 1
	Disabled = 2

	Active   = Enabled
	Inactive = Disabled
	Frozen   = Disabled
)
