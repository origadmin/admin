// Code generated by ent, DO NOT EDIT.

package departmentrole

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the departmentrole type in the database.
	Label = "department_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDepartmentID holds the string denoting the department_id field in the database.
	FieldDepartmentID = "department_id"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// Table holds the table name of the departmentrole in the database.
	Table = "sys_department_roles"
	// DepartmentTable is the table that holds the department relation/edge.
	DepartmentTable = "sys_department_roles"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "sys_departments"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
	// RoleTable is the table that holds the role relation/edge.
	RoleTable = "sys_department_roles"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "sys_roles"
	// RoleColumn is the table column denoting the role relation/edge.
	RoleColumn = "role_id"
)

// Columns holds all SQL columns for departmentrole fields.
var Columns = []string{
	FieldID,
	FieldDepartmentID,
	FieldRoleID,
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
	// DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	DepartmentIDValidator func(int64) error
	// RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	RoleIDValidator func(int64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)

// OrderOption defines the ordering options for the DepartmentRole queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDepartmentID orders the results by the department_id field.
func ByDepartmentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartmentID, opts...).ToFunc()
}

// ByRoleID orders the results by the role_id field.
func ByRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleID, opts...).ToFunc()
}

// ByDepartmentField orders the results by department field.
func ByDepartmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDepartmentStep(), sql.OrderByField(field, opts...))
	}
}

// ByRoleField orders the results by role field.
func ByRoleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleStep(), sql.OrderByField(field, opts...))
	}
}
func newDepartmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DepartmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DepartmentTable, DepartmentColumn),
	)
}
func newRoleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
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
