package cache

import (
	"VehicleSupervision/config"
	rc "VehicleSupervision/internal/redis"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	config.Setup("../../config/setting.yaml")
	redisConfig := &goRedis.ClusterOptions{
		Addrs:        config.CONF_INSTANCE.RedisConf.Addresses,
		IdleTimeout:  config.CONF_INSTANCE.RedisConf.Pool.IdleTimeout,
		MinIdleConns: config.CONF_INSTANCE.RedisConf.Pool.MinIdle,
		PoolSize:     config.CONF_INSTANCE.RedisConf.Pool.MaxActive,
	}
	rc.Setup(redisConfig)

}

func TestNewCacheManager(t *testing.T) {

	cm := NewCacheManager(rc.REDIS_CLIENT)
	assert.NotNil(t, cm)
}

func TestCacheManager_GetCache(t *testing.T) {

	cm := NewCacheManager(rc.REDIS_CLIENT)
	assert.NotNil(t, cm)
	c := cm.GetCache("test", false)
	assert.Nil(t, c)
	c = cm.GetCache("test", true)
	assert.NotNil(t, c)
}

func TestCacheManager_NewCache(t *testing.T) {

	cm := NewCacheManager(rc.REDIS_CLIENT)
	assert.NotNil(t, cm)
	c := cm.NewCache("test")
	assert.NotNil(t, c)
	c2 := cm.NewCache("test")
	assert.NotNil(t, c2)
	assert.Equal(t, c, c2)
}
