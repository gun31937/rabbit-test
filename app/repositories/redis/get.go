package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func (r *Repository) Get(ctx context.Context, key string) (*string, error) {

	value, err := r.RDB.Get(ctx, key).Result()

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &value, nil
}
