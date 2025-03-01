// Code generated by ent, DO NOT EDIT.

package position

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldName, v))
}

// Keyword applies equality check predicate on the "keyword" field. It's identical to KeywordEQ.
func Keyword(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldKeyword, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldDescription, v))
}

// DepartmentID applies equality check predicate on the "department_id" field. It's identical to DepartmentIDEQ.
func DepartmentID(v int64) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldDepartmentID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldName, v))
}

// KeywordEQ applies the EQ predicate on the "keyword" field.
func KeywordEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldKeyword, v))
}

// KeywordNEQ applies the NEQ predicate on the "keyword" field.
func KeywordNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldKeyword, v))
}

// KeywordIn applies the In predicate on the "keyword" field.
func KeywordIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldKeyword, vs...))
}

// KeywordNotIn applies the NotIn predicate on the "keyword" field.
func KeywordNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldKeyword, vs...))
}

// KeywordGT applies the GT predicate on the "keyword" field.
func KeywordGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldKeyword, v))
}

// KeywordGTE applies the GTE predicate on the "keyword" field.
func KeywordGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldKeyword, v))
}

// KeywordLT applies the LT predicate on the "keyword" field.
func KeywordLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldKeyword, v))
}

// KeywordLTE applies the LTE predicate on the "keyword" field.
func KeywordLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldKeyword, v))
}

// KeywordContains applies the Contains predicate on the "keyword" field.
func KeywordContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldKeyword, v))
}

// KeywordHasPrefix applies the HasPrefix predicate on the "keyword" field.
func KeywordHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldKeyword, v))
}

// KeywordHasSuffix applies the HasSuffix predicate on the "keyword" field.
func KeywordHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldKeyword, v))
}

// KeywordEqualFold applies the EqualFold predicate on the "keyword" field.
func KeywordEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldKeyword, v))
}

// KeywordContainsFold applies the ContainsFold predicate on the "keyword" field.
func KeywordContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldKeyword, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldDescription, v))
}

// DepartmentIDEQ applies the EQ predicate on the "department_id" field.
func DepartmentIDEQ(v int64) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldDepartmentID, v))
}

// DepartmentIDNEQ applies the NEQ predicate on the "department_id" field.
func DepartmentIDNEQ(v int64) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldDepartmentID, v))
}

// DepartmentIDIn applies the In predicate on the "department_id" field.
func DepartmentIDIn(vs ...int64) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldDepartmentID, vs...))
}

// DepartmentIDNotIn applies the NotIn predicate on the "department_id" field.
func DepartmentIDNotIn(vs ...int64) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldDepartmentID, vs...))
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := newDepartmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPermissions applies the HasEdge predicate on the "permissions" edge.
func HasPermissions() predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PermissionsTable, PermissionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionsWith applies the HasEdge predicate on the "permissions" edge with a given conditions (other predicates).
func HasPermissionsWith(preds ...predicate.Permission) predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := newPermissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserPositions applies the HasEdge predicate on the "user_positions" edge.
func HasUserPositions() predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserPositionsTable, UserPositionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserPositionsWith applies the HasEdge predicate on the "user_positions" edge with a given conditions (other predicates).
func HasUserPositionsWith(preds ...predicate.UserPosition) predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := newUserPositionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPositionPermissions applies the HasEdge predicate on the "position_permissions" edge.
func HasPositionPermissions() predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PositionPermissionsTable, PositionPermissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPositionPermissionsWith applies the HasEdge predicate on the "position_permissions" edge with a given conditions (other predicates).
func HasPositionPermissionsWith(preds ...predicate.PositionPermission) predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := newPositionPermissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Position) predicate.Position {
	return predicate.Position(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Position) predicate.Position {
	return predicate.Position(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Position) predicate.Position {
	return predicate.Position(sql.NotPredicates(p))
}
