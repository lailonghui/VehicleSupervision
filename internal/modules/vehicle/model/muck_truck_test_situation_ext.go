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

//go:generate go run github.com/vektah/dataloaden MuckTruckTestSituationUnionPkLoader string *VehicleSupervision/internal/modules/vehicle/model.MuckTruckTestSituation

// 数据库表名
func (t *MuckTruckTestSituation) TableName() string {
	return "muck_truck_test_situation"
}

// 主键列名
func (t *MuckTruckTestSituation) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *MuckTruckTestSituation) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *MuckTruckTestSituationPkLoader) NewLoader(ctx context.Context) *MuckTruckTestSituationPkLoader {
	return &MuckTruckTestSituationPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckTestSituation, []error) {
			var rs []*MuckTruckTestSituation = make([]*MuckTruckTestSituation, len(keys))
			var m MuckTruckTestSituation
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
						var entity MuckTruckTestSituation
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

// 联合主键列名
func (t *MuckTruckTestSituation) UnionPrimaryColumnName() string {
	return "muck_truck_test_situation_id"
}

// 获取联合主键
func (t *MuckTruckTestSituation) GetUnionPrimary() string {
	return t.MuckTruckTestSituationID
}

// 新建联合主键dataloader
func (t *MuckTruckTestSituationUnionPkLoader) NewLoader(ctx context.Context) *MuckTruckTestSituationUnionPkLoader {
	return &MuckTruckTestSituationUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckTestSituation, []error) {
			var rs []*MuckTruckTestSituation = make([]*MuckTruckTestSituation, len(keys))
			var m MuckTruckTestSituation
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
						var entity MuckTruckTestSituation
						cacheKey := keys[i]
						if cacheErr == nil {
							exist, err := cacheAspect.OnUnionPkQuery(ctx, cacheKey, &entity)
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
						tx := db.DB.Model(m).Where(m.UnionPrimaryColumnName()+" = ?", keys[i]).First(&entity)
						err := tx.Error
						if err != nil {
							if err == gorm.ErrRecordNotFound {
								if cacheErr == nil {
									_ = cacheAspect.SetNotExistUnionPkQueryCache(ctx, cacheKey, "")
								}
								return
							}
							return
						}
						// 设置数据到缓存
						if cacheErr == nil {
							_ = cacheAspect.SetUnionPkQueryCache(ctx, cacheKey, entity)
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
