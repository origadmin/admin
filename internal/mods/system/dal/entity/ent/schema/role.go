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

// Role holds the schema definition for the Role domain.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").MaxLen(32).Default(""),          // Code of role (unique)
		field.String("name").MaxLen(128).Default(""),         // Display name of role
		field.String("description").MaxLen(1024).Default(""), // Details about role
		field.Int("sequence"),                                // Sequence for sorting
		field.Int8("status").Default(0),
		// Role menu list
	}
}

// Mixin of the Role.
func (Role) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the Role.
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
		index.Fields("name"),
		index.Fields("sequence"),
		index.Fields("status"),
	}
}

// Annotations of the Role.
func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_roles"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("menus", Menu.Type).
			StorageKey(edge.Columns("role_id", "menu_id")).
			Through("role_menus", RoleMenu.Type),
		// edge.To("menus", Menu.Dialect).StorageKey(edge.Column("menu_id")),
		// edge.To("users", User.Type).
		//     StorageKey(edge.Columns("user_id", "role_id")).
		//     Through("user_role", UserRole.Type), // .Field("user_id"),
		edge.From("users", User.Type).
			Ref("roles").
			Through("user_roles", UserRole.Type), // .Field("user_id"),
		// edge.From("users", User.Dialect).Ref("roles").StorageKey(edge.Table(config.c.FormatTableName("user_role")), edge.Column("role_id")).Unique(),
	}
}
