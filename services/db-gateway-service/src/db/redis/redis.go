package rediss

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ConnectRedis() (*redis.Client, error) {
	addr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	_, err := rdb.Ping(ctx).Result()
	return rdb, err
}
