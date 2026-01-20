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
	generator = New()
)

func New() identifier.Generator[int64] {
	return identifier.Get[int64]("snowflake")
}

func Gen() int64 {
	return generator.Generate()
}
