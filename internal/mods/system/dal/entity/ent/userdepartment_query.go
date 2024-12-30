// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userdepartment"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserDepartmentQuery is the builder for querying UserDepartment entities.
type UserDepartmentQuery struct {
	config
	ctx            *QueryContext
	order          []userdepartment.OrderOption
	inters         []Interceptor
	predicates     []predicate.UserDepartment
	withUser       *UserQuery
	withDepartment *DepartmentQuery
	modifiers      []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserDepartmentQuery builder.
func (udq *UserDepartmentQuery) Where(ps ...predicate.UserDepartment) *UserDepartmentQuery {
	udq.predicates = append(udq.predicates, ps...)
	return udq
}

// Limit the number of records to be returned by this query.
func (udq *UserDepartmentQuery) Limit(limit int) *UserDepartmentQuery {
	udq.ctx.Limit = &limit
	return udq
}

// Offset to start from.
func (udq *UserDepartmentQuery) Offset(offset int) *UserDepartmentQuery {
	udq.ctx.Offset = &offset
	return udq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (udq *UserDepartmentQuery) Unique(unique bool) *UserDepartmentQuery {
	udq.ctx.Unique = &unique
	return udq
}

// Order specifies how the records should be ordered.
func (udq *UserDepartmentQuery) Order(o ...userdepartment.OrderOption) *UserDepartmentQuery {
	udq.order = append(udq.order, o...)
	return udq
}

// QueryUser chains the current query on the "user" edge.
func (udq *UserDepartmentQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: udq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := udq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := udq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userdepartment.Table, userdepartment.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userdepartment.UserTable, userdepartment.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(udq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDepartment chains the current query on the "department" edge.
func (udq *UserDepartmentQuery) QueryDepartment() *DepartmentQuery {
	query := (&DepartmentClient{config: udq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := udq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := udq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userdepartment.Table, userdepartment.FieldID, selector),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userdepartment.DepartmentTable, userdepartment.DepartmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(udq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserDepartment entity from the query.
// Returns a *NotFoundError when no UserDepartment was found.
func (udq *UserDepartmentQuery) First(ctx context.Context) (*UserDepartment, error) {
	nodes, err := udq.Limit(1).All(setContextOp(ctx, udq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userdepartment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (udq *UserDepartmentQuery) FirstX(ctx context.Context) *UserDepartment {
	node, err := udq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserDepartment ID from the query.
// Returns a *NotFoundError when no UserDepartment ID was found.
func (udq *UserDepartmentQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = udq.Limit(1).IDs(setContextOp(ctx, udq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userdepartment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (udq *UserDepartmentQuery) FirstIDX(ctx context.Context) int64 {
	id, err := udq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserDepartment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserDepartment entity is found.
// Returns a *NotFoundError when no UserDepartment entities are found.
func (udq *UserDepartmentQuery) Only(ctx context.Context) (*UserDepartment, error) {
	nodes, err := udq.Limit(2).All(setContextOp(ctx, udq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userdepartment.Label}
	default:
		return nil, &NotSingularError{userdepartment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (udq *UserDepartmentQuery) OnlyX(ctx context.Context) *UserDepartment {
	node, err := udq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserDepartment ID in the query.
// Returns a *NotSingularError when more than one UserDepartment ID is found.
// Returns a *NotFoundError when no entities are found.
func (udq *UserDepartmentQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = udq.Limit(2).IDs(setContextOp(ctx, udq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userdepartment.Label}
	default:
		err = &NotSingularError{userdepartment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (udq *UserDepartmentQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := udq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserDepartments.
func (udq *UserDepartmentQuery) All(ctx context.Context) ([]*UserDepartment, error) {
	ctx = setContextOp(ctx, udq.ctx, ent.OpQueryAll)
	if err := udq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserDepartment, *UserDepartmentQuery]()
	return withInterceptors[[]*UserDepartment](ctx, udq, qr, udq.inters)
}

// AllX is like All, but panics if an error occurs.
func (udq *UserDepartmentQuery) AllX(ctx context.Context) []*UserDepartment {
	nodes, err := udq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserDepartment IDs.
func (udq *UserDepartmentQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if udq.ctx.Unique == nil && udq.path != nil {
		udq.Unique(true)
	}
	ctx = setContextOp(ctx, udq.ctx, ent.OpQueryIDs)
	if err = udq.Select(userdepartment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (udq *UserDepartmentQuery) IDsX(ctx context.Context) []int64 {
	ids, err := udq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (udq *UserDepartmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, udq.ctx, ent.OpQueryCount)
	if err := udq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, udq, querierCount[*UserDepartmentQuery](), udq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (udq *UserDepartmentQuery) CountX(ctx context.Context) int {
	count, err := udq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (udq *UserDepartmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, udq.ctx, ent.OpQueryExist)
	switch _, err := udq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (udq *UserDepartmentQuery) ExistX(ctx context.Context) bool {
	exist, err := udq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserDepartmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (udq *UserDepartmentQuery) Clone() *UserDepartmentQuery {
	if udq == nil {
		return nil
	}
	return &UserDepartmentQuery{
		config:         udq.config,
		ctx:            udq.ctx.Clone(),
		order:          append([]userdepartment.OrderOption{}, udq.order...),
		inters:         append([]Interceptor{}, udq.inters...),
		predicates:     append([]predicate.UserDepartment{}, udq.predicates...),
		withUser:       udq.withUser.Clone(),
		withDepartment: udq.withDepartment.Clone(),
		// clone intermediate query.
		sql:       udq.sql.Clone(),
		path:      udq.path,
		modifiers: append([]func(*sql.Selector){}, udq.modifiers...),
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (udq *UserDepartmentQuery) WithUser(opts ...func(*UserQuery)) *UserDepartmentQuery {
	query := (&UserClient{config: udq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	udq.withUser = query
	return udq
}

// WithDepartment tells the query-builder to eager-load the nodes that are connected to
// the "department" edge. The optional arguments are used to configure the query builder of the edge.
func (udq *UserDepartmentQuery) WithDepartment(opts ...func(*DepartmentQuery)) *UserDepartmentQuery {
	query := (&DepartmentClient{config: udq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	udq.withDepartment = query
	return udq
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
//	client.UserDepartment.Query().
//		GroupBy(userdepartment.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (udq *UserDepartmentQuery) GroupBy(field string, fields ...string) *UserDepartmentGroupBy {
	udq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserDepartmentGroupBy{build: udq}
	grbuild.flds = &udq.ctx.Fields
	grbuild.label = userdepartment.Label
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
//	client.UserDepartment.Query().
//		Select(userdepartment.FieldUserID).
//		Scan(ctx, &v)
func (udq *UserDepartmentQuery) Select(fields ...string) *UserDepartmentSelect {
	udq.ctx.Fields = append(udq.ctx.Fields, fields...)
	sbuild := &UserDepartmentSelect{UserDepartmentQuery: udq}
	sbuild.label = userdepartment.Label
	sbuild.flds, sbuild.scan = &udq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserDepartmentSelect configured with the given aggregations.
func (udq *UserDepartmentQuery) Aggregate(fns ...AggregateFunc) *UserDepartmentSelect {
	return udq.Select().Aggregate(fns...)
}

func (udq *UserDepartmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range udq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, udq); err != nil {
				return err
			}
		}
	}
	for _, f := range udq.ctx.Fields {
		if !userdepartment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if udq.path != nil {
		prev, err := udq.path(ctx)
		if err != nil {
			return err
		}
		udq.sql = prev
	}
	return nil
}

func (udq *UserDepartmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserDepartment, error) {
	var (
		nodes       = []*UserDepartment{}
		_spec       = udq.querySpec()
		loadedTypes = [2]bool{
			udq.withUser != nil,
			udq.withDepartment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserDepartment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserDepartment{config: udq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(udq.modifiers) > 0 {
		_spec.Modifiers = udq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, udq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := udq.withUser; query != nil {
		if err := udq.loadUser(ctx, query, nodes, nil,
			func(n *UserDepartment, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := udq.withDepartment; query != nil {
		if err := udq.loadDepartment(ctx, query, nodes, nil,
			func(n *UserDepartment, e *Department) { n.Edges.Department = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (udq *UserDepartmentQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserDepartment, init func(*UserDepartment), assign func(*UserDepartment, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*UserDepartment)
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
func (udq *UserDepartmentQuery) loadDepartment(ctx context.Context, query *DepartmentQuery, nodes []*UserDepartment, init func(*UserDepartment), assign func(*UserDepartment, *Department)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*UserDepartment)
	for i := range nodes {
		fk := nodes[i].DepartmentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(department.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "department_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (udq *UserDepartmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := udq.querySpec()
	if len(udq.modifiers) > 0 {
		_spec.Modifiers = udq.modifiers
	}
	_spec.Node.Columns = udq.ctx.Fields
	if len(udq.ctx.Fields) > 0 {
		_spec.Unique = udq.ctx.Unique != nil && *udq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, udq.driver, _spec)
}

func (udq *UserDepartmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userdepartment.Table, userdepartment.Columns, sqlgraph.NewFieldSpec(userdepartment.FieldID, field.TypeInt64))
	_spec.From = udq.sql
	if unique := udq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if udq.path != nil {
		_spec.Unique = true
	}
	if fields := udq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userdepartment.FieldID)
		for i := range fields {
			if fields[i] != userdepartment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if udq.withUser != nil {
			_spec.Node.AddColumnOnce(userdepartment.FieldUserID)
		}
		if udq.withDepartment != nil {
			_spec.Node.AddColumnOnce(userdepartment.FieldDepartmentID)
		}
	}
	if ps := udq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := udq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := udq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := udq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (udq *UserDepartmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(udq.driver.Dialect())
	t1 := builder.Table(userdepartment.Table)
	columns := udq.ctx.Fields
	if len(columns) == 0 {
		columns = userdepartment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if udq.sql != nil {
		selector = udq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if udq.ctx.Unique != nil && *udq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range udq.modifiers {
		m(selector)
	}
	for _, p := range udq.predicates {
		p(selector)
	}
	for _, p := range udq.order {
		p(selector)
	}
	if offset := udq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := udq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (udq *UserDepartmentQuery) ForUpdate(opts ...sql.LockOption) *UserDepartmentQuery {
	if udq.driver.Dialect() == dialect.Postgres {
		udq.Unique(false)
	}
	udq.modifiers = append(udq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return udq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (udq *UserDepartmentQuery) ForShare(opts ...sql.LockOption) *UserDepartmentQuery {
	if udq.driver.Dialect() == dialect.Postgres {
		udq.Unique(false)
	}
	udq.modifiers = append(udq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return udq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (udq *UserDepartmentQuery) Modify(modifiers ...func(s *sql.Selector)) *UserDepartmentSelect {
	udq.modifiers = append(udq.modifiers, modifiers...)
	return udq.Select()
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
// Example:
//
//	var v []struct {
//	  UserID int64 `json:"user_id,omitempty"`
//	  DepartmentID int64 `json:"department_id,omitempty"`
//	}
//
//	client.UserDepartment.Query().
//	  Omit(
//	  userdepartment.FieldUserID,
//	  userdepartment.FieldDepartmentID,
//	  ).
//	  Scan(ctx, &v)
func (udq *UserDepartmentQuery) Omit(fields ...string) *UserDepartmentSelect {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	for _, col := range userdepartment.Columns {
		if _, ok := omits[col]; !ok {
			udq.ctx.Fields = append(udq.ctx.Fields, col)
		}
	}

	sbuild := &UserDepartmentSelect{UserDepartmentQuery: udq}
	sbuild.label = userdepartment.Label
	sbuild.flds, sbuild.scan = &udq.ctx.Fields, sbuild.Scan
	return sbuild
}

// UserDepartmentGroupBy is the group-by builder for UserDepartment entities.
type UserDepartmentGroupBy struct {
	selector
	build *UserDepartmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (udgb *UserDepartmentGroupBy) Aggregate(fns ...AggregateFunc) *UserDepartmentGroupBy {
	udgb.fns = append(udgb.fns, fns...)
	return udgb
}

// Scan applies the selector query and scans the result into the given value.
func (udgb *UserDepartmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, udgb.build.ctx, ent.OpQueryGroupBy)
	if err := udgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserDepartmentQuery, *UserDepartmentGroupBy](ctx, udgb.build, udgb, udgb.build.inters, v)
}

func (udgb *UserDepartmentGroupBy) sqlScan(ctx context.Context, root *UserDepartmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(udgb.fns))
	for _, fn := range udgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*udgb.flds)+len(udgb.fns))
		for _, f := range *udgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*udgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := udgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserDepartmentSelect is the builder for selecting fields of UserDepartment entities.
type UserDepartmentSelect struct {
	*UserDepartmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uds *UserDepartmentSelect) Aggregate(fns ...AggregateFunc) *UserDepartmentSelect {
	uds.fns = append(uds.fns, fns...)
	return uds
}

// Scan applies the selector query and scans the result into the given value.
func (uds *UserDepartmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uds.ctx, ent.OpQuerySelect)
	if err := uds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserDepartmentQuery, *UserDepartmentSelect](ctx, uds.UserDepartmentQuery, uds, uds.inters, v)
}

func (uds *UserDepartmentSelect) sqlScan(ctx context.Context, root *UserDepartmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uds.fns))
	for _, fn := range uds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (uds *UserDepartmentSelect) Modify(modifiers ...func(s *sql.Selector)) *UserDepartmentSelect {
	uds.modifiers = append(uds.modifiers, modifiers...)
	return uds
}