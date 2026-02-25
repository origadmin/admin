/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package idutil implements the functions, types, and contracts for the module.
package idutil

import (
	"github.com/origadmin/toolkits/identifier"
	_ "github.com/origadmin/toolkits/identifier/snowflake"
)

var (
	generator = identifier.Get("snowflake")
)

func Gen() int64 {
	v, _ := generator.GenerateNumber()
	return v
}
