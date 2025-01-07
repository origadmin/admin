// Code generated by ent, DO NOT EDIT.

package department

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdateTime, v))
}

// Keyword applies equality check predicate on the "keyword" field. It's identical to KeywordEQ.
func Keyword(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldKeyword, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldName, v))
}

// TreePath applies equality check predicate on the "tree_path" field. It's identical to TreePathEQ.
func TreePath(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldTreePath, v))
}

// Sequence applies equality check predicate on the "sequence" field. It's identical to SequenceEQ.
func Sequence(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldSequence, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldStatus, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldLevel, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldDescription, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldParentID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldUpdateTime, v))
}

// KeywordEQ applies the EQ predicate on the "keyword" field.
func KeywordEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldKeyword, v))
}

// KeywordNEQ applies the NEQ predicate on the "keyword" field.
func KeywordNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldKeyword, v))
}

// KeywordIn applies the In predicate on the "keyword" field.
func KeywordIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldKeyword, vs...))
}

// KeywordNotIn applies the NotIn predicate on the "keyword" field.
func KeywordNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldKeyword, vs...))
}

// KeywordGT applies the GT predicate on the "keyword" field.
func KeywordGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldKeyword, v))
}

// KeywordGTE applies the GTE predicate on the "keyword" field.
func KeywordGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldKeyword, v))
}

// KeywordLT applies the LT predicate on the "keyword" field.
func KeywordLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldKeyword, v))
}

// KeywordLTE applies the LTE predicate on the "keyword" field.
func KeywordLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldKeyword, v))
}

// KeywordContains applies the Contains predicate on the "keyword" field.
func KeywordContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldKeyword, v))
}

// KeywordHasPrefix applies the HasPrefix predicate on the "keyword" field.
func KeywordHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldKeyword, v))
}

// KeywordHasSuffix applies the HasSuffix predicate on the "keyword" field.
func KeywordHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldKeyword, v))
}

// KeywordEqualFold applies the EqualFold predicate on the "keyword" field.
func KeywordEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldKeyword, v))
}

// KeywordContainsFold applies the ContainsFold predicate on the "keyword" field.
func KeywordContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldKeyword, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldName, v))
}

// TreePathEQ applies the EQ predicate on the "tree_path" field.
func TreePathEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldTreePath, v))
}

// TreePathNEQ applies the NEQ predicate on the "tree_path" field.
func TreePathNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldTreePath, v))
}

// TreePathIn applies the In predicate on the "tree_path" field.
func TreePathIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldTreePath, vs...))
}

// TreePathNotIn applies the NotIn predicate on the "tree_path" field.
func TreePathNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldTreePath, vs...))
}

// TreePathGT applies the GT predicate on the "tree_path" field.
func TreePathGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldTreePath, v))
}

// TreePathGTE applies the GTE predicate on the "tree_path" field.
func TreePathGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldTreePath, v))
}

// TreePathLT applies the LT predicate on the "tree_path" field.
func TreePathLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldTreePath, v))
}

// TreePathLTE applies the LTE predicate on the "tree_path" field.
func TreePathLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldTreePath, v))
}

// TreePathContains applies the Contains predicate on the "tree_path" field.
func TreePathContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldTreePath, v))
}

// TreePathHasPrefix applies the HasPrefix predicate on the "tree_path" field.
func TreePathHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldTreePath, v))
}

// TreePathHasSuffix applies the HasSuffix predicate on the "tree_path" field.
func TreePathHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldTreePath, v))
}

// TreePathEqualFold applies the EqualFold predicate on the "tree_path" field.
func TreePathEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldTreePath, v))
}

// TreePathContainsFold applies the ContainsFold predicate on the "tree_path" field.
func TreePathContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldTreePath, v))
}

// SequenceEQ applies the EQ predicate on the "sequence" field.
func SequenceEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldSequence, v))
}

// SequenceNEQ applies the NEQ predicate on the "sequence" field.
func SequenceNEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldSequence, v))
}

// SequenceIn applies the In predicate on the "sequence" field.
func SequenceIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldSequence, vs...))
}

// SequenceNotIn applies the NotIn predicate on the "sequence" field.
func SequenceNotIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldSequence, vs...))
}

// SequenceGT applies the GT predicate on the "sequence" field.
func SequenceGT(v int) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldSequence, v))
}

// SequenceGTE applies the GTE predicate on the "sequence" field.
func SequenceGTE(v int) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldSequence, v))
}

// SequenceLT applies the LT predicate on the "sequence" field.
func SequenceLT(v int) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldSequence, v))
}

// SequenceLTE applies the LTE predicate on the "sequence" field.
func SequenceLTE(v int) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldSequence, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldStatus, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldLevel, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldDescription, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int64) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int64) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int64) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldParentID))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPositions applies the HasEdge predicate on the "positions" edge.
func HasPositions() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PositionsTable, PositionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPositionsWith applies the HasEdge predicate on the "positions" edge with a given conditions (other predicates).
func HasPositionsWith(preds ...predicate.Position) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newPositionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserDepartments applies the HasEdge predicate on the "user_departments" edge.
func HasUserDepartments() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserDepartmentsTable, UserDepartmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserDepartmentsWith applies the HasEdge predicate on the "user_departments" edge with a given conditions (other predicates).
func HasUserDepartmentsWith(preds ...predicate.UserDepartment) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := newUserDepartmentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Department) predicate.Department {
	return predicate.Department(sql.NotPredicates(p))
}
