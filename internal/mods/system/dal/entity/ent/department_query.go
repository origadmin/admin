// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userdepartment"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DepartmentQuery is the builder for querying Department entities.
type DepartmentQuery struct {
	config
	ctx                 *QueryContext
	order               []department.OrderOption
	inters              []Interceptor
	predicates          []predicate.Department
	withUsers           *UserQuery
	withPositions       *PositionQuery
	withParent          *DepartmentQuery
	withChildren        *DepartmentQuery
	withUserDepartments *UserDepartmentQuery
	modifiers           []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DepartmentQuery builder.
func (dq *DepartmentQuery) Where(ps ...predicate.Department) *DepartmentQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DepartmentQuery) Limit(limit int) *DepartmentQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DepartmentQuery) Offset(offset int) *DepartmentQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DepartmentQuery) Unique(unique bool) *DepartmentQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DepartmentQuery) Order(o ...department.OrderOption) *DepartmentQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryUsers chains the current query on the "users" edge.
func (dq *DepartmentQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, department.UsersTable, department.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPositions chains the current query on the "positions" edge.
func (dq *DepartmentQuery) QueryPositions() *PositionQuery {
	query := (&PositionClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.PositionsTable, department.PositionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (dq *DepartmentQuery) QueryParent() *DepartmentQuery {
	query := (&DepartmentClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, department.ParentTable, department.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (dq *DepartmentQuery) QueryChildren() *DepartmentQuery {
	query := (&DepartmentClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.ChildrenTable, department.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserDepartments chains the current query on the "user_departments" edge.
func (dq *DepartmentQuery) QueryUserDepartments() *UserDepartmentQuery {
	query := (&UserDepartmentClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(userdepartment.Table, userdepartment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, department.UserDepartmentsTable, department.UserDepartmentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Department entity from the query.
// Returns a *NotFoundError when no Department was found.
func (dq *DepartmentQuery) First(ctx context.Context) (*Department, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{department.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DepartmentQuery) FirstX(ctx context.Context) *Department {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Department ID from the query.
// Returns a *NotFoundError when no Department ID was found.
func (dq *DepartmentQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{department.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DepartmentQuery) FirstIDX(ctx context.Context) int64 {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Department entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Department entity is found.
// Returns a *NotFoundError when no Department entities are found.
func (dq *DepartmentQuery) Only(ctx context.Context) (*Department, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{department.Label}
	default:
		return nil, &NotSingularError{department.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DepartmentQuery) OnlyX(ctx context.Context) *Department {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Department ID in the query.
// Returns a *NotSingularError when more than one Department ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DepartmentQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = &NotSingularError{department.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DepartmentQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Departments.
func (dq *DepartmentQuery) All(ctx context.Context) ([]*Department, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryAll)
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Department, *DepartmentQuery]()
	return withInterceptors[[]*Department](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DepartmentQuery) AllX(ctx context.Context) []*Department {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Department IDs.
func (dq *DepartmentQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryIDs)
	if err = dq.Select(department.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DepartmentQuery) IDsX(ctx context.Context) []int64 {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DepartmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryCount)
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DepartmentQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DepartmentQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DepartmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryExist)
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DepartmentQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DepartmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DepartmentQuery) Clone() *DepartmentQuery {
	if dq == nil {
		return nil
	}
	return &DepartmentQuery{
		config:              dq.config,
		ctx:                 dq.ctx.Clone(),
		order:               append([]department.OrderOption{}, dq.order...),
		inters:              append([]Interceptor{}, dq.inters...),
		predicates:          append([]predicate.Department{}, dq.predicates...),
		withUsers:           dq.withUsers.Clone(),
		withPositions:       dq.withPositions.Clone(),
		withParent:          dq.withParent.Clone(),
		withChildren:        dq.withChildren.Clone(),
		withUserDepartments: dq.withUserDepartments.Clone(),
		// clone intermediate query.
		sql:       dq.sql.Clone(),
		path:      dq.path,
		modifiers: append([]func(*sql.Selector){}, dq.modifiers...),
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithUsers(opts ...func(*UserQuery)) *DepartmentQuery {
	query := (&UserClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withUsers = query
	return dq
}

// WithPositions tells the query-builder to eager-load the nodes that are connected to
// the "positions" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithPositions(opts ...func(*PositionQuery)) *DepartmentQuery {
	query := (&PositionClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withPositions = query
	return dq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithParent(opts ...func(*DepartmentQuery)) *DepartmentQuery {
	query := (&DepartmentClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withParent = query
	return dq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithChildren(opts ...func(*DepartmentQuery)) *DepartmentQuery {
	query := (&DepartmentClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withChildren = query
	return dq
}

// WithUserDepartments tells the query-builder to eager-load the nodes that are connected to
// the "user_departments" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithUserDepartments(opts ...func(*UserDepartmentQuery)) *DepartmentQuery {
	query := (&UserDepartmentClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withUserDepartments = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Department.Query().
//		GroupBy(department.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DepartmentQuery) GroupBy(field string, fields ...string) *DepartmentGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DepartmentGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = department.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Department.Query().
//		Select(department.FieldCreateTime).
//		Scan(ctx, &v)
func (dq *DepartmentQuery) Select(fields ...string) *DepartmentSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DepartmentSelect{DepartmentQuery: dq}
	sbuild.label = department.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DepartmentSelect configured with the given aggregations.
func (dq *DepartmentQuery) Aggregate(fns ...AggregateFunc) *DepartmentSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DepartmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !department.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DepartmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Department, error) {
	var (
		nodes       = []*Department{}
		_spec       = dq.querySpec()
		loadedTypes = [5]bool{
			dq.withUsers != nil,
			dq.withPositions != nil,
			dq.withParent != nil,
			dq.withChildren != nil,
			dq.withUserDepartments != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Department).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Department{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withUsers; query != nil {
		if err := dq.loadUsers(ctx, query, nodes,
			func(n *Department) { n.Edges.Users = []*User{} },
			func(n *Department, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withPositions; query != nil {
		if err := dq.loadPositions(ctx, query, nodes,
			func(n *Department) { n.Edges.Positions = []*Position{} },
			func(n *Department, e *Position) { n.Edges.Positions = append(n.Edges.Positions, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withParent; query != nil {
		if err := dq.loadParent(ctx, query, nodes, nil,
			func(n *Department, e *Department) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withChildren; query != nil {
		if err := dq.loadChildren(ctx, query, nodes,
			func(n *Department) { n.Edges.Children = []*Department{} },
			func(n *Department, e *Department) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withUserDepartments; query != nil {
		if err := dq.loadUserDepartments(ctx, query, nodes,
			func(n *Department) { n.Edges.UserDepartments = []*UserDepartment{} },
			func(n *Department, e *UserDepartment) { n.Edges.UserDepartments = append(n.Edges.UserDepartments, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DepartmentQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Department, init func(*Department), assign func(*Department, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*Department)
	nids := make(map[int64]map[*Department]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(department.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(department.UsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(department.UsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(department.UsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*Department]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dq *DepartmentQuery) loadPositions(ctx context.Context, query *PositionQuery, nodes []*Department, init func(*Department), assign func(*Department, *Position)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Department)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(position.FieldDepartmentID)
	}
	query.Where(predicate.Position(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(department.PositionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DepartmentID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "department_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dq *DepartmentQuery) loadParent(ctx context.Context, query *DepartmentQuery, nodes []*Department, init func(*Department), assign func(*Department, *Department)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Department)
	for i := range nodes {
		fk := nodes[i].ParentID
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
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DepartmentQuery) loadChildren(ctx context.Context, query *DepartmentQuery, nodes []*Department, init func(*Department), assign func(*Department, *Department)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Department)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(department.FieldParentID)
	}
	query.Where(predicate.Department(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(department.ChildrenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ParentID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "parent_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dq *DepartmentQuery) loadUserDepartments(ctx context.Context, query *UserDepartmentQuery, nodes []*Department, init func(*Department), assign func(*Department, *UserDepartment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Department)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(userdepartment.FieldDepartmentID)
	}
	query.Where(predicate.UserDepartment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(department.UserDepartmentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DepartmentID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "department_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DepartmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DepartmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(department.Table, department.Columns, sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt64))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, department.FieldID)
		for i := range fields {
			if fields[i] != department.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dq.withParent != nil {
			_spec.Node.AddColumnOnce(department.FieldParentID)
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DepartmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(department.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = department.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dq.modifiers {
		m(selector)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dq *DepartmentQuery) ForUpdate(opts ...sql.LockOption) *DepartmentQuery {
	if dq.driver.Dialect() == dialect.Postgres {
		dq.Unique(false)
	}
	dq.modifiers = append(dq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dq *DepartmentQuery) ForShare(opts ...sql.LockOption) *DepartmentQuery {
	if dq.driver.Dialect() == dialect.Postgres {
		dq.Unique(false)
	}
	dq.modifiers = append(dq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dq *DepartmentQuery) Modify(modifiers ...func(s *sql.Selector)) *DepartmentSelect {
	dq.modifiers = append(dq.modifiers, modifiers...)
	return dq.Select()
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
// Example:
//
//	var v []struct {
//	  CreateTime time.Time `json:"create_time,omitempty"`
//	  UpdateTime time.Time `json:"update_time,omitempty"`
//	  Keyword string `json:"keyword,omitempty"`
//	  Name string `json:"name,omitempty"`
//	  TreePath string `json:"tree_path,omitempty"`
//	  Sequence int `json:"sequence,omitempty"`
//	  Status int8 `json:"status,omitempty"`
//	  Level int `json:"level,omitempty"`
//	  Description string `json:"description,omitempty"`
//	  ParentID int64 `json:"parent_id,omitempty"`
//	}
//
//	client.Department.Query().
//	  Omit(
//	  department.FieldCreateTime,
//	  department.FieldUpdateTime,
//	  department.FieldKeyword,
//	  department.FieldName,
//	  department.FieldTreePath,
//	  department.FieldSequence,
//	  department.FieldStatus,
//	  department.FieldLevel,
//	  department.FieldDescription,
//	  department.FieldParentID,
//	  ).
//	  Scan(ctx, &v)
func (dq *DepartmentQuery) Omit(fields ...string) *DepartmentSelect {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	for _, col := range department.Columns {
		if _, ok := omits[col]; !ok {
			dq.ctx.Fields = append(dq.ctx.Fields, col)
		}
	}

	sbuild := &DepartmentSelect{DepartmentQuery: dq}
	sbuild.label = department.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// DepartmentGroupBy is the group-by builder for Department entities.
type DepartmentGroupBy struct {
	selector
	build *DepartmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DepartmentGroupBy) Aggregate(fns ...AggregateFunc) *DepartmentGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DepartmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, ent.OpQueryGroupBy)
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DepartmentQuery, *DepartmentGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DepartmentGroupBy) sqlScan(ctx context.Context, root *DepartmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DepartmentSelect is the builder for selecting fields of Department entities.
type DepartmentSelect struct {
	*DepartmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DepartmentSelect) Aggregate(fns ...AggregateFunc) *DepartmentSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DepartmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, ent.OpQuerySelect)
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DepartmentQuery, *DepartmentSelect](ctx, ds.DepartmentQuery, ds, ds.inters, v)
}

func (ds *DepartmentSelect) sqlScan(ctx context.Context, root *DepartmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ds *DepartmentSelect) Modify(modifiers ...func(s *sql.Selector)) *DepartmentSelect {
	ds.modifiers = append(ds.modifiers, modifiers...)
	return ds
}
