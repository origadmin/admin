// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userdepartment"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userrole"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreateAuthor sets the "create_author" field.
func (uc *UserCreate) SetCreateAuthor(i int64) *UserCreate {
	uc.mutation.SetCreateAuthor(i)
	return uc
}

// SetNillableCreateAuthor sets the "create_author" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateAuthor(i *int64) *UserCreate {
	if i != nil {
		uc.SetCreateAuthor(*i)
	}
	return uc
}

// SetUpdateAuthor sets the "update_author" field.
func (uc *UserCreate) SetUpdateAuthor(i int64) *UserCreate {
	uc.mutation.SetUpdateAuthor(i)
	return uc
}

// SetNillableUpdateAuthor sets the "update_author" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateAuthor(i *int64) *UserCreate {
	if i != nil {
		uc.SetUpdateAuthor(*i)
	}
	return uc
}

// SetCreateTime sets the "create_time" field.
func (uc *UserCreate) SetCreateTime(t time.Time) *UserCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreateTime(*t)
	}
	return uc
}

// SetUpdateTime sets the "update_time" field.
func (uc *UserCreate) SetUpdateTime(t time.Time) *UserCreate {
	uc.mutation.SetUpdateTime(t)
	return uc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdateTime(*t)
	}
	return uc
}

// SetUUID sets the "uuid" field.
func (uc *UserCreate) SetUUID(s string) *UserCreate {
	uc.mutation.SetUUID(s)
	return uc
}

// SetAllowedIP sets the "allowed_ip" field.
func (uc *UserCreate) SetAllowedIP(s string) *UserCreate {
	uc.mutation.SetAllowedIP(s)
	return uc
}

// SetNillableAllowedIP sets the "allowed_ip" field if the given value is not nil.
func (uc *UserCreate) SetNillableAllowedIP(s *string) *UserCreate {
	if s != nil {
		uc.SetAllowedIP(*s)
	}
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetNickname sets the "nickname" field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetGender sets the "gender" field.
func (uc *UserCreate) SetGender(u user.Gender) *UserCreate {
	uc.mutation.SetGender(u)
	return uc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uc *UserCreate) SetNillableGender(u *user.Gender) *UserCreate {
	if u != nil {
		uc.SetGender(*u)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetSalt sets the "salt" field.
func (uc *UserCreate) SetSalt(s string) *UserCreate {
	uc.mutation.SetSalt(s)
	return uc
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (uc *UserCreate) SetNillableSalt(s *string) *UserCreate {
	if s != nil {
		uc.SetSalt(*s)
	}
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetRemark sets the "remark" field.
func (uc *UserCreate) SetRemark(s string) *UserCreate {
	uc.mutation.SetRemark(s)
	return uc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (uc *UserCreate) SetNillableRemark(s *string) *UserCreate {
	if s != nil {
		uc.SetRemark(*s)
	}
	return uc
}

// SetToken sets the "token" field.
func (uc *UserCreate) SetToken(s string) *UserCreate {
	uc.mutation.SetToken(s)
	return uc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (uc *UserCreate) SetNillableToken(s *string) *UserCreate {
	if s != nil {
		uc.SetToken(*s)
	}
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(i int8) *UserCreate {
	uc.mutation.SetStatus(i)
	return uc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(i *int8) *UserCreate {
	if i != nil {
		uc.SetStatus(*i)
	}
	return uc
}

// SetLastLoginIP sets the "last_login_ip" field.
func (uc *UserCreate) SetLastLoginIP(s string) *UserCreate {
	uc.mutation.SetLastLoginIP(s)
	return uc
}

// SetNillableLastLoginIP sets the "last_login_ip" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastLoginIP(s *string) *UserCreate {
	if s != nil {
		uc.SetLastLoginIP(*s)
	}
	return uc
}

// SetLastLoginTime sets the "last_login_time" field.
func (uc *UserCreate) SetLastLoginTime(t time.Time) *UserCreate {
	uc.mutation.SetLastLoginTime(t)
	return uc
}

// SetNillableLastLoginTime sets the "last_login_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastLoginTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastLoginTime(*t)
	}
	return uc
}

// SetSanctionDate sets the "sanction_date" field.
func (uc *UserCreate) SetSanctionDate(t time.Time) *UserCreate {
	uc.mutation.SetSanctionDate(t)
	return uc
}

// SetNillableSanctionDate sets the "sanction_date" field if the given value is not nil.
func (uc *UserCreate) SetNillableSanctionDate(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetSanctionDate(*t)
	}
	return uc
}

// SetManagerID sets the "manager_id" field.
func (uc *UserCreate) SetManagerID(i int64) *UserCreate {
	uc.mutation.SetManagerID(i)
	return uc
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableManagerID(i *int64) *UserCreate {
	if i != nil {
		uc.SetManagerID(*i)
	}
	return uc
}

// SetManager sets the "manager" field.
func (uc *UserCreate) SetManager(s string) *UserCreate {
	uc.mutation.SetManager(s)
	return uc
}

// SetNillableManager sets the "manager" field if the given value is not nil.
func (uc *UserCreate) SetNillableManager(s *string) *UserCreate {
	if s != nil {
		uc.SetManager(*s)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int64) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(i *int64) *UserCreate {
	if i != nil {
		uc.SetID(*i)
	}
	return uc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uc *UserCreate) AddRoleIDs(ids ...int64) *UserCreate {
	uc.mutation.AddRoleIDs(ids...)
	return uc
}

// AddRoles adds the "roles" edges to the Role entity.
func (uc *UserCreate) AddRoles(r ...*Role) *UserCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRoleIDs(ids...)
}

// AddDepartmentIDs adds the "departments" edge to the Department entity by IDs.
func (uc *UserCreate) AddDepartmentIDs(ids ...int64) *UserCreate {
	uc.mutation.AddDepartmentIDs(ids...)
	return uc
}

// AddDepartments adds the "departments" edges to the Department entity.
func (uc *UserCreate) AddDepartments(d ...*Department) *UserCreate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDepartmentIDs(ids...)
}

// AddUserRoleIDs adds the "user_roles" edge to the UserRole entity by IDs.
func (uc *UserCreate) AddUserRoleIDs(ids ...int64) *UserCreate {
	uc.mutation.AddUserRoleIDs(ids...)
	return uc
}

// AddUserRoles adds the "user_roles" edges to the UserRole entity.
func (uc *UserCreate) AddUserRoles(u ...*UserRole) *UserCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserRoleIDs(ids...)
}

// AddUserDepartmentIDs adds the "user_departments" edge to the UserDepartment entity by IDs.
func (uc *UserCreate) AddUserDepartmentIDs(ids ...int64) *UserCreate {
	uc.mutation.AddUserDepartmentIDs(ids...)
	return uc
}

// AddUserDepartments adds the "user_departments" edges to the UserDepartment entity.
func (uc *UserCreate) AddUserDepartments(u ...*UserDepartment) *UserCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserDepartmentIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if err := uc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() error {
	if _, ok := uc.mutation.CreateAuthor(); !ok {
		v := user.DefaultCreateAuthor
		uc.mutation.SetCreateAuthor(v)
	}
	if _, ok := uc.mutation.UpdateAuthor(); !ok {
		v := user.DefaultUpdateAuthor
		uc.mutation.SetUpdateAuthor(v)
	}
	if _, ok := uc.mutation.CreateTime(); !ok {
		if user.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := user.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		if user.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := user.DefaultUpdateTime()
		uc.mutation.SetUpdateTime(v)
	}
	if _, ok := uc.mutation.AllowedIP(); !ok {
		v := user.DefaultAllowedIP
		uc.mutation.SetAllowedIP(v)
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		v := user.DefaultNickname
		uc.mutation.SetNickname(v)
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		v := user.DefaultAvatar
		uc.mutation.SetAvatar(v)
	}
	if _, ok := uc.mutation.Name(); !ok {
		v := user.DefaultName
		uc.mutation.SetName(v)
	}
	if _, ok := uc.mutation.Gender(); !ok {
		v := user.DefaultGender
		uc.mutation.SetGender(v)
	}
	if _, ok := uc.mutation.Password(); !ok {
		v := user.DefaultPassword
		uc.mutation.SetPassword(v)
	}
	if _, ok := uc.mutation.Salt(); !ok {
		v := user.DefaultSalt
		uc.mutation.SetSalt(v)
	}
	if _, ok := uc.mutation.Phone(); !ok {
		v := user.DefaultPhone
		uc.mutation.SetPhone(v)
	}
	if _, ok := uc.mutation.Email(); !ok {
		v := user.DefaultEmail
		uc.mutation.SetEmail(v)
	}
	if _, ok := uc.mutation.Remark(); !ok {
		v := user.DefaultRemark
		uc.mutation.SetRemark(v)
	}
	if _, ok := uc.mutation.Token(); !ok {
		v := user.DefaultToken
		uc.mutation.SetToken(v)
	}
	if _, ok := uc.mutation.Status(); !ok {
		v := user.DefaultStatus
		uc.mutation.SetStatus(v)
	}
	if _, ok := uc.mutation.LastLoginIP(); !ok {
		v := user.DefaultLastLoginIP
		uc.mutation.SetLastLoginIP(v)
	}
	if _, ok := uc.mutation.LastLoginTime(); !ok {
		if user.DefaultLastLoginTime == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultLastLoginTime (forgotten import ent/runtime?)")
		}
		v := user.DefaultLastLoginTime()
		uc.mutation.SetLastLoginTime(v)
	}
	if _, ok := uc.mutation.SanctionDate(); !ok {
		if user.DefaultSanctionDate == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultSanctionDate (forgotten import ent/runtime?)")
		}
		v := user.DefaultSanctionDate()
		uc.mutation.SetSanctionDate(v)
	}
	if _, ok := uc.mutation.Manager(); !ok {
		v := user.DefaultManager
		uc.mutation.SetManager(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		if user.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultID (forgotten import ent/runtime?)")
		}
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "User.create_time"`)}
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "User.update_time"`)}
	}
	if _, ok := uc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "User.uuid"`)}
	}
	if v, ok := uc.mutation.UUID(); ok {
		if err := user.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf(`ent: validator failed for field "User.uuid": %w`, err)}
		}
	}
	if _, ok := uc.mutation.AllowedIP(); !ok {
		return &ValidationError{Name: "allowed_ip", err: errors.New(`ent: missing required field "User.allowed_ip"`)}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if v, ok := uc.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`ent: missing required field "User.nickname"`)}
	}
	if v, ok := uc.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "User.avatar"`)}
	}
	if v, ok := uc.mutation.Avatar(); ok {
		if err := user.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "User.avatar": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "User.gender"`)}
	}
	if v, ok := uc.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "User.gender": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New(`ent: missing required field "User.salt"`)}
	}
	if v, ok := uc.mutation.Salt(); ok {
		if err := user.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "User.salt": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "User.phone"`)}
	}
	if v, ok := uc.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "User.remark"`)}
	}
	if v, ok := uc.mutation.Remark(); ok {
		if err := user.RemarkValidator(v); err != nil {
			return &ValidationError{Name: "remark", err: fmt.Errorf(`ent: validator failed for field "User.remark": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "User.token"`)}
	}
	if v, ok := uc.mutation.Token(); ok {
		if err := user.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "User.token": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "User.status"`)}
	}
	if _, ok := uc.mutation.LastLoginIP(); !ok {
		return &ValidationError{Name: "last_login_ip", err: errors.New(`ent: missing required field "User.last_login_ip"`)}
	}
	if v, ok := uc.mutation.LastLoginIP(); ok {
		if err := user.LastLoginIPValidator(v); err != nil {
			return &ValidationError{Name: "last_login_ip", err: fmt.Errorf(`ent: validator failed for field "User.last_login_ip": %w`, err)}
		}
	}
	if _, ok := uc.mutation.LastLoginTime(); !ok {
		return &ValidationError{Name: "last_login_time", err: errors.New(`ent: missing required field "User.last_login_time"`)}
	}
	if _, ok := uc.mutation.SanctionDate(); !ok {
		return &ValidationError{Name: "sanction_date", err: errors.New(`ent: missing required field "User.sanction_date"`)}
	}
	if v, ok := uc.mutation.ManagerID(); ok {
		if err := user.ManagerIDValidator(v); err != nil {
			return &ValidationError{Name: "manager_id", err: fmt.Errorf(`ent: validator failed for field "User.manager_id": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Manager(); !ok {
		return &ValidationError{Name: "manager", err: errors.New(`ent: missing required field "User.manager"`)}
	}
	if v, ok := uc.mutation.ID(); ok {
		if err := user.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "User.id": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreateAuthor(); ok {
		_spec.SetField(user.FieldCreateAuthor, field.TypeInt64, value)
		_node.CreateAuthor = value
	}
	if value, ok := uc.mutation.UpdateAuthor(); ok {
		_spec.SetField(user.FieldUpdateAuthor, field.TypeInt64, value)
		_node.UpdateAuthor = value
	}
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := uc.mutation.UUID(); ok {
		_spec.SetField(user.FieldUUID, field.TypeString, value)
		_node.UUID = value
	}
	if value, ok := uc.mutation.AllowedIP(); ok {
		_spec.SetField(user.FieldAllowedIP, field.TypeString, value)
		_node.AllowedIP = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeEnum, value)
		_node.Gender = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Salt(); ok {
		_spec.SetField(user.FieldSalt, field.TypeString, value)
		_node.Salt = value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Remark(); ok {
		_spec.SetField(user.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := uc.mutation.Token(); ok {
		_spec.SetField(user.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.LastLoginIP(); ok {
		_spec.SetField(user.FieldLastLoginIP, field.TypeString, value)
		_node.LastLoginIP = value
	}
	if value, ok := uc.mutation.LastLoginTime(); ok {
		_spec.SetField(user.FieldLastLoginTime, field.TypeTime, value)
		_node.LastLoginTime = value
	}
	if value, ok := uc.mutation.SanctionDate(); ok {
		_spec.SetField(user.FieldSanctionDate, field.TypeTime, value)
		_node.SanctionDate = value
	}
	if value, ok := uc.mutation.ManagerID(); ok {
		_spec.SetField(user.FieldManagerID, field.TypeInt64, value)
		_node.ManagerID = value
	}
	if value, ok := uc.mutation.Manager(); ok {
		_spec.SetField(user.FieldManager, field.TypeString, value)
		_node.Manager = value
	}
	if nodes := uc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserRoleCreate{config: uc.config, mutation: newUserRoleMutation(uc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.DepartmentsTable,
			Columns: user.DepartmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserDepartmentCreate{config: uc.config, mutation: newUserDepartmentMutation(uc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.UserRolesTable,
			Columns: []string{user.UserRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userrole.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserDepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.UserDepartmentsTable,
			Columns: []string{user.UserDepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userdepartment.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetUser set the User
func (uc *UserCreate) SetUser(input *User, fields ...string) *UserCreate {
	m := uc.mutation
	if len(fields) == 0 {
		fields = user.Columns
	}
	_ = m.SetFields(input, fields...)
	return uc
}

// SetUserWithZero set the User
func (uc *UserCreate) SetUserWithZero(input *User, fields ...string) *UserCreate {
	m := uc.mutation
	if len(fields) == 0 {
		fields = user.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return uc
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
