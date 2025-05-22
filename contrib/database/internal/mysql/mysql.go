/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/goexts/generic/types"
)

const (
	databaseCreateQuery = "CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET '%s' DEFAULT COLLATE '%s';"
	defaultCharSet      = "utf8mb4"
	defaultCollate      = "utf8mb4_general_ci"
)

// CreateDatabase creates a MySQL database with the given DSN.
func CreateDatabase(dsn string, name string) error {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr))
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close database: %v\n", err)
		}
	}(db)

	charset := types.ZeroOr(cfg.Params["charset"], defaultCharSet)
	collate := types.ZeroOr(cfg.Collation, defaultCollate)
	if name == "" {
		name = cfg.DBName
	}
	query := fmt.Sprintf(databaseCreateQuery, name, charset, collate)
	_, err = db.Exec(query)
	return err
}
