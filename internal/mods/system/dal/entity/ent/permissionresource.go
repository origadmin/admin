// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// permission_resource.table.comment
type PermissionResource struct {
	config `json:"-"`
	// ID of the ent.
	// field.primary_key.comment
	ID int64 `json:"id,omitempty"`
	// field.foreign_key.comment
	PermissionID int64 `json:"permission_id,omitempty"`
	// field.foreign_key.comment
	ResourceID int64 `json:"resource_id,omitempty"`
	// permission_resource.field.actions
	Actions string `json:"actions,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionResourceQuery when eager-loading is set.
	Edges        PermissionResourceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PermissionResourceEdges holds the relations/edges for other nodes in the graph.
type PermissionResourceEdges struct {
	// Permission holds the value of the permission edge.
	Permission *Permission `json:"permission,omitempty"`
	// Resource holds the value of the resource edge.
	Resource *Resource `json:"resource,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PermissionOrErr returns the Permission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionResourceEdges) PermissionOrErr() (*Permission, error) {
	if e.Permission != nil {
		return e.Permission, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: permission.Label}
	}
	return nil, &NotLoadedError{edge: "permission"}
}

// ResourceOrErr returns the Resource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionResourceEdges) ResourceOrErr() (*Resource, error) {
	if e.Resource != nil {
		return e.Resource, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: resource.Label}
	}
	return nil, &NotLoadedError{edge: "resource"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PermissionResource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permissionresource.FieldID, permissionresource.FieldPermissionID, permissionresource.FieldResourceID:
			values[i] = new(sql.NullInt64)
		case permissionresource.FieldActions:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PermissionResource fields.
func (pr *PermissionResource) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permissionresource.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int64(value.Int64)
		case permissionresource.FieldPermissionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field permission_id", values[i])
			} else if value.Valid {
				pr.PermissionID = value.Int64
			}
		case permissionresource.FieldResourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field resource_id", values[i])
			} else if value.Valid {
				pr.ResourceID = value.Int64
			}
		case permissionresource.FieldActions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actions", values[i])
			} else if value.Valid {
				pr.Actions = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PermissionResource.
// This includes values selected through modifiers, order, etc.
func (pr *PermissionResource) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryPermission queries the "permission" edge of the PermissionResource entity.
func (pr *PermissionResource) QueryPermission() *PermissionQuery {
	return NewPermissionResourceClient(pr.config).QueryPermission(pr)
}

// QueryResource queries the "resource" edge of the PermissionResource entity.
func (pr *PermissionResource) QueryResource() *ResourceQuery {
	return NewPermissionResourceClient(pr.config).QueryResource(pr)
}

// Update returns a builder for updating this PermissionResource.
// Note that you need to call PermissionResource.Unwrap() before calling this method if this PermissionResource
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *PermissionResource) Update() *PermissionResourceUpdateOne {
	return NewPermissionResourceClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the PermissionResource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *PermissionResource) Unwrap() *PermissionResource {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PermissionResource is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *PermissionResource) String() string {
	var builder strings.Builder
	builder.WriteString("PermissionResource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("permission_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.PermissionID))
	builder.WriteString(", ")
	builder.WriteString("resource_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.ResourceID))
	builder.WriteString(", ")
	builder.WriteString("actions=")
	builder.WriteString(pr.Actions)
	builder.WriteByte(')')
	return builder.String()
}

// PermissionResources is a parsable slice of PermissionResource.
type PermissionResources []*PermissionResource
