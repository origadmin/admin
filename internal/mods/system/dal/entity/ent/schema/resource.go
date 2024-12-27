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

// Resource holds the schema definition for the Resource domain.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("method").MaxLen(20).Default(""),    // HTTP method (e.g. GET, POST, PUT, DELETE)
		field.String("operation").MaxLen(20).Default(""), // grpc operation (e.g. CreateUser, GetUser, UpdateUser, DeleteUser)
		field.String("path").MaxLen(255),                 // API request path (e.g. /users/:id or /users/{id})
		mixin.OP("menu_id"),                              // From Menu.ID
	}
}

// Mixin of the Resource.
func (Resource) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the Resource.
func (Resource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("menu_id"),
	}
}

// Annotations of the Menu.
func (Resource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_resources"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("resource:table:comment")),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("menu", Menu.Type).
			Ref("resources").
			Field("menu_id").
			Unique(),
		edge.From("permission", Permission.Type).
			Ref("resources").
			Through("permission_resources", PermissionResource.Type),
	}
}
