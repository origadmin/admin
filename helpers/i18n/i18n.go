/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package i18n implements the functions, types, and interfaces for the module.
package i18n

const (
	// DefaultLanguage defines the default language for the system
	DefaultLanguage = "en_US"
)

// KeyTextMap defines all internationalized key-value pairs used in the system
var KeyTextMap = map[string]map[string]string{
	"zh_CN": {
		"test": "测试",
		// 权限相关
		"permission.system.user.manage":  "系统用户管理权限",
		"permission.dept.staff.transfer": "部门人员调动权限",
		"permission.resource.api.access": "API访问权限",

		"department.name":    "部门名称",
		"department.keyword": "部门编码",

		// 菜单相关
		"menu.system.user": "系统用户管理",
		"menu.system.role": "系统角色管理",
		"menu.dept.manage": "部门管理",

		// 按钮相关
		"button.add":    "添加",
		"button.edit":   "编辑",
		"button.delete": "删除",

		// 通用字段
		"field.name":        "名称",
		"field.description": "描述",
		"field.status":      "状态",
		"field.created_at":  "创建时间",
		"field.updated_at":  "更新时间",
	},
	"en_US": {
		"test": "test",
		// Permissions
		"permission.system.user.manage":  "System User Management Permission",
		"permission.dept.staff.transfer": "Department Staff Transfer Permission",
		"permission.resource.api.access": "API Access Permission",

		// Menus
		"menu.system.user": "System Users",
		"menu.system.role": "System Roles",
		"menu.dept.manage": "Department Management",

		// Buttons
		"button.add":    "Add",
		"button.edit":   "Edit",
		"button.delete": "Delete",

		// Common Fields
		"field.name":        "Name",
		"field.description": "Description",
		"field.status":      "Status",
		"field.created_at":  "Created At",
		"field.updated_at":  "Updated At",

		// Department
		"department:id":            "Primary key of Department",
		"department:department_id": "Foreign key of Department",
		"department:keyword":       "Keyword of Department",
		"department:name":          "Display name of Department",
		"department:description":   "Details about Department",
		"department:sequence":      "Sequence for sorting",
		"department:status":        "Status of the department",
		"department:inherit_roles": "Whether to inherit roles from the parent department",
		"department:ancestors":     "Ancestor list (format: ,1,2,3,)",
		"department:parent_id":     "Parent department ID",
		"department:level":         "Department level",

		// Menu
		"menu:id":          "Primary key of the menu item",
		"menu:menu_id":     "Foreign key of the menu item",
		"menu:keyword":     "Unique keyword for the menu item",
		"menu:name":        "Display name of the menu item",
		"menu:description": "Description of the menu item",
		"menu:type":        "Type of the menu item (e.g., page, link)",
		"menu:icon":        "Icon for the menu item",
		"menu:path":        "Path associated with the menu item",
		"menu:status":      "Status of the menu item (e.g., activated, deactivated)",
		"menu:parent_path": "Parent path of the menu item",
		"menu:sequence":    "Sequence for sorting the menu item",
		"menu:properties":  "Additional properties of the menu item",
		"menu:parent_id":   "Parent ID of the menu item",
	},
}

// LocaleText Obtain the corresponding translated text based on the key and language
func LocaleText(locale string, key string) string {
	// If you cannot find the specified language,
	// Chinese is used by default
	if translations, ok := KeyTextMap[locale]; ok {
		if text, exists := translations[key]; exists {
			//fmt.Println("locale:", locale, "key:", key, "text:", text)
			return text
		}
	}
	// If the specified language cannot be found and
	// the language is the default language,
	// the key itself is returned
	if locale == DefaultLanguage {
		//fmt.Println("locale:", locale, "key:", key, "text:", "default")
		return key
	}
	if translations, ok := KeyTextMap[DefaultLanguage]; ok {
		if text, exists := translations[key]; exists {
			//fmt.Println("locale:", "default", "key:", key, "text:", text)
			return text
		}
	}
	//fmt.Println("locale:", "default", "key:", key, "text:", "default")
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

var locale = "en_US"

// Locale 获取当前系统语言设置
func Locale() string {
	//i18n.PreferredLocale(i18n.Locales, "zh_CN", "en_US")
	return locale
}
