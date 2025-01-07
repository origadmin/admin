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

//
//type (
//	Menu   = ent.Menu
//	MenuPB = pb.Menu
//)
//
//// ConvertMenu2PB menu.table.comment
//func ConvertMenu2PB(goModel *Menu) (pbModel *MenuPB) {
//	pbModel = &MenuPB{}
//	if goModel == nil {
//		return pbModel
//	}
//
//	pbModel.Id = int64(goModel.ID)
//	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
//	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
//	pbModel.Keyword = goModel.Keyword
//	pbModel.Name = goModel.Name
//	pbModel.I18NKey = goModel.I18nKey
//	pbModel.Description = goModel.Description
//	pbModel.Type = goModel.Type
//	pbModel.Icon = goModel.Icon
//	pbModel.Path = goModel.Path
//	pbModel.Status = int32(goModel.Status)
//	pbModel.ParentPath = goModel.ParentPath
//	pbModel.Sequence = int32(goModel.Sequence)
//	pbModel.Properties = goModel.Properties
//	pbModel.ParentId = goModel.ParentID
//	return pbModel
//}
//
//// ConvertMenuPB2Object menu.table.comment
//func ConvertMenuPB2Object(pbModel *MenuPB) (goModel *Menu) {
//	goModel = &Menu{}
//	if pbModel == nil {
//		return goModel
//	}
//
//	goModel.ID = int64(pbModel.Id)
//	goModel.CreateTime = pbModel.CreateTime.AsTime()
//	goModel.UpdateTime = pbModel.UpdateTime.AsTime()
//	goModel.Keyword = pbModel.Keyword
//	goModel.Name = pbModel.Name
//	goModel.I18nKey = pbModel.I18NKey
//	goModel.Description = pbModel.Description
//	goModel.Type = pbModel.Type
//	goModel.Icon = pbModel.Icon
//	goModel.Path = pbModel.Path
//	goModel.Status = int8(pbModel.Status)
//	goModel.ParentPath = pbModel.ParentPath
//	goModel.Sequence = int(pbModel.Sequence)
//	goModel.Properties = pbModel.Properties
//	goModel.ParentID = pbModel.ParentId
//	return goModel
//}

type (
	Resource   = ent.Resource
	ResourcePB = pb.Resource
)

// ConvertResource2PB resource.table.comment
func ConvertResource2PB(goModel *Resource) (pbModel *ResourcePB) {
	pbModel = &ResourcePB{}
	if goModel == nil {
		return pbModel
	}

	pbModel.Id = goModel.ID
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Method = goModel.Method
	pbModel.Operation = goModel.Operation
	pbModel.Uri = goModel.URI
	//pbModel.MenuId = goModel.MenuID
	return pbModel
}

// ConvertResourcePB2Object resource.table.comment
func ConvertResourcePB2Object(pbModel *ResourcePB) (goModel *Resource) {
	goModel = &Resource{}
	if pbModel == nil {
		return goModel
	}

	goModel.ID = pbModel.Id
	goModel.CreateTime = pbModel.CreateTime.AsTime()
	goModel.UpdateTime = pbModel.UpdateTime.AsTime()
	goModel.Method = pbModel.Method
	goModel.Operation = pbModel.Operation
	goModel.URI = pbModel.Uri
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
