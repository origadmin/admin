// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package main is the main package
package main

import (
	"context"
	"fmt"
	"os"

	goversion "github.com/caarlos0/go-version"
	"github.com/spf13/cobra"

	"origadmin/backend/cmd"
	"origadmin/backend/helpers/interrupt"
	"origadmin/backend/internal/config"
)

// build tool goreleaser tags
//
//nolint:gochecknoglobals
var (
	version   = ""
	commit    = ""
	treeState = ""
	date      = ""
	builtBy   = ""
	debug     = false
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "origadmin",
	Short: "OrigAdmin Backend is a distributed backend management system with a focus on scalability, security, and flexibility.",
	Run: func(cmd *cobra.Command, args []string) {
		// Place your logic here
		// _ = cmd.Help()
	},
}

func init() {
	//cobra.AddTemplateFuncs(usage.Template)
	goinfo := buildVersion(version, commit, date, builtBy, treeState)
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.CompletionOptions.DisableNoDescFlag = true
	//rootCmd.SetUsageTemplate(resources.Usage)
	//rootCmd.SetHelpCommand(cmd.HelpCmd(goinfo))
	rootCmd.AddCommand(cmd.StartCmd())
	//rootCmd.AddCommand(cmd.StopCmd())
	//rootCmd.AddCommand(cmd.RestartCmd())
	//rootCmd.AddCommand(cmd.ConfigCmd())
	rootCmd.Version = goinfo.String()
}

// @title						OrigAdmin Backend API
// @version					v1.0.0
// @description				A
// @contact.name				OrigAdmin
// @contact.url				https://github.com/origadmin
// @license.name				MIT
// @license.url				https://github.com/origadmin/origadmin/blob/main/LICENSE.md
//
// @host						localhost:28080
// @basepath					/api/v1
// @schemes					http https
//
// @securitydefinitions.basic	Basic
//
// @securitydefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	Execute()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctx := interrupt.WithContext(context.Background())
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		fmt.Printf("Command executed with error:\n%v\n", err)
		os.Exit(1)
	}
}

func buildVersion(version, commit, date, builtBy, treeState string) goversion.Info {
	return goversion.GetVersionInfo(
		goversion.WithAppDetails(config.Application, config.Description, config.WebSite),
		func(i *goversion.Info) {
			i.ASCIIName = config.UI
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
