// Code generated by ent, DO NOT EDIT.

package ent

/* Additional dependencies injected to config. */

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/origadmin/toolkits/database"
)

// Database is the client that holds all ent builders.
type Database struct {
	client *Client
}

// NewDatabase creates a new database configured with the given options.
func NewDatabase(client *Client, opts ...Option) *Database {
	if client == nil {
		client = NewClient(opts...)
	}
	return &Database{client: client}
}

func (db *Database) clientDriver(ctx context.Context) dialect.Driver {
	tx := TxFromContext(ctx)
	c := db.client
	if tx != nil {
		c = tx.Client()
	}
	return c.driver
}

// Tx runs the given function f within a transaction.
func (db *Database) Tx(ctx context.Context, fn func(context.Context) error) error {
	tx := TxFromContext(ctx)
	if tx != nil {
		return fn(ctx)
	}

	return db.InTx(ctx, func(tx database.Tx) error {
		txv, ok := tx.(*Tx)
		if !ok {
			return fmt.Errorf("ent: expected tx context")
		}
		return fn(NewTxContext(ctx, txv))
	})
}

// InTx runs the given function f within a transaction.
func (db *Database) InTx(ctx context.Context, fn func(tx database.Tx) error) error {
	tx := TxFromContext(ctx)
	if tx != nil {
		return fn(tx)
	}
	tx, err := db.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	if err = fn(tx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return fmt.Errorf("rolling back transaction: %v (original error: %w)", txerr, err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

// Client returns the client that holds all ent builders.
func (db *Database) Client(ctx context.Context) *Client {
	tx := TxFromContext(ctx)
	if tx != nil {
		return tx.Client()
	}
	return db.client
}

// Exec executes a query that doesn't return rows. For example, in SQL, INSERT or UPDATE.
func (db *Database) Exec(ctx context.Context, query string, args ...interface{}) (*sql.Result, error) {
	var res sql.Result
	err := db.clientDriver(ctx).Exec(ctx, query, args, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Query executes a query that returns rows, typically a SELECT in SQL.
func (db *Database) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var rows sql.Rows
	err := db.clientDriver(ctx).Query(ctx, query, args, &rows)
	if err != nil {
		return nil, err
	}
	return &rows, nil
}

// CasbinRule is the client for interacting with the CasbinRule builders.
func (db *Database) CasbinRule(ctx context.Context) *CasbinRuleClient {
	return db.Client(ctx).CasbinRule
}

// Department is the client for interacting with the Department builders.
func (db *Database) Department(ctx context.Context) *DepartmentClient {
	return db.Client(ctx).Department
}

// Permission is the client for interacting with the Permission builders.
func (db *Database) Permission(ctx context.Context) *PermissionClient {
	return db.Client(ctx).Permission
}

// PermissionResource is the client for interacting with the PermissionResource builders.
func (db *Database) PermissionResource(ctx context.Context) *PermissionResourceClient {
	return db.Client(ctx).PermissionResource
}

// Position is the client for interacting with the Position builders.
func (db *Database) Position(ctx context.Context) *PositionClient {
	return db.Client(ctx).Position
}

// PositionPermission is the client for interacting with the PositionPermission builders.
func (db *Database) PositionPermission(ctx context.Context) *PositionPermissionClient {
	return db.Client(ctx).PositionPermission
}

// Resource is the client for interacting with the Resource builders.
func (db *Database) Resource(ctx context.Context) *ResourceClient {
	return db.Client(ctx).Resource
}

// Role is the client for interacting with the Role builders.
func (db *Database) Role(ctx context.Context) *RoleClient {
	return db.Client(ctx).Role
}

// RolePermission is the client for interacting with the RolePermission builders.
func (db *Database) RolePermission(ctx context.Context) *RolePermissionClient {
	return db.Client(ctx).RolePermission
}

// User is the client for interacting with the User builders.
func (db *Database) User(ctx context.Context) *UserClient {
	return db.Client(ctx).User
}

// UserDepartment is the client for interacting with the UserDepartment builders.
func (db *Database) UserDepartment(ctx context.Context) *UserDepartmentClient {
	return db.Client(ctx).UserDepartment
}

// UserPosition is the client for interacting with the UserPosition builders.
func (db *Database) UserPosition(ctx context.Context) *UserPositionClient {
	return db.Client(ctx).UserPosition
}

// UserRole is the client for interacting with the UserRole builders.
func (db *Database) UserRole(ctx context.Context) *UserRoleClient {
	return db.Client(ctx).UserRole
}
