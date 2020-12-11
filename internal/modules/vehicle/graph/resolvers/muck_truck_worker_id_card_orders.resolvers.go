package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteMuckTruckWorkerIDCardOrders(ctx context.Context, where model.MuckTruckWorkerIDCardOrdersBoolExp) (*model.MuckTruckWorkerIDCardOrdersMutationResponse, error) {
	var err error
	var rs []*model.MuckTruckWorkerIDCardOrders
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckWorkerIDCardOrders{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util2.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx := tx.Select(preloads)
		tx = tx.Find(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err = tx.Error; err != nil {
		return nil, err
	}
	return &model.MuckTruckWorkerIDCardOrdersMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteMuckTruckWorkerIDCardOrdersByPk(ctx context.Context, id int64) (*model.MuckTruckWorkerIDCardOrders, error) {
	var rs model.MuckTruckWorkerIDCardOrders
	var err error
	tx := db.DB.Model(&model.MuckTruckWorkerIDCardOrders{})
	preloads := util2.GetPreloads(ctx)
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ? ", id).First(&rs)
		// 如果查询结果含有错误，则返回错误
		if err := tx.Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = errors.New("0 rows affected or returned")
				return nil, err
			}
		}
	}
	// 删除
	tx = tx.Delete(nil)
	if err = tx.Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *mutationResolver) InsertMuckTruckWorkerIDCardOrders(ctx context.Context, objects []*model.MuckTruckWorkerIDCardOrdersInsertInput, onConflict *model.MuckTruckWorkerIDCardOrdersOnConflict) (*model.MuckTruckWorkerIDCardOrdersMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertMuckTruckWorkerIDCardOrdersOne(ctx context.Context, object model.MuckTruckWorkerIDCardOrdersInsertInput, onConflict *model.MuckTruckWorkerIDCardOrdersOnConflict) (*model.MuckTruckWorkerIDCardOrders, error) {
	v := &model.MuckTruckWorkerIDCardOrders{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateMuckTruckWorkerIDCardOrders(ctx context.Context, inc *model.MuckTruckWorkerIDCardOrdersIncInput, set *model.MuckTruckWorkerIDCardOrdersSetInput, where model.MuckTruckWorkerIDCardOrdersBoolExp) (*model.MuckTruckWorkerIDCardOrdersMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckWorkerIDCardOrders{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.MuckTruckWorkerIDCardOrdersMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.MuckTruckWorkerIDCardOrders
	tx.Scan(&vehicleList)
	return &model.MuckTruckWorkerIDCardOrdersMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateMuckTruckWorkerIDCardOrdersByPk(ctx context.Context, inc *model.MuckTruckWorkerIDCardOrdersIncInput, set *model.MuckTruckWorkerIDCardOrdersSetInput, pkColumns model.MuckTruckWorkerIDCardOrdersPkColumnsInput) (*model.MuckTruckWorkerIDCardOrders, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.MuckTruckWorkerIDCardOrders{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.MuckTruckWorkerIDCardOrders
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) MuckTruckWorkerIDCardOrders(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) ([]*model.MuckTruckWorkerIDCardOrders, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckWorkerIDCardOrders{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.MuckTruckWorkerIDCardOrders
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) MuckTruckWorkerIDCardOrdersAggregate(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) (*model.MuckTruckWorkerIDCardOrdersAggregate, error) {
	var rs model.MuckTruckWorkerIDCardOrdersAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.MuckTruckWorkerIDCardOrders{})
	tx, err := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Aggregate(&rs, ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	return &rs, err
}

func (r *queryResolver) MuckTruckWorkerIDCardOrdersByPk(ctx context.Context, id int64) (*model.MuckTruckWorkerIDCardOrders, error) {
	var rs model.MuckTruckWorkerIDCardOrders
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) MuckTruckWorkerIDCardOrders(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) (<-chan []*model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckWorkerIDCardOrdersAggregate(ctx context.Context, distinctOn []model.MuckTruckWorkerIDCardOrdersSelectColumn, limit *int, offset *int, orderBy []*model.MuckTruckWorkerIDCardOrdersOrderBy, where *model.MuckTruckWorkerIDCardOrdersBoolExp) (<-chan *model.MuckTruckWorkerIDCardOrdersAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MuckTruckWorkerIDCardOrdersByPk(ctx context.Context, id int64) (<-chan *model.MuckTruckWorkerIDCardOrders, error) {
	panic(fmt.Errorf("not implemented"))
}
