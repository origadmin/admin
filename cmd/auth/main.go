package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/joho/godotenv"

	_ "github.com/sqlite3ent/sqlite3" // Import for sqlite3 driver

	_ "github.com/origadmin/contrib/config/consul"
	_ "github.com/origadmin/contrib/registry/consul"
	"github.com/origadmin/runtime"
	runtimebootstrap "github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/conf"
	_ "origadmin/application/admin/internal/data/entity/ent/runtime"
	confhelper "origadmin/application/admin/internal/helpers/conf"
)

var (
	// Name is the name of the compiled software.
	Name = "origadmin.service.auth"
	// Version is the version of the compiled software.
	Version = "v1.0.0"

	// flagconf is the config flag.
	flagconf string
)

func init() {
	// The config path should be the directory containing configuration files.
	// The default is empty, so we can detect if the user has provided it.
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf bootstrap.yaml")
}

func NewApp(app *runtime.App, servers []transport.Server) *kratos.App {
	return app.NewApp(servers)
}

func main() {
	// Load .env file for local development from resources directory.
	// It's safe to ignore the error, as the file may not exist in production.
	_ = godotenv.Load("resources/.env.auth")

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
	log.Infof("Starting %s %s (ID: %s)\n", rt.AppInfo().Name(), rt.AppInfo().Version(), rt.AppInfo().ID())

	// Get bootstrap config
	bootstrapConfig, ok := rt.StructuredConfig().(*conf.Config)
	if !ok {
		log.Fatalf("failed to get bootstrap config")
	}

	// wireApp now takes the runtime instance and builds the kratos app.
	kratosApp, cleanupApp, err := wireApp(rt, bootstrapConfig)
	if err != nil {
		log.Fatalf("failed to wire app: %v", err)
	}
	defer cleanupApp()

	// Run the application
	if err := kratosApp.Run(); err != nil {
		log.Fatalf("app run failed: %v", err)
	}
}
