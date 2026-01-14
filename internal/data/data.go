/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data provides the foundational infrastructure for data access.
package data

import (
	"context"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/data/storage"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, ProvideDatabase, NewStorageProvider, NewAdapter)

// SystemUserID holds the ID of the system user. It is 0 if no system user is found.
var SystemUserID int64

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
	database := ent.NewDatabase(
		ent.Driver(activeDB),
		ent.Debug(),
	)
	ctx := context.Background()
	// === The migration logic is moved here ===
	if err := database.Migration(ctx,
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
		schema.WithForeignKeys(false),
	); err != nil {
		return nil, nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	// Fetch and cache the system user ID on startup.
	systemUser, err := database.User(ctx).Query().Where(user.IsSystem(true)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			logHelper.Info("System user not found, this is expected on initial startup.")
		} else {
			return nil, nil, fmt.Errorf("failed to query system user: %w", err)
		}
	}
	if systemUser != nil {
		SystemUserID = systemUser.ID
		logHelper.Infof("System user ID cached: %d", SystemUserID)
	}

	return database, func() {
		if database != nil {
			if err := database.Client(ctx).Close(); err != nil {
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
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	d := &Data{
		DB:  database,
		log: logHelper,
	}
	return d, nil
}
