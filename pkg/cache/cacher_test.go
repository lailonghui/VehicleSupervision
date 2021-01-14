package cache

import (
	"VehicleSupervision/config"
	rc "VehicleSupervision/internal/redis"
	"context"
	"fmt"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	cf := config.Setup("../../config/setting.yaml")
	//rc.Setup()
	// 启动redis
	redisConfig := &goRedis.UniversalOptions{
		Addrs:        cf.RedisConf.Addresses,
		IdleTimeout:  cf.RedisConf.Pool.IdleTimeout,
		MinIdleConns: cf.RedisConf.Pool.MinIdle,
		PoolSize:     cf.RedisConf.Pool.MaxActive,
	}
	 rc.Setup(redisConfig)
}

func TestNewCacher(t *testing.T) {

	c := NewCacher("test", rc.REDIS_CLIENT)
	assert.NotNil(t, c)
}

func TestCacher_Set(t *testing.T) {
	fmt.Println("test")
	var ctx context.Context = context.Background()

	c := NewCacher("test005", rc.REDIS_CLIENT)

	t.Run("test not exist key", func(t *testing.T) {
		var dest interface{}
		exist, err := c.Get(ctx, "dsda", dest)
		assert.Nil(t, err)
		assert.False(t, exist)
		assert.Nil(t, dest)
	})
	t.Run("test exist key", func(t *testing.T) {
		var data string = "hello world"
		var key string = "key1"
		err := c.Set(ctx, key, data, -1)
		assert.Nil(t, err)
		var dest string
		exist, err := c.Get(ctx, key, &dest)
		assert.Nil(t, err)
		assert.True(t, exist)
		assert.Equal(t, data, dest)
	})
}

func TestCacher_Clear(t *testing.T) {
	var ctx context.Context = context.Background()

	c := NewCacher("enterprise:upk", rc.REDIS_CLIENT)
	assert.NotNil(t, c)
	err := c.Clear(ctx)
	assert.Nil(t, err)
}
