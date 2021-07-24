package redis

import (
	"context"
)

func (r *Repository) Delete(ctx context.Context, key string) error {

	err := r.RDB.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
