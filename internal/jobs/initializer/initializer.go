package initializer

import (
	"context"
	"fmt"

	"github.com/origadmin/runtime/log"
)

// Initializer defines an interface for components that perform idempotent initialization tasks.
// These tasks can include infrastructure provisioning (e.g., NATS Streams) and data seeding.
type Initializer interface {
	// Init executes the initialization logic.
	// It must be idempotent, meaning it can be run multiple times without causing errors.
	Init(ctx context.Context) error
}

// compositeInitializer executes a list of Initializers in order.
type compositeInitializer struct {
	initializers []Initializer
	log          *log.Helper
}

// NewCompositeInitializer creates a new CompositeInitializer that runs a series of initializers.
// This is the primary entry point for creating an initialization service.
func NewCompositeInitializer(initializers []Initializer, logger log.Logger) Initializer {
	return &compositeInitializer{
		initializers: initializers,
		log:          log.NewHelper(log.With(logger, "module", "initializer.composite")),
	}
}

// Init executes all registered initializers sequentially.
func (c *compositeInitializer) Init(ctx context.Context) error {
	c.log.Info("Starting composite initialization...")
	for i, init := range c.initializers {
		c.log.Infof("Executing initializer %d/%d: %T", i+1, len(c.initializers), init)
		if err := init.Init(ctx); err != nil {
			return fmt.Errorf("initializer %T failed: %w", init, err)
		}
	}
	c.log.Info("Composite initialization completed successfully.")
	return nil
}
