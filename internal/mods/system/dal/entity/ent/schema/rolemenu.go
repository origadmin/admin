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

// RoleMenu holds the schema definition for the RoleMenu domain.
type RoleMenu struct {
	ent.Schema
}

// Fields of the RoleMenu.
func (RoleMenu) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_id").MaxLen(36), // From Role.ID
		field.String("menu_id").MaxLen(36), // From Menu.ID
	}
}

// Mixin of the RoleMenu.
func (RoleMenu) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the RoleMenu.
func (RoleMenu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id"),
		index.Fields("menu_id"),
	}
}

// Annotations of the Role.
func (RoleMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_role_menus"),
	}
}

// Edges of the RoleMenu.
func (RoleMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).
			Field("role_id").
			// StorageKey(edge.Column("role_id")).
			Required().
			Unique(),
		edge.To("menu", Menu.Type).
			Field("menu_id").
			// StorageKey(edge.Column("menu_id")).
			Required().
			Unique(),
		// edge.From("role", Role.Type).
		//     Ref("role_menus").
		//     Field("role_id"),
		// edge.From("menu", Menu.Type).
		//     Ref("role_menus").
		//     Field("role_id"),
	}
}
