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

// MenuResource holds the schema definition for the MenuResource domain.
type MenuResource struct {
	ent.Schema
}

// Fields of the MenuResource.
func (MenuResource) Fields() []ent.Field {
	return []ent.Field{
		field.String("menu_id").MaxLen(36).Optional(), // From Menu.ID
		field.String("method").MaxLen(20).Default(""), // HTTP method
		field.String("path").MaxLen(255),              // API request path (e.g. /users/:id or /users/{id})
	}
}

// Mixin of the MenuResource.
func (MenuResource) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the MenuResource.
func (MenuResource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("menu_id"),
	}
}

// Annotations of the Menu.
func (MenuResource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_menu_resources"),
	}
}

// Edges of the MenuResource.
func (MenuResource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("menu", Menu.Type).Ref("resources").Field("menu_id").Unique(),
	}
}
