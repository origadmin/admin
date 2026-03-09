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

	"github.com/cenkalti/backoff/v5"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	ProvideDatabase,
	NewAdapter,
)

// systemUserID holds the ID of the system user. It is 0 if no system user is found.
var systemUserID atomic.Int64

// GetSystemUserID returns the ID of the system user.
func GetSystemUserID() int64 {
	return systemUserID.Load()
}

// refreshSystemUserID checks and caches the system user ID.
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

// startSystemUserIDWatcher starts a background worker to periodically check for system user.
func startSystemUserIDWatcher(ctx context.Context, database *ent.Database, logger log.Logger) {
	go func() {
		b := backoff.NewExponentialBackOff()
		b.InitialInterval = 1 * time.Second
		b.MaxInterval = 30 * time.Second
		b.Reset()

		_, _ = backoff.Retry(ctx, func() (struct{}, error) {
			if id := GetSystemUserID(); id != 0 {
				return struct{}{}, nil
			}
			if refreshSystemUserID(ctx, database) {
				return struct{}{}, nil
			}
			return struct{}{}, fmt.Errorf("system user not found")
		}, backoff.WithBackOff(b), backoff.WithMaxElapsedTime(0))
	}()
}

// Data encapsulates the core data access components.
type Data struct {
	DB  *ent.Database
	log *log.Helper
}

// ProvideDatabase extracts the initialized *ent.Database from the runtime container
// and performs business-level initialization (System User caching).
func ProvideDatabase(app *runtime.App) (*ent.Database, func(), error) {
	db, err := comp.GetDefault[*ent.Database](app.Context(), app.Container().In("infrastructure/ent"))
	if err != nil {
		return nil, nil, err
	}

	// Business Initialization: Handle System User Identity
	ctx := context.Background()
	logHelper := log.NewHelper(app.Logger())
	if !refreshSystemUserID(ctx, db) {
		logHelper.Info("System user not found on startup, starting background watcher.")
		startSystemUserIDWatcher(ctx, db, app.Logger())
	} else {
		logHelper.Infof("System user ID cached on startup: %d", GetSystemUserID())
	}

	cleanup := func() {
		if db != nil {
			if err := db.Client(ctx).Close(); err != nil {
				logHelper.Errorf("failed to close ent client: %v", err)
			}
		}
	}

	return db, cleanup, nil
}

// NewData creates a new Data instance.
func NewData(database *ent.Database, logger log.Logger) (*Data, error) {
	return &Data{DB: database, log: log.NewHelper(logger)}, nil
}
