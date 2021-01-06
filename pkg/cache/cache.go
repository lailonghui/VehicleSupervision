package cache

import (
	"time"
)

//Cache 是缓存接口
type Cache interface {



	//Get 获取缓存中保存的值
	// @param cacheKey 缓存的key
	// @return []byte 缓存中的值
	// @return error 错误
	Get(cacheKey string) ([]byte, error)

	//Set 设置缓存中的值
	// @param cacheKey 缓存的key
	// @param value 缓存的值
	// @param expire 过期时间
	// @return error 错误
	Set(cacheKey string, value []byte, expire time.Duration) error

	//Expire 设置缓存的过期时间
	// @param cacheKey 缓存的key
	// @param expire 过期时间
	// @return error 错误
	Expire(cacheKey string, expire time.Duration) error

	//Del 删除缓存
	// @param cacheKey 缓存的key
	// @param expire 过期时间
	// @return error 错误
	Del(cacheKey string) error

	//Clear 清楚缓存组中所有的缓存
	// @return error 错误
	Clear() error
}
