// Code generated by ent, DO NOT EDIT.

package ent

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menupermission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolemenu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolepermission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/schema"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userdepartment"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userposition"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userrole"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	departmentMixin := schema.Department{}.Mixin()
	departmentMixinFields0 := departmentMixin[0].Fields()
	_ = departmentMixinFields0
	departmentMixinFields1 := departmentMixin[1].Fields()
	_ = departmentMixinFields1
	departmentMixinFields2 := departmentMixin[2].Fields()
	_ = departmentMixinFields2
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescCreateTime is the schema descriptor for create_time field.
	departmentDescCreateTime := departmentMixinFields1[0].Descriptor()
	// department.DefaultCreateTime holds the default value on creation for the create_time field.
	department.DefaultCreateTime = departmentDescCreateTime.Default.(func() time.Time)
	// departmentDescUpdateTime is the schema descriptor for update_time field.
	departmentDescUpdateTime := departmentMixinFields2[0].Descriptor()
	// department.DefaultUpdateTime holds the default value on creation for the update_time field.
	department.DefaultUpdateTime = departmentDescUpdateTime.Default.(func() time.Time)
	// department.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	department.UpdateDefaultUpdateTime = departmentDescUpdateTime.UpdateDefault.(func() time.Time)
	// departmentDescKeyword is the schema descriptor for keyword field.
	departmentDescKeyword := departmentFields[0].Descriptor()
	// department.DefaultKeyword holds the default value on creation for the keyword field.
	department.DefaultKeyword = departmentDescKeyword.Default.(string)
	// department.KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	department.KeywordValidator = departmentDescKeyword.Validators[0].(func(string) error)
	// departmentDescName is the schema descriptor for name field.
	departmentDescName := departmentFields[1].Descriptor()
	// department.DefaultName holds the default value on creation for the name field.
	department.DefaultName = departmentDescName.Default.(string)
	// department.NameValidator is a validator for the "name" field. It is called by the builders before save.
	department.NameValidator = departmentDescName.Validators[0].(func(string) error)
	// departmentDescDescription is the schema descriptor for description field.
	departmentDescDescription := departmentFields[2].Descriptor()
	// department.DefaultDescription holds the default value on creation for the description field.
	department.DefaultDescription = departmentDescDescription.Default.(string)
	// department.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	department.DescriptionValidator = departmentDescDescription.Validators[0].(func(string) error)
	// departmentDescStatus is the schema descriptor for status field.
	departmentDescStatus := departmentFields[4].Descriptor()
	// department.DefaultStatus holds the default value on creation for the status field.
	department.DefaultStatus = departmentDescStatus.Default.(int8)
	// departmentDescID is the schema descriptor for id field.
	departmentDescID := departmentMixinFields0[0].Descriptor()
	// department.IDValidator is a validator for the "id" field. It is called by the builders before save.
	department.IDValidator = departmentDescID.Validators[0].(func(int) error)
	menuMixin := schema.Menu{}.Mixin()
	menuMixinFields0 := menuMixin[0].Fields()
	_ = menuMixinFields0
	menuMixinFields1 := menuMixin[1].Fields()
	_ = menuMixinFields1
	menuMixinFields2 := menuMixin[2].Fields()
	_ = menuMixinFields2
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescCreateTime is the schema descriptor for create_time field.
	menuDescCreateTime := menuMixinFields1[0].Descriptor()
	// menu.DefaultCreateTime holds the default value on creation for the create_time field.
	menu.DefaultCreateTime = menuDescCreateTime.Default.(func() time.Time)
	// menuDescUpdateTime is the schema descriptor for update_time field.
	menuDescUpdateTime := menuMixinFields2[0].Descriptor()
	// menu.DefaultUpdateTime holds the default value on creation for the update_time field.
	menu.DefaultUpdateTime = menuDescUpdateTime.Default.(func() time.Time)
	// menu.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	menu.UpdateDefaultUpdateTime = menuDescUpdateTime.UpdateDefault.(func() time.Time)
	// menuDescKeyword is the schema descriptor for keyword field.
	menuDescKeyword := menuFields[0].Descriptor()
	// menu.DefaultKeyword holds the default value on creation for the keyword field.
	menu.DefaultKeyword = menuDescKeyword.Default.(string)
	// menu.KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	menu.KeywordValidator = menuDescKeyword.Validators[0].(func(string) error)
	// menuDescName is the schema descriptor for name field.
	menuDescName := menuFields[1].Descriptor()
	// menu.DefaultName holds the default value on creation for the name field.
	menu.DefaultName = menuDescName.Default.(string)
	// menu.NameValidator is a validator for the "name" field. It is called by the builders before save.
	menu.NameValidator = menuDescName.Validators[0].(func(string) error)
	// menuDescDescription is the schema descriptor for description field.
	menuDescDescription := menuFields[2].Descriptor()
	// menu.DefaultDescription holds the default value on creation for the description field.
	menu.DefaultDescription = menuDescDescription.Default.(string)
	// menu.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	menu.DescriptionValidator = menuDescDescription.Validators[0].(func(string) error)
	// menuDescType is the schema descriptor for type field.
	menuDescType := menuFields[3].Descriptor()
	// menu.DefaultType holds the default value on creation for the type field.
	menu.DefaultType = menuDescType.Default.(uint8)
	// menuDescIcon is the schema descriptor for icon field.
	menuDescIcon := menuFields[4].Descriptor()
	// menu.DefaultIcon holds the default value on creation for the icon field.
	menu.DefaultIcon = menuDescIcon.Default.(string)
	// menu.IconValidator is a validator for the "icon" field. It is called by the builders before save.
	menu.IconValidator = menuDescIcon.Validators[0].(func(string) error)
	// menuDescPath is the schema descriptor for path field.
	menuDescPath := menuFields[5].Descriptor()
	// menu.DefaultPath holds the default value on creation for the path field.
	menu.DefaultPath = menuDescPath.Default.(string)
	// menu.PathValidator is a validator for the "path" field. It is called by the builders before save.
	menu.PathValidator = menuDescPath.Validators[0].(func(string) error)
	// menuDescStatus is the schema descriptor for status field.
	menuDescStatus := menuFields[6].Descriptor()
	// menu.DefaultStatus holds the default value on creation for the status field.
	menu.DefaultStatus = menuDescStatus.Default.(int8)
	// menuDescParentID is the schema descriptor for parent_id field.
	menuDescParentID := menuFields[7].Descriptor()
	// menu.ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	menu.ParentIDValidator = menuDescParentID.Validators[0].(func(int) error)
	// menuDescParentPath is the schema descriptor for parent_path field.
	menuDescParentPath := menuFields[8].Descriptor()
	// menu.DefaultParentPath holds the default value on creation for the parent_path field.
	menu.DefaultParentPath = menuDescParentPath.Default.(string)
	// menu.ParentPathValidator is a validator for the "parent_path" field. It is called by the builders before save.
	menu.ParentPathValidator = menuDescParentPath.Validators[0].(func(string) error)
	// menuDescSequence is the schema descriptor for sequence field.
	menuDescSequence := menuFields[9].Descriptor()
	// menu.DefaultSequence holds the default value on creation for the sequence field.
	menu.DefaultSequence = menuDescSequence.Default.(int)
	// menuDescProperties is the schema descriptor for properties field.
	menuDescProperties := menuFields[10].Descriptor()
	// menu.DefaultProperties holds the default value on creation for the properties field.
	menu.DefaultProperties = menuDescProperties.Default.(string)
	// menuDescID is the schema descriptor for id field.
	menuDescID := menuMixinFields0[0].Descriptor()
	// menu.IDValidator is a validator for the "id" field. It is called by the builders before save.
	menu.IDValidator = menuDescID.Validators[0].(func(int) error)
	menupermissionMixin := schema.MenuPermission{}.Mixin()
	menupermissionMixinFields0 := menupermissionMixin[0].Fields()
	_ = menupermissionMixinFields0
	menupermissionFields := schema.MenuPermission{}.Fields()
	_ = menupermissionFields
	// menupermissionDescMenuID is the schema descriptor for menu_id field.
	menupermissionDescMenuID := menupermissionFields[0].Descriptor()
	// menupermission.MenuIDValidator is a validator for the "menu_id" field. It is called by the builders before save.
	menupermission.MenuIDValidator = menupermissionDescMenuID.Validators[0].(func(int) error)
	// menupermissionDescPermissionID is the schema descriptor for permission_id field.
	menupermissionDescPermissionID := menupermissionFields[1].Descriptor()
	// menupermission.PermissionIDValidator is a validator for the "permission_id" field. It is called by the builders before save.
	menupermission.PermissionIDValidator = menupermissionDescPermissionID.Validators[0].(func(int) error)
	// menupermissionDescID is the schema descriptor for id field.
	menupermissionDescID := menupermissionMixinFields0[0].Descriptor()
	// menupermission.IDValidator is a validator for the "id" field. It is called by the builders before save.
	menupermission.IDValidator = menupermissionDescID.Validators[0].(func(int) error)
	permissionMixin := schema.Permission{}.Mixin()
	permissionMixinFields0 := permissionMixin[0].Fields()
	_ = permissionMixinFields0
	permissionMixinFields1 := permissionMixin[1].Fields()
	_ = permissionMixinFields1
	permissionMixinFields2 := permissionMixin[2].Fields()
	_ = permissionMixinFields2
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescCreateTime is the schema descriptor for create_time field.
	permissionDescCreateTime := permissionMixinFields1[0].Descriptor()
	// permission.DefaultCreateTime holds the default value on creation for the create_time field.
	permission.DefaultCreateTime = permissionDescCreateTime.Default.(func() time.Time)
	// permissionDescUpdateTime is the schema descriptor for update_time field.
	permissionDescUpdateTime := permissionMixinFields2[0].Descriptor()
	// permission.DefaultUpdateTime holds the default value on creation for the update_time field.
	permission.DefaultUpdateTime = permissionDescUpdateTime.Default.(func() time.Time)
	// permission.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	permission.UpdateDefaultUpdateTime = permissionDescUpdateTime.UpdateDefault.(func() time.Time)
	// permissionDescName is the schema descriptor for name field.
	permissionDescName := permissionFields[0].Descriptor()
	// permission.NameValidator is a validator for the "name" field. It is called by the builders before save.
	permission.NameValidator = permissionDescName.Validators[0].(func(string) error)
	// permissionDescDescription is the schema descriptor for description field.
	permissionDescDescription := permissionFields[1].Descriptor()
	// permission.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	permission.DescriptionValidator = permissionDescDescription.Validators[0].(func(string) error)
	// permissionDescID is the schema descriptor for id field.
	permissionDescID := permissionMixinFields0[0].Descriptor()
	// permission.IDValidator is a validator for the "id" field. It is called by the builders before save.
	permission.IDValidator = permissionDescID.Validators[0].(func(int) error)
	permissionresourceMixin := schema.PermissionResource{}.Mixin()
	permissionresourceMixinFields0 := permissionresourceMixin[0].Fields()
	_ = permissionresourceMixinFields0
	permissionresourceFields := schema.PermissionResource{}.Fields()
	_ = permissionresourceFields
	// permissionresourceDescPermissionID is the schema descriptor for permission_id field.
	permissionresourceDescPermissionID := permissionresourceFields[0].Descriptor()
	// permissionresource.PermissionIDValidator is a validator for the "permission_id" field. It is called by the builders before save.
	permissionresource.PermissionIDValidator = permissionresourceDescPermissionID.Validators[0].(func(int) error)
	// permissionresourceDescResourceID is the schema descriptor for resource_id field.
	permissionresourceDescResourceID := permissionresourceFields[1].Descriptor()
	// permissionresource.ResourceIDValidator is a validator for the "resource_id" field. It is called by the builders before save.
	permissionresource.ResourceIDValidator = permissionresourceDescResourceID.Validators[0].(func(int) error)
	// permissionresourceDescID is the schema descriptor for id field.
	permissionresourceDescID := permissionresourceMixinFields0[0].Descriptor()
	// permissionresource.IDValidator is a validator for the "id" field. It is called by the builders before save.
	permissionresource.IDValidator = permissionresourceDescID.Validators[0].(func(int) error)
	positionMixin := schema.Position{}.Mixin()
	positionMixinFields0 := positionMixin[0].Fields()
	_ = positionMixinFields0
	positionMixinFields1 := positionMixin[1].Fields()
	_ = positionMixinFields1
	positionMixinFields2 := positionMixin[2].Fields()
	_ = positionMixinFields2
	positionFields := schema.Position{}.Fields()
	_ = positionFields
	// positionDescCreateTime is the schema descriptor for create_time field.
	positionDescCreateTime := positionMixinFields1[0].Descriptor()
	// position.DefaultCreateTime holds the default value on creation for the create_time field.
	position.DefaultCreateTime = positionDescCreateTime.Default.(func() time.Time)
	// positionDescUpdateTime is the schema descriptor for update_time field.
	positionDescUpdateTime := positionMixinFields2[0].Descriptor()
	// position.DefaultUpdateTime holds the default value on creation for the update_time field.
	position.DefaultUpdateTime = positionDescUpdateTime.Default.(func() time.Time)
	// position.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	position.UpdateDefaultUpdateTime = positionDescUpdateTime.UpdateDefault.(func() time.Time)
	// positionDescName is the schema descriptor for name field.
	positionDescName := positionFields[0].Descriptor()
	// position.NameValidator is a validator for the "name" field. It is called by the builders before save.
	position.NameValidator = positionDescName.Validators[0].(func(string) error)
	// positionDescDescription is the schema descriptor for description field.
	positionDescDescription := positionFields[1].Descriptor()
	// position.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	position.DescriptionValidator = positionDescDescription.Validators[0].(func(string) error)
	// positionDescDepartmentID is the schema descriptor for department_id field.
	positionDescDepartmentID := positionFields[2].Descriptor()
	// position.DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	position.DepartmentIDValidator = positionDescDepartmentID.Validators[0].(func(int) error)
	// positionDescID is the schema descriptor for id field.
	positionDescID := positionMixinFields0[0].Descriptor()
	// position.IDValidator is a validator for the "id" field. It is called by the builders before save.
	position.IDValidator = positionDescID.Validators[0].(func(int) error)
	resourceMixin := schema.Resource{}.Mixin()
	resourceMixinFields0 := resourceMixin[0].Fields()
	_ = resourceMixinFields0
	resourceMixinFields1 := resourceMixin[1].Fields()
	_ = resourceMixinFields1
	resourceMixinFields2 := resourceMixin[2].Fields()
	_ = resourceMixinFields2
	resourceFields := schema.Resource{}.Fields()
	_ = resourceFields
	// resourceDescCreateTime is the schema descriptor for create_time field.
	resourceDescCreateTime := resourceMixinFields1[0].Descriptor()
	// resource.DefaultCreateTime holds the default value on creation for the create_time field.
	resource.DefaultCreateTime = resourceDescCreateTime.Default.(func() time.Time)
	// resourceDescUpdateTime is the schema descriptor for update_time field.
	resourceDescUpdateTime := resourceMixinFields2[0].Descriptor()
	// resource.DefaultUpdateTime holds the default value on creation for the update_time field.
	resource.DefaultUpdateTime = resourceDescUpdateTime.Default.(func() time.Time)
	// resource.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	resource.UpdateDefaultUpdateTime = resourceDescUpdateTime.UpdateDefault.(func() time.Time)
	// resourceDescMethod is the schema descriptor for method field.
	resourceDescMethod := resourceFields[0].Descriptor()
	// resource.DefaultMethod holds the default value on creation for the method field.
	resource.DefaultMethod = resourceDescMethod.Default.(string)
	// resource.MethodValidator is a validator for the "method" field. It is called by the builders before save.
	resource.MethodValidator = resourceDescMethod.Validators[0].(func(string) error)
	// resourceDescOperation is the schema descriptor for operation field.
	resourceDescOperation := resourceFields[1].Descriptor()
	// resource.DefaultOperation holds the default value on creation for the operation field.
	resource.DefaultOperation = resourceDescOperation.Default.(string)
	// resource.OperationValidator is a validator for the "operation" field. It is called by the builders before save.
	resource.OperationValidator = resourceDescOperation.Validators[0].(func(string) error)
	// resourceDescPath is the schema descriptor for path field.
	resourceDescPath := resourceFields[2].Descriptor()
	// resource.PathValidator is a validator for the "path" field. It is called by the builders before save.
	resource.PathValidator = resourceDescPath.Validators[0].(func(string) error)
	// resourceDescMenuID is the schema descriptor for menu_id field.
	resourceDescMenuID := resourceFields[3].Descriptor()
	// resource.MenuIDValidator is a validator for the "menu_id" field. It is called by the builders before save.
	resource.MenuIDValidator = resourceDescMenuID.Validators[0].(func(int) error)
	// resourceDescID is the schema descriptor for id field.
	resourceDescID := resourceMixinFields0[0].Descriptor()
	// resource.IDValidator is a validator for the "id" field. It is called by the builders before save.
	resource.IDValidator = resourceDescID.Validators[0].(func(int) error)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleMixinFields1 := roleMixin[1].Fields()
	_ = roleMixinFields1
	roleMixinFields2 := roleMixin[2].Fields()
	_ = roleMixinFields2
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreateTime is the schema descriptor for create_time field.
	roleDescCreateTime := roleMixinFields1[0].Descriptor()
	// role.DefaultCreateTime holds the default value on creation for the create_time field.
	role.DefaultCreateTime = roleDescCreateTime.Default.(func() time.Time)
	// roleDescUpdateTime is the schema descriptor for update_time field.
	roleDescUpdateTime := roleMixinFields2[0].Descriptor()
	// role.DefaultUpdateTime holds the default value on creation for the update_time field.
	role.DefaultUpdateTime = roleDescUpdateTime.Default.(func() time.Time)
	// role.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	role.UpdateDefaultUpdateTime = roleDescUpdateTime.UpdateDefault.(func() time.Time)
	// roleDescKeyword is the schema descriptor for keyword field.
	roleDescKeyword := roleFields[0].Descriptor()
	// role.DefaultKeyword holds the default value on creation for the keyword field.
	role.DefaultKeyword = roleDescKeyword.Default.(string)
	// role.KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	role.KeywordValidator = roleDescKeyword.Validators[0].(func(string) error)
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[1].Descriptor()
	// role.DefaultName holds the default value on creation for the name field.
	role.DefaultName = roleDescName.Default.(string)
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescDescription is the schema descriptor for description field.
	roleDescDescription := roleFields[2].Descriptor()
	// role.DefaultDescription holds the default value on creation for the description field.
	role.DefaultDescription = roleDescDescription.Default.(string)
	// role.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	role.DescriptionValidator = roleDescDescription.Validators[0].(func(string) error)
	// roleDescStatus is the schema descriptor for status field.
	roleDescStatus := roleFields[4].Descriptor()
	// role.DefaultStatus holds the default value on creation for the status field.
	role.DefaultStatus = roleDescStatus.Default.(int8)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleMixinFields0[0].Descriptor()
	// role.IDValidator is a validator for the "id" field. It is called by the builders before save.
	role.IDValidator = roleDescID.Validators[0].(func(int) error)
	rolemenuMixin := schema.RoleMenu{}.Mixin()
	rolemenuMixinFields0 := rolemenuMixin[0].Fields()
	_ = rolemenuMixinFields0
	rolemenuFields := schema.RoleMenu{}.Fields()
	_ = rolemenuFields
	// rolemenuDescRoleID is the schema descriptor for role_id field.
	rolemenuDescRoleID := rolemenuFields[0].Descriptor()
	// rolemenu.RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	rolemenu.RoleIDValidator = rolemenuDescRoleID.Validators[0].(func(int) error)
	// rolemenuDescMenuID is the schema descriptor for menu_id field.
	rolemenuDescMenuID := rolemenuFields[1].Descriptor()
	// rolemenu.MenuIDValidator is a validator for the "menu_id" field. It is called by the builders before save.
	rolemenu.MenuIDValidator = rolemenuDescMenuID.Validators[0].(func(int) error)
	// rolemenuDescID is the schema descriptor for id field.
	rolemenuDescID := rolemenuMixinFields0[0].Descriptor()
	// rolemenu.IDValidator is a validator for the "id" field. It is called by the builders before save.
	rolemenu.IDValidator = rolemenuDescID.Validators[0].(func(int) error)
	rolepermissionMixin := schema.RolePermission{}.Mixin()
	rolepermissionMixinFields0 := rolepermissionMixin[0].Fields()
	_ = rolepermissionMixinFields0
	rolepermissionFields := schema.RolePermission{}.Fields()
	_ = rolepermissionFields
	// rolepermissionDescRoleID is the schema descriptor for role_id field.
	rolepermissionDescRoleID := rolepermissionFields[0].Descriptor()
	// rolepermission.RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	rolepermission.RoleIDValidator = rolepermissionDescRoleID.Validators[0].(func(int) error)
	// rolepermissionDescPermissionID is the schema descriptor for permission_id field.
	rolepermissionDescPermissionID := rolepermissionFields[1].Descriptor()
	// rolepermission.PermissionIDValidator is a validator for the "permission_id" field. It is called by the builders before save.
	rolepermission.PermissionIDValidator = rolepermissionDescPermissionID.Validators[0].(func(int) error)
	// rolepermissionDescID is the schema descriptor for id field.
	rolepermissionDescID := rolepermissionMixinFields0[0].Descriptor()
	// rolepermission.IDValidator is a validator for the "id" field. It is called by the builders before save.
	rolepermission.IDValidator = rolepermissionDescID.Validators[0].(func(int) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userMixinFields2 := userMixin[2].Fields()
	_ = userMixinFields2
	userMixinFields3 := userMixin[3].Fields()
	_ = userMixinFields3
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateAuthor is the schema descriptor for create_author field.
	userDescCreateAuthor := userMixinFields1[0].Descriptor()
	// user.DefaultCreateAuthor holds the default value on creation for the create_author field.
	user.DefaultCreateAuthor = userDescCreateAuthor.Default.(string)
	// userDescUpdateAuthor is the schema descriptor for update_author field.
	userDescUpdateAuthor := userMixinFields1[1].Descriptor()
	// user.DefaultUpdateAuthor holds the default value on creation for the update_author field.
	user.DefaultUpdateAuthor = userDescUpdateAuthor.Default.(string)
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields2[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields3[0].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[0].Descriptor()
	// user.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	user.UUIDValidator = userDescUUID.Validators[0].(func(string) error)
	// userDescAllowedIP is the schema descriptor for allowed_ip field.
	userDescAllowedIP := userFields[1].Descriptor()
	// user.DefaultAllowedIP holds the default value on creation for the allowed_ip field.
	user.DefaultAllowedIP = userDescAllowedIP.Default.(string)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[2].Descriptor()
	// user.DefaultUsername holds the default value on creation for the username field.
	user.DefaultUsername = userDescUsername.Default.(string)
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[3].Descriptor()
	// user.DefaultNickname holds the default value on creation for the nickname field.
	user.DefaultNickname = userDescNickname.Default.(string)
	// user.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	user.NicknameValidator = userDescNickname.Validators[0].(func(string) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[4].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = userDescAvatar.Validators[0].(func(string) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[5].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescGender is the schema descriptor for gender field.
	userDescGender := userFields[6].Descriptor()
	// user.DefaultGender holds the default value on creation for the gender field.
	user.DefaultGender = userDescGender.Default.(string)
	// user.GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	user.GenderValidator = userDescGender.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[7].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescSalt is the schema descriptor for salt field.
	userDescSalt := userFields[8].Descriptor()
	// user.DefaultSalt holds the default value on creation for the salt field.
	user.DefaultSalt = userDescSalt.Default.(string)
	// user.SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	user.SaltValidator = userDescSalt.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[9].Descriptor()
	// user.DefaultPhone holds the default value on creation for the phone field.
	user.DefaultPhone = userDescPhone.Default.(string)
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[10].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescRemark is the schema descriptor for remark field.
	userDescRemark := userFields[11].Descriptor()
	// user.DefaultRemark holds the default value on creation for the remark field.
	user.DefaultRemark = userDescRemark.Default.(string)
	// user.RemarkValidator is a validator for the "remark" field. It is called by the builders before save.
	user.RemarkValidator = userDescRemark.Validators[0].(func(string) error)
	// userDescToken is the schema descriptor for token field.
	userDescToken := userFields[12].Descriptor()
	// user.DefaultToken holds the default value on creation for the token field.
	user.DefaultToken = userDescToken.Default.(string)
	// user.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	user.TokenValidator = userDescToken.Validators[0].(func(string) error)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[13].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(int8)
	// userDescLastLoginIP is the schema descriptor for last_login_ip field.
	userDescLastLoginIP := userFields[14].Descriptor()
	// user.DefaultLastLoginIP holds the default value on creation for the last_login_ip field.
	user.DefaultLastLoginIP = userDescLastLoginIP.Default.(string)
	// user.LastLoginIPValidator is a validator for the "last_login_ip" field. It is called by the builders before save.
	user.LastLoginIPValidator = userDescLastLoginIP.Validators[0].(func(string) error)
	// userDescLastLoginTime is the schema descriptor for last_login_time field.
	userDescLastLoginTime := userFields[15].Descriptor()
	// user.DefaultLastLoginTime holds the default value on creation for the last_login_time field.
	user.DefaultLastLoginTime = userDescLastLoginTime.Default.(func() time.Time)
	// userDescSanctionDate is the schema descriptor for sanction_date field.
	userDescSanctionDate := userFields[16].Descriptor()
	// user.DefaultSanctionDate holds the default value on creation for the sanction_date field.
	user.DefaultSanctionDate = userDescSanctionDate.Default.(func() time.Time)
	// userDescDepartmentID is the schema descriptor for department_id field.
	userDescDepartmentID := userFields[17].Descriptor()
	// user.DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	user.DepartmentIDValidator = userDescDepartmentID.Validators[0].(func(int) error)
	// userDescManagerID is the schema descriptor for manager_id field.
	userDescManagerID := userFields[18].Descriptor()
	// user.ManagerIDValidator is a validator for the "manager_id" field. It is called by the builders before save.
	user.ManagerIDValidator = userDescManagerID.Validators[0].(func(int) error)
	// userDescManager is the schema descriptor for manager field.
	userDescManager := userFields[19].Descriptor()
	// user.DefaultManager holds the default value on creation for the manager field.
	user.DefaultManager = userDescManager.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(int) error)
	userdepartmentMixin := schema.UserDepartment{}.Mixin()
	userdepartmentMixinFields0 := userdepartmentMixin[0].Fields()
	_ = userdepartmentMixinFields0
	userdepartmentFields := schema.UserDepartment{}.Fields()
	_ = userdepartmentFields
	// userdepartmentDescUserID is the schema descriptor for user_id field.
	userdepartmentDescUserID := userdepartmentFields[0].Descriptor()
	// userdepartment.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	userdepartment.UserIDValidator = userdepartmentDescUserID.Validators[0].(func(int) error)
	// userdepartmentDescDepartmentID is the schema descriptor for department_id field.
	userdepartmentDescDepartmentID := userdepartmentFields[1].Descriptor()
	// userdepartment.DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	userdepartment.DepartmentIDValidator = userdepartmentDescDepartmentID.Validators[0].(func(int) error)
	// userdepartmentDescID is the schema descriptor for id field.
	userdepartmentDescID := userdepartmentMixinFields0[0].Descriptor()
	// userdepartment.IDValidator is a validator for the "id" field. It is called by the builders before save.
	userdepartment.IDValidator = userdepartmentDescID.Validators[0].(func(int) error)
	userpositionMixin := schema.UserPosition{}.Mixin()
	userpositionMixinFields0 := userpositionMixin[0].Fields()
	_ = userpositionMixinFields0
	userpositionFields := schema.UserPosition{}.Fields()
	_ = userpositionFields
	// userpositionDescUserID is the schema descriptor for user_id field.
	userpositionDescUserID := userpositionFields[0].Descriptor()
	// userposition.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	userposition.UserIDValidator = userpositionDescUserID.Validators[0].(func(int) error)
	// userpositionDescPositionID is the schema descriptor for position_id field.
	userpositionDescPositionID := userpositionFields[1].Descriptor()
	// userposition.PositionIDValidator is a validator for the "position_id" field. It is called by the builders before save.
	userposition.PositionIDValidator = userpositionDescPositionID.Validators[0].(func(int) error)
	// userpositionDescID is the schema descriptor for id field.
	userpositionDescID := userpositionMixinFields0[0].Descriptor()
	// userposition.IDValidator is a validator for the "id" field. It is called by the builders before save.
	userposition.IDValidator = userpositionDescID.Validators[0].(func(int) error)
	userroleMixin := schema.UserRole{}.Mixin()
	userroleMixinFields0 := userroleMixin[0].Fields()
	_ = userroleMixinFields0
	userroleFields := schema.UserRole{}.Fields()
	_ = userroleFields
	// userroleDescUserID is the schema descriptor for user_id field.
	userroleDescUserID := userroleFields[0].Descriptor()
	// userrole.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	userrole.UserIDValidator = userroleDescUserID.Validators[0].(func(int) error)
	// userroleDescRoleID is the schema descriptor for role_id field.
	userroleDescRoleID := userroleFields[1].Descriptor()
	// userrole.RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	userrole.RoleIDValidator = userroleDescRoleID.Validators[0].(func(int) error)
	// userroleDescID is the schema descriptor for id field.
	userroleDescID := userroleMixinFields0[0].Descriptor()
	// userrole.IDValidator is a validator for the "id" field. It is called by the builders before save.
	userrole.IDValidator = userroleDescID.Validators[0].(func(int) error)
}
