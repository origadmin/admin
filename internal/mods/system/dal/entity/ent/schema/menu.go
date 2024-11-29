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

// Menu holds the schema definition for the Menu domain.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").MaxLen(32).Default(""),          // Code of menu (unique for each level)
		field.String("name").MaxLen(128).Default(""),         // Display name of menu
		field.String("description").MaxLen(1024).Default(""), // Details about menu
		field.String("type").MaxLen(20).Default(""),          // Dialect of menu (page, button)
		field.String("path").MaxLen(255).Default(""),         // Access path of menu
		field.Int8("status").Default(0),
		field.String("parent_id").MaxLen(36).Optional(),     // Parent ID (From Menu.ID)
		field.String("parent_path").MaxLen(255).Default(""), // Parent path (split by .)
		field.Int("sequence"),                               // Sequence for sorting (Order by desc)
		field.Text("properties"),                            // Properties of menu (JSON)
		// Resources of menu
	}
}

// Mixin of the Menu.
func (Menu) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the Menu.
func (Menu) Indexes() []ent.Index {
	return []ent.Index{
		// Unique index
		index.Fields("code"),
		index.Fields("name"),
		index.Fields("type"),
		index.Fields("status"),
		index.Fields("parent_id"),
		index.Fields("parent_path"),
	}
}

// Annotations of the Menu.
func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_menus"),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		// Child menus
		edge.To("children", Menu.Type),
		// Parent menus
		edge.From("parent", Menu.Type).
			Ref("children").
			Field("parent_id").
			Unique(),
		edge.To("resources", MenuResource.Type),
		edge.From("roles", Role.Type).
			Ref("menus").
			// Field("role_id").
			// Unique().
			Through("role_menus", RoleMenu.Type),
	}
}
