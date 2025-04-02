/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package schema implements the functions, types, and interfaces for the module.
package schema

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent"
	"github.com/origadmin/runtime/log"

	"origadmin/application/admin/internal/mods/system/dal/entity/ent/schema/audit"
)

func UserAuditHook(service audit.Service) ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			value, err := mutator.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			go func() {
				defer func() {
					if r := recover(); r != nil {
						log.Error("audit panic recovered", r)
					}
				}()
				action := strings.ToUpper(strings.TrimPrefix(fmt.Sprintf("%T", m), "*ent."))
				service.Log(context.Background(), action, value)
			}()

			return value, nil
		})
	}
}
