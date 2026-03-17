/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package captcha implements the functions, types, and interfaces for the module.
package captcha

import (
	"context"
	"fmt"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/contracts/component"
	storageiface "github.com/origadmin/runtime/contracts/storage"
	"github.com/origadmin/runtime/helpers/comp"
	confpb "origadmin/application/admin/internal/conf/pb"
)

const (
	NameCaptcha = "captcha"
)

// NewCaptchaHandle is the engine provider for *captcha.Captcha.
func NewCaptchaHandle(ctx context.Context, h component.Handle) (any, error) {
	cfg, err := comp.AsConfig[confpb.Captcha](h)
	if err != nil {
		return nil, err
	}

	cacheName := cfg.GetCacheName()
	l := h.Locator().In(runtime.CategoryCache)

	cache, err := comp.Get[storageiface.Cache](ctx, l, cacheName)
	if err != nil || cache == nil {
		cache, err = comp.GetDefault[storageiface.Cache](ctx, l)
	}

	if err != nil || cache == nil {
		return nil, fmt.Errorf("engine: failed to resolve any cache for captcha (requested: %s)", cacheName)
	}

	c := &Config{
		Store:   NewStore(cache),
		Captcha: cfg,
	}
	return NewCaptcha(c), nil
}

func Resolver(ctx context.Context, source any, opts *component.LoadOptions) (*component.ModuleConfig, error) {
	v, ok := source.(*confpb.Bootstrap)
	if !ok {
		return nil, fmt.Errorf("source is not a *confpb.Bootstrap")
	}
	return &component.ModuleConfig{
		Entries: []component.ConfigEntry{
			{
				Name:  NameCaptcha,
				Value: v.GetCaptcha(),
			},
		},
	}, nil
}
