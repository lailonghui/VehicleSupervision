package main

import (
	"VehicleSupervision/config"
	"VehicleSupervision/internal/cache"
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/redis"
	goRedis "github.com/go-redis/redis/v8"

	"VehicleSupervision/internal/server"
	"VehicleSupervision/pkg/logger"
	"flag"
)

var configFile string

func main() {
	resolveCmdParam()

	setup()
}

// 解析命令行参数
func resolveCmdParam() {
	flag.StringVar(&configFile, "f", "config/setting.yaml", "配置文件")
	flag.Parse()
}

// 启动
func setup() {
	// 加载配置信息
	cf := config.Setup(configFile)
	// 加载日志
	logger.Setup(&logger.ConfigOption{
		Path:    cf.LogConf.Path,
		Level:   cf.LogConf.Level,
		Mode:    cf.LogConf.Mode,
		AppName: cf.AppConf.Name,
	})
	defer logger.Sync()
	// 启动db
	dbConfig := db.ConfigOption{}
	dbConfig.PoolInfo = db.PoolInfo{
		MaxIdleConn: cf.DbPoolConf.MaxIdleConn,
		MaxOpenConn: cf.DbPoolConf.MaxOpenConn,
		MaxLifeTime: cf.DbPoolConf.MaxLifeTime,
	}
	dbConfig.MasterNode = db.NodeInfo{
		Host:     cf.Master.Host,
		Username: cf.Master.Username,
		Password: cf.Master.Password,
		Dbname:   cf.Master.Dbname,
		Port:     cf.Master.Port,
		Sslmode:  cf.Master.SslMode,
		Timezone: cf.Master.Timezone,
	}
	if len(cf.Slaves) > 0 {
		dbConfig.SlaveNodes = make([]db.NodeInfo, 0)
		for i := 0; i < len(cf.Slaves); i++ {
			slave := cf.Slaves[i]
			dbConfig.SlaveNodes = append(dbConfig.SlaveNodes, db.NodeInfo{
				Host:     slave.Host,
				Username: slave.Username,
				Password: slave.Password,
				Dbname:   slave.Dbname,
				Port:     slave.Port,
				Sslmode:  slave.SslMode,
				Timezone: slave.Timezone,
			})
		}
	}
	db.Setup(dbConfig)
	// 启动redis
	redisConfig := &goRedis.ClusterOptions{
		Addrs:        cf.RedisConf.Addresses,
		IdleTimeout:  cf.RedisConf.Pool.IdleTimeout,
		MinIdleConns: cf.RedisConf.Pool.MinIdle,
		PoolSize:     cf.RedisConf.Pool.MaxActive,
	}
	rc := redis.Setup(redisConfig)
	defer redis.Close()
	// 启动gqlcache
	cache.GqlCacheSetup(cf.GqlConf.CacheConfigFile, rc)
	// 启动服务器
	server.Setup(cf.ServerConf.Host, cf.ServerConf.Port)

}
