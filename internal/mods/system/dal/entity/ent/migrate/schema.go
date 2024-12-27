// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysDepartmentsColumns holds the columns for the "sys_departments" table.
	SysDepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "keyword", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "name", Type: field.TypeString, Size: 64, Default: ""},
		{Name: "description", Type: field.TypeString, Size: 256, Default: ""},
		{Name: "sequence", Type: field.TypeInt},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "ancestors", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "level", Type: field.TypeInt, Default: 1},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
	}
	// SysDepartmentsTable holds the schema information for the "sys_departments" table.
	SysDepartmentsTable = &schema.Table{
		Name:       "sys_departments",
		Columns:    SysDepartmentsColumns,
		PrimaryKey: []*schema.Column{SysDepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_departments_sys_departments_children",
				Columns:    []*schema.Column{SysDepartmentsColumns[10]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "department_create_time",
				Unique:  false,
				Columns: []*schema.Column{SysDepartmentsColumns[1]},
			},
			{
				Name:    "department_update_time",
				Unique:  false,
				Columns: []*schema.Column{SysDepartmentsColumns[2]},
			},
			{
				Name:    "department_keyword",
				Unique:  false,
				Columns: []*schema.Column{SysDepartmentsColumns[3]},
			},
			{
				Name:    "department_name",
				Unique:  false,
				Columns: []*schema.Column{SysDepartmentsColumns[4]},
			},
			{
				Name:    "department_sequence",
				Unique:  false,
				Columns: []*schema.Column{SysDepartmentsColumns[6]},
			},
			{
				Name:    "department_status",
				Unique:  false,
				Columns: []*schema.Column{SysDepartmentsColumns[7]},
			},
		},
	}
	// SysDepartmentRolesColumns holds the columns for the "sys_department_roles" table.
	SysDepartmentRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "department_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// SysDepartmentRolesTable holds the schema information for the "sys_department_roles" table.
	SysDepartmentRolesTable = &schema.Table{
		Name:       "sys_department_roles",
		Columns:    SysDepartmentRolesColumns,
		PrimaryKey: []*schema.Column{SysDepartmentRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_department_roles_sys_departments_department",
				Columns:    []*schema.Column{SysDepartmentRolesColumns[1]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "sys_department_roles_sys_roles_role",
				Columns:    []*schema.Column{SysDepartmentRolesColumns[2]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "departmentrole_department_id",
				Unique:  false,
				Columns: []*schema.Column{SysDepartmentRolesColumns[1]},
			},
			{
				Name:    "departmentrole_role_id",
				Unique:  false,
				Columns: []*schema.Column{SysDepartmentRolesColumns[2]},
			},
			{
				Name:    "departmentrole_department_id_role_id",
				Unique:  true,
				Columns: []*schema.Column{SysDepartmentRolesColumns[1], SysDepartmentRolesColumns[2]},
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "keyword", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "name", Type: field.TypeString, Size: 128, Default: ""},
		{Name: "description", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "type", Type: field.TypeInt32, Default: 80},
		{Name: "icon", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "path", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "parent_path", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "sequence", Type: field.TypeInt, Default: 0},
		{Name: "properties", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menus_sys_menus_children",
				Columns:    []*schema.Column{SysMenusColumns[13]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "menu_create_time",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[1]},
			},
			{
				Name:    "menu_update_time",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[2]},
			},
			{
				Name:    "menu_keyword",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[3]},
			},
			{
				Name:    "menu_name",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[4]},
			},
			{
				Name:    "menu_type",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[6]},
			},
			{
				Name:    "menu_status",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[9]},
			},
			{
				Name:    "menu_parent_id",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[13]},
			},
			{
				Name:    "menu_parent_path",
				Unique:  false,
				Columns: []*schema.Column{SysMenusColumns[10]},
			},
		},
	}
	// SysMenuPermissionsColumns holds the columns for the "sys_menu_permissions" table.
	SysMenuPermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "menu_id", Type: field.TypeInt},
		{Name: "permission_id", Type: field.TypeInt},
	}
	// SysMenuPermissionsTable holds the schema information for the "sys_menu_permissions" table.
	SysMenuPermissionsTable = &schema.Table{
		Name:       "sys_menu_permissions",
		Columns:    SysMenuPermissionsColumns,
		PrimaryKey: []*schema.Column{SysMenuPermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menu_permissions_sys_menus_menu",
				Columns:    []*schema.Column{SysMenuPermissionsColumns[1]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "sys_menu_permissions_permissions_permission",
				Columns:    []*schema.Column{SysMenuPermissionsColumns[2]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "menupermission_menu_id",
				Unique:  false,
				Columns: []*schema.Column{SysMenuPermissionsColumns[1]},
			},
			{
				Name:    "menupermission_permission_id",
				Unique:  false,
				Columns: []*schema.Column{SysMenuPermissionsColumns[2]},
			},
			{
				Name:    "menupermission_menu_id_permission_id",
				Unique:  true,
				Columns: []*schema.Column{SysMenuPermissionsColumns[1], SysMenuPermissionsColumns[2]},
			},
		},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 64},
		{Name: "keyword", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 256},
		{Name: "i18n_key", Type: field.TypeString, Unique: true, Size: 128},
		{Name: "type", Type: field.TypeInt8, Default: 2},
		{Name: "scope", Type: field.TypeString, Default: "self"},
		{Name: "scope_depts", Type: field.TypeJSON, Nullable: true},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "permission_create_time",
				Unique:  false,
				Columns: []*schema.Column{PermissionsColumns[1]},
			},
			{
				Name:    "permission_update_time",
				Unique:  false,
				Columns: []*schema.Column{PermissionsColumns[2]},
			},
		},
	}
	// SysPermissionResourcesColumns holds the columns for the "sys_permission_resources" table.
	SysPermissionResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "permission_id", Type: field.TypeInt},
		{Name: "resource_id", Type: field.TypeInt},
	}
	// SysPermissionResourcesTable holds the schema information for the "sys_permission_resources" table.
	SysPermissionResourcesTable = &schema.Table{
		Name:       "sys_permission_resources",
		Columns:    SysPermissionResourcesColumns,
		PrimaryKey: []*schema.Column{SysPermissionResourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_permission_resources_permissions_permission",
				Columns:    []*schema.Column{SysPermissionResourcesColumns[1]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "sys_permission_resources_sys_resources_resource",
				Columns:    []*schema.Column{SysPermissionResourcesColumns[2]},
				RefColumns: []*schema.Column{SysResourcesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "permissionresource_permission_id",
				Unique:  false,
				Columns: []*schema.Column{SysPermissionResourcesColumns[1]},
			},
			{
				Name:    "permissionresource_resource_id",
				Unique:  false,
				Columns: []*schema.Column{SysPermissionResourcesColumns[2]},
			},
			{
				Name:    "permissionresource_permission_id_resource_id",
				Unique:  true,
				Columns: []*schema.Column{SysPermissionResourcesColumns[1], SysPermissionResourcesColumns[2]},
			},
		},
	}
	// PositionsColumns holds the columns for the "positions" table.
	PositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "description", Type: field.TypeString, Size: 256},
		{Name: "department_id", Type: field.TypeInt},
	}
	// PositionsTable holds the schema information for the "positions" table.
	PositionsTable = &schema.Table{
		Name:       "positions",
		Columns:    PositionsColumns,
		PrimaryKey: []*schema.Column{PositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "positions_sys_departments_positions",
				Columns:    []*schema.Column{PositionsColumns[5]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "position_create_time",
				Unique:  false,
				Columns: []*schema.Column{PositionsColumns[1]},
			},
			{
				Name:    "position_update_time",
				Unique:  false,
				Columns: []*schema.Column{PositionsColumns[2]},
			},
		},
	}
	// SysResourcesColumns holds the columns for the "sys_resources" table.
	SysResourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "method", Type: field.TypeString, Size: 20, Default: ""},
		{Name: "operation", Type: field.TypeString, Size: 20, Default: ""},
		{Name: "path", Type: field.TypeString, Size: 255},
		{Name: "menu_id", Type: field.TypeInt, Nullable: true},
	}
	// SysResourcesTable holds the schema information for the "sys_resources" table.
	SysResourcesTable = &schema.Table{
		Name:       "sys_resources",
		Columns:    SysResourcesColumns,
		PrimaryKey: []*schema.Column{SysResourcesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_resources_sys_menus_resources",
				Columns:    []*schema.Column{SysResourcesColumns[6]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "resource_create_time",
				Unique:  false,
				Columns: []*schema.Column{SysResourcesColumns[1]},
			},
			{
				Name:    "resource_update_time",
				Unique:  false,
				Columns: []*schema.Column{SysResourcesColumns[2]},
			},
			{
				Name:    "resource_menu_id",
				Unique:  false,
				Columns: []*schema.Column{SysResourcesColumns[6]},
			},
		},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "keyword", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "name", Type: field.TypeString, Size: 128, Default: ""},
		{Name: "description", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "type", Type: field.TypeInt8, Default: 2},
		{Name: "sequence", Type: field.TypeInt},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "is_system", Type: field.TypeBool, Default: false},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:       "sys_roles",
		Columns:    SysRolesColumns,
		PrimaryKey: []*schema.Column{SysRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_create_time",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[1]},
			},
			{
				Name:    "role_update_time",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[2]},
			},
			{
				Name:    "role_keyword",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[3]},
			},
			{
				Name:    "role_name",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[4]},
			},
			{
				Name:    "role_sequence",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[7]},
			},
			{
				Name:    "role_status",
				Unique:  false,
				Columns: []*schema.Column{SysRolesColumns[8]},
			},
		},
	}
	// SysRoleMenusColumns holds the columns for the "sys_role_menus" table.
	SysRoleMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_id", Type: field.TypeInt},
		{Name: "menu_id", Type: field.TypeInt},
	}
	// SysRoleMenusTable holds the schema information for the "sys_role_menus" table.
	SysRoleMenusTable = &schema.Table{
		Name:       "sys_role_menus",
		Columns:    SysRoleMenusColumns,
		PrimaryKey: []*schema.Column{SysRoleMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_role_menus_sys_roles_role",
				Columns:    []*schema.Column{SysRoleMenusColumns[1]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "sys_role_menus_sys_menus_menu",
				Columns:    []*schema.Column{SysRoleMenusColumns[2]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "rolemenu_role_id",
				Unique:  false,
				Columns: []*schema.Column{SysRoleMenusColumns[1]},
			},
			{
				Name:    "rolemenu_menu_id",
				Unique:  false,
				Columns: []*schema.Column{SysRoleMenusColumns[2]},
			},
			{
				Name:    "rolemenu_role_id_menu_id",
				Unique:  true,
				Columns: []*schema.Column{SysRoleMenusColumns[1], SysRoleMenusColumns[2]},
			},
		},
	}
	// RolePermissionsColumns holds the columns for the "role_permissions" table.
	RolePermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_id", Type: field.TypeInt},
		{Name: "permission_id", Type: field.TypeInt},
	}
	// RolePermissionsTable holds the schema information for the "role_permissions" table.
	RolePermissionsTable = &schema.Table{
		Name:       "role_permissions",
		Columns:    RolePermissionsColumns,
		PrimaryKey: []*schema.Column{RolePermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_permissions_sys_roles_role",
				Columns:    []*schema.Column{RolePermissionsColumns[1]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "role_permissions_permissions_permission",
				Columns:    []*schema.Column{RolePermissionsColumns[2]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "rolepermission_role_id",
				Unique:  false,
				Columns: []*schema.Column{RolePermissionsColumns[1]},
			},
			{
				Name:    "rolepermission_permission_id",
				Unique:  false,
				Columns: []*schema.Column{RolePermissionsColumns[2]},
			},
			{
				Name:    "rolepermission_role_id_permission_id",
				Unique:  true,
				Columns: []*schema.Column{RolePermissionsColumns[1], RolePermissionsColumns[2]},
			},
		},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_author", Type: field.TypeString, Default: ""},
		{Name: "update_author", Type: field.TypeString, Default: ""},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "uuid", Type: field.TypeString, Size: 36},
		{Name: "allowed_ip", Type: field.TypeString, Default: "0.0.0.0"},
		{Name: "username", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "nickname", Type: field.TypeString, Size: 64, Default: ""},
		{Name: "avatar", Type: field.TypeString, Size: 256, Default: ""},
		{Name: "name", Type: field.TypeString, Size: 64, Default: ""},
		{Name: "gender", Type: field.TypeString, Size: 16, Default: ""},
		{Name: "password", Type: field.TypeString, Size: 256, Default: ""},
		{Name: "salt", Type: field.TypeString, Size: 64, Default: ""},
		{Name: "phone", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "email", Type: field.TypeString, Size: 64, Default: ""},
		{Name: "remark", Type: field.TypeString, Size: 1024, Default: ""},
		{Name: "token", Type: field.TypeString, Size: 512, Default: ""},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "last_login_ip", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "last_login_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "sanction_date", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "manager_id", Type: field.TypeInt},
		{Name: "manager", Type: field.TypeString, Default: ""},
	}
	// SysUsersTable holds the schema information for the "sys_users" table.
	SysUsersTable = &schema.Table{
		Name:       "sys_users",
		Columns:    SysUsersColumns,
		PrimaryKey: []*schema.Column{SysUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_create_author",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[1]},
			},
			{
				Name:    "user_update_author",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[2]},
			},
			{
				Name:    "user_create_time",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[3]},
			},
			{
				Name:    "user_update_time",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[4]},
			},
			{
				Name:    "user_username",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[7]},
			},
			{
				Name:    "user_phone",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[14]},
			},
			{
				Name:    "user_email",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[15]},
			},
			{
				Name:    "user_status",
				Unique:  false,
				Columns: []*schema.Column{SysUsersColumns[18]},
			},
		},
	}
	// SysUserDepartmentsColumns holds the columns for the "sys_user_departments" table.
	SysUserDepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "department_id", Type: field.TypeInt},
	}
	// SysUserDepartmentsTable holds the schema information for the "sys_user_departments" table.
	SysUserDepartmentsTable = &schema.Table{
		Name:       "sys_user_departments",
		Columns:    SysUserDepartmentsColumns,
		PrimaryKey: []*schema.Column{SysUserDepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_user_departments_sys_users_user",
				Columns:    []*schema.Column{SysUserDepartmentsColumns[1]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "sys_user_departments_sys_departments_department",
				Columns:    []*schema.Column{SysUserDepartmentsColumns[2]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userdepartment_user_id",
				Unique:  false,
				Columns: []*schema.Column{SysUserDepartmentsColumns[1]},
			},
			{
				Name:    "userdepartment_department_id",
				Unique:  false,
				Columns: []*schema.Column{SysUserDepartmentsColumns[2]},
			},
			{
				Name:    "userdepartment_user_id_department_id",
				Unique:  true,
				Columns: []*schema.Column{SysUserDepartmentsColumns[1], SysUserDepartmentsColumns[2]},
			},
		},
	}
	// SysUserPositionsColumns holds the columns for the "sys_user_positions" table.
	SysUserPositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "position_user_positions", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "position_id", Type: field.TypeInt},
	}
	// SysUserPositionsTable holds the schema information for the "sys_user_positions" table.
	SysUserPositionsTable = &schema.Table{
		Name:       "sys_user_positions",
		Columns:    SysUserPositionsColumns,
		PrimaryKey: []*schema.Column{SysUserPositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_user_positions_positions_user_positions",
				Columns:    []*schema.Column{SysUserPositionsColumns[1]},
				RefColumns: []*schema.Column{PositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "sys_user_positions_sys_users_user",
				Columns:    []*schema.Column{SysUserPositionsColumns[2]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "sys_user_positions_positions_position",
				Columns:    []*schema.Column{SysUserPositionsColumns[3]},
				RefColumns: []*schema.Column{PositionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userposition_user_id",
				Unique:  false,
				Columns: []*schema.Column{SysUserPositionsColumns[2]},
			},
			{
				Name:    "userposition_position_id",
				Unique:  false,
				Columns: []*schema.Column{SysUserPositionsColumns[3]},
			},
			{
				Name:    "userposition_user_id_position_id",
				Unique:  true,
				Columns: []*schema.Column{SysUserPositionsColumns[2], SysUserPositionsColumns[3]},
			},
		},
	}
	// SysUserRolesColumns holds the columns for the "sys_user_roles" table.
	SysUserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// SysUserRolesTable holds the schema information for the "sys_user_roles" table.
	SysUserRolesTable = &schema.Table{
		Name:       "sys_user_roles",
		Columns:    SysUserRolesColumns,
		PrimaryKey: []*schema.Column{SysUserRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_user_roles_sys_users_user",
				Columns:    []*schema.Column{SysUserRolesColumns[1]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "sys_user_roles_sys_roles_role",
				Columns:    []*schema.Column{SysUserRolesColumns[2]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userrole_user_id",
				Unique:  false,
				Columns: []*schema.Column{SysUserRolesColumns[1]},
			},
			{
				Name:    "userrole_role_id",
				Unique:  false,
				Columns: []*schema.Column{SysUserRolesColumns[2]},
			},
			{
				Name:    "userrole_user_id_role_id",
				Unique:  true,
				Columns: []*schema.Column{SysUserRolesColumns[1], SysUserRolesColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysDepartmentsTable,
		SysDepartmentRolesTable,
		SysMenusTable,
		SysMenuPermissionsTable,
		PermissionsTable,
		SysPermissionResourcesTable,
		PositionsTable,
		SysResourcesTable,
		SysRolesTable,
		SysRoleMenusTable,
		RolePermissionsTable,
		SysUsersTable,
		SysUserDepartmentsTable,
		SysUserPositionsTable,
		SysUserRolesTable,
	}
)

func init() {
	SysDepartmentsTable.ForeignKeys[0].RefTable = SysDepartmentsTable
	SysDepartmentsTable.Annotation = &entsql.Annotation{
		Table: "sys_departments",
	}
	SysDepartmentRolesTable.ForeignKeys[0].RefTable = SysDepartmentsTable
	SysDepartmentRolesTable.ForeignKeys[1].RefTable = SysRolesTable
	SysDepartmentRolesTable.Annotation = &entsql.Annotation{
		Table: "sys_department_roles",
	}
	SysMenusTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenusTable.Annotation = &entsql.Annotation{
		Table: "sys_menus",
	}
	SysMenuPermissionsTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenuPermissionsTable.ForeignKeys[1].RefTable = PermissionsTable
	SysMenuPermissionsTable.Annotation = &entsql.Annotation{
		Table: "sys_menu_permissions",
	}
	SysPermissionResourcesTable.ForeignKeys[0].RefTable = PermissionsTable
	SysPermissionResourcesTable.ForeignKeys[1].RefTable = SysResourcesTable
	SysPermissionResourcesTable.Annotation = &entsql.Annotation{
		Table: "sys_permission_resources",
	}
	PositionsTable.ForeignKeys[0].RefTable = SysDepartmentsTable
	SysResourcesTable.ForeignKeys[0].RefTable = SysMenusTable
	SysResourcesTable.Annotation = &entsql.Annotation{
		Table: "sys_resources",
	}
	SysRolesTable.Annotation = &entsql.Annotation{
		Table: "sys_roles",
	}
	SysRoleMenusTable.ForeignKeys[0].RefTable = SysRolesTable
	SysRoleMenusTable.ForeignKeys[1].RefTable = SysMenusTable
	SysRoleMenusTable.Annotation = &entsql.Annotation{
		Table: "sys_role_menus",
	}
	RolePermissionsTable.ForeignKeys[0].RefTable = SysRolesTable
	RolePermissionsTable.ForeignKeys[1].RefTable = PermissionsTable
	SysUsersTable.Annotation = &entsql.Annotation{
		Table: "sys_users",
	}
	SysUserDepartmentsTable.ForeignKeys[0].RefTable = SysUsersTable
	SysUserDepartmentsTable.ForeignKeys[1].RefTable = SysDepartmentsTable
	SysUserDepartmentsTable.Annotation = &entsql.Annotation{
		Table: "sys_user_departments",
	}
	SysUserPositionsTable.ForeignKeys[0].RefTable = PositionsTable
	SysUserPositionsTable.ForeignKeys[1].RefTable = SysUsersTable
	SysUserPositionsTable.ForeignKeys[2].RefTable = PositionsTable
	SysUserPositionsTable.Annotation = &entsql.Annotation{
		Table: "sys_user_positions",
	}
	SysUserRolesTable.ForeignKeys[0].RefTable = SysUsersTable
	SysUserRolesTable.ForeignKeys[1].RefTable = SysRolesTable
	SysUserRolesTable.Annotation = &entsql.Annotation{
		Table: "sys_user_roles",
	}
}
