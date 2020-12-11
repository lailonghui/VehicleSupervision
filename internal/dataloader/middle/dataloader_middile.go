package middle

import (
	adminModel "VehicleSupervision/internal/modules/admin/model"
	deviceModel "VehicleSupervision/internal/modules/device/model"
	"context"
	"github.com/gin-gonic/gin"
)

const DATA_LOADER_CONTEXT_KEY = "C_DATALOADER"

// dataloader集合
type Loaders struct {
	*adminModel.EnterpriseLoader
	*adminModel.DepartmentLoader
	*adminModel.SystemUserLoader
	*deviceModel.SimCardFlowLoader
	*deviceModel.SimCardLoader
	*deviceModel.TerminalLoader
	*deviceModel.TerminalFactoryLoader
	*deviceModel.TerminalTypeLoader
	*deviceModel.TerminalModalLoader
}

// dataloader 中间件
func DataloaderMiddle(contextKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextKey, &Loaders{
			EnterpriseLoader:      (&adminModel.Enterprise{}).NewLoader(),
			DepartmentLoader:      (&adminModel.Department{}).NewLoader(),
			SystemUserLoader:      (&adminModel.SystemUser{}).NewLoader(),
			SimCardFlowLoader:     (&deviceModel.SimCardFlow{}).NewLoader(),
			SimCardLoader:         (&deviceModel.SimCard{}).NewLoader(),
			TerminalLoader:        (&deviceModel.Terminal{}).NewLoader(),
			TerminalFactoryLoader: (&deviceModel.TerminalFactory{}).NewLoader(),
			TerminalTypeLoader:    (&deviceModel.TerminalType{}).NewLoader(),
			TerminalModalLoader:   (&deviceModel.TerminalModal{}).NewLoader(),
		})
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}

}
