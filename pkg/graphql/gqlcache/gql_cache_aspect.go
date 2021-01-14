package gqlcache

import (
	"VehicleSupervision/pkg/cache"
	"context"
)

type GqlCacheAspect struct {
	GqlCacheConf   *GqlCacheConf
	TableName      string
	PkCacher       *cache.Cacher
	UnionPkCacher  *cache.Cacher
	ListCacher     *cache.Cacher
	AggregateCache *cache.Cacher
}

//OnPkQuery
// 主键查询时候，查询缓存，如果缓存有内容，则将缓存内容设置到dest，并返回true
// 反之，返回false
func (n *GqlCacheAspect) OnPkQuery(ctx context.Context, key string, dest interface{}) (bool, error) {
	if !n.GqlCacheConf.EnablePkCache {
		return false, nil
	}
	return n.PkCacher.Get(ctx, key, dest)
}

//OnUnionPkQuery
// 联合主键查询时候，查询缓存，如果缓存有内容，则将缓存内容设置到dest，并返回true
// 反之，返回false
func (n *GqlCacheAspect) OnUnionPkQuery(ctx context.Context, key string, dest interface{}) (bool, error) {
	if !n.GqlCacheConf.EnablePkCache {
		return false, nil
	}
	return n.UnionPkCacher.Get(ctx, key, dest)
}

//OnListQuery
// 列表查询时候，查询缓存，如果缓存有内容，则将缓存内容设置到dest中，并返回true
// 反之，返回false
func (n *GqlCacheAspect) OnListQuery(ctx context.Context, key string, dest interface{}) (bool, error) {
	if !n.GqlCacheConf.EnableListCache {
		return false, nil
	}
	return n.ListCacher.Get(ctx, key, dest)
}

//OnArrgegateQuery
// 聚合查询时候，查询缓存，如果缓存有内容，则将缓存内容设置到dest中，并返回true
// 反之，返回false
func (n *GqlCacheAspect) OnArrgegateQuery(ctx context.Context, key string, dest interface{}) (bool, error) {
	if !n.GqlCacheConf.EnableAggregateCache {
		return false, nil
	}
	return n.AggregateCache.Get(ctx, key, dest)
}

//SetPkQueryCache
// 设置主键查询的缓存
func (n *GqlCacheAspect) SetPkQueryCache(ctx context.Context, key string, data interface{}) error {
	if !n.GqlCacheConf.EnablePkCache {
		return nil
	}
	return n.PkCacher.Set(ctx, key, data, n.GqlCacheConf.PkCacheTimeout)
}

//SetPkQueryCache
// 设置主键查询的缓存
func (n *GqlCacheAspect) SetUnionPkQueryCache(ctx context.Context, key string, data interface{}) error {
	if !n.GqlCacheConf.EnablePkCache {
		return nil
	}
	return n.UnionPkCacher.Set(ctx, key, data, n.GqlCacheConf.PkCacheTimeout)
}

//SetNotExistPkQueryCache
// 设置不存在的主键的缓存
func (n *GqlCacheAspect) SetNotExistPkQueryCache(ctx context.Context, key string, data interface{}) error {
	if !n.GqlCacheConf.EnablePkCache {
		return nil
	}
	return n.PkCacher.Set(ctx, key, data, n.GqlCacheConf.NotExistRecordTimeout)
}

//SetNotExistUnionPkQueryCache
// 设置不存在的联合主键的缓存
func (n *GqlCacheAspect) SetNotExistUnionPkQueryCache(ctx context.Context, key string, data interface{}) error {
	if !n.GqlCacheConf.EnablePkCache {
		return nil
	}
	return n.UnionPkCacher.Set(ctx, key, data, n.GqlCacheConf.NotExistRecordTimeout)
}

//SetListQueryCache
// 设置列表查询缓存
func (n *GqlCacheAspect) SetListQueryCache(ctx context.Context, key string, data interface{}) error {
	if !n.GqlCacheConf.EnableListCache {
		return nil
	}

	return n.ListCacher.Set(ctx, key, data, n.GqlCacheConf.ListCacheTimeout)
}

//SetArrgegateQueryCache
// 设置聚合查询缓存
func (n *GqlCacheAspect) SetArrgegateQueryCache(ctx context.Context, key string, data interface{}) error {
	if !n.GqlCacheConf.EnableAggregateCache {
		return nil
	}
	return n.AggregateCache.Set(ctx, key, data, n.GqlCacheConf.AggregateCacheTimeout)
}

//OnPkRemove
// 根据主键删除时候，删除对应的缓存
func (n *GqlCacheAspect) OnPkRemove(ctx context.Context, key string) error {
	if n.GqlCacheConf.EnablePkCache {
		err := n.PkCacher.Del(ctx, key)
		if err != nil {
			return err
		}
		err = n.UnionPkCacher.Clear(ctx)
		if err != nil {
			return err
		}
	}
	if n.GqlCacheConf.EnableListCache {
		err := n.ListCacher.Clear(ctx)
		if err != nil {
			return err
		}
	}
	if n.GqlCacheConf.EnableAggregateCache {
		err := n.AggregateCache.Clear(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

//OnUnionPkRemove
// 根据联合主键删除时候，删除对应的缓存
func (n *GqlCacheAspect) OnUnionPkRemove(ctx context.Context, key string) error {
	if n.GqlCacheConf.EnablePkCache {
		err := n.UnionPkCacher.Del(ctx, key)
		if err != nil {
			return err
		}
	}
	if n.GqlCacheConf.EnableListCache {
		err := n.ListCacher.Clear(ctx)
		if err != nil {
			return err
		}
	}
	if n.GqlCacheConf.EnableAggregateCache {
		err := n.AggregateCache.Clear(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

//OnListRemove
// 根据列表删除时候，删除对应的缓存
func (n *GqlCacheAspect) OnListRemove(ctx context.Context, ids []string, unionIds []string) error {
	if n.GqlCacheConf.EnablePkCache {
		if len(ids) > 0 {
			err := n.PkCacher.BatchDel(ctx, ids)
			if err != nil {
				return err
			}
		}
		if len(unionIds) > 0 {
			err := n.UnionPkCacher.Clear(ctx)
			if err != nil {
				return err
			}
		}
	}
	if n.GqlCacheConf.EnableListCache {
		err := n.ListCacher.Clear(ctx)
		if err != nil {
			return err
		}
	}
	if n.GqlCacheConf.EnableAggregateCache {
		err := n.AggregateCache.Clear(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

//OnInsert
// 插入记录时候，删除对应的缓存
func (n *GqlCacheAspect) OnInsert(ctx context.Context) error {
	if n.GqlCacheConf.EnableListCache {
		err := n.ListCacher.Clear(ctx)
		if err != nil {
			return err
		}
	}
	if n.GqlCacheConf.EnableAggregateCache {
		err := n.AggregateCache.Clear(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

//OnPkUpdate
// 根据主键更新时候，删除对应的缓存
func (n *GqlCacheAspect) OnPkUpdate(ctx context.Context, key string) error {
	if n.GqlCacheConf.EnablePkCache {
		err := n.PkCacher.Del(ctx, key)
		if err != nil {
			return err
		}
		err = n.UnionPkCacher.Clear(ctx)
		if err != nil {
			return err
		}
	}
	if n.GqlCacheConf.EnableListCache {
		err := n.ListCacher.Clear(ctx)
		if err != nil {
			return err
		}
	}
	if n.GqlCacheConf.EnableAggregateCache {
		err := n.AggregateCache.Clear(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

//OnListUpdate
// 根据列表更新时候，删除对应的缓存
func (n *GqlCacheAspect) OnListUpdate(ctx context.Context, ids []string, unionIds []string) error {
	if n.GqlCacheConf.EnablePkCache {
		if len(ids) > 0 {
			err := n.PkCacher.BatchDel(ctx, ids)
			if err != nil {
				return err
			}
		}
		if len(unionIds) > 0 {
			err := n.UnionPkCacher.Clear(ctx)
			if err != nil {
				return err
			}
		}
	}
	if n.GqlCacheConf.EnableListCache {
		err := n.ListCacher.Clear(ctx)
		if err != nil {
			return err
		}
	}
	if n.GqlCacheConf.EnableAggregateCache {
		err := n.AggregateCache.Clear(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
