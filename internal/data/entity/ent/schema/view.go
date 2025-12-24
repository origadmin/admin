package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
			Comment(i18n.Text("view.keyword.comment")).
			Unique().
			NotEmpty(),
		field.String("scope").
			Comment(i18n.Text("view.scope.comment")).
			Default("default"),
		field.String("name").
			Comment(i18n.Text("view.name.comment")),
		field.String("type").
			Comment(i18n.Text("view.type.comment")).
			MaxLen(1).
			Default("U"),
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
	}
}

// Edges of the View.
func (View) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", View.Type).
			From("parent").
			Field("parent_id").
			Unique(),
		edge.From("resources", Resource.Type).
			Ref("views"),
		edge.From("permissions", Permission.Type).
			Ref("views"),
	}
}

// Mixin of the View.
func (View) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}
