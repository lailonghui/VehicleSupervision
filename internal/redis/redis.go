package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var REDIS_CLIENT *redis.ClusterClient

func Setup(option *redis.ClusterOptions) *redis.ClusterClient {
	var ctx = context.Background()
	REDIS_CLIENT = redis.NewClusterClient(option)
	REDIS_CLIENT.Ping(ctx)
	return REDIS_CLIENT
}

func Close() {
	REDIS_CLIENT.Close()
}
