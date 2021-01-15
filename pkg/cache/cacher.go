package cache

import (
	"context"
	rc "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

//Cacher 缓存信息
type Cacher struct {

	//KeyPrefix key前缀
	KeyPrefix string

	// redis客户端
	RedisClient redis.UniversalClient

	//Cache 缓存实例
	Cache *rc.Cache

	mutex *sync.Mutex
}

//NewCacher 新建Cacher
func NewCacher(keyPrefix string, redisClient redis.UniversalClient) *Cacher {
	c := rc.New(&rc.Options{
		Redis: redisClient,
	})
	return &Cacher{
		KeyPrefix:   keyPrefix,
		RedisClient: redisClient,
		Cache:       c,
		mutex:       &sync.Mutex{},
	}
}

//getKeySetKey 获取keyset中key的值
func (r *Cacher) getKeySetKey(ctx context.Context) string {
	return "keyset:" + r.KeyPrefix
}

//getKey 获取缓存的key
func (r *Cacher) getKey(ctx context.Context, cacheKey string) string {
	key := r.KeyPrefix + ":" + cacheKey
	r.RedisClient.SAdd(ctx, r.getKeySetKey(ctx), key)
	return key
}

//Get 获取缓存内容
func (r *Cacher) Get(ctx context.Context, cacheKey string, dest interface{}) (bool, error) {
	err := r.Cache.Get(ctx, r.getKey(ctx, cacheKey), dest)
	if err != nil {
		if err == rc.ErrCacheMiss {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

//Set 设置缓存的值
func (r *Cacher) Set(ctx context.Context, cacheKey string, value interface{}, expire time.Duration) error {

	return r.Cache.Set(&rc.Item{
		Ctx:            ctx,
		Key:            r.getKey(ctx, cacheKey),
		Value:          value,
		TTL:            expire,
		SkipLocalCache: true,
	})
}

//Del 删除缓存的值
func (r *Cacher) Del(ctx context.Context, cacheKey string) error {
	return r.Cache.Delete(ctx, r.getKey(ctx, cacheKey))
}

//BatchDel 批量删除
func (r *Cacher) BatchDel(ctx context.Context, cacheKeys []string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	setKeys := make([]interface{}, 0, len(cacheKeys))
	// 使用管道发送删除命令
	pipe := r.RedisClient.Pipeline()
	for i := range cacheKeys {
		pipe.Del(ctx, cacheKeys[i])
		setKeys = append(setKeys, cacheKeys[i])
	}
	pipe.SRem(ctx, r.getKeySetKey(ctx), setKeys...)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

//Clear 清空缓存
func (r *Cacher) Clear(ctx context.Context) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var cursor uint64
	for true {
		v, c, err := r.RedisClient.SScan(ctx, r.getKeySetKey(ctx), cursor, "", 1000).Result()
		cursor = c
		if err != nil {
			if err == redis.Nil {
				return nil
			}
			return err
		}
		if v == nil {
			return nil
		}

		if len(v) == 0 {
			break
		}
		vvi := make([]interface{}, 0, len(v))
		// 使用管道发送删除命令
		pipe := r.RedisClient.Pipeline()
		for j := range v {
			pipe.Del(ctx, v[j])
			vvi = append(vvi, v[j])
		}
		pipe.SRem(ctx, r.getKeySetKey(ctx), vvi...)
		_, err = pipe.Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

//IsEmpty 判断缓存是否为空
func (r Cacher) IsEmpty(ctx context.Context) (bool, error) {
	i, err := r.RedisClient.SCard(ctx, r.getKeySetKey(ctx)).Result()
	if err != nil {
		if err == redis.Nil {
			return true, nil
		}
		return false, err
	}
	if i <= 0 {
		return false, nil
	}
	return true, nil
}
