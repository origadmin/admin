/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"origadmin/application/admin/helpers/ent/mixin"
)

const (
	// UserStatusActivated represents the activated status of a user
	UserStatusActivated = 0
	// UserStatusFreezed represents the freezed status of a user
	UserStatusFreezed = 1
)

// User holds the schema definition for the User domain.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		mixin.FieldIndex("index"),
		field.String("department").Default(""),
		field.String("allowed_ip").Default("0.0.0.0"),
		field.String("username").MaxLen(64).Default(""),  // Username for login
		field.String("name").MaxLen(64).Default(""),      // Name of user
		field.String("user_id").Default(""),              // User ID
		field.String("avatar").MaxLen(255).Default(""),   // Avatar of user
		field.String("password").MaxLen(128).Default(""), // Password for login (encrypted)
		field.String("salt").MaxLen(255).Default(""),     // Salt
		field.String("phone").MaxLen(32).Default(""),     // Phone number of user
		field.String("email").MaxLen(128).Default(""),    // Email of user
		field.String("remark").MaxLen(1024).Default(""),  // Remark of user
		mixin.FieldTime("last_login_time"),
		mixin.FieldTime("sanction_date"),
		field.Int8("status").Default(UserStatusActivated),
		mixin.FieldID("manager_id"),         // 管理员ID
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
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// Roles of user
		edge.To("roles", Role.Type).
			// Ref("users").
			// Field("role_id").
			// Unique().
			Through("user_roles", UserRole.Type),
	}
}
