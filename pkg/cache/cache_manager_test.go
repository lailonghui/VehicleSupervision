package cache

import (
	"VehicleSupervision/config"
	rc "VehicleSupervision/internal/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)


func init() {
	config.Setup("../../config/setting.yaml")
	//rc.Setup()

}

func TestNewCacheManager(t *testing.T) {

	cm := NewCacheManager( rc.REDIS_CLIENT)
	assert.NotNil(t, cm)
}

func TestCacheManager_GetCache(t *testing.T) {

	cm := NewCacheManager( rc.REDIS_CLIENT)
	assert.NotNil(t, cm)
	c := cm.GetCache( "test", false)
	assert.Nil(t, c)
	c = cm.GetCache( "test", true)
	assert.NotNil(t, c)
}

func TestCacheManager_NewCache(t *testing.T) {

	cm := NewCacheManager( rc.REDIS_CLIENT)
	assert.NotNil(t, cm)
	c := cm.NewCache( "test")
	assert.NotNil(t, c)
	c2 := cm.NewCache( "test")
	assert.NotNil(t, c2)
	assert.Equal(t, c, c2)
}

