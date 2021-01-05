package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
	"time"
)

// redis缓存
type RedisCache struct {
	// 缓存组
	RedisGroup Group
	// redis实例
	RedisInstance *redisc.Cluster
}

func (r *RedisCache) GetConn() *redis.Conn {
	c := r.RedisInstance.Get()
	return &c
}

func (r *RedisCache) Get(cacheKey string) ([]byte, error) {
	panic("implement me")
}

func (r *RedisCache) Set(cacheKey string, value []byte, expire time.Duration) error {
	panic("implement me")
}

func (r *RedisCache) Expire(cacheKey string, expire time.Duration) error {
	panic("implement me")
}

func (r *RedisCache) Del(cacheKey string) error {
	panic("implement me")
}

func (r *RedisCache) Clear() error {
	panic("implement me")
}
