/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"context"
	"fmt"
	"os"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/origadmin/runtime/contracts/component"
	storageiface "github.com/origadmin/runtime/contracts/storage"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
)

// NewEnt is the engine provider for *ent.Database.
// It handles purely infrastructural initialization: creating the client connection.
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
	logHelper.Infof("Connecting to ent database with dialect: %s", dbInst.Dialect())

	database := ent.NewDatabase(activeDB)

	// 3. Conditional Migration (Only for Initializer Module)
	if os.Getenv("DB_AUTO_MIGRATION") == "true" {
		logHelper.Info("DB_AUTO_MIGRATION is enabled. Executing schema migration...")
		if err := database.Migration(ctx,
			schema.WithDropIndex(true),
			schema.WithDropColumn(true),
			schema.WithForeignKeys(false),
		); err != nil {
			return nil, fmt.Errorf("engine: schema migration failed: %w", err)
		}

		// CRITICAL FIX: Repair the 'admin' user is_system flag if it was incorrectly created as 'false'.
		// This uses raw SQL to bypass Biz layer immutability constraints and preserve all ID-based relations.
		fixResult, err := dbInst.DB().ExecContext(ctx,
			"UPDATE sys_users SET is_system = true WHERE username = 'admin' AND is_system = false")
		if err == nil {
			rows, _ := fixResult.RowsAffected()
			if rows > 0 {
				logHelper.Warnf("Successfully repaired root user 'admin' status: set is_system=true for existing record.")
			}
		}

		logHelper.Info("Infrastructure initialization completed successfully.")
	}

	return database, nil
}
