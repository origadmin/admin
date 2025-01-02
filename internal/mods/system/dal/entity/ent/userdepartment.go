// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userdepartment"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// user_department.table.comment
type UserDepartment struct {
	config `json:"-"`
	// ID of the ent.
	// field.primary_key.comment
	ID int64 `json:"id,omitempty"`
	// field.foreign_key.comment
	UserID int64 `json:"user_id,omitempty"`
	// field.foreign_key.comment
	DepartmentID int64 `json:"department_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserDepartmentQuery when eager-loading is set.
	Edges        UserDepartmentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserDepartmentEdges holds the relations/edges for other nodes in the graph.
type UserDepartmentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Department holds the value of the department edge.
	Department *Department `json:"department,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserDepartmentEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserDepartmentEdges) DepartmentOrErr() (*Department, error) {
	if e.Department != nil {
		return e.Department, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: department.Label}
	}
	return nil, &NotLoadedError{edge: "department"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserDepartment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userdepartment.FieldID, userdepartment.FieldUserID, userdepartment.FieldDepartmentID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserDepartment fields.
func (ud *UserDepartment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userdepartment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ud.ID = int64(value.Int64)
		case userdepartment.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ud.UserID = value.Int64
			}
		case userdepartment.FieldDepartmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field department_id", values[i])
			} else if value.Valid {
				ud.DepartmentID = value.Int64
			}
		default:
			ud.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserDepartment.
// This includes values selected through modifiers, order, etc.
func (ud *UserDepartment) Value(name string) (ent.Value, error) {
	return ud.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserDepartment entity.
func (ud *UserDepartment) QueryUser() *UserQuery {
	return NewUserDepartmentClient(ud.config).QueryUser(ud)
}

// QueryDepartment queries the "department" edge of the UserDepartment entity.
func (ud *UserDepartment) QueryDepartment() *DepartmentQuery {
	return NewUserDepartmentClient(ud.config).QueryDepartment(ud)
}

// Update returns a builder for updating this UserDepartment.
// Note that you need to call UserDepartment.Unwrap() before calling this method if this UserDepartment
// was returned from a transaction, and the transaction was committed or rolled back.
func (ud *UserDepartment) Update() *UserDepartmentUpdateOne {
	return NewUserDepartmentClient(ud.config).UpdateOne(ud)
}

// Unwrap unwraps the UserDepartment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ud *UserDepartment) Unwrap() *UserDepartment {
	_tx, ok := ud.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserDepartment is not a transactional entity")
	}
	ud.config.driver = _tx.drv
	return ud
}

// String implements the fmt.Stringer.
func (ud *UserDepartment) String() string {
	var builder strings.Builder
	builder.WriteString("UserDepartment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ud.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ud.UserID))
	builder.WriteString(", ")
	builder.WriteString("department_id=")
	builder.WriteString(fmt.Sprintf("%v", ud.DepartmentID))
	builder.WriteByte(')')
	return builder.String()
}

// UserDepartments is a parsable slice of UserDepartment.
type UserDepartments []*UserDepartment
