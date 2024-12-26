// Code generated by ent, DO NOT EDIT.

package userposition

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldUserID, v))
}

// PositionID applies equality check predicate on the "position_id" field. It's identical to PositionIDEQ.
func PositionID(v int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldPositionID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNotIn(FieldUserID, vs...))
}

// PositionIDEQ applies the EQ predicate on the "position_id" field.
func PositionIDEQ(v int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldPositionID, v))
}

// PositionIDNEQ applies the NEQ predicate on the "position_id" field.
func PositionIDNEQ(v int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNEQ(FieldPositionID, v))
}

// PositionIDIn applies the In predicate on the "position_id" field.
func PositionIDIn(vs ...int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldIn(FieldPositionID, vs...))
}

// PositionIDNotIn applies the NotIn predicate on the "position_id" field.
func PositionIDNotIn(vs ...int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNotIn(FieldPositionID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserPosition {
	return predicate.UserPosition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserPosition {
	return predicate.UserPosition(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosition applies the HasEdge predicate on the "position" edge.
func HasPosition() predicate.UserPosition {
	return predicate.UserPosition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PositionTable, PositionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPositionWith applies the HasEdge predicate on the "position" edge with a given conditions (other predicates).
func HasPositionWith(preds ...predicate.Position) predicate.UserPosition {
	return predicate.UserPosition(func(s *sql.Selector) {
		step := newPositionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserPosition) predicate.UserPosition {
	return predicate.UserPosition(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserPosition) predicate.UserPosition {
	return predicate.UserPosition(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserPosition) predicate.UserPosition {
	return predicate.UserPosition(sql.NotPredicates(p))
}
