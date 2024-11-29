// Code generated by ent, DO NOT EDIT.

package rolemenu

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldContainsFold(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldUpdateTime, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldRoleID, v))
}

// MenuID applies equality check predicate on the "menu_id" field. It's identical to MenuIDEQ.
func MenuID(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldMenuID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLTE(FieldUpdateTime, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLTE(FieldRoleID, v))
}

// RoleIDContains applies the Contains predicate on the "role_id" field.
func RoleIDContains(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldContains(FieldRoleID, v))
}

// RoleIDHasPrefix applies the HasPrefix predicate on the "role_id" field.
func RoleIDHasPrefix(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldHasPrefix(FieldRoleID, v))
}

// RoleIDHasSuffix applies the HasSuffix predicate on the "role_id" field.
func RoleIDHasSuffix(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldHasSuffix(FieldRoleID, v))
}

// RoleIDEqualFold applies the EqualFold predicate on the "role_id" field.
func RoleIDEqualFold(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEqualFold(FieldRoleID, v))
}

// RoleIDContainsFold applies the ContainsFold predicate on the "role_id" field.
func RoleIDContainsFold(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldContainsFold(FieldRoleID, v))
}

// MenuIDEQ applies the EQ predicate on the "menu_id" field.
func MenuIDEQ(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEQ(FieldMenuID, v))
}

// MenuIDNEQ applies the NEQ predicate on the "menu_id" field.
func MenuIDNEQ(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNEQ(FieldMenuID, v))
}

// MenuIDIn applies the In predicate on the "menu_id" field.
func MenuIDIn(vs ...string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldIn(FieldMenuID, vs...))
}

// MenuIDNotIn applies the NotIn predicate on the "menu_id" field.
func MenuIDNotIn(vs ...string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldNotIn(FieldMenuID, vs...))
}

// MenuIDGT applies the GT predicate on the "menu_id" field.
func MenuIDGT(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGT(FieldMenuID, v))
}

// MenuIDGTE applies the GTE predicate on the "menu_id" field.
func MenuIDGTE(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldGTE(FieldMenuID, v))
}

// MenuIDLT applies the LT predicate on the "menu_id" field.
func MenuIDLT(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLT(FieldMenuID, v))
}

// MenuIDLTE applies the LTE predicate on the "menu_id" field.
func MenuIDLTE(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldLTE(FieldMenuID, v))
}

// MenuIDContains applies the Contains predicate on the "menu_id" field.
func MenuIDContains(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldContains(FieldMenuID, v))
}

// MenuIDHasPrefix applies the HasPrefix predicate on the "menu_id" field.
func MenuIDHasPrefix(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldHasPrefix(FieldMenuID, v))
}

// MenuIDHasSuffix applies the HasSuffix predicate on the "menu_id" field.
func MenuIDHasSuffix(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldHasSuffix(FieldMenuID, v))
}

// MenuIDEqualFold applies the EqualFold predicate on the "menu_id" field.
func MenuIDEqualFold(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldEqualFold(FieldMenuID, v))
}

// MenuIDContainsFold applies the ContainsFold predicate on the "menu_id" field.
func MenuIDContainsFold(v string) predicate.RoleMenu {
	return predicate.RoleMenu(sql.FieldContainsFold(FieldMenuID, v))
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
