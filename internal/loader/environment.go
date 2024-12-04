/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"os"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

func SetupEnv(args map[string]string, s string) {
	for k, v := range args {
		key := strings.Join([]string{k, s}, "_")
		err := os.Setenv(key, v)
		if err != nil {
			log.Warnf("failed to set environment variable: %s \n", err.Error())
		}
	}
}
