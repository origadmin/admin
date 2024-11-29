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
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = httperr.New("http.response.status."+pb.SystemErrorReason_SYSTEM_ERROR_REASON_USER_NOT_FOUND.String(), http.StatusNotFound, "user not found")
)

// ConvertMenu2PB Menu is the model entity for the Menu schema.
func ConvertMenu2PB(goModel *Menu) (pbModel *MenuPB) {
	pbModel = &MenuPB{}
	pbModel.Id = goModel.ID
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Code = goModel.Code
	pbModel.Name = goModel.Name
	pbModel.Description = goModel.Description
	pbModel.Sequence = int32(goModel.Sequence)
	pbModel.Type = goModel.Type
	pbModel.Path = goModel.Path
	pbModel.Properties = goModel.Properties
	pbModel.Status = int32(goModel.Status)
	pbModel.ParentId = goModel.ParentID
	pbModel.ParentPath = goModel.ParentPath
	//pbModel.Edges = goModel.Edges
	return pbModel
}

// ConvertMenuEdges2PB MenuEdges holds the relations/edges for other nodes in the graph.
func ConvertMenuEdges2PB(goModel *MenuEdges) (pbModel *MenuEdgesPB) {
	pbModel = &MenuEdgesPB{}
	pbModel.Children = ConvertMenus(goModel.Children)
	pbModel.Parent = ConvertMenu2PB(goModel.Parent)
	pbModel.Resources = ConvertMenuResources(goModel.Resources)
	pbModel.Roles = ConvertRoles(goModel.Roles)
	pbModel.RoleMenus = ConvertRoleMenus(goModel.RoleMenus)
	return pbModel
}

// ConvertRole2PB RolePB is the model entity for the RolePB schema.
func ConvertRole2PB(goModel *Role) (pbModel *RolePB) {
	pbModel = &RolePB{}
	pbModel.Id = goModel.ID
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Code = goModel.Code
	pbModel.Name = goModel.Name
	pbModel.Description = goModel.Description
	pbModel.Sequence = int32(goModel.Sequence)
	pbModel.Status = int32(goModel.Status)
	pbModel.Edges = ConvertRoleEdges2PB(&goModel.Edges)
	return pbModel
}

// ConvertRoleEdges2PB RoleEdges holds the relations/edges for other nodes in the graph.
func ConvertRoleEdges2PB(goModel *RoleEdges) (pbModel *RoleEdgesPB) {
	pbModel = &RoleEdgesPB{}
	pbModel.Menus = ConvertMenus(goModel.Menus)
	pbModel.Users = ConvertUsers(goModel.Users)
	pbModel.RoleMenus = ConvertRoleMenus(goModel.RoleMenus)
	pbModel.UserRoles = ConvertUserRoles(goModel.UserRoles)
	return pbModel
}

// ConvertUser2PB UserPB is the model entity for the UserPB schema.
func ConvertUser2PB(goModel *User) (pbModel *UserPB) {
	pbModel = &UserPB{}
	pbModel.Id = goModel.ID
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.Username = goModel.Username
	pbModel.Name = goModel.Name
	pbModel.Avatar = goModel.Avatar
	pbModel.Password = goModel.Password
	pbModel.Salt = goModel.Salt
	pbModel.Phone = goModel.Phone
	pbModel.Email = goModel.Email
	pbModel.Remark = goModel.Remark
	pbModel.Status = int32(goModel.Status)
	pbModel.Edges = ConvertUserEdges2PB(&goModel.Edges)
	return pbModel
}

// ConvertUserEdges2PB UserEdges holds the relations/edges for other nodes in the graph.
func ConvertUserEdges2PB(goModel *UserEdges) (pbModel *UserEdgesPB) {
	pbModel = &UserEdgesPB{}
	pbModel.Roles = ConvertRoles(goModel.Roles)
	pbModel.UserRoles = ConvertUserRoles(goModel.UserRoles)
	return pbModel
}

// ConvertUserRole2PB UserRole is the model entity for the UserRole schema.
func ConvertUserRole2PB(goModel *UserRole) (pbModel *UserRolePB) {
	pbModel = &UserRolePB{}
	pbModel.Id = goModel.ID
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.UserId = goModel.UserID
	pbModel.RoleId = goModel.RoleID
	pbModel.RoleName = goModel.RoleName
	pbModel.Edges = ConvertUserRoleEdges2PB(&goModel.Edges)
	return pbModel
}

// ConvertUserRoleEdges2PB UserRoleEdges holds the relations/edges for other nodes in the graph.
func ConvertUserRoleEdges2PB(goModel *UserRoleEdges) (pbModel *UserRoleEdgesPB) {
	pbModel = &UserRoleEdgesPB{}
	pbModel.User = ConvertUser2PB(goModel.User)
	pbModel.Role = ConvertRole2PB(goModel.Role)
	return pbModel
}

// ConvertRoleMenu2PB RoleMenu is the model entity for the RoleMenu schema.
func ConvertRoleMenu2PB(goModel *RoleMenu) (pbModel *RoleMenuPB) {
	pbModel = &RoleMenuPB{}
	pbModel.Id = goModel.ID
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.RoleId = goModel.RoleID
	pbModel.MenuId = goModel.MenuID
	pbModel.Edges = ConvertRoleMenuEdges2PB(&goModel.Edges)
	return pbModel
}

// ConvertRoleMenuEdges2PB RoleMenuEdges holds the relations/edges for other nodes in the graph.
func ConvertRoleMenuEdges2PB(goModel *RoleMenuEdges) (pbModel *RoleMenuEdgesPB) {
	pbModel = &RoleMenuEdgesPB{}
	pbModel.Role = ConvertRole2PB(goModel.Role)
	pbModel.Menu = ConvertMenu2PB(goModel.Menu)
	return pbModel
}

// ConvertMenuResource2PB MenuResource is the model entity for the MenuResource schema.
func ConvertMenuResource2PB(goModel *MenuResource) (pbModel *MenuResourcePB) {
	pbModel = &MenuResourcePB{}
	pbModel.Id = goModel.ID
	pbModel.CreateTime = timestamppb.New(goModel.CreateTime)
	pbModel.UpdateTime = timestamppb.New(goModel.UpdateTime)
	pbModel.MenuId = goModel.MenuID
	pbModel.Method = goModel.Method
	pbModel.Path = goModel.Path
	pbModel.Edges = ConvertMenuResourceEdges2PB(&goModel.Edges)
	return pbModel
}

// ConvertMenuResourceEdges2PB MenuResourceEdges holds the relations/edges for other nodes in the graph.
func ConvertMenuResourceEdges2PB(goModel *MenuResourceEdges) (pbModel *MenuResourceEdgesPB) {
	pbModel = &MenuResourceEdgesPB{}
	pbModel.Menu = ConvertMenu2PB(goModel.Menu)
	return pbModel
}
