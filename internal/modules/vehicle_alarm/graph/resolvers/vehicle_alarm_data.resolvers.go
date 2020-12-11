package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleAlarmData(ctx context.Context, where model.VehicleAlarmDataBoolExp) (*model.VehicleAlarmDataMutationResponse, error) {
	var err error
	var rs []*model.VehicleAlarmData
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleAlarmData{})
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
	return &model.VehicleAlarmDataMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteVehicleAlarmDataByPk(ctx context.Context, id int64, vehicleAlarmDataID string) (*model.VehicleAlarmData, error) {
	var rs model.VehicleAlarmData
	var err error
	tx := db.DB.Model(&model.VehicleAlarmData{})
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

func (r *mutationResolver) InsertVehicleAlarmData(ctx context.Context, objects []*model.VehicleAlarmDataInsertInput, onConflict *model.VehicleAlarmDataOnConflict) (*model.VehicleAlarmDataMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleAlarmDataOne(ctx context.Context, object model.VehicleAlarmDataInsertInput, onConflict *model.VehicleAlarmDataOnConflict) (*model.VehicleAlarmData, error) {
	v := &model.VehicleAlarmData{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateVehicleAlarmData(ctx context.Context, inc *model.VehicleAlarmDataIncInput, set *model.VehicleAlarmDataSetInput, where model.VehicleAlarmDataBoolExp) (*model.VehicleAlarmDataMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleAlarmData{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.VehicleAlarmDataMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.VehicleAlarmData
	tx.Scan(&vehicleList)
	return &model.VehicleAlarmDataMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateVehicleAlarmDataByPk(ctx context.Context, inc *model.VehicleAlarmDataIncInput, set *model.VehicleAlarmDataSetInput, pkColumns model.VehicleAlarmDataPkColumnsInput) (*model.VehicleAlarmData, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.VehicleAlarmData{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.VehicleAlarmData
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) VehicleAlarmData(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) ([]*model.VehicleAlarmData, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleAlarmData{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.VehicleAlarmData
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) VehicleAlarmDataAggregate(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) (*model.VehicleAlarmDataAggregate, error) {
	var rs model.VehicleAlarmDataAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleAlarmData{})
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

func (r *queryResolver) VehicleAlarmDataByPk(ctx context.Context, id int64, vehicleAlarmDataID string) (*model.VehicleAlarmData, error) {
	var rs model.VehicleAlarmData
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) VehicleAlarmData(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) (<-chan []*model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleAlarmDataAggregate(ctx context.Context, distinctOn []model.VehicleAlarmDataSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmDataOrderBy, where *model.VehicleAlarmDataBoolExp) (<-chan *model.VehicleAlarmDataAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleAlarmDataByPk(ctx context.Context, id int64, vehicleAlarmDataID string) (<-chan *model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}
