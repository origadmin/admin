/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data implements the functions, types, and interfaces for the module.
package data

import (
	"errors"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/google/wire"
	"github.com/origadmin/contrib/database"
	"github.com/origadmin/entslog/v3"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/interfaces/security"
	"github.com/origadmin/runtime/log"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/data/entity/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	//wire.Struct(new(LoginData), "*"),
	NewDataWithClient,
)

type Data struct {
	*ent.Database
}

type LoginData struct {
	Captcha   *configs.Captcha
	RootUser  *configs.RootUser
	Tokenizer security.RefreshTokenizer
	//Resource  systemdto.ResourceRepo
	//Role      systemdto.RoleRepo
	//User      systemdto.UserRepo
}

func NewDataWithClient(client *ent.Client) *Data {
	return &Data{
		Database: ent.NewDatabaseWithClient(client),
	}
}

func debugDatabase(driver dialect.Driver, debug bool) dialect.Driver {
	if debug {
		return entslog.New(driver)
	}
	return driver
}

// NewData .

// NewData .
func NewData(r runtime.Runtime, bootstrap *configs.Bootstrap) (*Data, func(), error) {
	if bootstrap == nil {
		return nil, nil, errors.New("bootstrap is nil")
	}

	cfg := bootstrap.GetStorage().GetDatabase()
	if cfg == nil {
		return nil, nil, errors.New("data source not found")
	}

	drv, err := database.Open(cfg)
	log.Infow("msg", "connecting to database", "dialect", cfg.Dialect, "source", cfg.Source)
	if err != nil {
		log.Errorw("msg", "failed opening connection to database", "error", err)
		return nil, nil, err
	}

	// Run the auto migration tool.
	//sqldb := debugDatabase(sql.OpenDB(cfg.Dialect, drv), cfg.Debug)

	db := ent.NewDatabase(ent.Driver(sql.OpenDB(cfg.Dialect, drv)), ent.WithDebug(func(driver dialect.Driver, f ...func(...any)) dialect.Driver {
		return debugDatabase(driver, cfg.Debug)
	}))
	if true || cfg.GetMigration().GetEnabled() {
		if err := db.Migration(
			r.Context(),
			schema.WithDropIndex(true),
			schema.WithDropColumn(true),
			schema.WithForeignKeys(false)); err != nil {
			log.Errorw("msg", "failed creating schema resources", "error", err)
			return nil, nil, err
		}
	}

	data := &Data{
		Database: db,
	}

	return data, func() {
		log.Info("closing the data resources")
		if err := drv.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
