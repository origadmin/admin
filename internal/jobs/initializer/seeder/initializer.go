/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package seeder

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/origadmin/runtime/log"
	taskseeder "origadmin/application/admin/internal/jobs/tasks/seeder"
	initizertypes "origadmin/application/admin/internal/jobs/initializer/types"
)

// ProviderSet exports the TaskSeeder initializer and its dependencies.
var ProviderSet = wire.NewSet(
	taskseeder.ProviderSet, // Include the actual task implementation
	NewInitializer,
)

// Initializer implements the Task interface for data seeding.
type Initializer struct {
	seeder *taskseeder.TaskSeeder
	log    *log.Helper
}

// NewInitializer creates a new DataSeederInitializer.
func NewInitializer(s *taskseeder.TaskSeeder, logger log.Logger) *Initializer {
	return &Initializer{
		seeder: s,
		log:    log.NewHelper(log.With(logger, "module", "initializer.data_seeder")),
	}
}

// Name returns the task name.
func (d *Initializer) Name() string {
	return "seeder"
}

// Phase returns the execution phase.
func (d *Initializer) Phase() initizertypes.Phase {
	return initizertypes.PhaseData
}

// Priority returns the execution priority within the phase.
func (d *Initializer) Priority() int {
	return 100
}

// Init executes the data seeding logic by calling the TaskSeeder's Run method.
func (d *Initializer) Init(ctx context.Context) error {
	d.log.Info("Starting data seeding initialization...")
	if err := d.seeder.Run(); err != nil {
		return fmt.Errorf("seeder failed: %w", err)
	}
	d.log.Info("Data seeding initialization completed successfully.")
	return nil
}
