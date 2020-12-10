package server

import (
	"VehicleSupervision/config"
	adasMutation "VehicleSupervision/internal/modules/adas/mutation"
	adasQuery "VehicleSupervision/internal/modules/adas/query"

	admin "VehicleSupervision/internal/modules/admin"
	areaMutation "VehicleSupervision/internal/modules/area/mutation"
	areaQuery "VehicleSupervision/internal/modules/area/query"
	blacklistRecord "VehicleSupervision/internal/modules/blacklist"
	device "VehicleSupervision/internal/modules/device"
	ridehailing "VehicleSupervision/internal/modules/ridehailing"

	dictionary "VehicleSupervision/internal/modules/dictionary"

	"VehicleSupervision/internal/modules/driver"
	driving "VehicleSupervision/internal/modules/driving"
	trainingMutation "VehicleSupervision/internal/modules/training/mutation"
	trainingQuery "VehicleSupervision/internal/modules/training/query"
	"VehicleSupervision/internal/modules/vehicle"
	vehicleLocationHisMutation "VehicleSupervision/internal/modules/vehiclelocation/his/mutation"
	vehicleLocationHisQuery "VehicleSupervision/internal/modules/vehiclelocation/his/query"
	vehicleLocationLastMutation "VehicleSupervision/internal/modules/vehiclelocation/last/mutation"
	vehicleLocationLastQuery "VehicleSupervision/internal/modules/vehiclelocation/last/query"
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
	router.Any("/vehicle", vehicle.GinEndpoint())
	//驾驶员模块端点
	router.Any("/driver", driver.GinEndpoint())

	// 系统管理端点
	router.Any("/admin", admin.GinEndpoint())
	// 字典管理端点
	router.Any("/dictionary", dictionary.GinEndpoint())
	// 黑名单端点
	router.Any("/blacklist", blacklistRecord.GinEndpoint())
	// 行车端点
	router.Any("/driving", driving.GinEndpoint())
	// 车辆历史位置端点
	router.Any("/vehicle_location_his/query", vehicleLocationHisQuery.GinEndpoint())
	router.Any("/vehicle_location_his/mutation", vehicleLocationHisMutation.GinEndpoint())
	// 车辆最新位置端点
	router.Any("/vehicle_location_last/query", vehicleLocationLastQuery.GinEndpoint())
	router.Any("/vehicle_location_last/mutation", vehicleLocationLastMutation.GinEndpoint())
	// 设备管理端点
	router.Any("/device", device.GinEndpoint())
	// 网约车模块
	router.Any("/ride_hailing", ridehailing.GinEndpoint())

	//adas模块端点
	router.Any("/adas/query", adasQuery.GinEndpoint())
	router.Any("/adas/mutation", adasMutation.GinEndpoint())

	//地区模块端点
	router.Any("/area/query", areaQuery.GinEndpoint())
	router.Any("/area/mutation", areaMutation.GinEndpoint())

	//教育模块端点
	router.Any("/training/query", trainingQuery.GinEndpoint())
	router.Any("/training/mutation", trainingMutation.GinEndpoint())

	addr := fmt.Sprintf("%s:%d", host, port)
	logger.Info("server will run on " + addr)
	logger.Info(fmt.Sprintf("connect to http://localhost:%d/ for GraphQL playground", port))
	log.Fatal(router.Run(addr))
}
