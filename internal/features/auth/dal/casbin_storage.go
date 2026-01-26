/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/casbin/casbin/v3/model"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data"
)

// CasbinStorageManager defines the interface for bulk/storage-level operations on Casbin policies.
// This is used by the initial synchronizer.
type CasbinStorageManager interface {
	// OverwriteAllPolicies atomically replaces all policies in the storage with the given model.
	OverwriteAllPolicies(ctx context.Context, m model.Model) error
}

// casbinStorageManager implements the CasbinStorageManager interface.
type casbinStorageManager struct {
	adapter *data.CasbinAdapter
	log     *log.Helper
}

// NewCasbinStorageManager creates a new storage manager.
func NewCasbinStorageManager(adapter *data.CasbinAdapter, logger log.Logger) (CasbinStorageManager, error) {
	return &casbinStorageManager{
		adapter: adapter,
		log:     log.NewHelper(log.With(logger, "module", "auth.dal.casbin_storage")),
	}, nil
}

// OverwriteAllPolicies calls the underlying adapter's SavePolicy method to perform a full overwrite.
func (s *casbinStorageManager) OverwriteAllPolicies(ctx context.Context, m model.Model) error {
	s.log.WithContext(ctx).Info("DAL: Calling adapter to save/overwrite all policies.")
	return s.adapter.SavePolicy(m)
}
