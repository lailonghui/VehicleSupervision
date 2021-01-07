package cache

import (
	"VehicleSupervision/config"
	rc "VehicleSupervision/internal/redis"
	"context"
	"github.com/dgryski/trifles/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	config.Setup("../../config/setting.yaml")
	rc.Setup()

}

func TestNewCacher(t *testing.T) {
	var ctx context.Context = context.Background()

	c := NewCacher(ctx, "test", rc.REDIS_CLIENT)
	assert.NotNil(t, c)
}

func TestCacher_Set(t *testing.T) {
	var ctx context.Context = context.Background()

	c := NewCacher(ctx, "test005", rc.REDIS_CLIENT)
	assert.NotNil(t, c)
	for i := 0; i < 10000; i++ {
		key := uuid.UUIDv4()
		err := c.Set(ctx, key, key, -1)

		assert.Nil(t, err)
	}

}

func TestCacher_Clear(t *testing.T) {
	var ctx context.Context = context.Background()

	c := NewCacher(ctx, "test005", rc.REDIS_CLIENT)
	assert.NotNil(t, c)
	err := c.Clear(ctx)
	assert.Nil(t, err)
}
