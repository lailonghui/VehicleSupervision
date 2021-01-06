package redis

import (
	"VehicleSupervision/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var REDIS_CLIENT *redis.ClusterClient

func Setup() {
	var ctx = context.Background()
	REDIS_CLIENT = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        config.CONF_INSTANCE.RedisConf.Addresses,
		IdleTimeout:  time.Duration(config.CONF_INSTANCE.RedisConf.Pool.IdleTimeout) * time.Millisecond,
		MinIdleConns: config.CONF_INSTANCE.RedisConf.Pool.MinIdle,
		PoolSize:     config.CONF_INSTANCE.RedisConf.Pool.MaxActive,
	})
	REDIS_CLIENT.Ping(ctx)

}

func Close() {
}
