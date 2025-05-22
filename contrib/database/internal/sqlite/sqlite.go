/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package sqlite implements the functions, types, and interfaces for the module.
package sqlite

import (
	"strings"
)

const FKSuffix = "_fk=1"

func SourceForeignKeys(source string) string {
	// Check if the source already contains the FK parameter
	if strings.Contains(source, FKSuffix) {
		return source
	}

	// Check if the source already contains parameters
	if strings.Contains(source, "?") {
		// If parameters exist, append with &
		if !strings.HasSuffix(source, "&") {
			source += "&"
		}
		source += FKSuffix
	} else {
		// If no parameters exist, append with ?
		source += "?" + FKSuffix
	}
	return source
}
