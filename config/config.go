package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var CONF_INSTANCE *Conf

// 配置信息
type Conf struct {
	ServerConf `yaml:"server"`
	AppConf    `yaml:"app"`
	DbConf     `yaml:"db"`
	LogConf    `yaml:"log"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type AppConf struct {
	Name string `yaml:"name"`
}

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

type DbPoolConf struct {
	MaxIdleConn int `yaml:"maxIdleConn"`
	MaxOpenConn int `yaml:"maxOpenConn"`
	MaxLifeTime int `yaml:"maxLifeTime"`
}

type LogConf struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
	Mode  string `yaml:"mode"`
}

func Setup() {
	goPath := os.Getenv("GOPATH")
	filename := goPath + strings.ReplaceAll("/src/lai.com/VehicleSupervision/config/setting.yaml", "/", "\\")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &CONF_INSTANCE)
	if err != nil {
		log.Fatal(err)
	}
}
