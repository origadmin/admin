/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package main is the main package
package main

import (
	goversion "github.com/caarlos0/go-version"

	"origadmin/application/admin/cmd"
	"origadmin/application/admin/internal/loader"
)

var (
	version   = ""
	commit    = ""
	treeState = ""
	date      = ""
	builtBy   = ""
	debug     = false
)

func buildVersion(version, commit, date, builtBy, treeState string) goversion.Info {
	return goversion.GetVersionInfo(
		goversion.WithAppDetails(loader.Application, loader.Description, loader.WebSite),
		func(i *goversion.Info) {
			i.ASCIIName = loader.UI
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

// @title						OrigAdmin Backend API
// @version					v1.0.0
// @description				A distributed backend management system with a focus on scalability, security, and flexibility.
// @contact.name				OrigAdmin
// @contact.url				https://github.com/origadmin
// @license.name				MIT
// @license.url				https://github.com/origadmin/origadmin/blob/main/LICENSE.md
//
// @host						localhost:10080
// @basepath					/api/v1
// @schemes					http https
//
// @securitydefinitions.basic	Basic
//
// @securitydefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	cmd.Execute(buildVersion(version, commit, date, builtBy, treeState))
}
