/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/origadmin/toolkits/database"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMenuDal, NewLoginDal)

// Data .
type Data struct {
	*ent.Database
}

// NewTrans returns a transaction wit data
func NewTrans(data *Data) database.Trans {
	return data
}

// NewData .
func NewData(bootstrap *configs.Bootstrap, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	database := bootstrap.GetData().GetDatabase()
	drv, err := sql.Open(
		database.Driver,
		database.Source,
	)
	if err != nil {
		log.Errorw("failed opening connection to sqlite", "error", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	client := ent.NewClient(ent.Driver(drv))
	if database.GetMigrate() {
		if err := client.Schema.Create(context.Background()); err != nil {
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
