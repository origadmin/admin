/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"time"

	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/goexts/generic/cmp"
	"github.com/goexts/generic/maps"
	"github.com/goexts/generic/settings"
	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/security"
	"github.com/prometheus/client_golang/prometheus"
)

// Authorizer is a struct that implements the Authorizer interface.
type Authorizer struct {
	options          *AuthorizerOptions
	enforcer         *casbin.SyncedEnforcer
	updater          *PolicyUpdater
	wildcardItem     string
	model            casbinmodel.Model
	adapter          persist.Adapter
	watcher          persist.Watcher
	enablePrometheus bool
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

func (auth *Authorizer) Authorized(ctx context.Context, policy security.Policy, object string, action string) (bool, error) {
	domain := cmp.Or(policy.GetDomain(), "*")
	object = cmp.Or(object, policy.GetObject())
	action = cmp.Or(action, policy.GetAction())
	return auth.enforce(ctx, policy.GetSubject(), object, action, domain)
}

func (auth *Authorizer) AuthorizedWithDomain(ctx context.Context, policy security.Policy, domain string, object string, action string) (bool, error) {
	domain = cmp.Or(domain, policy.GetDomain(), "*")
	object = cmp.Or(object, policy.GetObject())
	action = cmp.Or(action, policy.GetAction())
	return auth.enforce(ctx, policy.GetSubject(), object, action, domain)
}

func (auth *Authorizer) AuthorizedWithExtra(ctx context.Context, data security.ExtraData) (bool, error) {
	policy, ok := data.GetPolicy()
	if !ok {
		return false, errors.New("policy not found in extra data")
	}
	return auth.enforce(ctx, policy.GetSubject(), policy.GetObject(), policy.GetAction(), policy.GetDomain())
}

func (auth *Authorizer) enforce(ctx context.Context, subject, object, action, domain string) (bool, error) {
	allowed, err := auth.enforcer.Enforce(subject, object, action, domain)
	if err != nil {
		log.Errorf("Authorization error: %auth", err)
		return false, err
	}
	log.Debugf("Authorization result: %t for %s %s %s %s", allowed, subject, object, action, domain)
	return allowed, nil
}

func (auth *Authorizer) SetPolicies(ctx context.Context, policies map[string]any, roles map[string]any) error {
	merged := make(map[string][][]string)

	// Merge policy and role data
	maps.Transform(policies, func(k string, v any) (string, [][]string, bool) {
		if vv, ok := v.([][]string); ok {
			merged[k] = append(merged[k], vv...)
			return k, vv, true
		}
		return "", nil, false
	})

	maps.Transform(roles, func(k string, v any) (string, [][]string, bool) {
		if vv, ok := v.([][]string); ok {
			merged[k] = append(merged[k], vv...)
			return k, vv, true
		}
		return "", nil, false
	})

	// Incremental update policy
	if ps, ok := auth.adapter.(security.PolicyRegistry); ok {
		err := ps.SetPolicyRoles(ctx, policies, roles)
		if err != nil {
			return err
		}
	}
	err := auth.watcher.Update()
	if err != nil {
		return err
	}

	return nil
}

func (auth *Authorizer) Apply() error {
	var err error
	auth.adapter = NewAdapter(nil)
	if auth.options.Adapter != nil {
		auth.adapter = auth.options.Adapter
	}
	auth.model, err = casbinmodel.NewModelFromString(DefaultModel())
	if err != nil {
		return err
	}
	if auth.options.Model != nil {
		auth.model = auth.options.Model
	}
	auth.watcher = NewWatcher()
	if auth.options.Watcher != nil {
		auth.watcher = auth.options.Watcher
	}
	if auth.model == nil || auth.adapter == nil {
		return errors.New("model and adapter cannot be nil")
	}
	if auth.options.WildcardItem == "" {
		auth.wildcardItem = "*"
	}
	return nil
}

func NewDefaultAuthorizer() *Authorizer {
	model, _ := casbinmodel.NewModelFromString(DefaultModel())
	return &Authorizer{
		model:            model,
		adapter:          NewAdapter(nil),
		enablePrometheus: false,
	}
}

func NewAuthorizer(cfg *configv1.Security, ss ...AuthorizerOption) (security.Authorizer, error) {
	config := cfg.GetAuthz().GetCasbin()
	if config == nil {
		return nil, errors.New("authorizer casbin config is empty")
	}

	options := settings.ApplyDefault(DefaultAuthorizerOptions, ss)
	if options.ServiceClient == nil {
		return nil, errors.New("authorizer casbin client is empty")
	}

	updater := &PolicyUpdater{
		client:   options.ServiceClient,
		adapter:  options.Adapter,
		interval: options.SyncInterval,
		metric:   options.EnablePrometheus,
	}

	auth := &Authorizer{
		options:      options,
		updater:      updater,
		wildcardItem: options.WildcardItem,
	}

	if err := auth.Apply(); err != nil {
		return nil, err
	}
	_, err := updater.Sync(context.Background())
	if err != nil {
		return nil, err
	}

	enforcer, err := casbin.NewSyncedEnforcer(auth.model, auth.adapter)
	if err != nil {
		return nil, err
	}
	auth.enforcer = enforcer

	if err := auth.enforcer.SetWatcher(auth.watcher); err != nil {
		return nil, err
	}

	go updater.Watch(context.Background(), auth.watcher)

	if options.EnablePrometheus {
		prometheus.MustRegister(
			policySyncCounter,
			policyCountGauge,
			policySyncDuration,
		)
	}

	return auth, nil
}
