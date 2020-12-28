package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/muck_truck_recommend_catalog/graph/model"
	model1 "VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteSellerFiling(ctx context.Context, where model.SellerFilingBoolExp) (*model.SellerFilingMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SellerFiling{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.SellerFiling
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
	return &model.SellerFilingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteSellerFilingByPk(ctx context.Context, Id int64) (*model1.SellerFiling, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.SellerFiling
	tx := db.DB.Model(&model1.SellerFiling{})
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

func (r *mutationResolver) InsertSellerFiling(ctx context.Context, objects []*model.SellerFilingInsertInput) (*model.SellerFilingMutationResponse, error) {
	rs := []*model1.SellerFiling{}
	for _, object := range objects {
		v := &model1.SellerFiling{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.SellerFiling{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.SellerFilingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertSellerFilingOne(ctx context.Context, object model.SellerFilingInsertInput) (*model1.SellerFiling, error) {
	rs := &model1.SellerFiling{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.SellerFiling{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateSellerFiling(ctx context.Context, inc *model.SellerFilingIncInput, set *model.SellerFilingSetInput, where model.SellerFilingBoolExp) (*model.SellerFilingMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SellerFiling{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SellerFilingMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.SellerFilingMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateSellerFilingByPk(ctx context.Context, inc *model.SellerFilingIncInput, set *model.SellerFilingSetInput, Id int64) (*model1.SellerFiling, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.SellerFiling{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.SellerFiling
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) SellerFiling(ctx context.Context, distinctOn []model.SellerFilingSelectColumn, limit *int, offset *int, orderBy []*model.SellerFilingOrderBy, where *model.SellerFilingBoolExp) ([]*model1.SellerFiling, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SellerFiling{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.SellerFiling
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) SellerFilingAggregate(ctx context.Context, distinctOn []model.SellerFilingSelectColumn, limit *int, offset *int, orderBy []*model.SellerFilingOrderBy, where *model.SellerFilingBoolExp) (*model.SellerFilingAggregate, error) {
	var rs model.SellerFilingAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.SellerFiling{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		return nil, err
	}
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &rs, nil
}

func (r *queryResolver) SellerFilingByPk(ctx context.Context, Id int64) (*model1.SellerFiling, error) {
	var rs model1.SellerFiling
	tx := db.DB.Model(&model1.SellerFiling{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
