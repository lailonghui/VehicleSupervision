package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//CONF_INSTANCE 配置信息实例
var CONF_INSTANCE *Conf

//Conf 配置信息
type Conf struct {
	//ServerConf 服务器配置信息
	ServerConf `yaml:"server"`
	//AppConf 应用配置信息
	AppConf `yaml:"app"`
	//DbConf 数据库配置信息
	DbConf `yaml:"db"`
	//LogConf 日志配置信息
	LogConf `yaml:"log"`
	//RedisConf redis配置信息
	RedisConf `yaml:"redis"`
}

//ServerConf 服务器配置
type ServerConf struct {
	//Host 绑定的host
	Host string `yaml:"host"`
	//Port 绑定的端口
	Port int `yaml:"port"`
}

//AppConf 应用配置
type AppConf struct {
	//Name 应用名称
	Name string `yaml:"name"`
}

//DbConf 数据库配置
type DbConf struct {
	//Host 数据库连接地址
	Host string `yaml:"host"`
	//Port 数据库端口
	Port int `yaml:"port"`
	//Username 用户名
	Username string `yaml:"username"`
	//Password 密码
	Password string `yaml:"password"`
	//DbName 数据库连接
	Dbname string `yaml:"dbname"`
	//SslMode ssl模式
	SslMode string `yaml:"sslMode"`
	//TimeZone 时区
	Timezone string `yaml:"timezone"`
	//DbPoolConf 连接池设置
	DbPoolConf `yaml:"pool"`
}

//DbPoolConf 数据库pool配置
type DbPoolConf struct {
	//MaxIdleConn 最多空闲连接
	MaxIdleConn int `yaml:"maxIdleConn"`
	//MaxOpenConn 最多连接数
	MaxOpenConn int `yaml:"maxOpenConn"`
	//MaxLifeTime 连接多久过期
	MaxLifeTime int `yaml:"maxLifeTime"`
}

// 日志配置
type LogConf struct {
	//Path 日志放置位置
	Path string `yaml:"path"`
	//Level 日志级别,支持 debug,warn,info,error
	Level string `yaml:"level"`
	//Mode 日志模式,支持console,file,all
	Mode string `yaml:"mode"`
}

//RedisConf redis配置
type RedisConf struct {
	//Addresses redis集群地址
	Addresses []string `yaml:"addresses"`
	//Pool redis连接池配置
	Pool RedisPoolConf `yaml:"pool"`
}

//RedisPoolConf redis pool 配置
type RedisPoolConf struct {
	//MaxIdle 最少空闲连接数
	MinIdle int `yaml:"minIdle"`
	//MaxActive 最多活跃连接数
	MaxActive int `yaml:"maxActive"`
	//IdleTimeout 空闲连接多久过期
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
