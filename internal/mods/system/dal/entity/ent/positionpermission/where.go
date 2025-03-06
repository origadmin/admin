// Code generated by ent, DO NOT EDIT.

package positionpermission

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldLTE(FieldID, id))
}

// PositionID applies equality check predicate on the "position_id" field. It's identical to PositionIDEQ.
func PositionID(v int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldEQ(FieldPositionID, v))
}

// PermissionID applies equality check predicate on the "permission_id" field. It's identical to PermissionIDEQ.
func PermissionID(v int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldEQ(FieldPermissionID, v))
}

// PositionIDEQ applies the EQ predicate on the "position_id" field.
func PositionIDEQ(v int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldEQ(FieldPositionID, v))
}

// PositionIDNEQ applies the NEQ predicate on the "position_id" field.
func PositionIDNEQ(v int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldNEQ(FieldPositionID, v))
}

// PositionIDIn applies the In predicate on the "position_id" field.
func PositionIDIn(vs ...int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldIn(FieldPositionID, vs...))
}

// PositionIDNotIn applies the NotIn predicate on the "position_id" field.
func PositionIDNotIn(vs ...int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldNotIn(FieldPositionID, vs...))
}

// PermissionIDEQ applies the EQ predicate on the "permission_id" field.
func PermissionIDEQ(v int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldEQ(FieldPermissionID, v))
}

// PermissionIDNEQ applies the NEQ predicate on the "permission_id" field.
func PermissionIDNEQ(v int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldNEQ(FieldPermissionID, v))
}

// PermissionIDIn applies the In predicate on the "permission_id" field.
func PermissionIDIn(vs ...int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldIn(FieldPermissionID, vs...))
}

// PermissionIDNotIn applies the NotIn predicate on the "permission_id" field.
func PermissionIDNotIn(vs ...int64) predicate.PositionPermission {
	return predicate.PositionPermission(sql.FieldNotIn(FieldPermissionID, vs...))
}

// HasPosition applies the HasEdge predicate on the "position" edge.
func HasPosition() predicate.PositionPermission {
	return predicate.PositionPermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PositionTable, PositionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPositionWith applies the HasEdge predicate on the "position" edge with a given conditions (other predicates).
func HasPositionWith(preds ...predicate.Position) predicate.PositionPermission {
	return predicate.PositionPermission(func(s *sql.Selector) {
		step := newPositionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPermission applies the HasEdge predicate on the "permission" edge.
func HasPermission() predicate.PositionPermission {
	return predicate.PositionPermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PermissionTable, PermissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionWith applies the HasEdge predicate on the "permission" edge with a given conditions (other predicates).
func HasPermissionWith(preds ...predicate.Permission) predicate.PositionPermission {
	return predicate.PositionPermission(func(s *sql.Selector) {
		step := newPermissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PositionPermission) predicate.PositionPermission {
	return predicate.PositionPermission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PositionPermission) predicate.PositionPermission {
	return predicate.PositionPermission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PositionPermission) predicate.PositionPermission {
	return predicate.PositionPermission(sql.NotPredicates(p))
}
