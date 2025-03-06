// Code generated by ent, DO NOT EDIT.

package positionpermission

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the positionpermission type in the database.
	Label = "position_permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPositionID holds the string denoting the position_id field in the database.
	FieldPositionID = "position_id"
	// FieldPermissionID holds the string denoting the permission_id field in the database.
	FieldPermissionID = "permission_id"
	// EdgePosition holds the string denoting the position edge name in mutations.
	EdgePosition = "position"
	// EdgePermission holds the string denoting the permission edge name in mutations.
	EdgePermission = "permission"
	// Table holds the table name of the positionpermission in the database.
	Table = "sys_position_permissions"
	// PositionTable is the table that holds the position relation/edge.
	PositionTable = "sys_position_permissions"
	// PositionInverseTable is the table name for the Position entity.
	// It exists in this package in order to avoid circular dependency with the "position" package.
	PositionInverseTable = "sys_positions"
	// PositionColumn is the table column denoting the position relation/edge.
	PositionColumn = "position_id"
	// PermissionTable is the table that holds the permission relation/edge.
	PermissionTable = "sys_position_permissions"
	// PermissionInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionInverseTable = "sys_permissions"
	// PermissionColumn is the table column denoting the permission relation/edge.
	PermissionColumn = "permission_id"
)

// Columns holds all SQL columns for positionpermission fields.
var Columns = []string{
	FieldID,
	FieldPositionID,
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
	// PositionIDValidator is a validator for the "position_id" field. It is called by the builders before save.
	PositionIDValidator func(int64) error
	// PermissionIDValidator is a validator for the "permission_id" field. It is called by the builders before save.
	PermissionIDValidator func(int64) error
)

// OrderOption defines the ordering options for the PositionPermission queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPositionID orders the results by the position_id field.
func ByPositionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPositionID, opts...).ToFunc()
}

// ByPermissionID orders the results by the permission_id field.
func ByPermissionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPermissionID, opts...).ToFunc()
}

// ByPositionField orders the results by position field.
func ByPositionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPositionStep(), sql.OrderByField(field, opts...))
	}
}

// ByPermissionField orders the results by permission field.
func ByPermissionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionStep(), sql.OrderByField(field, opts...))
	}
}
func newPositionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PositionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PositionTable, PositionColumn),
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
	return omitColumns(Columns, fields, true)
}

// OmitCustomColumns returns all fields that are not in the list of fields.
func OmitCustomColumns(src []string, fields ...string) []string {
	if len(src) == 0 {
		src = Columns
	}
	// Default removal FieldID
	return omitColumns(src, fields, true)
}

// OmitColumnsWithID returns all fields that are not in the list of fields.
func OmitColumnsWithID(fields ...string) []string {
	// Not remove FieldID
	return omitColumns(Columns, fields, false)
}

// OmitCustomColumns returns all fields that are not in the list of fields.
func OmitCustomColumnsWithID(src []string, fields ...string) []string {
	if len(src) == 0 {
		src = Columns
	}
	// Not remove FieldID
	return omitColumns(src, fields, false)
}

func omitColumns(src []string, fields []string, omitID bool) []string {
	// Default removal FieldID
	filteredFields := make([]string, 0, len(src))
	for _, field := range src {
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
