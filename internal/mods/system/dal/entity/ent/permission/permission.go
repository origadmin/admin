// Code generated by ent, DO NOT EDIT.

package permission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the permission type in the database.
	Label = "permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldKeyword holds the string denoting the keyword field in the database.
	FieldKeyword = "keyword"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldI18nKey holds the string denoting the i18n_key field in the database.
	FieldI18nKey = "i18n_key"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldScope holds the string denoting the scope field in the database.
	FieldScope = "scope"
	// FieldScopeDepts holds the string denoting the scope_depts field in the database.
	FieldScopeDepts = "scope_depts"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeMenus holds the string denoting the menus edge name in mutations.
	EdgeMenus = "menus"
	// EdgeResources holds the string denoting the resources edge name in mutations.
	EdgeResources = "resources"
	// EdgeRolePermissions holds the string denoting the role_permissions edge name in mutations.
	EdgeRolePermissions = "role_permissions"
	// EdgeMenuPermissions holds the string denoting the menu_permissions edge name in mutations.
	EdgeMenuPermissions = "menu_permissions"
	// EdgePermissionResources holds the string denoting the permission_resources edge name in mutations.
	EdgePermissionResources = "permission_resources"
	// Table holds the table name of the permission in the database.
	Table = "sys_permissions"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "sys_role_permission"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "sys_roles"
	// MenusTable is the table that holds the menus relation/edge. The primary key declared below.
	MenusTable = "sys_menu_permissions"
	// MenusInverseTable is the table name for the Menu entity.
	// It exists in this package in order to avoid circular dependency with the "menu" package.
	MenusInverseTable = "sys_menus"
	// ResourcesTable is the table that holds the resources relation/edge. The primary key declared below.
	ResourcesTable = "sys_permission_resources"
	// ResourcesInverseTable is the table name for the Resource entity.
	// It exists in this package in order to avoid circular dependency with the "resource" package.
	ResourcesInverseTable = "sys_resources"
	// RolePermissionsTable is the table that holds the role_permissions relation/edge.
	RolePermissionsTable = "sys_role_permission"
	// RolePermissionsInverseTable is the table name for the RolePermission entity.
	// It exists in this package in order to avoid circular dependency with the "rolepermission" package.
	RolePermissionsInverseTable = "sys_role_permission"
	// RolePermissionsColumn is the table column denoting the role_permissions relation/edge.
	RolePermissionsColumn = "permission_id"
	// MenuPermissionsTable is the table that holds the menu_permissions relation/edge.
	MenuPermissionsTable = "sys_menu_permissions"
	// MenuPermissionsInverseTable is the table name for the MenuPermission entity.
	// It exists in this package in order to avoid circular dependency with the "menupermission" package.
	MenuPermissionsInverseTable = "sys_menu_permissions"
	// MenuPermissionsColumn is the table column denoting the menu_permissions relation/edge.
	MenuPermissionsColumn = "permission_id"
	// PermissionResourcesTable is the table that holds the permission_resources relation/edge.
	PermissionResourcesTable = "sys_permission_resources"
	// PermissionResourcesInverseTable is the table name for the PermissionResource entity.
	// It exists in this package in order to avoid circular dependency with the "permissionresource" package.
	PermissionResourcesInverseTable = "sys_permission_resources"
	// PermissionResourcesColumn is the table column denoting the permission_resources relation/edge.
	PermissionResourcesColumn = "permission_id"
)

// Columns holds all SQL columns for permission fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldKeyword,
	FieldDescription,
	FieldI18nKey,
	FieldType,
	FieldScope,
	FieldScopeDepts,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"role_id", "permission_id"}
	// MenusPrimaryKey and MenusColumn2 are the table columns denoting the
	// primary key for the menus relation (M2M).
	MenusPrimaryKey = []string{"menu_id", "permission_id"}
	// ResourcesPrimaryKey and ResourcesColumn2 are the table columns denoting the
	// primary key for the resources relation (M2M).
	ResourcesPrimaryKey = []string{"permission_id", "resource_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	KeywordValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// I18nKeyValidator is a validator for the "i18n_key" field. It is called by the builders before save.
	I18nKeyValidator func(string) error
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int8
	// DefaultScope holds the default value on creation for the "scope" field.
	DefaultScope string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)

// OrderOption defines the ordering options for the Permission queries.
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

// ByKeyword orders the results by the keyword field.
func ByKeyword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeyword, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByI18nKey orders the results by the i18n_key field.
func ByI18nKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldI18nKey, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByScope orders the results by the scope field.
func ByScope(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScope, opts...).ToFunc()
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

// ByRolePermissionsCount orders the results by role_permissions count.
func ByRolePermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolePermissionsStep(), opts...)
	}
}

// ByRolePermissions orders the results by role_permissions terms.
func ByRolePermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolePermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMenuPermissionsCount orders the results by menu_permissions count.
func ByMenuPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMenuPermissionsStep(), opts...)
	}
}

// ByMenuPermissions orders the results by menu_permissions terms.
func ByMenuPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMenuPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPermissionResourcesCount orders the results by permission_resources count.
func ByPermissionResourcesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPermissionResourcesStep(), opts...)
	}
}

// ByPermissionResources orders the results by permission_resources terms.
func ByPermissionResources(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionResourcesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
	)
}
func newMenusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MenusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, MenusTable, MenusPrimaryKey...),
	)
}
func newResourcesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResourcesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ResourcesTable, ResourcesPrimaryKey...),
	)
}
func newRolePermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolePermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, RolePermissionsTable, RolePermissionsColumn),
	)
}
func newMenuPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MenuPermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, MenuPermissionsTable, MenuPermissionsColumn),
	)
}
func newPermissionResourcesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PermissionResourcesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, PermissionResourcesTable, PermissionResourcesColumn),
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
