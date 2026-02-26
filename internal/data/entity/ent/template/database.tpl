{{/* The line below tells Intellij/GoLand to enable the autocompletion based *gen.Type type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Type */}}


{{ define "database" }}
    {{- $pkg := base $.Config.Package }}
    {{ template "header" $ }}

		/* Additional dependencies injected to config. */
    {{- $deps := list }}{{ with $.Config.Annotations }}{{ $deps = $.Config.Annotations.Dependencies }}{{ end }}

		import (
		"context"
		"fmt"
		"entgo.io/ent/dialect"
		"entgo.io/ent/dialect/sql"
    "github.com/origadmin/runtime/contracts/storage/database"
		)

		// Database is the client that holds all ent builders.
		type Database struct {
		client *Client
		}

		// NewDatabase creates a new database with the required driver and optional options.
		// The driver parameter is required and specifies the database dialect driver to use.
		// Additional options can be passed to configure the client behavior.
		func NewDatabase(driver dialect.Driver, opts ...Option) *Database {
		client := NewClient(append([]Option{Driver(driver)}, opts...)...)
		return &Database{client: client}
		}

		// NewDatabaseWithClient creates a new database with an existing client.
		// If client is nil, it will create a new client with the provided options.
		// Note: This function is deprecated, use NewDatabase instead.
		func NewDatabaseWithClient(client *Client) *Database {
		return &Database{client: client}
		}

		// clientDriver returns the underlying database driver, considering transaction context.
		func (db *Database) clientDriver(ctx context.Context) dialect.Driver {
		if tx := TxFromContext(ctx); tx != nil {
		return tx.Client().driver
		}
		return db.client.driver
		}

		// Tx runs the given function within a new transaction.
		// This function always creates a new transaction, even if one already exists in the context.
		// The transaction is committed after the function completes successfully, or rolled back on error.
		func (db *Database) Tx(ctx context.Context, fn func(context.Context) error) error {
		return db.inTx(ctx, func(tx *Tx) error {
		return fn(NewTxContext(ctx, tx))
		})
		}

		// InTx runs the given function within a transaction.
		// If a transaction already exists in the context, it will be reused (and not committed here).
		// Otherwise, a new transaction is created, committed, or rolled back automatically.
		func (db *Database) InTx(ctx context.Context, fn func(tx database.Tx) error) error {
		if tx := TxFromContext(ctx); tx != nil {
		// Reuse existing transaction, don't commit/rollback
		return fn(tx)
		}
		// Create new transaction, handle commit/rollback
		return db.inTx(ctx, func(tx *Tx) error { return fn(tx) })
		}

		// inTx creates a new transaction and runs the given function within it.
		// The transaction is committed after the function completes successfully, or rolled back on error.
		// This is a private helper function.
		func (db *Database) inTx(ctx context.Context, fn func(tx *Tx) error) error {
		tx, err := db.client.Tx(ctx)
		if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
		}

		if err = fn(tx); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
		return fmt.Errorf("rolling back transaction: %w (original error: %w)", txerr, err)
		}
		return err
		}

		if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
		}
		return nil
		}

		// Client returns the ent Client, considering transaction context.
		func (db *Database) Client(ctx context.Context) *Client {
		if tx := TxFromContext(ctx); tx != nil {
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

    {{- range $n := $.Nodes }}
        {{- $client := print $n.Name "Client" }}
				// {{ $n.Name }} is the client for interacting with the {{ $n.Name }} builders.
				func (db *Database) {{ $n.Name }}(ctx context.Context) *{{ $client }} {
				return db.Client(ctx).{{ $n.Name }}
				}
    {{- end }}

		// Migration executes database schema migrations with the given options.
		//
		// Common migration options:
		//   - schema.WithDropIndex(true): Drops and recreates indexes
		//   - schema.WithDropColumn(true): Drops columns that are no longer defined in schema
		//   - schema.WithForeignKeys(false): Disables foreign key constraints
		//   - schema.WithHooks(hook...): Adds migration hooks
		//
		// Example:
		//
		//   db.Migration(ctx,
		//       schema.WithDropIndex(true),
		//       schema.WithDropColumn(true),
		//   )
		//
		// Note: This method modifies the database schema and should be used with caution in production.
		// It's recommended to review and test migrations in a staging environment first.
		//
		// Warning: WithDropColumn(true) will permanently delete columns and their data.
		// WithDropIndex(true) may temporarily affect query performance.
		func (db *Database) Migration(ctx context.Context, opts ...schema.MigrateOption) error {
		return db.Client(ctx).Schema.Create(ctx, opts...)
		}

{{ end }}