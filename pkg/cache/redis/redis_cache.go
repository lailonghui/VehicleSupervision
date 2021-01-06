package redis

import (
	"VehicleSupervision/pkg/cache"
	"time"
)

// redis缓存
type RedisCache struct {
	// 缓存组
	RedisGroup cache.Group
	// redis实例
	RedisInstance *redisc.Cluster
	// 序列化方式
	Serialization *RedisSerialization
}

func (r *RedisCache) GetConn() redis.Conn {
	c := r.RedisInstance.Get()
	return c
}

func (r *RedisCache) Get(cacheKey string) ([]byte, error) {
	c := r.GetConn()
	defer c.Close()
	bs, err := redis.Bytes(c.Do("Get", cacheKey))
	return bs, err
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
