package redis

import (
	"github.com/go-redis/redis/v8"
)

type Repository struct {
	RDB *redis.Client
}
