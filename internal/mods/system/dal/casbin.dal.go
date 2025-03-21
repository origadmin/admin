/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dal implements the functions, types, and interfaces for the module.
package dal

import (
	"context"
	"strconv"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dto"
)

type CasbinSourceConfig struct {
	PrefixNumberID func(prefix string, id int64) string
}

type casbinSourceRepo struct {
	ctx    context.Context
	data   *Data
	config *CasbinSourceConfig
}

func (c casbinSourceRepo) mustEmbedUnimplementedCasbinSourceServiceServer() {
	// This method is useless,
	// it is just automatically generated when using the inheritance implementation interface
}

func permissionResourceQuery(query *ent.PermissionQuery) {
	query.WithResources()
}
func (c casbinSourceRepo) ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	rolePermissions, err := c.data.RolePermission(ctx).Query().WithPermission(permissionResourceQuery).WithRole().All(ctx)
	if err != nil {
		return nil, err
	}
	var rules []*pb.PolicyRule
	for _, rolePermission := range rolePermissions {
		permission, err := rolePermission.Edges.PermissionOrErr()
		if err != nil {
			continue
		}
		resources, err := permission.Edges.ResourcesOrErr()
		if err != nil {
			continue
		}
		for _, resource := range resources {
			rules = append(rules, &pb.PolicyRule{
				PType: "p",
				Params: []string{
					c.config.PrefixNumberID("role", rolePermission.RoleID),
					resource.Path,
					resource.Method,
					//todo: add domain support
				},
			})
		}
	}
	return &pb.ListPoliciesResponse{Rules: rules}, nil
}

func (c casbinSourceRepo) ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	userRoles, err := c.data.UserRole(ctx).Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var rules []*pb.GroupingRule
	for _, userRole := range userRoles {
		rules = append(rules, &pb.GroupingRule{
			PType: "g",
			Params: []string{
				c.config.PrefixNumberID("user", userRole.UserID),
				c.config.PrefixNumberID("role", userRole.RoleID),
				//todo: add domain support
			},
		})
	}
	return &pb.ListGroupingsResponse{
		Rules: rules,
	}, nil
}

// NewCasbinSourceRepo returns a new CasbinSourceRepo
func NewCasbinSourceRepo(data *Data) (dto.CasbinSourceRepo, error) {
	c := &casbinSourceRepo{
		data: data,
		config: &CasbinSourceConfig{
			PrefixNumberID: func(prefix string, id int64) string {
				return prefix + "_" + strconv.FormatInt(id, 10)
			},
		},
	}
	return c, nil
}

// NewCasbinSourceWithClient create a new CasbinSourceRepo with given client.
// This method does not ensure the existence of database, user should create database manually.
func NewCasbinSourceWithClient(client *ent.Client) (dto.CasbinSourceRepo, error) {
	c := &casbinSourceRepo{
		data: NewDataWithClient(client),
		config: &CasbinSourceConfig{
			PrefixNumberID: func(prefix string, id int64) string {
				return prefix + "_" + strconv.FormatInt(id, 10)
			},
		},
	}
	return c, nil
}
