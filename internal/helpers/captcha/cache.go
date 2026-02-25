// Package captcha implements the functions, types, and contracts for the module.
package captcha

import (
	"context"
	"time"

	"github.com/mojocn/base64Captcha"

	storageiface "github.com/origadmin/runtime/contracts/storage"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/helpers/contextutil"
)

const (
	// captchaPrefix is the prefix for captcha keys in the cache.
	captchaPrefix = "captcha:"
	// defaultExpiration is the default expiration time for captcha keys.
	defaultExpiration = 5 * time.Minute
)

type store struct {
	ctx   context.Context
	cache storageiface.Cache
}

// NewStore creates a new captcha store backed by the provided storage.Cache.
// It holds a background context for cache operations.
func NewStore(cache storageiface.Cache) base64Captcha.Store {
	return &store{
		ctx:   context.Background(),
		cache: cache,
	}
}

// NewStoreWithContext creates a new captcha store backed by the provided storage.Cache.
// It holds a provided context for cache operations.
func NewStoreWithContext(ctx context.Context, cache storageiface.Cache) base64Captcha.Store {
	return &store{
		ctx:   ctx,
		cache: cache,
	}
}

// Set stores the captcha value with a default expiration.
func (s *store) Set(id string, value string) error {
	key := captchaPrefix + id
	// The cache's Set method expects a string value.
	return s.cache.Set(s.ctx, key, value, defaultExpiration)
}

// Get retrieves the captcha value.
func (s *store) Get(id string, clear bool) string {
	key := captchaPrefix + id
	helper := log.NewHelper(contextutil.GetLogger(s.ctx))
	// The cache's Get method returns a string value.
	val, err := s.cache.Get(s.ctx, key)
	if err != nil {
		helper.Errorf("failed to get captcha from cache: %v", err)
		return ""
	}
	if clear {
		if err := s.cache.Delete(s.ctx, key); err != nil {
			helper.Errorf("failed to delete captcha from cache: %v", err)
		}
	}
	return val
}

// Verify verifies the captcha value.
func (s *store) Verify(id, answer string, clear bool) bool {
	val := s.Get(id, clear)
	return val == answer
}
