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

// 角色类型常量
const (
	RoleTypeSystem     int8 = 1 // 系统角色（如：超级管理员、访客）
	RoleTypeUser       int8 = 2 // 用户角色（如：普通用户、运营、客服）
	RoleTypeDepartment int8 = 3 // 部门角色（如：部门主管、部门成员）
)

// Role holds the schema definition for the Role domain.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword").MaxLen(32).Default(""),       // Code of role (unique)
		field.String("name").MaxLen(128).Default(""),         // Display name of role
		field.String("description").MaxLen(1024).Default(""), // Details about role
		field.Int8("type").
			Default(RoleTypeUser).
			Comment("角色类型：1-系统角色 2-用户角色 3-部门角色"),
		field.Int("sequence"), // Sequence for sorting
		field.Int8("status").Default(0),
		field.Bool("is_system").
			Default(false).
			Comment("是否系统内置（系统内置角色不可删除）"),
	}
}

// Mixin of the Role.
func (Role) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the Role.
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("keyword"),
		index.Fields("name"),
		index.Fields("sequence"),
		index.Fields("status"),
	}
}

// Annotations of the Role.
func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_roles"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("role:table:comment")),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("menus", Menu.Type).
			StorageKey(edge.Columns("role_id", "menu_id")).
			Through("role_menus", RoleMenu.Type),
		edge.From("users", User.Type).
			Ref("roles").
			Through("user_roles", UserRole.Type),
		edge.To("permissions", Permission.Type).
			StorageKey(edge.Columns("role_id", "permission_id")).
			Through("role_permissions", RolePermission.Type),
		edge.From("departments", Department.Type).
			Ref("roles").
			Through("department_roles", DepartmentRole.Type),
	}
}
