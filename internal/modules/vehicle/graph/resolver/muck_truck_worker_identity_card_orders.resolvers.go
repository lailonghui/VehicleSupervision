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

func (r *mutationResolver) DeleteMuckTruckWorkerIdentityCardOrders(ctx context.Context, where model.MuckTruckWorkerIdentityCardOrdersBoolExp) (*model.MuckTruckWorkerIdentityCardOrdersMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckWorkerIdentityCardOrders{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.MuckTruckWorkerIdentityCardOrders
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
	return &model.MuckTruckWorkerIdentityCardOrdersMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteMuckTruckWorkerIdentityCardOrdersByPk(ctx context.Context, Id int64) (*model1.MuckTruckWorkerIdentityCardOrders, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.MuckTruckWorkerIdentityCardOrders
	tx := db.DB.Model(&model1.MuckTruckWorkerIdentityCardOrders{})
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

func (r *mutationResolver) InsertMuckTruckWorkerIdentityCardOrders(ctx context.Context, objects []*model.MuckTruckWorkerIdentityCardOrdersInsertInput) (*model.MuckTruckWorkerIdentityCardOrdersMutationResponse, error) {
	rs := make([]*model1.MuckTruckWorkerIdentityCardOrders, 0)
	for _, object := range objects {
		v := &model1.MuckTruckWorkerIdentityCardOrders{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.MuckTruckWorkerIdentityCardOrders{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.MuckTruckWorkerIdentityCardOrdersMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertMuckTruckWorkerIdentityCardOrdersOne(ctx context.Context, object model.MuckTruckWorkerIdentityCardOrdersInsertInput) (*model1.MuckTruckWorkerIdentityCardOrders, error) {
	rs := &model1.MuckTruckWorkerIdentityCardOrders{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.MuckTruckWorkerIdentityCardOrders{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateMuckTruckWorkerIdentityCardOrders(ctx context.Context, inc *model.MuckTruckWorkerIdentityCardOrdersIncInput, set *model.MuckTruckWorkerIdentityCardOrdersSetInput, where model.MuckTruckWorkerIdentityCardOrdersBoolExp) (*model.MuckTruckWorkerIdentityCardOrdersMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckWorkerIdentityCardOrders{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.MuckTruckWorkerIdentityCardOrdersMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.MuckTruckWorkerIdentityCardOrdersMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateMuckTruckWorkerIdentityCardOrdersByPk(ctx context.Context, inc *model.MuckTruckWorkerIdentityCardOrdersIncInput, set *model.MuckTruckWorkerIdentityCardOrdersSetInput, Id int64) (*model1.MuckTruckWorkerIdentityCardOrders, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.MuckTruckWorkerIdentityCardOrders{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.MuckTruckWorkerIdentityCardOrders
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) MuckTruckWorkerIdentityCardOrders(ctx context.Context, distinctOn []model.MuckTruckWorkerIdentityCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIdentityCardOrdersOrderBy, where *model.MuckTruckWorkerIdentityCardOrdersBoolExp) ([]*model1.MuckTruckWorkerIdentityCardOrders, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckWorkerIdentityCardOrders{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.MuckTruckWorkerIdentityCardOrders
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) MuckTruckWorkerIdentityCardOrdersAggregate(ctx context.Context, distinctOn []model.MuckTruckWorkerIdentityCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIdentityCardOrdersOrderBy, where *model.MuckTruckWorkerIdentityCardOrdersBoolExp) (*model.MuckTruckWorkerIdentityCardOrdersAggregate, error) {
	var rs model.MuckTruckWorkerIdentityCardOrdersAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.MuckTruckWorkerIdentityCardOrders{})
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

func (r *queryResolver) MuckTruckWorkerIdentityCardOrdersByPk(ctx context.Context, Id int64) (*model1.MuckTruckWorkerIdentityCardOrders, error) {
	var rs model1.MuckTruckWorkerIdentityCardOrders
	tx := db.DB.Model(&model1.MuckTruckWorkerIdentityCardOrders{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}