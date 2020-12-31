package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var CONF_INSTANCE *Conf

// 配置信息
type Conf struct {
	ServerConf `yaml:"server"`
	AppConf    `yaml:"app"`
	DbConf     `yaml:"db"`
	LogConf    `yaml:"log"`
	RedisConf  `yaml:"redis"`
}

// 服务器配置
type ServerConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// 应用配置
type AppConf struct {
	Name string `yaml:"name"`
}

// 数据库配置
type DbConf struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Dbname     string `yaml:"dbname"`
	SslMode    string `yaml:"sslMode"`
	Timezone   string `yaml:"timezone"`
	DbPoolConf `yaml:"pool"`
}

// 数据库pool配置
type DbPoolConf struct {
	MaxIdleConn int `yaml:"maxIdleConn"`
	MaxOpenConn int `yaml:"maxOpenConn"`
	MaxLifeTime int `yaml:"maxLifeTime"`
}

// 日志配置
type LogConf struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
	Mode  string `yaml:"mode"`
}

// redis配置
type RedisConf struct {
	Addresses []string      `yaml:"addresses"`
	Database  int           `yaml:"database"`
	Pool      RedisPoolConf `yaml:"pool"`
}

// redis pool 配置
type RedisPoolConf struct {
	MaxIdle     int `yaml:"maxIdle"`
	MaxActive   int `yaml:"maxActive"`
	IdleTimeout int `yaml:"idleTimeout"`
}

func Setup(configFile string) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &CONF_INSTANCE)
	if err != nil {
		log.Fatal(err)
	}
}
