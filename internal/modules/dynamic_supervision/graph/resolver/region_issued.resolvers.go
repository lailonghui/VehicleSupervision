package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	model1 "VehicleSupervision/internal/modules/dynamic_supervision/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteRegionIssued(ctx context.Context, where model.RegionIssuedBoolExp) (*model.RegionIssuedMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RegionIssued{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.RegionIssued
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
	return &model.RegionIssuedMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteRegionIssuedByPk(ctx context.Context, Id int64) (*model1.RegionIssued, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.RegionIssued
	tx := db.DB.Model(&model1.RegionIssued{})
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

func (r *mutationResolver) InsertRegionIssued(ctx context.Context, objects []*model.RegionIssuedInsertInput) (*model.RegionIssuedMutationResponse, error) {
	rs := make([]*model1.RegionIssued, 0)
	for _, object := range objects {
		v := &model1.RegionIssued{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.RegionIssued{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.RegionIssuedMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertRegionIssuedOne(ctx context.Context, object model.RegionIssuedInsertInput) (*model1.RegionIssued, error) {
	rs := &model1.RegionIssued{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.RegionIssued{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateRegionIssued(ctx context.Context, inc *model.RegionIssuedIncInput, set *model.RegionIssuedSetInput, where model.RegionIssuedBoolExp) (*model.RegionIssuedMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RegionIssued{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.RegionIssuedMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.RegionIssuedMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateRegionIssuedByPk(ctx context.Context, inc *model.RegionIssuedIncInput, set *model.RegionIssuedSetInput, Id int64) (*model1.RegionIssued, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.RegionIssued{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.RegionIssued
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) RegionIssued(ctx context.Context, distinctOn []model.RegionIssuedSelectColumn, limit *int, offset *int, orderBy []*model.RegionIssuedOrderBy, where *model.RegionIssuedBoolExp) ([]*model1.RegionIssued, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.RegionIssued{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.RegionIssued
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) RegionIssuedAggregate(ctx context.Context, distinctOn []model.RegionIssuedSelectColumn, limit *int, offset *int, orderBy []*model.RegionIssuedOrderBy, where *model.RegionIssuedBoolExp) (*model.RegionIssuedAggregate, error) {
	var rs model.RegionIssuedAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.RegionIssued{})
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

func (r *queryResolver) RegionIssuedByPk(ctx context.Context, Id int64) (*model1.RegionIssued, error) {
	var rs model1.RegionIssued
	tx := db.DB.Model(&model1.RegionIssued{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
