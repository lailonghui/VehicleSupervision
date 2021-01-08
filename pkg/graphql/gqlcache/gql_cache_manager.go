package gqlcache

import (
	"VehicleSupervision/pkg/cache"
	"errors"
	"sync"
	"time"
)

var (
	ErrCacheNameEmpty = errors.New("cache name can not be empty")
)

//GqlCacheManager gql缓存管理
type GqlCacheManager struct {
	AspectMap    map[string]*GqlCacheAspect
	Mutex        *sync.Mutex
	CacheManager *cache.CacheManager
}

//GqlCacheConfs gql缓存配置信息
type GqlCacheConf struct {
	// 是否启用主键缓存
	EnablePkCache bool
	// 是否启用列表缓存
	EnableListCache bool
	// 是否启用聚合查询缓存
	EnableAggregateCache bool
	// 主键缓存过期时间
	PkCacheTimeout time.Duration
	// 列表缓存过期时间
	ListCacheTimeout time.Duration
	// 聚合查询缓存过期时间
	AggregateCacheTimeout time.Duration
}

//NewGqlCacheManager 新建 GqlCacheManager 实例
func NewGqlCacheManager(cacheManager *cache.CacheManager) *GqlCacheManager {
	aspectMap := make(map[string]*GqlCacheAspect, 0)
	return &GqlCacheManager{
		AspectMap:    aspectMap,
		Mutex:        &sync.Mutex{},
		CacheManager: cacheManager,
	}
}

//GetGqlCacheAspect 获取gqlCacheAspect
func (m *GqlCacheManager) GetGqlCacheAspect(cacheName string, conf *GqlCacheConf) (*GqlCacheAspect, error) {
	if cacheName == "" {
		return nil, ErrCacheNameEmpty
	}
	aspect, ok := m.AspectMap[cacheName]
	if ok {
		return aspect, nil
	}
	return m.NewGqlCacheAspect(cacheName, conf)
}

//NewGqlCacheAspect 新建gql缓存切面
func (m *GqlCacheManager) NewGqlCacheAspect(cacheName string, conf *GqlCacheConf) (*GqlCacheAspect, error) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	aspect, ok := m.AspectMap[cacheName]
	if ok {
		return aspect, nil
	}

	return nil, nil
}
