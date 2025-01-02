/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/internal/configs"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bootstrap *configs.Bootstrap, l log.KLogger, ss ...service.OptionSetting) *service.HTTPServer {
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetService(), ss...)
	if err != nil {
		panic(err)
	}
	return srv
}
