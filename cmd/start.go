// Copyright (c) 2024 KasaAdmin. All rights reserved.

// Package cmd defines a CLI command to start a server with various flags and options, including the
// ability to run as a daemon.
// It includes functions to create and manage the command, as well as the logic to run the server.
// It also includes a function to create a new server instance and start it.
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"origadmin/backend/bootstrap"
	"origadmin/backend/internal/errors"
)

const (
	startRandom  = `random`
	startWorkdir = `workdir`
	startConfig  = `config`
	startStatic  = `static`
	startDaemon  = `daemon`
)

// StartCmd The function defines a CLI command to start a server with various flags and options, including the
// ability to run as a daemon.
func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the backend server",
		RunE: func(cmd *cobra.Command, args []string) error {
			workDir, _ := cmd.Flags().GetString(startWorkdir)
			staticDir, _ := cmd.Flags().GetString(startStatic)
			configs, _ := cmd.Flags().GetString(startConfig)
			random, _ := cmd.Flags().GetBool(startRandom)

			if daemon, _ := cmd.Flags().GetBool("daemon"); daemon {
				bin, err := filepath.Abs(os.Args[0])
				if err != nil {
					return err
				}

				cmdArgs := []string{"start"}
				cmdArgs = append(cmdArgs, "-d", strings.TrimSpace(workDir))
				cmdArgs = append(cmdArgs, "-c", strings.TrimSpace(configs))
				cmdArgs = append(cmdArgs, "-s", strings.TrimSpace(staticDir))
				command := exec.Command(bin, cmdArgs...)
				err = command.Start()
				if err != nil {
					return errors.RunError("Backend", "failed to start daemon thread", err)
				}

				pid := command.Process.Pid
				err = os.WriteFile(
					fmt.Sprintf("%s.lock", strings.ToLower(cmd.Root().Name())),
					[]byte(fmt.Sprintf("%d", pid)),
					0o600)
				if err != nil {
					//_, _ = ppfmt.Printf("failed to write pid file: %s \n", err.Error())

				}
				//_, _ = ppfmt.Printf("service %s daemon thread started with pid %d \n", config.C.General.AppName, pid)
				return nil
			}

			_ = os.WriteFile(
				fmt.Sprintf("%s.lock", strings.ToLower(cmd.Root().Name())),
				[]byte(fmt.Sprintf("%d", os.Getpid())),
				0o600)
			err := bootstrap.Run(cmd.Context(), bootstrap.Config{
				Random: random,
			})
			return err
		},
	}
	cmd.Flags().BoolP(startRandom, "r", false, "Start with random password")
	cmd.Flags().StringP(startWorkdir, "d", ".", "Working directory")
	cmd.Flags().StringP(startConfig, "c", "configs",
		"Runtime configuration files or directory (relative to workdir, multiple separated by commas)")
	cmd.Flags().StringP(startStatic, "s", "", "Static files directory")
	cmd.Flags().Bool(startDaemon, false, "Run as a daemon")
	return cmd
}
