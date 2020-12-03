package server

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/playground"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"log"
	"system-manage/config"
	"system-manage/internal/graphql/system"
	"system-manage/pkg/logger"
	"time"
)

// 监听的host
var host string

// 监听的port
var port int

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func Setup() {
	host = config.CONF_INSTANCE.ServerConf.Host
	port = config.CONF_INSTANCE.ServerConf.Port

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// gin中间件配置
	router.Use(ginzap.Ginzap(logger.GinLogger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger.GinLogger, true))
	// 路由配置
	router.GET("/", playgroundHandler())
	router.POST("/graphql", system.GinEndpoint())

	log.Fatal(router.Run(fmt.Sprintf("%s:%d", host, port)))
}
