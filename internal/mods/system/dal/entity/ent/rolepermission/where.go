// Code generated by ent, DO NOT EDIT.

package rolepermission

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldID, id))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldRoleID, v))
}

// PermissionID applies equality check predicate on the "permission_id" field. It's identical to PermissionIDEQ.
func PermissionID(v int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldPermissionID, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldRoleID, vs...))
}

// PermissionIDEQ applies the EQ predicate on the "permission_id" field.
func PermissionIDEQ(v int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldPermissionID, v))
}

// PermissionIDNEQ applies the NEQ predicate on the "permission_id" field.
func PermissionIDNEQ(v int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldPermissionID, v))
}

// PermissionIDIn applies the In predicate on the "permission_id" field.
func PermissionIDIn(vs ...int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldPermissionID, vs...))
}

// PermissionIDNotIn applies the NotIn predicate on the "permission_id" field.
func PermissionIDNotIn(vs ...int) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldPermissionID, vs...))
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPermission applies the HasEdge predicate on the "permission" edge.
func HasPermission() predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PermissionTable, PermissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionWith applies the HasEdge predicate on the "permission" edge with a given conditions (other predicates).
func HasPermissionWith(preds ...predicate.Permission) predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := newPermissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RolePermission) predicate.RolePermission {
	return predicate.RolePermission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RolePermission) predicate.RolePermission {
	return predicate.RolePermission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RolePermission) predicate.RolePermission {
	return predicate.RolePermission(sql.NotPredicates(p))
}
