// Code generated by ent, DO NOT EDIT.

package role

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
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
	// FieldSequence holds the string denoting the sequence field in the database.
	FieldSequence = "sequence"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeMenus holds the string denoting the menus edge name in mutations.
	EdgeMenus = "menus"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeRoleMenus holds the string denoting the role_menus edge name in mutations.
	EdgeRoleMenus = "role_menus"
	// EdgeUserRoles holds the string denoting the user_roles edge name in mutations.
	EdgeUserRoles = "user_roles"
	// Table holds the table name of the role in the database.
	Table = "sys_roles"
	// MenusTable is the table that holds the menus relation/edge. The primary key declared below.
	MenusTable = "sys_role_menus"
	// MenusInverseTable is the table name for the Menu entity.
	// It exists in this package in order to avoid circular dependency with the "menu" package.
	MenusInverseTable = "sys_menus"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "sys_user_roles"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "sys_users"
	// RoleMenusTable is the table that holds the role_menus relation/edge.
	RoleMenusTable = "sys_role_menus"
	// RoleMenusInverseTable is the table name for the RoleMenu entity.
	// It exists in this package in order to avoid circular dependency with the "rolemenu" package.
	RoleMenusInverseTable = "sys_role_menus"
	// RoleMenusColumn is the table column denoting the role_menus relation/edge.
	RoleMenusColumn = "role_id"
	// UserRolesTable is the table that holds the user_roles relation/edge.
	UserRolesTable = "sys_user_roles"
	// UserRolesInverseTable is the table name for the UserRole entity.
	// It exists in this package in order to avoid circular dependency with the "userrole" package.
	UserRolesInverseTable = "sys_user_roles"
	// UserRolesColumn is the table column denoting the user_roles relation/edge.
	UserRolesColumn = "role_id"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldKeyword,
	FieldName,
	FieldDescription,
	FieldSequence,
	FieldStatus,
}

var (
	// MenusPrimaryKey and MenusColumn2 are the table columns denoting the
	// primary key for the menus relation (M2M).
	MenusPrimaryKey = []string{"role_id", "menu_id"}
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "role_id"}
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Role queries.
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

// BySequence orders the results by the sequence field.
func BySequence(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSequence, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByMenusCount orders the results by menus count.
func ByMenusCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMenusStep(), opts...)
	}
}

// ByMenus orders the results by menus terms.
func ByMenus(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMenusStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByUserRolesCount orders the results by user_roles count.
func ByUserRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserRolesStep(), opts...)
	}
}

// ByUserRoles orders the results by user_roles terms.
func ByUserRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMenusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MenusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, MenusTable, MenusPrimaryKey...),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
	)
}
func newRoleMenusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleMenusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, RoleMenusTable, RoleMenusColumn),
	)
}
func newUserRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UserRolesTable, UserRolesColumn),
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
