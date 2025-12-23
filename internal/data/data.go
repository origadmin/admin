/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data implements the functions, types, and interfaces for the module.
package data

import (
	"context"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/data/storage"
	"github.com/origadmin/runtime/interfaces"
	ifacestorage "github.com/origadmin/runtime/interfaces/storage"
	"github.com/origadmin/runtime/log"

	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/data/entity/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data encapsulates ent client and cache.
type Data struct {
	database *ent.Database
	cache    ifacestorage.Cache
	provider storage.Provider
	config   interfaces.StructuredConfig
	Log      *log.Helper
}

// NewData creates a new Data instance.
func NewData(rt *runtime.App, conf *conf.Config) (*Data, func(), error) {
	logHelper := log.NewHelper(rt.Logger())

	provider, err := storage.New(rt.StructuredConfig())
	if err != nil {
		return nil, nil, err
	}

	db, err := provider.DefaultDatabase()
	if err != nil {
		return nil, nil, err
	}

	activeDB := entsql.OpenDB(db.Dialect(), db.DB())
	database := ent.NewDatabase(ent.Driver(activeDB))

	// Run the auto migration tool.
	// Note: context.Background() is used here as the schema creation is a one-time setup.
	if err := database.Client(context.Background()).Schema.Create(context.Background(),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
		schema.WithForeignKeys(false)); err != nil {
		logHelper.Fatalf("failed creating schema resources: %v", err)
	}

	cache, err := provider.DefaultCache()
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		logHelper.Info("closing the data resources")
		if database != nil {
			if err := database.Client(context.Background()).Close(); err != nil {
				logHelper.Errorf("failed to close ent client: %v", err)
			}
		}

	}

	return &Data{
		config:   rt.StructuredConfig(),
		provider: provider,
		database: database,
		cache:    cache,
		Log:      logHelper,
	}, cleanup, nil
}

// DB returns the ent.Client instance.
func (d *Data) DB() *ent.Database {
	return d.database
}
