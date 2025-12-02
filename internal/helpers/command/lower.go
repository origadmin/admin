/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package command implements the functions, types, and interfaces for the module.
package command

import (
	"strings"

	"github.com/spf13/cobra"
)

func ToLower(cmd *cobra.Command) string {
	return strings.ToLower(cmd.Root().Name())
}
