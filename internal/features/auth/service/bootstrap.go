/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/features/auth/biz"
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

// Bootstrap performs one-time initialization of authorization policies.
// This should be called at application startup via a BeforeStart hook.
func (b *PolicyBootstrap) Bootstrap(ctx context.Context) error {
	b.log.WithContext(ctx).Info("Starting authorization policy bootstrap...")
	if err := b.syncer.Sync(ctx); err != nil {
		// We log the error but return it to stop the application startup if strict consistency is required.
		// If you want to allow the app to start even if sync fails, return nil here.
		b.log.WithContext(ctx).Errorf("Authorization policy bootstrap failed: %v", err)
		return err
	}
	b.log.WithContext(ctx).Info("Authorization policy bootstrap completed successfully.")
	return nil
}
