/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package id implements the functions, types, and interfaces for the module.
package id

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
