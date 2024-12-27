// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menupermission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuPermissionQuery is the builder for querying MenuPermission entities.
type MenuPermissionQuery struct {
	config
	ctx            *QueryContext
	order          []menupermission.OrderOption
	inters         []Interceptor
	predicates     []predicate.MenuPermission
	withMenu       *MenuQuery
	withPermission *PermissionQuery
	modifiers      []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MenuPermissionQuery builder.
func (mpq *MenuPermissionQuery) Where(ps ...predicate.MenuPermission) *MenuPermissionQuery {
	mpq.predicates = append(mpq.predicates, ps...)
	return mpq
}

// Limit the number of records to be returned by this query.
func (mpq *MenuPermissionQuery) Limit(limit int) *MenuPermissionQuery {
	mpq.ctx.Limit = &limit
	return mpq
}

// Offset to start from.
func (mpq *MenuPermissionQuery) Offset(offset int) *MenuPermissionQuery {
	mpq.ctx.Offset = &offset
	return mpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mpq *MenuPermissionQuery) Unique(unique bool) *MenuPermissionQuery {
	mpq.ctx.Unique = &unique
	return mpq
}

// Order specifies how the records should be ordered.
func (mpq *MenuPermissionQuery) Order(o ...menupermission.OrderOption) *MenuPermissionQuery {
	mpq.order = append(mpq.order, o...)
	return mpq
}

// QueryMenu chains the current query on the "menu" edge.
func (mpq *MenuPermissionQuery) QueryMenu() *MenuQuery {
	query := (&MenuClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menupermission.Table, menupermission.FieldID, selector),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, menupermission.MenuTable, menupermission.MenuColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPermission chains the current query on the "permission" edge.
func (mpq *MenuPermissionQuery) QueryPermission() *PermissionQuery {
	query := (&PermissionClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menupermission.Table, menupermission.FieldID, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, menupermission.PermissionTable, menupermission.PermissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MenuPermission entity from the query.
// Returns a *NotFoundError when no MenuPermission was found.
func (mpq *MenuPermissionQuery) First(ctx context.Context) (*MenuPermission, error) {
	nodes, err := mpq.Limit(1).All(setContextOp(ctx, mpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{menupermission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mpq *MenuPermissionQuery) FirstX(ctx context.Context) *MenuPermission {
	node, err := mpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MenuPermission ID from the query.
// Returns a *NotFoundError when no MenuPermission ID was found.
func (mpq *MenuPermissionQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(1).IDs(setContextOp(ctx, mpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{menupermission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mpq *MenuPermissionQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MenuPermission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MenuPermission entity is found.
// Returns a *NotFoundError when no MenuPermission entities are found.
func (mpq *MenuPermissionQuery) Only(ctx context.Context) (*MenuPermission, error) {
	nodes, err := mpq.Limit(2).All(setContextOp(ctx, mpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{menupermission.Label}
	default:
		return nil, &NotSingularError{menupermission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mpq *MenuPermissionQuery) OnlyX(ctx context.Context) *MenuPermission {
	node, err := mpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MenuPermission ID in the query.
// Returns a *NotSingularError when more than one MenuPermission ID is found.
// Returns a *NotFoundError when no entities are found.
func (mpq *MenuPermissionQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(2).IDs(setContextOp(ctx, mpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{menupermission.Label}
	default:
		err = &NotSingularError{menupermission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mpq *MenuPermissionQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MenuPermissions.
func (mpq *MenuPermissionQuery) All(ctx context.Context) ([]*MenuPermission, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryAll)
	if err := mpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MenuPermission, *MenuPermissionQuery]()
	return withInterceptors[[]*MenuPermission](ctx, mpq, qr, mpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mpq *MenuPermissionQuery) AllX(ctx context.Context) []*MenuPermission {
	nodes, err := mpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MenuPermission IDs.
func (mpq *MenuPermissionQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mpq.ctx.Unique == nil && mpq.path != nil {
		mpq.Unique(true)
	}
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryIDs)
	if err = mpq.Select(menupermission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mpq *MenuPermissionQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mpq *MenuPermissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryCount)
	if err := mpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mpq, querierCount[*MenuPermissionQuery](), mpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mpq *MenuPermissionQuery) CountX(ctx context.Context) int {
	count, err := mpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mpq *MenuPermissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryExist)
	switch _, err := mpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mpq *MenuPermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := mpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MenuPermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mpq *MenuPermissionQuery) Clone() *MenuPermissionQuery {
	if mpq == nil {
		return nil
	}
	return &MenuPermissionQuery{
		config:         mpq.config,
		ctx:            mpq.ctx.Clone(),
		order:          append([]menupermission.OrderOption{}, mpq.order...),
		inters:         append([]Interceptor{}, mpq.inters...),
		predicates:     append([]predicate.MenuPermission{}, mpq.predicates...),
		withMenu:       mpq.withMenu.Clone(),
		withPermission: mpq.withPermission.Clone(),
		// clone intermediate query.
		sql:       mpq.sql.Clone(),
		path:      mpq.path,
		modifiers: append([]func(*sql.Selector){}, mpq.modifiers...),
	}
}

// WithMenu tells the query-builder to eager-load the nodes that are connected to
// the "menu" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MenuPermissionQuery) WithMenu(opts ...func(*MenuQuery)) *MenuPermissionQuery {
	query := (&MenuClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMenu = query
	return mpq
}

// WithPermission tells the query-builder to eager-load the nodes that are connected to
// the "permission" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MenuPermissionQuery) WithPermission(opts ...func(*PermissionQuery)) *MenuPermissionQuery {
	query := (&PermissionClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withPermission = query
	return mpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MenuID int64 `json:"menu_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MenuPermission.Query().
//		GroupBy(menupermission.FieldMenuID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mpq *MenuPermissionQuery) GroupBy(field string, fields ...string) *MenuPermissionGroupBy {
	mpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MenuPermissionGroupBy{build: mpq}
	grbuild.flds = &mpq.ctx.Fields
	grbuild.label = menupermission.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MenuID int64 `json:"menu_id,omitempty"`
//	}
//
//	client.MenuPermission.Query().
//		Select(menupermission.FieldMenuID).
//		Scan(ctx, &v)
func (mpq *MenuPermissionQuery) Select(fields ...string) *MenuPermissionSelect {
	mpq.ctx.Fields = append(mpq.ctx.Fields, fields...)
	sbuild := &MenuPermissionSelect{MenuPermissionQuery: mpq}
	sbuild.label = menupermission.Label
	sbuild.flds, sbuild.scan = &mpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MenuPermissionSelect configured with the given aggregations.
func (mpq *MenuPermissionQuery) Aggregate(fns ...AggregateFunc) *MenuPermissionSelect {
	return mpq.Select().Aggregate(fns...)
}

func (mpq *MenuPermissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mpq); err != nil {
				return err
			}
		}
	}
	for _, f := range mpq.ctx.Fields {
		if !menupermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mpq.path != nil {
		prev, err := mpq.path(ctx)
		if err != nil {
			return err
		}
		mpq.sql = prev
	}
	return nil
}

func (mpq *MenuPermissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MenuPermission, error) {
	var (
		nodes       = []*MenuPermission{}
		_spec       = mpq.querySpec()
		loadedTypes = [2]bool{
			mpq.withMenu != nil,
			mpq.withPermission != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MenuPermission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MenuPermission{config: mpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mpq.modifiers) > 0 {
		_spec.Modifiers = mpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mpq.withMenu; query != nil {
		if err := mpq.loadMenu(ctx, query, nodes, nil,
			func(n *MenuPermission, e *Menu) { n.Edges.Menu = e }); err != nil {
			return nil, err
		}
	}
	if query := mpq.withPermission; query != nil {
		if err := mpq.loadPermission(ctx, query, nodes, nil,
			func(n *MenuPermission, e *Permission) { n.Edges.Permission = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mpq *MenuPermissionQuery) loadMenu(ctx context.Context, query *MenuQuery, nodes []*MenuPermission, init func(*MenuPermission), assign func(*MenuPermission, *Menu)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MenuPermission)
	for i := range nodes {
		fk := nodes[i].MenuID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(menu.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "menu_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mpq *MenuPermissionQuery) loadPermission(ctx context.Context, query *PermissionQuery, nodes []*MenuPermission, init func(*MenuPermission), assign func(*MenuPermission, *Permission)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MenuPermission)
	for i := range nodes {
		fk := nodes[i].PermissionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(permission.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "permission_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mpq *MenuPermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mpq.querySpec()
	if len(mpq.modifiers) > 0 {
		_spec.Modifiers = mpq.modifiers
	}
	_spec.Node.Columns = mpq.ctx.Fields
	if len(mpq.ctx.Fields) > 0 {
		_spec.Unique = mpq.ctx.Unique != nil && *mpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mpq.driver, _spec)
}

func (mpq *MenuPermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(menupermission.Table, menupermission.Columns, sqlgraph.NewFieldSpec(menupermission.FieldID, field.TypeInt64))
	_spec.From = mpq.sql
	if unique := mpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mpq.path != nil {
		_spec.Unique = true
	}
	if fields := mpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menupermission.FieldID)
		for i := range fields {
			if fields[i] != menupermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mpq.withMenu != nil {
			_spec.Node.AddColumnOnce(menupermission.FieldMenuID)
		}
		if mpq.withPermission != nil {
			_spec.Node.AddColumnOnce(menupermission.FieldPermissionID)
		}
	}
	if ps := mpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mpq *MenuPermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mpq.driver.Dialect())
	t1 := builder.Table(menupermission.Table)
	columns := mpq.ctx.Fields
	if len(columns) == 0 {
		columns = menupermission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mpq.sql != nil {
		selector = mpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mpq.ctx.Unique != nil && *mpq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mpq.modifiers {
		m(selector)
	}
	for _, p := range mpq.predicates {
		p(selector)
	}
	for _, p := range mpq.order {
		p(selector)
	}
	if offset := mpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (mpq *MenuPermissionQuery) ForUpdate(opts ...sql.LockOption) *MenuPermissionQuery {
	if mpq.driver.Dialect() == dialect.Postgres {
		mpq.Unique(false)
	}
	mpq.modifiers = append(mpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return mpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (mpq *MenuPermissionQuery) ForShare(opts ...sql.LockOption) *MenuPermissionQuery {
	if mpq.driver.Dialect() == dialect.Postgres {
		mpq.Unique(false)
	}
	mpq.modifiers = append(mpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return mpq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mpq *MenuPermissionQuery) Modify(modifiers ...func(s *sql.Selector)) *MenuPermissionSelect {
	mpq.modifiers = append(mpq.modifiers, modifiers...)
	return mpq.Select()
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
// Example:
//
//	var v []struct {
//	  MenuID int64 `json:"menu_id,omitempty"`
//	  PermissionID int64 `json:"permission_id,omitempty"`
//	}
//
//	client.MenuPermission.Query().
//	  Omit(
//	  menupermission.FieldMenuID,
//	  menupermission.FieldPermissionID,
//	  ).
//	  Scan(ctx, &v)
func (mpq *MenuPermissionQuery) Omit(fields ...string) *MenuPermissionSelect {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	for _, col := range menupermission.Columns {
		if _, ok := omits[col]; !ok {
			mpq.ctx.Fields = append(mpq.ctx.Fields, col)
		}
	}

	sbuild := &MenuPermissionSelect{MenuPermissionQuery: mpq}
	sbuild.label = menupermission.Label
	sbuild.flds, sbuild.scan = &mpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// MenuPermissionGroupBy is the group-by builder for MenuPermission entities.
type MenuPermissionGroupBy struct {
	selector
	build *MenuPermissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mpgb *MenuPermissionGroupBy) Aggregate(fns ...AggregateFunc) *MenuPermissionGroupBy {
	mpgb.fns = append(mpgb.fns, fns...)
	return mpgb
}

// Scan applies the selector query and scans the result into the given value.
func (mpgb *MenuPermissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mpgb.build.ctx, ent.OpQueryGroupBy)
	if err := mpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MenuPermissionQuery, *MenuPermissionGroupBy](ctx, mpgb.build, mpgb, mpgb.build.inters, v)
}

func (mpgb *MenuPermissionGroupBy) sqlScan(ctx context.Context, root *MenuPermissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mpgb.fns))
	for _, fn := range mpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mpgb.flds)+len(mpgb.fns))
		for _, f := range *mpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MenuPermissionSelect is the builder for selecting fields of MenuPermission entities.
type MenuPermissionSelect struct {
	*MenuPermissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mps *MenuPermissionSelect) Aggregate(fns ...AggregateFunc) *MenuPermissionSelect {
	mps.fns = append(mps.fns, fns...)
	return mps
}

// Scan applies the selector query and scans the result into the given value.
func (mps *MenuPermissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mps.ctx, ent.OpQuerySelect)
	if err := mps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MenuPermissionQuery, *MenuPermissionSelect](ctx, mps.MenuPermissionQuery, mps, mps.inters, v)
}

func (mps *MenuPermissionSelect) sqlScan(ctx context.Context, root *MenuPermissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mps.fns))
	for _, fn := range mps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mps *MenuPermissionSelect) Modify(modifiers ...func(s *sql.Selector)) *MenuPermissionSelect {
	mps.modifiers = append(mps.modifiers, modifiers...)
	return mps
}
