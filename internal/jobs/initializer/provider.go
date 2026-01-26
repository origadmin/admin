package initializer

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime/log"

	jobsbroker "origadmin/application/admin/internal/jobs/initializer/broker"
	"origadmin/application/admin/internal/jobs/initializer/seeder"
)

// ProviderSet is the main provider set for the initializer job.
// It aggregates all specific initializer providers and assembles the composite initializer.
var ProviderSet = wire.NewSet(
	jobsbroker.ProviderSet,
	seeder.ProviderSet,
	ProvideCompositeInitializer,
)

// ProvideCompositeInitializer assembles the composite initializer from individual initializer components.
// The order of arguments determines the execution order if we construct the slice manually here.
func ProvideCompositeInitializer(
	logger log.Logger,
	brokerInit *jobsbroker.Initializer,
	seederInit *seeder.Initializer,
) Initializer {
	// The order of initializers is explicitly defined here.
	// 1. Infrastructure (Broker)
	// 2. Data (Seeder)
	initializers := []Initializer{
		brokerInit,
		seederInit,
	}
	return NewCompositeInitializer(initializers, logger)
}
