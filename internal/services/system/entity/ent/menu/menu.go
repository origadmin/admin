// Code generated by ent, DO NOT EDIT.

package menu

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the menu type in the database.
	Label = "menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSequence holds the string denoting the sequence field in the database.
	FieldSequence = "sequence"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldProperties holds the string denoting the properties field in the database.
	FieldProperties = "properties"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldParentPath holds the string denoting the parent_path field in the database.
	FieldParentPath = "parent_path"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the menu in the database.
	Table = "sys_menus"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "sys_menus"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "sys_menus"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
)

// Columns holds all SQL columns for menu fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldName,
	FieldDescription,
	FieldSequence,
	FieldType,
	FieldPath,
	FieldProperties,
	FieldStatus,
	FieldParentID,
	FieldParentPath,
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
	// DefaultCode holds the default value on creation for the "code" field.
	DefaultCode string
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultPath holds the default value on creation for the "path" field.
	DefaultPath string
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(string) error
	// DefaultParentPath holds the default value on creation for the "parent_path" field.
	DefaultParentPath string
	// ParentPathValidator is a validator for the "parent_path" field. It is called by the builders before save.
	ParentPathValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusDisabled is the default value of the Status enum.
const DefaultStatus = StatusDisabled

// Status values.
const (
	StatusDisabled Status = "disabled"
	StatusEnabled  Status = "enabled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDisabled, StatusEnabled:
		return nil
	default:
		return fmt.Errorf("menu: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Menu queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// BySequence orders the results by the sequence field.
func BySequence(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSequence, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByProperties orders the results by the properties field.
func ByProperties(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProperties, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByParentPath orders the results by the parent_path field.
func ByParentPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentPath, opts...).ToFunc()
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}