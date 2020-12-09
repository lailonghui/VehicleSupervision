package db

import (
	"VehicleSupervision/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	l "VehicleSupervision/pkg/logger"
	"time"
)

var DB *gorm.DB

// 初始化db
func Setup() {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.CONF_INSTANCE.DbConf.Host,
		config.CONF_INSTANCE.DbConf.Username,
		config.CONF_INSTANCE.DbConf.Password,
		config.CONF_INSTANCE.DbConf.Dbname,
		config.CONF_INSTANCE.DbConf.Port,
		config.CONF_INSTANCE.DbConf.SslMode,
		config.CONF_INSTANCE.DbConf.Timezone)

	var err error

	newLogger := l.GormLogger{
		LogLevel:      logger.Info,
		SlowThreshold: 100 * time.Millisecond,
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}

	setConnectionPool()
	autoMigrate()
}

// 设置连接处
func setConnectionPool() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.CONF_INSTANCE.DbConf.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.CONF_INSTANCE.DbConf.MaxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(config.CONF_INSTANCE.DbConf.MaxLifeTime) * time.Millisecond)
}

// 维护表结构
func autoMigrate() {

}
