package cache

import (
	"VehicleSupervision/pkg/cache"
	gCache "VehicleSupervision/pkg/graphql/gqlcache"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

//GqlCacheConfs gql缓存配置信息
type GqlCacheConfs struct {
	// 是否启用gqlCache
	Enable bool `yaml:"enable"`
	// 缓存配置信息
	GqlCaches map[string]GqlCacheConf `yaml:"caches"`
	// 默认缓存配置信息
	DefaultCacheConf GqlCacheConf `yaml:"defaultCache"`
}

//GqlCacheConf gql缓存配置信息
type GqlCacheConf struct {
	// 是否启用主键缓存
	EnablePkCache bool `yaml:"enablePkCache"`
	// 是否启用列表缓存
	EnableListCache bool `yaml:"enableListCache"`
	// 是否启用聚合查询缓存
	EnableAggregateCache bool `yaml:"enableAggregateCache"`
	// 主键缓存过期时间
	PkCacheTimeout time.Duration `yaml:"pkCacheTimeout"`
	// 列表缓存过期时间
	ListCacheTimeout time.Duration `yaml:"listCacheTimeout"`
	// 聚合查询缓存过期时间
	AggregateCacheTimeout time.Duration `yaml:"aggregateCacheTimeout"`
	// 不存在记录缓存过期时间
	NotExistRecordTimeout time.Duration `yaml:"notExistRecordTimeout"`
}

var GqlCacheManager *gCache.GqlCacheManager
var CacheManager *cache.CacheManager
var conf *GqlCacheConfs

//GqlCacheSetup gqlCache 启动
func GqlCacheSetup(configFile string, client redis.UniversalClient) {
	// 加载配置文件
	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(bs, &conf)
	if err != nil {
		log.Fatal(err)
	}
	CacheManager = cache.NewCacheManager(client)
	GqlCacheManager = gCache.NewGqlCacheManager(CacheManager)

	for tableName, cacheConf := range conf.GqlCaches {
		_, err := GqlCacheManager.NewGqlCacheAspect(tableName, &gCache.GqlCacheConf{
			EnablePkCache:         cacheConf.EnablePkCache,
			EnableListCache:       cacheConf.EnableListCache,
			EnableAggregateCache:  cacheConf.EnableAggregateCache,
			PkCacheTimeout:        cacheConf.PkCacheTimeout,
			ListCacheTimeout:      cacheConf.ListCacheTimeout,
			AggregateCacheTimeout: cacheConf.AggregateCacheTimeout,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

//GetGqlCacheAspect 获取GqlCacheAspect实例
func GetGqlCacheAspect(tableName string) (*gCache.GqlCacheAspect, error) {
	cacheAspect, err := GqlCacheManager.GetGqlCacheAspect(tableName)
	if err != nil {
		if err == gCache.ErrNotCacheConfig {
			return GqlCacheManager.GetGqlCacheAspectIfNotExistNew(tableName, &gCache.GqlCacheConf{
				EnablePkCache:         conf.DefaultCacheConf.EnablePkCache,
				EnableListCache:       conf.DefaultCacheConf.EnableListCache,
				EnableAggregateCache:  conf.DefaultCacheConf.EnableAggregateCache,
				PkCacheTimeout:        conf.DefaultCacheConf.PkCacheTimeout,
				ListCacheTimeout:      conf.DefaultCacheConf.ListCacheTimeout,
				AggregateCacheTimeout: conf.DefaultCacheConf.AggregateCacheTimeout,
			})
		}
		return nil, err
	}
	return cacheAspect, nil
}
