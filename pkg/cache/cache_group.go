package cache

import "time"

// 缓存组
type Group struct {
	// 组名
	Name string
	// 默认过期时间
	ExpireTime time.Duration
	// 缓存前缀
	CachePrefix string
}
