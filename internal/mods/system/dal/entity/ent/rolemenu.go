// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolemenu"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// RoleMenu is the model entity for the RoleMenu schema.
type RoleMenu struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID string `json:"role_id,omitempty"`
	// MenuID holds the value of the "menu_id" field.
	MenuID string `json:"menu_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleMenuQuery when eager-loading is set.
	Edges        RoleMenuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoleMenuEdges holds the relations/edges for other nodes in the graph.
type RoleMenuEdges struct {
	// Role holds the value of the role edge.
	Role *Role `json:"role,omitempty"`
	// Menu holds the value of the menu edge.
	Menu *Menu `json:"menu,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleMenuEdges) RoleOrErr() (*Role, error) {
	if e.Role != nil {
		return e.Role, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: role.Label}
	}
	return nil, &NotLoadedError{edge: "role"}
}

// MenuOrErr returns the Menu value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleMenuEdges) MenuOrErr() (*Menu, error) {
	if e.Menu != nil {
		return e.Menu, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: menu.Label}
	}
	return nil, &NotLoadedError{edge: "menu"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoleMenu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rolemenu.FieldID, rolemenu.FieldRoleID, rolemenu.FieldMenuID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoleMenu fields.
func (rm *RoleMenu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rolemenu.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				rm.ID = value.String
			}
		case rolemenu.FieldRoleID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				rm.RoleID = value.String
			}
		case rolemenu.FieldMenuID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menu_id", values[i])
			} else if value.Valid {
				rm.MenuID = value.String
			}
		default:
			rm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoleMenu.
// This includes values selected through modifiers, order, etc.
func (rm *RoleMenu) Value(name string) (ent.Value, error) {
	return rm.selectValues.Get(name)
}

// QueryRole queries the "role" edge of the RoleMenu entity.
func (rm *RoleMenu) QueryRole() *RoleQuery {
	return NewRoleMenuClient(rm.config).QueryRole(rm)
}

// QueryMenu queries the "menu" edge of the RoleMenu entity.
func (rm *RoleMenu) QueryMenu() *MenuQuery {
	return NewRoleMenuClient(rm.config).QueryMenu(rm)
}

// Update returns a builder for updating this RoleMenu.
// Note that you need to call RoleMenu.Unwrap() before calling this method if this RoleMenu
// was returned from a transaction, and the transaction was committed or rolled back.
func (rm *RoleMenu) Update() *RoleMenuUpdateOne {
	return NewRoleMenuClient(rm.config).UpdateOne(rm)
}

// Unwrap unwraps the RoleMenu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rm *RoleMenu) Unwrap() *RoleMenu {
	_tx, ok := rm.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoleMenu is not a transactional entity")
	}
	rm.config.driver = _tx.drv
	return rm
}

// String implements the fmt.Stringer.
func (rm *RoleMenu) String() string {
	var builder strings.Builder
	builder.WriteString("RoleMenu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rm.ID))
	builder.WriteString("role_id=")
	builder.WriteString(rm.RoleID)
	builder.WriteString(", ")
	builder.WriteString("menu_id=")
	builder.WriteString(rm.MenuID)
	builder.WriteByte(')')
	return builder.String()
}

// RoleMenus is a parsable slice of RoleMenu.
type RoleMenus []*RoleMenu
