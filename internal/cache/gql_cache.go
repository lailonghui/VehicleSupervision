package cache

import (
	"VehicleSupervision/pkg/cache"
	gCache "VehicleSupervision/pkg/graphql/gqlcache"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type GqlCacheConfs struct {
	Confs map[string]GqlCacheConf `yaml:"gqlcache"`
}

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
}

var GqlCacheManager *gCache.GqlCacheManager
var CacheManager *cache.CacheManager

//GqlCacheSetup gqlCache 启动
func GqlCacheSetup(configFile string, client *redis.ClusterClient) {
	// 加载配置文件
	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	var confs *GqlCacheConfs
	err = yaml.Unmarshal(bs, &confs)
	if err != nil {
		panic(err)
	}
	CacheManager = cache.NewCacheManager(client)
	GqlCacheManager = &gCache.GqlCacheManager{CacheManager: CacheManager}
}
