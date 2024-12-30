// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// role:table:comment
type Role struct {
	config `json:"-"`
	// ID of the ent.
	// primary_key:comment
	ID int64 `json:"id,omitempty"`
	// create_time:comment
	CreateTime time.Time `json:"create_time,omitempty"`
	// update_time:comment
	UpdateTime time.Time `json:"update_time,omitempty"`
	// role:field:keyword
	Keyword string `json:"keyword,omitempty"`
	// role:field:name
	Name string `json:"name,omitempty"`
	// role:field:description
	Description string `json:"description,omitempty"`
	// role:field:type
	Type int8 `json:"type,omitempty"`
	// role:field:sequence
	Sequence int `json:"sequence,omitempty"`
	// role:field:status
	Status int8 `json:"status,omitempty"`
	// role:field:is_system
	IsSystem bool `json:"is_system,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges        RoleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Menus holds the value of the menus edge.
	Menus []*Menu `json:"menus,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Permissions holds the value of the permissions edge.
	Permissions []*Permission `json:"permissions,omitempty"`
	// Departments holds the value of the departments edge.
	Departments []*Department `json:"departments,omitempty"`
	// RoleMenus holds the value of the role_menus edge.
	RoleMenus []*RoleMenu `json:"role_menus,omitempty"`
	// UserRoles holds the value of the user_roles edge.
	UserRoles []*UserRole `json:"user_roles,omitempty"`
	// RolePermissions holds the value of the role_permissions edge.
	RolePermissions []*RolePermission `json:"role_permissions,omitempty"`
	// DepartmentRoles holds the value of the department_roles edge.
	DepartmentRoles []*DepartmentRole `json:"department_roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// MenusOrErr returns the Menus value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) MenusOrErr() ([]*Menu, error) {
	if e.loadedTypes[0] {
		return e.Menus, nil
	}
	return nil, &NotLoadedError{edge: "menus"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) PermissionsOrErr() ([]*Permission, error) {
	if e.loadedTypes[2] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// DepartmentsOrErr returns the Departments value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) DepartmentsOrErr() ([]*Department, error) {
	if e.loadedTypes[3] {
		return e.Departments, nil
	}
	return nil, &NotLoadedError{edge: "departments"}
}

// RoleMenusOrErr returns the RoleMenus value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) RoleMenusOrErr() ([]*RoleMenu, error) {
	if e.loadedTypes[4] {
		return e.RoleMenus, nil
	}
	return nil, &NotLoadedError{edge: "role_menus"}
}

// UserRolesOrErr returns the UserRoles value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) UserRolesOrErr() ([]*UserRole, error) {
	if e.loadedTypes[5] {
		return e.UserRoles, nil
	}
	return nil, &NotLoadedError{edge: "user_roles"}
}

// RolePermissionsOrErr returns the RolePermissions value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) RolePermissionsOrErr() ([]*RolePermission, error) {
	if e.loadedTypes[6] {
		return e.RolePermissions, nil
	}
	return nil, &NotLoadedError{edge: "role_permissions"}
}

// DepartmentRolesOrErr returns the DepartmentRoles value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) DepartmentRolesOrErr() ([]*DepartmentRole, error) {
	if e.loadedTypes[7] {
		return e.DepartmentRoles, nil
	}
	return nil, &NotLoadedError{edge: "department_roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldIsSystem:
			values[i] = new(sql.NullBool)
		case role.FieldID, role.FieldType, role.FieldSequence, role.FieldStatus:
			values[i] = new(sql.NullInt64)
		case role.FieldKeyword, role.FieldName, role.FieldDescription:
			values[i] = new(sql.NullString)
		case role.FieldCreateTime, role.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int64(value.Int64)
		case role.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = value.Time
			}
		case role.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				r.UpdateTime = value.Time
			}
		case role.FieldKeyword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keyword", values[i])
			} else if value.Valid {
				r.Keyword = value.String
			}
		case role.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case role.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case role.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				r.Type = int8(value.Int64)
			}
		case role.FieldSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence", values[i])
			} else if value.Valid {
				r.Sequence = int(value.Int64)
			}
		case role.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = int8(value.Int64)
			}
		case role.FieldIsSystem:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_system", values[i])
			} else if value.Valid {
				r.IsSystem = value.Bool
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Role.
// This includes values selected through modifiers, order, etc.
func (r *Role) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryMenus queries the "menus" edge of the Role entity.
func (r *Role) QueryMenus() *MenuQuery {
	return NewRoleClient(r.config).QueryMenus(r)
}

// QueryUsers queries the "users" edge of the Role entity.
func (r *Role) QueryUsers() *UserQuery {
	return NewRoleClient(r.config).QueryUsers(r)
}

// QueryPermissions queries the "permissions" edge of the Role entity.
func (r *Role) QueryPermissions() *PermissionQuery {
	return NewRoleClient(r.config).QueryPermissions(r)
}

// QueryDepartments queries the "departments" edge of the Role entity.
func (r *Role) QueryDepartments() *DepartmentQuery {
	return NewRoleClient(r.config).QueryDepartments(r)
}

// QueryRoleMenus queries the "role_menus" edge of the Role entity.
func (r *Role) QueryRoleMenus() *RoleMenuQuery {
	return NewRoleClient(r.config).QueryRoleMenus(r)
}

// QueryUserRoles queries the "user_roles" edge of the Role entity.
func (r *Role) QueryUserRoles() *UserRoleQuery {
	return NewRoleClient(r.config).QueryUserRoles(r)
}

// QueryRolePermissions queries the "role_permissions" edge of the Role entity.
func (r *Role) QueryRolePermissions() *RolePermissionQuery {
	return NewRoleClient(r.config).QueryRolePermissions(r)
}

// QueryDepartmentRoles queries the "department_roles" edge of the Role entity.
func (r *Role) QueryDepartmentRoles() *DepartmentRoleQuery {
	return NewRoleClient(r.config).QueryDepartmentRoles(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return NewRoleClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("create_time=")
	builder.WriteString(r.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(r.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("keyword=")
	builder.WriteString(r.Keyword)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", r.Type))
	builder.WriteString(", ")
	builder.WriteString("sequence=")
	builder.WriteString(fmt.Sprintf("%v", r.Sequence))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteString(", ")
	builder.WriteString("is_system=")
	builder.WriteString(fmt.Sprintf("%v", r.IsSystem))
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role
