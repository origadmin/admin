//go:build debug

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package start implements the functions, types, and interfaces for the module.
package start

import (
	"fmt"
)

func init() {
	fmt.Println("debug mode")
	flags.Env = "debug"
}
