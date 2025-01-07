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

	"origadmin/application/admin/helpers/ent/mixin"
	"origadmin/application/admin/helpers/i18n"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

const (
	// PermTypeSystem 权限类型
	PermTypeSystem   int8 = 1 // 系统权限（如：登录、修改密码）
	PermTypeMenu     int8 = 2 // 菜单权限（如：查看菜单、操作按钮）
	PermTypeData     int8 = 3 // 数据权限（如：查看数据、修改数据）
	PermTypeDept     int8 = 4 // 部门管理权限（如：管理部门人员、分配角色）
	PermTypeResource int8 = 5 // 资源权限（如：API访问权限）
)

const (
	// ScopeSelf 数据范围类型
	ScopeSelf       string = "self"     // 仅本人数据
	ScopeDept       string = "dept"     // 本部门数据
	ScopeSubDept    string = "sub_dept" // 本部门及下级数据
	ScopeCustomDept string = "custom"   // 自定义部门数据
	ScopeAll        string = "all"      // 所有数据
)

// Fields of the Permission.
// // 部门管理权限示例
//
//	{
//	    "name": "部门人员管理",
//	    "keyword": "dept.user.manage",
//	    "type": PermTypeDept,
//	    "scope": ScopeSubDept,
//	}
//
// // 自定义部门管理权限示例
//
//	{
//	    "name": "指定部门资源管理",
//	    "keyword": "dept.resource.manage",
//	    "type": PermTypeDept,
//	    "scope": ScopeCustomDept,
//	    "scope_depts": ["dept1", "dept2"]
//	}
//
// // 数据查看权限示例
//
//	{
//	    "name": "查看财务数据",
//	    "keyword": "finance.view",
//	    "type": PermTypeData,
//	    "scope": ScopeDept,
//	}
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(64).
			Comment("permission.field.name"),
		field.String("keyword").
			MaxLen(64).
			Unique().
			Comment("permission.field.keyword"),
		field.String("description").
			MaxLen(256).
			Optional().
			Comment("permission.field.description"),
		field.String("i18n_key").
			MaxLen(128).
			Default("").
			Comment("permission.field.i18n_key"), //Comment("国际化标识符(如：permission.system.user.manage)").
		field.Int8("type").
			Default(PermTypeMenu).Comment("permission.field.type"), //Comment("权限类型：1-系统 2-菜单 3-数据 4-部门 5-资源"),
		field.String("scope").
			Default(ScopeSelf).Comment("permission.field.scope"), //Comment("数据范围：self-仅本人 dept-本部门 sub_dept-本部门及下级 custom-自定义部门 all-所有"),
		field.JSON("scope_depts", []string{}).
			Optional().
			Comment("permission.field.scope_depts"), //Comment("自定义数据范围的部门ID列表，当scope为custom时有效"),
	}
}

// Annotations of the Permission.
func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_permissions"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("permission.table.comment")),
	}
}

// Mixin of the Permission.
func (Permission) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).
			Ref("permissions").
			Through("role_permissions", RolePermission.Type),
		edge.To("menus", Menu.Type).
			Through("permission_menus", PermissionMenu.Type),
		edge.To("resources", Resource.Type).
			Through("permission_resources", PermissionResource.Type),
	}
}
