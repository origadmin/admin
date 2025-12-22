package dto

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
)

type (
	Department                = ent.Department
	DepartmentEdges           = ent.DepartmentEdges
	DepartmentEdgesPB         = types.DepartmentEdges
	DepartmentPB              = types.Department
	Departments               = []*ent.Department
	DepartmentsPB             = []*types.Department
	MenuPB                    = types.Menu
	MenusPB                   = []*types.Menu
	Permission                = ent.Permission
	PermissionEdges           = ent.PermissionEdges
	PermissionEdgesPB         = types.PermissionEdges
	PermissionPB              = types.Permission
	PermissionResource        = ent.PermissionResource
	PermissionResourceEdges   = ent.PermissionResourceEdges
	PermissionResourceEdgesPB = types.PermissionResourceEdges
	PermissionResourcePB      = types.PermissionResource
	PermissionResources       = []*ent.PermissionResource
	PermissionResourcesPB     = []*types.PermissionResource
	Permissions               = []*ent.Permission
	PermissionsPB             = []*types.Permission
	Position                  = ent.Position
	PositionEdges             = ent.PositionEdges
	PositionEdgesPB           = types.PositionEdges
	PositionPB                = types.Position
	PositionPermission        = ent.PositionPermission
	PositionPermissionEdges   = ent.PositionPermissionEdges
	PositionPermissionEdgesPB = types.PositionPermissionEdges
	PositionPermissionPB      = types.PositionPermission
	PositionPermissions       = []*ent.PositionPermission
	PositionPermissionsPB     = []*types.PositionPermission
	Positions                 = []*ent.Position
	PositionsPB               = []*types.Position
	Resource                  = ent.Resource
	ResourceEdges             = ent.ResourceEdges
	ResourceEdgesPB           = types.ResourceEdges
	ResourcePB                = types.Resource
	Resources                 = []*ent.Resource
	ResourcesPB               = []*types.Resource
	Role                      = ent.Role
	RoleEdges                 = ent.RoleEdges
	RoleEdgesPB               = types.RoleEdges
	RoleMenuPB                = types.RoleMenu
	RoleMenusPB               = []*types.RoleMenu
	RolePB                    = types.Role
	RolePermission            = ent.RolePermission
	RolePermissionEdges       = ent.RolePermissionEdges
	RolePermissionEdgesPB     = types.RolePermissionEdges
	RolePermissionPB          = types.RolePermission
	RolePermissions           = []*ent.RolePermission
	RolePermissionsPB         = []*types.RolePermission
	Roles                     = []*ent.Role
	RolesPB                   = []*types.Role
	TimestampPB               = timestamppb.Timestamp
	User                      = ent.User
	UserDepartment            = ent.UserDepartment
	UserDepartmentEdges       = ent.UserDepartmentEdges
	UserDepartmentEdgesPB     = types.UserDepartmentEdges
	UserDepartmentPB          = types.UserDepartment
	UserDepartments           = []*ent.UserDepartment
	UserDepartmentsPB         = []*types.UserDepartment
	UserEdges                 = ent.UserEdges
	UserEdgesPB               = types.UserEdges
	UserPB                    = types.User
	UserPosition              = ent.UserPosition
	UserPositionEdges         = ent.UserPositionEdges
	UserPositionEdgesPB       = types.UserPositionEdges
	UserPositionPB            = types.UserPosition
	UserPositions             = []*ent.UserPosition
	UserPositionsPB           = []*types.UserPosition
	UserRole                  = ent.UserRole
	UserRoleEdges             = ent.UserRoleEdges
	UserRoleEdgesPB           = types.UserRoleEdges
	UserRolePB                = types.UserRole
	UserRoles                 = []*ent.UserRole
	UserRolesPB               = []*types.UserRole
	Users                     = []*ent.User
	UsersPB                   = []*types.User
)

func ConvertDepartmentEdgesPBToDepartmentEdges(from *DepartmentEdgesPB) *DepartmentEdges {
	if from == nil {
		return nil
	}

	to := &DepartmentEdges{
		Users:           ConvertUsersPBToUsers(from.Users),
		Positions:       ConvertPositionsPBToPositions(from.Positions),
		Children:        ConvertDepartmentsPBToDepartments(from.Children),
		Parent:          ConvertDepartmentPBToDepartment(from.Parent),
		UserDepartments: ConvertUserDepartmentsPBToUserDepartments(from.UserDepartments),
	}
	return to
}

func ConvertDepartmentEdgesToDepartmentEdgesPB(from *DepartmentEdges) *DepartmentEdgesPB {
	if from == nil {
		return nil
	}

	to := &DepartmentEdgesPB{
		Users:           ConvertUsersToUsersPB(from.Users),
		Positions:       ConvertPositionsToPositionsPB(from.Positions),
		Parent:          ConvertDepartmentToDepartmentPB(from.Parent),
		Children:        ConvertDepartmentsToDepartmentsPB(from.Children),
		UserDepartments: ConvertUserDepartmentsToUserDepartmentsPB(from.UserDepartments),
	}
	return to
}

func ConvertDepartmentPBToDepartment(from *DepartmentPB) *Department {
	if from == nil {
		return nil
	}

	to := &Department{
		ID:          from.Id,
		CreateTime:  ConvertTimestampToTime(from.CreateTime),
		UpdateTime:  ConvertTimestampToTime(from.UpdateTime),
		Keyword:     from.Keyword,
		Name:        from.Name,
		TreePath:    from.TreePath,
		Sequence:    int(from.Sequence),
		Status:      int8(from.Status),
		Level:       int(from.Level),
		Description: from.Description,
		ParentID:    from.ParentId,
	}
	return to
}

func ConvertDepartmentToDepartmentPB(from *Department) *DepartmentPB {
	if from == nil {
		return nil
	}

	to := &DepartmentPB{
		Id:          from.ID,
		CreateTime:  ConvertTimeToTimestamp(from.CreateTime),
		UpdateTime:  ConvertTimeToTimestamp(from.UpdateTime),
		Keyword:     from.Keyword,
		Name:        from.Name,
		TreePath:    from.TreePath,
		Sequence:    int32(from.Sequence),
		Status:      int32(from.Status),
		Level:       int32(from.Level),
		Description: from.Description,
		ParentId:    from.ParentID,
	}
	return to
}

func ConvertDepartmentsPBToDepartments(froms DepartmentsPB) Departments {
	if froms == nil {
		return nil
	}
	tos := make(Departments, len(froms))
	for i, f := range froms {
		tos[i] = ConvertDepartmentPBToDepartment(f)
	}
	return tos
}

func ConvertDepartmentsToDepartmentsPB(froms Departments) DepartmentsPB {
	if froms == nil {
		return nil
	}
	tos := make(DepartmentsPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertDepartmentToDepartmentPB(f)
	}
	return tos
}

func ConvertMenuPBToResource(from *MenuPB) *Resource {
	if from == nil {
		return nil
	}

	to := &Resource{
		ID:          from.Id,
		CreateTime:  ConvertTimestampToTime(from.CreateTime),
		UpdateTime:  ConvertTimestampToTime(from.UpdateTime),
		Keyword:     from.Keyword,
		Name:        from.Name,
		I18nKey:     from.I18NKey,
		Description: from.Description,
		Sequence:    int(from.Sequence),
		Type:        from.Type,
		Icon:        from.Icon,
		Path:        from.Path,
		Properties:  ConvertMenuPBPropertiesToResourceProperties(from.Properties),
		Status:      int8(from.Status),
		ParentID:    from.ParentId,
	}
	return to
}

func ConvertPermissionEdgesPBToPermissionEdges(from *PermissionEdgesPB) *PermissionEdges {
	if from == nil {
		return nil
	}

	to := &PermissionEdges{
		Roles:               ConvertRolesPBToRoles(from.Roles),
		Resources:           ConvertResourcesPBToResources(from.Resources),
		Positions:           ConvertPositionsPBToPositions(from.Positions),
		RolePermissions:     ConvertRolePermissionsPBToRolePermissions(from.RolePermissions),
		PermissionResources: ConvertPermissionResourcesPBToPermissionResources(from.PermissionResources),
		PositionPermissions: ConvertPositionPermissionsPBToPositionPermissions(from.PositionPermissions),
	}
	return to
}

func ConvertPermissionEdgesToPermissionEdgesPB(from *PermissionEdges) *PermissionEdgesPB {
	if from == nil {
		return nil
	}

	to := &PermissionEdgesPB{
		Roles:               ConvertRolesToRolesPB(from.Roles),
		Positions:           ConvertPositionsToPositionsPB(from.Positions),
		Resources:           ConvertResourcesToResourcesPB(from.Resources),
		RolePermissions:     ConvertRolePermissionsToRolePermissionsPB(from.RolePermissions),
		PositionPermissions: ConvertPositionPermissionsToPositionPermissionsPB(from.PositionPermissions),
		PermissionResources: ConvertPermissionResourcesToPermissionResourcesPB(from.PermissionResources),
	}
	return to
}

func ConvertPermissionPBToPermission(from *PermissionPB) *Permission {
	if from == nil {
		return nil
	}

	to := &Permission{
		ID:          from.Id,
		CreateTime:  ConvertTimestampToTime(from.CreateTime),
		UpdateTime:  ConvertTimestampToTime(from.UpdateTime),
		Name:        from.Name,
		Keyword:     from.Keyword,
		Description: from.Description,
		DataScope:   from.DataScope,
		DataRules:   from.DataRules,
	}
	return to
}

func ConvertPermissionResourceEdgesPBToPermissionResourceEdges(from *PermissionResourceEdgesPB) *PermissionResourceEdges {
	if from == nil {
		return nil
	}

	to := &PermissionResourceEdges{
		Permission: ConvertPermissionPBToPermission(from.Permission),
		Resource:   ConvertResourcePBToResource(from.Resource),
	}
	return to
}

func ConvertPermissionResourceEdgesToPermissionResourceEdgesPB(from *PermissionResourceEdges) *PermissionResourceEdgesPB {
	if from == nil {
		return nil
	}

	to := &PermissionResourceEdgesPB{
		Permission: ConvertPermissionToPermissionPB(from.Permission),
		Resource:   ConvertResourceToResourcePB(from.Resource),
	}
	return to
}

func ConvertPermissionResourcePBToPermissionResource(from *PermissionResourcePB) *PermissionResource {
	if from == nil {
		return nil
	}

	to := &PermissionResource{
		ID:           int(from.Id),
		PermissionID: from.PermissionId,
		ResourceID:   from.ResourceId,
	}
	return to
}

func ConvertPermissionResourceToPermissionResourcePB(from *PermissionResource) *PermissionResourcePB {
	if from == nil {
		return nil
	}

	to := &PermissionResourcePB{
		Id:           int64(from.ID),
		PermissionId: from.PermissionID,
		ResourceId:   from.ResourceID,
	}
	return to
}

func ConvertPermissionResourcesPBToPermissionResources(froms PermissionResourcesPB) PermissionResources {
	if froms == nil {
		return nil
	}
	tos := make(PermissionResources, len(froms))
	for i, f := range froms {
		tos[i] = ConvertPermissionResourcePBToPermissionResource(f)
	}
	return tos
}

func ConvertPermissionResourcesToPermissionResourcesPB(froms PermissionResources) PermissionResourcesPB {
	if froms == nil {
		return nil
	}
	tos := make(PermissionResourcesPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertPermissionResourceToPermissionResourcePB(f)
	}
	return tos
}

func ConvertPermissionToPermissionPB(from *Permission) *PermissionPB {
	if from == nil {
		return nil
	}

	to := &PermissionPB{
		Id:          from.ID,
		CreateTime:  ConvertTimeToTimestamp(from.CreateTime),
		UpdateTime:  ConvertTimeToTimestamp(from.UpdateTime),
		Name:        from.Name,
		Keyword:     from.Keyword,
		Description: from.Description,
		DataScope:   from.DataScope,
		DataRules:   from.DataRules,
	}
	return to
}

func ConvertPermissionsPBToPermissions(froms PermissionsPB) Permissions {
	if froms == nil {
		return nil
	}
	tos := make(Permissions, len(froms))
	for i, f := range froms {
		tos[i] = ConvertPermissionPBToPermission(f)
	}
	return tos
}

func ConvertPermissionsToPermissionsPB(froms Permissions) PermissionsPB {
	if froms == nil {
		return nil
	}
	tos := make(PermissionsPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertPermissionToPermissionPB(f)
	}
	return tos
}

func ConvertPositionEdgesPBToPositionEdges(from *PositionEdgesPB) *PositionEdges {
	if from == nil {
		return nil
	}

	to := &PositionEdges{
		Department:          ConvertDepartmentPBToDepartment(from.Department),
		Users:               ConvertUsersPBToUsers(from.Users),
		Permissions:         ConvertPermissionsPBToPermissions(from.Permissions),
		UserPositions:       ConvertUserPositionsPBToUserPositions(from.UserPositions),
		PositionPermissions: ConvertPositionPermissionsPBToPositionPermissions(from.PositionPermissions),
	}
	return to
}

func ConvertPositionEdgesToPositionEdgesPB(from *PositionEdges) *PositionEdgesPB {
	if from == nil {
		return nil
	}

	to := &PositionEdgesPB{
		Department:          ConvertDepartmentToDepartmentPB(from.Department),
		Users:               ConvertUsersToUsersPB(from.Users),
		Permissions:         ConvertPermissionsToPermissionsPB(from.Permissions),
		UserPositions:       ConvertUserPositionsToUserPositionsPB(from.UserPositions),
		PositionPermissions: ConvertPositionPermissionsToPositionPermissionsPB(from.PositionPermissions),
	}
	return to
}

func ConvertPositionPBToPosition(from *PositionPB) *Position {
	if from == nil {
		return nil
	}

	to := &Position{
		ID:           from.Id,
		CreateTime:   ConvertTimestampToTime(from.CreateTime),
		UpdateTime:   ConvertTimestampToTime(from.UpdateTime),
		Name:         from.Name,
		Keyword:      from.Keyword,
		Description:  from.Description,
		DepartmentID: from.DepartmentId,
	}
	return to
}

func ConvertPositionPermissionEdgesPBToPositionPermissionEdges(from *PositionPermissionEdgesPB) *PositionPermissionEdges {
	if from == nil {
		return nil
	}

	to := &PositionPermissionEdges{
		Position:   ConvertPositionPBToPosition(from.Position),
		Permission: ConvertPermissionPBToPermission(from.Permission),
	}
	return to
}

func ConvertPositionPermissionEdgesToPositionPermissionEdgesPB(from *PositionPermissionEdges) *PositionPermissionEdgesPB {
	if from == nil {
		return nil
	}

	to := &PositionPermissionEdgesPB{
		Position:   ConvertPositionToPositionPB(from.Position),
		Permission: ConvertPermissionToPermissionPB(from.Permission),
	}
	return to
}

func ConvertPositionPermissionPBToPositionPermission(from *PositionPermissionPB) *PositionPermission {
	if from == nil {
		return nil
	}

	to := &PositionPermission{
		ID:           int(from.Id),
		PositionID:   from.PositionId,
		PermissionID: from.PermissionId,
	}
	return to
}

func ConvertPositionPermissionToPositionPermissionPB(from *PositionPermission) *PositionPermissionPB {
	if from == nil {
		return nil
	}

	to := &PositionPermissionPB{
		Id:           int64(from.ID),
		PositionId:   from.PositionID,
		PermissionId: from.PermissionID,
	}
	return to
}

func ConvertPositionPermissionsPBToPositionPermissions(froms PositionPermissionsPB) PositionPermissions {
	if froms == nil {
		return nil
	}
	tos := make(PositionPermissions, len(froms))
	for i, f := range froms {
		tos[i] = ConvertPositionPermissionPBToPositionPermission(f)
	}
	return tos
}

func ConvertPositionPermissionsToPositionPermissionsPB(froms PositionPermissions) PositionPermissionsPB {
	if froms == nil {
		return nil
	}
	tos := make(PositionPermissionsPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertPositionPermissionToPositionPermissionPB(f)
	}
	return tos
}

func ConvertPositionToPositionPB(from *Position) *PositionPB {
	if from == nil {
		return nil
	}

	to := &PositionPB{
		Id:           from.ID,
		CreateTime:   ConvertTimeToTimestamp(from.CreateTime),
		UpdateTime:   ConvertTimeToTimestamp(from.UpdateTime),
		Name:         from.Name,
		Keyword:      from.Keyword,
		Description:  from.Description,
		DepartmentId: from.DepartmentID,
	}
	return to
}

func ConvertPositionsPBToPositions(froms PositionsPB) Positions {
	if froms == nil {
		return nil
	}
	tos := make(Positions, len(froms))
	for i, f := range froms {
		tos[i] = ConvertPositionPBToPosition(f)
	}
	return tos
}

func ConvertPositionsToPositionsPB(froms Positions) PositionsPB {
	if froms == nil {
		return nil
	}
	tos := make(PositionsPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertPositionToPositionPB(f)
	}
	return tos
}

func ConvertResourceEdgesPBToResourceEdges(from *ResourceEdgesPB) *ResourceEdges {
	if from == nil {
		return nil
	}

	to := &ResourceEdges{}
	return to
}

func ConvertResourceEdgesToResourceEdgesPB(from *ResourceEdges) *ResourceEdgesPB {
	if from == nil {
		return nil
	}

	to := &ResourceEdgesPB{}
	return to
}

func ConvertResourcePBToResource(from *ResourcePB) *Resource {
	if from == nil {
		return nil
	}

	to := &Resource{
		ID:          from.Id,
		CreateTime:  ConvertTimestampToTime(from.CreateTime),
		UpdateTime:  ConvertTimestampToTime(from.UpdateTime),
		Name:        from.Name,
		Keyword:     from.Keyword,
		I18nKey:     from.I18NKey,
		Type:        from.Type,
		Status:      int8(from.Status),
		Path:        from.Path,
		Operation:   from.Operation,
		Method:      from.Method,
		Component:   from.Component,
		Icon:        from.Icon,
		Sequence:    int(from.Sequence),
		Visible:     from.Visible,
		TreePath:    from.TreePath,
		Properties:  from.Properties,
		Description: from.Description,
		ParentID:    from.ParentId,
	}
	return to
}

func ConvertResourceToMenuPB(from *Resource) *MenuPB {
	if from == nil {
		return nil
	}

	to := &MenuPB{
		Id:          from.ID,
		CreateTime:  ConvertTimeToTimestamp(from.CreateTime),
		UpdateTime:  ConvertTimeToTimestamp(from.UpdateTime),
		Name:        from.Name,
		Keyword:     from.Keyword,
		I18NKey:     from.I18nKey,
		Type:        from.Type,
		Status:      int32(from.Status),
		Path:        from.Path,
		Icon:        from.Icon,
		Sequence:    int32(from.Sequence),
		Properties:  ConvertResourcePropertiesToMenuPBProperties(from.Properties),
		Description: from.Description,
		ParentId:    from.ParentID,
	}
	return to
}

func ConvertResourceToResourcePB(from *Resource) *ResourcePB {
	if from == nil {
		return nil
	}

	to := &ResourcePB{
		Id:          from.ID,
		CreateTime:  ConvertTimeToTimestamp(from.CreateTime),
		UpdateTime:  ConvertTimeToTimestamp(from.UpdateTime),
		Name:        from.Name,
		Keyword:     from.Keyword,
		I18NKey:     from.I18nKey,
		Type:        from.Type,
		Status:      int32(from.Status),
		Path:        from.Path,
		Operation:   from.Operation,
		Method:      from.Method,
		Component:   from.Component,
		Icon:        from.Icon,
		Sequence:    int32(from.Sequence),
		Visible:     from.Visible,
		TreePath:    from.TreePath,
		Properties:  from.Properties,
		Description: from.Description,
		ParentId:    from.ParentID,
	}
	return to
}

func ConvertResourcesPBToResources(froms ResourcesPB) Resources {
	if froms == nil {
		return nil
	}
	tos := make(Resources, len(froms))
	for i, f := range froms {
		tos[i] = ConvertResourcePBToResource(f)
	}
	return tos
}

func ConvertResourcesToResourcesPB(froms Resources) ResourcesPB {
	if froms == nil {
		return nil
	}
	tos := make(ResourcesPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertResourceToResourcePB(f)
	}
	return tos
}

func ConvertRoleEdgesPBToRoleEdges(from *RoleEdgesPB) *RoleEdges {
	if from == nil {
		return nil
	}

	to := &RoleEdges{
		Users:     ConvertUsersPBToUsers(from.Users),
		UserRoles: ConvertUserRolesPBToUserRoles(from.UserRoles),
	}
	return to
}

func ConvertRoleEdgesToRoleEdgesPB(from *RoleEdges) *RoleEdgesPB {
	if from == nil {
		return nil
	}

	to := &RoleEdgesPB{
		Users:     ConvertUsersToUsersPB(from.Users),
		UserRoles: ConvertUserRolesToUserRolesPB(from.UserRoles),
	}
	return to
}

func ConvertRolePBToRole(from *RolePB) *Role {
	if from == nil {
		return nil
	}

	to := &Role{
		ID:          from.Id,
		CreateTime:  ConvertTimestampToTime(from.CreateTime),
		UpdateTime:  ConvertTimestampToTime(from.UpdateTime),
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		Type:        int8(from.Type),
		Sequence:    int(from.Sequence),
		Status:      int8(from.Status),
	}
	return to
}

func ConvertRolePermissionEdgesPBToRolePermissionEdges(from *RolePermissionEdgesPB) *RolePermissionEdges {
	if from == nil {
		return nil
	}

	to := &RolePermissionEdges{
		Role:       ConvertRolePBToRole(from.Role),
		Permission: ConvertPermissionPBToPermission(from.Permission),
	}
	return to
}

func ConvertRolePermissionEdgesToRolePermissionEdgesPB(from *RolePermissionEdges) *RolePermissionEdgesPB {
	if from == nil {
		return nil
	}

	to := &RolePermissionEdgesPB{
		Role:       ConvertRoleToRolePB(from.Role),
		Permission: ConvertPermissionToPermissionPB(from.Permission),
	}
	return to
}

func ConvertRolePermissionPBToRolePermission(from *RolePermissionPB) *RolePermission {
	if from == nil {
		return nil
	}

	to := &RolePermission{
		ID:           int(from.Id),
		RoleID:       from.RoleId,
		PermissionID: from.PermissionId,
	}
	return to
}

func ConvertRolePermissionToRolePermissionPB(from *RolePermission) *RolePermissionPB {
	if from == nil {
		return nil
	}

	to := &RolePermissionPB{
		Id:           int64(from.ID),
		RoleId:       from.RoleID,
		PermissionId: from.PermissionID,
	}
	return to
}

func ConvertRolePermissionsPBToRolePermissions(froms RolePermissionsPB) RolePermissions {
	if froms == nil {
		return nil
	}
	tos := make(RolePermissions, len(froms))
	for i, f := range froms {
		tos[i] = ConvertRolePermissionPBToRolePermission(f)
	}
	return tos
}

func ConvertRolePermissionsToRolePermissionsPB(froms RolePermissions) RolePermissionsPB {
	if froms == nil {
		return nil
	}
	tos := make(RolePermissionsPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertRolePermissionToRolePermissionPB(f)
	}
	return tos
}

func ConvertRoleToRolePB(from *Role) *RolePB {
	if from == nil {
		return nil
	}

	to := &RolePB{
		Id:          from.ID,
		CreateTime:  ConvertTimeToTimestamp(from.CreateTime),
		UpdateTime:  ConvertTimeToTimestamp(from.UpdateTime),
		Keyword:     from.Keyword,
		Name:        from.Name,
		Description: from.Description,
		Type:        int32(from.Type),
		Sequence:    int32(from.Sequence),
		Status:      int32(from.Status),
	}
	return to
}

func ConvertRolesPBToRoles(froms RolesPB) Roles {
	if froms == nil {
		return nil
	}
	tos := make(Roles, len(froms))
	for i, f := range froms {
		tos[i] = ConvertRolePBToRole(f)
	}
	return tos
}

func ConvertRolesToRolesPB(froms Roles) RolesPB {
	if froms == nil {
		return nil
	}
	tos := make(RolesPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertRoleToRolePB(f)
	}
	return tos
}

func ConvertUserDepartmentEdgesPBToUserDepartmentEdges(from *UserDepartmentEdgesPB) *UserDepartmentEdges {
	if from == nil {
		return nil
	}

	to := &UserDepartmentEdges{
		User:       ConvertUserPBToUser(from.User),
		Department: ConvertDepartmentPBToDepartment(from.Department),
	}
	return to
}

func ConvertUserDepartmentEdgesToUserDepartmentEdgesPB(from *UserDepartmentEdges) *UserDepartmentEdgesPB {
	if from == nil {
		return nil
	}

	to := &UserDepartmentEdgesPB{
		User:       ConvertUserToUserPB(from.User),
		Department: ConvertDepartmentToDepartmentPB(from.Department),
	}
	return to
}

func ConvertUserDepartmentPBToUserDepartment(from *UserDepartmentPB) *UserDepartment {
	if from == nil {
		return nil
	}

	to := &UserDepartment{
		ID:           int(from.Id),
		UserID:       from.UserId,
		DepartmentID: from.DepartmentId,
		Edges:        *ConvertUserDepartmentEdgesPBToUserDepartmentEdges(from.Edges),
	}
	return to
}

func ConvertUserDepartmentToUserDepartmentPB(from *UserDepartment) *UserDepartmentPB {
	if from == nil {
		return nil
	}

	to := &UserDepartmentPB{
		Id:           int64(from.ID),
		UserId:       from.UserID,
		DepartmentId: from.DepartmentID,
		Edges:        ConvertUserDepartmentEdgesToUserDepartmentEdgesPB(&from.Edges),
	}
	return to
}

func ConvertUserDepartmentsPBToUserDepartments(froms UserDepartmentsPB) UserDepartments {
	if froms == nil {
		return nil
	}
	tos := make(UserDepartments, len(froms))
	for i, f := range froms {
		tos[i] = ConvertUserDepartmentPBToUserDepartment(f)
	}
	return tos
}

func ConvertUserDepartmentsToUserDepartmentsPB(froms UserDepartments) UserDepartmentsPB {
	if froms == nil {
		return nil
	}
	tos := make(UserDepartmentsPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertUserDepartmentToUserDepartmentPB(f)
	}
	return tos
}

func ConvertUserEdgesPBToUserEdges(from *UserEdgesPB) *UserEdges {
	if from == nil {
		return nil
	}

	to := &UserEdges{
		Roles:     ConvertRolesPBToRoles(from.Roles),
		UserRoles: ConvertUserRolesPBToUserRoles(from.UserRoles),
	}
	return to
}

func ConvertUserEdgesToUserEdgesPB(from *UserEdges) *UserEdgesPB {
	if from == nil {
		return nil
	}

	to := &UserEdgesPB{
		Roles:     ConvertRolesToRolesPB(from.Roles),
		UserRoles: ConvertUserRolesToUserRolesPB(from.UserRoles),
	}
	return to
}

func ConvertUserPBToUser(from *UserPB) *User {
	if from == nil {
		return nil
	}

	to := &User{
		ID:            from.Id,
		CreateAuthor:  from.CreateAuthor,
		UpdateAuthor:  from.UpdateAuthor,
		CreateTime:    ConvertTimestampToTime(from.CreateTime),
		UpdateTime:    ConvertTimestampToTime(from.UpdateTime),
		UUID:          from.Uuid,
		AllowedIP:     from.AllowedIp,
		Username:      from.Username,
		Nickname:      from.Nickname,
		Avatar:        from.Avatar,
		Name:          from.Name,
		Gender:        ConvertUserPBGenderToUserGender(from.Gender),
		Salt:          from.Salt,
		Phone:         from.Phone,
		Email:         from.Email,
		Remark:        from.Remark,
		Token:         from.Token,
		Status:        int8(from.Status),
		LastLoginIP:   from.LastLoginIp,
		LastLoginTime: ConvertTimestampToTime(from.LastLoginTime),
		SanctionDate:  ConvertTimestampToTime(from.SanctionDate),
		ManagerID:     from.ManagerId,
		Manager:       from.Manager,
	}
	return to
}

func ConvertUserPositionEdgesPBToUserPositionEdges(from *UserPositionEdgesPB) *UserPositionEdges {
	if from == nil {
		return nil
	}

	to := &UserPositionEdges{
		User:     ConvertUserPBToUser(from.User),
		Position: ConvertPositionPBToPosition(from.Position),
	}
	return to
}

func ConvertUserPositionEdgesToUserPositionEdgesPB(from *UserPositionEdges) *UserPositionEdgesPB {
	if from == nil {
		return nil
	}

	to := &UserPositionEdgesPB{
		User:     ConvertUserToUserPB(from.User),
		Position: ConvertPositionToPositionPB(from.Position),
	}
	return to
}

func ConvertUserPositionPBToUserPosition(from *UserPositionPB) *UserPosition {
	if from == nil {
		return nil
	}

	to := &UserPosition{
		ID:         int(from.Id),
		UserID:     from.UserId,
		PositionID: from.PositionId,
	}
	return to
}

func ConvertUserPositionToUserPositionPB(from *UserPosition) *UserPositionPB {
	if from == nil {
		return nil
	}

	to := &UserPositionPB{
		Id:         int64(from.ID),
		UserId:     from.UserID,
		PositionId: from.PositionID,
	}
	return to
}

func ConvertUserPositionsPBToUserPositions(froms UserPositionsPB) UserPositions {
	if froms == nil {
		return nil
	}
	tos := make(UserPositions, len(froms))
	for i, f := range froms {
		tos[i] = ConvertUserPositionPBToUserPosition(f)
	}
	return tos
}

func ConvertUserPositionsToUserPositionsPB(froms UserPositions) UserPositionsPB {
	if froms == nil {
		return nil
	}
	tos := make(UserPositionsPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertUserPositionToUserPositionPB(f)
	}
	return tos
}

func ConvertUserRoleEdgesPBToUserRoleEdges(from *UserRoleEdgesPB) *UserRoleEdges {
	if from == nil {
		return nil
	}

	to := &UserRoleEdges{
		User: ConvertUserPBToUser(from.User),
		Role: ConvertRolePBToRole(from.Role),
	}
	return to
}

func ConvertUserRoleEdgesToUserRoleEdgesPB(from *UserRoleEdges) *UserRoleEdgesPB {
	if from == nil {
		return nil
	}

	to := &UserRoleEdgesPB{
		User: ConvertUserToUserPB(from.User),
		Role: ConvertRoleToRolePB(from.Role),
	}
	return to
}

func ConvertUserRolePBToUserRole(from *UserRolePB) *UserRole {
	if from == nil {
		return nil
	}

	to := &UserRole{
		ID:     int(from.Id),
		UserID: from.UserId,
		RoleID: from.RoleId,
	}
	return to
}

func ConvertUserRoleToUserRolePB(from *UserRole) *UserRolePB {
	if from == nil {
		return nil
	}

	to := &UserRolePB{
		Id:     int64(from.ID),
		UserId: from.UserID,
		RoleId: from.RoleID,
	}
	return to
}

func ConvertUserRolesPBToUserRoles(froms UserRolesPB) UserRoles {
	if froms == nil {
		return nil
	}
	tos := make(UserRoles, len(froms))
	for i, f := range froms {
		tos[i] = ConvertUserRolePBToUserRole(f)
	}
	return tos
}

func ConvertUserRolesToUserRolesPB(froms UserRoles) UserRolesPB {
	if froms == nil {
		return nil
	}
	tos := make(UserRolesPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertUserRoleToUserRolePB(f)
	}
	return tos
}

func ConvertUserToUserPB(from *User) *UserPB {
	if from == nil {
		return nil
	}

	to := &UserPB{
		Id:            from.ID,
		CreateAuthor:  from.CreateAuthor,
		UpdateAuthor:  from.UpdateAuthor,
		CreateTime:    ConvertTimeToTimestamp(from.CreateTime),
		UpdateTime:    ConvertTimeToTimestamp(from.UpdateTime),
		Uuid:          from.UUID,
		AllowedIp:     from.AllowedIP,
		Username:      from.Username,
		Nickname:      from.Nickname,
		Avatar:        from.Avatar,
		Name:          from.Name,
		Gender:        ConvertUserGenderToUserPBGender(from.Gender),
		Salt:          from.Salt,
		Phone:         from.Phone,
		Email:         from.Email,
		Remark:        from.Remark,
		Token:         from.Token,
		Status:        int32(from.Status),
		LastLoginIp:   from.LastLoginIP,
		LastLoginTime: ConvertTimeToTimestamp(from.LastLoginTime),
		SanctionDate:  ConvertTimeToTimestamp(from.SanctionDate),
		ManagerId:     from.ManagerID,
		Manager:       from.Manager,
	}
	return to
}

func ConvertUsersPBToUsers(froms UsersPB) Users {
	if froms == nil {
		return nil
	}
	tos := make(Users, len(froms))
	for i, f := range froms {
		tos[i] = ConvertUserPBToUser(f)
	}
	return tos
}

func ConvertUsersToUsersPB(froms Users) UsersPB {
	if froms == nil {
		return nil
	}
	tos := make(UsersPB, len(froms))
	for i, f := range froms {
		tos[i] = ConvertUserToUserPB(f)
	}
	return tos
}

func ConvertTimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}
func ConvertTimestampToTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}
