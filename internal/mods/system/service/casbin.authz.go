/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the moduls.enforcer.
package service

import (
	"sync"
	"time"

	"github.com/casbin/casbin/v2/persist"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/contrib/security/authz/casbin"
)

type CasbinAuthorizerService struct {
	client       pb.CasbinSourceServiceClient
	adapter      persist.Adapter
	mu           sync.RWMutex
	callback     func(string)
	interval     time.Duration
	wildcardItem string
	lastModified int64
}

func (s *CasbinAuthorizerService) SetUpdateCallback(callback func(string)) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.callback = callback
	return nil
}

func (s *CasbinAuthorizerService) Update() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.callback != nil {
		s.callback("update")
	}
	return nil
}

func (s *CasbinAuthorizerService) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.callback = nil
	//s.adapter.Close()
}

// NewCasbinAuthorizerService new a casbin service.
func NewCasbinAuthorizerService(client pb.CasbinSourceServiceClient) *CasbinAuthorizerService {
	return &CasbinAuthorizerService{
		client:       client,
		adapter:      casbin.NewAdapter(nil),
		callback:     nil,
		interval:     5 * time.Second,
		wildcardItem: "*",
		lastModified: 0,
	}
}

func NewCasbinSourceServiceClient(client *service.GRPCClient) pb.CasbinSourceServiceClient {
	return pb.NewCasbinSourceServiceClient(client)
}

var _ persist.Watcher = (*CasbinAuthorizerService)(nil)
