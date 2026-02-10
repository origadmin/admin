/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/features/system/biz"
)

const (
	maxRetries    = 10
	retryInterval = 3 * time.Second
)

// PolicyBootstrap is responsible for bootstrapping authorization policies at service startup.
type PolicyBootstrap struct {
	syncer *biz.PolicySyncUseCase
	log    *log.Helper
}

// NewPolicyBootstrap creates a new PolicyBootstrap instance.
func NewPolicyBootstrap(syncer *biz.PolicySyncUseCase, logger log.Logger) *PolicyBootstrap {
	return &PolicyBootstrap{
		syncer: syncer,
		log:    log.NewHelper(log.With(logger, "module", "system.service.policy_bootstrap")),
	}
}

// Bootstrap performs one-time initialization of authorization policies with a retry mechanism.
// This should be called at application startup via a BeforeStart hook.
//
// Bootstrap Process:
// 1. Fetch all policies from URPR data source (User-Role-Permission-Resource)
// 2. Clear existing casbin_rule table
// 3. Write new policies to casbin_rule table using casbin_policy_modifier
// 4. Trigger watcher to notify other instances of policy updates
func (b *PolicyBootstrap) Bootstrap(ctx context.Context) error {
	b.log.WithContext(ctx).Info("Starting authorization policy bootstrap...")

	var lastErr error
	for i := 0; i < maxRetries; i++ {
		// Use SyncNow to ensure immediate execution without debounce delay during startup
		err := b.syncer.SyncNow(ctx)
		if err == nil {
			b.log.WithContext(ctx).Info("Authorization policy bootstrap completed successfully.")
			return nil
		}

		lastErr = err
		st, ok := status.FromError(err)
		// Retry on transient errors (Unavailable, DeadlineExceeded) or if the database is not yet ready
		if ok && (st.Code() == codes.Unavailable || st.Code() == codes.DeadlineExceeded) {
			b.log.WithContext(ctx).Warnf("Policy sync failed (transient error), retrying in %v... (attempt %d/%d): %v", retryInterval, i+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}

		// Also retry on generic errors during startup, as DB might be initializing
		b.log.WithContext(ctx).Warnf("Policy sync failed, retrying in %v... (attempt %d/%d): %v", retryInterval, i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}

	b.log.WithContext(ctx).Errorf("Authorization policy bootstrap failed after %d attempts: %v", maxRetries, lastErr)
	return lastErr
}
