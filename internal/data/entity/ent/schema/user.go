/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package schema implements the functions, types, and interfaces for the module.
package schema

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	gen "origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/hook"
	"origadmin/application/admin/internal/data/entity/ent/user"

	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// User holds the schema definition for the User domain.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		mixin.FieldUUIDFK("uuid", i18n.Text("entity.user.field.uuid")),
		field.String("allowed_ip").
			Default("0.0.0.0").
			Comment(i18n.Text("entity.user.field.allowed_ip")),
		field.String("username").
			MaxLen(32).
			Unique().
			Comment(i18n.Text("entity.user.field.username")), // login username of user
		field.String("nickname").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("entity.user.field.nickname")), // Nickname display name of user
		field.String("avatar").
			MaxLen(256).
			Default("").
			Comment("entity.user.field.avatar"), // Avatar display avatar of user
		field.String("name").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("entity.user.field.nickname")), // Name of user
		field.Enum("gender").
			Values(
				string(enums.GenderMale),
				string(enums.GenderFemale),
				string(enums.GenderUnknown),
			).
			Default(string(enums.GenderUnknown)).
			Comment(i18n.Text("entity.user.field.gender")), // Gender of user
		field.String("encrypted_password").
			MaxLen(256).
			Default("").
			Comment(i18n.Text("entity.user.field.encrypted_password")),
		field.String("salt").
			MaxLen(64).
			Default("").
			Deprecated("toolkits/crypto includes salt management").
			Comment(i18n.Text("entity.user.field.salt")),
		field.String("phone").
			MaxLen(32).
			Default("").
			Comment(i18n.Text("entity.user.field.phone")), // login phone number of user
		field.String("email").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("entity.user.field.email")), // login email of user
		field.String("department").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("entity.user.field.department")), // Department of user
		field.String("remark").
			MaxLen(1024).
			Default("").
			Comment(i18n.Text("entity.user.field.remark")), // Remark of user
		field.String("token").
			MaxLen(512).
			Default("").
			Comment(i18n.Text("entity.user.field.token")), // Token for login
		field.Int8("status").
			GoType(enums.Status(0)).
			Default(int8(enums.StatusActive)).
			Comment(i18n.Text("entity.user.field.status")),
		field.Bool("is_system").
			Default(false).
			Comment("entity.user.field.is_system"), // Whether the system is built-in (the built-in user cannot be deleted, but can be disabled)
		field.String("last_login_ip").
			MaxLen(32).
			Default("").
			Comment(i18n.Text("entity.user.field.last_login_ip")),
		field.String("login_ip").
			MaxLen(32).
			Default("").
			Comment(i18n.Text("entity.user.field.login_ip")),
		mixin.Time("last_login_time", i18n.Text("entity.user.field.last_login_time")),
		mixin.Time("login_time", i18n.Text("entity.user.field.login_time")),
		mixin.TimeOptional("sanction_date", i18n.Text("entity.user.field.sanction_date")),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return append(mixin.AuditModelMixin, mixin.SoftDeleteMixin{})
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("phone"),
		index.Fields("email"),
		index.Fields("status"),
	}
}

// Annotations of the Role.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_users"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.user.table.comment")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).
			Through("user_roles", UserRole.Type),
		edge.To("positions", Position.Type).
			Through("user_positions", UserPosition.Type),
		edge.To("departments", Department.Type).
			Through("user_departments", UserDepartment.Type),
	}
}

// Interceptors of the User.
func (User) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		mixin.SoftDeleteInterceptor(mixin.SoftDeleteMixin{}),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		// On CREATE, prevent creating more than one system user.
		hook.On(preventDuplicateSystemUser, ent.OpCreate),
		// On DELETE, prevent system users from being deleted.
		hook.On(preventDeleteSystemUser, ent.OpDelete|ent.OpDeleteOne),
		// On UPDATE, convert DELETE operations to UPDATE operations.
		mixin.SoftDeleteHook(mixin.SoftDeleteMixin{}),
	}
}

// preventDuplicateSystemUser is a hook that prevents creating more than one system user.
func preventDuplicateSystemUser(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
		isSystem, ok := m.IsSystem()
		if !ok || !isSystem {
			return next.Mutate(ctx, m)
		}
		// If creating a system user, check if one already exists.
		count, err := m.Client().User.
			Query().
			Where(user.IsSystem(true)).
			Count(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to check for existing system user: %w", err)
		}
		if count > 0 {
			return nil, errors.New("only one system user is allowed")
		}
		return next.Mutate(ctx, m)
	})
}

// preventDeleteSystemUser is a hook that prevents deleting a system user.
func preventDeleteSystemUser(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
		// Add a predicate to ensure system users are not included in the delete operation.
		m.Where(user.IsSystem(false))
		return next.Mutate(ctx, m)
	})
}
