package gqlcache

import "context"

//GqlCacheAspect gql缓存切面
type GqlCacheAspect interface {
	
	//OnPkQuery
	// 主键查询时候，查询缓存，如果缓存有内容，则将缓存内容设置到dest，并返回true
	// 反之，返回false
	OnPkQuery(ctx context.Context, key string, dest interface{}) (bool, error)

	//OnListQuery
	// 列表查询时候，查询缓存，如果缓存有内容，则将缓存内容设置到dest中，并返回true
	// 反之，返回false
	OnListQuery(ctx context.Context, key string, dest interface{}) (bool, error)

	//OnArrgegateQuery
	// 聚合查询时候，查询缓存，如果缓存有内容，则将缓存内容设置到dest中，并返回true
	// 反之，返回false
	OnArrgegateQuery(ctx context.Context, key string, dest interface{}) (bool, error)

	//SetPkQueryCache
	// 设置主键查询的缓存
	SetPkQueryCache(ctx context.Context, key string, data interface{}) error

	//SetListQueryCache
	// 设置列表查询缓存
	SetListQueryCache(ctx context.Context, key string, data interface{}) error

	//SetArrgegateQueryCache
	// 设置聚合查询缓存
	SetArrgegateQueryCache(ctx context.Context, key string, data interface{}) error

	//OnPkRemove
	// 根据主键删除时候，删除对应的缓存
	OnPkRemove(ctx context.Context, key string) error

	//OnListRemove
	// 根据列表删除时候，删除对应的缓存
	OnListRemove(ctx context.Context, key string) error

	//OnInsert
	// 插入记录时候，删除对应的缓存
	OnInsert(ctx context.Context) error

	//OnPkUpdate
	// 根据主键更新时候，删除对应的缓存
	OnPkUpdate(ctx context.Context, key string) error

	//OnListUpdate
	// 根据列表更新时候，删除对应的缓存
	OnListUpdate(ctx context.Context, key string) error
}
