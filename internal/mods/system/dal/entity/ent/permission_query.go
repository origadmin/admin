// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/positionpermission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolepermission"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionQuery is the builder for querying Permission entities.
type PermissionQuery struct {
	config
	ctx                     *QueryContext
	order                   []permission.OrderOption
	inters                  []Interceptor
	predicates              []predicate.Permission
	withRoles               *RoleQuery
	withPositions           *PositionQuery
	withResources           *ResourceQuery
	withRolePermissions     *RolePermissionQuery
	withPositionPermissions *PositionPermissionQuery
	withPermissionResources *PermissionResourceQuery
	modifiers               []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PermissionQuery builder.
func (pq *PermissionQuery) Where(ps ...predicate.Permission) *PermissionQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PermissionQuery) Limit(limit int) *PermissionQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PermissionQuery) Offset(offset int) *PermissionQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PermissionQuery) Unique(unique bool) *PermissionQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PermissionQuery) Order(o ...permission.OrderOption) *PermissionQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryRoles chains the current query on the "roles" edge.
func (pq *PermissionQuery) QueryRoles() *RoleQuery {
	query := (&RoleClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, permission.RolesTable, permission.RolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPositions chains the current query on the "positions" edge.
func (pq *PermissionQuery) QueryPositions() *PositionQuery {
	query := (&PositionClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, selector),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, permission.PositionsTable, permission.PositionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryResources chains the current query on the "resources" edge.
func (pq *PermissionQuery) QueryResources() *ResourceQuery {
	query := (&ResourceClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, selector),
			sqlgraph.To(resource.Table, resource.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, permission.ResourcesTable, permission.ResourcesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRolePermissions chains the current query on the "role_permissions" edge.
func (pq *PermissionQuery) QueryRolePermissions() *RolePermissionQuery {
	query := (&RolePermissionClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, selector),
			sqlgraph.To(rolepermission.Table, rolepermission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, permission.RolePermissionsTable, permission.RolePermissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPositionPermissions chains the current query on the "position_permissions" edge.
func (pq *PermissionQuery) QueryPositionPermissions() *PositionPermissionQuery {
	query := (&PositionPermissionClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, selector),
			sqlgraph.To(positionpermission.Table, positionpermission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, permission.PositionPermissionsTable, permission.PositionPermissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPermissionResources chains the current query on the "permission_resources" edge.
func (pq *PermissionQuery) QueryPermissionResources() *PermissionResourceQuery {
	query := (&PermissionResourceClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, selector),
			sqlgraph.To(permissionresource.Table, permissionresource.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, permission.PermissionResourcesTable, permission.PermissionResourcesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Permission entity from the query.
// Returns a *NotFoundError when no Permission was found.
func (pq *PermissionQuery) First(ctx context.Context) (*Permission, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{permission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PermissionQuery) FirstX(ctx context.Context) *Permission {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Permission ID from the query.
// Returns a *NotFoundError when no Permission ID was found.
func (pq *PermissionQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{permission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PermissionQuery) FirstIDX(ctx context.Context) int64 {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Permission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Permission entity is found.
// Returns a *NotFoundError when no Permission entities are found.
func (pq *PermissionQuery) Only(ctx context.Context) (*Permission, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{permission.Label}
	default:
		return nil, &NotSingularError{permission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PermissionQuery) OnlyX(ctx context.Context) *Permission {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Permission ID in the query.
// Returns a *NotSingularError when more than one Permission ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PermissionQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{permission.Label}
	default:
		err = &NotSingularError{permission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PermissionQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Permissions.
func (pq *PermissionQuery) All(ctx context.Context) ([]*Permission, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Permission, *PermissionQuery]()
	return withInterceptors[[]*Permission](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PermissionQuery) AllX(ctx context.Context) []*Permission {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Permission IDs.
func (pq *PermissionQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(permission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PermissionQuery) IDsX(ctx context.Context) []int64 {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PermissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PermissionQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PermissionQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PermissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PermissionQuery) Clone() *PermissionQuery {
	if pq == nil {
		return nil
	}
	return &PermissionQuery{
		config:                  pq.config,
		ctx:                     pq.ctx.Clone(),
		order:                   append([]permission.OrderOption{}, pq.order...),
		inters:                  append([]Interceptor{}, pq.inters...),
		predicates:              append([]predicate.Permission{}, pq.predicates...),
		withRoles:               pq.withRoles.Clone(),
		withPositions:           pq.withPositions.Clone(),
		withResources:           pq.withResources.Clone(),
		withRolePermissions:     pq.withRolePermissions.Clone(),
		withPositionPermissions: pq.withPositionPermissions.Clone(),
		withPermissionResources: pq.withPermissionResources.Clone(),
		// clone intermediate query.
		sql:       pq.sql.Clone(),
		path:      pq.path,
		modifiers: append([]func(*sql.Selector){}, pq.modifiers...),
	}
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PermissionQuery) WithRoles(opts ...func(*RoleQuery)) *PermissionQuery {
	query := (&RoleClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withRoles = query
	return pq
}

// WithPositions tells the query-builder to eager-load the nodes that are connected to
// the "positions" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PermissionQuery) WithPositions(opts ...func(*PositionQuery)) *PermissionQuery {
	query := (&PositionClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPositions = query
	return pq
}

// WithResources tells the query-builder to eager-load the nodes that are connected to
// the "resources" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PermissionQuery) WithResources(opts ...func(*ResourceQuery)) *PermissionQuery {
	query := (&ResourceClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withResources = query
	return pq
}

// WithRolePermissions tells the query-builder to eager-load the nodes that are connected to
// the "role_permissions" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PermissionQuery) WithRolePermissions(opts ...func(*RolePermissionQuery)) *PermissionQuery {
	query := (&RolePermissionClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withRolePermissions = query
	return pq
}

// WithPositionPermissions tells the query-builder to eager-load the nodes that are connected to
// the "position_permissions" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PermissionQuery) WithPositionPermissions(opts ...func(*PositionPermissionQuery)) *PermissionQuery {
	query := (&PositionPermissionClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPositionPermissions = query
	return pq
}

// WithPermissionResources tells the query-builder to eager-load the nodes that are connected to
// the "permission_resources" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PermissionQuery) WithPermissionResources(opts ...func(*PermissionResourceQuery)) *PermissionQuery {
	query := (&PermissionResourceClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPermissionResources = query
	return pq
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
//	client.Permission.Query().
//		GroupBy(permission.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PermissionQuery) GroupBy(field string, fields ...string) *PermissionGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PermissionGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = permission.Label
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
//	client.Permission.Query().
//		Select(permission.FieldCreateTime).
//		Scan(ctx, &v)
func (pq *PermissionQuery) Select(fields ...string) *PermissionSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PermissionSelect{PermissionQuery: pq}
	sbuild.label = permission.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PermissionSelect configured with the given aggregations.
func (pq *PermissionQuery) Aggregate(fns ...AggregateFunc) *PermissionSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PermissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !permission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PermissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Permission, error) {
	var (
		nodes       = []*Permission{}
		_spec       = pq.querySpec()
		loadedTypes = [6]bool{
			pq.withRoles != nil,
			pq.withPositions != nil,
			pq.withResources != nil,
			pq.withRolePermissions != nil,
			pq.withPositionPermissions != nil,
			pq.withPermissionResources != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Permission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Permission{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withRoles; query != nil {
		if err := pq.loadRoles(ctx, query, nodes,
			func(n *Permission) { n.Edges.Roles = []*Role{} },
			func(n *Permission, e *Role) { n.Edges.Roles = append(n.Edges.Roles, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withPositions; query != nil {
		if err := pq.loadPositions(ctx, query, nodes,
			func(n *Permission) { n.Edges.Positions = []*Position{} },
			func(n *Permission, e *Position) { n.Edges.Positions = append(n.Edges.Positions, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withResources; query != nil {
		if err := pq.loadResources(ctx, query, nodes,
			func(n *Permission) { n.Edges.Resources = []*Resource{} },
			func(n *Permission, e *Resource) { n.Edges.Resources = append(n.Edges.Resources, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withRolePermissions; query != nil {
		if err := pq.loadRolePermissions(ctx, query, nodes,
			func(n *Permission) { n.Edges.RolePermissions = []*RolePermission{} },
			func(n *Permission, e *RolePermission) { n.Edges.RolePermissions = append(n.Edges.RolePermissions, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withPositionPermissions; query != nil {
		if err := pq.loadPositionPermissions(ctx, query, nodes,
			func(n *Permission) { n.Edges.PositionPermissions = []*PositionPermission{} },
			func(n *Permission, e *PositionPermission) {
				n.Edges.PositionPermissions = append(n.Edges.PositionPermissions, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := pq.withPermissionResources; query != nil {
		if err := pq.loadPermissionResources(ctx, query, nodes,
			func(n *Permission) { n.Edges.PermissionResources = []*PermissionResource{} },
			func(n *Permission, e *PermissionResource) {
				n.Edges.PermissionResources = append(n.Edges.PermissionResources, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PermissionQuery) loadRoles(ctx context.Context, query *RoleQuery, nodes []*Permission, init func(*Permission), assign func(*Permission, *Role)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*Permission)
	nids := make(map[int64]map[*Permission]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(permission.RolesTable)
		s.Join(joinT).On(s.C(role.FieldID), joinT.C(permission.RolesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(permission.RolesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(permission.RolesPrimaryKey[1]))
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
					nids[inValue] = map[*Permission]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Role](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "roles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PermissionQuery) loadPositions(ctx context.Context, query *PositionQuery, nodes []*Permission, init func(*Permission), assign func(*Permission, *Position)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*Permission)
	nids := make(map[int64]map[*Permission]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(permission.PositionsTable)
		s.Join(joinT).On(s.C(position.FieldID), joinT.C(permission.PositionsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(permission.PositionsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(permission.PositionsPrimaryKey[1]))
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
					nids[inValue] = map[*Permission]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Position](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "positions" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PermissionQuery) loadResources(ctx context.Context, query *ResourceQuery, nodes []*Permission, init func(*Permission), assign func(*Permission, *Resource)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*Permission)
	nids := make(map[int64]map[*Permission]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(permission.ResourcesTable)
		s.Join(joinT).On(s.C(resource.FieldID), joinT.C(permission.ResourcesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(permission.ResourcesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(permission.ResourcesPrimaryKey[0]))
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
					nids[inValue] = map[*Permission]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Resource](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "resources" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PermissionQuery) loadRolePermissions(ctx context.Context, query *RolePermissionQuery, nodes []*Permission, init func(*Permission), assign func(*Permission, *RolePermission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Permission)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(rolepermission.FieldPermissionID)
	}
	query.Where(predicate.RolePermission(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(permission.RolePermissionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PermissionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "permission_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PermissionQuery) loadPositionPermissions(ctx context.Context, query *PositionPermissionQuery, nodes []*Permission, init func(*Permission), assign func(*Permission, *PositionPermission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Permission)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(positionpermission.FieldPermissionID)
	}
	query.Where(predicate.PositionPermission(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(permission.PositionPermissionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PermissionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "permission_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PermissionQuery) loadPermissionResources(ctx context.Context, query *PermissionResourceQuery, nodes []*Permission, init func(*Permission), assign func(*Permission, *PermissionResource)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Permission)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(permissionresource.FieldPermissionID)
	}
	query.Where(predicate.PermissionResource(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(permission.PermissionResourcesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PermissionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "permission_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt64))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for i := range fields {
			if fields[i] != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(permission.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = permission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pq.modifiers {
		m(selector)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (pq *PermissionQuery) ForUpdate(opts ...sql.LockOption) *PermissionQuery {
	if pq.driver.Dialect() == dialect.Postgres {
		pq.Unique(false)
	}
	pq.modifiers = append(pq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return pq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (pq *PermissionQuery) ForShare(opts ...sql.LockOption) *PermissionQuery {
	if pq.driver.Dialect() == dialect.Postgres {
		pq.Unique(false)
	}
	pq.modifiers = append(pq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return pq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pq *PermissionQuery) Modify(modifiers ...func(s *sql.Selector)) *PermissionSelect {
	pq.modifiers = append(pq.modifiers, modifiers...)
	return pq.Select()
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
// Example:
//
//	var v []struct {
//	  CreateTime time.Time `json:"create_time,omitempty"`
//	  UpdateTime time.Time `json:"update_time,omitempty"`
//	  Name string `json:"name,omitempty"`
//	  Keyword string `json:"keyword,omitempty"`
//	  Description string `json:"description,omitempty"`
//	  DataScope string `json:"data_scope,omitempty"`
//	  DataRules map[string]string `json:"data_rules,omitempty"`
//	  Actions permission.Actions `json:"actions,omitempty"`
//	}
//
//	client.Permission.Query().
//	  Omit(
//	  permission.FieldCreateTime,
//	  permission.FieldUpdateTime,
//	  permission.FieldName,
//	  permission.FieldKeyword,
//	  permission.FieldDescription,
//	  permission.FieldDataScope,
//	  permission.FieldDataRules,
//	  permission.FieldActions,
//	  ).
//	  Scan(ctx, &v)
func (pq *PermissionQuery) Omit(fields ...string) *PermissionSelect {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	for _, col := range permission.Columns {
		if _, ok := omits[col]; !ok {
			pq.ctx.Fields = append(pq.ctx.Fields, col)
		}
	}

	sbuild := &PermissionSelect{PermissionQuery: pq}
	sbuild.label = permission.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// PermissionGroupBy is the group-by builder for Permission entities.
type PermissionGroupBy struct {
	selector
	build *PermissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PermissionGroupBy) Aggregate(fns ...AggregateFunc) *PermissionGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PermissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PermissionQuery, *PermissionGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PermissionGroupBy) sqlScan(ctx context.Context, root *PermissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PermissionSelect is the builder for selecting fields of Permission entities.
type PermissionSelect struct {
	*PermissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PermissionSelect) Aggregate(fns ...AggregateFunc) *PermissionSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PermissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PermissionQuery, *PermissionSelect](ctx, ps.PermissionQuery, ps, ps.inters, v)
}

func (ps *PermissionSelect) sqlScan(ctx context.Context, root *PermissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ps *PermissionSelect) Modify(modifiers ...func(s *sql.Selector)) *PermissionSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
