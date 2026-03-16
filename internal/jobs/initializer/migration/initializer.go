/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package migration

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
	initizertypes "origadmin/application/admin/internal/jobs/initializer/types"
)

// ProviderSet exports the Migration initializer and its dependencies.
var ProviderSet = wire.NewSet(
	NewInitializer,
	wire.Bind(new(initizertypes.Task), new(*Initializer)),
)

// Initializer implements the Task interface for database schema migration.
type Initializer struct {
	db  *ent.Database
	log *log.Helper
}

// NewInitializer creates a new MigrationInitializer.
func NewInitializer(db *ent.Database, logger log.Logger) *Initializer {
	return &Initializer{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "initializer.migration")),
	}
}

// Name returns the task name.
func (m *Initializer) Name() string {
	return "migration"
}

// Phase returns the execution phase.
func (m *Initializer) Phase() initizertypes.Phase {
	return initizertypes.PhaseSchema
}

// Priority returns the execution priority within the phase.
func (m *Initializer) Priority() int {
	return 100
}

// Init executes the database migration logic.
func (m *Initializer) Init(ctx context.Context) error {
	m.log.Info("Starting database migration initialization...")
	if err := m.db.Migration(ctx); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	m.log.Info("Database migration initialization completed successfully.")
	return nil
}
