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

	// Create a copy to avoid modifying the original slice
	tasks := make([]initizertypes.Task, len(m.tasks))
	copy(tasks, m.tasks)

	// Sort all tasks by phase first, then by priority (both descending)
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].Phase() != tasks[j].Phase() {
			return tasks[i].Phase() < tasks[j].Phase() // Lower phase first
		}
		return tasks[i].Priority() > tasks[j].Priority() // Higher priority first
	})

	// Execute tasks in sorted order
	for _, task := range tasks {
		m.log.Infof("Executing task: %s (Phase: %d, Priority: %d)", task.Name(), task.Phase(), task.Priority())
		if err := task.Init(ctx); err != nil {
			return fmt.Errorf("task '%s' failed: %w", task.Name(), err)
		}
	}

	m.log.Info("Layered initialization completed successfully.")
	return nil
}
