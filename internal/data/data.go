/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data provides the foundational infrastructure for data access.
package data

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/cenkalti/backoff/v5"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/data/storage"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	ProvideDatabase,
	NewStorageProvider,
	NewAdapterFromApp,
)

// systemUserID holds the ID of the system user. It is 0 if no system user is found.
var systemUserID atomic.Int64

// GetSystemUserID returns the ID of the system user.
// This is a thread-safe atomic load.
func GetSystemUserID() int64 {
	return systemUserID.Load()
}

// refreshSystemUserID attempts to fetch the system user ID from the database and updates the cache.
func refreshSystemUserID(ctx context.Context, database *ent.Database) bool {
	if database == nil {
		return false
	}
	systemUser, err := database.User(ctx).Query().Where(user.IsSystem(true)).Only(ctx)
	if err != nil {
		return false
	}
	systemUserID.Store(systemUser.ID)
	return true
}

// startSystemUserIDWatcher starts a background goroutine to poll for the system user ID using exponential backoff.
func startSystemUserIDWatcher(ctx context.Context, database *ent.Database, logger log.Logger) {
	logHelper := log.NewHelper(logger)
	go func() {
		// Create a strategy that starts at 1s and backs off to 30s
		b := backoff.NewExponentialBackOff()
		b.InitialInterval = 1 * time.Second
		b.MaxInterval = 30 * time.Second
		b.Reset()

		_, _ = backoff.Retry(ctx, func() (struct{}, error) {
			if id := GetSystemUserID(); id != 0 {
				return struct{}{}, nil // Already found, perhaps by startup grace period
			}

			if refreshSystemUserID(ctx, database) {
				logHelper.Infof("System user ID cached from background watcher: %d", GetSystemUserID())
				return struct{}{}, nil // Found and cached
			}

			return struct{}{}, fmt.Errorf("system user not found, retrying...")
		}, backoff.WithBackOff(b), backoff.WithMaxElapsedTime(0))
	}()
}

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

	logHelper.Infof("Database dialect: %s", db.Dialect())

	database := ent.NewDatabase(activeDB)
	// === DEBUG ===
	//database := ent.NewDatabase(entslog.New(activeDB, entslog.WithLogger(log.GetSlogLogger())))

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
	// If it's not found (e.g., because the initializer hasn't run yet), 
	// start a background watcher to keep trying without blocking the app startup.
	if !refreshSystemUserID(ctx, database) {
		logHelper.Info("System user not found on startup, starting background watcher.")
		startSystemUserIDWatcher(context.Background(), database, logger)
	} else {
		logHelper.Infof("System user ID cached on startup: %d", GetSystemUserID())
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
	d := &Data{
		DB:  database,
		log: logHelper,
	}
	return d, nil
}
