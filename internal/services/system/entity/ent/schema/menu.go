// Copyright (c) 2024 KasaAdmin. All rights reserved.

package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

const (
	MenuStatusDisabled = "disabled"
	MenuStatusEnabled  = "enabled"
)

// Menu holds the schema definition for the Menu domain.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Annotations(entproto.Field(1)).MaxLen(36).Default(""),            // ID of menu
		field.String("code").Annotations(entproto.Field(2)).MaxLen(32).Default(""),          // Code of menu (unique for each level)
		field.String("name").Annotations(entproto.Field(3)).MaxLen(128).Default(""),         // Display name of menu
		field.String("description").Annotations(entproto.Field(4)).MaxLen(1024).Default(""), // Details about menu
		field.Int("sequence").Annotations(entproto.Field(5)),                                // Sequence for sorting (Order by desc)
		field.String("type").Annotations(entproto.Field(6)).MaxLen(20).Default(""),          // Dialect of menu (page, button)
		field.String("path").Annotations(entproto.Field(7)).MaxLen(255).Default(""),         // Access path of menu
		field.Text("properties").Annotations(entproto.Field(8)),                             // Properties of menu (JSON)
		field.Enum("status").Annotations(entproto.Field(9), entproto.Enum(map[string]int32{
			MenuStatusDisabled: 0,
			MenuStatusEnabled:  1,
		})).
			Values(MenuStatusDisabled, MenuStatusEnabled).
			Default(MenuStatusDisabled), // Status of menu (enabled, disabled)
		field.String("parent_id").Annotations(entproto.Field(10)).MaxLen(36).Optional(),     // Parent ID (From Menu.ID)
		field.String("parent_path").Annotations(entproto.Field(11)).MaxLen(255).Default(""), // Parent path (split by .)
		// Resources of menu
	}
}

// Mixin of the Menu.
func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		//mixin.Model{},
		//mixin.CreatedSchema{},
		//mixin.UpdatedSchema{},
	}
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
		entproto.Message(),
		entproto.Service(), // <-- add this
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		// Child menus
		edge.To("children", Menu.Type).Annotations(entproto.Field(12)),
		// Parent menus
		edge.From("parent", Menu.Type).Annotations(entproto.Field(13)).
			Ref("children").
			Field("parent_id").
			Unique(),
		//edge.To("resources", MenuResource.Type),
		//edge.From("roles", Role.Type).
		//	Ref("menus").
		//	// Field("role_id").
		//	// Unique().
		//	Through("role_menus", RoleMenu.Type),
	}
}
