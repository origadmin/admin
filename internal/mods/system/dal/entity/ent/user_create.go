// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
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
func (uc *UserCreate) SetCreateAuthor(s string) *UserCreate {
	uc.mutation.SetCreateAuthor(s)
	return uc
}

// SetNillableCreateAuthor sets the "create_author" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateAuthor(s *string) *UserCreate {
	if s != nil {
		uc.SetCreateAuthor(*s)
	}
	return uc
}

// SetUpdateAuthor sets the "update_author" field.
func (uc *UserCreate) SetUpdateAuthor(s string) *UserCreate {
	uc.mutation.SetUpdateAuthor(s)
	return uc
}

// SetNillableUpdateAuthor sets the "update_author" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateAuthor(s *string) *UserCreate {
	if s != nil {
		uc.SetUpdateAuthor(*s)
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

// SetIndex sets the "index" field.
func (uc *UserCreate) SetIndex(i int) *UserCreate {
	uc.mutation.SetIndex(i)
	return uc
}

// SetDepartment sets the "department" field.
func (uc *UserCreate) SetDepartment(s string) *UserCreate {
	uc.mutation.SetDepartment(s)
	return uc
}

// SetNillableDepartment sets the "department" field if the given value is not nil.
func (uc *UserCreate) SetNillableDepartment(s *string) *UserCreate {
	if s != nil {
		uc.SetDepartment(*s)
	}
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

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uc *UserCreate) SetNillableUsername(s *string) *UserCreate {
	if s != nil {
		uc.SetUsername(*s)
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

// SetUserID sets the "user_id" field.
func (uc *UserCreate) SetUserID(s string) *UserCreate {
	uc.mutation.SetUserID(s)
	return uc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableUserID(s *string) *UserCreate {
	if s != nil {
		uc.SetUserID(*s)
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

// SetManagerID sets the "manager_id" field.
func (uc *UserCreate) SetManagerID(s string) *UserCreate {
	uc.mutation.SetManagerID(s)
	return uc
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableManagerID(s *string) *UserCreate {
	if s != nil {
		uc.SetManagerID(*s)
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
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uc *UserCreate) AddRoleIDs(ids ...string) *UserCreate {
	uc.mutation.AddRoleIDs(ids...)
	return uc
}

// AddRoles adds the "roles" edges to the Role entity.
func (uc *UserCreate) AddRoles(r ...*Role) *UserCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRoleIDs(ids...)
}

// AddUserRoleIDs adds the "user_roles" edge to the UserRole entity by IDs.
func (uc *UserCreate) AddUserRoleIDs(ids ...string) *UserCreate {
	uc.mutation.AddUserRoleIDs(ids...)
	return uc
}

// AddUserRoles adds the "user_roles" edges to the UserRole entity.
func (uc *UserCreate) AddUserRoles(u ...*UserRole) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
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
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreateAuthor(); !ok {
		v := user.DefaultCreateAuthor
		uc.mutation.SetCreateAuthor(v)
	}
	if _, ok := uc.mutation.UpdateAuthor(); !ok {
		v := user.DefaultUpdateAuthor
		uc.mutation.SetUpdateAuthor(v)
	}
	if _, ok := uc.mutation.CreateTime(); !ok {
		v := user.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		v := user.DefaultUpdateTime()
		uc.mutation.SetUpdateTime(v)
	}
	if _, ok := uc.mutation.Department(); !ok {
		v := user.DefaultDepartment
		uc.mutation.SetDepartment(v)
	}
	if _, ok := uc.mutation.AllowedIP(); !ok {
		v := user.DefaultAllowedIP
		uc.mutation.SetAllowedIP(v)
	}
	if _, ok := uc.mutation.Username(); !ok {
		v := user.DefaultUsername
		uc.mutation.SetUsername(v)
	}
	if _, ok := uc.mutation.Name(); !ok {
		v := user.DefaultName
		uc.mutation.SetName(v)
	}
	if _, ok := uc.mutation.UserID(); !ok {
		v := user.DefaultUserID
		uc.mutation.SetUserID(v)
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		v := user.DefaultAvatar
		uc.mutation.SetAvatar(v)
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
	if _, ok := uc.mutation.LastLoginTime(); !ok {
		v := user.DefaultLastLoginTime()
		uc.mutation.SetLastLoginTime(v)
	}
	if _, ok := uc.mutation.SanctionDate(); !ok {
		v := user.DefaultSanctionDate()
		uc.mutation.SetSanctionDate(v)
	}
	if _, ok := uc.mutation.Status(); !ok {
		v := user.DefaultStatus
		uc.mutation.SetStatus(v)
	}
	if _, ok := uc.mutation.Manager(); !ok {
		v := user.DefaultManager
		uc.mutation.SetManager(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreateAuthor(); !ok {
		return &ValidationError{Name: "create_author", err: errors.New(`ent: missing required field "User.create_author"`)}
	}
	if _, ok := uc.mutation.UpdateAuthor(); !ok {
		return &ValidationError{Name: "update_author", err: errors.New(`ent: missing required field "User.update_author"`)}
	}
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "User.create_time"`)}
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "User.update_time"`)}
	}
	if _, ok := uc.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New(`ent: missing required field "User.index"`)}
	}
	if _, ok := uc.mutation.Department(); !ok {
		return &ValidationError{Name: "department", err: errors.New(`ent: missing required field "User.department"`)}
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
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "User.user_id"`)}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "User.avatar"`)}
	}
	if v, ok := uc.mutation.Avatar(); ok {
		if err := user.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "User.avatar": %w`, err)}
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
	if _, ok := uc.mutation.LastLoginTime(); !ok {
		return &ValidationError{Name: "last_login_time", err: errors.New(`ent: missing required field "User.last_login_time"`)}
	}
	if _, ok := uc.mutation.SanctionDate(); !ok {
		return &ValidationError{Name: "sanction_date", err: errors.New(`ent: missing required field "User.sanction_date"`)}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "User.status"`)}
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
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreateAuthor(); ok {
		_spec.SetField(user.FieldCreateAuthor, field.TypeString, value)
		_node.CreateAuthor = value
	}
	if value, ok := uc.mutation.UpdateAuthor(); ok {
		_spec.SetField(user.FieldUpdateAuthor, field.TypeString, value)
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
	if value, ok := uc.mutation.Index(); ok {
		_spec.SetField(user.FieldIndex, field.TypeInt, value)
		_node.Index = value
	}
	if value, ok := uc.mutation.Department(); ok {
		_spec.SetField(user.FieldDepartment, field.TypeString, value)
		_node.Department = value
	}
	if value, ok := uc.mutation.AllowedIP(); ok {
		_spec.SetField(user.FieldAllowedIP, field.TypeString, value)
		_node.AllowedIP = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.UserID(); ok {
		_spec.SetField(user.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
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
	if value, ok := uc.mutation.LastLoginTime(); ok {
		_spec.SetField(user.FieldLastLoginTime, field.TypeTime, value)
		_node.LastLoginTime = value
	}
	if value, ok := uc.mutation.SanctionDate(); ok {
		_spec.SetField(user.FieldSanctionDate, field.TypeTime, value)
		_node.SanctionDate = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.ManagerID(); ok {
		_spec.SetField(user.FieldManagerID, field.TypeString, value)
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
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserRoleCreate{config: uc.config, mutation: newUserRoleMutation(uc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
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
				IDSpec: sqlgraph.NewFieldSpec(userrole.FieldID, field.TypeString),
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
	_ = m.SetFields(input, fields...)
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
