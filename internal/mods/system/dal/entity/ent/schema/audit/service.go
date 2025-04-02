/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package audit implements the functions, types, and interfaces for the module.
package audit

import (
	"context"

	"github.com/origadmin/runtime/log"
)

type Service interface {
	Log(ctx context.Context, action string, entity interface{})
}

type service struct {
}

func (s service) Log(ctx context.Context, action string, entity interface{}) {
	log.Info("audit", "action", action, "entity", entity)
}

func NewService() Service {
	return &service{}
}
