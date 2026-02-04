/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package debounce

import (
	"sync"
	"time"
)

// Executor defines an interface for a debounced function executor.
// It allows scheduling a function to run after a certain delay, with any subsequent
// calls within that delay resetting the timer.
type Executor interface {
	// Schedule arranges for the function f to be called after the configured delay.
	// If Schedule is called again before the delay has passed, the previous timer is reset.
	Schedule(f func())

	// Cancel stops any pending scheduled function from executing.
	Cancel()

	// IsPending returns true if there is a function scheduled to run.
	IsPending() bool
}

// debouncedExecutor implements the Executor interface.
type debouncedExecutor struct {
	delay time.Duration
	mu    sync.Mutex
	timer *time.Timer
}

// NewExecutor creates a new debounced executor with the specified delay.
// If the provided delay is zero or negative, a default of 1 second will be used.
func NewExecutor(delay time.Duration) Executor {
	if delay <= 0 {
		delay = time.Second // A sensible default to prevent zero-delay issues.
	}
	return &debouncedExecutor{delay: delay}
}

// Schedule implements the Executor interface.
func (e *debouncedExecutor) Schedule(f func()) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.timer != nil {
		e.timer.Stop()
	}
	e.timer = time.AfterFunc(e.delay, f)
}

// Cancel implements the Executor interface.
func (e *debouncedExecutor) Cancel() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.timer != nil {
		e.timer.Stop()
		e.timer = nil
	}
}

// IsPending implements the Executor interface.
func (e *debouncedExecutor) IsPending() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.timer != nil
}
