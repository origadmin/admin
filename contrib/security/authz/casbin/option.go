/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"

	"github.com/origadmin/contrib/security/authz/casbin/internal/model"
	"github.com/origadmin/contrib/security/authz/casbin/internal/policy"
)

// Setting is a function type for setting the Authenticator.
type Setting = func(*Authorizer)

func DefaultModel() string {
	return model.DefaultRestfullWithRoleModel
}

func DefaultPolicy() []byte {
	return policy.MustPolicy("keymatch_with_rbac_in_domain.csv")
}

func WithModel(model casbinmodel.Model) Setting {
	return func(s *Authorizer) {
		s.model = model
	}
}

func WithStringModel(str string) Setting {
	return func(s *Authorizer) {
		s.model, _ = casbinmodel.NewModelFromString(str)
	}
}

func WithFileModel(path string) Setting {
	return func(s *Authorizer) {
		s.model, _ = casbinmodel.NewModelFromFile(path)
	}
}

func WithNameModel(name string) Setting {
	return func(s *Authorizer) {
		s.model, _ = casbinmodel.NewModelFromString(model.MustModel(name))
	}
}

func WithPolicyAdapter(policy persist.Adapter) Setting {
	return func(s *Authorizer) {
		s.policy = policy
	}
}
