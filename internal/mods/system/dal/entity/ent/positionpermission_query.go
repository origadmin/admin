// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/positionpermission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PositionPermissionQuery is the builder for querying PositionPermission entities.
type PositionPermissionQuery struct {
	config
	ctx            *QueryContext
	order          []positionpermission.OrderOption
	inters         []Interceptor
	predicates     []predicate.PositionPermission
	withPosition   *PositionQuery
	withPermission *PermissionQuery
	modifiers      []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PositionPermissionQuery builder.
func (ppq *PositionPermissionQuery) Where(ps ...predicate.PositionPermission) *PositionPermissionQuery {
	ppq.predicates = append(ppq.predicates, ps...)
	return ppq
}

// Limit the number of records to be returned by this query.
func (ppq *PositionPermissionQuery) Limit(limit int) *PositionPermissionQuery {
	ppq.ctx.Limit = &limit
	return ppq
}

// Offset to start from.
func (ppq *PositionPermissionQuery) Offset(offset int) *PositionPermissionQuery {
	ppq.ctx.Offset = &offset
	return ppq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ppq *PositionPermissionQuery) Unique(unique bool) *PositionPermissionQuery {
	ppq.ctx.Unique = &unique
	return ppq
}

// Order specifies how the records should be ordered.
func (ppq *PositionPermissionQuery) Order(o ...positionpermission.OrderOption) *PositionPermissionQuery {
	ppq.order = append(ppq.order, o...)
	return ppq
}

// QueryPosition chains the current query on the "position" edge.
func (ppq *PositionPermissionQuery) QueryPosition() *PositionQuery {
	query := (&PositionClient{config: ppq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(positionpermission.Table, positionpermission.FieldID, selector),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, positionpermission.PositionTable, positionpermission.PositionColumn),
		)
		fromU = sqlgraph.SetNeighbors(ppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPermission chains the current query on the "permission" edge.
func (ppq *PositionPermissionQuery) QueryPermission() *PermissionQuery {
	query := (&PermissionClient{config: ppq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(positionpermission.Table, positionpermission.FieldID, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, positionpermission.PermissionTable, positionpermission.PermissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(ppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PositionPermission entity from the query.
// Returns a *NotFoundError when no PositionPermission was found.
func (ppq *PositionPermissionQuery) First(ctx context.Context) (*PositionPermission, error) {
	nodes, err := ppq.Limit(1).All(setContextOp(ctx, ppq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{positionpermission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ppq *PositionPermissionQuery) FirstX(ctx context.Context) *PositionPermission {
	node, err := ppq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PositionPermission ID from the query.
// Returns a *NotFoundError when no PositionPermission ID was found.
func (ppq *PositionPermissionQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ppq.Limit(1).IDs(setContextOp(ctx, ppq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{positionpermission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ppq *PositionPermissionQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ppq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PositionPermission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PositionPermission entity is found.
// Returns a *NotFoundError when no PositionPermission entities are found.
func (ppq *PositionPermissionQuery) Only(ctx context.Context) (*PositionPermission, error) {
	nodes, err := ppq.Limit(2).All(setContextOp(ctx, ppq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{positionpermission.Label}
	default:
		return nil, &NotSingularError{positionpermission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ppq *PositionPermissionQuery) OnlyX(ctx context.Context) *PositionPermission {
	node, err := ppq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PositionPermission ID in the query.
// Returns a *NotSingularError when more than one PositionPermission ID is found.
// Returns a *NotFoundError when no entities are found.
func (ppq *PositionPermissionQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ppq.Limit(2).IDs(setContextOp(ctx, ppq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{positionpermission.Label}
	default:
		err = &NotSingularError{positionpermission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ppq *PositionPermissionQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ppq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PositionPermissions.
func (ppq *PositionPermissionQuery) All(ctx context.Context) ([]*PositionPermission, error) {
	ctx = setContextOp(ctx, ppq.ctx, ent.OpQueryAll)
	if err := ppq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PositionPermission, *PositionPermissionQuery]()
	return withInterceptors[[]*PositionPermission](ctx, ppq, qr, ppq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ppq *PositionPermissionQuery) AllX(ctx context.Context) []*PositionPermission {
	nodes, err := ppq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PositionPermission IDs.
func (ppq *PositionPermissionQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if ppq.ctx.Unique == nil && ppq.path != nil {
		ppq.Unique(true)
	}
	ctx = setContextOp(ctx, ppq.ctx, ent.OpQueryIDs)
	if err = ppq.Select(positionpermission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ppq *PositionPermissionQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ppq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ppq *PositionPermissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ppq.ctx, ent.OpQueryCount)
	if err := ppq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ppq, querierCount[*PositionPermissionQuery](), ppq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ppq *PositionPermissionQuery) CountX(ctx context.Context) int {
	count, err := ppq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ppq *PositionPermissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ppq.ctx, ent.OpQueryExist)
	switch _, err := ppq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ppq *PositionPermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := ppq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PositionPermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ppq *PositionPermissionQuery) Clone() *PositionPermissionQuery {
	if ppq == nil {
		return nil
	}
	return &PositionPermissionQuery{
		config:         ppq.config,
		ctx:            ppq.ctx.Clone(),
		order:          append([]positionpermission.OrderOption{}, ppq.order...),
		inters:         append([]Interceptor{}, ppq.inters...),
		predicates:     append([]predicate.PositionPermission{}, ppq.predicates...),
		withPosition:   ppq.withPosition.Clone(),
		withPermission: ppq.withPermission.Clone(),
		// clone intermediate query.
		sql:       ppq.sql.Clone(),
		path:      ppq.path,
		modifiers: append([]func(*sql.Selector){}, ppq.modifiers...),
	}
}

// WithPosition tells the query-builder to eager-load the nodes that are connected to
// the "position" edge. The optional arguments are used to configure the query builder of the edge.
func (ppq *PositionPermissionQuery) WithPosition(opts ...func(*PositionQuery)) *PositionPermissionQuery {
	query := (&PositionClient{config: ppq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ppq.withPosition = query
	return ppq
}

// WithPermission tells the query-builder to eager-load the nodes that are connected to
// the "permission" edge. The optional arguments are used to configure the query builder of the edge.
func (ppq *PositionPermissionQuery) WithPermission(opts ...func(*PermissionQuery)) *PositionPermissionQuery {
	query := (&PermissionClient{config: ppq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ppq.withPermission = query
	return ppq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PositionID int64 `json:"position_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PositionPermission.Query().
//		GroupBy(positionpermission.FieldPositionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ppq *PositionPermissionQuery) GroupBy(field string, fields ...string) *PositionPermissionGroupBy {
	ppq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PositionPermissionGroupBy{build: ppq}
	grbuild.flds = &ppq.ctx.Fields
	grbuild.label = positionpermission.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PositionID int64 `json:"position_id,omitempty"`
//	}
//
//	client.PositionPermission.Query().
//		Select(positionpermission.FieldPositionID).
//		Scan(ctx, &v)
func (ppq *PositionPermissionQuery) Select(fields ...string) *PositionPermissionSelect {
	ppq.ctx.Fields = append(ppq.ctx.Fields, fields...)
	sbuild := &PositionPermissionSelect{PositionPermissionQuery: ppq}
	sbuild.label = positionpermission.Label
	sbuild.flds, sbuild.scan = &ppq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PositionPermissionSelect configured with the given aggregations.
func (ppq *PositionPermissionQuery) Aggregate(fns ...AggregateFunc) *PositionPermissionSelect {
	return ppq.Select().Aggregate(fns...)
}

func (ppq *PositionPermissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ppq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ppq); err != nil {
				return err
			}
		}
	}
	for _, f := range ppq.ctx.Fields {
		if !positionpermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ppq.path != nil {
		prev, err := ppq.path(ctx)
		if err != nil {
			return err
		}
		ppq.sql = prev
	}
	return nil
}

func (ppq *PositionPermissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PositionPermission, error) {
	var (
		nodes       = []*PositionPermission{}
		_spec       = ppq.querySpec()
		loadedTypes = [2]bool{
			ppq.withPosition != nil,
			ppq.withPermission != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PositionPermission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PositionPermission{config: ppq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ppq.modifiers) > 0 {
		_spec.Modifiers = ppq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ppq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ppq.withPosition; query != nil {
		if err := ppq.loadPosition(ctx, query, nodes, nil,
			func(n *PositionPermission, e *Position) { n.Edges.Position = e }); err != nil {
			return nil, err
		}
	}
	if query := ppq.withPermission; query != nil {
		if err := ppq.loadPermission(ctx, query, nodes, nil,
			func(n *PositionPermission, e *Permission) { n.Edges.Permission = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ppq *PositionPermissionQuery) loadPosition(ctx context.Context, query *PositionQuery, nodes []*PositionPermission, init func(*PositionPermission), assign func(*PositionPermission, *Position)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*PositionPermission)
	for i := range nodes {
		fk := nodes[i].PositionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(position.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "position_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ppq *PositionPermissionQuery) loadPermission(ctx context.Context, query *PermissionQuery, nodes []*PositionPermission, init func(*PositionPermission), assign func(*PositionPermission, *Permission)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*PositionPermission)
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

func (ppq *PositionPermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ppq.querySpec()
	if len(ppq.modifiers) > 0 {
		_spec.Modifiers = ppq.modifiers
	}
	_spec.Node.Columns = ppq.ctx.Fields
	if len(ppq.ctx.Fields) > 0 {
		_spec.Unique = ppq.ctx.Unique != nil && *ppq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ppq.driver, _spec)
}

func (ppq *PositionPermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(positionpermission.Table, positionpermission.Columns, sqlgraph.NewFieldSpec(positionpermission.FieldID, field.TypeInt64))
	_spec.From = ppq.sql
	if unique := ppq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ppq.path != nil {
		_spec.Unique = true
	}
	if fields := ppq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, positionpermission.FieldID)
		for i := range fields {
			if fields[i] != positionpermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ppq.withPosition != nil {
			_spec.Node.AddColumnOnce(positionpermission.FieldPositionID)
		}
		if ppq.withPermission != nil {
			_spec.Node.AddColumnOnce(positionpermission.FieldPermissionID)
		}
	}
	if ps := ppq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ppq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ppq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ppq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ppq *PositionPermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ppq.driver.Dialect())
	t1 := builder.Table(positionpermission.Table)
	columns := ppq.ctx.Fields
	if len(columns) == 0 {
		columns = positionpermission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ppq.sql != nil {
		selector = ppq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ppq.ctx.Unique != nil && *ppq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ppq.modifiers {
		m(selector)
	}
	for _, p := range ppq.predicates {
		p(selector)
	}
	for _, p := range ppq.order {
		p(selector)
	}
	if offset := ppq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ppq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ppq *PositionPermissionQuery) ForUpdate(opts ...sql.LockOption) *PositionPermissionQuery {
	if ppq.driver.Dialect() == dialect.Postgres {
		ppq.Unique(false)
	}
	ppq.modifiers = append(ppq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ppq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ppq *PositionPermissionQuery) ForShare(opts ...sql.LockOption) *PositionPermissionQuery {
	if ppq.driver.Dialect() == dialect.Postgres {
		ppq.Unique(false)
	}
	ppq.modifiers = append(ppq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ppq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ppq *PositionPermissionQuery) Modify(modifiers ...func(s *sql.Selector)) *PositionPermissionSelect {
	ppq.modifiers = append(ppq.modifiers, modifiers...)
	return ppq.Select()
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
// Example:
//
//	var v []struct {
//	  PositionID int64 `json:"position_id,omitempty"`
//	  PermissionID int64 `json:"permission_id,omitempty"`
//	}
//
//	client.PositionPermission.Query().
//	  Omit(
//	  positionpermission.FieldPositionID,
//	  positionpermission.FieldPermissionID,
//	  ).
//	  Scan(ctx, &v)
func (ppq *PositionPermissionQuery) Omit(fields ...string) *PositionPermissionSelect {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	for _, col := range positionpermission.Columns {
		if _, ok := omits[col]; !ok {
			ppq.ctx.Fields = append(ppq.ctx.Fields, col)
		}
	}

	sbuild := &PositionPermissionSelect{PositionPermissionQuery: ppq}
	sbuild.label = positionpermission.Label
	sbuild.flds, sbuild.scan = &ppq.ctx.Fields, sbuild.Scan
	return sbuild
}

// PositionPermissionGroupBy is the group-by builder for PositionPermission entities.
type PositionPermissionGroupBy struct {
	selector
	build *PositionPermissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ppgb *PositionPermissionGroupBy) Aggregate(fns ...AggregateFunc) *PositionPermissionGroupBy {
	ppgb.fns = append(ppgb.fns, fns...)
	return ppgb
}

// Scan applies the selector query and scans the result into the given value.
func (ppgb *PositionPermissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ppgb.build.ctx, ent.OpQueryGroupBy)
	if err := ppgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PositionPermissionQuery, *PositionPermissionGroupBy](ctx, ppgb.build, ppgb, ppgb.build.inters, v)
}

func (ppgb *PositionPermissionGroupBy) sqlScan(ctx context.Context, root *PositionPermissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ppgb.fns))
	for _, fn := range ppgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ppgb.flds)+len(ppgb.fns))
		for _, f := range *ppgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ppgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ppgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PositionPermissionSelect is the builder for selecting fields of PositionPermission entities.
type PositionPermissionSelect struct {
	*PositionPermissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pps *PositionPermissionSelect) Aggregate(fns ...AggregateFunc) *PositionPermissionSelect {
	pps.fns = append(pps.fns, fns...)
	return pps
}

// Scan applies the selector query and scans the result into the given value.
func (pps *PositionPermissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pps.ctx, ent.OpQuerySelect)
	if err := pps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PositionPermissionQuery, *PositionPermissionSelect](ctx, pps.PositionPermissionQuery, pps, pps.inters, v)
}

func (pps *PositionPermissionSelect) sqlScan(ctx context.Context, root *PositionPermissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pps.fns))
	for _, fn := range pps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pps *PositionPermissionSelect) Modify(modifiers ...func(s *sql.Selector)) *PositionPermissionSelect {
	pps.modifiers = append(pps.modifiers, modifiers...)
	return pps
}
