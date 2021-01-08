package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var REDIS_CLIENT *redis.ClusterClient

func Setup(option *redis.ClusterOptions) *redis.ClusterClient {
	var ctx = context.Background()
	REDIS_CLIENT = redis.NewClusterClient(option)
	err := REDIS_CLIENT.Ping(ctx).Err()
	if err != nil {
		log.Fatal(err)
	}
	return REDIS_CLIENT
}

func Close() {
	REDIS_CLIENT.Close()
}
