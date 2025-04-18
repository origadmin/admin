// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userposition"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// entity.user_position.table.comment
type UserPosition struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// field.foreign_key.comment
	UserID int64 `json:"user_id,omitempty"`
	// field.foreign_key.comment
	PositionID int64 `json:"position_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserPositionQuery when eager-loading is set.
	Edges        UserPositionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserPositionEdges holds the relations/edges for other nodes in the graph.
type UserPositionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Position holds the value of the position edge.
	Position *Position `json:"position,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserPositionEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PositionOrErr returns the Position value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserPositionEdges) PositionOrErr() (*Position, error) {
	if e.Position != nil {
		return e.Position, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: position.Label}
	}
	return nil, &NotLoadedError{edge: "position"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserPosition) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userposition.FieldID, userposition.FieldUserID, userposition.FieldPositionID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserPosition fields.
func (up *UserPosition) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userposition.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			up.ID = int(value.Int64)
		case userposition.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				up.UserID = value.Int64
			}
		case userposition.FieldPositionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field position_id", values[i])
			} else if value.Valid {
				up.PositionID = value.Int64
			}
		default:
			up.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserPosition.
// This includes values selected through modifiers, order, etc.
func (up *UserPosition) Value(name string) (ent.Value, error) {
	return up.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserPosition entity.
func (up *UserPosition) QueryUser() *UserQuery {
	return NewUserPositionClient(up.config).QueryUser(up)
}

// QueryPosition queries the "position" edge of the UserPosition entity.
func (up *UserPosition) QueryPosition() *PositionQuery {
	return NewUserPositionClient(up.config).QueryPosition(up)
}

// Update returns a builder for updating this UserPosition.
// Note that you need to call UserPosition.Unwrap() before calling this method if this UserPosition
// was returned from a transaction, and the transaction was committed or rolled back.
func (up *UserPosition) Update() *UserPositionUpdateOne {
	return NewUserPositionClient(up.config).UpdateOne(up)
}

// Unwrap unwraps the UserPosition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (up *UserPosition) Unwrap() *UserPosition {
	_tx, ok := up.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserPosition is not a transactional entity")
	}
	up.config.driver = _tx.drv
	return up
}

// String implements the fmt.Stringer.
func (up *UserPosition) String() string {
	var builder strings.Builder
	builder.WriteString("UserPosition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", up.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", up.UserID))
	builder.WriteString(", ")
	builder.WriteString("position_id=")
	builder.WriteString(fmt.Sprintf("%v", up.PositionID))
	builder.WriteByte(')')
	return builder.String()
}

// UserPositions is a parsable slice of UserPosition.
type UserPositions []*UserPosition
