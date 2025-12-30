/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data provides the foundational infrastructure for data access.
package data

import (
	"context"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/data/storage"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, ProvideDatabase, NewStorageProvider, NewAdapter)

// Data encapsulates the core data access components.
type Data struct {
	DB  *ent.Database
	log *log.Helper
}

// ProvideDatabase extracts and provides the *ent.Database from the *Data object.
func ProvideDatabase(pv storage.Provider, logger log.Logger) (*ent.Database, func(), error) {
	logHelper := log.NewHelper(logger)
	db, err := pv.DefaultDatabase()
	if err != nil {
		return nil, nil, err
	}

	activeDB := entsql.OpenDB(db.Dialect(), db.DB())
	database := ent.NewDatabase(ent.Driver(activeDB))
	return database, func() {
		if database != nil {
			if err := database.Client(context.Background()).Close(); err != nil {
				logHelper.Errorf("failed to close ent client: %v", err)
			}
		}
		if activeDB != nil {
			if err := activeDB.Close(); err != nil {
				logHelper.Errorf("failed to close database: %v", err)
			}
		}
	}, nil
}

// NewStorageProvider creates a new storage provider from the application's structured config.
func NewStorageProvider(rt *runtime.App) (storage.Provider, error) {
	return storage.New(rt.StructuredConfig())
}

// NewData creates a new Data instance, which encapsulates the core database object.
func NewData(database *ent.Database, logger log.Logger) (*Data, error) {
	logHelper := log.NewHelper(logger)
	// Run the auto migration tool.
	if err := database.Migration(context.Background(),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
		schema.WithForeignKeys(false),
	); err != nil {
		logHelper.Fatalf("failed creating schema resources: %v", err)
	}
	ent.Debug()
	d := &Data{
		DB:  database,
		log: logHelper,
	}
	return d, nil
}
