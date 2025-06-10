/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"time"

	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"

	"origadmin/application/admin/contrib/security/authz/casbin/internal/model"
)

// AuthorizerOptions contains configuration parameters for Casbin authorizer
// Model:        Required, Casbin model definition
// Adapter:     Required, policy persistence adapter
// Watcher:     Optional, policy change watcher
// Enforcer:    Optional, existing synced enforcer instance
// SyncInterval: Optional, policy sync interval (default 5s)
// Source: gRPC source for policy data service
// WildcardItem: Permission matching wildcard (default "*")
type AuthorizerOptions struct {
	Model            casbinmodel.Model
	Adapter          persist.Adapter
	Watcher          persist.Watcher
	Enforcer         *casbin.SyncedEnforcer
	SyncInterval     time.Duration
	Source           RuleSource
	WildcardItem     string
	EnablePrometheus bool
}

// AuthorizerOption function type for configuring AuthorizerOptions
type AuthorizerOption = func(*AuthorizerOptions)

// DefaultAuthorizerOptions parameters for authorizer
// Model:        Creates new empty model
// Watcher:     Initializes new watcher instance
// SyncInterval: 5s sync interval
// WildcardItem: Wildcard "*"
var (
	DefaultAuthorizerOptions = AuthorizerOptions{
		Watcher:      NewWatcher(),
		SyncInterval: 5 * time.Second,
		WildcardItem: "*",
	}
)

// DefaultModel provides default RESTful role-based model definition
// Returns: Predefined RBAC with RESTful model string
func DefaultModel() string {
	return model.DefaultRestfullWithRoleModel
}

// WithModel sets custom Casbin model configuration
// model: Casbin model instance to use
func WithModel(model casbinmodel.Model) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Model = model
	}
}

// WithStringModel configures model from definition string
// str: Model definition string in Casbin syntax
func WithStringModel(str string) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Model, _ = casbinmodel.NewModelFromString(str)
	}
}

// WithFileModel loads model configuration from file
// path: Path to model configuration file
func WithFileModel(path string) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Model, _ = casbinmodel.NewModelFromFile(path)
	}
}

// WithNameModel sets model using predefined model name
// name: Predefined model name from internal/model package
func WithNameModel(name string) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Model, _ = casbinmodel.NewModelFromString(model.MustModel(name))
	}
}

// WithPolicyAdapter sets policy storage adapter
// adapter: Persistence adapter instance (database/file/etc)
func WithPolicyAdapter(adapter persist.Adapter) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Adapter = adapter
	}
}

// WithWatcher sets policy change watcher
// watcher: Watcher implementation for cluster synchronization
func WithWatcher(watcher persist.Watcher) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Watcher = watcher
	}
}

// WithSyncInterval sets policy synchronization interval
// interval: Duration between policy sync operations
func WithSyncInterval(interval time.Duration) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.SyncInterval = interval
	}
}

// WithEnforcer reuses existing enforcer instance
// enforcer: Preconfigured synced enforcer instance
func WithEnforcer(enforcer *casbin.SyncedEnforcer) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Enforcer = enforcer
	}
}

// WithWildcardItem sets permission matching wildcard
// item: Wildcard symbol for policy matching (default "*")
func WithWildcardItem(item string) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.WildcardItem = item
	}
}

// WithSource sets gRPC policy source service source
// source: gRPC source implementing CasbinSourceService
func WithSource(source RuleSource) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Source = source
	}
}

// WithPrometheusMetrics enables Prometheus metrics collection
// enable: Enable Prometheus metrics collection (default false)
func WithPrometheusMetrics(enable bool) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.EnablePrometheus = enable
	}
}

func (s *AuthorizerOptions) Setup() error {
	if s.Adapter == nil {
		s.Adapter = NewAdapter(nil)
	}

	if s.Model == nil {
		var err error
		s.Model, err = casbinmodel.NewModelFromString(DefaultModel())
		if err != nil {
			return err
		}
	}

	if s.Watcher == nil {
		s.Watcher = NewWatcher()
	}
	return nil
}
