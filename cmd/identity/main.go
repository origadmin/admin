package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
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
	identityservice "origadmin/application/admin/internal/features/identity/service"
	confhelper "origadmin/application/admin/internal/helpers/conf"
)

var (
	// Name is the name of the compiled software.
	Name = "origadmin.service.identity"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
	// envPath is the path to the .env file.
	envPath = "resources/.env.identity"

	// flagconf is the config flag.
	flagconf string
)

func init() {
	// The config path should be the directory containing configuration files.
	// The default is empty, so we can detect if the user has provided it.
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf bootstrap.yaml")
}

// NewBootstrapOptions creates Kratos options from the PolicyBootstrap.
// This acts as a transformer for wire to inject the BeforeStart hook.
func NewBootstrapOptions(bootstrap *identityservice.PolicyBootstrap) []kratos.Option {
	return []kratos.Option{kratos.BeforeStart(bootstrap.Bootstrap)}
}

// NewApp creates a new Kratos application.
func NewApp(app *runtime.App, servers []transport.Server, opts []kratos.Option) *kratos.App {
	// Prepend the bootstrap options to any other options.
	log.SetLogger(app.Logger())
	return app.NewApp(servers, opts...)
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
	log.Infof("Starting %s %s (ID: %s)\n", rt.AppInfo().Name(), rt.AppInfo().Version(), rt.AppInfo().ID())

	// Get bootstrap config
	bootstrapConfig, ok := rt.StructuredConfig().(*conf.Config)
	if !ok {
		log.Fatalf("failed to get bootstrap config")
	}

	// wireApp now builds the entire application, including options.
	kratosApp, cleanupApp, err := wireApp(rt, bootstrapConfig)
	if err != nil {
		log.Fatalf("failed to wire app: %v", err)
	}
	defer cleanupApp()

	// Run the application. The BeforeStart hook is now injected via wire.
	if err := kratosApp.Run(); err != nil {
		log.Fatalf("app run failed: %v", err)
	}
}
