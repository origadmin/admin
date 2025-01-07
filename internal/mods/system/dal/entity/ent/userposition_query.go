// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userposition"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPositionQuery is the builder for querying UserPosition entities.
type UserPositionQuery struct {
	config
	ctx          *QueryContext
	order        []userposition.OrderOption
	inters       []Interceptor
	predicates   []predicate.UserPosition
	withUser     *UserQuery
	withPosition *PositionQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserPositionQuery builder.
func (upq *UserPositionQuery) Where(ps ...predicate.UserPosition) *UserPositionQuery {
	upq.predicates = append(upq.predicates, ps...)
	return upq
}

// Limit the number of records to be returned by this query.
func (upq *UserPositionQuery) Limit(limit int) *UserPositionQuery {
	upq.ctx.Limit = &limit
	return upq
}

// Offset to start from.
func (upq *UserPositionQuery) Offset(offset int) *UserPositionQuery {
	upq.ctx.Offset = &offset
	return upq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (upq *UserPositionQuery) Unique(unique bool) *UserPositionQuery {
	upq.ctx.Unique = &unique
	return upq
}

// Order specifies how the records should be ordered.
func (upq *UserPositionQuery) Order(o ...userposition.OrderOption) *UserPositionQuery {
	upq.order = append(upq.order, o...)
	return upq
}

// QueryUser chains the current query on the "user" edge.
func (upq *UserPositionQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: upq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := upq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := upq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userposition.Table, userposition.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userposition.UserTable, userposition.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(upq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPosition chains the current query on the "position" edge.
func (upq *UserPositionQuery) QueryPosition() *PositionQuery {
	query := (&PositionClient{config: upq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := upq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := upq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userposition.Table, userposition.FieldID, selector),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userposition.PositionTable, userposition.PositionColumn),
		)
		fromU = sqlgraph.SetNeighbors(upq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserPosition entity from the query.
// Returns a *NotFoundError when no UserPosition was found.
func (upq *UserPositionQuery) First(ctx context.Context) (*UserPosition, error) {
	nodes, err := upq.Limit(1).All(setContextOp(ctx, upq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userposition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (upq *UserPositionQuery) FirstX(ctx context.Context) *UserPosition {
	node, err := upq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserPosition ID from the query.
// Returns a *NotFoundError when no UserPosition ID was found.
func (upq *UserPositionQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = upq.Limit(1).IDs(setContextOp(ctx, upq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userposition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (upq *UserPositionQuery) FirstIDX(ctx context.Context) int64 {
	id, err := upq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserPosition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserPosition entity is found.
// Returns a *NotFoundError when no UserPosition entities are found.
func (upq *UserPositionQuery) Only(ctx context.Context) (*UserPosition, error) {
	nodes, err := upq.Limit(2).All(setContextOp(ctx, upq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userposition.Label}
	default:
		return nil, &NotSingularError{userposition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (upq *UserPositionQuery) OnlyX(ctx context.Context) *UserPosition {
	node, err := upq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserPosition ID in the query.
// Returns a *NotSingularError when more than one UserPosition ID is found.
// Returns a *NotFoundError when no entities are found.
func (upq *UserPositionQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = upq.Limit(2).IDs(setContextOp(ctx, upq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userposition.Label}
	default:
		err = &NotSingularError{userposition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (upq *UserPositionQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := upq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserPositions.
func (upq *UserPositionQuery) All(ctx context.Context) ([]*UserPosition, error) {
	ctx = setContextOp(ctx, upq.ctx, ent.OpQueryAll)
	if err := upq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserPosition, *UserPositionQuery]()
	return withInterceptors[[]*UserPosition](ctx, upq, qr, upq.inters)
}

// AllX is like All, but panics if an error occurs.
func (upq *UserPositionQuery) AllX(ctx context.Context) []*UserPosition {
	nodes, err := upq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserPosition IDs.
func (upq *UserPositionQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if upq.ctx.Unique == nil && upq.path != nil {
		upq.Unique(true)
	}
	ctx = setContextOp(ctx, upq.ctx, ent.OpQueryIDs)
	if err = upq.Select(userposition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (upq *UserPositionQuery) IDsX(ctx context.Context) []int64 {
	ids, err := upq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (upq *UserPositionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, upq.ctx, ent.OpQueryCount)
	if err := upq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, upq, querierCount[*UserPositionQuery](), upq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (upq *UserPositionQuery) CountX(ctx context.Context) int {
	count, err := upq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (upq *UserPositionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, upq.ctx, ent.OpQueryExist)
	switch _, err := upq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (upq *UserPositionQuery) ExistX(ctx context.Context) bool {
	exist, err := upq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserPositionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (upq *UserPositionQuery) Clone() *UserPositionQuery {
	if upq == nil {
		return nil
	}
	return &UserPositionQuery{
		config:       upq.config,
		ctx:          upq.ctx.Clone(),
		order:        append([]userposition.OrderOption{}, upq.order...),
		inters:       append([]Interceptor{}, upq.inters...),
		predicates:   append([]predicate.UserPosition{}, upq.predicates...),
		withUser:     upq.withUser.Clone(),
		withPosition: upq.withPosition.Clone(),
		// clone intermediate query.
		sql:       upq.sql.Clone(),
		path:      upq.path,
		modifiers: append([]func(*sql.Selector){}, upq.modifiers...),
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (upq *UserPositionQuery) WithUser(opts ...func(*UserQuery)) *UserPositionQuery {
	query := (&UserClient{config: upq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	upq.withUser = query
	return upq
}

// WithPosition tells the query-builder to eager-load the nodes that are connected to
// the "position" edge. The optional arguments are used to configure the query builder of the edge.
func (upq *UserPositionQuery) WithPosition(opts ...func(*PositionQuery)) *UserPositionQuery {
	query := (&PositionClient{config: upq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	upq.withPosition = query
	return upq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserPosition.Query().
//		GroupBy(userposition.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (upq *UserPositionQuery) GroupBy(field string, fields ...string) *UserPositionGroupBy {
	upq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserPositionGroupBy{build: upq}
	grbuild.flds = &upq.ctx.Fields
	grbuild.label = userposition.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//	}
//
//	client.UserPosition.Query().
//		Select(userposition.FieldUserID).
//		Scan(ctx, &v)
func (upq *UserPositionQuery) Select(fields ...string) *UserPositionSelect {
	upq.ctx.Fields = append(upq.ctx.Fields, fields...)
	sbuild := &UserPositionSelect{UserPositionQuery: upq}
	sbuild.label = userposition.Label
	sbuild.flds, sbuild.scan = &upq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserPositionSelect configured with the given aggregations.
func (upq *UserPositionQuery) Aggregate(fns ...AggregateFunc) *UserPositionSelect {
	return upq.Select().Aggregate(fns...)
}

func (upq *UserPositionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range upq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, upq); err != nil {
				return err
			}
		}
	}
	for _, f := range upq.ctx.Fields {
		if !userposition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if upq.path != nil {
		prev, err := upq.path(ctx)
		if err != nil {
			return err
		}
		upq.sql = prev
	}
	return nil
}

func (upq *UserPositionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserPosition, error) {
	var (
		nodes       = []*UserPosition{}
		_spec       = upq.querySpec()
		loadedTypes = [2]bool{
			upq.withUser != nil,
			upq.withPosition != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserPosition).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserPosition{config: upq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(upq.modifiers) > 0 {
		_spec.Modifiers = upq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, upq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := upq.withUser; query != nil {
		if err := upq.loadUser(ctx, query, nodes, nil,
			func(n *UserPosition, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := upq.withPosition; query != nil {
		if err := upq.loadPosition(ctx, query, nodes, nil,
			func(n *UserPosition, e *Position) { n.Edges.Position = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (upq *UserPositionQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserPosition, init func(*UserPosition), assign func(*UserPosition, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*UserPosition)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (upq *UserPositionQuery) loadPosition(ctx context.Context, query *PositionQuery, nodes []*UserPosition, init func(*UserPosition), assign func(*UserPosition, *Position)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*UserPosition)
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

func (upq *UserPositionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := upq.querySpec()
	if len(upq.modifiers) > 0 {
		_spec.Modifiers = upq.modifiers
	}
	_spec.Node.Columns = upq.ctx.Fields
	if len(upq.ctx.Fields) > 0 {
		_spec.Unique = upq.ctx.Unique != nil && *upq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, upq.driver, _spec)
}

func (upq *UserPositionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userposition.Table, userposition.Columns, sqlgraph.NewFieldSpec(userposition.FieldID, field.TypeInt64))
	_spec.From = upq.sql
	if unique := upq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if upq.path != nil {
		_spec.Unique = true
	}
	if fields := upq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userposition.FieldID)
		for i := range fields {
			if fields[i] != userposition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if upq.withUser != nil {
			_spec.Node.AddColumnOnce(userposition.FieldUserID)
		}
		if upq.withPosition != nil {
			_spec.Node.AddColumnOnce(userposition.FieldPositionID)
		}
	}
	if ps := upq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := upq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := upq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := upq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (upq *UserPositionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(upq.driver.Dialect())
	t1 := builder.Table(userposition.Table)
	columns := upq.ctx.Fields
	if len(columns) == 0 {
		columns = userposition.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if upq.sql != nil {
		selector = upq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if upq.ctx.Unique != nil && *upq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range upq.modifiers {
		m(selector)
	}
	for _, p := range upq.predicates {
		p(selector)
	}
	for _, p := range upq.order {
		p(selector)
	}
	if offset := upq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := upq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (upq *UserPositionQuery) ForUpdate(opts ...sql.LockOption) *UserPositionQuery {
	if upq.driver.Dialect() == dialect.Postgres {
		upq.Unique(false)
	}
	upq.modifiers = append(upq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return upq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (upq *UserPositionQuery) ForShare(opts ...sql.LockOption) *UserPositionQuery {
	if upq.driver.Dialect() == dialect.Postgres {
		upq.Unique(false)
	}
	upq.modifiers = append(upq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return upq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (upq *UserPositionQuery) Modify(modifiers ...func(s *sql.Selector)) *UserPositionSelect {
	upq.modifiers = append(upq.modifiers, modifiers...)
	return upq.Select()
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
// Example:
//
//	var v []struct {
//	  UserID int64 `json:"user_id,omitempty"`
//	  PositionID int64 `json:"position_id,omitempty"`
//	}
//
//	client.UserPosition.Query().
//	  Omit(
//	  userposition.FieldUserID,
//	  userposition.FieldPositionID,
//	  ).
//	  Scan(ctx, &v)
func (upq *UserPositionQuery) Omit(fields ...string) *UserPositionSelect {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	for _, col := range userposition.Columns {
		if _, ok := omits[col]; !ok {
			upq.ctx.Fields = append(upq.ctx.Fields, col)
		}
	}

	sbuild := &UserPositionSelect{UserPositionQuery: upq}
	sbuild.label = userposition.Label
	sbuild.flds, sbuild.scan = &upq.ctx.Fields, sbuild.Scan
	return sbuild
}

// UserPositionGroupBy is the group-by builder for UserPosition entities.
type UserPositionGroupBy struct {
	selector
	build *UserPositionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (upgb *UserPositionGroupBy) Aggregate(fns ...AggregateFunc) *UserPositionGroupBy {
	upgb.fns = append(upgb.fns, fns...)
	return upgb
}

// Scan applies the selector query and scans the result into the given value.
func (upgb *UserPositionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, upgb.build.ctx, ent.OpQueryGroupBy)
	if err := upgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserPositionQuery, *UserPositionGroupBy](ctx, upgb.build, upgb, upgb.build.inters, v)
}

func (upgb *UserPositionGroupBy) sqlScan(ctx context.Context, root *UserPositionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(upgb.fns))
	for _, fn := range upgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*upgb.flds)+len(upgb.fns))
		for _, f := range *upgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*upgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := upgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserPositionSelect is the builder for selecting fields of UserPosition entities.
type UserPositionSelect struct {
	*UserPositionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ups *UserPositionSelect) Aggregate(fns ...AggregateFunc) *UserPositionSelect {
	ups.fns = append(ups.fns, fns...)
	return ups
}

// Scan applies the selector query and scans the result into the given value.
func (ups *UserPositionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ups.ctx, ent.OpQuerySelect)
	if err := ups.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserPositionQuery, *UserPositionSelect](ctx, ups.UserPositionQuery, ups, ups.inters, v)
}

func (ups *UserPositionSelect) sqlScan(ctx context.Context, root *UserPositionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ups.fns))
	for _, fn := range ups.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ups.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ups.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ups *UserPositionSelect) Modify(modifiers ...func(s *sql.Selector)) *UserPositionSelect {
	ups.modifiers = append(ups.modifiers, modifiers...)
	return ups
}
