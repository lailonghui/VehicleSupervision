package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driving/graph/model"
	model1 "VehicleSupervision/internal/modules/driving/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteLimitSpeedLineTime(ctx context.Context, where model.LimitSpeedLineTimeBoolExp) (*model.LimitSpeedLineTimeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedLineTime{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.LimitSpeedLineTime
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &model.LimitSpeedLineTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteLimitSpeedLineTimeByPk(ctx context.Context, Id int64) (*model1.LimitSpeedLineTime, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.LimitSpeedLineTime
	tx := db.DB.Model(&model1.LimitSpeedLineTime{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ?", Id).First(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *mutationResolver) DeleteLimitSpeedLineTimeByUnionPk(ctx context.Context, unionId string) (*model1.LimitSpeedLineTime, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.LimitSpeedLineTime
	tx := db.DB.Model(&model1.LimitSpeedLineTime{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *mutationResolver) InsertLimitSpeedLineTime(ctx context.Context, objects []*model.LimitSpeedLineTimeInsertInput) (*model.LimitSpeedLineTimeMutationResponse, error) {
	rs := make([]*model1.LimitSpeedLineTime, 0)
	for _, object := range objects {
		v := &model1.LimitSpeedLineTime{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.LimitSpeedLineTime{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.LimitSpeedLineTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertLimitSpeedLineTimeOne(ctx context.Context, object model.LimitSpeedLineTimeInsertInput) (*model1.LimitSpeedLineTime, error) {
	rs := &model1.LimitSpeedLineTime{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.LimitSpeedLineTime{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateLimitSpeedLineTime(ctx context.Context, inc *model.LimitSpeedLineTimeIncInput, set *model.LimitSpeedLineTimeSetInput, where model.LimitSpeedLineTimeBoolExp) (*model.LimitSpeedLineTimeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedLineTime{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.LimitSpeedLineTimeMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.LimitSpeedLineTime
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
	}
	return &model.LimitSpeedLineTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateLimitSpeedLineTimeByPk(ctx context.Context, inc *model.LimitSpeedLineTimeIncInput, set *model.LimitSpeedLineTimeSetInput, Id int64) (*model1.LimitSpeedLineTime, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.LimitSpeedLineTime{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.LimitSpeedLineTime
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateLimitSpeedLineTimeByUnionPk(ctx context.Context, inc *model.LimitSpeedLineTimeIncInput, set *model.LimitSpeedLineTimeSetInput, unionId string) (*model1.LimitSpeedLineTime, error) {
	var rs model1.LimitSpeedLineTime
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.LimitSpeedLineTime{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}

	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) LimitSpeedLineTime(ctx context.Context, distinctOn []model.LimitSpeedLineTimeSelectColumn, limit *int, offset *int, orderBy []*model.LimitSpeedLineTimeOrderBy, where *model.LimitSpeedLineTimeBoolExp) ([]*model1.LimitSpeedLineTime, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedLineTime{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.LimitSpeedLineTime
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) LimitSpeedLineTimeAggregate(ctx context.Context, distinctOn []model.LimitSpeedLineTimeSelectColumn, limit *int, offset *int, orderBy []*model.LimitSpeedLineTimeOrderBy, where *model.LimitSpeedLineTimeBoolExp) (*model.LimitSpeedLineTimeAggregate, error) {
	var rs model.LimitSpeedLineTimeAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedLineTime{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		return nil, err
	}
	err = tx.Error
	return &rs, err
}

func (r *queryResolver) LimitSpeedLineTimeByPk(ctx context.Context, Id int64) (*model1.LimitSpeedLineTime, error) {
	var rs model1.LimitSpeedLineTime
	tx := db.DB.Model(&model1.LimitSpeedLineTime{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) LimitSpeedLineTimeByUnionPk(ctx context.Context, unionId string) (*model1.LimitSpeedLineTime, error) {
	var rs model1.LimitSpeedLineTime
	tx := db.DB.Model(&model1.LimitSpeedLineTime{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
