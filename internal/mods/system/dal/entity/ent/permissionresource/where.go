// Code generated by ent, DO NOT EDIT.

package permissionresource

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldLTE(FieldID, id))
}

// PermissionID applies equality check predicate on the "permission_id" field. It's identical to PermissionIDEQ.
func PermissionID(v int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldEQ(FieldPermissionID, v))
}

// ResourceID applies equality check predicate on the "resource_id" field. It's identical to ResourceIDEQ.
func ResourceID(v int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldEQ(FieldResourceID, v))
}

// Actions applies equality check predicate on the "actions" field. It's identical to ActionsEQ.
func Actions(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldEQ(FieldActions, v))
}

// PermissionIDEQ applies the EQ predicate on the "permission_id" field.
func PermissionIDEQ(v int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldEQ(FieldPermissionID, v))
}

// PermissionIDNEQ applies the NEQ predicate on the "permission_id" field.
func PermissionIDNEQ(v int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldNEQ(FieldPermissionID, v))
}

// PermissionIDIn applies the In predicate on the "permission_id" field.
func PermissionIDIn(vs ...int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldIn(FieldPermissionID, vs...))
}

// PermissionIDNotIn applies the NotIn predicate on the "permission_id" field.
func PermissionIDNotIn(vs ...int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldNotIn(FieldPermissionID, vs...))
}

// ResourceIDEQ applies the EQ predicate on the "resource_id" field.
func ResourceIDEQ(v int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldEQ(FieldResourceID, v))
}

// ResourceIDNEQ applies the NEQ predicate on the "resource_id" field.
func ResourceIDNEQ(v int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldNEQ(FieldResourceID, v))
}

// ResourceIDIn applies the In predicate on the "resource_id" field.
func ResourceIDIn(vs ...int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldIn(FieldResourceID, vs...))
}

// ResourceIDNotIn applies the NotIn predicate on the "resource_id" field.
func ResourceIDNotIn(vs ...int64) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldNotIn(FieldResourceID, vs...))
}

// ActionsEQ applies the EQ predicate on the "actions" field.
func ActionsEQ(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldEQ(FieldActions, v))
}

// ActionsNEQ applies the NEQ predicate on the "actions" field.
func ActionsNEQ(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldNEQ(FieldActions, v))
}

// ActionsIn applies the In predicate on the "actions" field.
func ActionsIn(vs ...string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldIn(FieldActions, vs...))
}

// ActionsNotIn applies the NotIn predicate on the "actions" field.
func ActionsNotIn(vs ...string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldNotIn(FieldActions, vs...))
}

// ActionsGT applies the GT predicate on the "actions" field.
func ActionsGT(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldGT(FieldActions, v))
}

// ActionsGTE applies the GTE predicate on the "actions" field.
func ActionsGTE(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldGTE(FieldActions, v))
}

// ActionsLT applies the LT predicate on the "actions" field.
func ActionsLT(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldLT(FieldActions, v))
}

// ActionsLTE applies the LTE predicate on the "actions" field.
func ActionsLTE(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldLTE(FieldActions, v))
}

// ActionsContains applies the Contains predicate on the "actions" field.
func ActionsContains(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldContains(FieldActions, v))
}

// ActionsHasPrefix applies the HasPrefix predicate on the "actions" field.
func ActionsHasPrefix(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldHasPrefix(FieldActions, v))
}

// ActionsHasSuffix applies the HasSuffix predicate on the "actions" field.
func ActionsHasSuffix(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldHasSuffix(FieldActions, v))
}

// ActionsEqualFold applies the EqualFold predicate on the "actions" field.
func ActionsEqualFold(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldEqualFold(FieldActions, v))
}

// ActionsContainsFold applies the ContainsFold predicate on the "actions" field.
func ActionsContainsFold(v string) predicate.PermissionResource {
	return predicate.PermissionResource(sql.FieldContainsFold(FieldActions, v))
}

// HasPermission applies the HasEdge predicate on the "permission" edge.
func HasPermission() predicate.PermissionResource {
	return predicate.PermissionResource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PermissionTable, PermissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionWith applies the HasEdge predicate on the "permission" edge with a given conditions (other predicates).
func HasPermissionWith(preds ...predicate.Permission) predicate.PermissionResource {
	return predicate.PermissionResource(func(s *sql.Selector) {
		step := newPermissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResource applies the HasEdge predicate on the "resource" edge.
func HasResource() predicate.PermissionResource {
	return predicate.PermissionResource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ResourceTable, ResourceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResourceWith applies the HasEdge predicate on the "resource" edge with a given conditions (other predicates).
func HasResourceWith(preds ...predicate.Resource) predicate.PermissionResource {
	return predicate.PermissionResource(func(s *sql.Selector) {
		step := newResourceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PermissionResource) predicate.PermissionResource {
	return predicate.PermissionResource(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PermissionResource) predicate.PermissionResource {
	return predicate.PermissionResource(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PermissionResource) predicate.PermissionResource {
	return predicate.PermissionResource(sql.NotPredicates(p))
}
