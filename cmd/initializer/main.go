package main

import (
	"context"
	"flag"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	_ "github.com/origadmin/contrib/config/consul"
	_ "github.com/origadmin/contrib/registry/consul"
	"github.com/origadmin/runtime"
	runtimebootstrap "github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/log"
	_ "github.com/sqlite3ent/sqlite3"
	"origadmin/application/admin/internal/conf"
	_ "origadmin/application/admin/internal/data/entity/ent/runtime"
	confhelper "origadmin/application/admin/internal/helpers/conf"
)

var (
	// Name is the name of the compiled software.
	Name = "origadmin.job.initializer"
	// Version is the version of the compiled software.
	Version = "v1.0.0"

	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf bootstrap.yaml")
}

func main() {
	_ = godotenv.Load("resources/.env.initializer")

	flag.Parse()

	confPath := confhelper.FindConfPath(flagconf)
	if confPath == "" {
		log.Fatalf("Could not find configuration file. Searched -conf flag, executable path, and development path.")
	}

	log.Infof("Loading configuration from: %s\n", confPath)

	rt := runtime.New(Name, Version)
	err := rt.Load(confPath, runtimebootstrap.WithConfigTransformer(conf.New()))
	if err != nil {
		log.Fatalf("failed to create runtime: %v", err)
	}
	defer rt.Config().Close()
	log.Infof("Starting %s %s (ID: %s)\n", rt.AppInfo().Name(), rt.AppInfo().Version(), rt.AppInfo().ID())

	bootstrapConfig, ok := rt.StructuredConfig().(*conf.Config)
	if !ok {
		log.Fatalf("failed to get bootstrap config")
	}

	initSvc, cleanup, err := wireApp(rt, bootstrapConfig)
	if err != nil {
		log.Fatalf("failed to wire app: %v", err)
	}
	defer cleanup()

	// Execute the initialization logic
	if err := initSvc.Init(context.Background()); err != nil {
		log.Fatalf("initialization failed: %v", err)
	}

	log.Info("All initialization tasks completed successfully.")
}
