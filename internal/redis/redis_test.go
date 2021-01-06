package redis

import (
	"VehicleSupervision/config"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	config.Setup("../../config/setting.yaml")
	Setup()
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
}
