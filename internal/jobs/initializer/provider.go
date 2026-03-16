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
	ProvideManager,
)

// ProvideManager provides the Manager with all registered initialization tasks.
// Tasks will be executed in phase order (Schema -> Infrastructure -> Data -> Finalize),
// and by priority within each phase.
func ProvideManager(
	logger log.Logger,
	brokerInit initizertypes.Task,
	migrationInit initizertypes.Task,
	seederInit initizertypes.Task,
) *Manager {
	tasks := []initizertypes.Task{
		migrationInit,
		brokerInit,
		seederInit,
	}
	return NewManager(tasks, logger)
}
