/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"io"
	"time"

	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/goexts/generic/cmp"
	"github.com/goexts/generic/maps"
	"github.com/goexts/generic/settings"
	"github.com/origadmin/runtime/log"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/status"

	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/security"

	pb "origadmin/application/admin/api/v1/services/system"
)

// Authorizer is a struct that implements the Authorizer interface.
type Authorizer struct {
	client       pb.CasbinSourceServiceClient
	options      *AuthorizerOptions
	enforcer     *casbin.SyncedEnforcer
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
	prometheus.MustRegister(
		policySyncCounter,
		policyCountGauge,
		policySyncDuration,
	)
}

func (auth *Authorizer) Authorized(ctx context.Context, policy security.Policy, object string, action string) (bool, error) {
	log.Debugf("Authorizing user with adapter: %+v", policy)
	var err error
	var allowed bool
	if object == "" {
		object = policy.GetObject()
	}
	if action == "" {
		action = policy.GetAction()
	}
	if allowed, err = auth.enforcer.Enforce(policy.GetSubject(), object, action); err != nil {
		log.Errorf("Authorization failed with error: %v", err)
		return false, err
	} else if allowed {
		log.Debugf("Authorization successful for user with adapter: %+v", policy)
		return true, nil
	}
	log.Debugf("Authorization failed for user with adapter: %+v", policy)
	return false, nil
}

func (auth *Authorizer) AuthorizedWithDomain(ctx context.Context, policy security.Policy, domain string, object string, action string) (bool, error) {
	log.Debugf("Authorizing user with adapter: %+v", policy)
	domain = cmp.Or(domain, policy.GetDomain(), auth.wildcardItem)
	object = cmp.Or(object, policy.GetObject())
	action = cmp.Or(action, policy.GetAction())

	if allowed, err := auth.enforcer.Enforce(policy.GetSubject(), object, action, domain); err != nil {
		log.Errorf("Authorization failed with error: %v", err)
		return false, err
	} else if allowed {
		log.Debugf("Authorization successful for user with adapter: %+v", policy)
		return true, nil
	}
	log.Debugf("Authorization failed for user with adapter: %+v", policy)
	return false, nil
}

func (auth *Authorizer) AuthorizedWithExtra(ctx context.Context, data security.ExtraData) (bool, error) {
	log.Debugf("Authorizing user with extra data: %+v", data)
	policy, ok := data.GetPolicy()
	if !ok {
		return false, errors.New("adapter is empty")
	}
	if allowed, err := auth.enforcer.Enforce(policy.GetSubject(), policy.GetObject(), policy.GetAction(), policy.GetDomain()); err != nil {
		log.Errorf("Authorization failed with error: %v", err)
		return false, err
	} else if allowed {
		log.Debugf("Authorization successful for user with adapter: %+v", policy)
		return true, nil
	}
	log.Debugf("Authorization failed for user with adapter: %+v", policy)
	return false, nil
}

func (auth *Authorizer) SetPolicies(ctx context.Context, policies map[string]any, roles map[string]any) error {
	p := maps.Transform(policies, func(k string, v any) (string, [][]string, bool) {
		if vv, ok := v.([][]string); ok {
			return k, vv, ok
		}
		return k, [][]string{}, false
	})

	adapter := NewAdapterWithPolicies(p)
	auth.enforcer.SetAdapter(adapter)
	//err := auth.enforcer.LoadPolicy()
	err := auth.options.Watcher.Update()
	if err != nil {
		return errors.Wrap(err, "failed to load adapter")
	}
	return nil
}

func (auth *Authorizer) SyncPolicy(ctx context.Context) error {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		policySyncDuration.Observe(duration)
	}()
	stream, err := auth.client.StreamRules(ctx, &pb.StreamRulesRequest{
		WithGroupings: true,
		WithPolicies:  true,
	})
	if err != nil {
		return err
	}
	log.Infof("Syncing policies...")
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
	log.Infof("Updated %d policies", pLen)
	for s, p := range policies {
		log.Debugf("Updated %d %s policies", len(p), s)
	}
	if pLen > 0 {
		auth.lastModified = time.Now().Unix()
		if setter, ok := auth.options.Adapter.(interface {
			SetPolicies(policies map[string][][]string)
		}); ok {
			log.Infof("Setting policies...")
			setter.SetPolicies(policies)
		}
		//err := auth.options.Watcher.Update()
		//if err != nil {
		//	policySyncCounter.WithLabelValues("failed").Inc()
		//	return err
		//}
		//for ptype, policy := range policies {
		//	switch ptype {
		//	case "p":
		//		log.Debugf("Updating %d policies", len(policy))
		//		added, err := auth.enforcer.AddPolicies(policy)
		//		if err != nil {
		//			log.Warnf("Failed to add policies: %v", err)
		//			return err
		//		}
		//		log.Debugf("Added %b policies", added)
		//	case "g":
		//		log.Debugf("Updating %d groupings", len(policy))
		//		added, err := auth.enforcer.AddGroupingPolicies(policy)
		//		if err != nil {
		//			log.Warnf("Failed to add groupings: %v", err)
		//			return err
		//		}
		//		log.Debugf("Added %b groupings", added)
		//	}
		//
		//}
		err := auth.enforcer.LoadPolicy()
		if err != nil {
			log.Warnf("Failed to load policy: %v", err)
			return err
		}
		policies, err := auth.enforcer.GetPolicy()
		if err != nil {
			log.Warnf("Failed to get policy: %v", err)
			return err
		}
		//for _, policy := range policies {
		//	log.Debugf("Updated policy: %+v", policy)
		//}
		log.Infof("Updated %d policies", len(policies))
		policyCountGauge.Set(float64(pLen))
		policySyncCounter.WithLabelValues("success").Inc()
		//auth.enforcer.update
	}

	return nil
}

func (auth *Authorizer) WatchUpdate() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	interval := time.Duration(auth.interval) * time.Second
	retryDelay := 3 * time.Second
	timer := time.NewTimer(interval)
	defer timer.Stop()
	log.Infof("Watching for updates every %v", interval)
	for {
		select {
		case <-timer.C:
			response, err := auth.client.WatchUpdate(ctx, &pb.WatchUpdateRequest{
				LastModified: auth.lastModified,
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

			retryDelay = 3 * time.Second
			timer.Reset(interval)
			lastDate := response.ModifiedDate
			if lastDate > auth.lastModified || auth.lastModified == 0 {
				log.Infof("Update detected, last modified: %v", lastDate)
				//auth.lastModified = lastDate
				//_ = auth.watcher.Update()
				//auth.enforcer
				go auth.SyncPolicy(ctx)
				continue
			}
			log.Debugf("No update detected, last modified: %v", lastDate)
		case <-ctx.Done():
			return
		}
	}
}

func (auth *Authorizer) Apply() error {
	if auth.options.Adapter == nil {
		return errors.New("adapter adapter is nil")
	}
	if auth.options.Model == nil {
		auth.options.Model, _ = casbinmodel.NewModelFromString(DefaultModel())
	}
	if auth.options.Watcher == nil {
		auth.options.Watcher = NewWatcher()
	}
	if auth.options.Adapter == nil {
		auth.options.Adapter = NewAdapter()
	}
	return nil
}

func NewAuthorizer(cfg *configv1.Security, ss ...Setting) (security.Authorizer, error) {
	config := cfg.GetAuthz().GetCasbin()
	if config == nil {
		return nil, errors.New("authorizer casbin config is empty")
	}
	var err error
	if config.ModelFile != "" {
		ss = append(ss, WithFileModel(config.ModelFile))
	}
	options := settings.Apply(&AuthorizerOptions{
		Interval:   5,
		RetryDelay: 3,
	}, ss)
	auth := &Authorizer{
		interval: 5,
		options:  options,
		client:   options.ServiceCli,
	}
	if err := auth.Apply(); err != nil {
		return nil, err
	}
	auth.enforcer, err = casbin.NewSyncedEnforcer(auth.options.Model, auth.options.Adapter)
	if err != nil {
		return nil, err
	}
	err = auth.enforcer.SetWatcher(auth.options.Watcher)
	if err != nil {
		return nil, err
	}
	go auth.WatchUpdate()
	return auth, nil
}
