package main

import (
	"system-manage/config"
	"system-manage/internal/db"
	"system-manage/internal/server"
	"system-manage/pkg/logger"
)

func main() {
	setup()
}

func setup() {
	// 加载配置信息
	config.Setup()
	// 加载日志
	logger.Setup()
	defer logger.Sync()
	logger.Info("加载成功")
	// 启动db
	db.Setup()
	// 启动服务器
	server.Setup()
}
