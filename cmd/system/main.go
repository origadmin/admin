/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"context"
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	_ "github.com/origadmin/contrib/config/consul"
	_ "github.com/origadmin/contrib/registry/consul"
	"github.com/origadmin/runtime"
	runtimebootstrap "github.com/origadmin/runtime/engine/bootstrap"
	"github.com/origadmin/runtime/log"
	_ "github.com/sqlite3ent/sqlite3"
	"origadmin/application/admin/internal/conf"
	_ "origadmin/application/admin/internal/data/entity/ent/runtime"
	"origadmin/application/admin/internal/features/system/service"
	confhelper "origadmin/application/admin/internal/helpers/conf"
)

var (
	// Name is the name of the compiled software.
	Name = "origadmin.service.system"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
	// envPath is the path to the .env file.
	envPath = "resources/.env.system"

	// flagconf is the config flag.
	flagconf string
)

func init() {
	// The config path should be the directory containing configuration files.
	// The default is empty, so we can detect if the user has provided it.
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf bootstrap.yaml")
}

func NewApp(app *runtime.App, servers []transport.Server, kratosOpts ...kratos.Option) *kratos.App {
	log.SetLogger(app.Logger())
	return app.NewApp(servers, kratosOpts...)
}

// NewBootstrapOptions creates a new bootstrap options.
func NewBootstrapOptions(bootstrap *service.PolicyBootstrap) []kratos.Option {
	return []kratos.Option{
		kratos.BeforeStart(func(ctx context.Context) error {
			return bootstrap.Bootstrap(ctx)
		}),
	}
}

func main() {
	if err := godotenv.Load(envPath); err != nil {
		wd, _ := os.Getwd()
		log.Warnf("godotenv: failed to load '%s' (PWD: %s): %v", envPath, wd, err)
	}

	flag.Parse()

	confPath := confhelper.FindConfPath(flagconf)
	if confPath == "" {
		log.Fatalf("Could not find configuration file. Searched -conf flag, executable path, and development path.")
	}

	// Log the config path for debugging
	log.Infof("Loading configuration from: %s\n", confPath)

	// NewFromBootstrap handles config loading, logging, and container setup.
	rt := runtime.New(Name, Version)
	err := rt.Load(confPath, runtimebootstrap.WithConfigTransformer(conf.New()))
	if err != nil {
		log.Fatalf("failed to create runtime: %v", err)
	}
	defer rt.Config().Close()
	rt.ShowAppInfo()

	// Get bootstrap config
	bootstrapConfig, ok := rt.StructuredConfig().(*conf.Config)
	if !ok {
		log.Fatalf("failed to get bootstrap config")
	}
	// wireApp now takes the runtime instance and builds the kratos app.
	app, cleanupApp, err := wireApp(rt, bootstrapConfig)
	if err != nil {
		log.Fatalf("failed to wire app: %v", err)
	}
	defer cleanupApp()

	// Run the application
	if err := app.Run(); err != nil {
		log.Fatalf("app run failed: %v", err)
	}
}
