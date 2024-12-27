// Code generated by ent, DO NOT EDIT.

package permission

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldName, v))
}

// Keyword applies equality check predicate on the "keyword" field. It's identical to KeywordEQ.
func Keyword(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldKeyword, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldDescription, v))
}

// I18nKey applies equality check predicate on the "i18n_key" field. It's identical to I18nKeyEQ.
func I18nKey(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldI18nKey, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int8) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldType, v))
}

// Scope applies equality check predicate on the "scope" field. It's identical to ScopeEQ.
func Scope(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldScope, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldName, v))
}

// KeywordEQ applies the EQ predicate on the "keyword" field.
func KeywordEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldKeyword, v))
}

// KeywordNEQ applies the NEQ predicate on the "keyword" field.
func KeywordNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldKeyword, v))
}

// KeywordIn applies the In predicate on the "keyword" field.
func KeywordIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldKeyword, vs...))
}

// KeywordNotIn applies the NotIn predicate on the "keyword" field.
func KeywordNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldKeyword, vs...))
}

// KeywordGT applies the GT predicate on the "keyword" field.
func KeywordGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldKeyword, v))
}

// KeywordGTE applies the GTE predicate on the "keyword" field.
func KeywordGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldKeyword, v))
}

// KeywordLT applies the LT predicate on the "keyword" field.
func KeywordLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldKeyword, v))
}

// KeywordLTE applies the LTE predicate on the "keyword" field.
func KeywordLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldKeyword, v))
}

// KeywordContains applies the Contains predicate on the "keyword" field.
func KeywordContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldKeyword, v))
}

// KeywordHasPrefix applies the HasPrefix predicate on the "keyword" field.
func KeywordHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldKeyword, v))
}

// KeywordHasSuffix applies the HasSuffix predicate on the "keyword" field.
func KeywordHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldKeyword, v))
}

// KeywordEqualFold applies the EqualFold predicate on the "keyword" field.
func KeywordEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldKeyword, v))
}

// KeywordContainsFold applies the ContainsFold predicate on the "keyword" field.
func KeywordContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldKeyword, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldDescription, v))
}

// I18nKeyEQ applies the EQ predicate on the "i18n_key" field.
func I18nKeyEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldI18nKey, v))
}

// I18nKeyNEQ applies the NEQ predicate on the "i18n_key" field.
func I18nKeyNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldI18nKey, v))
}

// I18nKeyIn applies the In predicate on the "i18n_key" field.
func I18nKeyIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldI18nKey, vs...))
}

// I18nKeyNotIn applies the NotIn predicate on the "i18n_key" field.
func I18nKeyNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldI18nKey, vs...))
}

// I18nKeyGT applies the GT predicate on the "i18n_key" field.
func I18nKeyGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldI18nKey, v))
}

// I18nKeyGTE applies the GTE predicate on the "i18n_key" field.
func I18nKeyGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldI18nKey, v))
}

// I18nKeyLT applies the LT predicate on the "i18n_key" field.
func I18nKeyLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldI18nKey, v))
}

// I18nKeyLTE applies the LTE predicate on the "i18n_key" field.
func I18nKeyLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldI18nKey, v))
}

// I18nKeyContains applies the Contains predicate on the "i18n_key" field.
func I18nKeyContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldI18nKey, v))
}

// I18nKeyHasPrefix applies the HasPrefix predicate on the "i18n_key" field.
func I18nKeyHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldI18nKey, v))
}

// I18nKeyHasSuffix applies the HasSuffix predicate on the "i18n_key" field.
func I18nKeyHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldI18nKey, v))
}

// I18nKeyEqualFold applies the EqualFold predicate on the "i18n_key" field.
func I18nKeyEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldI18nKey, v))
}

// I18nKeyContainsFold applies the ContainsFold predicate on the "i18n_key" field.
func I18nKeyContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldI18nKey, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int8) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int8) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int8) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int8) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int8) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int8) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int8) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int8) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldType, v))
}

// ScopeEQ applies the EQ predicate on the "scope" field.
func ScopeEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldScope, v))
}

// ScopeNEQ applies the NEQ predicate on the "scope" field.
func ScopeNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldScope, v))
}

// ScopeIn applies the In predicate on the "scope" field.
func ScopeIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldScope, vs...))
}

// ScopeNotIn applies the NotIn predicate on the "scope" field.
func ScopeNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldScope, vs...))
}

// ScopeGT applies the GT predicate on the "scope" field.
func ScopeGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldScope, v))
}

// ScopeGTE applies the GTE predicate on the "scope" field.
func ScopeGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldScope, v))
}

// ScopeLT applies the LT predicate on the "scope" field.
func ScopeLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldScope, v))
}

// ScopeLTE applies the LTE predicate on the "scope" field.
func ScopeLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldScope, v))
}

// ScopeContains applies the Contains predicate on the "scope" field.
func ScopeContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldScope, v))
}

// ScopeHasPrefix applies the HasPrefix predicate on the "scope" field.
func ScopeHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldScope, v))
}

// ScopeHasSuffix applies the HasSuffix predicate on the "scope" field.
func ScopeHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldScope, v))
}

// ScopeEqualFold applies the EqualFold predicate on the "scope" field.
func ScopeEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldScope, v))
}

// ScopeContainsFold applies the ContainsFold predicate on the "scope" field.
func ScopeContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldScope, v))
}

// ScopeDeptsIsNil applies the IsNil predicate on the "scope_depts" field.
func ScopeDeptsIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldScopeDepts))
}

// ScopeDeptsNotNil applies the NotNil predicate on the "scope_depts" field.
func ScopeDeptsNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldScopeDepts))
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newRolesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenus applies the HasEdge predicate on the "menus" edge.
func HasMenus() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MenusTable, MenusPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenusWith applies the HasEdge predicate on the "menus" edge with a given conditions (other predicates).
func HasMenusWith(preds ...predicate.Menu) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newMenusStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResources applies the HasEdge predicate on the "resources" edge.
func HasResources() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ResourcesTable, ResourcesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResourcesWith applies the HasEdge predicate on the "resources" edge with a given conditions (other predicates).
func HasResourcesWith(preds ...predicate.Resource) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newResourcesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRolePermissions applies the HasEdge predicate on the "role_permissions" edge.
func HasRolePermissions() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RolePermissionsTable, RolePermissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolePermissionsWith applies the HasEdge predicate on the "role_permissions" edge with a given conditions (other predicates).
func HasRolePermissionsWith(preds ...predicate.RolePermission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newRolePermissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenuPermissions applies the HasEdge predicate on the "menu_permissions" edge.
func HasMenuPermissions() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, MenuPermissionsTable, MenuPermissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenuPermissionsWith applies the HasEdge predicate on the "menu_permissions" edge with a given conditions (other predicates).
func HasMenuPermissionsWith(preds ...predicate.MenuPermission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newMenuPermissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPermissionResources applies the HasEdge predicate on the "permission_resources" edge.
func HasPermissionResources() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PermissionResourcesTable, PermissionResourcesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionResourcesWith applies the HasEdge predicate on the "permission_resources" edge with a given conditions (other predicates).
func HasPermissionResourcesWith(preds ...predicate.PermissionResource) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newPermissionResourcesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.NotPredicates(p))
}
