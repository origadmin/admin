/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/google/wire"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/database"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAuthRepo, NewMenuRepo, NewRoleRepo, NewUserRepo)

// Data .
type Data struct {
	*ent.Database
}

// NewTrans returns a transaction wit data
func NewTrans(data *Data) database.Trans {
	return data
}

// NewData .
func NewData(bootstrap *configs.Bootstrap, logger log.KLogger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	//s := data.NewService(bootstrap.GetData(), nil)
	cfg := bootstrap.GetData().GetDatabase()
	if cfg == nil {
		return nil, nil, errors.New("data source not found")
	}
	drv, err := sql.Open(
		cfg.Driver,
		cfg.Source,
	)

	drv.Exec(context.Background(), "PRAGMA foreign_keys = ON;", []any{}, nil)
	if err != nil {
		log.Errorw("failed opening connection to sqlite", "error", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	client := ent.NewClient(ent.Driver(drv))
	if cfg.GetMigrate() {
		if err := client.Schema.Create(
			context.Background(),
			schema.WithDropIndex(true),
			schema.WithDropColumn(true),
			schema.WithForeignKeys(true)); err != nil {
			log.Errorw("failed creating schema resources", "error", err)
			return nil, nil, err
		}
	}

	d := &Data{
		Database: ent.NewDatabase(client),
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := drv.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
