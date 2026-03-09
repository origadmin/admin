/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"context"
	"fmt"

	"github.com/origadmin/runtime/contracts/component"
	storageiface "github.com/origadmin/runtime/contracts/storage"
	"github.com/origadmin/runtime/helpers/comp"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/captcha"
)

// NewCaptcha is the engine provider for *captcha.Captcha.
// It retrieves the designated cache instance with a fallback to the default cache.
func NewCaptcha(ctx context.Context, h component.Handle) (any, error) {
	cfg, err := comp.AsConfig[confpb.Captcha](h)
	if err != nil {
		return nil, err
	}

	// 1. Get Cache instance with fallback logic
	// Strategy: Specified Name -> Default Name
	cacheName := cfg.GetCacheName()
	l := h.Locator().In(component.CategoryCache)

	cache, err := comp.Get[storageiface.Cache](ctx, l, cacheName)
	if err != nil || cache == nil {
		// If named cache not found, fallback to the default one
		cache, err = comp.GetDefault[storageiface.Cache](ctx, l)
	}

	if err != nil || cache == nil {
		return nil, fmt.Errorf("engine: failed to resolve any cache for captcha (requested: %s)", cacheName)
	}

	// 2. Initialize Captcha
	c := &captcha.Config{
		Store:   captcha.NewStore(cache),
		Captcha: cfg,
	}
	return captcha.NewCaptcha(c), nil
}
