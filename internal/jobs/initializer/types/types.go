/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types defines core interfaces and types for initialization tasks.
package types

import "context"

// Phase defines the execution phase for an initialization task.
// Tasks are executed in ascending order of their phase.
type Phase int

const (
	// PhaseSchema is for low-level storage structure setup.
	// Examples: DB Migrations, NATS Stream/KV creation, Directory structure setup.
	PhaseSchema Phase = 100

	// PhaseInfrastructure is for infrastructure-level connectivity or basic provisioning.
	// Examples: Broker client warmup, External API health checks.
	PhaseInfrastructure Phase = 200

	// PhaseData is for populating the system with base application data.
	// Examples: Default admin creation, Permission seeding, I18n loading.
	PhaseData Phase = 300

	// PhaseFinalize is for post-initialization cleanup or readiness signaling.
	// Examples: Cache invalidation, Starting background watchers.
	PhaseFinalize Phase = 400
)

// Task defines an interface for a modular, self-describing initialization task.
// The Manager will use the metadata provided by Phase() and Priority() to execute tasks
// in the correct, deterministic order.
type Task interface {
	// Name returns a unique, human-readable name for the task, used for logging.
	Name() string

	// Phase returns the execution phase of the task.
	Phase() Phase

	// Priority returns the execution priority of the task within its phase.
	// Higher values are executed first.
	Priority() int

	// Init executes the initialization logic of the task.
	// It must be idempotent.
	Init(ctx context.Context) error
}
