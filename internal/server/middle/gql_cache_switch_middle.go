package middle

import (
	"context"
	"github.com/gin-gonic/gin"
)

// qqlCache开关 上下文key
const GQL_CACHE_SWITCH_CONTEXT_KEY = "C_GQL_CACHE_SWITCH"

// gqlCache开关 头部
const GQL_CACHE_SWITCH_HEADER = "gql_cache_switch"

// gqlcache开关中间件-默认打开
// 传false时候才关闭gqlcache
func GqlCacheSwitchMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerValue := c.Request.Header.Get(GQL_CACHE_SWITCH_HEADER)
		if headerValue == "false" {
			ctx := context.WithValue(c.Request.Context(), GQL_CACHE_SWITCH_CONTEXT_KEY, false)
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}

// 是否启用gql缓存
func IsEnableGqlCache(ctx context.Context) bool {
	v := ctx.Value(GQL_CACHE_SWITCH_CONTEXT_KEY)
	if v == nil {
		return true
	}
	value, ok := v.(bool)
	if ok {
		return value
	}
	return true
}
