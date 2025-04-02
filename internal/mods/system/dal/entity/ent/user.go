// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// entity.user.table.comment
type User struct {
	config `json:"-"`
	// ID of the ent.
	// field.primary_key.comment
	ID int64 `json:"id,omitempty"`
	// create_author.field.comment
	CreateAuthor int64 `json:"create_author,omitempty"`
	// update_author.field.comment
	UpdateAuthor int64 `json:"update_author,omitempty"`
	// create_time.field.comment
	CreateTime time.Time `json:"create_time,omitempty"`
	// update_time.field.comment
	UpdateTime time.Time `json:"update_time,omitempty"`
	// delete_time.field.comment
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// entity.user.field.uuid
	UUID string `json:"uuid,omitempty"`
	// entity.user.field.allowed_ip
	AllowedIP string `json:"allowed_ip,omitempty"`
	// entity.user.field.username
	Username string `json:"username,omitempty"`
	// entity.user.field.nickname
	Nickname string `json:"nickname,omitempty"`
	// user.field.avatar
	Avatar string `json:"avatar,omitempty"`
	// entity.user.field.nickname
	Name string `json:"name,omitempty"`
	// entity.user.field.gender
	Gender user.Gender `json:"gender,omitempty"`
	// entity.user.field.password
	Password string `json:"password,omitempty"`
	// entity.user.field.salt
	Salt string `json:"salt,omitempty"`
	// entity.user.field.phone
	Phone string `json:"phone,omitempty"`
	// entity.user.field.email
	Email string `json:"email,omitempty"`
	// entity.user.field.department
	Department string `json:"department,omitempty"`
	// entity.user.field.remark
	Remark string `json:"remark,omitempty"`
	// entity.user.field.token
	Token string `json:"token,omitempty"`
	// entity.user.field.status
	Status int8 `json:"status,omitempty"`
	// user.field.is_system
	IsSystem bool `json:"is_system,omitempty"`
	// entity.user.field.last_login_ip
	LastLoginIP string `json:"last_login_ip,omitempty"`
	// entity.user.field.last_login_time
	LastLoginTime time.Time `json:"last_login_time,omitempty"`
	// entity.user.field.login_time
	LoginTime time.Time `json:"login_time,omitempty"`
	// entity.user.field.sanction_date
	SanctionDate time.Time `json:"sanction_date,omitempty"`
	// entity.user.field.manager_id
	ManagerID int64 `json:"manager_id,omitempty"`
	// entity.user.field.manager
	Manager string `json:"manager,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Positions holds the value of the positions edge.
	Positions []*Position `json:"positions,omitempty"`
	// Departments holds the value of the departments edge.
	Departments []*Department `json:"departments,omitempty"`
	// UserRoles holds the value of the user_roles edge.
	UserRoles []*UserRole `json:"user_roles,omitempty"`
	// UserPositions holds the value of the user_positions edge.
	UserPositions []*UserPosition `json:"user_positions,omitempty"`
	// UserDepartments holds the value of the user_departments edge.
	UserDepartments []*UserDepartment `json:"user_departments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// PositionsOrErr returns the Positions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PositionsOrErr() ([]*Position, error) {
	if e.loadedTypes[1] {
		return e.Positions, nil
	}
	return nil, &NotLoadedError{edge: "positions"}
}

// DepartmentsOrErr returns the Departments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DepartmentsOrErr() ([]*Department, error) {
	if e.loadedTypes[2] {
		return e.Departments, nil
	}
	return nil, &NotLoadedError{edge: "departments"}
}

// UserRolesOrErr returns the UserRoles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserRolesOrErr() ([]*UserRole, error) {
	if e.loadedTypes[3] {
		return e.UserRoles, nil
	}
	return nil, &NotLoadedError{edge: "user_roles"}
}

// UserPositionsOrErr returns the UserPositions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserPositionsOrErr() ([]*UserPosition, error) {
	if e.loadedTypes[4] {
		return e.UserPositions, nil
	}
	return nil, &NotLoadedError{edge: "user_positions"}
}

// UserDepartmentsOrErr returns the UserDepartments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserDepartmentsOrErr() ([]*UserDepartment, error) {
	if e.loadedTypes[5] {
		return e.UserDepartments, nil
	}
	return nil, &NotLoadedError{edge: "user_departments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsSystem:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldCreateAuthor, user.FieldUpdateAuthor, user.FieldStatus, user.FieldManagerID:
			values[i] = new(sql.NullInt64)
		case user.FieldUUID, user.FieldAllowedIP, user.FieldUsername, user.FieldNickname, user.FieldAvatar, user.FieldName, user.FieldGender, user.FieldPassword, user.FieldSalt, user.FieldPhone, user.FieldEmail, user.FieldDepartment, user.FieldRemark, user.FieldToken, user.FieldLastLoginIP, user.FieldManager:
			values[i] = new(sql.NullString)
		case user.FieldCreateTime, user.FieldUpdateTime, user.FieldDeleteTime, user.FieldLastLoginTime, user.FieldLoginTime, user.FieldSanctionDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int64(value.Int64)
		case user.FieldCreateAuthor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_author", values[i])
			} else if value.Valid {
				u.CreateAuthor = value.Int64
			}
		case user.FieldUpdateAuthor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_author", values[i])
			} else if value.Valid {
				u.UpdateAuthor = value.Int64
			}
		case user.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				u.CreateTime = value.Time
			}
		case user.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				u.UpdateTime = value.Time
			}
		case user.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				u.DeleteTime = new(time.Time)
				*u.DeleteTime = value.Time
			}
		case user.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				u.UUID = value.String
			}
		case user.FieldAllowedIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field allowed_ip", values[i])
			} else if value.Valid {
				u.AllowedIP = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = user.Gender(value.String)
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				u.Salt = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldDepartment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field department", values[i])
			} else if value.Valid {
				u.Department = value.String
			}
		case user.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				u.Remark = value.String
			}
		case user.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				u.Token = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = int8(value.Int64)
			}
		case user.FieldIsSystem:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_system", values[i])
			} else if value.Valid {
				u.IsSystem = value.Bool
			}
		case user.FieldLastLoginIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_ip", values[i])
			} else if value.Valid {
				u.LastLoginIP = value.String
			}
		case user.FieldLastLoginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_time", values[i])
			} else if value.Valid {
				u.LastLoginTime = value.Time
			}
		case user.FieldLoginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field login_time", values[i])
			} else if value.Valid {
				u.LoginTime = value.Time
			}
		case user.FieldSanctionDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sanction_date", values[i])
			} else if value.Valid {
				u.SanctionDate = value.Time
			}
		case user.FieldManagerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field manager_id", values[i])
			} else if value.Valid {
				u.ManagerID = value.Int64
			}
		case user.FieldManager:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manager", values[i])
			} else if value.Valid {
				u.Manager = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the User entity.
func (u *User) QueryRoles() *RoleQuery {
	return NewUserClient(u.config).QueryRoles(u)
}

// QueryPositions queries the "positions" edge of the User entity.
func (u *User) QueryPositions() *PositionQuery {
	return NewUserClient(u.config).QueryPositions(u)
}

// QueryDepartments queries the "departments" edge of the User entity.
func (u *User) QueryDepartments() *DepartmentQuery {
	return NewUserClient(u.config).QueryDepartments(u)
}

// QueryUserRoles queries the "user_roles" edge of the User entity.
func (u *User) QueryUserRoles() *UserRoleQuery {
	return NewUserClient(u.config).QueryUserRoles(u)
}

// QueryUserPositions queries the "user_positions" edge of the User entity.
func (u *User) QueryUserPositions() *UserPositionQuery {
	return NewUserClient(u.config).QueryUserPositions(u)
}

// QueryUserDepartments queries the "user_departments" edge of the User entity.
func (u *User) QueryUserDepartments() *UserDepartmentQuery {
	return NewUserClient(u.config).QueryUserDepartments(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("create_author=")
	builder.WriteString(fmt.Sprintf("%v", u.CreateAuthor))
	builder.WriteString(", ")
	builder.WriteString("update_author=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdateAuthor))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(u.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(u.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := u.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(u.UUID)
	builder.WriteString(", ")
	builder.WriteString("allowed_ip=")
	builder.WriteString(u.AllowedIP)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", u.Gender))
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("salt=")
	builder.WriteString(u.Salt)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("department=")
	builder.WriteString(u.Department)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(u.Remark)
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(u.Token)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("is_system=")
	builder.WriteString(fmt.Sprintf("%v", u.IsSystem))
	builder.WriteString(", ")
	builder.WriteString("last_login_ip=")
	builder.WriteString(u.LastLoginIP)
	builder.WriteString(", ")
	builder.WriteString("last_login_time=")
	builder.WriteString(u.LastLoginTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("login_time=")
	builder.WriteString(u.LoginTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sanction_date=")
	builder.WriteString(u.SanctionDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("manager_id=")
	builder.WriteString(fmt.Sprintf("%v", u.ManagerID))
	builder.WriteString(", ")
	builder.WriteString("manager=")
	builder.WriteString(u.Manager)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
