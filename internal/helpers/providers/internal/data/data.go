/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"context"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/contracts/component"
	storageiface "github.com/origadmin/runtime/contracts/storage"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	internaldata "origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/data/entity/ent"
)

const (
	CategoryEnt       component.Category = "ent"
	NameCasbinAdapter                    = "casbin-adapter"
	NameEnt                              = "ent"
)

// NewEnt is the engine provider for *ent.Database.
func NewEnt(ctx context.Context, h component.Handle) (any, error) {
	// 1. Retrieve Raw Infrastructure DB
	dbInst, err := comp.GetDefault[storageiface.Database](ctx, h.Locator().In(runtime.CategoryDatabase))
	if err != nil {
		return nil, fmt.Errorf("engine: failed to get infrastructure database: %w", err)
	}

	// 2. Initialize Ent Database
	activeDB := entsql.OpenDB(dbInst.Dialect(), dbInst.DB())
	return ent.NewDatabase(activeDB), nil
}

// EntResolver resolves Ent configuration.
func EntResolver(ctx context.Context, root any, opts *component.LoadOptions) (*component.ModuleConfig, error) {
	return &component.ModuleConfig{
		Entries: []component.ConfigEntry{{Name: NameEnt, Value: root}},
		Active:  NameEnt,
	}, nil
}

// CasbinAdapterProvider creates a Casbin adapter.
func CasbinAdapterProvider(ctx context.Context, h component.Handle) (any, error) {
	if h.Name() != NameCasbinAdapter {
		return nil, nil
	}
	// CategoryEnt is in GlobalScope
	dbH := h.Locator().In(CategoryEnt)
	dbInst, err := comp.GetDefault[*ent.Database](ctx, dbH)
	if err != nil {
		return nil, err
	}
	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(runtime.CategoryLogger))
	return internaldata.NewAdapter(ctx, dbInst, logger)
}

// CasbinAdapterResolver resolves Casbin adapter configuration.
func CasbinAdapterResolver(ctx context.Context, root any, opts *component.LoadOptions) (*component.ModuleConfig, error) {
	return &component.ModuleConfig{Entries: []component.ConfigEntry{
		{Name: NameCasbinAdapter, Value: root},
	}}, nil
}
