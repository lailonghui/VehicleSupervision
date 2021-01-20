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

//go:generate go run github.com/vektah/dataloaden VioCodewfdmPkLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.VioCodewfdm

// 数据库表名
func (t *VioCodewfdm) TableName() string {
	return "vio_codewfdm"
}

// 主键列名
func (t *VioCodewfdm) PrimaryColumnName() string {
	return "xh"
}

// 获取主键
func (t *VioCodewfdm) GetPrimary() string {
	return t.Xh
}

// 新建主键dataloader
func (t *VioCodewfdmPkLoader) NewLoader(ctx context.Context) *VioCodewfdmPkLoader {
	return &VioCodewfdmPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VioCodewfdm, []error) {
			var rs []*VioCodewfdm = make([]*VioCodewfdm, len(keys))
			var m VioCodewfdm
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
						var entity VioCodewfdm
						cacheKey := keys[i]
						if cacheErr == nil {
							exist, err := cacheAspect.OnPkQuery(ctx, cacheKey, &entity)
							logger.Info("dataloader获取缓存值", zap.Error(err), zap.Bool("exist", exist), zap.Any("entity", entity))
							if err != nil {
								return
							}
							if exist {

								if entity.GetPrimary() != "" {
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
