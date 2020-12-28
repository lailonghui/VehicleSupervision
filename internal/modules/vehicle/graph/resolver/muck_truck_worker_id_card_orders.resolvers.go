package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteMuckTruckWorkerIdCardOrders(ctx context.Context, where model.MuckTruckWorkerIdCardOrdersBoolExp) (*model.MuckTruckWorkerIdCardOrdersMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckWorkerIdCardOrders{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.MuckTruckWorkerIdCardOrders
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
	return &model.MuckTruckWorkerIdCardOrdersMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckWorkerIdCardOrdersByPk(ctx context.Context, Id int64) (*model1.MuckTruckWorkerIdCardOrders, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.MuckTruckWorkerIdCardOrders
	tx := db.DB.Model(&model1.MuckTruckWorkerIdCardOrders{})
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

func (r *mutationResolver) InsertMuckTruckWorkerIdCardOrders(ctx context.Context, objects []*model.MuckTruckWorkerIdCardOrdersInsertInput) (*model.MuckTruckWorkerIdCardOrdersMutationResponse, error) {
	rs := []*model1.MuckTruckWorkerIdCardOrders{}
	for _, object := range objects {
		v := &model1.MuckTruckWorkerIdCardOrders{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.MuckTruckWorkerIdCardOrders{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.MuckTruckWorkerIdCardOrdersMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertMuckTruckWorkerIdCardOrdersOne(ctx context.Context, object model.MuckTruckWorkerIdCardOrdersInsertInput) (*model1.MuckTruckWorkerIdCardOrders, error) {
	rs := &model1.MuckTruckWorkerIdCardOrders{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.MuckTruckWorkerIdCardOrders{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateMuckTruckWorkerIdCardOrders(ctx context.Context, inc *model.MuckTruckWorkerIdCardOrdersIncInput, set *model.MuckTruckWorkerIdCardOrdersSetInput, where model.MuckTruckWorkerIdCardOrdersBoolExp) (*model.MuckTruckWorkerIdCardOrdersMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckWorkerIdCardOrders{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.MuckTruckWorkerIdCardOrdersMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.MuckTruckWorkerIdCardOrdersMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateMuckTruckWorkerIdCardOrdersByPk(ctx context.Context, inc *model.MuckTruckWorkerIdCardOrdersIncInput, set *model.MuckTruckWorkerIdCardOrdersSetInput, Id int64) (*model1.MuckTruckWorkerIdCardOrders, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.MuckTruckWorkerIdCardOrders{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.MuckTruckWorkerIdCardOrders
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) MuckTruckWorkerIdCardOrders(ctx context.Context, distinctOn []model.MuckTruckWorkerIdCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIdCardOrdersOrderBy, where *model.MuckTruckWorkerIdCardOrdersBoolExp) ([]*model1.MuckTruckWorkerIdCardOrders, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckWorkerIdCardOrders{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.MuckTruckWorkerIdCardOrders
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) MuckTruckWorkerIdCardOrdersAggregate(ctx context.Context, distinctOn []model.MuckTruckWorkerIdCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIdCardOrdersOrderBy, where *model.MuckTruckWorkerIdCardOrdersBoolExp) (*model.MuckTruckWorkerIdCardOrdersAggregate, error) {
	var rs model.MuckTruckWorkerIdCardOrdersAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckWorkerIdCardOrders{})
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

func (r *queryResolver) MuckTruckWorkerIdCardOrdersByPk(ctx context.Context, Id int64) (*model1.MuckTruckWorkerIdCardOrders, error) {
	var rs model1.MuckTruckWorkerIdCardOrders
	tx := db.DB.Model(&model1.MuckTruckWorkerIdCardOrders{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
