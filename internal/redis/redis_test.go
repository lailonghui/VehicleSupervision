package redis

import (
	"VehicleSupervision/config"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	config.Setup("../../config/setting.yaml")
	//Setup()
}

func TestSetup(t *testing.T) {
	ctx := context.Background()
	key := "test"
	value := "hello world"
	err := REDIS_CLIENT.Set(ctx, key, value, -1).Err()
	assert.Nil(t, err)
	val, err := REDIS_CLIENT.Get(ctx, key).Result()
	assert.Nil(t, err)
	assert.Equal(t, value, val)
	err = REDIS_CLIENT.Del(ctx, key).Err()
	assert.Nil(t, err)
}


func TestSetup2(t *testing.T) {
	key := "test"
	value := "hello world"
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.3.130:6379",
		IdleTimeout:  time.Duration(config.CONF_INSTANCE.RedisConf.Pool.IdleTimeout) * time.Millisecond,
		MinIdleConns: config.CONF_INSTANCE.RedisConf.Pool.MinIdle,
		PoolSize:     config.CONF_INSTANCE.RedisConf.Pool.MaxActive,
	})
	client.Ping(ctx)
	err := REDIS_CLIENT.Set(ctx, key, value, -1).Err()
	assert.Nil(t, err)
	val, err := REDIS_CLIENT.Get(ctx, key).Result()
	assert.Nil(t, err)
	assert.Equal(t, value, val)
	err = REDIS_CLIENT.Del(ctx, key).Err()
	assert.Nil(t, err)
}