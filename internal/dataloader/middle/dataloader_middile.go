package middle

import (
	"VehicleSupervision/internal/modules/admin/model"
	"context"
	"github.com/gin-gonic/gin"
)

const DATA_LOADER_CONTEXT_KEY = "C_DATALOADER"

// dataloader集合
type Loaders struct {
	*model.EnterpriseLoader
	*model.DepartmentLoader
	*model.SystemUserLoader
}

// dataloader 中间件
func DataloaderMiddle(contextKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextKey, &Loaders{
			EnterpriseLoader: (&model.Enterprise{}).NewLoader(),
			DepartmentLoader: (&model.Department{}).NewLoader(),
			SystemUserLoader: (&model.SystemUser{}).NewLoader(),
		})
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}

}
