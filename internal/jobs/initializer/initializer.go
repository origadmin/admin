package initializer

import (
	"context"
	"fmt"
	"sort"

	"github.com/origadmin/runtime/log"
	initizertypes "origadmin/application/admin/internal/jobs/initializer/types"
)

// Initializer defines an interface for components that perform idempotent initialization tasks.
// These tasks can include infrastructure provisioning (e.g., NATS Streams) and data seeding.
type Initializer interface {
	// Init executes the initialization logic.
	// It must be idempotent, meaning it can be run multiple times without causing errors.
	Init(ctx context.Context) error
}

// Manager coordinates the execution of multiple initialization tasks.
type Manager struct {
	tasks []initizertypes.Task
	log   *log.Helper
}

// NewManager creates a new Manager that executes tasks in the correct order.
func NewManager(tasks []initizertypes.Task, logger log.Logger) *Manager {
	return &Manager{
		tasks: tasks,
		log:   log.NewHelper(log.With(logger, "module", "initializer.manager")),
	}
}

// Init executes all registered tasks in phase order, then by priority within each phase.
func (m *Manager) Init(ctx context.Context) error {
	m.log.Info("Starting layered initialization...")

	// Group tasks by phase
	phases := make(map[initizertypes.Phase][]initizertypes.Task)
	for _, task := range m.tasks {
		phases[task.Phase()] = append(phases[task.Phase()], task)
	}

	// Get sorted phase list
	sortedPhases := make([]initizertypes.Phase, 0, len(phases))
	for phase := range phases {
		sortedPhases = append(sortedPhases, phase)
	}
	sort.Slice(sortedPhases, func(i, j int) bool {
		return sortedPhases[i] < sortedPhases[j]
	})

	// Execute tasks in phase order
	for _, phase := range sortedPhases {
		tasksInPhase := phases[phase]
		// Sort tasks by priority (higher first)
		sort.Slice(tasksInPhase, func(i, j int) bool {
			return tasksInPhase[i].Priority() > tasksInPhase[j].Priority()
		})

		m.log.Infof("Executing Phase %d with %d tasks...", phase, len(tasksInPhase))
		for _, task := range tasksInPhase {
			m.log.Infof("Executing task: %s (Phase: %d, Priority: %d)", task.Name(), task.Phase(), task.Priority())
			if err := task.Init(ctx); err != nil {
				return fmt.Errorf("task '%s' failed: %w", task.Name(), err)
			}
		}
	}

	m.log.Info("Layered initialization completed successfully.")
	return nil
}
