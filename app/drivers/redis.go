package drivers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"rabbit-test/app/env"
)

// ConnectRedis connect to Redis
var ctx = context.Background()

func ConnectRedis() *redis.Client {

	addr := fmt.Sprintf(
		"%s:%s",
		env.RedisHost,
		env.RedisPort,
	)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: env.RedisPassword,
		DB:       env.RedisDB,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redisEngine error : %+v", err.Error()))
	}

	return rdb
}
