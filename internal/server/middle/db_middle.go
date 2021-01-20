package middle

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/pkg/logger"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sync"
)

// 数据库对象 上下文key
const DB_CONTEXT_KEY = "C_DB"

//事务开关 头部
const TX_SWITCH_HEADER = "tx_switch"

// 上下文中的数据库信息
type ContextDB struct {
	// 数据库对象
	tx *gorm.DB
	// 是否启用事务
	txEnable bool
	// 互斥锁
	mutex sync.Mutex
}

func DbMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		txSwitchHeader := c.Request.Header.Get(TX_SWITCH_HEADER)
		cdb := &ContextDB{
			tx:       nil,
			txEnable: false,
			mutex:    sync.Mutex{},
		}
		if txSwitchHeader == "false" {
			cdb.txEnable = false

		} else {
			cdb.txEnable = true
		}
		ctx := context.WithValue(c.Request.Context(), DB_CONTEXT_KEY, cdb)
		c.Request = c.Request.WithContext(ctx)
		c.Next()


	}
}

// 返回数据库对象
func GetTx(ctx context.Context) *gorm.DB {
	contextDB := ctx.Value(DB_CONTEXT_KEY).(*ContextDB)
	if contextDB == nil {
		logger.Error("获取contextDB失败")
		return db.DB
	}
	if contextDB.tx != nil {
		return contextDB.tx
	}
	contextDB.mutex.Lock()
	defer contextDB.mutex.Unlock()
	if contextDB.txEnable {
		contextDB.tx = db.DB
	}
	return db.DB
}
