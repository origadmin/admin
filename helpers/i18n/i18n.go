/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package i18n implements the functions, types, and interfaces for the module.
package i18n

const (
	// DefaultLanguage defines the default language for the system
	DefaultLanguage = "en_US"
)

var locale = "en_US"

// KeyTextMap defines all internationalized key-value pairs used in the system
var KeyTextMap = map[string]map[string]string{
	"zh_CN": {
		"entity.test": "测试",
		// 权限相关
		"entity.permission.system.user.manage":  "系统用户管理权限",
		"entity.permission.dept.staff.transfer": "部门人员调动权限",
		"entity.permission.resource.api.access": "API访问权限",

		"entity.department.name":    "部门名称",
		"entity.department.keyword": "部门编码",

		// 菜单相关
		"entity.menu.system.user": "系统用户管理",
		"entity.menu.system.role": "系统角色管理",
		"entity.menu.dept.manage": "部门管理",

		// 按钮相关
		"entity.button.add":    "添加",
		"entity.button.edit":   "编辑",
		"entity.button.delete": "删除",

		// 通用字段
		"entity.field.name":        "名称",
		"entity.field.description": "描述",
		"entity.field.status":      "状态",
		"entity.field.created_at":  "创建时间",
		"entity.field.updated_at":  "更新时间",
	},
	"en_US": {
		"entity.test": "test",
		// Permissions
		"entity.permission.system.user.manage":  "System User Management Permission",
		"entity.permission.dept.staff.transfer": "Department Staff Transfer Permission",
		"entity.permission.resource.api.access": "API Access Permission",

		// Menus
		"entity.menu.system.user": "System Users",
		"entity.menu.system.role": "System Roles",
		"entity.menu.dept.manage": "Department Management",

		// Buttons
		"entity.button.add":    "Add",
		"entity.button.edit":   "Edit",
		"entity.button.delete": "Delete",

		// Common Fields
		"entity.field.name":        "Name",
		"entity.field.description": "Description",
		"entity.field.status":      "Status",
		"entity.field.created_at":  "Created At",
		"entity.field.updated_at":  "Updated At",

		// Department
		"entity.department.id":            "Primary key of Department",
		"entity.department.department_id": "Foreign key of Department",
		"entity.department.keyword":       "Keyword of Department",
		"entity.department.name":          "Display name of Department",
		"entity.department.description":   "Details about Department",
		"entity.department.sequence":      "Sequence for sorting",
		"entity.department.status":        "Status of the department",
		"entity.department.inherit_roles": "Whether to inherit roles from the parent department",
		"entity.department.ancestors":     "Ancestor list (format: ,1,2,3,)",
		"entity.department.parent_id":     "Parent department ID",
		"entity.department.level":         "Department level",

		// Menu
		"entity.menu.id":          "Primary key of the menu item",
		"entity.menu.menu_id":     "Foreign key of the menu item",
		"entity.menu.keyword":     "Unique keyword for the menu item",
		"entity.menu.name":        "Display name of the menu item",
		"entity.menu.description": "Description of the menu item",
		"entity.menu.type":        "Type of the menu item (e.g., page, link)",
		"entity.menu.icon":        "Icon for the menu item",
		"entity.menu.path":        "Path associated with the menu item",
		"entity.menu.status":      "Status of the menu item (e.g., activated, deactivated)",
		"entity.menu.parent_path": "Parent path of the menu item",
		"entity.menu.sequence":    "Sequence for sorting the menu item",
		"entity.menu.properties":  "Additional properties of the menu item",
		"entity.menu.parent_id":   "Parent ID of the menu item",
	},
}

// LocaleText Obtain the corresponding translated text based on the key and language
func LocaleText(locale string, key string) string {
	// If you cannot find the specified language,
	// Chinese is used by default
	if translations, ok := KeyTextMap[locale]; ok {
		if text, exists := translations[key]; exists {
			//fmt.Println("locale.", locale, "key.", key, "text.", text)
			return text
		}
	}
	// If the specified language cannot be found and
	// the language is the default language,
	// the key itself is returned
	if locale == DefaultLanguage {
		//fmt.Println("locale.", locale, "key.", key, "text.", "default")
		return key
	}
	if translations, ok := KeyTextMap[DefaultLanguage]; ok {
		if text, exists := translations[key]; exists {
			//fmt.Println("locale.", "default", "key.", key, "text.", text)
			return text
		}
	}
	//fmt.Println("locale.", "default", "key.", key, "text.", "default")
	return key
}

func Text(key string) string {
	return LocaleText(locale, key)
}

// Docs
// // 权限相关
// permission.system.user.manage      // 系统用户管理权限
// permission.dept.staff.transfer     // 部门人员调动权限
// permission.resource.api.access     // API访问权限

// // 菜单相关
// menu.system.user                  // 系统用户菜单
// menu.system.role                  // 系统角色菜单
// menu.dept.manage                  // 部门管理菜单

// // 按钮相关
// button.add                        // 添加按钮
// button.edit                       // 编辑按钮
// button.delete                     // 删除按钮

// Locale 获取当前系统语言设置
func Locale() string {
	//i18n.PreferredLocale(i18n.Locales, "zh_CN", "en_US")
	return locale
}
