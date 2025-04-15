/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package casbin implements the functions, types, and interfaces for the module.
package casbin

import (
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/casbin/casbin/v2/persist"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
)

type RetryPolicy struct {
	BaseDelay  time.Duration
	MaxDelay   time.Duration
	MaxRetries int
}

type PBWatcher struct {
	client       pb.CasbinSourceServiceClient
	setter       PolicySetter
	callback     func(string)
	mu           sync.RWMutex
	ctx          context.Context
	cancel       context.CancelFunc
	retryPolicy  RetryPolicy
	lastModified int64
}

func (w *PBWatcher) SetUpdateCallback(cb func(string)) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.callback = cb
	return nil
}

func (w *PBWatcher) Update() error {
	w.mu.RLock()
	defer w.mu.RUnlock()
	if w.callback != nil {
		w.callback("")
	}
	return nil
}

func (w *PBWatcher) Close() {
	w.cancel()
}

func (w *PBWatcher) Watch() {
	go w.watchLoop()
}

func (w *PBWatcher) handleError(err error, retryCount int) int {
	if retryCount >= w.retryPolicy.MaxRetries {
		log.Fatal("Max retries exceeded")
		//return retryCount
	}

	delay := w.retryPolicy.BaseDelay * (1 << retryCount)
	if delay > w.retryPolicy.MaxDelay {
		delay = w.retryPolicy.MaxDelay
	}

	log.Warnf("Retrying in %v (attempt %d)", delay, retryCount+1)
	time.Sleep(delay)
	return retryCount + 1
}

func (w *PBWatcher) watchLoop() {
	retryCount := 0
	lastModified := int64(0)
	for {
		select {
		case <-w.ctx.Done():
			return
		default:
			err := w.streamUpdates(lastModified)
			if err != nil {
				retryCount = w.handleError(err, retryCount)
			} else {
				retryCount = 0
			}
		}
	}
}

func (w *PBWatcher) streamUpdates(lastVersion int64) error {
	resp, err := w.client.WatchUpdate(w.ctx, &pb.WatchUpdateRequest{
		LastModified: atomic.LoadInt64(&w.lastModified),
	})

	if err != nil {
		return err
	}

	if resp.ModifiedDate > atomic.LoadInt64(&w.lastModified) {
		policies, err := w.fetchPolicies()
		if err != nil {
			return err
		}

		w.mu.Lock()
		defer w.mu.Unlock()
		w.setter.SetPolicies(policies)
		atomic.StoreInt64(&w.lastModified, resp.ModifiedDate)

		if w.callback != nil {
			w.callback("")
		}
	}
	return nil
}

func (w *PBWatcher) fetchPolicies() (map[string][][]string, error) {
	ctx, cancel := context.WithTimeout(w.ctx, 5*time.Second)
	defer cancel()

	stream, err := w.client.StreamRules(ctx, &pb.StreamRulesRequest{
		WithGroupings: true,
		WithPolicies:  true,
	})
	if err != nil {
		return nil, err
	}

	policies := make(map[string][][]string)
	for {
		rule, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch v := rule.RuleType.(type) {
		case *pb.StreamRulesResponse_Policy:
			policies[v.Policy.PType] = append(policies[v.Policy.PType], v.Policy.Params)
		case *pb.StreamRulesResponse_Grouping:
			policies[v.Grouping.PType] = append(policies[v.Grouping.PType], v.Grouping.Params)
		}
	}
	return policies, nil
}

func NewPBWatcher(ctx context.Context, client pb.CasbinSourceServiceClient, setter PolicySetter) (persist.Watcher, error) {
	ctx, cancel := context.WithCancel(ctx)
	return &PBWatcher{
		client: client,
		setter: setter,
		ctx:    ctx,
		cancel: cancel,
		retryPolicy: RetryPolicy{
			BaseDelay:  1 * time.Second,
			MaxDelay:   30 * time.Second,
			MaxRetries: 5,
		},
	}, nil
}
