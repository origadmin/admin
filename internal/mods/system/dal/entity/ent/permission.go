// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// permission:table:comment
type Permission struct {
	config `json:"-"`
	// ID of the ent.
	// primary_key:comment
	ID int64 `json:"id,omitempty"`
	// create_time:comment
	CreateTime time.Time `json:"create_time,omitempty"`
	// update_time:comment
	UpdateTime time.Time `json:"update_time,omitempty"`
	// 权限名称
	Name string `json:"name,omitempty"`
	// 权限标识符
	Keyword string `json:"keyword,omitempty"`
	// 权限描述
	Description string `json:"description,omitempty"`
	// 国际化标识符(如：permission.system.user.manage)
	I18nKey string `json:"i18n_key,omitempty"`
	// 权限类型：1-系统 2-菜单 3-数据 4-部门 5-资源
	Type int8 `json:"type,omitempty"`
	// 数据范围：self-仅本人 dept-本部门 sub_dept-本部门及下级 custom-自定义部门 all-所有
	Scope string `json:"scope,omitempty"`
	// 自定义数据范围的部门ID列表，当scope为custom时有效
	ScopeDepts []string `json:"scope_depts,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionQuery when eager-loading is set.
	Edges        PermissionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PermissionEdges holds the relations/edges for other nodes in the graph.
type PermissionEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Menus holds the value of the menus edge.
	Menus []*Menu `json:"menus,omitempty"`
	// Resources holds the value of the resources edge.
	Resources []*Resource `json:"resources,omitempty"`
	// RolePermissions holds the value of the role_permissions edge.
	RolePermissions []*RolePermission `json:"role_permissions,omitempty"`
	// MenuPermissions holds the value of the menu_permissions edge.
	MenuPermissions []*MenuPermission `json:"menu_permissions,omitempty"`
	// PermissionResources holds the value of the permission_resources edge.
	PermissionResources []*PermissionResource `json:"permission_resources,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// MenusOrErr returns the Menus value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) MenusOrErr() ([]*Menu, error) {
	if e.loadedTypes[1] {
		return e.Menus, nil
	}
	return nil, &NotLoadedError{edge: "menus"}
}

// ResourcesOrErr returns the Resources value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) ResourcesOrErr() ([]*Resource, error) {
	if e.loadedTypes[2] {
		return e.Resources, nil
	}
	return nil, &NotLoadedError{edge: "resources"}
}

// RolePermissionsOrErr returns the RolePermissions value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) RolePermissionsOrErr() ([]*RolePermission, error) {
	if e.loadedTypes[3] {
		return e.RolePermissions, nil
	}
	return nil, &NotLoadedError{edge: "role_permissions"}
}

// MenuPermissionsOrErr returns the MenuPermissions value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) MenuPermissionsOrErr() ([]*MenuPermission, error) {
	if e.loadedTypes[4] {
		return e.MenuPermissions, nil
	}
	return nil, &NotLoadedError{edge: "menu_permissions"}
}

// PermissionResourcesOrErr returns the PermissionResources value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) PermissionResourcesOrErr() ([]*PermissionResource, error) {
	if e.loadedTypes[5] {
		return e.PermissionResources, nil
	}
	return nil, &NotLoadedError{edge: "permission_resources"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldScopeDepts:
			values[i] = new([]byte)
		case permission.FieldID, permission.FieldType:
			values[i] = new(sql.NullInt64)
		case permission.FieldName, permission.FieldKeyword, permission.FieldDescription, permission.FieldI18nKey, permission.FieldScope:
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
			pe.ID = int64(value.Int64)
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
		case permission.FieldKeyword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keyword", values[i])
			} else if value.Valid {
				pe.Keyword = value.String
			}
		case permission.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pe.Description = value.String
			}
		case permission.FieldI18nKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field i18n_key", values[i])
			} else if value.Valid {
				pe.I18nKey = value.String
			}
		case permission.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pe.Type = int8(value.Int64)
			}
		case permission.FieldScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value.Valid {
				pe.Scope = value.String
			}
		case permission.FieldScopeDepts:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field scope_depts", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.ScopeDepts); err != nil {
					return fmt.Errorf("unmarshal field scope_depts: %w", err)
				}
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

// QueryMenus queries the "menus" edge of the Permission entity.
func (pe *Permission) QueryMenus() *MenuQuery {
	return NewPermissionClient(pe.config).QueryMenus(pe)
}

// QueryResources queries the "resources" edge of the Permission entity.
func (pe *Permission) QueryResources() *ResourceQuery {
	return NewPermissionClient(pe.config).QueryResources(pe)
}

// QueryRolePermissions queries the "role_permissions" edge of the Permission entity.
func (pe *Permission) QueryRolePermissions() *RolePermissionQuery {
	return NewPermissionClient(pe.config).QueryRolePermissions(pe)
}

// QueryMenuPermissions queries the "menu_permissions" edge of the Permission entity.
func (pe *Permission) QueryMenuPermissions() *MenuPermissionQuery {
	return NewPermissionClient(pe.config).QueryMenuPermissions(pe)
}

// QueryPermissionResources queries the "permission_resources" edge of the Permission entity.
func (pe *Permission) QueryPermissionResources() *PermissionResourceQuery {
	return NewPermissionClient(pe.config).QueryPermissionResources(pe)
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
	builder.WriteString("keyword=")
	builder.WriteString(pe.Keyword)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pe.Description)
	builder.WriteString(", ")
	builder.WriteString("i18n_key=")
	builder.WriteString(pe.I18nKey)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", pe.Type))
	builder.WriteString(", ")
	builder.WriteString("scope=")
	builder.WriteString(pe.Scope)
	builder.WriteString(", ")
	builder.WriteString("scope_depts=")
	builder.WriteString(fmt.Sprintf("%v", pe.ScopeDepts))
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission
