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

func (r *mutationResolver) DeleteControlLineTime(ctx context.Context, where model.ControlLineTimeBoolExp) (*model.ControlLineTimeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ControlLineTime{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ControlLineTime
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
	return &model.ControlLineTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteControlLineTimeByPk(ctx context.Context, Id int64) (*model1.ControlLineTime, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.ControlLineTime
	tx := db.DB.Model(&model1.ControlLineTime{})
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

func (r *mutationResolver) InsertControlLineTime(ctx context.Context, objects []*model.ControlLineTimeInsertInput) (*model.ControlLineTimeMutationResponse, error) {
	rs := make([]*model1.ControlLineTime, 0)
	for _, object := range objects {
		v := &model1.ControlLineTime{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.ControlLineTime{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.ControlLineTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertControlLineTimeOne(ctx context.Context, object model.ControlLineTimeInsertInput) (*model1.ControlLineTime, error) {
	rs := &model1.ControlLineTime{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.ControlLineTime{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateControlLineTime(ctx context.Context, inc *model.ControlLineTimeIncInput, set *model.ControlLineTimeSetInput, where model.ControlLineTimeBoolExp) (*model.ControlLineTimeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ControlLineTime{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.ControlLineTimeMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.ControlLineTime
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
	return &model.ControlLineTimeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateControlLineTimeByPk(ctx context.Context, inc *model.ControlLineTimeIncInput, set *model.ControlLineTimeSetInput, Id int64) (*model1.ControlLineTime, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.ControlLineTime{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.ControlLineTime
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) ControlLineTime(ctx context.Context, distinctOn []model.ControlLineTimeSelectColumn, limit *int, offset *int, orderBy []*model.ControlLineTimeOrderBy, where *model.ControlLineTimeBoolExp) ([]*model1.ControlLineTime, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.ControlLineTime{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.ControlLineTime
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) ControlLineTimeAggregate(ctx context.Context, distinctOn []model.ControlLineTimeSelectColumn, limit *int, offset *int, orderBy []*model.ControlLineTimeOrderBy, where *model.ControlLineTimeBoolExp) (*model.ControlLineTimeAggregate, error) {
	var rs model.ControlLineTimeAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.ControlLineTime{})
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

func (r *queryResolver) ControlLineTimeByPk(ctx context.Context, Id int64) (*model1.ControlLineTime, error) {
	var rs model1.ControlLineTime
	tx := db.DB.Model(&model1.ControlLineTime{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
