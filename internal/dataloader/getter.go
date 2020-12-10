package dataloader

import (
	"VehicleSupervision/internal/dataloader/middle"
	"context"
)

// 获取dataloader实例集合
func GetLoaders(ctx context.Context) *middle.Loaders {
	return ctx.Value(middle.DATA_LOADER_CONTEXT_KEY).(*middle.Loaders)
}
