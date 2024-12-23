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

//func (w *watcher) UpdateForAddPolicy(sec, ptype string, params ...string) error {
//	return w.Update()
//}
//
//func (w *watcher) UpdateForRemovePolicy(sec, ptype string, params ...string) error {
//	return w.Update()
//}
//
//func (w *watcher) UpdateForRemoveFilteredPolicy(sec, ptype string, fieldIndex int, fieldValues ...string) error {
//	return w.Update()
//}
//
//func (w *watcher) UpdateForSavePolicy(model model.Model) error {
//	w.mu.Lock()
//	defer w.mu.Unlock()
//	if w.callback != nil {
//		w.callback(model.ToText())
//	}
//	return nil
//}
//
//func (w *watcher) UpdateForAddPolicies(sec string, ptype string, rules ...[]string) error {
//	return w.Update()
//}
//
//func (w *watcher) UpdateForRemovePolicies(sec string, ptype string, rules ...[]string) error {
//	return w.Update()
//}

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

//var _ persist.WatcherEx = &watcher{}
