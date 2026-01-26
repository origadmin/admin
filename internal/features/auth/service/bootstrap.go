/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/features/auth/biz"
)

// CasbinBootstrap is responsible for bootstrapping Casbin at service startup.
type CasbinBootstrap struct {
	synchronizer *biz.CasbinSynchronizer
	log          *log.Helper
}

// NewCasbinBootstrap creates a new Casbin bootstrap instance.
func NewCasbinBootstrap(synchronizer *biz.CasbinSynchronizer, logger log.Logger) *CasbinBootstrap {
	return &CasbinBootstrap{
		synchronizer: synchronizer,
		log:          log.NewHelper(log.With(logger, "module", "auth.service.bootstrap")),
	}
}

// Bootstrap performs one-time initialization of Casbin policies.
// This should be called at application startup via a BeforeStart hook.
func (b *CasbinBootstrap) Bootstrap(ctx context.Context) error {
	b.log.WithContext(ctx).Info("Starting Casbin policy bootstrap...")
	if err := b.synchronizer.Sync(ctx); err != nil {
		return err
	}
	b.log.WithContext(ctx).Info("Casbin policy bootstrap completed successfully.")
	return nil
}
