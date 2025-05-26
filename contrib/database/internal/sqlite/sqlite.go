/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package sqlite implements the functions, types, and interfaces for the module.
package sqlite

import (
	"os"
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

func MakeSourceDirectory(source string) {
	if strings.HasPrefix(source, "file://") {
		source = strings.TrimPrefix(source, "file://")
	}
	idx := strings.Index(source, "?")
	if idx > 0 {
		source = source[:idx]
	}
	dirs := strings.Split(source, "/")
	if len(dirs) > 1 {
		dirs = dirs[:len(dirs)-1]
		dir := strings.Join(dirs, "/")
		_, err := os.Stat(dir)
		if err != nil {
			os.MkdirAll(dir, 0755)
			return
		}
	}
}
