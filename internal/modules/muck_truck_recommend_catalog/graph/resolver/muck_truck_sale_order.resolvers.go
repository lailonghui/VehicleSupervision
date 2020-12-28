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

func (r *mutationResolver) DeleteMuckTruckSaleOrder(ctx context.Context, where model.MuckTruckSaleOrderBoolExp) (*model.MuckTruckSaleOrderMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckSaleOrder{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.MuckTruckSaleOrder
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
	return &model.MuckTruckSaleOrderMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckSaleOrderByPk(ctx context.Context, Id int64) (*model1.MuckTruckSaleOrder, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.MuckTruckSaleOrder
	tx := db.DB.Model(&model1.MuckTruckSaleOrder{})
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

func (r *mutationResolver) InsertMuckTruckSaleOrder(ctx context.Context, objects []*model.MuckTruckSaleOrderInsertInput) (*model.MuckTruckSaleOrderMutationResponse, error) {
	rs := []*model1.MuckTruckSaleOrder{}
	for _, object := range objects {
		v := &model1.MuckTruckSaleOrder{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.MuckTruckSaleOrder{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.MuckTruckSaleOrderMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertMuckTruckSaleOrderOne(ctx context.Context, object model.MuckTruckSaleOrderInsertInput) (*model1.MuckTruckSaleOrder, error) {
	rs := &model1.MuckTruckSaleOrder{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.MuckTruckSaleOrder{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateMuckTruckSaleOrder(ctx context.Context, inc *model.MuckTruckSaleOrderIncInput, set *model.MuckTruckSaleOrderSetInput, where model.MuckTruckSaleOrderBoolExp) (*model.MuckTruckSaleOrderMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckSaleOrder{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.MuckTruckSaleOrderMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.MuckTruckSaleOrderMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateMuckTruckSaleOrderByPk(ctx context.Context, inc *model.MuckTruckSaleOrderIncInput, set *model.MuckTruckSaleOrderSetInput, Id int64) (*model1.MuckTruckSaleOrder, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.MuckTruckSaleOrder{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.MuckTruckSaleOrder
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) MuckTruckSaleOrder(ctx context.Context, distinctOn []model.MuckTruckSaleOrderSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckSaleOrderOrderBy, where *model.MuckTruckSaleOrderBoolExp) ([]*model1.MuckTruckSaleOrder, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckSaleOrder{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.MuckTruckSaleOrder
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) MuckTruckSaleOrderAggregate(ctx context.Context, distinctOn []model.MuckTruckSaleOrderSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckSaleOrderOrderBy, where *model.MuckTruckSaleOrderBoolExp) (*model.MuckTruckSaleOrderAggregate, error) {
	var rs model.MuckTruckSaleOrderAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckSaleOrder{})
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

func (r *queryResolver) MuckTruckSaleOrderByPk(ctx context.Context, Id int64) (*model1.MuckTruckSaleOrder, error) {
	var rs model1.MuckTruckSaleOrder
	tx := db.DB.Model(&model1.MuckTruckSaleOrder{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
