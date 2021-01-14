package model

import (
	"VehicleSupervision/internal/cache"
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/server/middle"
	"VehicleSupervision/pkg/logger"
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

//go:generate go run github.com/vektah/dataloaden DynamicSpotCheckDisposalPkLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.DynamicSpotCheckDisposal

// 数据库表名
func (t *DynamicSpotCheckDisposal) TableName() string {
	return "dynamic_spot_check_disposal"
}

// 主键列名
func (t *DynamicSpotCheckDisposal) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *DynamicSpotCheckDisposal) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *DynamicSpotCheckDisposalPkLoader) NewLoader(ctx context.Context) *DynamicSpotCheckDisposalPkLoader {
	return &DynamicSpotCheckDisposalPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DynamicSpotCheckDisposal, []error) {
			var rs []*DynamicSpotCheckDisposal = make([]*DynamicSpotCheckDisposal, len(keys))
			var m DynamicSpotCheckDisposal
			wg := sync.WaitGroup{}
			if middle.IsEnableGqlCache(ctx) {
				// 如果启用缓存，则从缓存中查询数据
				cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
				for i := range keys {
					wg.Add(1)
					go func(i int) {

						defer wg.Done()
						logger.Info("dataloader获取值", zap.Int("i", i))
						// 如果启用缓存，则从缓存中查询数据
						var entity DynamicSpotCheckDisposal
						cacheKey := keys[i]
						if cacheErr == nil {
							exist, err := cacheAspect.OnPkQuery(ctx, cacheKey, &entity)
							logger.Info("dataloader获取缓存值", zap.Error(err), zap.Bool("exist", exist), zap.Any("entity", entity))
							if err != nil {
								return
							}
							if exist {
								if entity.GetPrimary() != 0 {
									rs[i] = &entity
								}
								return
							}
						}
						// 缓存中找不到数据的话，查询数据库获取数据
						tx := db.DB.Model(m).Where(m.PrimaryColumnName()+" = ?", keys[i]).First(&entity)
						err := tx.Error
						if err != nil {
							if err == gorm.ErrRecordNotFound {
								if cacheErr == nil {
									_ = cacheAspect.SetNotExistPkQueryCache(ctx, cacheKey, "")
								}
								return
							}
							return
						}
						// 设置数据到缓存
						if cacheErr == nil {
							_ = cacheAspect.SetPkQueryCache(ctx, cacheKey, entity)
						}
						rs[i] = &entity
					}(i)
				}

			}
			wg.Wait()
			return rs, nil
		},
	}
}
