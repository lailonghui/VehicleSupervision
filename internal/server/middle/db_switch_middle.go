package middle

import (
	"context"
	"github.com/gin-gonic/gin"
)

// 数据库主从开关 上下文key
const DB_SWITCH_CONTEXT_KEY = "C_DB_SWITCH"

// 数据库主从开关 头部
const DB_SWITCH_HEADER = "db_switch"

// 数据库主从开关中间件-读请求默认走从库
// 传true时候才走主库
func DBSwitchMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerValue := c.Request.Header.Get(DB_SWITCH_HEADER)
		if headerValue == "true" {
			ctx := context.WithValue(c.Request.Context(), DB_SWITCH_CONTEXT_KEY, true)
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}

// 是否强制走主库
func IsForceMaster(ctx context.Context) bool {
	v := ctx.Value(DB_SWITCH_CONTEXT_KEY)
	if v == nil {
		return false
	}
	value, ok := v.(bool)
	if ok {
		return value
	}
	return false
}
