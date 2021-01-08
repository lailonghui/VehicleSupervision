package cache

import (
	"VehicleSupervision/config"
	"VehicleSupervision/internal/redis"
	goRedis "github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var client *goRedis.ClusterClient

var gqlCacheConfigFile string

func init() {
	cf := config.Setup("../../config/setting.yaml")
	redis.Setup(&goRedis.ClusterOptions{
		Addrs:        cf.RedisConf.Addresses,
		IdleTimeout:  time.Duration(cf.RedisConf.Pool.IdleTimeout) * time.Millisecond,
		MinIdleConns: cf.RedisConf.Pool.MinIdle,
		PoolSize:     cf.RedisConf.Pool.MaxActive,
	})
	gqlCacheConfigFile = "../../config/gqlcache.yaml"
}

func TestGqlCacheSetup(t *testing.T) {
	GqlCacheSetup(gqlCacheConfigFile, client)
}
