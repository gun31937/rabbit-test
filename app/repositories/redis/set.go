package redis

import (
	"context"
	"rabbit-test/app/env"

	"time"
)

func (r *Repository) Set(ctx context.Context, key string, value string) error {

	itemTTL := time.Duration(env.RedisItemTTL) * time.Minute

	err := r.RDB.Set(ctx, key, value, itemTTL).Err()

	if err != nil {
		return err
	}

	return nil
}
