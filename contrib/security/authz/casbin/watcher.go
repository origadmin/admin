/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package casbin implements the functions, types, and interfaces for the module.
package casbin

import (
	"sync"

	"github.com/casbin/casbin/v2/persist"
)

type Watcher interface {
	SetUpdateCallback(f func(string)) error
	Update() error
	Close()
}

type watcher struct {
	mu       sync.RWMutex
	callback func(string)
}

func (w *watcher) SetUpdateCallback(f func(string)) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.callback = f
	return nil
}

func (w *watcher) Update() error {
	w.mu.RLock()
	defer w.mu.RUnlock()
	if w.callback != nil {
		w.callback("")
	}
	return nil
}

func (w *watcher) Close() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.callback = nil
}

func NewWatcher() persist.Watcher {
	return &watcher{}
}

var _ persist.Watcher = &watcher{}
