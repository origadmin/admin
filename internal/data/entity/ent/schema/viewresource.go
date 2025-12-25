package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// ViewResource holds the schema definition for the ViewResource entity.
// It's a through-table for the M-N relationship between View and Resource.
type ViewResource struct {
	ent.Schema
}

// Fields of the ViewResource.
func (ViewResource) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("view_id", i18n.Text("view_resource.view_id.comment")),
		mixin.FK("resource_id", i18n.Text("view_resource.resource_id.comment")),
	}
}

// Edges of the ViewResource.
func (ViewResource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("view", View.Type).
			Field("view_id").
			Unique().
			Required(),
		edge.To("resource", Resource.Type).
			Field("resource_id").
			Unique().
			Required(),
	}
}

// Indexes of the ViewResource.
func (ViewResource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("view_id", "resource_id").
			Unique(),
	}
}

// Annotations of the ViewResource.
func (ViewResource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_view_resources"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.view_resource.table.comment")),
	}
}

// Mixin of the ViewResource.
func (ViewResource) Mixin() []ent.Mixin {
	// Using AuditModelMixin to automatically get id, create/update times, and created_by/updated_by fields.
	return mixin.AuditModelMixin
}
