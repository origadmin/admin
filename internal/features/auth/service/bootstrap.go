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
	"origadmin/application/admin/internal/features/auth/biz"
)

const (
	maxRetries    = 10
	retryInterval = 3 * time.Second
)

// PolicyBootstrap is responsible for bootstrapping authorization policies at service startup.
type PolicyBootstrap struct {
	syncer *biz.PolicySyncer
	log    *log.Helper
}

// NewPolicyBootstrap creates a new PolicyBootstrap instance.
func NewPolicyBootstrap(syncer *biz.PolicySyncer, logger log.Logger) *PolicyBootstrap {
	return &PolicyBootstrap{
		syncer: syncer,
		log:    log.NewHelper(log.With(logger, "module", "auth.service.policy_bootstrap")),
	}
}

// Bootstrap performs one-time initialization of authorization policies with a retry mechanism.
// This should be called at application startup via a BeforeStart hook.
func (b *PolicyBootstrap) Bootstrap(ctx context.Context) error {
	b.log.WithContext(ctx).Info("Starting authorization policy bootstrap...")

	var lastErr error
	for i := 0; i < maxRetries; i++ {
		// Use ForceSync to ensure immediate execution without debounce delay during startup
		err := b.syncer.ForceSync(ctx)
		if err == nil {
			b.log.WithContext(ctx).Info("Authorization policy bootstrap completed successfully.")
			return nil
		}

		lastErr = err
		st, ok := status.FromError(err)
		if ok && (st.Code() == codes.Unavailable || st.Code() == codes.DeadlineExceeded) {
			b.log.WithContext(ctx).Warnf("System service not available, retrying in %v... (attempt %d/%d)", retryInterval, i+1, maxRetries)
			time.Sleep(retryInterval)
			continue
		}

		// For other errors, fail immediately.
		b.log.WithContext(ctx).Errorf("Authorization policy bootstrap failed with a non-retriable error: %v", err)
		return err
	}

	b.log.WithContext(ctx).Errorf("Authorization policy bootstrap failed after %d attempts: %v", maxRetries, lastErr)
	return lastErr
}
