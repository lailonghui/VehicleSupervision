package server

import (
	"VehicleSupervision/config"
	adasMutation "VehicleSupervision/internal/modules/adas/mutation"
	adasQuery "VehicleSupervision/internal/modules/adas/query"
	departmentMutation "VehicleSupervision/internal/modules/admin/department/mutation"
	departmentQuery "VehicleSupervision/internal/modules/admin/department/query"
	enterpriseMutation "VehicleSupervision/internal/modules/admin/enterprise/mutation"
	enterpriseQuery "VehicleSupervision/internal/modules/admin/enterprise/query"
	systemUserMutation "VehicleSupervision/internal/modules/admin/systemuser/mutation"
	systemUserQuery "VehicleSupervision/internal/modules/admin/systemuser/query"
	areaMutation "VehicleSupervision/internal/modules/area/mutation"
	areaQuery "VehicleSupervision/internal/modules/area/query"
	blacklistRecordMutation "VehicleSupervision/internal/modules/blacklist/record/mutation"
	blacklistRecordQuery "VehicleSupervision/internal/modules/blacklist/record/query"
	dictionaryCategoryMutation "VehicleSupervision/internal/modules/dictionary/category/mutation"
	dictionaryCategoryQuery "VehicleSupervision/internal/modules/dictionary/category/query"
	dictionaryMutation "VehicleSupervision/internal/modules/dictionary/dict/mutation"
	dictionaryQuery "VehicleSupervision/internal/modules/dictionary/dict/query"
	trainingMutation "VehicleSupervision/internal/modules/training/mutation"
	trainingQuery "VehicleSupervision/internal/modules/training/query"
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
	// 企业端点
	router.Any("/enterprise/query", enterpriseQuery.GinEndpoint())
	router.Any("/enterprise/mutation", enterpriseMutation.GinEndpoint())
	// 系统用户端点
	router.Any("/system_user/query", systemUserQuery.GinEndpoint())
	router.Any("/system_user/mutation", systemUserMutation.GinEndpoint())
	// 部门端点
	router.Any("/department/query", departmentQuery.GinEndpoint())
	router.Any("/department/mutation", departmentMutation.GinEndpoint())
	// 字典类型端点
	router.Any("/dictionary_category/query", dictionaryCategoryQuery.GinEndpoint())
	router.Any("/dictionary_category/mutation", dictionaryCategoryMutation.GinEndpoint())
	// 字典端点
	router.Any("/dictionary/query", dictionaryQuery.GinEndpoint())
	router.Any("/dictionary/mutation", dictionaryMutation.GinEndpoint())
	// 黑名单记录端点
	router.Any("/blacklist_operation_record/query", blacklistRecordQuery.GinEndpoint())
	router.Any("/blacklist_operation_record/mutation", blacklistRecordMutation.GinEndpoint())

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
