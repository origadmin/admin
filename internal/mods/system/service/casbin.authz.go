/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the moduls.enforcer.
package service

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	casbinv2 "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/goexts/generic/cmp"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"
	"github.com/origadmin/toolkits/security"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/status"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/contrib/security/authz/casbin"
)

type CasbinAuthorizerService struct {
	adapter      persist.Adapter
	enforcer     *casbinv2.Enforcer
	client       pb.CasbinSourceServiceClient
	mu           sync.RWMutex
	callback     func(string)
	lastModified int64
	interval     int64
	wildcardItem string
}

const MaxRetryDelay = time.Minute

var (
	policySyncCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "casbin_policy_sync_total",
			Help: "Total number of policy sync operations",
		},
		[]string{"status"},
	)

	policyCountGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "casbin_policy_count",
			Help: "Current number of loaded policies",
		},
	)

	policySyncDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "casbin_sync_duration_seconds",
			Help:    "Histogram of policy sync durations",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	//prometheus.MustRegister(
	//	policySyncCounter,
	//	policyCountGauge,
	//	policySyncDuration,
	//)
}

func (s *CasbinAuthorizerService) Authorized(ctx context.Context, policy security.Policy, object string, action string) (bool, error) {
	log.Debugf("Authorizing user with adapter: %+v", policy)
	var err error
	var allowed bool
	if object == "" {
		object = policy.GetObject()
	}
	if action == "" {
		action = policy.GetAction()
	}
	if allowed, err = s.enforcer.Enforce(policy.GetSubject(), object, action); err != nil {
		log.Errorf("Authorization failed with error: %v", err)
		return false, err
	} else if allowed {
		log.Debugf("Authorization successful for user with adapter: %+v", policy)
		return true, nil
	}
	log.Debugf("Authorization failed for user with adapter: %+v", policy)
	return false, nil
}

func (s *CasbinAuthorizerService) AuthorizedWithDomain(ctx context.Context, policy security.Policy, domain string, object string, action string) (bool, error) {
	log.Debugf("Authorizing user with adapter: %+v", policy)
	domain = cmp.Or(domain, policy.GetDomain(), s.wildcardItem)
	object = cmp.Or(object, policy.GetObject())
	action = cmp.Or(action, policy.GetAction())

	if allowed, err := s.enforcer.Enforce(policy.GetSubject(), object, action, domain); err != nil {
		log.Errorf("Authorization failed with error: %v", err)
		return false, err
	} else if allowed {
		log.Debugf("Authorization successful for user with adapter: %+v", policy)
		return true, nil
	}
	log.Debugf("Authorization failed for user with adapter: %+v", policy)
	return false, nil
}

func (s *CasbinAuthorizerService) AuthorizedWithExtra(ctx context.Context, data security.ExtraData) (bool, error) {
	log.Debugf("Authorizing user with extra data: %+v", data)
	policy, ok := data.GetPolicy()
	if !ok {
		return false, errors.New("adapter is empty")
	}
	if allowed, err := s.enforcer.Enforce(policy.GetSubject(), policy.GetObject(), policy.GetAction(),
		policy.GetDomain()); err != nil {
		log.Errorf("Authorization failed with error: %v", err)
		return false, err
	} else if allowed {
		log.Debugf("Authorization successful for user with adapter: %+v", policy)
		return true, nil
	}
	log.Debugf("Authorization failed for user with adapter: %+v", policy)
	return false, nil
}

func (s *CasbinAuthorizerService) SetUpdateCallback(f func(string)) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.callback = f
	return nil
}

func (s *CasbinAuthorizerService) Update() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.callback != nil {
		s.callback("")
	}
	return nil
}

func (s *CasbinAuthorizerService) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.callback = nil
}

func (s *CasbinAuthorizerService) SyncPolicy(ctx context.Context) error {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		policySyncDuration.Observe(duration)
	}()
	stream, err := s.client.StreamRules(ctx, &pb.StreamRulesRequest{
		WithGroupings: true,
		WithPolicies:  true,
	})
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	var policies = make(map[string][][]string)
	for {
		rule, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			st, _ := status.FromError(err)
			return status.Errorf(st.Code(), "recvied error: %v", st.Message())
		}

		switch v := rule.RuleType.(type) {
		case *pb.StreamRulesResponse_Policy:
			policies[v.Policy.PType] = append(policies[v.Policy.PType], v.Policy.Params)
		case *pb.StreamRulesResponse_Grouping:
			policies[v.Grouping.PType] = append(policies[v.Grouping.PType], v.Grouping.Params)
		}
	}
	pLen := len(policies)
	if pLen > 0 {
		s.adapter = casbin.NewAdapterWithPolicies(policies)
		policyCountGauge.Set(float64(pLen))
		policySyncCounter.WithLabelValues("success").Inc()
	}

	return nil
}

func (s *CasbinAuthorizerService) WatchUpdate() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	retryDelay := time.Duration(s.interval) * time.Second
	timer := time.NewTimer(retryDelay)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			response, err := s.client.WatchUpdate(ctx, &pb.WatchUpdateRequest{
				LastModified: s.lastModified,
			})

			if err != nil {
				newDelay := time.Duration(float64(retryDelay) * 1.5)
				if newDelay > MaxRetryDelay {
					newDelay = MaxRetryDelay
				}
				retryDelay = newDelay
				log.Warnf("WatchUpdate failed, retrying in %v: %v", retryDelay, err)
				timer.Reset(retryDelay)
				continue
			}

			timer.Reset(time.Duration(s.interval) * time.Second)
			lastDate := response.ModifiedDate
			if lastDate > s.lastModified || s.lastModified == 0 {
				s.lastModified = lastDate
				_ = s.Update()
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}

// NewCasbinAuthorizerService new a casbin service.
func NewCasbinAuthorizerService(client pb.CasbinSourceServiceClient) security.Authorizer {
	return &CasbinAuthorizerService{
		client:       client,
		lastModified: 0,
		interval:     int64(5 * time.Minute),
	}
}

func NewCasbinSourceServiceClient(client *service.GRPCClient) pb.CasbinSourceServiceClient {
	return pb.NewCasbinSourceServiceClient(client)
}

//var _ pb.CasbinSourceServiceServer = (*CasbinAuthorizerService)(nil)
