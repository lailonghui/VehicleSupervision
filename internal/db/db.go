package db

import (
	"VehicleSupervision/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"

	l "VehicleSupervision/pkg/logger"
	"time"
)

var DB *gorm.DB

type ConfigOption struct {
	// 主节点
	MasterNode NodeInfo
	// 从节点
	SlaveNodes []NodeInfo
	// 连接池
	PoolInfo PoolInfo
}

type NodeInfo struct {
	Host     string
	Username string
	Password string
	Dbname   string
	Port     int
	Sslmode  string
	Timezone string
}

func (i NodeInfo) getDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		i.Host,
		i.Username,
		i.Password,
		i.Dbname,
		i.Port,
		i.Sslmode,
		i.Timezone)
}

type PoolInfo struct {
	//MaxIdleConn 最多空闲连接
	MaxIdleConn int
	//MaxOpenConn 最多连接数
	MaxOpenConn int
	//MaxLifeTime 连接多久过期
	MaxLifeTime time.Duration
}

// 初始化db
func Setup(config ConfigOption) {
	// 主节点
	var masterDsn = config.MasterNode.getDsn()

	var err error

	newLogger := l.GormLogger{
		LogLevel:      logger.Info,
		SlowThreshold: 100 * time.Millisecond,
	}
	DB, err = gorm.Open(postgres.Open(masterDsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	// 从节点
	var slaves = make([]gorm.Dialector, 0)
	for _, node := range config.SlaveNodes {
		slaves = append(slaves, postgres.Open(node.getDsn()))
	}
	err = DB.Use(dbresolver.Register(dbresolver.Config{Replicas: slaves}))
	if err != nil {
		panic(err)
	}
	// 连接池
	setConnectionPool()

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
	sqlDB.SetConnMaxLifetime(config.CONF_INSTANCE.DbConf.MaxLifeTime)
}
