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
)

// ProviderSet exports the Seeder initializer and its dependencies.
var ProviderSet = wire.NewSet(
	taskseeder.ProviderSet, // Include the actual task implementation
	NewInitializer,
)

// Initializer implements the Initializer interface for data seeding.
type Initializer struct {
	seeder *taskseeder.Seeder
	log    *log.Helper
}

// NewInitializer creates a new DataSeederInitializer.
func NewInitializer(s *taskseeder.Seeder, logger log.Logger) *Initializer {
	return &Initializer{
		seeder: s,
		log:    log.NewHelper(log.With(logger, "module", "initializer.data_seeder")),
	}
}

// Init executes the data seeding logic by calling the Seeder's Run method.
func (d *Initializer) Init(ctx context.Context) error {
	d.log.Info("Starting data seeding initialization...")
	if err := d.seeder.Run(); err != nil {
		return fmt.Errorf("seeder failed: %w", err)
	}
	d.log.Info("Data seeding initialization completed successfully.")
	return nil
}
