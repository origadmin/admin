// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// resource.table.comment
type Resource struct {
	config `json:"-"`
	// ID of the ent.
	// field.primary_key.comment
	ID int64 `json:"id,omitempty"`
	// create_time.field.comment
	CreateTime time.Time `json:"create_time,omitempty"`
	// update_time.field.comment
	UpdateTime time.Time `json:"update_time,omitempty"`
	// resource.field.name
	Name string `json:"name,omitempty"`
	// resource.field.keyword
	Keyword string `json:"keyword,omitempty"`
	// resource.field.i18n_key
	I18nKey string `json:"i18n_key,omitempty"`
	// resource.field.type
	Type string `json:"type,omitempty"`
	// resource.field.status
	Status int8 `json:"status,omitempty"`
	// resource.field.uri
	URI string `json:"uri,omitempty"`
	// resource.field.operation
	Operation string `json:"operation,omitempty"`
	// resource.field.method
	Method string `json:"method,omitempty"`
	// resource.field.component
	Component string `json:"component,omitempty"`
	// resource.field.icon
	Icon string `json:"icon,omitempty"`
	// resource.field.sequence
	Sequence int `json:"sequence,omitempty"`
	// resource.field.visible
	Visible bool `json:"visible,omitempty"`
	// resource.field.tree_path
	TreePath string `json:"tree_path,omitempty"`
	// resource.field.properties
	Properties map[string]string `json:"properties,omitempty"`
	// resource.field.description
	Description string `json:"description,omitempty"`
	// resource.field.parent_id
	ParentID int64 `json:"parent_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceQuery when eager-loading is set.
	Edges        ResourceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ResourceEdges holds the relations/edges for other nodes in the graph.
type ResourceEdges struct {
	// Children holds the value of the children edge.
	Children []*Resource `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Resource `json:"parent,omitempty"`
	// Permissions holds the value of the permissions edge.
	Permissions []*Permission `json:"permissions,omitempty"`
	// PermissionResources holds the value of the permission_resources edge.
	PermissionResources []*PermissionResource `json:"permission_resources,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceEdges) ChildrenOrErr() ([]*Resource, error) {
	if e.loadedTypes[0] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceEdges) ParentOrErr() (*Resource, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: resource.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceEdges) PermissionsOrErr() ([]*Permission, error) {
	if e.loadedTypes[2] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// PermissionResourcesOrErr returns the PermissionResources value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceEdges) PermissionResourcesOrErr() ([]*PermissionResource, error) {
	if e.loadedTypes[3] {
		return e.PermissionResources, nil
	}
	return nil, &NotLoadedError{edge: "permission_resources"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Resource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resource.FieldProperties:
			values[i] = new([]byte)
		case resource.FieldVisible:
			values[i] = new(sql.NullBool)
		case resource.FieldID, resource.FieldStatus, resource.FieldSequence, resource.FieldParentID:
			values[i] = new(sql.NullInt64)
		case resource.FieldName, resource.FieldKeyword, resource.FieldI18nKey, resource.FieldType, resource.FieldURI, resource.FieldOperation, resource.FieldMethod, resource.FieldComponent, resource.FieldIcon, resource.FieldTreePath, resource.FieldDescription:
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
			r.ID = int64(value.Int64)
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
		case resource.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case resource.FieldKeyword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keyword", values[i])
			} else if value.Valid {
				r.Keyword = value.String
			}
		case resource.FieldI18nKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field i18n_key", values[i])
			} else if value.Valid {
				r.I18nKey = value.String
			}
		case resource.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				r.Type = value.String
			}
		case resource.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = int8(value.Int64)
			}
		case resource.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				r.URI = value.String
			}
		case resource.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				r.Operation = value.String
			}
		case resource.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				r.Method = value.String
			}
		case resource.FieldComponent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field component", values[i])
			} else if value.Valid {
				r.Component = value.String
			}
		case resource.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				r.Icon = value.String
			}
		case resource.FieldSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence", values[i])
			} else if value.Valid {
				r.Sequence = int(value.Int64)
			}
		case resource.FieldVisible:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field visible", values[i])
			} else if value.Valid {
				r.Visible = value.Bool
			}
		case resource.FieldTreePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tree_path", values[i])
			} else if value.Valid {
				r.TreePath = value.String
			}
		case resource.FieldProperties:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field properties", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Properties); err != nil {
					return fmt.Errorf("unmarshal field properties: %w", err)
				}
			}
		case resource.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case resource.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				r.ParentID = value.Int64
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

// QueryChildren queries the "children" edge of the Resource entity.
func (r *Resource) QueryChildren() *ResourceQuery {
	return NewResourceClient(r.config).QueryChildren(r)
}

// QueryParent queries the "parent" edge of the Resource entity.
func (r *Resource) QueryParent() *ResourceQuery {
	return NewResourceClient(r.config).QueryParent(r)
}

// QueryPermissions queries the "permissions" edge of the Resource entity.
func (r *Resource) QueryPermissions() *PermissionQuery {
	return NewResourceClient(r.config).QueryPermissions(r)
}

// QueryPermissionResources queries the "permission_resources" edge of the Resource entity.
func (r *Resource) QueryPermissionResources() *PermissionResourceQuery {
	return NewResourceClient(r.config).QueryPermissionResources(r)
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
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("keyword=")
	builder.WriteString(r.Keyword)
	builder.WriteString(", ")
	builder.WriteString("i18n_key=")
	builder.WriteString(r.I18nKey)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(r.Type)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(r.URI)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(r.Operation)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(r.Method)
	builder.WriteString(", ")
	builder.WriteString("component=")
	builder.WriteString(r.Component)
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(r.Icon)
	builder.WriteString(", ")
	builder.WriteString("sequence=")
	builder.WriteString(fmt.Sprintf("%v", r.Sequence))
	builder.WriteString(", ")
	builder.WriteString("visible=")
	builder.WriteString(fmt.Sprintf("%v", r.Visible))
	builder.WriteString(", ")
	builder.WriteString("tree_path=")
	builder.WriteString(r.TreePath)
	builder.WriteString(", ")
	builder.WriteString("properties=")
	builder.WriteString(fmt.Sprintf("%v", r.Properties))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", r.ParentID))
	builder.WriteByte(')')
	return builder.String()
}

// Resources is a parsable slice of Resource.
type Resources []*Resource
