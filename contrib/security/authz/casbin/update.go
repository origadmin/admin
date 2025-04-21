/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package casbin implements the functions, types, and interfaces for the module.
package casbin

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/goexts/generic/maps"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/security"
	"google.golang.org/grpc/status"

	pb "origadmin/application/admin/api/v1/services/system"
)

type PolicyUpdater struct {
	client       pb.CasbinSourceServiceClient
	adapter      persist.Adapter
	enforcer     *casbin.SyncedEnforcer
	lastModified int64
	interval     time.Duration
}

func (u *PolicyUpdater) Sync(ctx context.Context) (bool, error) {
	start := time.Now()
	defer func() {
		policySyncDuration.Observe(time.Since(start).Seconds())
	}()

	update, err := u.client.WatchUpdate(ctx, &pb.WatchUpdateRequest{
		LastModified: u.lastModified,
	})
	if err != nil {
		return false, err
	}
	if u.lastModified >= update.ModifiedDate {
		return false, nil
	}
	u.lastModified = update.ModifiedDate

	stream, err := u.client.StreamRules(ctx, &pb.StreamRulesRequest{
		WithGroupings: true,
		WithPolicies:  true,
	})
	if err != nil {
		return false, err
	}

	policies := make(map[string][][]string)
	for {
		rule, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, status.Errorf(status.Code(err), "received stream error: %v", err)
		}

		switch v := rule.RuleType.(type) {
		case *pb.StreamRulesResponse_Policy:
			policies[v.Policy.PType] = append(policies[v.Policy.PType], v.Policy.Params)
		case *pb.StreamRulesResponse_Grouping:
			policies[v.Grouping.PType] = append(policies[v.Grouping.PType], v.Grouping.Params)
		}
	}

	if len(policies) > 0 {
		u.lastModified = time.Now().Unix()
		switch setter := u.adapter.(type) {
		case *adapter:
			setter.typedPolicies = policies
		case security.PolicyRegistry:
			pm := maps.Transform(policies, func(k string, v [][]string) (string, any, bool) {
				return k, any(v), true
			})
			if err := setter.SetPolicies(ctx, pm); err != nil {
				return false, err
			}
		default:
			return false, errors.New("unsupported adapter")
		}
		policyCountGauge.Set(float64(len(policies)))
		policySyncCounter.WithLabelValues("success").Inc()
		return true, nil
	}
	return false, nil
}

func (u *PolicyUpdater) Watch(ctx context.Context, notifier persist.Watcher) {
	ticker := time.NewTicker(u.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if update, err := u.Sync(ctx); err != nil || !update {
				log.Errorf("Policy sync failed: %v", err)
				continue
			}
			_ = notifier.Update()
		case <-ctx.Done():
			return
		}
	}
}
