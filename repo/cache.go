package repo

import (
	"time"

	"github.com/labib0x9/ProjectUnsafe/infra/cache/redis"
)

type CacheRepo interface {
	Set(key string, value string, expire time.Duration) error
	Get(key string) (string, error)
}

type cacheRepo struct {
	cache *redis.Redis
}

func NewCacheRepo(
	cache *redis.Redis,
) CacheRepo {
	return &cacheRepo{
		cache: cache,
	}
}

func (r *cacheRepo) Set(key string, value string, expire time.Duration) error {
	return r.cache.Client.Set(
		ctx,
		key,
		value,
		expire,
	).Err()
}

func (r *cacheRepo) Get(key string) (string, error) {
	return r.cache.Client.Get(ctx, key).Result()
}
