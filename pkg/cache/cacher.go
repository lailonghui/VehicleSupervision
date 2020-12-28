package cache

import (
	"time"
)

// 缓存接口
type Cacher interface {

	// 获取缓存中保存的值
	// @param cacheKey 缓存的key
	// @param cacheGroup
	// @return []byte 缓存中的值
	// @return error 错误
	get(cacheKey, cacheGroup string) ([]byte, error)

	// 设置缓存中的值
	// @param cacheKey 缓存的key
	// @param cacheGroup 缓存组
	// @param value 缓存的值
	// @param expire 过期时间
	// @return error 错误
	set(cacheKey, cacheGroup string, value []byte, expire time.Duration) error

	// 设置缓存的过期时间
	// @param cacheKey 缓存的key
	// @param cacheGroup 缓存组
	// @param expire 过期时间
	// @return error 错误
	expire(cacheKey, cacheGroup string, expire time.Duration) error

	// 删除缓存
	// @param cacheKey 缓存的key
	// @param cacheGroup 缓存组
	// @param expire 过期时间
	// @return error 错误
	del(cacheKey, cacheGroup string) error

	// 清楚缓存组中所有的缓存
	// @param cacheGroup 缓存组
	// @return error 错误
	clear(cacheGroup string) error
}
