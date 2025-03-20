/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/goexts/generic/cmp"
	"github.com/goexts/generic/maps"
	"github.com/goexts/generic/settings"
	"github.com/origadmin/runtime/log"

	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/security"
)

// Authorizer is a struct that implements the Authorizer interface.
type Authorizer struct {
	model        model.Model
	adapter      persist.Adapter
	watcher      persist.Watcher
	enforcer     *casbin.SyncedEnforcer
	wildcardItem string
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

	auth.adapter = NewAdapterWithPolicies(p)
	err := auth.watcher.Update()
	if err != nil {
		return errors.Wrap(err, "failed to load adapter")
	}
	return nil
}

func (auth *Authorizer) ApplyDefaults() error {
	if auth.adapter == nil {
		return errors.New("adapter adapter is nil")
	}
	if auth.wildcardItem == "" {
		auth.wildcardItem = "*"
	}
	if auth.model == nil {
		auth.model, _ = model.NewModelFromString(DefaultModel())
	}
	if auth.watcher == nil {
		auth.watcher = NewWatcher()
	}
	if auth.enforcer == nil {
		auth.enforcer, _ = casbin.NewSyncedEnforcer(auth.model, auth.adapter)
	}
	err := auth.enforcer.SetWatcher(auth.watcher)
	if err != nil {
		return err
	}
	if err := auth.enforcer.LoadPolicy(); err != nil {
		return err
	}
	return nil
}

func (auth *Authorizer) WithConfig(config *configv1.AuthZConfig_CasbinConfig) error {
	var err error
	if config.ModelFile != "" {
		auth.model, err = model.NewModelFromFile(config.ModelFile)
	}
	return err
}

func NewAuthorizer(cfg *configv1.Security, ss ...Setting) (security.Authorizer, error) {
	config := cfg.GetAuthz().GetCasbin()
	if config == nil {
		return nil, errors.New("authorizer casbin config is empty")
	}
	var err error
	auth := &Authorizer{}
	err = auth.WithConfig(config)
	if err != nil {
		return nil, err
	}
	return settings.ApplyErrorDefaults(auth, ss)
}
