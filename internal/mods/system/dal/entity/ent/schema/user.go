/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package schema implements the functions, types, and interfaces for the module.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"origadmin/application/admin/helpers/ent/mixin"
	"origadmin/application/admin/helpers/i18n"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/schema/constants"
)

const (
	UserStatusActive = constants.Active
	UserStatusFrozen = constants.Frozen
)

const (
	UserGenderMale    = "male"
	UserGenderFemale  = "female"
	UserGenderUnknown = "unknown"
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
			Comment("user.field.avatar"), // Avatar display avatar of user
		field.String("name").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("entity.user.field.nickname")), // Name of user
		field.Enum("gender").
			Values(UserGenderMale, UserGenderFemale, UserGenderUnknown).
			Default(UserGenderUnknown).
			Comment(i18n.Text("entity.user.field.gender")), // Gender of user
		field.String("password").
			MaxLen(256).
			Default("").
			Comment(i18n.Text("entity.user.field.password")),
		field.String("salt").
			MaxLen(64).
			Default("").
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
			Default(UserStatusActive).
			Comment(i18n.Text("entity.user.field.status")),
		field.Bool("is_system").
			Default(false).
			Comment("user.field.is_system"), // Whether the system is built-in (the built-in user cannot be deleted, but can be disabled)
		field.String("last_login_ip").
			MaxLen(32).
			Default("").
			Comment(i18n.Text("entity.user.field.last_login_ip")),
		mixin.Time("last_login_time", i18n.Text("entity.user.field.last_login_time")),
		mixin.Time("login_time", i18n.Text("entity.user.field.login_time")),
		mixin.TimeOP("sanction_date", i18n.Text("entity.user.field.sanction_date")),
		mixin.OP("manager_id", i18n.Text("entity.user.field.manager_id")), // 管理员ID
		field.String("manager").
			Default("").
			Comment(i18n.Text("entity.user.field.manager")), // 管理员
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return append(mixin.AuditModelMixin, SoftDelete{})
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
		// Roles of user
		edge.To("roles", Role.Type).
			Through("user_roles", UserRole.Type),
		// Posts of user
		// edge.To("posts", Post.Type).
		//	Through("user_posts", UserPost.Type),
		// Departments of user
		edge.To("positions", Position.Type).
			Through("user_positions", UserPosition.Type),
		//// Departments of user
		edge.To("departments", Department.Type).
			Through("user_departments", UserDepartment.Type),
		//edge.To("user_departments", UserDepartment.Type),
	}
}
