// Code generated by ent, DO NOT EDIT.

package position

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the position type in the database.
	Label = "position"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldDepartmentID holds the string denoting the department_id field in the database.
	FieldDepartmentID = "department_id"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"
	// EdgeUserPositions holds the string denoting the user_positions edge name in mutations.
	EdgeUserPositions = "user_positions"
	// Table holds the table name of the position in the database.
	Table = "sys_positions"
	// DepartmentTable is the table that holds the department relation/edge.
	DepartmentTable = "sys_positions"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "sys_departments"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
	// UserPositionsTable is the table that holds the user_positions relation/edge.
	UserPositionsTable = "sys_user_positions"
	// UserPositionsInverseTable is the table name for the UserPosition entity.
	// It exists in this package in order to avoid circular dependency with the "userposition" package.
	UserPositionsInverseTable = "sys_user_positions"
	// UserPositionsColumn is the table column denoting the user_positions relation/edge.
	UserPositionsColumn = "position_user_positions"
)

// Columns holds all SQL columns for position fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldDescription,
	FieldDepartmentID,
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultDepartmentID holds the default value on creation for the "department_id" field.
	DefaultDepartmentID func() int64
	// DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	DepartmentIDValidator func(int64) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)

// OrderOption defines the ordering options for the Position queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByDepartmentID orders the results by the department_id field.
func ByDepartmentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartmentID, opts...).ToFunc()
}

// ByDepartmentField orders the results by department field.
func ByDepartmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDepartmentStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserPositionsCount orders the results by user_positions count.
func ByUserPositionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserPositionsStep(), opts...)
	}
}

// ByUserPositions orders the results by user_positions terms.
func ByUserPositions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserPositionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDepartmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DepartmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
	)
}
func newUserPositionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserPositionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserPositionsTable, UserPositionsColumn),
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
