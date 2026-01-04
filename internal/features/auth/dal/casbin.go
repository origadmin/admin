/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/features/auth/dto"
)

// CasbinRepo is a repository for casbin rules that implements
// the application's internal CasbinRepo interface.
type CasbinRepo struct {
	db *ent.Database
}

// NewCasbinRepo creates a new casbin repository.
func NewCasbinRepo(db *ent.Database) (dto.CasbinRepo, error) {
	return &CasbinRepo{db: db}, nil
}

// ListPolicies retrieves policy rules ("p" type) from the storage.
func (r *CasbinRepo) ListPolicies(ctx context.Context, in *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	rules, err := r.db.CasbinRule(ctx).Query().Where(casbinrule.PtypeEQ("p")).All(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.ListPoliciesResponse{
		Rules: make([]*pb.PolicyRule, len(rules)),
	}
	for i, rule := range rules {
		resp.Rules[i] = &pb.PolicyRule{
			PType:  rule.Ptype,
			Params: []string{rule.V0, rule.V1, rule.V2, rule.V3, rule.V4, rule.V5},
		}
	}
	return resp, nil
}

// ListGroupings retrieves grouping rules ("g" type) from the storage.
func (r *CasbinRepo) ListGroupings(ctx context.Context, in *pb.ListGroupingsRequest) (*pb.ListGroupingsResponse, error) {
	rules, err := r.db.CasbinRule(ctx).Query().Where(casbinrule.PtypeEQ("g")).All(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.ListGroupingsResponse{
		Rules: make([]*pb.GroupingRule, len(rules)),
	}
	for i, rule := range rules {
		resp.Rules[i] = &pb.GroupingRule{
			PType:  rule.Ptype,
			Params: []string{rule.V0, rule.V1, rule.V2, rule.V3, rule.V4, rule.V5},
		}
	}
	return resp, nil
}
