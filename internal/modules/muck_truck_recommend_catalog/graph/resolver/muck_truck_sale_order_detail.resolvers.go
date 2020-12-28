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

func (r *mutationResolver) DeleteMuckTruckSaleOrderDetail(ctx context.Context, where model.MuckTruckSaleOrderDetailBoolExp) (*model.MuckTruckSaleOrderDetailMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckSaleOrderDetail{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.MuckTruckSaleOrderDetail
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
	return &model.MuckTruckSaleOrderDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckSaleOrderDetailByPk(ctx context.Context, Id int64) (*model1.MuckTruckSaleOrderDetail, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.MuckTruckSaleOrderDetail
	tx := db.DB.Model(&model1.MuckTruckSaleOrderDetail{})
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

func (r *mutationResolver) InsertMuckTruckSaleOrderDetail(ctx context.Context, objects []*model.MuckTruckSaleOrderDetailInsertInput) (*model.MuckTruckSaleOrderDetailMutationResponse, error) {
	rs := []*model1.MuckTruckSaleOrderDetail{}
	for _, object := range objects {
		v := &model1.MuckTruckSaleOrderDetail{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.MuckTruckSaleOrderDetail{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.MuckTruckSaleOrderDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertMuckTruckSaleOrderDetailOne(ctx context.Context, object model.MuckTruckSaleOrderDetailInsertInput) (*model1.MuckTruckSaleOrderDetail, error) {
	rs := &model1.MuckTruckSaleOrderDetail{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.MuckTruckSaleOrderDetail{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateMuckTruckSaleOrderDetail(ctx context.Context, inc *model.MuckTruckSaleOrderDetailIncInput, set *model.MuckTruckSaleOrderDetailSetInput, where model.MuckTruckSaleOrderDetailBoolExp) (*model.MuckTruckSaleOrderDetailMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckSaleOrderDetail{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.MuckTruckSaleOrderDetailMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.MuckTruckSaleOrderDetailMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateMuckTruckSaleOrderDetailByPk(ctx context.Context, inc *model.MuckTruckSaleOrderDetailIncInput, set *model.MuckTruckSaleOrderDetailSetInput, Id int64) (*model1.MuckTruckSaleOrderDetail, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.MuckTruckSaleOrderDetail{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.MuckTruckSaleOrderDetail
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) MuckTruckSaleOrderDetail(ctx context.Context, distinctOn []model.MuckTruckSaleOrderDetailSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckSaleOrderDetailOrderBy, where *model.MuckTruckSaleOrderDetailBoolExp) ([]*model1.MuckTruckSaleOrderDetail, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckSaleOrderDetail{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.MuckTruckSaleOrderDetail
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) MuckTruckSaleOrderDetailAggregate(ctx context.Context, distinctOn []model.MuckTruckSaleOrderDetailSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckSaleOrderDetailOrderBy, where *model.MuckTruckSaleOrderDetailBoolExp) (*model.MuckTruckSaleOrderDetailAggregate, error) {
	var rs model.MuckTruckSaleOrderDetailAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckSaleOrderDetail{})
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

func (r *queryResolver) MuckTruckSaleOrderDetailByPk(ctx context.Context, Id int64) (*model1.MuckTruckSaleOrderDetail, error) {
	var rs model1.MuckTruckSaleOrderDetail
	tx := db.DB.Model(&model1.MuckTruckSaleOrderDetail{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
