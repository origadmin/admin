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
)

const (
	UserStatusActivated = 0
	UserStatusFreezed   = 1
)

// User holds the schema definition for the User domain.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		mixin.FieldUUIDFK("uuid"),
		field.String("allowed_ip").Default("0.0.0.0"),
		field.String("username").MaxLen(32).Default(""), // login username of user
		field.String("nickname").MaxLen(64).Default(""), // Nickname display name of user
		field.String("avatar").MaxLen(256).Default(""),  // Avatar display avatar of user
		field.String("name").MaxLen(64).Default(""),     // Name of user
		field.String("gender").MaxLen(16).Default(""),   // Gender of user
		field.String("password").
			MaxLen(256).
			Default("").
			Sensitive(),
		field.String("salt").
			MaxLen(64).
			Default("").
			Sensitive(),
		field.String("phone").MaxLen(32).Default(""),    // login phone number of user
		field.String("email").MaxLen(64).Default(""),    // login email of user
		field.String("remark").MaxLen(1024).Default(""), // Remark of user
		field.String("token").MaxLen(512).Default(""),   // Token for login
		field.Int8("status").Default(UserStatusActivated),
		field.String("last_login_ip").MaxLen(32).Default(""),
		mixin.Time("last_login_time", i18n.Text("user:field:last_login_time")),
		mixin.Time("sanction_date", i18n.Text("user:field:sanction_date")),
		mixin.FK("manager_id"),              // 管理员ID
		field.String("manager").Default(""), // 管理员
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return mixin.AuditModelMixin
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
		schema.Comment(i18n.Text("user:table:comment")),
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
		//edge.To("user_departments", UserDepartment.Type).
		//	From("user").
		//	Ref("user"),
		//// Departments of user
		edge.To("departments", Department.Type).
			Through("user_departments", UserDepartment.Type),
		//edge.To("user_departments", UserDepartment.Type),
	}
}
