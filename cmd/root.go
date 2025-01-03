/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package cmd represents the base command when called without any subcommands
package cmd

import (
	"os"

	goversion "github.com/caarlos0/go-version"
	"github.com/spf13/cobra"

	"origadmin/application/admin/cmd/internal/start"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin is a distributed backend management system with a focus on scalability, security, and flexibility for OrigAdmin.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(info goversion.Info) {
	rootCmd.Version = info.String()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.admin.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(start.Cmd())
}
