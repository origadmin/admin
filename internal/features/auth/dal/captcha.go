package dal

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"github.com/origadmin/runtime/data/storage"

	confpb "origadmin/application/admin/internal/conf/pb"
)

const (
	captchaPrefix = "captcha:"
)

type captchaRepo struct {
	rdb *redis.Client
	log *log.Helper
}

// NewCaptchaRepo creates a new captcha repository that implements the base64Captcha.Store interface.
func NewCaptchaRepo(provider storage.Provider, cfg *confpb.Captcha, logger log.Logger) (base64Captcha.Store, error) {
	cacheName := cfg.GetCacheName()
	if cacheName == "" {
		return nil, fmt.Errorf("captcha cache_name is not configured")
	}

	cache, err := provider.Cache(cacheName)
	if err != nil {
		return nil, fmt.Errorf("failed to get cache '%s': %w", cacheName, err)
	}

	rdb := cache.Redis()
	if rdb == nil {
		return nil, fmt.Errorf("the cache '%s' is not a Redis client", cacheName)
	}

	return &captchaRepo{
		rdb: rdb,
		log: log.NewHelper(logger),
	}, nil
}

// Set stores the captcha value.
func (r *captchaRepo) Set(id string, value string) error {
	return r.rdb.Set(context.Background(), captchaPrefix+id, value, time.Minute*5).Err()
}

// Get retrieves the captcha value.
func (r *captchaRepo) Get(id string, clear bool) string {
	ctx := context.Background()
	key := captchaPrefix + id
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		r.log.Errorf("failed to get captcha from redis: %v", err)
		return ""
	}
	if clear {
		if err := r.rdb.Del(ctx, key).Err(); err != nil {
			r.log.Errorf("failed to delete captcha from redis: %v", err)
		}
	}
	return val
}

// Verify verifies the captcha value.
func (r *captchaRepo) Verify(id, answer string, clear bool) bool {
	val := r.Get(id, clear)
	return val == answer
}
