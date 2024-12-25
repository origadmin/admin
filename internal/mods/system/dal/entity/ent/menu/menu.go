// Code generated by ent, DO NOT EDIT.

package menu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the menu type in the database.
	Label = "menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldKeyword holds the string denoting the keyword field in the database.
	FieldKeyword = "keyword"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldParentPath holds the string denoting the parent_path field in the database.
	FieldParentPath = "parent_path"
	// FieldSequence holds the string denoting the sequence field in the database.
	FieldSequence = "sequence"
	// FieldProperties holds the string denoting the properties field in the database.
	FieldProperties = "properties"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeResources holds the string denoting the resources edge name in mutations.
	EdgeResources = "resources"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeRoleMenus holds the string denoting the role_menus edge name in mutations.
	EdgeRoleMenus = "role_menus"
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
	// ResourcesTable is the table that holds the resources relation/edge.
	ResourcesTable = "sys_resources"
	// ResourcesInverseTable is the table name for the Resource entity.
	// It exists in this package in order to avoid circular dependency with the "resource" package.
	ResourcesInverseTable = "sys_resources"
	// ResourcesColumn is the table column denoting the resources relation/edge.
	ResourcesColumn = "menu_id"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "sys_role_menus"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "sys_roles"
	// RoleMenusTable is the table that holds the role_menus relation/edge.
	RoleMenusTable = "sys_role_menus"
	// RoleMenusInverseTable is the table name for the RoleMenu entity.
	// It exists in this package in order to avoid circular dependency with the "rolemenu" package.
	RoleMenusInverseTable = "sys_role_menus"
	// RoleMenusColumn is the table column denoting the role_menus relation/edge.
	RoleMenusColumn = "menu_id"
)

// Columns holds all SQL columns for menu fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldKeyword,
	FieldName,
	FieldDescription,
	FieldType,
	FieldIcon,
	FieldPath,
	FieldStatus,
	FieldParentID,
	FieldParentPath,
	FieldSequence,
	FieldProperties,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"role_id", "menu_id"}
)

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
	// DefaultKeyword holds the default value on creation for the "keyword" field.
	DefaultKeyword string
	// KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	KeywordValidator func(string) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType uint8
	// DefaultIcon holds the default value on creation for the "icon" field.
	DefaultIcon string
	// IconValidator is a validator for the "icon" field. It is called by the builders before save.
	IconValidator func(string) error
	// DefaultPath holds the default value on creation for the "path" field.
	DefaultPath string
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(int) error
	// DefaultParentPath holds the default value on creation for the "parent_path" field.
	DefaultParentPath string
	// ParentPathValidator is a validator for the "parent_path" field. It is called by the builders before save.
	ParentPathValidator func(string) error
	// DefaultSequence holds the default value on creation for the "sequence" field.
	DefaultSequence int
	// DefaultProperties holds the default value on creation for the "properties" field.
	DefaultProperties string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)

// OrderOption defines the ordering options for the Menu queries.
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

// ByKeyword orders the results by the keyword field.
func ByKeyword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeyword, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByIcon orders the results by the icon field.
func ByIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIcon, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
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

// BySequence orders the results by the sequence field.
func BySequence(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSequence, opts...).ToFunc()
}

// ByProperties orders the results by the properties field.
func ByProperties(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProperties, opts...).ToFunc()
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

// ByResourcesCount orders the results by resources count.
func ByResourcesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newResourcesStep(), opts...)
	}
}

// ByResources orders the results by resources terms.
func ByResources(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResourcesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRolesCount orders the results by roles count.
func ByRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolesStep(), opts...)
	}
}

// ByRoles orders the results by roles terms.
func ByRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoleMenusCount orders the results by role_menus count.
func ByRoleMenusCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoleMenusStep(), opts...)
	}
}

// ByRoleMenus orders the results by role_menus terms.
func ByRoleMenus(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleMenusStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newResourcesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResourcesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ResourcesTable, ResourcesColumn),
	)
}
func newRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
	)
}
func newRoleMenusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleMenusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, RoleMenusTable, RoleMenusColumn),
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
