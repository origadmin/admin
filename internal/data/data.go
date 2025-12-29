/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data provides the foundational infrastructure for data access.
package data

import (
	"context"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/data/storage"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/data/entity/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, ProvideDatabase, ProvideCache)

// Data encapsulates the core data access components.
type Data struct {
	DB  *ent.Database
	RDB *redis.Client
	log *log.Helper
}

// ProvideDatabase extracts and provides the *ent.Database from the *Data object.
func ProvideDatabase(d *Data) *ent.Database {
	return d.DB
}

// ProvideCache extracts and provides the *redis.Client from the *Data object.
func ProvideCache(d *Data) *redis.Client {
	return d.RDB
}

// NewData creates a new Data instance, which encapsulates the core database and cache objects.
func NewData(rt *runtime.App, conf *conf.Config) (*Data, func(), error) {
	logHelper := log.NewHelper(rt.Logger())

	provider, err := storage.New(rt.StructuredConfig())
	if err != nil {
		return nil, nil, err
	}

	// --- Database ---
	db, err := provider.DefaultDatabase()
	if err != nil {
		return nil, nil, err
	}
	activeDB := entsql.OpenDB(db.Dialect(), db.DB())
	database := ent.NewDatabase(ent.Driver(activeDB))

	// Run the auto migration tool.
	if err := database.Migration(context.Background(),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
		schema.WithForeignKeys(false),
	); err != nil {
		logHelper.Fatalf("failed creating schema resources: %v", err)
	}

	// --- Cache ---
	cache, err := provider.DefaultCache()
	if err != nil {
		return nil, nil, err
	}
	rdb := cache.Redis()

	d := &Data{
		DB:  database,
		RDB: rdb,
		log: logHelper,
	}

	cleanup := func() {
		logHelper.Info("closing the data resources")
		if d.DB != nil {
			if err := d.DB.Client(context.Background()).Close(); err != nil {
				logHelper.Errorf("failed to close ent client: %v", err)
			}
		}
		if activeDB != nil {
			if err := activeDB.Close(); err != nil {
				logHelper.Errorf("failed to close database: %v", err)
			}
		}
		if d.RDB != nil {
			if err := d.RDB.Close(); err != nil {
				logHelper.Errorf("failed to close redis client: %v", err)
			}
		}
	}
	return d, cleanup, nil
}
