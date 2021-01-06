package cache

import (
	"VehicleSupervision/config"
	rc "VehicleSupervision/internal/redis"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)


func init() {
	config.Setup("../../config/setting.yaml")
	rc.Setup()

}

func TestNewCacheManager(t *testing.T) {
	var ctx context.Context = context.Background()

	cm := NewCacheManager(ctx, rc.REDIS_CLIENT)
	assert.NotNil(t, cm)
}

func TestCacheManager_GetCache(t *testing.T) {
	var ctx context.Context = context.Background()

	cm := NewCacheManager(ctx, rc.REDIS_CLIENT)
	assert.NotNil(t, cm)
	c := cm.GetCache(ctx, "test", false)
	assert.Nil(t, c)
	c = cm.GetCache(ctx, "test", true)
	assert.NotNil(t, c)
}

func TestCacheManager_NewCache(t *testing.T) {
	var ctx context.Context = context.Background()

	cm := NewCacheManager(ctx, rc.REDIS_CLIENT)
	assert.NotNil(t, cm)
	c := cm.NewCache(ctx, "test")
	assert.NotNil(t, c)
	c2 := cm.NewCache(ctx, "test")
	assert.NotNil(t, c2)
	assert.Equal(t, c, c2)
}

