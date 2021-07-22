package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
	redisRepo "rabbit-test/app/repositories/redis"
)

type Redis interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (*string, error)
}

func InitRedis(rdb *redis.Client) Redis {
	return &redisRepo.Repository{
		RDB: rdb,
	}
}
