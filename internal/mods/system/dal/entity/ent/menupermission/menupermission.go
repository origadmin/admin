// Code generated by ent, DO NOT EDIT.

package menupermission

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the menupermission type in the database.
	Label = "menu_permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMenuID holds the string denoting the menu_id field in the database.
	FieldMenuID = "menu_id"
	// FieldPermissionID holds the string denoting the permission_id field in the database.
	FieldPermissionID = "permission_id"
	// EdgeMenu holds the string denoting the menu edge name in mutations.
	EdgeMenu = "menu"
	// EdgePermission holds the string denoting the permission edge name in mutations.
	EdgePermission = "permission"
	// Table holds the table name of the menupermission in the database.
	Table = "sys_menu_permissions"
	// MenuTable is the table that holds the menu relation/edge.
	MenuTable = "sys_menu_permissions"
	// MenuInverseTable is the table name for the Menu entity.
	// It exists in this package in order to avoid circular dependency with the "menu" package.
	MenuInverseTable = "sys_menus"
	// MenuColumn is the table column denoting the menu relation/edge.
	MenuColumn = "menu_id"
	// PermissionTable is the table that holds the permission relation/edge.
	PermissionTable = "sys_menu_permissions"
	// PermissionInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionInverseTable = "sys_permissions"
	// PermissionColumn is the table column denoting the permission relation/edge.
	PermissionColumn = "permission_id"
)

// Columns holds all SQL columns for menupermission fields.
var Columns = []string{
	FieldID,
	FieldMenuID,
	FieldPermissionID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// MenuIDValidator is a validator for the "menu_id" field. It is called by the builders before save.
	MenuIDValidator func(string) error
	// PermissionIDValidator is a validator for the "permission_id" field. It is called by the builders before save.
	PermissionIDValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)

// OrderOption defines the ordering options for the MenuPermission queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMenuID orders the results by the menu_id field.
func ByMenuID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMenuID, opts...).ToFunc()
}

// ByPermissionID orders the results by the permission_id field.
func ByPermissionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPermissionID, opts...).ToFunc()
}

// ByMenuField orders the results by menu field.
func ByMenuField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMenuStep(), sql.OrderByField(field, opts...))
	}
}

// ByPermissionField orders the results by permission field.
func ByPermissionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionStep(), sql.OrderByField(field, opts...))
	}
}
func newMenuStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MenuInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, MenuTable, MenuColumn),
	)
}
func newPermissionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PermissionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PermissionTable, PermissionColumn),
	)
}

// SelectColumns returns all selected fields.
func SelectColumns(fields []string) []string {
	// Default removal FieldID
	filteredFields := make([]string, 0, len(fields))
	for _, field := range fields {
		if field != FieldID {
			filteredFields = append(filteredFields, field)
		}
	}
	return filteredFields
}

// OmitColumns returns all fields that are not in the list of fields.
func OmitColumns(fields ...string) []string {
	// Default removal FieldID
	return omitColumns(fields, true)
}

// OmitColumnsWithID returns all fields that are not in the list of fields.
func OmitColumnsWithID(fields ...string) []string {
	// Not remove FieldID
	return omitColumns(fields, false)
}

func omitColumns(fields []string, omitID bool) []string {
	// Default removal FieldID
	allFields := Columns
	filteredFields := make([]string, 0, len(allFields))
	for _, field := range allFields {
		if !(omitID && field == FieldID) && !contains(fields, field) {
			filteredFields = append(filteredFields, field)
		}
	}
	return filteredFields
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
