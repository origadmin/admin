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
)

const (
	MenuStatusActivated = 0
	MenuStatusFreezed   = 1
)

const (
	MenuTypeAction = 'A'
	MenuTypeButton = 'B'
	MenuTypeLink   = 'L'
	MenuTypeMenu   = 'M'
	MenuTypePage   = 'P'
)

// Menu holds the schema definition for the Menu domain.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword").MaxLen(32).Default(""),       // Code of menu (unique for each level)
		field.String("name").MaxLen(128).Default(""),         // Display name of menu
		field.String("description").MaxLen(1024).Default(""), // Details about menu
		field.Uint8("type").Default(MenuTypePage),            // Dialect of menu (Page,Button,Action,Menu,Link)
		field.String("icon").MaxLen(32).Default(""),
		field.String("path").MaxLen(255).Default(""), // Access path of menu
		field.Int8("status").Default(MenuStatusActivated),
		//field.String("parent_id").MaxLen(36).Default("").Optional(), // Parent UUID (From Menu.UUID)
		mixin.FieldOpID("parent_id"),
		field.String("parent_path").MaxLen(255).Default(""), // Parent path (split by .)
		field.Int("sequence").Default(0),                    // Sequence for sorting (Order by desc)
		field.Text("properties").Default("").Optional(),     // Properties of menu (JSON)
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
		index.Fields("keyword"),
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
		edge.To("resources", Resource.Type),
		edge.From("roles", Role.Type).
			Ref("menus").
			//Field("role_id").
			//Unique().
			Through("role_menus", RoleMenu.Type),
	}
}
