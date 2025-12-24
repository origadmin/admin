package schema

import (
	"entgo.io/ent"
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
		edge.To("views", View.Type),
		edge.From("permissions", Permission.Type).
			Ref("resources"),
	}
}

// Mixin of the Resource.
func (Resource) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}
