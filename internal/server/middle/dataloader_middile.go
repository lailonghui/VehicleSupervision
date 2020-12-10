package middle

import (
	"VehicleSupervision/internal/modules/admin/model"
	"context"
	"github.com/gin-gonic/gin"
)

const DATA_LOADER_CONTEXT_KEY = "C_DATALOADER"

// dataloader集合
type Loaders struct {
	EnterpriseById *model.EnterpriseLoader
	DepartmentById *model.DepartmentLoader
	SystemUserById *model.SystemUserLoader
}

func DataloaderMiddle(contextKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextKey, &Loaders{
			EnterpriseById: (&model.Enterprise{}).NewLoader(),
			DepartmentById: (&model.Department{}).NewLoader(),
			SystemUserById: (&model.SystemUser{}).NewLoader(),
		})
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}

}
