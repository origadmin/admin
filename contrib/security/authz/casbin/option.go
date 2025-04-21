/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"time"

	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/contrib/security/authz/casbin/internal/model"
)

type AuthorizerOptions struct {
	Model        casbinmodel.Model            // Need
	Adapter      persist.Adapter              // Need
	Watcher      persist.Watcher              // Optional
	Enforcer     *casbin.SyncedEnforcer       // Optional
	SyncInterval time.Duration                // Optional
	Client       pb.CasbinSourceServiceClient // gRPC client
	WildcardItem string
}

// AuthorizerOption is a function type for setting the Authenticator.
type AuthorizerOption = func(*AuthorizerOptions)

var (
	DefaultAuthorizerOptions = AuthorizerOptions{
		Model:        casbinmodel.NewModel(),
		Watcher:      NewWatcher(),
		SyncInterval: 5 * time.Second,
		WildcardItem: "*",
	}
)

func DefaultModel() string {
	return model.DefaultRestfullWithRoleModel
}

func WithModel(model casbinmodel.Model) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Model = model
	}
}

func WithStringModel(str string) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Model, _ = casbinmodel.NewModelFromString(str)
	}
}

func WithFileModel(path string) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Model, _ = casbinmodel.NewModelFromFile(path)
	}
}

func WithNameModel(name string) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Model, _ = casbinmodel.NewModelFromString(model.MustModel(name))
	}
}

func WithPolicyAdapter(adapter persist.Adapter) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Adapter = adapter
	}
}

func WithWatcher(watcher persist.Watcher) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Watcher = watcher
	}
}

func WithSyncInterval(interval time.Duration) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.SyncInterval = interval
	}
}

func WithEnforcer(enforcer *casbin.SyncedEnforcer) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Enforcer = enforcer
	}
}

func WithWildcardItem(item string) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.WildcardItem = item
	}
}

func WithClient(client pb.CasbinSourceServiceClient) AuthorizerOption {
	return func(s *AuthorizerOptions) {
		s.Client = client
	}
}
