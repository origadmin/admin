// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/positionpermission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
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
	// department.KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	department.KeywordValidator = departmentDescKeyword.Validators[0].(func(string) error)
	// departmentDescName is the schema descriptor for name field.
	departmentDescName := departmentFields[1].Descriptor()
	// department.DefaultName holds the default value on creation for the name field.
	department.DefaultName = departmentDescName.Default.(string)
	// department.NameValidator is a validator for the "name" field. It is called by the builders before save.
	department.NameValidator = departmentDescName.Validators[0].(func(string) error)
	// departmentDescTreePath is the schema descriptor for tree_path field.
	departmentDescTreePath := departmentFields[2].Descriptor()
	// department.DefaultTreePath holds the default value on creation for the tree_path field.
	department.DefaultTreePath = departmentDescTreePath.Default.(string)
	// department.TreePathValidator is a validator for the "tree_path" field. It is called by the builders before save.
	department.TreePathValidator = departmentDescTreePath.Validators[0].(func(string) error)
	// departmentDescStatus is the schema descriptor for status field.
	departmentDescStatus := departmentFields[4].Descriptor()
	// department.DefaultStatus holds the default value on creation for the status field.
	department.DefaultStatus = departmentDescStatus.Default.(int8)
	// departmentDescLevel is the schema descriptor for level field.
	departmentDescLevel := departmentFields[5].Descriptor()
	// department.DefaultLevel holds the default value on creation for the level field.
	department.DefaultLevel = departmentDescLevel.Default.(int)
	// departmentDescDescription is the schema descriptor for description field.
	departmentDescDescription := departmentFields[6].Descriptor()
	// department.DefaultDescription holds the default value on creation for the description field.
	department.DefaultDescription = departmentDescDescription.Default.(string)
	// department.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	department.DescriptionValidator = departmentDescDescription.Validators[0].(func(string) error)
	// departmentDescParentID is the schema descriptor for parent_id field.
	departmentDescParentID := departmentFields[7].Descriptor()
	// department.ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	department.ParentIDValidator = departmentDescParentID.Validators[0].(func(int64) error)
	// departmentDescID is the schema descriptor for id field.
	departmentDescID := departmentMixinFields0[0].Descriptor()
	// department.DefaultID holds the default value on creation for the id field.
	department.DefaultID = departmentDescID.Default.(func() int64)
	// department.IDValidator is a validator for the "id" field. It is called by the builders before save.
	department.IDValidator = departmentDescID.Validators[0].(func(int64) error)
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
	// permission.DefaultName holds the default value on creation for the name field.
	permission.DefaultName = permissionDescName.Default.(string)
	// permission.NameValidator is a validator for the "name" field. It is called by the builders before save.
	permission.NameValidator = permissionDescName.Validators[0].(func(string) error)
	// permissionDescKeyword is the schema descriptor for keyword field.
	permissionDescKeyword := permissionFields[1].Descriptor()
	// permission.KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	permission.KeywordValidator = permissionDescKeyword.Validators[0].(func(string) error)
	// permissionDescDescription is the schema descriptor for description field.
	permissionDescDescription := permissionFields[2].Descriptor()
	// permission.DefaultDescription holds the default value on creation for the description field.
	permission.DefaultDescription = permissionDescDescription.Default.(string)
	// permission.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	permission.DescriptionValidator = permissionDescDescription.Validators[0].(func(string) error)
	// permissionDescDataScope is the schema descriptor for data_scope field.
	permissionDescDataScope := permissionFields[3].Descriptor()
	// permission.DefaultDataScope holds the default value on creation for the data_scope field.
	permission.DefaultDataScope = permissionDescDataScope.Default.(string)
	// permissionDescID is the schema descriptor for id field.
	permissionDescID := permissionMixinFields0[0].Descriptor()
	// permission.DefaultID holds the default value on creation for the id field.
	permission.DefaultID = permissionDescID.Default.(func() int64)
	// permission.IDValidator is a validator for the "id" field. It is called by the builders before save.
	permission.IDValidator = permissionDescID.Validators[0].(func(int64) error)
	permissionresourceFields := schema.PermissionResource{}.Fields()
	_ = permissionresourceFields
	// permissionresourceDescPermissionID is the schema descriptor for permission_id field.
	permissionresourceDescPermissionID := permissionresourceFields[0].Descriptor()
	// permissionresource.PermissionIDValidator is a validator for the "permission_id" field. It is called by the builders before save.
	permissionresource.PermissionIDValidator = permissionresourceDescPermissionID.Validators[0].(func(int64) error)
	// permissionresourceDescResourceID is the schema descriptor for resource_id field.
	permissionresourceDescResourceID := permissionresourceFields[1].Descriptor()
	// permissionresource.ResourceIDValidator is a validator for the "resource_id" field. It is called by the builders before save.
	permissionresource.ResourceIDValidator = permissionresourceDescResourceID.Validators[0].(func(int64) error)
	// permissionresourceDescActions is the schema descriptor for actions field.
	permissionresourceDescActions := permissionresourceFields[2].Descriptor()
	// permissionresource.DefaultActions holds the default value on creation for the actions field.
	permissionresource.DefaultActions = permissionresourceDescActions.Default.(string)
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
	// positionDescKeyword is the schema descriptor for keyword field.
	positionDescKeyword := positionFields[1].Descriptor()
	// position.KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	position.KeywordValidator = positionDescKeyword.Validators[0].(func(string) error)
	// positionDescDescription is the schema descriptor for description field.
	positionDescDescription := positionFields[2].Descriptor()
	// position.DefaultDescription holds the default value on creation for the description field.
	position.DefaultDescription = positionDescDescription.Default.(string)
	// position.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	position.DescriptionValidator = positionDescDescription.Validators[0].(func(string) error)
	// positionDescDepartmentID is the schema descriptor for department_id field.
	positionDescDepartmentID := positionFields[3].Descriptor()
	// position.DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	position.DepartmentIDValidator = positionDescDepartmentID.Validators[0].(func(int64) error)
	// positionDescID is the schema descriptor for id field.
	positionDescID := positionMixinFields0[0].Descriptor()
	// position.DefaultID holds the default value on creation for the id field.
	position.DefaultID = positionDescID.Default.(func() int64)
	// position.IDValidator is a validator for the "id" field. It is called by the builders before save.
	position.IDValidator = positionDescID.Validators[0].(func(int64) error)
	positionpermissionFields := schema.PositionPermission{}.Fields()
	_ = positionpermissionFields
	// positionpermissionDescPositionID is the schema descriptor for position_id field.
	positionpermissionDescPositionID := positionpermissionFields[0].Descriptor()
	// positionpermission.PositionIDValidator is a validator for the "position_id" field. It is called by the builders before save.
	positionpermission.PositionIDValidator = positionpermissionDescPositionID.Validators[0].(func(int64) error)
	// positionpermissionDescPermissionID is the schema descriptor for permission_id field.
	positionpermissionDescPermissionID := positionpermissionFields[1].Descriptor()
	// positionpermission.PermissionIDValidator is a validator for the "permission_id" field. It is called by the builders before save.
	positionpermission.PermissionIDValidator = positionpermissionDescPermissionID.Validators[0].(func(int64) error)
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
	// resourceDescName is the schema descriptor for name field.
	resourceDescName := resourceFields[0].Descriptor()
	// resource.DefaultName holds the default value on creation for the name field.
	resource.DefaultName = resourceDescName.Default.(string)
	// resource.NameValidator is a validator for the "name" field. It is called by the builders before save.
	resource.NameValidator = resourceDescName.Validators[0].(func(string) error)
	// resourceDescKeyword is the schema descriptor for keyword field.
	resourceDescKeyword := resourceFields[1].Descriptor()
	// resource.KeywordValidator is a validator for the "keyword" field. It is called by the builders before save.
	resource.KeywordValidator = resourceDescKeyword.Validators[0].(func(string) error)
	// resourceDescI18nKey is the schema descriptor for i18n_key field.
	resourceDescI18nKey := resourceFields[2].Descriptor()
	// resource.DefaultI18nKey holds the default value on creation for the i18n_key field.
	resource.DefaultI18nKey = resourceDescI18nKey.Default.(string)
	// resource.I18nKeyValidator is a validator for the "i18n_key" field. It is called by the builders before save.
	resource.I18nKeyValidator = resourceDescI18nKey.Validators[0].(func(string) error)
	// resourceDescType is the schema descriptor for type field.
	resourceDescType := resourceFields[3].Descriptor()
	// resource.DefaultType holds the default value on creation for the type field.
	resource.DefaultType = resourceDescType.Default.(string)
	// resource.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	resource.TypeValidator = resourceDescType.Validators[0].(func(string) error)
	// resourceDescStatus is the schema descriptor for status field.
	resourceDescStatus := resourceFields[4].Descriptor()
	// resource.DefaultStatus holds the default value on creation for the status field.
	resource.DefaultStatus = resourceDescStatus.Default.(int8)
	// resourceDescPath is the schema descriptor for path field.
	resourceDescPath := resourceFields[5].Descriptor()
	// resource.DefaultPath holds the default value on creation for the path field.
	resource.DefaultPath = resourceDescPath.Default.(string)
	// resource.PathValidator is a validator for the "path" field. It is called by the builders before save.
	resource.PathValidator = resourceDescPath.Validators[0].(func(string) error)
	// resourceDescOperation is the schema descriptor for operation field.
	resourceDescOperation := resourceFields[6].Descriptor()
	// resource.DefaultOperation holds the default value on creation for the operation field.
	resource.DefaultOperation = resourceDescOperation.Default.(string)
	// resource.OperationValidator is a validator for the "operation" field. It is called by the builders before save.
	resource.OperationValidator = resourceDescOperation.Validators[0].(func(string) error)
	// resourceDescMethod is the schema descriptor for method field.
	resourceDescMethod := resourceFields[7].Descriptor()
	// resource.DefaultMethod holds the default value on creation for the method field.
	resource.DefaultMethod = resourceDescMethod.Default.(string)
	// resource.MethodValidator is a validator for the "method" field. It is called by the builders before save.
	resource.MethodValidator = resourceDescMethod.Validators[0].(func(string) error)
	// resourceDescComponent is the schema descriptor for component field.
	resourceDescComponent := resourceFields[8].Descriptor()
	// resource.DefaultComponent holds the default value on creation for the component field.
	resource.DefaultComponent = resourceDescComponent.Default.(string)
	// resource.ComponentValidator is a validator for the "component" field. It is called by the builders before save.
	resource.ComponentValidator = resourceDescComponent.Validators[0].(func(string) error)
	// resourceDescIcon is the schema descriptor for icon field.
	resourceDescIcon := resourceFields[9].Descriptor()
	// resource.DefaultIcon holds the default value on creation for the icon field.
	resource.DefaultIcon = resourceDescIcon.Default.(string)
	// resource.IconValidator is a validator for the "icon" field. It is called by the builders before save.
	resource.IconValidator = resourceDescIcon.Validators[0].(func(string) error)
	// resourceDescSequence is the schema descriptor for sequence field.
	resourceDescSequence := resourceFields[10].Descriptor()
	// resource.DefaultSequence holds the default value on creation for the sequence field.
	resource.DefaultSequence = resourceDescSequence.Default.(int)
	// resourceDescVisible is the schema descriptor for visible field.
	resourceDescVisible := resourceFields[11].Descriptor()
	// resource.DefaultVisible holds the default value on creation for the visible field.
	resource.DefaultVisible = resourceDescVisible.Default.(bool)
	// resourceDescTreePath is the schema descriptor for tree_path field.
	resourceDescTreePath := resourceFields[12].Descriptor()
	// resource.DefaultTreePath holds the default value on creation for the tree_path field.
	resource.DefaultTreePath = resourceDescTreePath.Default.(string)
	// resource.TreePathValidator is a validator for the "tree_path" field. It is called by the builders before save.
	resource.TreePathValidator = resourceDescTreePath.Validators[0].(func(string) error)
	// resourceDescDescription is the schema descriptor for description field.
	resourceDescDescription := resourceFields[14].Descriptor()
	// resource.DefaultDescription holds the default value on creation for the description field.
	resource.DefaultDescription = resourceDescDescription.Default.(string)
	// resource.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	resource.DescriptionValidator = resourceDescDescription.Validators[0].(func(string) error)
	// resourceDescParentID is the schema descriptor for parent_id field.
	resourceDescParentID := resourceFields[15].Descriptor()
	// resource.ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	resource.ParentIDValidator = resourceDescParentID.Validators[0].(func(int64) error)
	// resourceDescID is the schema descriptor for id field.
	resourceDescID := resourceMixinFields0[0].Descriptor()
	// resource.DefaultID holds the default value on creation for the id field.
	resource.DefaultID = resourceDescID.Default.(func() int64)
	// resource.IDValidator is a validator for the "id" field. It is called by the builders before save.
	resource.IDValidator = resourceDescID.Validators[0].(func(int64) error)
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
	// roleDescType is the schema descriptor for type field.
	roleDescType := roleFields[3].Descriptor()
	// role.DefaultType holds the default value on creation for the type field.
	role.DefaultType = roleDescType.Default.(int8)
	// roleDescSequence is the schema descriptor for sequence field.
	roleDescSequence := roleFields[4].Descriptor()
	// role.DefaultSequence holds the default value on creation for the sequence field.
	role.DefaultSequence = roleDescSequence.Default.(int)
	// roleDescStatus is the schema descriptor for status field.
	roleDescStatus := roleFields[5].Descriptor()
	// role.DefaultStatus holds the default value on creation for the status field.
	role.DefaultStatus = roleDescStatus.Default.(int8)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleMixinFields0[0].Descriptor()
	// role.DefaultID holds the default value on creation for the id field.
	role.DefaultID = roleDescID.Default.(func() int64)
	// role.IDValidator is a validator for the "id" field. It is called by the builders before save.
	role.IDValidator = roleDescID.Validators[0].(func(int64) error)
	rolepermissionFields := schema.RolePermission{}.Fields()
	_ = rolepermissionFields
	// rolepermissionDescRoleID is the schema descriptor for role_id field.
	rolepermissionDescRoleID := rolepermissionFields[0].Descriptor()
	// rolepermission.RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	rolepermission.RoleIDValidator = rolepermissionDescRoleID.Validators[0].(func(int64) error)
	// rolepermissionDescPermissionID is the schema descriptor for permission_id field.
	rolepermissionDescPermissionID := rolepermissionFields[1].Descriptor()
	// rolepermission.PermissionIDValidator is a validator for the "permission_id" field. It is called by the builders before save.
	rolepermission.PermissionIDValidator = rolepermissionDescPermissionID.Validators[0].(func(int64) error)
	userMixin := schema.User{}.Mixin()
	userMixinHooks4 := userMixin[4].Hooks()
	user.Hooks[0] = userMixinHooks4[0]
	userMixinInters4 := userMixin[4].Interceptors()
	user.Interceptors[0] = userMixinInters4[0]
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
	user.DefaultCreateAuthor = userDescCreateAuthor.Default.(int64)
	// userDescUpdateAuthor is the schema descriptor for update_author field.
	userDescUpdateAuthor := userMixinFields1[1].Descriptor()
	// user.DefaultUpdateAuthor holds the default value on creation for the update_author field.
	user.DefaultUpdateAuthor = userDescUpdateAuthor.Default.(int64)
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
	// userDescDepartment is the schema descriptor for department field.
	userDescDepartment := userFields[11].Descriptor()
	// user.DefaultDepartment holds the default value on creation for the department field.
	user.DefaultDepartment = userDescDepartment.Default.(string)
	// user.DepartmentValidator is a validator for the "department" field. It is called by the builders before save.
	user.DepartmentValidator = userDescDepartment.Validators[0].(func(string) error)
	// userDescRemark is the schema descriptor for remark field.
	userDescRemark := userFields[12].Descriptor()
	// user.DefaultRemark holds the default value on creation for the remark field.
	user.DefaultRemark = userDescRemark.Default.(string)
	// user.RemarkValidator is a validator for the "remark" field. It is called by the builders before save.
	user.RemarkValidator = userDescRemark.Validators[0].(func(string) error)
	// userDescToken is the schema descriptor for token field.
	userDescToken := userFields[13].Descriptor()
	// user.DefaultToken holds the default value on creation for the token field.
	user.DefaultToken = userDescToken.Default.(string)
	// user.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	user.TokenValidator = userDescToken.Validators[0].(func(string) error)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[14].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(int8)
	// userDescIsSystem is the schema descriptor for is_system field.
	userDescIsSystem := userFields[15].Descriptor()
	// user.DefaultIsSystem holds the default value on creation for the is_system field.
	user.DefaultIsSystem = userDescIsSystem.Default.(bool)
	// userDescLastLoginIP is the schema descriptor for last_login_ip field.
	userDescLastLoginIP := userFields[16].Descriptor()
	// user.DefaultLastLoginIP holds the default value on creation for the last_login_ip field.
	user.DefaultLastLoginIP = userDescLastLoginIP.Default.(string)
	// user.LastLoginIPValidator is a validator for the "last_login_ip" field. It is called by the builders before save.
	user.LastLoginIPValidator = userDescLastLoginIP.Validators[0].(func(string) error)
	// userDescLastLoginTime is the schema descriptor for last_login_time field.
	userDescLastLoginTime := userFields[17].Descriptor()
	// user.DefaultLastLoginTime holds the default value on creation for the last_login_time field.
	user.DefaultLastLoginTime = userDescLastLoginTime.Default.(func() time.Time)
	// userDescLoginTime is the schema descriptor for login_time field.
	userDescLoginTime := userFields[18].Descriptor()
	// user.DefaultLoginTime holds the default value on creation for the login_time field.
	user.DefaultLoginTime = userDescLoginTime.Default.(func() time.Time)
	// userDescManagerID is the schema descriptor for manager_id field.
	userDescManagerID := userFields[20].Descriptor()
	// user.ManagerIDValidator is a validator for the "manager_id" field. It is called by the builders before save.
	user.ManagerIDValidator = userDescManagerID.Validators[0].(func(int64) error)
	// userDescManager is the schema descriptor for manager field.
	userDescManager := userFields[21].Descriptor()
	// user.DefaultManager holds the default value on creation for the manager field.
	user.DefaultManager = userDescManager.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() int64)
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(int64) error)
	userdepartmentFields := schema.UserDepartment{}.Fields()
	_ = userdepartmentFields
	// userdepartmentDescUserID is the schema descriptor for user_id field.
	userdepartmentDescUserID := userdepartmentFields[0].Descriptor()
	// userdepartment.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	userdepartment.UserIDValidator = userdepartmentDescUserID.Validators[0].(func(int64) error)
	// userdepartmentDescDepartmentID is the schema descriptor for department_id field.
	userdepartmentDescDepartmentID := userdepartmentFields[1].Descriptor()
	// userdepartment.DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	userdepartment.DepartmentIDValidator = userdepartmentDescDepartmentID.Validators[0].(func(int64) error)
	userpositionFields := schema.UserPosition{}.Fields()
	_ = userpositionFields
	// userpositionDescUserID is the schema descriptor for user_id field.
	userpositionDescUserID := userpositionFields[0].Descriptor()
	// userposition.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	userposition.UserIDValidator = userpositionDescUserID.Validators[0].(func(int64) error)
	// userpositionDescPositionID is the schema descriptor for position_id field.
	userpositionDescPositionID := userpositionFields[1].Descriptor()
	// userposition.PositionIDValidator is a validator for the "position_id" field. It is called by the builders before save.
	userposition.PositionIDValidator = userpositionDescPositionID.Validators[0].(func(int64) error)
	userroleFields := schema.UserRole{}.Fields()
	_ = userroleFields
	// userroleDescUserID is the schema descriptor for user_id field.
	userroleDescUserID := userroleFields[0].Descriptor()
	// userrole.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	userrole.UserIDValidator = userroleDescUserID.Validators[0].(func(int64) error)
	// userroleDescRoleID is the schema descriptor for role_id field.
	userroleDescRoleID := userroleFields[1].Descriptor()
	// userrole.RoleIDValidator is a validator for the "role_id" field. It is called by the builders before save.
	userrole.RoleIDValidator = userroleDescRoleID.Validators[0].(func(int64) error)
}

const (
	Version = "v0.14.4"                                         // Version of ent codegen.
	Sum     = "h1:/DhDraSLXIkBhyiVoJeSshr4ZYi7femzhj6/TckzZuI=" // Sum of ent codegen.
)
