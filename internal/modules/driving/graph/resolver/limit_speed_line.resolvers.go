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

func (r *mutationResolver) DeleteLimitSpeedLine(ctx context.Context, where model.LimitSpeedLineBoolExp) (*model.LimitSpeedLineMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedLine{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.LimitSpeedLine
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
	return &model.LimitSpeedLineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteLimitSpeedLineByPk(ctx context.Context, Id int64) (*model1.LimitSpeedLine, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.LimitSpeedLine
	tx := db.DB.Model(&model1.LimitSpeedLine{})
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

func (r *mutationResolver) InsertLimitSpeedLine(ctx context.Context, objects []*model.LimitSpeedLineInsertInput) (*model.LimitSpeedLineMutationResponse, error) {
	rs := make([]*model1.LimitSpeedLine, 0)
	for _, object := range objects {
		v := &model1.LimitSpeedLine{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.LimitSpeedLine{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.LimitSpeedLineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertLimitSpeedLineOne(ctx context.Context, object model.LimitSpeedLineInsertInput) (*model1.LimitSpeedLine, error) {
	rs := &model1.LimitSpeedLine{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.LimitSpeedLine{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateLimitSpeedLine(ctx context.Context, inc *model.LimitSpeedLineIncInput, set *model.LimitSpeedLineSetInput, where model.LimitSpeedLineBoolExp) (*model.LimitSpeedLineMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedLine{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.LimitSpeedLineMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.LimitSpeedLine
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
	return &model.LimitSpeedLineMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateLimitSpeedLineByPk(ctx context.Context, inc *model.LimitSpeedLineIncInput, set *model.LimitSpeedLineSetInput, Id int64) (*model1.LimitSpeedLine, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.LimitSpeedLine{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.LimitSpeedLine
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) LimitSpeedLine(ctx context.Context, distinctOn []model.LimitSpeedLineSelectColumn, limit *int, offset *int, orderBy []*model.LimitSpeedLineOrderBy, where *model.LimitSpeedLineBoolExp) ([]*model1.LimitSpeedLine, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedLine{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.LimitSpeedLine
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) LimitSpeedLineAggregate(ctx context.Context, distinctOn []model.LimitSpeedLineSelectColumn, limit *int, offset *int, orderBy []*model.LimitSpeedLineOrderBy, where *model.LimitSpeedLineBoolExp) (*model.LimitSpeedLineAggregate, error) {
	var rs model.LimitSpeedLineAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.LimitSpeedLine{})
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

func (r *queryResolver) LimitSpeedLineByPk(ctx context.Context, Id int64) (*model1.LimitSpeedLine, error) {
	var rs model1.LimitSpeedLine
	tx := db.DB.Model(&model1.LimitSpeedLine{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
