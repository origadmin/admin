/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"net/http"

	"github.com/origadmin/toolkits/errors/httperr"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/schema"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound      = httperr.New("http.response.status."+pb.SystemErrorReason_SYSTEM_ERROR_REASON_USER_NOT_FOUND.String(), http.StatusNotFound, "user not found")
	ErrInvalidCaptchaID  = httperr.New("http.response.status."+pb.SystemErrorReason_SYSTEM_ERROR_REASON_INVALID_CAPTCHA_ID.String(), http.StatusBadRequest, "invalid captcha id")
	ErrInvalidPassword   = httperr.New("http.response.status."+pb.SystemErrorReason_SYSTEM_ERROR_REASON_INVALID_PASSWORD.String(), http.StatusBadRequest, "invalid password")
	ErrInvalidUsername   = httperr.New("http.response.status."+pb.SystemErrorReason_SYSTEM_ERROR_REASON_INVALID_USERNAME.String(), http.StatusBadRequest, "invalid username")
	ErrCaptchaIDNotFound = httperr.New("http.response.status."+pb.SystemErrorReason_SYSTEM_ERROR_REASON_CAPTCHA_ID_NOT_FOUND.String(), http.StatusBadRequest, "captcha id not found")
)

const (
	UserStatusEnabled  = schema.UserStatusEnabled
	UserStatusDisabled = schema.UserStatusDisabled

	ResourceStatusEnabled  = schema.ResourceStatusEnabled
	ResourceStatusDisabled = schema.ResourceStatusDisabled
)

type (
	User   = ent.User
	UserPB = pb.User
)

// ConvertUser2PB user.table.comment
func ConvertUser2PB(goModel *User) (pbModel *UserPB) {
	pbModel = &UserPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = int64(goModel.ID)
	pbModel.CreateAuthor = int64(goModel.CreateAuthor)
	pbModel.UpdateAuthor = int64(goModel.UpdateAuthor)
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Uuid = goModel.UUID
	pbModel.AllowedIp = goModel.AllowedIP
	pbModel.Username = goModel.Username
	pbModel.Nickname = goModel.Nickname
	pbModel.Avatar = goModel.Avatar
	pbModel.Name = goModel.Name
	pbModel.Gender = ConvertGender2PB(goModel.Gender)
	pbModel.Password = goModel.Password
	pbModel.Salt = goModel.Salt
	pbModel.Phone = goModel.Phone
	pbModel.Email = goModel.Email
	pbModel.Remark = goModel.Remark
	pbModel.Token = goModel.Token
	pbModel.Status = int32(goModel.Status)
	pbModel.LastLoginIp = goModel.LastLoginIP
	pbModel.LastLoginTime = timestamppb.New(goModel.LastLoginTime)
	pbModel.SanctionDate = timestamppb.New(goModel.SanctionDate)
	pbModel.ManagerId = int64(goModel.ManagerID)
	pbModel.Manager = goModel.Manager
	return pbModel
}

func ConvertGender2PB(gender user.Gender) string {
	return gender.String()
}

// ConvertUserPB2Object user.table.comment
func ConvertUserPB2Object(pbModel *UserPB) (goModel *User) {
	goModel = &User{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = int64(pbModel.Id)
	goModel.CreateAuthor = int64(pbModel.CreateAuthor)
	goModel.UpdateAuthor = int64(pbModel.UpdateAuthor)
	goModel.CreateTime = pbModel.CreateTime.AsTime()
	goModel.UpdateTime = pbModel.UpdateTime.AsTime()
	goModel.UUID = pbModel.Uuid
	goModel.AllowedIP = pbModel.AllowedIp
	goModel.Username = pbModel.Username
	goModel.Nickname = pbModel.Nickname
	goModel.Avatar = pbModel.Avatar
	goModel.Name = pbModel.Name
	goModel.Gender = user.Gender(pbModel.Gender)
	goModel.Password = pbModel.Password
	goModel.Salt = pbModel.Salt
	goModel.Phone = pbModel.Phone
	goModel.Email = pbModel.Email
	goModel.Remark = pbModel.Remark
	goModel.Token = pbModel.Token
	goModel.Status = int8(pbModel.Status)
	goModel.LastLoginIP = pbModel.LastLoginIp
	goModel.LastLoginTime = pbModel.LastLoginTime.AsTime()
	goModel.SanctionDate = pbModel.SanctionDate.AsTime()
	goModel.ManagerID = int64(pbModel.ManagerId)
	goModel.Manager = pbModel.Manager
	return goModel
}

type (
	Resource   = ent.Resource
	ResourcePB = pb.Resource
)

func ConvertResource2PB(goModel *Resource) (pbModel *ResourcePB) {
	pbModel = &ResourcePB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = int64(goModel.ID)
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Name = goModel.Name
	pbModel.Keyword = goModel.Keyword
	pbModel.I18NKey = goModel.I18nKey
	pbModel.Type = goModel.Type
	pbModel.Status = int32(goModel.Status)
	pbModel.Uri = goModel.URI
	pbModel.Operation = goModel.Operation
	pbModel.Method = goModel.Method
	pbModel.Component = goModel.Component
	pbModel.Icon = goModel.Icon
	pbModel.Sequence = int32(goModel.Sequence)
	pbModel.Visible = goModel.Visible
	pbModel.TreePath = goModel.TreePath
	pbModel.Properties = goModel.Properties
	pbModel.Description = goModel.Description
	pbModel.ParentId = int64(goModel.ParentID)
	return pbModel
}

func ConvertResourcePB2Object(pbModel *ResourcePB) (goModel *Resource) {
	goModel = &Resource{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = pbModel.Id
	goModel.CreateTime = pbModel.CreateTime.AsTime()
	goModel.UpdateTime = pbModel.UpdateTime.AsTime()
	goModel.Name = pbModel.Name
	goModel.Keyword = pbModel.Keyword
	goModel.I18nKey = pbModel.I18NKey
	goModel.Type = pbModel.Type
	goModel.Status = int8(pbModel.Status)
	goModel.URI = pbModel.Uri
	goModel.Operation = pbModel.Operation
	goModel.Method = pbModel.Method
	goModel.Component = pbModel.Component
	goModel.Icon = pbModel.Icon
	goModel.Sequence = int(pbModel.Sequence)
	goModel.Visible = pbModel.Visible
	goModel.TreePath = pbModel.TreePath
	goModel.Properties = pbModel.Properties
	goModel.Description = pbModel.Description
	goModel.ParentID = pbModel.ParentId
	return goModel
}

type (
	Role   = ent.Role
	RolePB = pb.Role
)

// ConvertRole2PB role.table.comment
func ConvertRole2PB(goModel *Role) (pbModel *RolePB) {
	pbModel = &RolePB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = goModel.ID
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Keyword = goModel.Keyword
	pbModel.Name = goModel.Name
	pbModel.Description = goModel.Description
	pbModel.Type = int32(goModel.Type)
	pbModel.Sequence = int32(goModel.Sequence)
	pbModel.Status = int32(goModel.Status)

	//pbModel.IsSystem = goModel.IsSystem
	return pbModel
}

// ConvertRolePB2Object role.table.comment
func ConvertRolePB2Object(pbModel *RolePB) (goModel *Role) {
	goModel = &Role{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = pbModel.Id
	goModel.CreateTime = pbModel.CreateTime.AsTime()
	goModel.UpdateTime = pbModel.UpdateTime.AsTime()
	goModel.Keyword = pbModel.Keyword
	goModel.Name = pbModel.Name
	goModel.Description = pbModel.Description
	goModel.Type = int8(pbModel.Type)
	goModel.Sequence = int(pbModel.Sequence)
	goModel.Status = int8(pbModel.Status)

	//goModel.IsSystem = pbModel.IsSystem
	return goModel
}

type (
	Department   = ent.Department
	DepartmentPB = pb.Department
)

// ConvertDepartment2PB department.table.comment
func ConvertDepartment2PB(goModel *Department) (pbModel *DepartmentPB) {
	pbModel = &DepartmentPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = goModel.ID
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Keyword = goModel.Keyword
	pbModel.Name = goModel.Name
	pbModel.Description = goModel.Description
	pbModel.Sequence = int32(goModel.Sequence)
	pbModel.Status = int32(goModel.Status)
	pbModel.Level = int32(goModel.Level)
	pbModel.ParentId = goModel.ParentID
	return pbModel
}

// ConvertDepartmentPB2Object department.table.comment
func ConvertDepartmentPB2Object(pbModel *DepartmentPB) (goModel *Department) {
	goModel = &Department{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = pbModel.Id
	goModel.CreateTime = pbModel.CreateTime.AsTime()
	goModel.UpdateTime = pbModel.UpdateTime.AsTime()
	goModel.Keyword = pbModel.Keyword
	goModel.Name = pbModel.Name
	goModel.Description = pbModel.Description
	goModel.Sequence = int(pbModel.Sequence)
	goModel.Status = int8(pbModel.Status)
	goModel.Level = int(pbModel.Level)
	goModel.ParentID = pbModel.ParentId
	return goModel
}

type (
	Position   = ent.Position
	PositionPB = pb.Position
)

// ConvertPosition2PB position.table.comment
func ConvertPosition2PB(goModel *Position) (pbModel *PositionPB) {
	pbModel = &PositionPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = int64(goModel.ID)
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Name = goModel.Name
	pbModel.Description = goModel.Description
	pbModel.DepartmentId = int64(goModel.DepartmentID)
	return pbModel
}

// ConvertPositionPB2Object position.table.comment
func ConvertPositionPB2Object(pbModel *PositionPB) (goModel *Position) {
	goModel = &Position{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = int64(pbModel.Id)
	goModel.CreateTime = pbModel.CreateTime.AsTime()
	goModel.UpdateTime = pbModel.UpdateTime.AsTime()
	goModel.Name = pbModel.Name
	goModel.Description = pbModel.Description
	goModel.DepartmentID = int64(pbModel.DepartmentId)
	return goModel
}

type (
	Users   = []*ent.User
	UsersPB = []*pb.User
)

// ConvertUsers2PB Users holds the value of the users edge.
func ConvertUsers2PB(gosModel Users) (pbsModel UsersPB) {
	for _, model := range gosModel {
		pbsModel = append(pbsModel, ConvertUser2PB(model))
	}
	return pbsModel
}

// ConvertUsersPB2Object Users holds the value of the users edge.
func ConvertUsersPB2Object(pbsModel UsersPB) (gosModel Users) {
	for _, model := range pbsModel {
		gosModel = append(gosModel, ConvertUserPB2Object(model))
	}
	return gosModel
}

type (
	Permissions   = []*ent.Permission
	PermissionsPB = []*pb.Permission
)

// ConvertPermissions2PB Permissions holds the value of the permissions edge.
func ConvertPermissions2PB(gosModel Permissions) (pbsModel PermissionsPB) {
	for _, model := range gosModel {
		pbsModel = append(pbsModel, ConvertPermission2PB(model))
	}
	return pbsModel
}

// ConvertPermissionsPB2Object Permissions holds the value of the permissions edge.
func ConvertPermissionsPB2Object(pbsModel PermissionsPB) (gosModel Permissions) {
	for _, model := range pbsModel {
		gosModel = append(gosModel, ConvertPermissionPB2Object(model))
	}
	return gosModel
}

type (
	UserPositions   = []*ent.UserPosition
	UserPositionsPB = []*pb.UserPosition
)

// ConvertUserPositions2PB UserPositions holds the value of the user_positions edge.
func ConvertUserPositions2PB(gosModel UserPositions) (pbsModel UserPositionsPB) {
	for _, model := range gosModel {
		pbsModel = append(pbsModel, ConvertUserPosition2PB(model))
	}
	return pbsModel
}

// ConvertUserPositionsPB2Object UserPositions holds the value of the user_positions edge.
func ConvertUserPositionsPB2Object(pbsModel UserPositionsPB) (gosModel UserPositions) {
	for _, model := range pbsModel {
		gosModel = append(gosModel, ConvertUserPositionPB2Object(model))
	}
	return gosModel
}

type (
	PositionEdges   = ent.PositionEdges
	PositionEdgesPB = pb.PositionEdges
)

// ConvertPositionEdges2PB PositionEdges holds the relations/edges for other nodes in the graph.
func ConvertPositionEdges2PB(goModel *PositionEdges) (pbModel *PositionEdgesPB) {
	pbModel = &PositionEdgesPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Department = ConvertDepartment2PB(goModel.Department)
	pbModel.Users = ConvertUsers2PB(goModel.Users)
	pbModel.Permissions = ConvertPermissions2PB(goModel.Permissions)
	pbModel.UserPositions = ConvertUserPositions2PB(goModel.UserPositions)
	pbModel.PositionPermissions = ConvertPositionPermissions2PB(goModel.PositionPermissions)
	return pbModel
}

// ConvertPositionEdgesPB2Object PositionEdges holds the relations/edges for other nodes in the graph.
func ConvertPositionEdgesPB2Object(pbModel *PositionEdgesPB) (goModel *PositionEdges) {
	goModel = &PositionEdges{}
	if pbModel == nil {
		return goModel
	}

	goModel.Department = ConvertDepartmentPB2Object(pbModel.Department)
	goModel.Users = ConvertUsersPB2Object(pbModel.Users)
	goModel.Permissions = ConvertPermissionsPB2Object(pbModel.Permissions)
	goModel.UserPositions = ConvertUserPositionsPB2Object(pbModel.UserPositions)
	goModel.PositionPermissions = ConvertPositionPermissionsPB2Object(pbModel.PositionPermissions)
	return goModel
}

// ConvertDataRules2PB permission.field.data_rules
func ConvertDataRules2PB(gosModel map[string]string) map[string]string {
	return gosModel
}

// ConvertDataRulesPB2Object permission.field.data_rules
func ConvertDataRulesPB2Object(pbsModel map[string]string) map[string]string {
	return pbsModel
}

type (
	Permission   = ent.Permission
	PermissionPB = pb.Permission
)

// ConvertPermission2PB permission.table.comment
func ConvertPermission2PB(goModel *Permission) (pbModel *PermissionPB) {
	pbModel = &PermissionPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = int64(goModel.ID)
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Name = goModel.Name
	pbModel.Keyword = goModel.Keyword
	pbModel.Description = goModel.Description
	pbModel.DataScope = goModel.DataScope
	pbModel.DataRules = ConvertDataRules2PB(goModel.DataRules)
	return pbModel
}

// ConvertPermissionPB2Object permission.table.comment
func ConvertPermissionPB2Object(pbModel *PermissionPB) (goModel *Permission) {
	goModel = &Permission{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = int64(pbModel.Id)
	goModel.CreateTime = pbModel.CreateTime.AsTime()
	goModel.UpdateTime = pbModel.UpdateTime.AsTime()
	goModel.Name = pbModel.Name
	goModel.Keyword = pbModel.Keyword
	goModel.Description = pbModel.Description
	goModel.DataScope = pbModel.DataScope
	goModel.DataRules = ConvertDataRulesPB2Object(pbModel.DataRules)
	return goModel
}

type (
	Roles   = []*ent.Role
	RolesPB = []*pb.Role
)

// ConvertRoles2PB Roles holds the value of the roles edge.
func ConvertRoles2PB(gosModel Roles) (pbsModel RolesPB) {
	for _, model := range gosModel {
		pbsModel = append(pbsModel, ConvertRole2PB(model))
	}
	return pbsModel
}

// ConvertRolesPB2Object Roles holds the value of the roles edge.
func ConvertRolesPB2Object(pbsModel RolesPB) (gosModel Roles) {
	for _, model := range pbsModel {
		gosModel = append(gosModel, ConvertRolePB2Object(model))
	}
	return gosModel
}

type (
	Resources   = []*ent.Resource
	ResourcesPB = []*pb.Resource
)

// ConvertResources2PB Resources holds the value of the resources edge.
func ConvertResources2PB(gosModel Resources) (pbsModel ResourcesPB) {
	for _, model := range gosModel {
		pbsModel = append(pbsModel, ConvertResource2PB(model))
	}
	return pbsModel
}

// ConvertResourcesPB2Object Resources holds the value of the resources edge.
func ConvertResourcesPB2Object(pbsModel ResourcesPB) (gosModel Resources) {
	for _, model := range pbsModel {
		gosModel = append(gosModel, ConvertResourcePB2Object(model))
	}
	return gosModel
}

type (
	Positions   = []*ent.Position
	PositionsPB = []*pb.Position
)

// ConvertPositions2PB Positions holds the value of the positions edge.
func ConvertPositions2PB(gosModel Positions) (pbsModel PositionsPB) {
	for _, model := range gosModel {
		pbsModel = append(pbsModel, ConvertPosition2PB(model))
	}
	return pbsModel
}

// ConvertPositionsPB2Object Positions holds the value of the positions edge.
func ConvertPositionsPB2Object(pbsModel PositionsPB) (gosModel Positions) {
	for _, model := range pbsModel {
		gosModel = append(gosModel, ConvertPositionPB2Object(model))
	}
	return gosModel
}

type (
	RolePermissions   = []*ent.RolePermission
	RolePermissionsPB = []*pb.RolePermission
)

// ConvertRolePermissions2PB RolePermissions holds the value of the role_permissions edge.
func ConvertRolePermissions2PB(gosModel RolePermissions) (pbsModel RolePermissionsPB) {
	for _, model := range gosModel {
		pbsModel = append(pbsModel, ConvertRolePermission2PB(model))
	}
	return pbsModel
}

// ConvertRolePermissionsPB2Object RolePermissions holds the value of the role_permissions edge.
func ConvertRolePermissionsPB2Object(pbsModel RolePermissionsPB) (gosModel RolePermissions) {
	for _, model := range pbsModel {
		gosModel = append(gosModel, ConvertRolePermissionPB2Object(model))
	}
	return gosModel
}

type (
	PermissionResources   = []*ent.PermissionResource
	PermissionResourcesPB = []*pb.PermissionResource
)

// ConvertPermissionResources2PB PermissionResources holds the value of the permission_resources edge.
func ConvertPermissionResources2PB(gosModel PermissionResources) (pbsModel PermissionResourcesPB) {
	for _, model := range gosModel {
		pbsModel = append(pbsModel, ConvertPermissionResource2PB(model))
	}
	return pbsModel
}

// ConvertPermissionResourcesPB2Object PermissionResources holds the value of the permission_resources edge.
func ConvertPermissionResourcesPB2Object(pbsModel PermissionResourcesPB) (gosModel PermissionResources) {
	for _, model := range pbsModel {
		gosModel = append(gosModel, ConvertPermissionResourcePB2Object(model))
	}
	return gosModel
}

type (
	PositionPermissions   = []*ent.PositionPermission
	PositionPermissionsPB = []*pb.PositionPermission
)

// ConvertPositionPermissions2PB PositionPermissions holds the value of the position_permissions edge.
func ConvertPositionPermissions2PB(gosModel PositionPermissions) (pbsModel PositionPermissionsPB) {
	for _, model := range gosModel {
		pbsModel = append(pbsModel, ConvertPositionPermission2PB(model))
	}
	return pbsModel
}

// ConvertPositionPermissionsPB2Object PositionPermissions holds the value of the position_permissions edge.
func ConvertPositionPermissionsPB2Object(pbsModel PositionPermissionsPB) (gosModel PositionPermissions) {
	for _, model := range pbsModel {
		gosModel = append(gosModel, ConvertPositionPermissionPB2Object(model))
	}
	return gosModel
}

type (
	PermissionEdges   = ent.PermissionEdges
	PermissionEdgesPB = pb.PermissionEdges
)

// ConvertPermissionEdges2PB PermissionEdges holds the relations/edges for other nodes in the graph.
func ConvertPermissionEdges2PB(goModel *PermissionEdges) (pbModel *PermissionEdgesPB) {
	pbModel = &PermissionEdgesPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Roles = ConvertRoles2PB(goModel.Roles)
	pbModel.Resources = ConvertResources2PB(goModel.Resources)
	pbModel.Positions = ConvertPositions2PB(goModel.Positions)
	pbModel.RolePermissions = ConvertRolePermissions2PB(goModel.RolePermissions)
	pbModel.PermissionResources = ConvertPermissionResources2PB(goModel.PermissionResources)
	pbModel.PositionPermissions = ConvertPositionPermissions2PB(goModel.PositionPermissions)
	return pbModel
}

// ConvertPermissionEdgesPB2Object PermissionEdges holds the relations/edges for other nodes in the graph.
func ConvertPermissionEdgesPB2Object(pbModel *PermissionEdgesPB) (goModel *PermissionEdges) {
	goModel = &PermissionEdges{}
	if pbModel == nil {
		return goModel
	}

	goModel.Roles = ConvertRolesPB2Object(pbModel.Roles)
	goModel.Resources = ConvertResourcesPB2Object(pbModel.Resources)
	goModel.Positions = ConvertPositionsPB2Object(pbModel.Positions)
	goModel.RolePermissions = ConvertRolePermissionsPB2Object(pbModel.RolePermissions)
	goModel.PermissionResources = ConvertPermissionResourcesPB2Object(pbModel.PermissionResources)
	goModel.PositionPermissions = ConvertPositionPermissionsPB2Object(pbModel.PositionPermissions)
	return goModel
}

type (
	UserPosition   = ent.UserPosition
	UserPositionPB = pb.UserPosition
)

// ConvertUserPosition2PB user_position.table.comment
func ConvertUserPosition2PB(goModel *UserPosition) (pbModel *UserPositionPB) {
	pbModel = &UserPositionPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = int64(goModel.ID)
	pbModel.UserId = int64(goModel.UserID)
	pbModel.PositionId = int64(goModel.PositionID)
	return pbModel
}

// ConvertUserPositionPB2Object user_position.table.comment
func ConvertUserPositionPB2Object(pbModel *UserPositionPB) (goModel *UserPosition) {
	goModel = &UserPosition{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = int64(pbModel.Id)
	goModel.UserID = int64(pbModel.UserId)
	goModel.PositionID = int64(pbModel.PositionId)
	return goModel
}

type (
	UserPositionEdges   = ent.UserPositionEdges
	UserPositionEdgesPB = pb.UserPositionEdges
)

// ConvertUserPositionEdges2PB UserPositionEdges holds the relations/edges for other nodes in the graph.
func ConvertUserPositionEdges2PB(goModel *UserPositionEdges) (pbModel *UserPositionEdgesPB) {
	pbModel = &UserPositionEdgesPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.User = ConvertUser2PB(goModel.User)
	pbModel.Position = ConvertPosition2PB(goModel.Position)
	return pbModel
}

// ConvertUserPositionEdgesPB2Object UserPositionEdges holds the relations/edges for other nodes in the graph.
func ConvertUserPositionEdgesPB2Object(pbModel *UserPositionEdgesPB) (goModel *UserPositionEdges) {
	goModel = &UserPositionEdges{}
	if pbModel == nil {
		return goModel
	}

	goModel.User = ConvertUserPB2Object(pbModel.User)
	goModel.Position = ConvertPositionPB2Object(pbModel.Position)
	return goModel
}

type (
	PositionPermission   = ent.PositionPermission
	PositionPermissionPB = pb.PositionPermission
)

// ConvertPositionPermission2PB position_permission.table.comment
func ConvertPositionPermission2PB(goModel *PositionPermission) (pbModel *PositionPermissionPB) {
	pbModel = &PositionPermissionPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = int64(goModel.ID)
	pbModel.PositionId = int64(goModel.PositionID)
	pbModel.PermissionId = int64(goModel.PermissionID)
	return pbModel
}

// ConvertPositionPermissionPB2Object position_permission.table.comment
func ConvertPositionPermissionPB2Object(pbModel *PositionPermissionPB) (goModel *PositionPermission) {
	goModel = &PositionPermission{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = int64(pbModel.Id)
	goModel.PositionID = int64(pbModel.PositionId)
	goModel.PermissionID = int64(pbModel.PermissionId)
	return goModel
}

type (
	PositionPermissionEdges   = ent.PositionPermissionEdges
	PositionPermissionEdgesPB = pb.PositionPermissionEdges
)

// ConvertPositionPermissionEdges2PB PositionPermissionEdges holds the relations/edges for other nodes in the graph.
func ConvertPositionPermissionEdges2PB(goModel *PositionPermissionEdges) (pbModel *PositionPermissionEdgesPB) {
	pbModel = &PositionPermissionEdgesPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Position = ConvertPosition2PB(goModel.Position)
	pbModel.Permission = ConvertPermission2PB(goModel.Permission)
	return pbModel
}

// ConvertPositionPermissionEdgesPB2Object PositionPermissionEdges holds the relations/edges for other nodes in the graph.
func ConvertPositionPermissionEdgesPB2Object(pbModel *PositionPermissionEdgesPB) (goModel *PositionPermissionEdges) {
	goModel = &PositionPermissionEdges{}
	if pbModel == nil {
		return goModel
	}

	goModel.Position = ConvertPositionPB2Object(pbModel.Position)
	goModel.Permission = ConvertPermissionPB2Object(pbModel.Permission)
	return goModel
}

type (
	RolePermission   = ent.RolePermission
	RolePermissionPB = pb.RolePermission
)

// ConvertRolePermission2PB role_permission.table.comment
func ConvertRolePermission2PB(goModel *RolePermission) (pbModel *RolePermissionPB) {
	pbModel = &RolePermissionPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = int64(goModel.ID)
	pbModel.RoleId = int64(goModel.RoleID)
	pbModel.PermissionId = int64(goModel.PermissionID)
	return pbModel
}

// ConvertRolePermissionPB2Object role_permission.table.comment
func ConvertRolePermissionPB2Object(pbModel *RolePermissionPB) (goModel *RolePermission) {
	goModel = &RolePermission{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = int64(pbModel.Id)
	goModel.RoleID = int64(pbModel.RoleId)
	goModel.PermissionID = int64(pbModel.PermissionId)
	return goModel
}

type (
	RolePermissionEdges   = ent.RolePermissionEdges
	RolePermissionEdgesPB = pb.RolePermissionEdges
)

// ConvertRolePermissionEdges2PB RolePermissionEdges holds the relations/edges for other nodes in the graph.
func ConvertRolePermissionEdges2PB(goModel *RolePermissionEdges) (pbModel *RolePermissionEdgesPB) {
	pbModel = &RolePermissionEdgesPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Role = ConvertRole2PB(goModel.Role)
	pbModel.Permission = ConvertPermission2PB(goModel.Permission)
	return pbModel
}

// ConvertRolePermissionEdgesPB2Object RolePermissionEdges holds the relations/edges for other nodes in the graph.
func ConvertRolePermissionEdgesPB2Object(pbModel *RolePermissionEdgesPB) (goModel *RolePermissionEdges) {
	goModel = &RolePermissionEdges{}
	if pbModel == nil {
		return goModel
	}

	goModel.Role = ConvertRolePB2Object(pbModel.Role)
	goModel.Permission = ConvertPermissionPB2Object(pbModel.Permission)
	return goModel
}

type (
	PermissionResource   = ent.PermissionResource
	PermissionResourcePB = pb.PermissionResource
)

// ConvertPermissionResource2PB permission_resource.table.comment
func ConvertPermissionResource2PB(goModel *PermissionResource) (pbModel *PermissionResourcePB) {
	pbModel = &PermissionResourcePB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = int64(goModel.ID)
	pbModel.PermissionId = int64(goModel.PermissionID)
	pbModel.ResourceId = int64(goModel.ResourceID)
	pbModel.Actions = goModel.Actions
	return pbModel
}

// ConvertPermissionResourcePB2Object permission_resource.table.comment
func ConvertPermissionResourcePB2Object(pbModel *PermissionResourcePB) (goModel *PermissionResource) {
	goModel = &PermissionResource{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = int64(pbModel.Id)
	goModel.PermissionID = int64(pbModel.PermissionId)
	goModel.ResourceID = int64(pbModel.ResourceId)
	goModel.Actions = pbModel.Actions
	return goModel
}

type (
	PermissionResourceEdges   = ent.PermissionResourceEdges
	PermissionResourceEdgesPB = pb.PermissionResourceEdges
)

// ConvertPermissionResourceEdges2PB PermissionResourceEdges holds the relations/edges for other nodes in the graph.
func ConvertPermissionResourceEdges2PB(goModel *PermissionResourceEdges) (pbModel *PermissionResourceEdgesPB) {
	pbModel = &PermissionResourceEdgesPB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Permission = ConvertPermission2PB(goModel.Permission)
	pbModel.Resource = ConvertResource2PB(goModel.Resource)
	return pbModel
}

// ConvertPermissionResourceEdgesPB2Object PermissionResourceEdges holds the relations/edges for other nodes in the graph.
func ConvertPermissionResourceEdgesPB2Object(pbModel *PermissionResourceEdgesPB) (goModel *PermissionResourceEdges) {
	goModel = &PermissionResourceEdges{}
	if pbModel == nil {
		return goModel
	}

	goModel.Permission = ConvertPermissionPB2Object(pbModel.Permission)
	goModel.Resource = ConvertResourcePB2Object(pbModel.Resource)
	return goModel
}
