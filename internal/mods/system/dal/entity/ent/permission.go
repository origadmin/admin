// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Permission is the model entity for the Permission schema.
type Permission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionQuery when eager-loading is set.
	Edges        PermissionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PermissionEdges holds the relations/edges for other nodes in the graph.
type PermissionEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// RolePermissions holds the value of the role_permissions edge.
	RolePermissions []*RolePermission `json:"role_permissions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// RolePermissionsOrErr returns the RolePermissions value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) RolePermissionsOrErr() ([]*RolePermission, error) {
	if e.loadedTypes[1] {
		return e.RolePermissions, nil
	}
	return nil, &NotLoadedError{edge: "role_permissions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			values[i] = new(sql.NullInt64)
		case permission.FieldName, permission.FieldDescription:
			values[i] = new(sql.NullString)
		case permission.FieldCreateTime, permission.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case permission.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pe.CreateTime = value.Time
			}
		case permission.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pe.UpdateTime = value.Time
			}
		case permission.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case permission.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pe.Description = value.String
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Permission.
// This includes values selected through modifiers, order, etc.
func (pe *Permission) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the Permission entity.
func (pe *Permission) QueryRoles() *RoleQuery {
	return NewPermissionClient(pe.config).QueryRoles(pe)
}

// QueryRolePermissions queries the "role_permissions" edge of the Permission entity.
func (pe *Permission) QueryRolePermissions() *RolePermissionQuery {
	return NewPermissionClient(pe.config).QueryRolePermissions(pe)
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return NewPermissionClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Permission is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("create_time=")
	builder.WriteString(pe.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(pe.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pe.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission
