/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package conf provides helper functions for application startup.
package conf

import (
	"log"
	"os"
	"path/filepath"
)

// FindConfPath searches for the configuration file in a prioritized order.
//
// It checks the following locations in order:
// 1. The path provided by the -conf flag.
// 2. Deployed environment: 'configs/bootstrap.yaml' or 'bootstrap.yaml' relative to the executable.
// 3. Development environment: './resources/configs/bootstrap.yaml' relative to the project root.
//
// The `flagPath` argument should be the value from a command-line flag.
// It returns the found path or an empty string if not found.
func FindConfPath(flagPath string) string {
	// 1. Highest Priority: User-provided flag.
	if flagPath != "" {
		log.Printf("Using configuration from -conf flag: %s", flagPath)
		return flagPath
	}

	wd, _ := os.Getwd()
	log.Printf("Current working directory: %s", wd)

	// 2. Second Priority: Deployed environment (relative to executable).
	exec, err := os.Executable()
	if err == nil {
		execDir := filepath.Dir(exec)
		deployPaths := []string{
			filepath.Join(execDir, "configs", "bootstrap.yaml"),
			filepath.Join(execDir, "bootstrap.yaml"),
		}
		for _, p := range deployPaths {
			if _, err := os.Stat(p); err == nil {
				log.Printf("Found configuration in deployed environment: %s", p)
				return p
			}
		}
	}

	// 3. Lowest Priority: Development environment (relative to project root).
	devPath := "./resources/configs/bootstrap.yaml"
	if _, err := os.Stat(devPath); err == nil {
		log.Printf("Found configuration in development environment: %s", devPath)
		return devPath
	}

	return "" // Not found
}
