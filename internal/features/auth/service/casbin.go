package service

import (
	"context"

	v1 "origadmin/application/admin/api/v1/services/auth"
)

// CasbinService is a service for Casbin.
type CasbinService struct {
	v1.UnimplementedCasbinServiceServer
}

func (s *CasbinService) mustEmbedUnimplementedCasbinServiceServer() {
	//TODO implement me
	panic("implement me")
}

// NewCasbinService creates a new Casbin source service.
func NewCasbinService() *CasbinService {
	return &CasbinService{}
}

// ListPolicies returns a list of policies.
func (s *CasbinService) ListPolicies(ctx context.Context, req *v1.ListPoliciesRequest) (*v1.ListPoliciesResponse, error) {
	return &v1.ListPoliciesResponse{}, nil
}

// ListGroupings returns a list of groupings.
func (s *CasbinService) ListGroupings(ctx context.Context, req *v1.ListGroupingsRequest) (*v1.ListGroupingsResponse, error) {
	return &v1.ListGroupingsResponse{}, nil
}

// WatchUpdate returns a watch update.
func (s *CasbinService) WatchUpdate(ctx context.Context, req *v1.WatchUpdateRequest) (*v1.WatchUpdateResponse, error) {
	return &v1.WatchUpdateResponse{}, nil
}

// StreamRules returns a stream of rules.
func (s *CasbinService) StreamRules(req *v1.StreamRulesRequest, stream v1.CasbinService_StreamRulesServer) error {
	return nil
}
