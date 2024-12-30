// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// menu:table:comment
type Menu struct {
	config `json:"-"`
	// ID of the ent.
	// field:primary_key:comment
	ID int64 `json:"id,omitempty"`
	// create_time:field:comment
	CreateTime time.Time `json:"create_time,omitempty"`
	// update_time:field:comment
	UpdateTime time.Time `json:"update_time,omitempty"`
	// menu:field:keyword
	Keyword string `json:"keyword,omitempty"`
	// menu:field:name
	Name string `json:"name,omitempty"`
	// menu:field:i18n_key
	I18nKey string `json:"i18n_key,omitempty"`
	// menu:field:description
	Description string `json:"description,omitempty"`
	// menu:field:type
	Type string `json:"type,omitempty"`
	// menu:field:icon
	Icon string `json:"icon,omitempty"`
	// menu:field:path
	Path string `json:"path,omitempty"`
	// menu:field:status
	Status int8 `json:"status,omitempty"`
	// menu:field:parent_path
	ParentPath string `json:"parent_path,omitempty"`
	// menu:field:sequence
	Sequence int `json:"sequence,omitempty"`
	// menu:field:properties
	Properties string `json:"properties,omitempty"`
	// menu:field:parent_id
	ParentID int64 `json:"parent_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges        MenuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Children holds the value of the children edge.
	Children []*Menu `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Menu `json:"parent,omitempty"`
	// Resources holds the value of the resources edge.
	Resources []*Resource `json:"resources,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Permissions holds the value of the permissions edge.
	Permissions []*Permission `json:"permissions,omitempty"`
	// RoleMenus holds the value of the role_menus edge.
	RoleMenus []*RoleMenu `json:"role_menus,omitempty"`
	// MenuPermissions holds the value of the menu_permissions edge.
	MenuPermissions []*MenuPermission `json:"menu_permissions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) ChildrenOrErr() ([]*Menu, error) {
	if e.loadedTypes[0] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) ParentOrErr() (*Menu, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: menu.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ResourcesOrErr returns the Resources value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) ResourcesOrErr() ([]*Resource, error) {
	if e.loadedTypes[2] {
		return e.Resources, nil
	}
	return nil, &NotLoadedError{edge: "resources"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[3] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) PermissionsOrErr() ([]*Permission, error) {
	if e.loadedTypes[4] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// RoleMenusOrErr returns the RoleMenus value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) RoleMenusOrErr() ([]*RoleMenu, error) {
	if e.loadedTypes[5] {
		return e.RoleMenus, nil
	}
	return nil, &NotLoadedError{edge: "role_menus"}
}

// MenuPermissionsOrErr returns the MenuPermissions value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) MenuPermissionsOrErr() ([]*MenuPermission, error) {
	if e.loadedTypes[6] {
		return e.MenuPermissions, nil
	}
	return nil, &NotLoadedError{edge: "menu_permissions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menu.FieldID, menu.FieldStatus, menu.FieldSequence, menu.FieldParentID:
			values[i] = new(sql.NullInt64)
		case menu.FieldKeyword, menu.FieldName, menu.FieldI18nKey, menu.FieldDescription, menu.FieldType, menu.FieldIcon, menu.FieldPath, menu.FieldParentPath, menu.FieldProperties:
			values[i] = new(sql.NullString)
		case menu.FieldCreateTime, menu.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int64(value.Int64)
		case menu.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				m.CreateTime = value.Time
			}
		case menu.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				m.UpdateTime = value.Time
			}
		case menu.FieldKeyword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keyword", values[i])
			} else if value.Valid {
				m.Keyword = value.String
			}
		case menu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case menu.FieldI18nKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field i18n_key", values[i])
			} else if value.Valid {
				m.I18nKey = value.String
			}
		case menu.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				m.Description = value.String
			}
		case menu.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				m.Type = value.String
			}
		case menu.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				m.Icon = value.String
			}
		case menu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				m.Path = value.String
			}
		case menu.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = int8(value.Int64)
			}
		case menu.FieldParentPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_path", values[i])
			} else if value.Valid {
				m.ParentPath = value.String
			}
		case menu.FieldSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence", values[i])
			} else if value.Valid {
				m.Sequence = int(value.Int64)
			}
		case menu.FieldProperties:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field properties", values[i])
			} else if value.Valid {
				m.Properties = value.String
			}
		case menu.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				m.ParentID = value.Int64
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Menu.
// This includes values selected through modifiers, order, etc.
func (m *Menu) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryChildren queries the "children" edge of the Menu entity.
func (m *Menu) QueryChildren() *MenuQuery {
	return NewMenuClient(m.config).QueryChildren(m)
}

// QueryParent queries the "parent" edge of the Menu entity.
func (m *Menu) QueryParent() *MenuQuery {
	return NewMenuClient(m.config).QueryParent(m)
}

// QueryResources queries the "resources" edge of the Menu entity.
func (m *Menu) QueryResources() *ResourceQuery {
	return NewMenuClient(m.config).QueryResources(m)
}

// QueryRoles queries the "roles" edge of the Menu entity.
func (m *Menu) QueryRoles() *RoleQuery {
	return NewMenuClient(m.config).QueryRoles(m)
}

// QueryPermissions queries the "permissions" edge of the Menu entity.
func (m *Menu) QueryPermissions() *PermissionQuery {
	return NewMenuClient(m.config).QueryPermissions(m)
}

// QueryRoleMenus queries the "role_menus" edge of the Menu entity.
func (m *Menu) QueryRoleMenus() *RoleMenuQuery {
	return NewMenuClient(m.config).QueryRoleMenus(m)
}

// QueryMenuPermissions queries the "menu_permissions" edge of the Menu entity.
func (m *Menu) QueryMenuPermissions() *MenuPermissionQuery {
	return NewMenuClient(m.config).QueryMenuPermissions(m)
}

// Update returns a builder for updating this Menu.
// Note that you need to call Menu.Unwrap() before calling this method if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return NewMenuClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Menu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("create_time=")
	builder.WriteString(m.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(m.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("keyword=")
	builder.WriteString(m.Keyword)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("i18n_key=")
	builder.WriteString(m.I18nKey)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(m.Description)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(m.Type)
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(m.Icon)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(m.Path)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ")
	builder.WriteString("parent_path=")
	builder.WriteString(m.ParentPath)
	builder.WriteString(", ")
	builder.WriteString("sequence=")
	builder.WriteString(fmt.Sprintf("%v", m.Sequence))
	builder.WriteString(", ")
	builder.WriteString("properties=")
	builder.WriteString(m.Properties)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", m.ParentID))
	builder.WriteByte(')')
	return builder.String()
}

// Menus is a parsable slice of Menu.
type Menus []*Menu
