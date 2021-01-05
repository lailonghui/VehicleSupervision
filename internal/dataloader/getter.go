package dataloader

import (
	"context"
)

// 获取dataloader实例集合
func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(DATA_LOADER_CONTEXT_KEY).(*Loaders)
}
