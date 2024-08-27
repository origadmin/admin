// Copyright (c) 2024 OrigAdmin. All rights reserved.

//go:generate swag init --parseDependency --generalInfo ./main.go --output ./resources/docs
// #go:generate docker run --rm -v resources:/local openapitools/openapi-generator-cli generate -i /local/docs/swagger.yaml -g openapi -o /local/docs/v3

// Package main is the main package
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// VERSION usage: go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "v1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Admin",
	Short: "The Admin is a control backend for OrigAdmin",
	Run: func(cmd *cobra.Command, args []string) {
		// Place your logic here
		// _ = cmd.Help()
	},
}

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.CompletionOptions.DisableNoDescFlag = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.Version = fmt.Sprintf(
		"%s on %s/%s",
		VERSION, runtime.GOOS, runtime.GOARCH,
	)
}

// @title						Admin
// @version					v1.0.0
// @description				A control backend for OrigAdmin.
// @contact.name				OrigAdmin
// @contact.url				https://github.com/origadmin
// @license.name				MIT
// @license.url				https://github.com/origadmin/admin/blob/main/LICENSE.md
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
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Command executed with error:\n%v\n", err)
		os.Exit(1)
	}
}
