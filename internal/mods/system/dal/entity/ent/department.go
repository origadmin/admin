// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// department:table:comment
type Department struct {
	config `json:"-"`
	// ID of the ent.
	// field:primary_key:comment
	ID int64 `json:"id,omitempty"`
	// create_time:field:comment
	CreateTime time.Time `json:"create_time,omitempty"`
	// update_time:field:comment
	UpdateTime time.Time `json:"update_time,omitempty"`
	// department:field:keyword
	Keyword string `json:"keyword,omitempty"`
	// department:field:name
	Name string `json:"name,omitempty"`
	// department:field:description
	Description string `json:"description,omitempty"`
	// department:field:sequence
	Sequence int `json:"sequence,omitempty"`
	// department:field:status
	Status int8 `json:"status,omitempty"`
	// department:field:ancestors
	Ancestors string `json:"ancestors,omitempty"`
	// department:field:level
	Level int `json:"level,omitempty"`
	// department:field:parent_id
	ParentID int64 `json:"parent_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DepartmentQuery when eager-loading is set.
	Edges        DepartmentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DepartmentEdges holds the relations/edges for other nodes in the graph.
type DepartmentEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Positions holds the value of the positions edge.
	Positions []*Position `json:"positions,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Children holds the value of the children edge.
	Children []*Department `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Department `json:"parent,omitempty"`
	// UserDepartments holds the value of the user_departments edge.
	UserDepartments []*UserDepartment `json:"user_departments,omitempty"`
	// DepartmentRoles holds the value of the department_roles edge.
	DepartmentRoles []*DepartmentRole `json:"department_roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// PositionsOrErr returns the Positions value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) PositionsOrErr() ([]*Position, error) {
	if e.loadedTypes[1] {
		return e.Positions, nil
	}
	return nil, &NotLoadedError{edge: "positions"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[2] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) ChildrenOrErr() ([]*Department, error) {
	if e.loadedTypes[3] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DepartmentEdges) ParentOrErr() (*Department, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: department.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// UserDepartmentsOrErr returns the UserDepartments value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) UserDepartmentsOrErr() ([]*UserDepartment, error) {
	if e.loadedTypes[5] {
		return e.UserDepartments, nil
	}
	return nil, &NotLoadedError{edge: "user_departments"}
}

// DepartmentRolesOrErr returns the DepartmentRoles value or an error if the edge
// was not loaded in eager-loading.
func (e DepartmentEdges) DepartmentRolesOrErr() ([]*DepartmentRole, error) {
	if e.loadedTypes[6] {
		return e.DepartmentRoles, nil
	}
	return nil, &NotLoadedError{edge: "department_roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Department) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case department.FieldID, department.FieldSequence, department.FieldStatus, department.FieldLevel, department.FieldParentID:
			values[i] = new(sql.NullInt64)
		case department.FieldKeyword, department.FieldName, department.FieldDescription, department.FieldAncestors:
			values[i] = new(sql.NullString)
		case department.FieldCreateTime, department.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Department fields.
func (d *Department) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case department.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int64(value.Int64)
		case department.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				d.CreateTime = value.Time
			}
		case department.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				d.UpdateTime = value.Time
			}
		case department.FieldKeyword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keyword", values[i])
			} else if value.Valid {
				d.Keyword = value.String
			}
		case department.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case department.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = value.String
			}
		case department.FieldSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence", values[i])
			} else if value.Valid {
				d.Sequence = int(value.Int64)
			}
		case department.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				d.Status = int8(value.Int64)
			}
		case department.FieldAncestors:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ancestors", values[i])
			} else if value.Valid {
				d.Ancestors = value.String
			}
		case department.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				d.Level = int(value.Int64)
			}
		case department.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				d.ParentID = value.Int64
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Department.
// This includes values selected through modifiers, order, etc.
func (d *Department) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Department entity.
func (d *Department) QueryUsers() *UserQuery {
	return NewDepartmentClient(d.config).QueryUsers(d)
}

// QueryPositions queries the "positions" edge of the Department entity.
func (d *Department) QueryPositions() *PositionQuery {
	return NewDepartmentClient(d.config).QueryPositions(d)
}

// QueryRoles queries the "roles" edge of the Department entity.
func (d *Department) QueryRoles() *RoleQuery {
	return NewDepartmentClient(d.config).QueryRoles(d)
}

// QueryChildren queries the "children" edge of the Department entity.
func (d *Department) QueryChildren() *DepartmentQuery {
	return NewDepartmentClient(d.config).QueryChildren(d)
}

// QueryParent queries the "parent" edge of the Department entity.
func (d *Department) QueryParent() *DepartmentQuery {
	return NewDepartmentClient(d.config).QueryParent(d)
}

// QueryUserDepartments queries the "user_departments" edge of the Department entity.
func (d *Department) QueryUserDepartments() *UserDepartmentQuery {
	return NewDepartmentClient(d.config).QueryUserDepartments(d)
}

// QueryDepartmentRoles queries the "department_roles" edge of the Department entity.
func (d *Department) QueryDepartmentRoles() *DepartmentRoleQuery {
	return NewDepartmentClient(d.config).QueryDepartmentRoles(d)
}

// Update returns a builder for updating this Department.
// Note that you need to call Department.Unwrap() before calling this method if this Department
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Department) Update() *DepartmentUpdateOne {
	return NewDepartmentClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Department entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Department) Unwrap() *Department {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Department is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Department) String() string {
	var builder strings.Builder
	builder.WriteString("Department(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("create_time=")
	builder.WriteString(d.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(d.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("keyword=")
	builder.WriteString(d.Keyword)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(d.Description)
	builder.WriteString(", ")
	builder.WriteString("sequence=")
	builder.WriteString(fmt.Sprintf("%v", d.Sequence))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", d.Status))
	builder.WriteString(", ")
	builder.WriteString("ancestors=")
	builder.WriteString(d.Ancestors)
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", d.Level))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", d.ParentID))
	builder.WriteByte(')')
	return builder.String()
}

// Departments is a parsable slice of Department.
type Departments []*Department
