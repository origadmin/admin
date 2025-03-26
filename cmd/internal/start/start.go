/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package start is the start command for the application.
package start

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/golang-cz/devslog"
	_ "github.com/origadmin/contrib/consul/config"
	_ "github.com/origadmin/contrib/consul/registry"
	_ "github.com/origadmin/contrib/database"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/log"
	kslog "github.com/origadmin/slog-kratos"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/sloge"
	"github.com/spf13/cobra"

	"origadmin/application/admin/helpers/command"
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
	flags = bootstrap.Bootstrap{Env: "release"}
)

var cmd = &cobra.Command{
	Use:   "start",
	Short: "start the server",
	RunE:  startCommandRun,
}

func RandomID() string {
	id, err := os.Hostname()
	if err != nil {
		id = "unknown"
	}
	return id + "." + fmt.Sprintf("%08d", time.Now().UnixNano()%(1<<32))
}
func init() {
	//fmt.Println("total env: ", os.Environ(), len(os.Environ()))
	flags.Flags.ID = RandomID()
	flags.SetFlags(Name, Version)
}

// Cmd The function defines a CLI command to start a server with various flags and options, including the
// ability to run as a daemon.
func Cmd() *cobra.Command {
	cmd.Flags().BoolP(startRandom, "r", false, "start with random password")
	//cmd.Flags().StringP(startWorkDir, "d", ".", "working directory")
	//cmd.Flags().StringP(startWorkDir, "d", ".", "working directory")
	cmd.Flags().StringP(startConfig, "c", "bootstrap.toml",
		"runtime configuration files or directory (relative to workdir, multiple separated by commas)")
	cmd.Flags().StringP(startStatic, "s", "", "static files directory")
	cmd.Flags().BoolP(startDebug, "d", false, "set debug mode, eg: --debug")
	cmd.Flags().Bool(startDaemon, false, "run as a daemon")
	return cmd
}

func startCommandRun(cmd *cobra.Command, args []string) error {
	debug, _ := cmd.Flags().GetBool(startDebug)
	if debug {
		flags.Env = "debug"
		flags.WorkDir = "resources/configs"
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	staticDir, _ := cmd.Flags().GetString(startStatic)
	flags.ConfigPath, _ = cmd.Flags().GetString(startConfig)
	//random, _ := cmd.Flags().GetBool(startRandom)
	slogInstance := sloge.New(sloge.WithFile("logs/admin.log"), sloge.WithDevConfig(&sloge.DevConfig{
		HandlerOptions:    &slog.HandlerOptions{Level: slog.LevelDebug},
		MaxSlicePrintSize: 50,
		SortKeys:          false,
		TimeFormat:        "[15:04:05]",
		DebugColor:        devslog.Blue,
		InfoColor:         devslog.Green,
		WarnColor:         devslog.Yellow,
		ErrorColor:        devslog.Red,
	}))
	l := log.With(kslog.NewLogger(kslog.WithLogger(slogInstance)),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", flags.ID(),
		"service.name", flags.ServiceName(),
		"service.version", flags.Version(),
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(l)
	//path := filepath.Join(flags.WorkDir, flags.ConfigPath)
	//envpath := filepath.Join(flags.WorkDir, flags.EnvPath)
	log.Infow("msg", "start info", "workpath", flags.WorkPath(), startStatic, staticDir)
	//env, _ := bootstrap.LoadEnv(envpath)
	//bs, err := bootstrap.FromLocalPath(flags.ServiceName, path, l)
	//if err != nil {
	//	return errors.Wrap(err, "load config error")
	//}
	//src := loader.LoadSourceFiles(flags.WorkDir, flags.ConfigPath)
	//source := loader.FileSourceConfig(flags.WorkPath())
	if daemon, _ := cmd.Flags().GetBool("daemon"); daemon {
		bin, err := filepath.Abs(os.Args[0])
		if err != nil {
			log.Errorf("failed to get absolute path for cmd: %s \n", err.Error())
			return err
		}

		cmdArgs := []string{"start"}
		cmdArgs = append(cmdArgs, "-d", strings.TrimSpace(flags.WorkDir))
		cmdArgs = append(cmdArgs, "-c", strings.TrimSpace(flags.ConfigPath))
		cmdArgs = append(cmdArgs, "-s", strings.TrimSpace(staticDir))
		_, _ = fmt.Printf("execute cmd: %s %s \n", bin, strings.Join(cmdArgs, " "))
		cmd := exec.Command(bin, cmdArgs...)
		err = cmd.Start()
		if err != nil {
			_, _ = fmt.Printf("failed to start daemon thread: %s \n", err.Error())
			return err
		}

		pid := cmd.Process.Pid
		log.Errorf("service %s daemon thread started with pid %d \n", flags.ServiceName(), pid)
		return nil
	}
	bs, err := loader.LoadBootstrap(&flags)
	if err != nil {
		return errors.Wrap(err, "load config error")
	}
	if bs == nil {
		return fmt.Errorf("bootstrap config not found")
	}
	if bs.CryptoType == "argon2" {
		hash.UseArgon2()
	}

	//log.Infof("bootstrap: %+v", loader.PrintString(bs))
	lockfile := fmt.Sprintf("%s.lock", command.ToLower(cmd))
	if err = os.WriteFile(lockfile, []byte(fmt.Sprintf("%d", os.Getpid())), 0o600); err == nil {
		defer os.Remove(lockfile)
	}
	//engine := gin.New()
	//info to ctx
	app, cleanup, err := buildInjectors(cmd.Context(), bs, l)
	if err != nil {
		return err
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		return err
	}
	return nil
}

func NewApp(ctx context.Context, injector *loader.InjectorClient) *kratos.App {
	opts := []kratos.Option{
		kratos.ID(flags.ServiceID()),
		kratos.Name(flags.ServiceName()),
		kratos.Version(flags.Version()),
		kratos.Metadata(map[string]string{}),
		kratos.Context(ctx),
		kratos.Signal(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT),
		kratos.Logger(injector.Logger),
		kratos.Server(injector.Server),
		//kratos.Server(injector.ServerGINS),
	}
	//err := loader.InjectorGinServer(injector)
	//if err != nil {
	//	log.Errorf("injector gin server error: %v", err)
	//	os.Exit(1)
	//}
	if flags.Env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infow("msg", "GIN route", "method", httpMethod, "path", absolutePath, "operation", handlerName, "handlers", nuHandlers)
	}

	return kratos.New(opts...)
}
