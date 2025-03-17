// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"origadmin/application/admin/internal/mods/system/dal/entity/ent/schema\",\"Package\":\"origadmin/application/admin/internal/mods/system/dal/entity/ent\",\"Schemas\":[{\"name\":\"Department\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"users\",\"type\":\"User\",\"ref_name\":\"departments\",\"through\":{\"N\":\"user_departments\",\"T\":\"UserDepartment\"},\"inverse\":true},{\"name\":\"positions\",\"type\":\"Position\"},{\"name\":\"children\",\"type\":\"Department\"},{\"name\":\"parent\",\"type\":\"Department\",\"field\":\"parent_id\",\"ref_name\":\"children\",\"unique\":true,\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.primary_key.comment\"},{\"name\":\"create_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"create_time.field.comment\"},{\"name\":\"update_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"comment\":\"update_time.field.comment\"},{\"name\":\"keyword\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"department.field.keyword\"},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"department.field.name\"},{\"name\":\"tree_path\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":256,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"menu.field.tree_path\"},{\"name\":\"sequence\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"department.field.sequence\"},{\"name\":\"status\",\"type\":{\"Type\":9,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":3,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"department.field.status\"},{\"name\":\"level\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1,\"default_kind\":2,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"department.field.level\"},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":1024,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"department.field.description\"},{\"name\":\"parent_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"validators\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"department.field.parent_id\"}],\"indexes\":[{\"fields\":[\"create_time\"]},{\"fields\":[\"update_time\"]},{\"fields\":[\"keyword\"]},{\"fields\":[\"name\"]},{\"fields\":[\"sequence\"]},{\"fields\":[\"status\"]}],\"annotations\":{\"Comment\":{\"Text\":\"department.table.comment\"},\"EntSQL\":{\"table\":\"sys_departments\",\"with_comments\":true}}},{\"name\":\"Permission\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"roles\",\"type\":\"Role\",\"ref_name\":\"permissions\",\"through\":{\"N\":\"role_permissions\",\"T\":\"RolePermission\"},\"inverse\":true},{\"name\":\"resources\",\"type\":\"Resource\",\"through\":{\"N\":\"permission_resources\",\"T\":\"PermissionResource\"}},{\"name\":\"positions\",\"type\":\"Position\",\"ref_name\":\"permissions\",\"through\":{\"N\":\"position_permissions\",\"T\":\"PositionPermission\"},\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.primary_key.comment\"},{\"name\":\"create_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"create_time.field.comment\"},{\"name\":\"update_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"comment\":\"update_time.field.comment\"},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"permission.field.name\"},{\"name\":\"keyword\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"unique\":true,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"permission.field.keyword\"},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":1024,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"permission.field.description\"},{\"name\":\"data_scope\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"self\",\"default_kind\":24,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"permission.field.data_scope\"},{\"name\":\"method\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"GET\",\"default_kind\":24,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"permission.field.method\"},{\"name\":\"path\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"permission.field.path\"},{\"name\":\"data_rules\",\"type\":{\"Type\":3,\"Ident\":\"map[string]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"map[string]string\",\"Kind\":21,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"permission.field.data_rules\"}],\"indexes\":[{\"fields\":[\"create_time\"]},{\"fields\":[\"update_time\"]}],\"annotations\":{\"Comment\":{\"Text\":\"permission.table.comment\"},\"EntSQL\":{\"table\":\"sys_permissions\",\"with_comments\":true}}},{\"name\":\"PermissionResource\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"permission\",\"type\":\"Permission\",\"field\":\"permission_id\",\"unique\":true,\"required\":true},{\"name\":\"resource\",\"type\":\"Resource\",\"field\":\"resource_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"permission_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"},{\"name\":\"resource_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"},{\"name\":\"actions\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"*\",\"default_kind\":24,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"permission_resource.field.actions\"}],\"indexes\":[{\"unique\":true,\"fields\":[\"permission_id\",\"resource_id\"]}],\"annotations\":{\"Comment\":{\"Text\":\"permission_resource.table.comment\"},\"EntSQL\":{\"table\":\"sys_permission_resources\",\"with_comments\":true}}},{\"name\":\"Position\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"department\",\"type\":\"Department\",\"field\":\"department_id\",\"ref_name\":\"positions\",\"unique\":true,\"inverse\":true,\"required\":true},{\"name\":\"users\",\"type\":\"User\",\"ref_name\":\"positions\",\"through\":{\"N\":\"user_positions\",\"T\":\"UserPosition\"},\"inverse\":true},{\"name\":\"permissions\",\"type\":\"Permission\",\"through\":{\"N\":\"position_permissions\",\"T\":\"PositionPermission\"}}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.primary_key.comment\"},{\"name\":\"create_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"create_time.field.comment\"},{\"name\":\"update_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"comment\":\"update_time.field.comment\"},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"position.field.name\"},{\"name\":\"keyword\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"unique\":true,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"position.field.keyword\"},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":1024,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"position.field.description\"},{\"name\":\"department_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"department.field.department_id\"}],\"indexes\":[{\"fields\":[\"create_time\"]},{\"fields\":[\"update_time\"]}],\"annotations\":{\"Comment\":{\"Text\":\"position.table.comment\"},\"EntSQL\":{\"table\":\"sys_positions\",\"with_comments\":true}}},{\"name\":\"PositionPermission\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"position\",\"type\":\"Position\",\"field\":\"position_id\",\"unique\":true,\"required\":true},{\"name\":\"permission\",\"type\":\"Permission\",\"field\":\"permission_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"position_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"position_permission.field.position_id\"},{\"name\":\"permission_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"position_permission.field.permission_id\"}],\"indexes\":[{\"fields\":[\"permission_id\"]},{\"fields\":[\"position_id\"]},{\"unique\":true,\"fields\":[\"position_id\",\"permission_id\"]}],\"annotations\":{\"Comment\":{\"Text\":\"position_permission.table.comment\"},\"EntSQL\":{\"table\":\"sys_position_permissions\",\"with_comments\":true}}},{\"name\":\"Resource\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"children\",\"type\":\"Resource\"},{\"name\":\"parent\",\"type\":\"Resource\",\"field\":\"parent_id\",\"ref_name\":\"children\",\"unique\":true,\"inverse\":true},{\"name\":\"permissions\",\"type\":\"Permission\",\"ref_name\":\"resources\",\"through\":{\"N\":\"permission_resources\",\"T\":\"PermissionResource\"},\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.primary_key.comment\"},{\"name\":\"create_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"create_time.field.comment\"},{\"name\":\"update_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"comment\":\"update_time.field.comment\"},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":128,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.name\"},{\"name\":\"keyword\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"unique\":true,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.keyword\"},{\"name\":\"i18n_key\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":128,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.i18n_key\"},{\"name\":\"type\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":2,\"default\":true,\"default_value\":\"M\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.type\"},{\"name\":\"status\",\"type\":{\"Type\":9,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1,\"default_kind\":3,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.status\"},{\"name\":\"path\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":256,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.path\"},{\"name\":\"operation\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":128,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.operation\"},{\"name\":\"method\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":16,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.method\"},{\"name\":\"component\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":128,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.component\"},{\"name\":\"icon\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":9,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.icon\"},{\"name\":\"sequence\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":10,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.sequence\"},{\"name\":\"visible\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":true,\"default_kind\":1,\"position\":{\"Index\":11,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.visible\"},{\"name\":\"tree_path\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":256,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":12,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.tree_path\"},{\"name\":\"properties\",\"type\":{\"Type\":3,\"Ident\":\"map[string]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"map[string]string\",\"Kind\":21,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":13,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.properties\"},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":1024,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":14,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"resource.field.description\"},{\"name\":\"parent_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"validators\":1,\"position\":{\"Index\":15,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"resource.field.parent_id\"}],\"indexes\":[{\"fields\":[\"create_time\"]},{\"fields\":[\"update_time\"]},{\"fields\":[\"parent_id\"]}],\"annotations\":{\"Comment\":{\"Text\":\"resource.table.comment\"},\"EntSQL\":{\"table\":\"sys_resources\",\"with_comments\":true}}},{\"name\":\"Role\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"users\",\"type\":\"User\",\"ref_name\":\"roles\",\"through\":{\"N\":\"user_roles\",\"T\":\"UserRole\"},\"inverse\":true},{\"name\":\"permissions\",\"type\":\"Permission\",\"through\":{\"N\":\"role_permissions\",\"T\":\"RolePermission\"},\"storage_key\":{\"Table\":\"\",\"Symbols\":null,\"Columns\":[\"role_id\",\"permission_id\"]}}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.primary_key.comment\"},{\"name\":\"create_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"comment\":\"create_time.field.comment\"},{\"name\":\"update_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"comment\":\"update_time.field.comment\"},{\"name\":\"keyword\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"role.field.keyword\"},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":128,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"role.field.name\"},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":1024,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"role.field.description\"},{\"name\":\"type\",\"type\":{\"Type\":9,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":2,\"default_kind\":3,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"role.field.type\"},{\"name\":\"sequence\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"role.field.sequence\"},{\"name\":\"status\",\"type\":{\"Type\":9,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1,\"default_kind\":3,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"role.field.status\"}],\"indexes\":[{\"fields\":[\"create_time\"]},{\"fields\":[\"update_time\"]},{\"fields\":[\"keyword\"]},{\"fields\":[\"name\"]},{\"fields\":[\"sequence\"]},{\"fields\":[\"status\"]}],\"annotations\":{\"Comment\":{\"Text\":\"role.table.comment\"},\"EntSQL\":{\"table\":\"sys_roles\",\"with_comments\":true}}},{\"name\":\"RolePermission\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"role\",\"type\":\"Role\",\"field\":\"role_id\",\"unique\":true,\"required\":true},{\"name\":\"permission\",\"type\":\"Permission\",\"field\":\"permission_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"role_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"},{\"name\":\"permission_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"}],\"indexes\":[{\"fields\":[\"role_id\"]},{\"fields\":[\"permission_id\"]},{\"unique\":true,\"fields\":[\"role_id\",\"permission_id\"]}],\"annotations\":{\"Comment\":{\"Text\":\"role_permission.table.comment\"},\"EntSQL\":{\"table\":\"sys_role_permissions\",\"with_comments\":true}}},{\"name\":\"User\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"roles\",\"type\":\"Role\",\"through\":{\"N\":\"user_roles\",\"T\":\"UserRole\"}},{\"name\":\"positions\",\"type\":\"Position\",\"through\":{\"N\":\"user_positions\",\"T\":\"UserPosition\"}},{\"name\":\"departments\",\"type\":\"Department\",\"through\":{\"N\":\"user_departments\",\"T\":\"UserDepartment\"}}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.primary_key.comment\"},{\"name\":\"create_author\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_value\":0,\"default_kind\":6,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"create_author.field.comment\"},{\"name\":\"update_author\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_value\":0,\"default_kind\":6,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"update_author.field.comment\"},{\"name\":\"create_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"comment\":\"create_time.field.comment\"},{\"name\":\"update_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":3},\"comment\":\"update_time.field.comment\"},{\"name\":\"delete_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":4},\"comment\":\"delete_time.field.comment\"},{\"name\":\"uuid\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":36,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.uuid\"},{\"name\":\"allowed_ip\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"0.0.0.0\",\"default_kind\":24,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.allowed_ip\"},{\"name\":\"username\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"unique\":true,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.username\"},{\"name\":\"nickname\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.nickname\"},{\"name\":\"avatar\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":256,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.avatar\"},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.nickname\"},{\"name\":\"gender\",\"type\":{\"Type\":6,\"Ident\":\"user.Gender\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"enums\":[{\"N\":\"male\",\"V\":\"male\"},{\"N\":\"female\",\"V\":\"female\"},{\"N\":\"unknown\",\"V\":\"unknown\"}],\"default\":true,\"default_value\":\"unknown\",\"default_kind\":24,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.gender\"},{\"name\":\"password\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":256,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.password\"},{\"name\":\"salt\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.salt\"},{\"name\":\"phone\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":9,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.phone\"},{\"name\":\"email\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":10,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.email\"},{\"name\":\"department\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":11,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.department\"},{\"name\":\"remark\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":1024,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":12,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.remark\"},{\"name\":\"token\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":512,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":13,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.token\"},{\"name\":\"status\",\"type\":{\"Type\":9,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1,\"default_kind\":3,\"position\":{\"Index\":14,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.status\"},{\"name\":\"is_system\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":15,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.is_system\"},{\"name\":\"last_login_ip\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":16,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.last_login_ip\"},{\"name\":\"last_login_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":17,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"},\"comment\":\"user.field.last_login_time\"},{\"name\":\"login_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":18,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"},\"comment\":\"user.field.login_time\"},{\"name\":\"sanction_date\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":19,\"MixedIn\":false,\"MixinIndex\":0},\"schema_type\":{\"mysql\":\"datetime\"},\"comment\":\"user.field.sanction_date\"},{\"name\":\"manager_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"validators\":1,\"position\":{\"Index\":20,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"user.field.manager_id\"},{\"name\":\"manager\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"\",\"default_kind\":24,\"position\":{\"Index\":21,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"user.field.manager\"}],\"indexes\":[{\"fields\":[\"create_author\"]},{\"fields\":[\"update_author\"]},{\"fields\":[\"create_time\"]},{\"fields\":[\"update_time\"]},{\"fields\":[\"delete_time\"]},{\"fields\":[\"username\"]},{\"fields\":[\"phone\"]},{\"fields\":[\"email\"]},{\"fields\":[\"status\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":4}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":4}],\"annotations\":{\"Comment\":{\"Text\":\"user.table.comment\"},\"EntSQL\":{\"table\":\"sys_users\",\"with_comments\":true}}},{\"name\":\"UserDepartment\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"user\",\"type\":\"User\",\"field\":\"user_id\",\"unique\":true,\"required\":true},{\"name\":\"department\",\"type\":\"Department\",\"field\":\"department_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"user_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"},{\"name\":\"department_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"}],\"indexes\":[{\"fields\":[\"user_id\"]},{\"fields\":[\"department_id\"]},{\"unique\":true,\"fields\":[\"user_id\",\"department_id\"]}],\"annotations\":{\"Comment\":{\"Text\":\"user_department.table.comment\"},\"EntSQL\":{\"table\":\"sys_user_departments\",\"with_comments\":true}}},{\"name\":\"UserPosition\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"user\",\"type\":\"User\",\"field\":\"user_id\",\"unique\":true,\"required\":true},{\"name\":\"position\",\"type\":\"Position\",\"field\":\"position_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"user_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"},{\"name\":\"position_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"}],\"indexes\":[{\"fields\":[\"user_id\"]},{\"fields\":[\"position_id\"]},{\"unique\":true,\"fields\":[\"user_id\",\"position_id\"]}],\"annotations\":{\"Comment\":{\"Text\":\"user_position.table.comment\"},\"EntSQL\":{\"table\":\"sys_user_positions\",\"with_comments\":true}}},{\"name\":\"UserRole\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"user\",\"type\":\"User\",\"field\":\"user_id\",\"unique\":true,\"required\":true},{\"name\":\"role\",\"type\":\"Role\",\"field\":\"role_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"user_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"},{\"name\":\"role_id\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"annotations\":{\"EntSQL\":{\"incremental\":false}},\"comment\":\"field.foreign_key.comment\"}],\"indexes\":[{\"fields\":[\"user_id\"]},{\"fields\":[\"role_id\"]},{\"unique\":true,\"fields\":[\"user_id\",\"role_id\"]}],\"annotations\":{\"Comment\":{\"Text\":\"(1).(2).(3)\"},\"EntSQL\":{\"table\":\"sys_user_roles\",\"with_comments\":true}}}],\"Features\":[\"intercept\",\"schema/snapshot\",\"sql/versioned-migration\",\"sql/lock\",\"sql/modifier\"]}"
