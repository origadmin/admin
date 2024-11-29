// Code generated by ent, DO NOT EDIT.

package userrole

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.UserRole {
	return predicate.UserRole(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.UserRole {
	return predicate.UserRole(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.UserRole {
	return predicate.UserRole(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.UserRole {
	return predicate.UserRole(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.UserRole {
	return predicate.UserRole(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.UserRole {
	return predicate.UserRole(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.UserRole {
	return predicate.UserRole(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.UserRole {
	return predicate.UserRole(sql.FieldContainsFold(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldUpdateTime, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldUserID, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldRoleID, v))
}

// RoleName applies equality check predicate on the "role_name" field. It's identical to RoleNameEQ.
func RoleName(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldRoleName, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.UserRole {
	return predicate.UserRole(sql.FieldLTE(FieldUpdateTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.UserRole {
	return predicate.UserRole(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.UserRole {
	return predicate.UserRole(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldContainsFold(FieldUserID, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...string) predicate.UserRole {
	return predicate.UserRole(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...string) predicate.UserRole {
	return predicate.UserRole(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldLTE(FieldRoleID, v))
}

// RoleIDContains applies the Contains predicate on the "role_id" field.
func RoleIDContains(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldContains(FieldRoleID, v))
}

// RoleIDHasPrefix applies the HasPrefix predicate on the "role_id" field.
func RoleIDHasPrefix(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldHasPrefix(FieldRoleID, v))
}

// RoleIDHasSuffix applies the HasSuffix predicate on the "role_id" field.
func RoleIDHasSuffix(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldHasSuffix(FieldRoleID, v))
}

// RoleIDEqualFold applies the EqualFold predicate on the "role_id" field.
func RoleIDEqualFold(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEqualFold(FieldRoleID, v))
}

// RoleIDContainsFold applies the ContainsFold predicate on the "role_id" field.
func RoleIDContainsFold(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldContainsFold(FieldRoleID, v))
}

// RoleNameEQ applies the EQ predicate on the "role_name" field.
func RoleNameEQ(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEQ(FieldRoleName, v))
}

// RoleNameNEQ applies the NEQ predicate on the "role_name" field.
func RoleNameNEQ(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldNEQ(FieldRoleName, v))
}

// RoleNameIn applies the In predicate on the "role_name" field.
func RoleNameIn(vs ...string) predicate.UserRole {
	return predicate.UserRole(sql.FieldIn(FieldRoleName, vs...))
}

// RoleNameNotIn applies the NotIn predicate on the "role_name" field.
func RoleNameNotIn(vs ...string) predicate.UserRole {
	return predicate.UserRole(sql.FieldNotIn(FieldRoleName, vs...))
}

// RoleNameGT applies the GT predicate on the "role_name" field.
func RoleNameGT(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldGT(FieldRoleName, v))
}

// RoleNameGTE applies the GTE predicate on the "role_name" field.
func RoleNameGTE(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldGTE(FieldRoleName, v))
}

// RoleNameLT applies the LT predicate on the "role_name" field.
func RoleNameLT(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldLT(FieldRoleName, v))
}

// RoleNameLTE applies the LTE predicate on the "role_name" field.
func RoleNameLTE(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldLTE(FieldRoleName, v))
}

// RoleNameContains applies the Contains predicate on the "role_name" field.
func RoleNameContains(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldContains(FieldRoleName, v))
}

// RoleNameHasPrefix applies the HasPrefix predicate on the "role_name" field.
func RoleNameHasPrefix(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldHasPrefix(FieldRoleName, v))
}

// RoleNameHasSuffix applies the HasSuffix predicate on the "role_name" field.
func RoleNameHasSuffix(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldHasSuffix(FieldRoleName, v))
}

// RoleNameEqualFold applies the EqualFold predicate on the "role_name" field.
func RoleNameEqualFold(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldEqualFold(FieldRoleName, v))
}

// RoleNameContainsFold applies the ContainsFold predicate on the "role_name" field.
func RoleNameContainsFold(v string) predicate.UserRole {
	return predicate.UserRole(sql.FieldContainsFold(FieldRoleName, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.UserRole {
	return predicate.UserRole(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserRole) predicate.UserRole {
	return predicate.UserRole(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserRole) predicate.UserRole {
	return predicate.UserRole(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserRole) predicate.UserRole {
	return predicate.UserRole(sql.NotPredicates(p))
}
