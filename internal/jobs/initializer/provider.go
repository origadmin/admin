package initializer

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime/log"
	initizertypes "origadmin/application/admin/internal/jobs/initializer/types"

	jobsbroker "origadmin/application/admin/internal/jobs/initializer/broker"
	jobsmigration "origadmin/application/admin/internal/jobs/initializer/migration"
	jobsseeder "origadmin/application/admin/internal/jobs/initializer/seeder"
)

// ProviderSet is the main provider set for the initializer job.
var ProviderSet = wire.NewSet(
	jobsbroker.ProviderSet,
	jobsmigration.ProviderSet,
	jobsseeder.ProviderSet,
	ProvideTasksSlice,
	ProvideManager,
)

// ProvideTasksSlice collects all tasks into a slice.
// Directly accept concrete types since wire cannot collect multiple implementations of the same interface.
func ProvideTasksSlice(
	migrationInit *jobsmigration.Initializer,
	brokerInit *jobsbroker.Initializer,
	seederInit *jobsseeder.Initializer,
) []initizertypes.Task {
	return []initizertypes.Task{
		migrationInit,
		brokerInit,
		seederInit,
	}
}

// ProvideManager provides the Manager with all registered initialization tasks.
// Tasks will be executed in phase order (Schema -> Infrastructure -> Data -> Finalize),
// and by priority within each phase.
func ProvideManager(
	logger log.Logger,
	tasks []initizertypes.Task,
) *Manager {
	return NewManager(tasks, logger)
}
