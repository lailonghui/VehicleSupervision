package server

import (
	"VehicleSupervision/config"
	enterprise_mutation "VehicleSupervision/internal/modules/admin/enterprise/mutation"
	enterprise_query "VehicleSupervision/internal/modules/admin/enterprise/query"
	"VehicleSupervision/pkg/logger"
	"fmt"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"log"
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

	//车辆模块端点
	//router.Any("/vehicle", vehicle.GinEndpoint())
	//驾驶员模块端点
	//router.Any("/driver", driver.GinEndpoint())
	router.Any("/enterprise/query", enterprise_query.GinEndpoint())
	router.Any("/enterprise/mutation", enterprise_mutation.GinEndpoint())
	addr := fmt.Sprintf("%s:%d", host, port)
	logger.Info("server will run on " + addr)
	logger.Info(fmt.Sprintf("connect to http://localhost:%d/ for GraphQL playground", port))
	log.Fatal(router.Run(addr))
}
