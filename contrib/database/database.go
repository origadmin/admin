/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package database implements the functions, types, and interfaces for the module.
package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/errors"

	"origadmin/application/admin/contrib/database/internal/mysql"
	"origadmin/application/admin/contrib/database/internal/sqlite"
)

func Open(database *configv1.Database) (*sql.DB, error) {
	if database == nil {
		return nil, errors.New("config: database is nil")
	}
	switch database.Dialect {
	case "mysql":
		err := mysql.CreateDatabase(database.Source, "")
		if err != nil {
			return nil, errors.Wrap(err, "mysql: create database error")
		}
	case "pgx":
		database.Dialect = "postgres"
	case "sqlite3", "sqlite":
		database.Dialect = "sqlite3"
		if !strings.Contains(database.Source, ":memory:") {
			sqlite.MakeSourceDirectory(database.Source)
		}
		database.Source = sqlite.SourceForeignKeys(database.Source)
	default:

	}
	fmt.Printf("database: dialect: %s, source: %s\n", database.Dialect, database.Source)
	db, err := sql.Open(database.Dialect, database.Source)
	if err != nil {
		return nil, errors.Wrap(err, "database: open database error")
	}
	if database.MaxIdleConnections > 0 {
		db.SetMaxIdleConns(int(database.MaxIdleConnections))
	}
	if database.MaxOpenConnections > 0 {
		db.SetMaxOpenConns(int(database.MaxOpenConnections))
	}
	if t := database.ConnectionMaxLifetime; t > 0 {
		db.SetConnMaxLifetime(time.Duration(t))
	}
	if t := database.ConnectionMaxIdleTime; t > 0 {
		db.SetConnMaxIdleTime(time.Duration(t))
	}
	return db, nil
}
