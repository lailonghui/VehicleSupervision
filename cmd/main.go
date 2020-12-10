package main

import (
	"VehicleSupervision/config"
	"VehicleSupervision/internal/db"
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

func setup() {
	// 加载配置信息
	config.Setup(configFile)
	// 加载日志
	logger.Setup()
	defer logger.Sync()
	logger.Info("加载成功")
	// 启动db
	db.Setup()
	// 启动服务器
	server.Setup()

}
