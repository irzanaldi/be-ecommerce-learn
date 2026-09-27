package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	redis *redis.Client
	ctx   context.Context
}

func NewCacheService(redisClient *redis.Client) *CacheService {
	return &CacheService{
		redis: redisClient,
		ctx:   context.Background(),
	}
}

func (c *CacheService) Get(ket string, dest interface{}) (bool, error) {
	val, err := c.redis.Get(c.ctx, key)
}
