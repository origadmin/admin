// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolemenu"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleMenuQuery is the builder for querying RoleMenu entities.
type RoleMenuQuery struct {
	config
	ctx        *QueryContext
	order      []rolemenu.OrderOption
	inters     []Interceptor
	predicates []predicate.RoleMenu
	withRole   *RoleQuery
	withMenu   *MenuQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoleMenuQuery builder.
func (rmq *RoleMenuQuery) Where(ps ...predicate.RoleMenu) *RoleMenuQuery {
	rmq.predicates = append(rmq.predicates, ps...)
	return rmq
}

// Limit the number of records to be returned by this query.
func (rmq *RoleMenuQuery) Limit(limit int) *RoleMenuQuery {
	rmq.ctx.Limit = &limit
	return rmq
}

// Offset to start from.
func (rmq *RoleMenuQuery) Offset(offset int) *RoleMenuQuery {
	rmq.ctx.Offset = &offset
	return rmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rmq *RoleMenuQuery) Unique(unique bool) *RoleMenuQuery {
	rmq.ctx.Unique = &unique
	return rmq
}

// Order specifies how the records should be ordered.
func (rmq *RoleMenuQuery) Order(o ...rolemenu.OrderOption) *RoleMenuQuery {
	rmq.order = append(rmq.order, o...)
	return rmq
}

// QueryRole chains the current query on the "role" edge.
func (rmq *RoleMenuQuery) QueryRole() *RoleQuery {
	query := (&RoleClient{config: rmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rolemenu.Table, rolemenu.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, rolemenu.RoleTable, rolemenu.RoleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMenu chains the current query on the "menu" edge.
func (rmq *RoleMenuQuery) QueryMenu() *MenuQuery {
	query := (&MenuClient{config: rmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rolemenu.Table, rolemenu.FieldID, selector),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, rolemenu.MenuTable, rolemenu.MenuColumn),
		)
		fromU = sqlgraph.SetNeighbors(rmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RoleMenu entity from the query.
// Returns a *NotFoundError when no RoleMenu was found.
func (rmq *RoleMenuQuery) First(ctx context.Context) (*RoleMenu, error) {
	nodes, err := rmq.Limit(1).All(setContextOp(ctx, rmq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rolemenu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rmq *RoleMenuQuery) FirstX(ctx context.Context) *RoleMenu {
	node, err := rmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RoleMenu ID from the query.
// Returns a *NotFoundError when no RoleMenu ID was found.
func (rmq *RoleMenuQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rmq.Limit(1).IDs(setContextOp(ctx, rmq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rolemenu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rmq *RoleMenuQuery) FirstIDX(ctx context.Context) int64 {
	id, err := rmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RoleMenu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RoleMenu entity is found.
// Returns a *NotFoundError when no RoleMenu entities are found.
func (rmq *RoleMenuQuery) Only(ctx context.Context) (*RoleMenu, error) {
	nodes, err := rmq.Limit(2).All(setContextOp(ctx, rmq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rolemenu.Label}
	default:
		return nil, &NotSingularError{rolemenu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rmq *RoleMenuQuery) OnlyX(ctx context.Context) *RoleMenu {
	node, err := rmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RoleMenu ID in the query.
// Returns a *NotSingularError when more than one RoleMenu ID is found.
// Returns a *NotFoundError when no entities are found.
func (rmq *RoleMenuQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rmq.Limit(2).IDs(setContextOp(ctx, rmq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rolemenu.Label}
	default:
		err = &NotSingularError{rolemenu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rmq *RoleMenuQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := rmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoleMenus.
func (rmq *RoleMenuQuery) All(ctx context.Context) ([]*RoleMenu, error) {
	ctx = setContextOp(ctx, rmq.ctx, ent.OpQueryAll)
	if err := rmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RoleMenu, *RoleMenuQuery]()
	return withInterceptors[[]*RoleMenu](ctx, rmq, qr, rmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rmq *RoleMenuQuery) AllX(ctx context.Context) []*RoleMenu {
	nodes, err := rmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RoleMenu IDs.
func (rmq *RoleMenuQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if rmq.ctx.Unique == nil && rmq.path != nil {
		rmq.Unique(true)
	}
	ctx = setContextOp(ctx, rmq.ctx, ent.OpQueryIDs)
	if err = rmq.Select(rolemenu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rmq *RoleMenuQuery) IDsX(ctx context.Context) []int64 {
	ids, err := rmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rmq *RoleMenuQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rmq.ctx, ent.OpQueryCount)
	if err := rmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rmq, querierCount[*RoleMenuQuery](), rmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rmq *RoleMenuQuery) CountX(ctx context.Context) int {
	count, err := rmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rmq *RoleMenuQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rmq.ctx, ent.OpQueryExist)
	switch _, err := rmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rmq *RoleMenuQuery) ExistX(ctx context.Context) bool {
	exist, err := rmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoleMenuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rmq *RoleMenuQuery) Clone() *RoleMenuQuery {
	if rmq == nil {
		return nil
	}
	return &RoleMenuQuery{
		config:     rmq.config,
		ctx:        rmq.ctx.Clone(),
		order:      append([]rolemenu.OrderOption{}, rmq.order...),
		inters:     append([]Interceptor{}, rmq.inters...),
		predicates: append([]predicate.RoleMenu{}, rmq.predicates...),
		withRole:   rmq.withRole.Clone(),
		withMenu:   rmq.withMenu.Clone(),
		// clone intermediate query.
		sql:       rmq.sql.Clone(),
		path:      rmq.path,
		modifiers: append([]func(*sql.Selector){}, rmq.modifiers...),
	}
}

// WithRole tells the query-builder to eager-load the nodes that are connected to
// the "role" edge. The optional arguments are used to configure the query builder of the edge.
func (rmq *RoleMenuQuery) WithRole(opts ...func(*RoleQuery)) *RoleMenuQuery {
	query := (&RoleClient{config: rmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rmq.withRole = query
	return rmq
}

// WithMenu tells the query-builder to eager-load the nodes that are connected to
// the "menu" edge. The optional arguments are used to configure the query builder of the edge.
func (rmq *RoleMenuQuery) WithMenu(opts ...func(*MenuQuery)) *RoleMenuQuery {
	query := (&MenuClient{config: rmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rmq.withMenu = query
	return rmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RoleID int64 `json:"role_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoleMenu.Query().
//		GroupBy(rolemenu.FieldRoleID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rmq *RoleMenuQuery) GroupBy(field string, fields ...string) *RoleMenuGroupBy {
	rmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoleMenuGroupBy{build: rmq}
	grbuild.flds = &rmq.ctx.Fields
	grbuild.label = rolemenu.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RoleID int64 `json:"role_id,omitempty"`
//	}
//
//	client.RoleMenu.Query().
//		Select(rolemenu.FieldRoleID).
//		Scan(ctx, &v)
func (rmq *RoleMenuQuery) Select(fields ...string) *RoleMenuSelect {
	rmq.ctx.Fields = append(rmq.ctx.Fields, fields...)
	sbuild := &RoleMenuSelect{RoleMenuQuery: rmq}
	sbuild.label = rolemenu.Label
	sbuild.flds, sbuild.scan = &rmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoleMenuSelect configured with the given aggregations.
func (rmq *RoleMenuQuery) Aggregate(fns ...AggregateFunc) *RoleMenuSelect {
	return rmq.Select().Aggregate(fns...)
}

func (rmq *RoleMenuQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rmq); err != nil {
				return err
			}
		}
	}
	for _, f := range rmq.ctx.Fields {
		if !rolemenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rmq.path != nil {
		prev, err := rmq.path(ctx)
		if err != nil {
			return err
		}
		rmq.sql = prev
	}
	return nil
}

func (rmq *RoleMenuQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RoleMenu, error) {
	var (
		nodes       = []*RoleMenu{}
		_spec       = rmq.querySpec()
		loadedTypes = [2]bool{
			rmq.withRole != nil,
			rmq.withMenu != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RoleMenu).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RoleMenu{config: rmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rmq.modifiers) > 0 {
		_spec.Modifiers = rmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rmq.withRole; query != nil {
		if err := rmq.loadRole(ctx, query, nodes, nil,
			func(n *RoleMenu, e *Role) { n.Edges.Role = e }); err != nil {
			return nil, err
		}
	}
	if query := rmq.withMenu; query != nil {
		if err := rmq.loadMenu(ctx, query, nodes, nil,
			func(n *RoleMenu, e *Menu) { n.Edges.Menu = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rmq *RoleMenuQuery) loadRole(ctx context.Context, query *RoleQuery, nodes []*RoleMenu, init func(*RoleMenu), assign func(*RoleMenu, *Role)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*RoleMenu)
	for i := range nodes {
		fk := nodes[i].RoleID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(role.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "role_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rmq *RoleMenuQuery) loadMenu(ctx context.Context, query *MenuQuery, nodes []*RoleMenu, init func(*RoleMenu), assign func(*RoleMenu, *Menu)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*RoleMenu)
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

func (rmq *RoleMenuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rmq.querySpec()
	if len(rmq.modifiers) > 0 {
		_spec.Modifiers = rmq.modifiers
	}
	_spec.Node.Columns = rmq.ctx.Fields
	if len(rmq.ctx.Fields) > 0 {
		_spec.Unique = rmq.ctx.Unique != nil && *rmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rmq.driver, _spec)
}

func (rmq *RoleMenuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rolemenu.Table, rolemenu.Columns, sqlgraph.NewFieldSpec(rolemenu.FieldID, field.TypeInt64))
	_spec.From = rmq.sql
	if unique := rmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rmq.path != nil {
		_spec.Unique = true
	}
	if fields := rmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rolemenu.FieldID)
		for i := range fields {
			if fields[i] != rolemenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if rmq.withRole != nil {
			_spec.Node.AddColumnOnce(rolemenu.FieldRoleID)
		}
		if rmq.withMenu != nil {
			_spec.Node.AddColumnOnce(rolemenu.FieldMenuID)
		}
	}
	if ps := rmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rmq *RoleMenuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rmq.driver.Dialect())
	t1 := builder.Table(rolemenu.Table)
	columns := rmq.ctx.Fields
	if len(columns) == 0 {
		columns = rolemenu.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rmq.sql != nil {
		selector = rmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rmq.ctx.Unique != nil && *rmq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rmq.modifiers {
		m(selector)
	}
	for _, p := range rmq.predicates {
		p(selector)
	}
	for _, p := range rmq.order {
		p(selector)
	}
	if offset := rmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (rmq *RoleMenuQuery) ForUpdate(opts ...sql.LockOption) *RoleMenuQuery {
	if rmq.driver.Dialect() == dialect.Postgres {
		rmq.Unique(false)
	}
	rmq.modifiers = append(rmq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return rmq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (rmq *RoleMenuQuery) ForShare(opts ...sql.LockOption) *RoleMenuQuery {
	if rmq.driver.Dialect() == dialect.Postgres {
		rmq.Unique(false)
	}
	rmq.modifiers = append(rmq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return rmq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rmq *RoleMenuQuery) Modify(modifiers ...func(s *sql.Selector)) *RoleMenuSelect {
	rmq.modifiers = append(rmq.modifiers, modifiers...)
	return rmq.Select()
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
// Example:
//
//	var v []struct {
//	  RoleID int64 `json:"role_id,omitempty"`
//	  MenuID int64 `json:"menu_id,omitempty"`
//	}
//
//	client.RoleMenu.Query().
//	  Omit(
//	  rolemenu.FieldRoleID,
//	  rolemenu.FieldMenuID,
//	  ).
//	  Scan(ctx, &v)
func (rmq *RoleMenuQuery) Omit(fields ...string) *RoleMenuSelect {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	for _, col := range rolemenu.Columns {
		if _, ok := omits[col]; !ok {
			rmq.ctx.Fields = append(rmq.ctx.Fields, col)
		}
	}

	sbuild := &RoleMenuSelect{RoleMenuQuery: rmq}
	sbuild.label = rolemenu.Label
	sbuild.flds, sbuild.scan = &rmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// RoleMenuGroupBy is the group-by builder for RoleMenu entities.
type RoleMenuGroupBy struct {
	selector
	build *RoleMenuQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rmgb *RoleMenuGroupBy) Aggregate(fns ...AggregateFunc) *RoleMenuGroupBy {
	rmgb.fns = append(rmgb.fns, fns...)
	return rmgb
}

// Scan applies the selector query and scans the result into the given value.
func (rmgb *RoleMenuGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rmgb.build.ctx, ent.OpQueryGroupBy)
	if err := rmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleMenuQuery, *RoleMenuGroupBy](ctx, rmgb.build, rmgb, rmgb.build.inters, v)
}

func (rmgb *RoleMenuGroupBy) sqlScan(ctx context.Context, root *RoleMenuQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rmgb.fns))
	for _, fn := range rmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rmgb.flds)+len(rmgb.fns))
		for _, f := range *rmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoleMenuSelect is the builder for selecting fields of RoleMenu entities.
type RoleMenuSelect struct {
	*RoleMenuQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rms *RoleMenuSelect) Aggregate(fns ...AggregateFunc) *RoleMenuSelect {
	rms.fns = append(rms.fns, fns...)
	return rms
}

// Scan applies the selector query and scans the result into the given value.
func (rms *RoleMenuSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rms.ctx, ent.OpQuerySelect)
	if err := rms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleMenuQuery, *RoleMenuSelect](ctx, rms.RoleMenuQuery, rms, rms.inters, v)
}

func (rms *RoleMenuSelect) sqlScan(ctx context.Context, root *RoleMenuQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rms.fns))
	for _, fn := range rms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rms *RoleMenuSelect) Modify(modifiers ...func(s *sql.Selector)) *RoleMenuSelect {
	rms.modifiers = append(rms.modifiers, modifiers...)
	return rms
}
