/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/casbin/casbin/v3/model"
	"github.com/origadmin/runtime/log"
	v1 "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/internal/data"
)

// CasbinService is a service for Casbin.
type CasbinService struct {
	v1.UnimplementedCasbinServiceServer
	adapter *data.CasbinAdapter
	log     *log.Helper
}

// NewCasbinService creates a new Casbin source service.
func NewCasbinService(adapter *data.CasbinAdapter, logger log.Logger) *CasbinService {
	return &CasbinService{
		adapter: adapter,
		log:     log.NewHelper(log.With(logger, "module", "auth.service.casbin")),
	}
}

// ListPolicies returns a list of 'p' policies from the adapter.
func (s *CasbinService) ListPolicies(ctx context.Context, req *v1.ListPoliciesRequest) (*v1.ListPoliciesResponse, error) {
	m := model.NewModel()
	if err := s.adapter.LoadPolicy(m); err != nil {
		s.log.WithContext(ctx).Errorf("Failed to load policies: %v", err)
		return nil, err
	}

	var rules []*v1.PolicyRule
	for ptype, ast := range m["p"] {
		for _, policy := range ast.Policy {
			rules = append(rules, &v1.PolicyRule{
				Ptype:  ptype,
				Params: policy,
			})
		}
	}

	s.log.WithContext(ctx).Debugf("ListPolicies returned %d rules", len(rules))
	return &v1.ListPoliciesResponse{Rules: rules}, nil
}

// ListGroupings returns a list of 'g' grouping policies from the adapter.
func (s *CasbinService) ListGroupings(ctx context.Context, req *v1.ListGroupingsRequest) (*v1.ListGroupingsResponse, error) {
	m := model.NewModel()
	if err := s.adapter.LoadPolicy(m); err != nil {
		s.log.WithContext(ctx).Errorf("Failed to load groupings: %v", err)
		return nil, err
	}

	var rules []*v1.GroupingRule
	for ptype, ast := range m["g"] {
		for _, policy := range ast.Policy {
			rules = append(rules, &v1.GroupingRule{
				Ptype:  ptype,
				Params: policy,
			})
		}
	}

	s.log.WithContext(ctx).Debugf("ListGroupings returned %d rules", len(rules))
	return &v1.ListGroupingsResponse{Rules: rules}, nil
}

// WatchUpdate monitors for policy updates via the watcher.
// This is a placeholder - actual implementation should use the watcher to detect changes.
func (s *CasbinService) WatchUpdate(ctx context.Context, req *v1.WatchUpdateRequest) (*v1.WatchUpdateResponse, error) {
	s.log.WithContext(ctx).Debugf("WatchUpdate called with last_modified=%d", req.GetLastModified())
	// TODO: Implement actual watch mechanism using the watcher
	return &v1.WatchUpdateResponse{
		ModifiedDate: req.GetLastModified(), // Placeholder - return current timestamp
	}, nil
}

// StreamRules streams rules to the client.
// This is a placeholder - actual implementation should support streaming.
func (s *CasbinService) StreamRules(req *v1.StreamRulesRequest, stream v1.CasbinService_StreamRulesServer) error {
	ctx := stream.Context()

	m := model.NewModel()
	if err := s.adapter.LoadPolicy(m); err != nil {
		s.log.WithContext(ctx).Errorf("Failed to load policies for streaming: %v", err)
		return err
	}

	// Stream 'p' policies if requested
	if req.GetWithPolicies() {
		for ptype, ast := range m["p"] {
			for _, policy := range ast.Policy {
				if err := stream.Send(&v1.StreamRulesResponse{
					Rule: &v1.StreamRulesResponse_Policy{
						Policy: &v1.PolicyRule{
							Ptype:  ptype,
							Params: policy,
						},
					},
				}); err != nil {
					return err
				}
			}
		}
	}

	// Stream 'g' groupings if requested
	if req.GetWithGroupings() {
		for ptype, ast := range m["g"] {
			for _, policy := range ast.Policy {
				if err := stream.Send(&v1.StreamRulesResponse{
					Rule: &v1.StreamRulesResponse_Grouping{
						Grouping: &v1.GroupingRule{
							Ptype:  ptype,
							Params: policy,
						},
					},
				}); err != nil {
					return err
				}
			}
		}
	}

	s.log.WithContext(ctx).Debug("StreamRules completed")
	return nil
}
