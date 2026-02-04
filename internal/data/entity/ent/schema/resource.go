package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"origadmin/application/admin/internal/data/enums"
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
		field.String("keyword").
			MaxLen(255).
			Comment(i18n.Text("entity.resource.field.keyword")).
			Unique().
			NotEmpty(),
		field.String("name").
			Comment(i18n.Text("entity.resource.field.name")).
			Default(""),
		field.String("i18n").
			Comment(i18n.Text("entity.resource.field.i18n")).
			Default(""),
		field.String("type").
			Comment(i18n.Text("entity.resource.field.type")).
			Default("API"),
		field.Int8("status").
			GoType(enums.Status(0)).
			Default(int8(enums.StatusActive)).
			Comment(i18n.Text("entity.resource.field.status")),
		field.Int("sequence").
			Comment(i18n.Text("entity.resource.field.sequence")).
			Default(0),
		field.String("method").
			Comment(i18n.Text("entity.resource.field.method")).
			Default(""),
		field.String("path").
			Comment(i18n.Text("entity.resource.field.path")).
			Default(""),
		field.String("operation").
			Comment(i18n.Text("entity.resource.field.operation")).
			Default(""),
		field.String("service_name").
			Default("").
			Comment(i18n.Text("entity.resource.field.service_name")),
		field.String("policy").
			Comment(i18n.Text("entity.resource.field.policy")).
			Default(""),
		field.String("version_id").
			Comment(i18n.Text("entity.resource.field.version_id")).
			Default(""),
		field.String("last_sync_version_id").
			Comment(i18n.Text("entity.resource.field.last_sync_version_id")).
			Default(""),
		field.String("sync_status").
			Comment(i18n.Text("entity.resource.field.sync_status")).
			Default("Synced"),
		field.String("tree_path").
			Comment(i18n.Text("entity.resource.field.tree_path")).
			Default(""),
		mixin.OptionalFK("parent_id", i18n.Text("entity.resource.field.parent_id")),
		field.String("properties").
			Comment(i18n.Text("entity.resource.field.properties")).
			Default(""),
		field.String("description").
			Comment(i18n.Text("entity.resource.field.description")).
			Default(""),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Resource.Type).
			From("parent").
			Field("parent_id").
			Unique(),
		edge.From("views", View.Type).
			Ref("resources").
			Through("view_resources", ViewResource.Type),
		edge.From("permissions", Permission.Type).
			Ref("resources").
			Through("permission_resources", PermissionResource.Type),
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
