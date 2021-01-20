package resolver

import (
	"VehicleSupervision/internal/cache"
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle/model"
	"VehicleSupervision/internal/server/middle"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteTempIdentityCardDownloadLog(ctx context.Context, where model.TempIdentityCardDownloadLogBoolExp) (*model.TempIdentityCardDownloadLogMutationResponse, error) {
	var rs []*model1.TempIdentityCardDownloadLog
	var m = &model1.TempIdentityCardDownloadLog{}
	// 查询主键和联合主键
	qt := util.NewQueryTranslator(db.DB, m)
	tx := qt.Where(where).Finish()

	tx = tx.Select(m.PrimaryColumnName())

	tx = tx.Find(&rs)
	// 删除
	amount := len(rs)
	if amount == 0 {
		return &model.TempIdentityCardDownloadLogMutationResponse{
			AffectedRows: 0,
			Returning:    nil,
		}, nil
	}
	var idList = make([]int64, 0, amount)
	for i := 0; i < amount; i++ {
		idList[i] = rs[i].GetPrimary()
	}
	tx = db.DB.Model(m).Delete(idList)
	if err := tx.Error; err != nil {
		return nil, err
	}
	// 删除缓存
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		rsLen := len(rs)
		var ids = make([]string, 0, rsLen)

		for i := 0; i < rsLen; i++ {
			ids[i] = util2.ToStr(rs[i].GetPrimary())

		}

		_ = cacheAspect.OnListRemove(ctx, ids, nil)

	}
	return &model.TempIdentityCardDownloadLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    nil,
	}, nil
}

func (r *mutationResolver) DeleteTempIdentityCardDownloadLogByPk(ctx context.Context, id int64) (*model1.TempIdentityCardDownloadLog, error) {
	var rs model1.TempIdentityCardDownloadLog
	m := &model1.TempIdentityCardDownloadLog{}
	tx := db.DB.Model(m)
	// 查询记录
	tx = tx.Where(rs.PrimaryColumnName()+" = ?", id).First(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	// 删除缓存
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnPkRemove(ctx, util2.ToStr(rs.GetPrimary()))

	}
	return &rs, nil
}

func (r *mutationResolver) InsertTempIdentityCardDownloadLog(ctx context.Context, objects []*model.TempIdentityCardDownloadLogInsertInput) (*model.TempIdentityCardDownloadLogMutationResponse, error) {
	rs := make([]*model1.TempIdentityCardDownloadLog, 0)
	m := &model1.TempIdentityCardDownloadLog{}
	// 结构转换
	for _, object := range objects {
		v := &model1.TempIdentityCardDownloadLog{}
		util2.StructAssign(object, v)
		rs = append(rs, v)
	}
	// 插入记录
	tx := db.DB.Model(m).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	// 清理缓存
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnInsert(ctx)
	}
	return &model.TempIdentityCardDownloadLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTempIdentityCardDownloadLogOne(ctx context.Context, objects model.TempIdentityCardDownloadLogInsertInput) (*model1.TempIdentityCardDownloadLog, error) {
	rs := &model1.TempIdentityCardDownloadLog{}
	m := &model1.TempIdentityCardDownloadLog{}
	// 插入记录
	util2.StructAssign(rs, &objects)
	tx := db.DB.Model(m).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	// 清理缓存
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnInsert(ctx)
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTempIdentityCardDownloadLog(ctx context.Context, inc *model.TempIdentityCardDownloadLogIncInput, set *model.TempIdentityCardDownloadLogSetInput, where model.TempIdentityCardDownloadLogBoolExp) (*model.TempIdentityCardDownloadLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TempIdentityCardDownloadLog{})
	var m = &model1.TempIdentityCardDownloadLog{}
	// 更新数据库
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TempIdentityCardDownloadLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 查询数据库
	var rs []*model1.TempIdentityCardDownloadLog
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	// 清理缓存
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		rsLen := len(rs)
		var ids = make([]string, 0, rsLen)

		for i := 0; i < rsLen; i++ {
			ids[i] = util2.ToStr(rs[i].GetPrimary())

		}

		_ = cacheAspect.OnListUpdate(ctx, ids, nil)

	}
	return &model.TempIdentityCardDownloadLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTempIdentityCardDownloadLogByPk(ctx context.Context, inc *model.TempIdentityCardDownloadLogIncInput, set *model.TempIdentityCardDownloadLogSetInput, id int64) (*model1.TempIdentityCardDownloadLog, error) {
	var rs model1.TempIdentityCardDownloadLog
	var m = &model1.TempIdentityCardDownloadLog{}
	// 更新数据库
	tx := db.DB.Where(rs.PrimaryColumnName()+" = ?", id)
	qt := util.NewQueryTranslator(tx, &model1.TempIdentityCardDownloadLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	// 查询数据库
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	// 清理缓存
	cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
	if cacheErr == nil {
		_ = cacheAspect.OnPkRemove(ctx, util2.ToStr(rs.GetPrimary()))

	}
	return &rs, nil
}

func (r *queryResolver) TempIdentityCardDownloadLog(ctx context.Context, distinctOn []model.TempIdentityCardDownloadLogSelectColumn, limit *int, offset *int, orderBy []*model.TempIdentityCardDownloadLogOrderBy, where *model.TempIdentityCardDownloadLogBoolExp) ([]*model1.TempIdentityCardDownloadLog, error) {
	var rs []*model1.TempIdentityCardDownloadLog
	var m = &model1.TempIdentityCardDownloadLog{}
	if middle.IsEnableGqlCache(ctx) {
		// 如果启用缓存，则从缓存中查询数据
		cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
		cacheKey := util2.ToStr(map[string]interface{}{
			"distinctOn": distinctOn,
			"limit":      limit,
			"offset":     offset,
			"orderBy":    orderBy,
			"where":      where,
		})
		cacheKey = fmt.Sprintf("%x", sha256.Sum256([]byte(cacheKey)))
		if cacheErr == nil {
			exist, err := cacheAspect.OnListQuery(ctx, cacheKey, &rs)
			if err != nil {
				return nil, err
			}
			if exist {
				return rs, err
			}
		}
		// 缓存中找不到数据的话，查询数据库获取数据
		qt := util.NewQueryTranslator(db.DB, m)
		tx := qt.DistinctOn(distinctOn).
			Limit(limit).
			Offset(offset).
			OrderBy(orderBy).
			Where(where).
			Finish()
		tx = tx.Find(&rs)
		err := tx.Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return rs, nil
			}
			return rs, err
		}
		// 设置数据到缓存
		if cacheErr == nil {
			_ = cacheAspect.SetListQueryCache(ctx, cacheKey, rs)
		}
		return rs, err
	}
	// 如果没启用缓存，则直接从数据库中查询
	qt := util.NewQueryTranslator(db.DB, m)
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TempIdentityCardDownloadLogAggregate(ctx context.Context, distinctOn []model.TempIdentityCardDownloadLogSelectColumn, limit *int, offset *int, orderBy []*model.TempIdentityCardDownloadLogOrderBy, where *model.TempIdentityCardDownloadLogBoolExp) (*model.TempIdentityCardDownloadLogAggregate, error) {
	var rs model.TempIdentityCardDownloadLogAggregate
	var m = &model1.TempIdentityCardDownloadLog{}
	// 获取聚合查询项
	queryStrings := util.GetPreloadsMustPrefix(ctx, "aggregate.")
	if middle.IsEnableGqlCache(ctx) {
		// 如果启用缓存，则从缓存中查询数据
		cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
		cacheKey := util2.ToStr(map[string]interface{}{
			"distinctOn":   distinctOn,
			"limit":        limit,
			"offset":       offset,
			"orderBy":      orderBy,
			"where":        where,
			"queryStrings": queryStrings,
		})
		cacheKey = fmt.Sprintf("%x", sha256.Sum256([]byte(cacheKey)))
		if cacheErr == nil {
			exist, err := cacheAspect.OnArrgegateQuery(ctx, cacheKey, &rs)
			if err != nil {
				return nil, err
			}
			if exist {
				return &rs, err
			}
		}
		// 缓存中找不到数据的话，查询数据库获取数据
		qt := util.NewQueryTranslator(db.DB, &model1.TempIdentityCardDownloadLog{})
		tx, err := qt.DistinctOn(distinctOn).
			Limit(limit).
			Offset(offset).
			OrderBy(orderBy).
			Where(where).
			AggregateWithQueryString(&rs, queryStrings)
		if err != nil {
			return nil, err
		}
		// 设置数据到缓存
		if cacheErr == nil {
			_ = cacheAspect.SetArrgegateQueryCache(ctx, cacheKey, rs)
		}
		err = tx.Error
		return &rs, err
	}
	// 如果没启用缓存，则直接从数据库中查询
	qt := util.NewQueryTranslator(db.DB, &model1.TempIdentityCardDownloadLog{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		AggregateWithQueryString(&rs, queryStrings)
	if err != nil {
		return nil, err
	}
	err = tx.Error
	return &rs, err
}

func (r *queryResolver) TempIdentityCardDownloadLogByPk(ctx context.Context, id int64) (*model1.TempIdentityCardDownloadLog, error) {

	var m = &model1.TempIdentityCardDownloadLog{}
	var rs model1.TempIdentityCardDownloadLog
	if middle.IsEnableGqlCache(ctx) {
		// 如果启用缓存，则从缓存中查询数据
		cacheAspect, cacheErr := cache.GetGqlCacheAspect(m.TableName())
		cacheKey := util2.ToStr(id)
		if cacheErr == nil {
			exist, err := cacheAspect.OnPkQuery(ctx, cacheKey, &rs)
			if err != nil {
				return nil, err
			}
			if exist {

				if rs.GetPrimary() == 0 {
					return nil, nil
				}

				return &rs, nil
			}
		}
		// 缓存中找不到数据的话，查询数据库获取数据
		tx := db.DB.Model(m).First(&rs, id)
		err := tx.Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				if cacheErr == nil {
					_ = cacheAspect.SetNotExistPkQueryCache(ctx, cacheKey, "")
				}
				return nil, nil
			}
			return &rs, err
		}
		// 设置数据到缓存
		if cacheErr == nil {
			_ = cacheAspect.SetPkQueryCache(ctx, cacheKey, rs)
		}
		return &rs, nil
	}
	// 如果没启用缓存，则直接从数据库中查询
	tx := db.DB.Model(m).First(&rs, id)
	err := tx.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
