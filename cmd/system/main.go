/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"log/slog"
	"os"

	goversion "github.com/caarlos0/go-version"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/codec/toml"

	_ "github.com/origadmin/backend/internal/data/entity/ent/runtime" // Updated import
	"origadmin/application/admin/cmd/internal/start" // Import the start command
)

// go build -ldflags "-X main.Version=vx.y.z -X main.Name=origadmin.service.system.v1"
var (
	// Name is the Name of the compiled software.
	Name = "origadmin.service.system.v1"
	// Version is the Version of the compiled software.
	Version = "v1.0.0"
	// flags are the bootstrap flags.
	flags = bootstrap.New()

	version   = ""
	commit    = ""
	treeState = ""
	date      = ""
	builtBy   = ""
)

func buildVersion(version, commit, date, builtBy, treeState string) goversion.Info {
	return goversion.GetVersionInfo(
		goversion.WithAppDetails(Name, "System Service", ""), // Use Name for app details
		func(i *goversion.Info) {
			if commit != "" {
				i.GitCommit = commit
			}
			if version != "" {
				i.GitVersion = version
			}
			if treeState != "" {
				i.GitTreeState = treeState
			}
			if date != "" {
				i.BuildDate = date
			}
			if builtBy != "" {
				i.BuiltBy = builtBy
			}
		},
	)
}

func init() {
	encoding.RegisterCodec(toml.Codec)
	flags.SetServiceInfo(Name, Version)
}

func main() {
	// Initialize cobra command for the system service
	rootCmd := start.Cmd()
	rootCmd.Use = "system"
	rootCmd.Short = "System service for OrigAdmin backend."

	info := buildVersion(version, commit, date, builtBy, treeState)
	rootCmd.Version = info.String()

	if err := rootCmd.Execute(); err != nil {
		slog.Error("failed to execute system service command", "error", err)
		os.Exit(1)
	}
}
