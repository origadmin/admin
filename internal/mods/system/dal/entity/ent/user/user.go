// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateAuthor holds the string denoting the create_author field in the database.
	FieldCreateAuthor = "create_author"
	// FieldUpdateAuthor holds the string denoting the update_author field in the database.
	FieldUpdateAuthor = "update_author"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldAllowedIP holds the string denoting the allowed_ip field in the database.
	FieldAllowedIP = "allowed_ip"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLastLoginIP holds the string denoting the last_login_ip field in the database.
	FieldLastLoginIP = "last_login_ip"
	// FieldLastLoginTime holds the string denoting the last_login_time field in the database.
	FieldLastLoginTime = "last_login_time"
	// FieldSanctionDate holds the string denoting the sanction_date field in the database.
	FieldSanctionDate = "sanction_date"
	// FieldManagerID holds the string denoting the manager_id field in the database.
	FieldManagerID = "manager_id"
	// FieldManager holds the string denoting the manager field in the database.
	FieldManager = "manager"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeDepartments holds the string denoting the departments edge name in mutations.
	EdgeDepartments = "departments"
	// EdgeUserRoles holds the string denoting the user_roles edge name in mutations.
	EdgeUserRoles = "user_roles"
	// EdgeUserDepartments holds the string denoting the user_departments edge name in mutations.
	EdgeUserDepartments = "user_departments"
	// Table holds the table name of the user in the database.
	Table = "sys_users"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "sys_user_roles"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "sys_roles"
	// DepartmentsTable is the table that holds the departments relation/edge. The primary key declared below.
	DepartmentsTable = "sys_user_departments"
	// DepartmentsInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentsInverseTable = "sys_departments"
	// UserRolesTable is the table that holds the user_roles relation/edge.
	UserRolesTable = "sys_user_roles"
	// UserRolesInverseTable is the table name for the UserRole entity.
	// It exists in this package in order to avoid circular dependency with the "userrole" package.
	UserRolesInverseTable = "sys_user_roles"
	// UserRolesColumn is the table column denoting the user_roles relation/edge.
	UserRolesColumn = "user_id"
	// UserDepartmentsTable is the table that holds the user_departments relation/edge.
	UserDepartmentsTable = "sys_user_departments"
	// UserDepartmentsInverseTable is the table name for the UserDepartment entity.
	// It exists in this package in order to avoid circular dependency with the "userdepartment" package.
	UserDepartmentsInverseTable = "sys_user_departments"
	// UserDepartmentsColumn is the table column denoting the user_departments relation/edge.
	UserDepartmentsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateAuthor,
	FieldUpdateAuthor,
	FieldCreateTime,
	FieldUpdateTime,
	FieldUUID,
	FieldAllowedIP,
	FieldUsername,
	FieldNickname,
	FieldAvatar,
	FieldName,
	FieldGender,
	FieldPassword,
	FieldSalt,
	FieldPhone,
	FieldEmail,
	FieldRemark,
	FieldToken,
	FieldStatus,
	FieldLastLoginIP,
	FieldLastLoginTime,
	FieldSanctionDate,
	FieldManagerID,
	FieldManager,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"user_id", "role_id"}
	// DepartmentsPrimaryKey and DepartmentsColumn2 are the table columns denoting the
	// primary key for the departments relation (M2M).
	DepartmentsPrimaryKey = []string{"user_id", "department_id"}
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
	// DefaultCreateAuthor holds the default value on creation for the "create_author" field.
	DefaultCreateAuthor int64
	// DefaultUpdateAuthor holds the default value on creation for the "update_author" field.
	DefaultUpdateAuthor int64
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	UUIDValidator func(string) error
	// DefaultAllowedIP holds the default value on creation for the "allowed_ip" field.
	DefaultAllowedIP string
	// DefaultUsername holds the default value on creation for the "username" field.
	DefaultUsername string
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultNickname holds the default value on creation for the "nickname" field.
	DefaultNickname string
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// DefaultAvatar holds the default value on creation for the "avatar" field.
	DefaultAvatar string
	// AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	AvatarValidator func(string) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultGender holds the default value on creation for the "gender" field.
	DefaultGender string
	// GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	GenderValidator func(string) error
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultSalt holds the default value on creation for the "salt" field.
	DefaultSalt string
	// SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	SaltValidator func(string) error
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
	// RemarkValidator is a validator for the "remark" field. It is called by the builders before save.
	RemarkValidator func(string) error
	// DefaultToken holds the default value on creation for the "token" field.
	DefaultToken string
	// TokenValidator is a validator for the "token" field. It is called by the builders before save.
	TokenValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// DefaultLastLoginIP holds the default value on creation for the "last_login_ip" field.
	DefaultLastLoginIP string
	// LastLoginIPValidator is a validator for the "last_login_ip" field. It is called by the builders before save.
	LastLoginIPValidator func(string) error
	// DefaultLastLoginTime holds the default value on creation for the "last_login_time" field.
	DefaultLastLoginTime func() time.Time
	// DefaultSanctionDate holds the default value on creation for the "sanction_date" field.
	DefaultSanctionDate func() time.Time
	// DefaultManagerID holds the default value on creation for the "manager_id" field.
	DefaultManagerID func() int64
	// ManagerIDValidator is a validator for the "manager_id" field. It is called by the builders before save.
	ManagerIDValidator func(int64) error
	// DefaultManager holds the default value on creation for the "manager" field.
	DefaultManager string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateAuthor orders the results by the create_author field.
func ByCreateAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateAuthor, opts...).ToFunc()
}

// ByUpdateAuthor orders the results by the update_author field.
func ByUpdateAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAuthor, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByAllowedIP orders the results by the allowed_ip field.
func ByAllowedIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAllowedIP, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// BySalt orders the results by the salt field.
func BySalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSalt, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLastLoginIP orders the results by the last_login_ip field.
func ByLastLoginIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginIP, opts...).ToFunc()
}

// ByLastLoginTime orders the results by the last_login_time field.
func ByLastLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginTime, opts...).ToFunc()
}

// BySanctionDate orders the results by the sanction_date field.
func BySanctionDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSanctionDate, opts...).ToFunc()
}

// ByManagerID orders the results by the manager_id field.
func ByManagerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManagerID, opts...).ToFunc()
}

// ByManager orders the results by the manager field.
func ByManager(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManager, opts...).ToFunc()
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

// ByDepartmentsCount orders the results by departments count.
func ByDepartmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDepartmentsStep(), opts...)
	}
}

// ByDepartments orders the results by departments terms.
func ByDepartments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDepartmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByUserDepartmentsCount orders the results by user_departments count.
func ByUserDepartmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserDepartmentsStep(), opts...)
	}
}

// ByUserDepartments orders the results by user_departments terms.
func ByUserDepartments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserDepartmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, RolesTable, RolesPrimaryKey...),
	)
}
func newDepartmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DepartmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, DepartmentsTable, DepartmentsPrimaryKey...),
	)
}
func newUserRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UserRolesTable, UserRolesColumn),
	)
}
func newUserDepartmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserDepartmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UserDepartmentsTable, UserDepartmentsColumn),
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
