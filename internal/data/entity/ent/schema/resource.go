package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("service_name").
			Comment(i18n.Text("resource.service_name.comment")),
		field.String("keyword").
			MaxLen(255).
			Comment(i18n.Text("resource.keyword.comment")).
			Unique().
			NotEmpty(),
		field.String("path").
			Comment(i18n.Text("resource.path.comment")).
			Optional(),
		field.String("method").
			Comment(i18n.Text("resource.method.comment")).
			Optional(),
		field.String("operation").
			Comment(i18n.Text("resource.operation.comment")).
			Optional(),
		field.String("policy").
			Comment(i18n.Text("resource.policy.comment")).
			Default(""),
		field.String("version_id").
			Comment(i18n.Text("resource.version_id.comment")).
			Default(""),
		field.String("last_sync_version_id").
			Comment(i18n.Text("resource.last_sync_version_id.comment")).
			Default(""),
		field.String("sync_status").
			Comment(i18n.Text("resource.sync_status.comment")).
			Default("Synced"),
		field.Enum("status").
			Comment(i18n.Text("resource.status.comment")).
			Values("enabled", "disabled").
			Default("enabled"),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("views", View.Type).
			Ref("resources").
			Through("view_resources", ViewResource.Type),
		edge.From("permissions", Permission.Type).
			Ref("resources"),
	}
}

// Annotations of the Resource.
func (Resource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_resources"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.resource.table.comment")),
	}
}

// Mixin of the Resource.
func (Resource) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}
