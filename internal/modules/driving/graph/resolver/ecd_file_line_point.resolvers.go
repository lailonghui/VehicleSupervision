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

func (r *mutationResolver) DeleteEcdFileLinePoint(ctx context.Context, where model.EcdFileLinePointBoolExp) (*model.EcdFileLinePointMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileLinePoint{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileLinePoint
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
	return &model.EcdFileLinePointMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEcdFileLinePointByPk(ctx context.Context, Id int64) (*model1.EcdFileLinePoint, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.EcdFileLinePoint
	tx := db.DB.Model(&model1.EcdFileLinePoint{})
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

func (r *mutationResolver) InsertEcdFileLinePoint(ctx context.Context, objects []*model.EcdFileLinePointInsertInput) (*model.EcdFileLinePointMutationResponse, error) {
	rs := make([]*model1.EcdFileLinePoint, 0)
	for _, object := range objects {
		v := &model1.EcdFileLinePoint{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.EcdFileLinePoint{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EcdFileLinePointMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEcdFileLinePointOne(ctx context.Context, object model.EcdFileLinePointInsertInput) (*model1.EcdFileLinePoint, error) {
	rs := &model1.EcdFileLinePoint{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.EcdFileLinePoint{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEcdFileLinePoint(ctx context.Context, inc *model.EcdFileLinePointIncInput, set *model.EcdFileLinePointSetInput, where model.EcdFileLinePointBoolExp) (*model.EcdFileLinePointMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileLinePoint{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EcdFileLinePointMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.EcdFileLinePoint
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
	return &model.EcdFileLinePointMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateEcdFileLinePointByPk(ctx context.Context, inc *model.EcdFileLinePointIncInput, set *model.EcdFileLinePointSetInput, Id int64) (*model1.EcdFileLinePoint, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.EcdFileLinePoint{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.EcdFileLinePoint
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) EcdFileLinePoint(ctx context.Context, distinctOn []model.EcdFileLinePointSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileLinePointOrderBy, where *model.EcdFileLinePointBoolExp) ([]*model1.EcdFileLinePoint, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileLinePoint{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.EcdFileLinePoint
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) EcdFileLinePointAggregate(ctx context.Context, distinctOn []model.EcdFileLinePointSelectColumn, limit *int, offset *int, orderBy []*model.EcdFileLinePointOrderBy, where *model.EcdFileLinePointBoolExp) (*model.EcdFileLinePointAggregate, error) {
	var rs model.EcdFileLinePointAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.EcdFileLinePoint{})
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

func (r *queryResolver) EcdFileLinePointByPk(ctx context.Context, Id int64) (*model1.EcdFileLinePoint, error) {
	var rs model1.EcdFileLinePoint
	tx := db.DB.Model(&model1.EcdFileLinePoint{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
