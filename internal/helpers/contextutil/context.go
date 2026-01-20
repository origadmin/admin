// Package contextutil provides utility functions for working with context.
package contextutil

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/origadmin/contrib/security"
	"github.com/origadmin/runtime/log"
)

// ErrNoPrincipalInContext is returned when no principal is found in the context.
var ErrNoPrincipalInContext = errors.New("contextutil: no principal found in context")

// GetUserID extracts the user ID from the context.
// It encapsulates the logic of retrieving user information, which is expected
// to be stored as a principal.Principal. This decouples consumers from the
// specific implementation of the principal package.
func GetUserID(ctx context.Context) (int64, error) {
	p, ok := security.FromContext(ctx)
	if !ok {
		return 0, ErrNoPrincipalInContext
	}

	userID, err := strconv.ParseInt(p.GetID(), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("contextutil: failed to parse principal ID '%s': %w", p.GetID(), err)
	}
	return userID, nil
}

type adminCtx struct{}

func IsAdmin(ctx context.Context, id string) bool {
	if value, ok := ctx.Value(adminCtx{}).(string); ok {
		return value == id
	}
	return false
}

func NewAdmin(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, adminCtx{}, id)
}

type logKey struct{}

func GetLogger(ctx context.Context) log.Logger {
	logger, ok := ctx.Value(logKey{}).(log.Logger)
	if !ok {
		return log.DefaultLogger
	}
	return logger
}

func SetLogger(ctx context.Context, logger log.Logger) context.Context {
	return context.WithValue(ctx, logKey{}, logger)
}

type systemCtx struct{}

func NewSystemUser(ctx context.Context) context.Context {
	return context.WithValue(ctx, systemCtx{}, true)
}

func IsSystemUser(ctx context.Context) bool {
	ok, _ := ctx.Value(systemCtx{}).(bool)
	return ok
}
