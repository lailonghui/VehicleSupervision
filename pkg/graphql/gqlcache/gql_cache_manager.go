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
	aspectMap    map[string]*GqlCacheAspect
	mutex        *sync.Mutex
	cacheManager *cache.CacheManager
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

//GetGqlCacheAspect 获取gqlCacheAspect
func (m *GqlCacheManager) GetGqlCacheAspect(cacheName string, conf *GqlCacheConf) (*GqlCacheAspect, error) {
	if cacheName == "" {
		return nil, ErrCacheNameEmpty
	}
	aspect, ok := m.aspectMap[cacheName]
	if ok {
		return aspect, nil
	}
	return m.NewGqlCacheAspect(cacheName, conf)
}

//NewGqlCacheAspect 新建gql缓存切面
func (m *GqlCacheManager) NewGqlCacheAspect(cacheName string, conf *GqlCacheConf) (*GqlCacheAspect, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	aspect, ok := m.aspectMap[cacheName]
	if ok {
		return aspect, nil
	}

	return nil, nil
}
