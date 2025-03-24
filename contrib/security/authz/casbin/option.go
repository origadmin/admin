/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"errors"

	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/contrib/security/authz/casbin/internal/model"
	"origadmin/application/admin/contrib/security/authz/casbin/internal/policy"
)

type AuthorizerOptions struct {
	Model      casbinmodel.Model
	Adapter    persist.Adapter
	Watcher    persist.Watcher
	ServiceCli pb.CasbinSourceServiceClient
	Interval   int
	RetryDelay int
}

// Setting is a function type for setting the Authenticator.
type Setting = func(*AuthorizerOptions)

func DefaultModel() string {
	return model.DefaultRestfullWithRoleModel
}

func DefaultPolicy() []byte {
	return policy.MustPolicy("keymatch_with_rbac_in_domain.csv")
}

func WithModel(model casbinmodel.Model) Setting {
	return func(s *AuthorizerOptions) {
		s.Model = model
	}
}

func WithStringModel(str string) Setting {
	return func(s *AuthorizerOptions) {
		s.Model, _ = casbinmodel.NewModelFromString(str)
	}
}

func WithFileModel(path string) Setting {
	return func(s *AuthorizerOptions) {
		s.Model, _ = casbinmodel.NewModelFromFile(path)
	}
}

func WithNameModel(name string) Setting {
	return func(s *AuthorizerOptions) {
		s.Model, _ = casbinmodel.NewModelFromString(model.MustModel(name))
	}
}

func WithPolicyAdapter(adapter persist.Adapter) Setting {
	return func(s *AuthorizerOptions) {
		s.Adapter = adapter
	}
}

func WithWatcher(watcher persist.Watcher) Setting {
	return func(s *AuthorizerOptions) {
		s.Watcher = watcher
	}
}

func WithServiceClient(client pb.CasbinSourceServiceClient) Setting {
	return func(s *AuthorizerOptions) {
		s.ServiceCli = client
	}
}

func (auth *AuthorizerOptions) Apply() error {
	if auth.Adapter == nil {
		return errors.New("adapter adapter is nil")
	}
	if auth.Model == nil {
		auth.Model, _ = casbinmodel.NewModelFromString(DefaultModel())
	}
	if auth.Watcher == nil {
		auth.Watcher = NewWatcher()
	}
	if auth.Adapter == nil {
		auth.Adapter = NewAdapter()
	}
	return nil
}
