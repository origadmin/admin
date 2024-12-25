// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Resource is the model entity for the Resource schema.
type Resource struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation string `json:"operation,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// MenuID holds the value of the "menu_id" field.
	MenuID int `json:"menu_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceQuery when eager-loading is set.
	Edges        ResourceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ResourceEdges holds the relations/edges for other nodes in the graph.
type ResourceEdges struct {
	// Menu holds the value of the menu edge.
	Menu *Menu `json:"menu,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MenuOrErr returns the Menu value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceEdges) MenuOrErr() (*Menu, error) {
	if e.Menu != nil {
		return e.Menu, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: menu.Label}
	}
	return nil, &NotLoadedError{edge: "menu"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Resource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resource.FieldID, resource.FieldMenuID:
			values[i] = new(sql.NullInt64)
		case resource.FieldMethod, resource.FieldOperation, resource.FieldPath:
			values[i] = new(sql.NullString)
		case resource.FieldCreateTime, resource.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Resource fields.
func (r *Resource) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resource.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case resource.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = value.Time
			}
		case resource.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				r.UpdateTime = value.Time
			}
		case resource.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				r.Method = value.String
			}
		case resource.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				r.Operation = value.String
			}
		case resource.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				r.Path = value.String
			}
		case resource.FieldMenuID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menu_id", values[i])
			} else if value.Valid {
				r.MenuID = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Resource.
// This includes values selected through modifiers, order, etc.
func (r *Resource) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryMenu queries the "menu" edge of the Resource entity.
func (r *Resource) QueryMenu() *MenuQuery {
	return NewResourceClient(r.config).QueryMenu(r)
}

// Update returns a builder for updating this Resource.
// Note that you need to call Resource.Unwrap() before calling this method if this Resource
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Resource) Update() *ResourceUpdateOne {
	return NewResourceClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Resource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Resource) Unwrap() *Resource {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Resource is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Resource) String() string {
	var builder strings.Builder
	builder.WriteString("Resource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("create_time=")
	builder.WriteString(r.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(r.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(r.Method)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(r.Operation)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(r.Path)
	builder.WriteString(", ")
	builder.WriteString("menu_id=")
	builder.WriteString(fmt.Sprintf("%v", r.MenuID))
	builder.WriteByte(')')
	return builder.String()
}

// Resources is a parsable slice of Resource.
type Resources []*Resource
