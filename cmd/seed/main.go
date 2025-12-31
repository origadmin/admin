/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// The seed command is a one-off task to initialize the database with default data.
package main

import (
	"flag"

	"github.com/joho/godotenv"
	_ "github.com/sqlite3ent/sqlite3" // Import for sqlite3 driver

	"github.com/origadmin/runtime"
	runtimebootstrap "github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/conf"
	_ "origadmin/application/admin/internal/data/entity/ent/runtime"
	confhelper "origadmin/application/admin/internal/helpers/conf"
)

var (
	// Name is the name of the compiled software.
	Name = "origadmin.task.seed"
	// Version is the version of the compiled software.
	Version = "v1.0.0"

	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf bootstrap.yaml")
}

func main() {
	_ = godotenv.Load("resources/.env.system")
	flag.Parse()

	confPath := confhelper.FindConfPath(flagconf)
	if confPath == "" {
		log.Fatal("Could not find configuration file.")
	}

	rt := runtime.New(Name, Version)
	if err := rt.Load(confPath, runtimebootstrap.WithConfigTransformer(conf.New())); err != nil {
		log.Fatalf("failed to create runtime: %v", err)
	}
	defer rt.Config().Close()

	bootstrapConfig, ok := rt.StructuredConfig().(*conf.Config)
	if !ok {
		log.Fatalf("failed to get bootstrap config")
	}

	// wireApp builds the dependencies needed for the seed task.
	s, cleanup, err := wireApp(rt, bootstrapConfig)
	if err != nil {
		log.Fatalf("failed to wire app: %v", err)
	}
	defer cleanup()

	// Execute the seed task.
	if err := s.Run(); err != nil {
		log.Fatalf("seed task failed: %v", err)
	}

	log.Info("seed task completed successfully.")
}
