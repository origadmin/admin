/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mock implements the functions, types, and interfaces for the module.
package mock

import (
	"testing"

	"github.com/origadmin/runtime/bootstrap"

	"origadmin/application/admin/internal/loader"
	"origadmin/application/admin/internal/mods/system/server"
)

func GenerateTokenTest(t *testing.T) {
	bs, err := loader.LoadBootstrap(&loader.Bootstrap{
		Flags:      bootstrap.Flags{},
		WorkDir:    "",
		ConfigPath: "resources/configs/system",
		Env:        "",
		Daemon:     false,
	})
	if err != nil {
		t.Fatalf("failed to load bootstrap: %v", err)
	}
	v, err := server.NewSystemClient(bs, nil)
	if err != nil {
		t.Fatalf("failed to new system client: %v", err)
	}

}
