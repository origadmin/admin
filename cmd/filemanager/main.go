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
	runtimebootstrap "github.com/origadmin/runtime/engine/bootstrap"
	"github.com/origadmin/runtime/log"
	_ "github.com/sqlite3ent/sqlite3"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	_ "origadmin/application/admin/internal/data/entity/ent/runtime"
	confhelper "origadmin/application/admin/internal/helpers/conf"
	_ "origadmin/application/admin/internal/helpers/providers"
)

var (
	// Name is the name of the compiled software.
	Name = "origadmin.service.filemanager"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
	// envPath is the path to the .env file.
	envPath = "resources/.env.filemanager"

	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf bootstrap.yaml")
}

func NewApp(app *runtime.App, servers []transport.Server) *kratos.App {
	log.SetLogger(app.Logger())
	return app.NewApp(servers)
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

	log.Infof("Loading configuration from: %s\n", confPath)
	rt := runtime.New(Name, Version)
	if err := rt.Load(confPath, runtimebootstrap.WithConfigTransformer(conf.Transformer)); err != nil {
		log.Fatalf("failed to create runtime: %v", err)
	}
	defer rt.Config().Close()
	rt.ShowAppInfo()

	// Get bootstrap config
	bootstrapConfig, ok := rt.BusinessConfig().(*confpb.Bootstrap)
	if !ok {
		log.Fatalf("failed to get bootstrap config")
	}

	app, cleanupApp, err := wireApp(rt, bootstrapConfig)
	if err != nil {
		log.Fatalf("failed to wire app: %v", err)
	}
	defer cleanupApp()

	if err := app.Run(); err != nil {
		log.Fatalf("app run failed: %v", err)
	}
}
