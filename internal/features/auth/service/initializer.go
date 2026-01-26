/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"fmt"

	"github.com/casbin/casbin/v3"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/features/auth/biz"
)

// CasbinInitializer is responsible for seeding the Casbin enforcer with
// initial policies if the underlying storage (e.g., database) is empty.
type CasbinInitializer struct {
	enforcer     *casbin.Enforcer
	synchronizer *biz.CasbinSynchronizer
	log          *log.Helper
}

// NewCasbinInitializer creates a new initializer.
// It requires the core Authorizer (which should be a Casbin enforcer) and the synchronizer.
func NewCasbinInitializer(
	authorizer authz.Authorizer,
	synchronizer *biz.CasbinSynchronizer,
	logger log.Logger,
) (*CasbinInitializer, error) {
	enforcer, ok := authorizer.(*casbin.Enforcer)
	if !ok {
		return nil, fmt.Errorf("authorizer is not a *casbin.Enforcer, cannot initialize")
	}

	return &CasbinInitializer{
		enforcer:     enforcer,
		synchronizer: synchronizer,
		log:          log.NewHelper(log.With(logger, "module", "auth.service.casbin-initializer")),
	}, nil
}

// Initialize seeds the enforcer if no policies are currently loaded.
// This method should be called once at application startup.
func (i *CasbinInitializer) Initialize(ctx context.Context) error {
	// The enforcer automatically calls LoadPolicy() on creation.
	// We just need to check if it's empty.
	if len(i.enforcer.GetPolicy()) > 0 || len(i.enforcer.GetGroupingPolicy()) > 0 {
		i.log.Info("Casbin policies already exist. Skipping initialization.")
		return nil
	}

	i.log.Info("No Casbin policies found. Starting full synchronization...")

	// Use the synchronizer to fetch and sync policies from the system module
	if err := i.synchronizer.Sync(ctx); err != nil {
		return fmt.Errorf("failed to sync policies from system module: %w", err)
	}

	i.log.Info("Successfully seeded and saved initial Casbin policies.")
	return nil
}
