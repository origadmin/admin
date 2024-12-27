// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionResourceQuery is the builder for querying PermissionResource entities.
type PermissionResourceQuery struct {
	config
	ctx            *QueryContext
	order          []permissionresource.OrderOption
	inters         []Interceptor
	predicates     []predicate.PermissionResource
	withPermission *PermissionQuery
	withResource   *ResourceQuery
	modifiers      []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PermissionResourceQuery builder.
func (prq *PermissionResourceQuery) Where(ps ...predicate.PermissionResource) *PermissionResourceQuery {
	prq.predicates = append(prq.predicates, ps...)
	return prq
}

// Limit the number of records to be returned by this query.
func (prq *PermissionResourceQuery) Limit(limit int) *PermissionResourceQuery {
	prq.ctx.Limit = &limit
	return prq
}

// Offset to start from.
func (prq *PermissionResourceQuery) Offset(offset int) *PermissionResourceQuery {
	prq.ctx.Offset = &offset
	return prq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prq *PermissionResourceQuery) Unique(unique bool) *PermissionResourceQuery {
	prq.ctx.Unique = &unique
	return prq
}

// Order specifies how the records should be ordered.
func (prq *PermissionResourceQuery) Order(o ...permissionresource.OrderOption) *PermissionResourceQuery {
	prq.order = append(prq.order, o...)
	return prq
}

// QueryPermission chains the current query on the "permission" edge.
func (prq *PermissionResourceQuery) QueryPermission() *PermissionQuery {
	query := (&PermissionClient{config: prq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permissionresource.Table, permissionresource.FieldID, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, permissionresource.PermissionTable, permissionresource.PermissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryResource chains the current query on the "resource" edge.
func (prq *PermissionResourceQuery) QueryResource() *ResourceQuery {
	query := (&ResourceClient{config: prq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permissionresource.Table, permissionresource.FieldID, selector),
			sqlgraph.To(resource.Table, resource.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, permissionresource.ResourceTable, permissionresource.ResourceColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PermissionResource entity from the query.
// Returns a *NotFoundError when no PermissionResource was found.
func (prq *PermissionResourceQuery) First(ctx context.Context) (*PermissionResource, error) {
	nodes, err := prq.Limit(1).All(setContextOp(ctx, prq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{permissionresource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prq *PermissionResourceQuery) FirstX(ctx context.Context) *PermissionResource {
	node, err := prq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PermissionResource ID from the query.
// Returns a *NotFoundError when no PermissionResource ID was found.
func (prq *PermissionResourceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(1).IDs(setContextOp(ctx, prq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{permissionresource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prq *PermissionResourceQuery) FirstIDX(ctx context.Context) int {
	id, err := prq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PermissionResource entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PermissionResource entity is found.
// Returns a *NotFoundError when no PermissionResource entities are found.
func (prq *PermissionResourceQuery) Only(ctx context.Context) (*PermissionResource, error) {
	nodes, err := prq.Limit(2).All(setContextOp(ctx, prq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{permissionresource.Label}
	default:
		return nil, &NotSingularError{permissionresource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prq *PermissionResourceQuery) OnlyX(ctx context.Context) *PermissionResource {
	node, err := prq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PermissionResource ID in the query.
// Returns a *NotSingularError when more than one PermissionResource ID is found.
// Returns a *NotFoundError when no entities are found.
func (prq *PermissionResourceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(2).IDs(setContextOp(ctx, prq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{permissionresource.Label}
	default:
		err = &NotSingularError{permissionresource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prq *PermissionResourceQuery) OnlyIDX(ctx context.Context) int {
	id, err := prq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PermissionResources.
func (prq *PermissionResourceQuery) All(ctx context.Context) ([]*PermissionResource, error) {
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryAll)
	if err := prq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PermissionResource, *PermissionResourceQuery]()
	return withInterceptors[[]*PermissionResource](ctx, prq, qr, prq.inters)
}

// AllX is like All, but panics if an error occurs.
func (prq *PermissionResourceQuery) AllX(ctx context.Context) []*PermissionResource {
	nodes, err := prq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PermissionResource IDs.
func (prq *PermissionResourceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if prq.ctx.Unique == nil && prq.path != nil {
		prq.Unique(true)
	}
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryIDs)
	if err = prq.Select(permissionresource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prq *PermissionResourceQuery) IDsX(ctx context.Context) []int {
	ids, err := prq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prq *PermissionResourceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryCount)
	if err := prq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, prq, querierCount[*PermissionResourceQuery](), prq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (prq *PermissionResourceQuery) CountX(ctx context.Context) int {
	count, err := prq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prq *PermissionResourceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, prq.ctx, ent.OpQueryExist)
	switch _, err := prq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (prq *PermissionResourceQuery) ExistX(ctx context.Context) bool {
	exist, err := prq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PermissionResourceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prq *PermissionResourceQuery) Clone() *PermissionResourceQuery {
	if prq == nil {
		return nil
	}
	return &PermissionResourceQuery{
		config:         prq.config,
		ctx:            prq.ctx.Clone(),
		order:          append([]permissionresource.OrderOption{}, prq.order...),
		inters:         append([]Interceptor{}, prq.inters...),
		predicates:     append([]predicate.PermissionResource{}, prq.predicates...),
		withPermission: prq.withPermission.Clone(),
		withResource:   prq.withResource.Clone(),
		// clone intermediate query.
		sql:       prq.sql.Clone(),
		path:      prq.path,
		modifiers: append([]func(*sql.Selector){}, prq.modifiers...),
	}
}

// WithPermission tells the query-builder to eager-load the nodes that are connected to
// the "permission" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *PermissionResourceQuery) WithPermission(opts ...func(*PermissionQuery)) *PermissionResourceQuery {
	query := (&PermissionClient{config: prq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	prq.withPermission = query
	return prq
}

// WithResource tells the query-builder to eager-load the nodes that are connected to
// the "resource" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *PermissionResourceQuery) WithResource(opts ...func(*ResourceQuery)) *PermissionResourceQuery {
	query := (&ResourceClient{config: prq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	prq.withResource = query
	return prq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PermissionID string `json:"permission_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PermissionResource.Query().
//		GroupBy(permissionresource.FieldPermissionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (prq *PermissionResourceQuery) GroupBy(field string, fields ...string) *PermissionResourceGroupBy {
	prq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PermissionResourceGroupBy{build: prq}
	grbuild.flds = &prq.ctx.Fields
	grbuild.label = permissionresource.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PermissionID string `json:"permission_id,omitempty"`
//	}
//
//	client.PermissionResource.Query().
//		Select(permissionresource.FieldPermissionID).
//		Scan(ctx, &v)
func (prq *PermissionResourceQuery) Select(fields ...string) *PermissionResourceSelect {
	prq.ctx.Fields = append(prq.ctx.Fields, fields...)
	sbuild := &PermissionResourceSelect{PermissionResourceQuery: prq}
	sbuild.label = permissionresource.Label
	sbuild.flds, sbuild.scan = &prq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PermissionResourceSelect configured with the given aggregations.
func (prq *PermissionResourceQuery) Aggregate(fns ...AggregateFunc) *PermissionResourceSelect {
	return prq.Select().Aggregate(fns...)
}

func (prq *PermissionResourceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range prq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, prq); err != nil {
				return err
			}
		}
	}
	for _, f := range prq.ctx.Fields {
		if !permissionresource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if prq.path != nil {
		prev, err := prq.path(ctx)
		if err != nil {
			return err
		}
		prq.sql = prev
	}
	return nil
}

func (prq *PermissionResourceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PermissionResource, error) {
	var (
		nodes       = []*PermissionResource{}
		_spec       = prq.querySpec()
		loadedTypes = [2]bool{
			prq.withPermission != nil,
			prq.withResource != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PermissionResource).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PermissionResource{config: prq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(prq.modifiers) > 0 {
		_spec.Modifiers = prq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, prq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := prq.withPermission; query != nil {
		if err := prq.loadPermission(ctx, query, nodes, nil,
			func(n *PermissionResource, e *Permission) { n.Edges.Permission = e }); err != nil {
			return nil, err
		}
	}
	if query := prq.withResource; query != nil {
		if err := prq.loadResource(ctx, query, nodes, nil,
			func(n *PermissionResource, e *Resource) { n.Edges.Resource = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (prq *PermissionResourceQuery) loadPermission(ctx context.Context, query *PermissionQuery, nodes []*PermissionResource, init func(*PermissionResource), assign func(*PermissionResource, *Permission)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PermissionResource)
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
func (prq *PermissionResourceQuery) loadResource(ctx context.Context, query *ResourceQuery, nodes []*PermissionResource, init func(*PermissionResource), assign func(*PermissionResource, *Resource)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*PermissionResource)
	for i := range nodes {
		fk := nodes[i].ResourceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(resource.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "resource_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (prq *PermissionResourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prq.querySpec()
	if len(prq.modifiers) > 0 {
		_spec.Modifiers = prq.modifiers
	}
	_spec.Node.Columns = prq.ctx.Fields
	if len(prq.ctx.Fields) > 0 {
		_spec.Unique = prq.ctx.Unique != nil && *prq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, prq.driver, _spec)
}

func (prq *PermissionResourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(permissionresource.Table, permissionresource.Columns, sqlgraph.NewFieldSpec(permissionresource.FieldID, field.TypeInt))
	_spec.From = prq.sql
	if unique := prq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if prq.path != nil {
		_spec.Unique = true
	}
	if fields := prq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permissionresource.FieldID)
		for i := range fields {
			if fields[i] != permissionresource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if prq.withPermission != nil {
			_spec.Node.AddColumnOnce(permissionresource.FieldPermissionID)
		}
		if prq.withResource != nil {
			_spec.Node.AddColumnOnce(permissionresource.FieldResourceID)
		}
	}
	if ps := prq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prq *PermissionResourceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prq.driver.Dialect())
	t1 := builder.Table(permissionresource.Table)
	columns := prq.ctx.Fields
	if len(columns) == 0 {
		columns = permissionresource.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prq.sql != nil {
		selector = prq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if prq.ctx.Unique != nil && *prq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range prq.modifiers {
		m(selector)
	}
	for _, p := range prq.predicates {
		p(selector)
	}
	for _, p := range prq.order {
		p(selector)
	}
	if offset := prq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (prq *PermissionResourceQuery) ForUpdate(opts ...sql.LockOption) *PermissionResourceQuery {
	if prq.driver.Dialect() == dialect.Postgres {
		prq.Unique(false)
	}
	prq.modifiers = append(prq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return prq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (prq *PermissionResourceQuery) ForShare(opts ...sql.LockOption) *PermissionResourceQuery {
	if prq.driver.Dialect() == dialect.Postgres {
		prq.Unique(false)
	}
	prq.modifiers = append(prq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return prq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (prq *PermissionResourceQuery) Modify(modifiers ...func(s *sql.Selector)) *PermissionResourceSelect {
	prq.modifiers = append(prq.modifiers, modifiers...)
	return prq.Select()
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
// Example:
//
//	var v []struct {
//	  PermissionID string `json:"permission_id,omitempty"`
//	  ResourceID string `json:"resource_id,omitempty"`
//	}
//
//	client.PermissionResource.Query().
//	  Omit(
//	  permissionresource.FieldPermissionID,
//	  permissionresource.FieldResourceID,
//	  ).
//	  Scan(ctx, &v)
func (prq *PermissionResourceQuery) Omit(fields ...string) *PermissionResourceSelect {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	for _, col := range permissionresource.Columns {
		if _, ok := omits[col]; !ok {
			prq.ctx.Fields = append(prq.ctx.Fields, col)
		}
	}

	sbuild := &PermissionResourceSelect{PermissionResourceQuery: prq}
	sbuild.label = permissionresource.Label
	sbuild.flds, sbuild.scan = &prq.ctx.Fields, sbuild.Scan
	return sbuild
}

// PermissionResourceGroupBy is the group-by builder for PermissionResource entities.
type PermissionResourceGroupBy struct {
	selector
	build *PermissionResourceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prgb *PermissionResourceGroupBy) Aggregate(fns ...AggregateFunc) *PermissionResourceGroupBy {
	prgb.fns = append(prgb.fns, fns...)
	return prgb
}

// Scan applies the selector query and scans the result into the given value.
func (prgb *PermissionResourceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prgb.build.ctx, ent.OpQueryGroupBy)
	if err := prgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PermissionResourceQuery, *PermissionResourceGroupBy](ctx, prgb.build, prgb, prgb.build.inters, v)
}

func (prgb *PermissionResourceGroupBy) sqlScan(ctx context.Context, root *PermissionResourceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(prgb.fns))
	for _, fn := range prgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*prgb.flds)+len(prgb.fns))
		for _, f := range *prgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*prgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PermissionResourceSelect is the builder for selecting fields of PermissionResource entities.
type PermissionResourceSelect struct {
	*PermissionResourceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (prs *PermissionResourceSelect) Aggregate(fns ...AggregateFunc) *PermissionResourceSelect {
	prs.fns = append(prs.fns, fns...)
	return prs
}

// Scan applies the selector query and scans the result into the given value.
func (prs *PermissionResourceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prs.ctx, ent.OpQuerySelect)
	if err := prs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PermissionResourceQuery, *PermissionResourceSelect](ctx, prs.PermissionResourceQuery, prs, prs.inters, v)
}

func (prs *PermissionResourceSelect) sqlScan(ctx context.Context, root *PermissionResourceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(prs.fns))
	for _, fn := range prs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*prs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (prs *PermissionResourceSelect) Modify(modifiers ...func(s *sql.Selector)) *PermissionResourceSelect {
	prs.modifiers = append(prs.modifiers, modifiers...)
	return prs
}
