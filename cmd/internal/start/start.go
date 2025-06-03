/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package start is the start command for the application.
package start

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/spf13/cobra"

	_ "origadmin/application/admin/contrib/consul/config"
	_ "origadmin/application/admin/contrib/consul/registry"
	_ "origadmin/application/admin/contrib/database"
	"origadmin/application/admin/internal/loader"
)

const (
	startRandom  = `random`
	startWorkDir = `workdir`
	startConfig  = `config`
	startStatic  = `static`
	startDaemon  = `daemon`
	startDebug   = `debug`
)

var (
	// Name is the name of the compiled software.
	Name = "origadmin.server.v1.admin"
	// Version is the Version of the compiled software.
	Version = "v1.0.0"
	// flags are the bootstrap flags.
	flags = bootstrap.New()
)

var cmd = &cobra.Command{
	Use:   "start",
	Short: "start the server",
	RunE:  startCommandRun,
}

func init() {
	flags.SetServiceInfo(Name, Version)
}

// Cmd The function defines a CLI command to start a server with various flags and options, including the
// ability to run as a daemon.
func Cmd() *cobra.Command {
	cmd.Flags().BoolP(startRandom, "r", false, "start with random password")
	cmd.Flags().StringP(startConfig, "c", "bootstrap.toml",
		"runtime configuration files or directory (relative to workdir, multiple separated by commas)")
	cmd.Flags().StringP(startStatic, "s", "", "static files directory")
	cmd.Flags().BoolP(startDebug, "d", false, "set debug mode, eg: --debug")
	cmd.Flags().Bool(startDaemon, false, "run as a daemon")
	return cmd
}

func startCommandRun(cmd *cobra.Command, args []string) error {
	if err := loader.Bootstrap(cmd.Context(), flags, buildInjectors); err != nil {
		return err
	}
	//var registrar registry.KRegistrar
	//if flags.IsMainService() {
	//	registrar, _ = registry.NewConsulRegistrar()
	//}
	//
	//buildInjectors()
	//
	//r.CreateApp(cmd.Context())
	//
	//// 组合使用配置和服务
	//appInstance := loader.NewApp(cmd.Context(), loader.AppOptions{
	//	Name:    bs.ServiceName,
	//	Version: flags.Version(),
	//	Server:  grpcServer,
	//})
	return nil
}

func NewAppProvider(r runtime.Runtime, injector *loader.InjectorClient) *kratos.App {
	return r.CreateApp(injector.Server)
}
