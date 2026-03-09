/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"context"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/origadmin/runtime/contracts/component"
	storageiface "github.com/origadmin/runtime/contracts/storage"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
)

// NewEnt is the engine provider for *ent.Database.
// It handles purely infrastructural initialization: creating the client and running schema migrations.
func NewEnt(ctx context.Context, h component.Handle) (any, error) {
	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(component.CategoryLogger))
	logHelper := log.NewHelper(logger)

	// 1. Retrieve Raw Infrastructure DB
	dbInst, err := comp.GetDefault[storageiface.Database](ctx, h.Locator().In(component.CategoryDatabase))
	if err != nil {
		return nil, fmt.Errorf("engine: failed to get infrastructure database: %w", err)
	}

	// 2. Initialize Ent Database
	activeDB := entsql.OpenDB(dbInst.Dialect(), dbInst.DB())
	logHelper.Infof("Initializing ent database infrastructure with dialect: %s", dbInst.Dialect())

	database := ent.NewDatabase(activeDB)

	// 3. Run Migrations (Infrastructure Setup)
	if err := database.Migration(ctx,
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
		schema.WithForeignKeys(false),
	); err != nil {
		return nil, fmt.Errorf("engine: failed creating schema resources: %w", err)
	}

	return database, nil
}
