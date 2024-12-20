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

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Sequence holds the value of the "sequence" field.
	Sequence int `json:"sequence,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
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
	// RoleMenus holds the value of the role_menus edge.
	RoleMenus []*RoleMenu `json:"role_menus,omitempty"`
	// UserRoles holds the value of the user_roles edge.
	UserRoles []*UserRole `json:"user_roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
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

// RoleMenusOrErr returns the RoleMenus value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) RoleMenusOrErr() ([]*RoleMenu, error) {
	if e.loadedTypes[2] {
		return e.RoleMenus, nil
	}
	return nil, &NotLoadedError{edge: "role_menus"}
}

// UserRolesOrErr returns the UserRoles value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) UserRolesOrErr() ([]*UserRole, error) {
	if e.loadedTypes[3] {
		return e.UserRoles, nil
	}
	return nil, &NotLoadedError{edge: "user_roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldSequence, role.FieldStatus:
			values[i] = new(sql.NullInt64)
		case role.FieldID, role.FieldCode, role.FieldName, role.FieldDescription:
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
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
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
		case role.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				r.Code = value.String
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

// QueryRoleMenus queries the "role_menus" edge of the Role entity.
func (r *Role) QueryRoleMenus() *RoleMenuQuery {
	return NewRoleClient(r.config).QueryRoleMenus(r)
}

// QueryUserRoles queries the "user_roles" edge of the Role entity.
func (r *Role) QueryUserRoles() *UserRoleQuery {
	return NewRoleClient(r.config).QueryUserRoles(r)
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
	builder.WriteString("code=")
	builder.WriteString(r.Code)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteString(", ")
	builder.WriteString("sequence=")
	builder.WriteString(fmt.Sprintf("%v", r.Sequence))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role
