// Copyright 2024 OrigAdmin. All rights reserved.

package tools

import (
	"context"
	"fmt"
	"testing"

	"entgo.io/ent/dialect/sql"

	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
)

// SetupTestDatabase 创建测试数据库
// 修正了参数顺序: dialect 在第二个参数,DSN 在第三个参数
func SetupTestDatabase(t *testing.T, opts ...enttest.Option) *ent.Client {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	t.Cleanup(func() { client.Close() })
	return client
}

// SetupTestDatabaseWithPostgreSQL 使用PostgreSQL创建测试数据库
func SetupTestDatabaseWithPostgreSQL(t *testing.T, dsn string) *ent.Client {
	drv, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("failed opening database: %v", err)
	}

	db := drv.DB()
	t.Cleanup(func() {
		db.Close()
	})

	client := ent.NewClient(ent.Driver(drv))
	t.Cleanup(func() {
		client.Close()
	})

	// 自动迁移
	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatalf("failed creating schema: %v", err)
	}

	return client
}

// TruncateTables 清空指定表
func TruncateTables(ctx context.Context, database *ent.Database, tables ...string) error {
	// 按外键依赖顺序清空
	for _, table := range tables {
		// 使用 ent.Database 的 Exec 方法执行原生 SQL
		_, err := database.Exec(ctx, "DELETE FROM "+table)
		if err != nil {
			return err
		}
	}
	return nil
}

// InTransaction 在事务中执行操作并自动回滚
func InTransaction(ctx context.Context, client *ent.Client, fn func(context.Context, *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(ctx, tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return fmt.Errorf("rolling back transaction: %v", rerr)
		}
		return err
	}

	return tx.Commit()
}
