/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the moduls.enforcer.
package service

import (
	"context"
	"errors"

	"github.com/casbin/casbin/v2"

	pb "origadmin/application/admin/api/v1/services/casbin"
)

type CasbinServiceServer struct {
	pb.UnimplementedCasbinServiceServer
	enforcer *casbin.Enforcer
}

// GetDomains gets the domains that a user has.
func (s *CasbinServiceServer) GetDomains(ctx context.Context, in *pb.UserRoleRequest) (*pb.ArrayReply, error) {
	rm := s.enforcer.GetModel()["g"]["g"].RM
	if rm == nil {
		return nil, errors.New("RoleManager is nil")
	}

	res, _ := rm.GetDomains(in.User)

	return &pb.ArrayReply{Array: res}, nil
}

// GetRolesForUser gets the roles that a user has.
func (s *CasbinServiceServer) GetRolesForUser(ctx context.Context, in *pb.UserRoleRequest) (*pb.ArrayReply, error) {
	rm := s.enforcer.GetModel()["g"]["g"].RM
	if rm == nil {
		return nil, errors.New("RoleManager is nil")
	}

	res, _ := rm.GetRoles(in.User, in.Domain...)

	return &pb.ArrayReply{Array: res}, nil
}

// GetImplicitRolesForUser gets implicit roles that a user has.
func (s *CasbinServiceServer) GetImplicitRolesForUser(ctx context.Context, in *pb.UserRoleRequest) (*pb.ArrayReply, error) {
	res, err := s.enforcer.GetImplicitRolesForUser(in.User)
	return &pb.ArrayReply{Array: res}, err
}

// GetUsersForRole gets the users that have a rols.enforcer.
func (s *CasbinServiceServer) GetUsersForRole(ctx context.Context, in *pb.UserRoleRequest) (*pb.ArrayReply, error) {
	rm := s.enforcer.GetModel()["g"]["g"].RM
	if rm == nil {
		return nil, errors.New("RoleManager is nil")
	}

	res, _ := rm.GetUsers(in.Role)

	return &pb.ArrayReply{Array: res}, nil
}

// HasRoleForUser determines whether a user has a rols.enforcer.
func (s *CasbinServiceServer) HasRoleForUser(ctx context.Context, in *pb.UserRoleRequest) (*pb.BoolReply, error) {
	roles, err := s.enforcer.GetRolesForUser(in.User)
	if err != nil {
		return &pb.BoolReply{}, err
	}

	for _, r := range roles {
		if r == in.Role {
			return &pb.BoolReply{Res: true}, nil
		}
	}

	return &pb.BoolReply{}, nil
}

// AddRoleForUser adds a role for a user.
// Returns false if the user already has the role (aka not affected).
func (s *CasbinServiceServer) AddRoleForUser(ctx context.Context, in *pb.UserRoleRequest) (*pb.BoolReply, error) {
	ruleAdded, err := s.enforcer.AddGroupingPolicy(in.User, in.Role)
	return &pb.BoolReply{Res: ruleAdded}, err
}

// DeleteRoleForUser deletes a role for a user.
// Returns false if the user does not have the role (aka not affected).
func (s *CasbinServiceServer) DeleteRoleForUser(ctx context.Context, in *pb.UserRoleRequest) (*pb.BoolReply, error) {
	ruleRemoved, err := s.enforcer.RemoveGroupingPolicy(in.User, in.Role)
	return &pb.BoolReply{Res: ruleRemoved}, err
}

// DeleteRolesForUser deletes all roles for a user.
// Returns false if the user does not have any roles (aka not affected).
func (s *CasbinServiceServer) DeleteRolesForUser(ctx context.Context, in *pb.UserRoleRequest) (*pb.BoolReply, error) {
	ruleRemoved, err := s.enforcer.RemoveFilteredGroupingPolicy(0, in.User)
	return &pb.BoolReply{Res: ruleRemoved}, err
}

// DeleteUser deletes a user.
// Returns false if the user does not exist (aka not affected).
func (s *CasbinServiceServer) DeleteUser(ctx context.Context, in *pb.UserRoleRequest) (*pb.BoolReply, error) {
	ruleRemoved, err := s.enforcer.RemoveFilteredGroupingPolicy(0, in.User)
	return &pb.BoolReply{Res: ruleRemoved}, err
}

// DeleteRole deletes a rols.enforcer.
func (s *CasbinServiceServer) DeleteRole(ctx context.Context, in *pb.UserRoleRequest) (*pb.EmptyReply, error) {
	_, err := s.enforcer.DeleteRole(in.Role)
	return &pb.EmptyReply{}, err
}

// DeletePermission deletes a permission.
// Returns false if the permission does not exist (aka not affected).
func (s *CasbinServiceServer) DeletePermission(ctx context.Context, in *pb.PermissionRequest) (*pb.BoolReply, error) {
	ruleRemoved, err := s.enforcer.RemoveFilteredPolicy(1, in.Permissions...)
	return &pb.BoolReply{Res: ruleRemoved}, err
}

// AddPermissionForUser adds a permission for a user or rols.enforcer.
// Returns false if the user or role already has the permission (aka not affected).
func (s *CasbinServiceServer) AddPermissionForUser(ctx context.Context, in *pb.PermissionRequest) (*pb.BoolReply, error) {
	ruleAdded, err := s.enforcer.AddPolicy(s.convertPermissions(in.User, in.Permissions...)...)
	return &pb.BoolReply{Res: ruleAdded}, err
}

// DeletePermissionForUser deletes a permission for a user or rols.enforcer.
// Returns false if the user or role does not have the permission (aka not affected).
func (s *CasbinServiceServer) DeletePermissionForUser(ctx context.Context, in *pb.PermissionRequest) (*pb.BoolReply, error) {
	ruleRemoved, err := s.enforcer.RemovePolicy(s.convertPermissions(in.User, in.Permissions...)...)
	return &pb.BoolReply{Res: ruleRemoved}, err
}

// DeletePermissionsForUser deletes permissions for a user or rols.enforcer.
// Returns false if the user or role does not have any permissions (aka not affected).
func (s *CasbinServiceServer) DeletePermissionsForUser(ctx context.Context, in *pb.PermissionRequest) (*pb.BoolReply, error) {
	ruleRemoved, err := s.enforcer.RemoveFilteredPolicy(0, in.User)
	return &pb.BoolReply{Res: ruleRemoved}, err
}

// GetPermissionsForUser gets permissions for a user or rols.enforcer.
func (s *CasbinServiceServer) GetPermissionsForUser(ctx context.Context, in *pb.PermissionRequest) (*pb.Array2DReply, error) {
	filteredPolicy, err := s.enforcer.GetFilteredPolicy(0, in.User)
	if err != nil {
		return &pb.Array2DReply{}, err
	}

	return s.wrapPlainPolicy(filteredPolicy), nil
}

// GetImplicitPermissionsForUser gets all permissions(including children) for a user or rols.enforcer.
func (s *CasbinServiceServer) GetImplicitPermissionsForUser(ctx context.Context, in *pb.PermissionRequest) (*pb.Array2DReply, error) {
	resp, err := s.enforcer.GetImplicitPermissionsForUser(in.User, in.Domain...)
	return s.wrapPlainPolicy(resp), err
}

// HasPermissionForUser determines whether a user has a permission.
func (s *CasbinServiceServer) HasPermissionForUser(ctx context.Context, in *pb.PermissionRequest) (*pb.BoolReply, error) {
	hasPolicy, err := s.enforcer.HasPolicy(s.convertPermissions(in.User, in.Permissions...)...)
	if err != nil {
		return &pb.BoolReply{}, err
	}

	return &pb.BoolReply{Res: hasPolicy}, nil
}

func (s *CasbinServiceServer) wrapPlainPolicy(policy [][]string) *pb.Array2DReply {
	if len(policy) == 0 {
		return &pb.Array2DReply{}
	}

	policyReply := &pb.Array2DReply{}
	policyReply.D2 = make([]*pb.Array2DReply_D, len(policy))
	for e := range policy {
		policyReply.D2[e] = &pb.Array2DReply_D{D1: policy[e]}
	}

	return policyReply
}

func (s *CasbinServiceServer) convertPermissions(user string, permissions ...string) []interface{} {
	params := make([]interface{}, 0, len(permissions)+1)
	params = append(params, user)
	for _, perm := range permissions {
		params = append(params, perm)
	}

	return params
}
