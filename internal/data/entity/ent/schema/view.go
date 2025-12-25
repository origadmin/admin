package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// View holds the schema definition for the View entity.
type View struct {
	ent.Schema
}

// Fields of the View.
func (View) Fields() []ent.Field {
	return []ent.Field{
		// Use OptionalFK for an optional foreign key, as designed in the mixin package.
		mixin.OptionalFK("parent_id", i18n.Text("view.parent_id.comment")),
		field.String("keyword").
			MaxLen(255).
			Comment(i18n.Text("view.keyword.comment")).
			Unique().
			NotEmpty(),
		field.String("scope").
			Comment(i18n.Text("view.scope.comment")).
			Default("default"),
		field.String("name").
			Comment(i18n.Text("view.name.comment")),
		field.Enum("type").
			Comment(i18n.Text("view.type.comment")).
			Values(
				string(enums.ViewTypeRoot),
				string(enums.ViewTypeGroup),
				string(enums.ViewTypeMenu),
				string(enums.ViewTypeLink),
				string(enums.ViewTypePage),
				string(enums.ViewTypeButton),
				string(enums.ViewTypeElement),
				string(enums.ViewTypeRedirect),
				string(enums.ViewTypeUnknown),
			).
			Default(string(enums.ViewTypeUnknown)),
		field.String("component").
			Comment(i18n.Text("view.component.comment")).
			Optional(),
		field.String("path").
			Comment(i18n.Text("view.path.comment")).
			Optional(),
		field.String("icon").
			Comment(i18n.Text("view.icon.comment")).
			Optional(),
		field.Bool("visible").
			Comment(i18n.Text("view.visible.comment")).
			Default(true),
		field.Int("sequence").
			Comment(i18n.Text("view.sequence.comment")).
			Default(0),
		field.String("tree_path").
			Comment(i18n.Text("view.tree_path.comment")).
			Optional(),
	}
}

// Edges of the View.
func (View) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", View.Type).
			From("parent").
			Field("parent_id").
			Unique(),
		edge.To("resources", Resource.Type).
			Through("view_resources", ViewResource.Type),
		edge.To("permissions", Permission.Type).
			Through("view_permissions", ViewPermission.Type),
	}
}

// Mixin of the View.
func (View) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the View.
func (View) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("keyword", "scope").
			Unique(),
	}
}

// Annotations of the View.
func (View) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_views"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.view.table.comment")),
	}
}
