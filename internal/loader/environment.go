/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/origadmin/toolkits/env"
)

func SetupEnv(args map[string]string, prefix string) {
	for k, v := range args {
		key := env.Var(prefix, k)
		log.Infof("set environment variable: %s=%s", key, v)
		err := env.SetEnv(key, v)
		if err != nil {
			log.Warnf("failed to set environment variable: %s", err.Error())
		}
	}
}
