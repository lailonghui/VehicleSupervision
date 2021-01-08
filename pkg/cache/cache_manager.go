package cache

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

//CacheManager 缓存管理
type CacheManager struct {
	CacheMap    map[string]*Cacher
	mutex       *sync.Mutex
	redisClient *redis.ClusterClient
}

//NewCacheManager 新建缓存管理实例
func NewCacheManager(redisClient *redis.ClusterClient) *CacheManager {
	return &CacheManager{
		CacheMap:    make(map[string]*Cacher, 0),
		mutex:       &sync.Mutex{},
		redisClient: redisClient,
	}
}

//NewCache 新建缓存
func (m *CacheManager) NewCache(cacheName string) *Cacher {
	cacher, ok := m.CacheMap[cacheName]
	if ok {
		return cacher
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	cacher, ok = m.CacheMap[cacheName]
	if ok {
		return cacher
	}
	return NewCacher(cacheName, m.redisClient)
}

//GetCache 获取缓存
func (m *CacheManager) GetCache(cacheName string, autoNew bool) *Cacher {
	cacher := m.CacheMap[cacheName]
	if autoNew && cacher == nil {
		return m.NewCache(cacheName)
	}
	return cacher
}
