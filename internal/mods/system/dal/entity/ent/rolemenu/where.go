// Code generated by ent, DO NOT EDIT.

package rolemenu

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLTE(FieldID, id))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldRoleID, v))
}

// MenuID applies equality check predicate on the "menu_id" field. It's identical to MenuIDEQ.
func MenuID(v int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldMenuID, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNotIn(FieldRoleID, vs...))
}

// MenuIDEQ applies the EQ predicate on the "menu_id" field.
func MenuIDEQ(v int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldMenuID, v))
}

// MenuIDNEQ applies the NEQ predicate on the "menu_id" field.
func MenuIDNEQ(v int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNEQ(FieldMenuID, v))
}

// MenuIDIn applies the In predicate on the "menu_id" field.
func MenuIDIn(vs ...int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldIn(FieldMenuID, vs...))
}

// MenuIDNotIn applies the NotIn predicate on the "menu_id" field.
func MenuIDNotIn(vs ...int64) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNotIn(FieldMenuID, vs...))
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.RoleMenu {
	return predicate.RoleMenu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.RoleMenu {
	return predicate.RoleMenu(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenu applies the HasEdge predicate on the "menu" edge.
func HasMenu() predicate.RoleMenu {
	return predicate.RoleMenu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MenuTable, MenuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenuWith applies the HasEdge predicate on the "menu" edge with a given conditions (other predicates).
func HasMenuWith(preds ...predicate.Menu) predicate.RoleMenu {
	return predicate.RoleMenu(func(s *sql.Selector) {
		step := newMenuStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RoleMenu) predicate.RoleMenu {
	return predicate.RoleMenu(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RoleMenu) predicate.RoleMenu {
	return predicate.RoleMenu(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RoleMenu) predicate.RoleMenu {
	return predicate.RoleMenu(sql.NotPredicates(p))
}
