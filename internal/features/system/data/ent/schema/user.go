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
	"entgo.io/ent/schema/mixin"

	"origadmin/application/admin/internal/data/entity/ent/schema/types"
)

const (
	UserStatusActive = types.Active
	UserStatusFrozen = types.Frozen
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

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_users"),
		entsql.WithComments(true),
		schema.Comment("User table"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("ID").
			Immutable().
			Unique(),
		field.String("uuid").MaxLen(36).Unique().Comment("UUID"),
		field.String("allowed_ip").Default("0.0.0.0").Comment("Allowed IP"),
		field.String("username").MaxLen(32).Unique().Comment("login username of user"),
		field.String("nickname").MaxLen(64).Default("").Comment("Nickname display name of user"),
		field.String("avatar").MaxLen(256).Default("").Comment("Avatar display avatar of user"),
		field.String("name").MaxLen(64).Default("").Comment("Name of user"),
		field.Enum("gender").Values(UserGenderMale, UserGenderFemale, UserGenderUnknown).Default(UserGenderUnknown).Comment("Gender of user"),
		field.String("password").MaxLen(256).Default("").Sensitive().Comment("Encrypted password"),
		field.String("phone").MaxLen(32).Default("").Comment("login phone number of user"),
		field.String("email").MaxLen(64).Default("").Comment("login email of user"),
		field.String("department").MaxLen(64).Default("").Comment("Department of user"),
		field.String("remark").MaxLen(1024).Default("").Comment("Remark of user"),
		field.Int8("status").Default(UserStatusActive).Comment("status"),
		field.Bool("is_system").Default(false).Comment("Whether the system is built-in"),
		field.String("last_login_ip").MaxLen(32).Default("").Comment("Last login IP"),
		field.Time("last_login_time").Optional().Comment("Last login time"),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).
			Through("user_roles", UserRole.Type),
	}
}
