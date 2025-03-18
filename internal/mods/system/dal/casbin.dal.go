/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dal implements the functions, types, and interfaces for the module.
package dal

import (
	"context"
	"strconv"

	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dto"
)

type casbinSourceRepo struct {
	ctx  context.Context
	data *Data
}

func (c casbinSourceRepo) ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	roles, err := c.data.Role(ctx).Query().WithPermissions(func(query *ent.PermissionQuery) {
		query.WithResources()
	}).All(ctx)
	if err != nil {
		return nil, err
	}
	// 转换策略
	var rules []*pb.PolicyRule
	for _, role := range roles {
		for _, permission := range role.Edges.Permissions {
			for _, resource := range permission.Edges.Resources {
				rules = append(rules, &pb.PolicyRule{
					Ptype: "p",
					Params: []string{
						"role_" + strconv.FormatInt(role.ID, 10),
						resource.Path,
						resource.Method,
					},
				})
			}
		}
	}
	return &pb.ListPoliciesResponse{Rules: rules}, nil
}

func (c casbinSourceRepo) ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	users, err := c.data.User(ctx).Query().WithRoles().All(ctx)
	if err != nil {
		return nil, err
	}
	var rules []*pb.GroupingRule
	for _, user := range users {
		for _, role := range user.Edges.Roles {
			rules = append(rules, &pb.GroupingRule{
				Ptype: "g",
				Params: []string{
					"user_" + strconv.FormatInt(user.ID, 10),
					"role_" + strconv.FormatInt(role.ID, 10),
				},
			})
		}
	}
	return &pb.ListGroupingsResponse{
		Rules: rules,
	}, nil
}

func (c casbinSourceRepo) StreamRules(ctx context.Context, in *pb.StreamRulesRequest) (grpc.ServerStreamingClient[pb.StreamRulesResponse], error) {
	return nil, nil
}

// NewCasbinSourceRepo returns a new CasbinSourceRepo
func NewCasbinSourceRepo(data *Data) (dto.CasbinSourceRepo, error) {
	c := &casbinSourceRepo{
		data: data,
	}
	return c, nil
}

// NewCasbinSourceWithClient create a new CasbinSourceRepo with given client.
// This method does not ensure the existence of database, user should create database manually.
func NewCasbinSourceWithClient(client *ent.Client) (dto.CasbinSourceRepo, error) {
	a := &casbinSourceRepo{
		data: NewDataWithClient(client),
	}
	return a, nil
}
