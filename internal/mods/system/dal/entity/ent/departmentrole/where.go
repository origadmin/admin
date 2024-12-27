// Code generated by ent, DO NOT EDIT.

package departmentrole

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldLTE(FieldID, id))
}

// DepartmentID applies equality check predicate on the "department_id" field. It's identical to DepartmentIDEQ.
func DepartmentID(v int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldEQ(FieldDepartmentID, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldEQ(FieldRoleID, v))
}

// DepartmentIDEQ applies the EQ predicate on the "department_id" field.
func DepartmentIDEQ(v int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldEQ(FieldDepartmentID, v))
}

// DepartmentIDNEQ applies the NEQ predicate on the "department_id" field.
func DepartmentIDNEQ(v int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldNEQ(FieldDepartmentID, v))
}

// DepartmentIDIn applies the In predicate on the "department_id" field.
func DepartmentIDIn(vs ...int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldIn(FieldDepartmentID, vs...))
}

// DepartmentIDNotIn applies the NotIn predicate on the "department_id" field.
func DepartmentIDNotIn(vs ...int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldNotIn(FieldDepartmentID, vs...))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...int64) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.FieldNotIn(FieldRoleID, vs...))
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.DepartmentRole {
	return predicate.DepartmentRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.DepartmentRole {
	return predicate.DepartmentRole(func(s *sql.Selector) {
		step := newDepartmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.DepartmentRole {
	return predicate.DepartmentRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.DepartmentRole {
	return predicate.DepartmentRole(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DepartmentRole) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DepartmentRole) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DepartmentRole) predicate.DepartmentRole {
	return predicate.DepartmentRole(sql.NotPredicates(p))
}
