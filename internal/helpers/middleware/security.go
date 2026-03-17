/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package middleware

import (
	"context"
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v3/persist"

	contribsecurity "github.com/origadmin/contrib/security"
	"github.com/origadmin/contrib/security/request"
	"github.com/origadmin/contrib/security/skip"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/contracts/component"
	storageiface "github.com/origadmin/runtime/contracts/storage"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/security"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/helpers/captcha"
)

var (
	gatewaySkipMap = make(map[string]struct{})
	backendSkipMap = make(map[string]struct{})
)

const (
	PolicyNamePublic = "public"
	PolicyNameAuthN  = "authn"

	NameCasbinAdapter = "casbin-adapter"
)

const (
	TypeCasbin = "casbin"
)

const (
	CategoryWatcher = "watcher"
)

const (
	PurposeAdapter = "casbin-adapter"
	PurposeWatcher = "watcher"
)

// RegisterFilterPolicies filters registered policies into skip maps.
func RegisterFilterPolicies() {
	ps := security.RegisteredPolicies()
	for _, p := range ps {
		if p.Name == PolicyNamePublic || p.Name == PolicyNameAuthN {
			gatewaySkipMap[p.ServiceMethod] = struct{}{}
			backendSkipMap[p.ServiceMethod] = struct{}{}
		}
	}
}

// GatewaySkipperProvider creates the gateway skipper.
func GatewaySkipperProvider(ctx context.Context, h component.Handle) (any, error) {
	return contribsecurity.Skipper(func(ctx context.Context, req contribsecurity.Request) bool {
		if req, err := request.NewFromServerContext(ctx); err == nil {
			if _, ok := gatewaySkipMap[req.GetOperation()]; ok {
				return true
			}
		}
		return false
	}), nil
}

// BackendSkipperProvider creates the backend skipper.
func BackendSkipperProvider(ctx context.Context, h component.Handle) (any, error) {
	adminSkipper := skip.Principal(func(principal contribsecurity.Principal) bool {
		id := data.GetSystemUserID()
		if id == 0 {
			return false
		}
		pid := strconv.FormatInt(id, 10)
		return principal.GetID() == pid
	})
	pathSkipper := func(ctx context.Context, req contribsecurity.Request) bool {
		_, ok := backendSkipMap[req.GetOperation()]
		return ok
	}
	return skip.Composite(adminSkipper, pathSkipper), nil
}

func AuthzRequirementResolver(ctx context.Context, h component.Handle, purpose string) (any, error) {
	switch purpose {
	case PurposeAdapter:
		// CategoryStorage and CategoryWatcher are infrastructure components in GlobalScope
		storageH := h.Locator().In(runtime.CategoryStorage)
		adapterInst, err := comp.Get[persist.Adapter](ctx, storageH, NameCasbinAdapter)
		if err != nil {
			return nil, fmt.Errorf("engine: failed to get casbin adapter: %w", err)
		}
		return adapterInst, nil
	case PurposeWatcher:
		watcherH := h.Locator().In(CategoryWatcher)
		watcherInst, err := comp.GetDefault[persist.Watcher](ctx, watcherH)
		if err != nil {
			return nil, fmt.Errorf("engine: failed to get watcher: %w", err)
		}
		return watcherInst, nil
	default:
		return nil, nil
	}
}
