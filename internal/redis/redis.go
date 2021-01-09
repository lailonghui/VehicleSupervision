package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var REDIS_CLIENT redis.UniversalClient

func Setup(option *redis.UniversalOptions) redis.UniversalClient {
	var ctx = context.Background()
	client := redis.NewUniversalClient(option)
	REDIS_CLIENT = client
	err := REDIS_CLIENT.Ping(ctx).Err()
	if err != nil {
		log.Fatal(err)
	}
	return REDIS_CLIENT
}

func Close() {
	REDIS_CLIENT.Close()
}
