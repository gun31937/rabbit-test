package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func (r *Repository) Remove(ctx context.Context, key string) error {

	err, _ := r.RDB.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
