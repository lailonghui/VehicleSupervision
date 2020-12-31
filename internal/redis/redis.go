package redis

import (
	"VehicleSupervision/config"
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
	"time"
)

var REDIS *redisc.Cluster

func Setup() {
	REDIS = &redisc.Cluster{
		StartupNodes: config.CONF_INSTANCE.RedisConf.Addresses,
		DialOptions:  []redis.DialOption{redis.DialConnectTimeout(5 * time.Second)},
		CreatePool: func(address string, options ...redis.DialOption) (*redis.Pool, error) {
			return &redis.Pool{
				MaxIdle:     config.CONF_INSTANCE.Pool.MaxIdle,
				MaxActive:   config.CONF_INSTANCE.Pool.MaxActive,
				IdleTimeout: time.Duration(config.CONF_INSTANCE.Pool.IdleTimeout) * time.Millisecond,
				Dial: func() (redis.Conn, error) {
					return redis.Dial("tcp", address, options...)
				},
				TestOnBorrow: func(c redis.Conn, t time.Time) error {
					_, err := c.Do("PING")
					return err
				},
			}, nil
		},
	}
}

func Close() {
	REDIS.Close()
}
