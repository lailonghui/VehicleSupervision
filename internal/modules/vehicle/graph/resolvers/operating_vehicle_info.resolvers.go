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

func (r *mutationResolver) DeleteOperatingVehicleInfo(ctx context.Context, where model.OperatingVehicleInfoBoolExp) (*model.OperatingVehicleInfoMutationResponse, error) {
	var err error
	var rs []*model.OperatingVehicleInfo
	qt := util2.NewQueryTranslator(db.DB, &model.OperatingVehicleInfo{})
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
	return &model.OperatingVehicleInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteOperatingVehicleInfoByPk(ctx context.Context, operatingVehicleID int64) (*model.OperatingVehicleInfo, error) {
	var rs model.OperatingVehicleInfo
	var err error
	tx := db.DB.Model(&model.OperatingVehicleInfo{})
	preloads := util2.GetPreloads(ctx)
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("operation_vehicle_id = ? ", operatingVehicleID).First(&rs)
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

func (r *mutationResolver) InsertOperatingVehicleInfo(ctx context.Context, objects []*model.OperatingVehicleInfoInsertInput, onConflict *model.OperatingVehicleInfoOnConflict) (*model.OperatingVehicleInfoMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertOperatingVehicleInfoOne(ctx context.Context, object model.OperatingVehicleInfoInsertInput, onConflict *model.OperatingVehicleInfoOnConflict) (*model.OperatingVehicleInfo, error) {
	v := &model.OperatingVehicleInfo{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateOperatingVehicleInfo(ctx context.Context, inc *model.OperatingVehicleInfoIncInput, set *model.OperatingVehicleInfoSetInput, where model.OperatingVehicleInfoBoolExp) (*model.OperatingVehicleInfoMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.OperatingVehicleInfo{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.OperatingVehicleInfoMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.OperatingVehicleInfo
	tx.Scan(&vehicleList)
	return &model.OperatingVehicleInfoMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateOperatingVehicleInfoByPk(ctx context.Context, inc *model.OperatingVehicleInfoIncInput, set *model.OperatingVehicleInfoSetInput, pkColumns model.OperatingVehicleInfoPkColumnsInput) (*model.OperatingVehicleInfo, error) {
	var err error
	tx := db.DB.Where("operation_vehicle_id = ? ", pkColumns.OperatingVehicleID)
	qt := util2.NewQueryTranslator(tx, &model.OperatingVehicleInfo{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.OperatingVehicleInfo
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) OperatingVehicleInfo(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) ([]*model.OperatingVehicleInfo, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.OperatingVehicleInfo{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.OperatingVehicleInfo
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) OperatingVehicleInfoAggregate(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) (*model.OperatingVehicleInfoAggregate, error) {
	var rs model.OperatingVehicleInfoAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.OperatingVehicleInfo{})
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

func (r *queryResolver) OperatingVehicleInfoByPk(ctx context.Context, operatingVehicleID int64) (*model.OperatingVehicleInfo, error) {
	var rs model.OperatingVehicleInfo
	err := db.DB.Where("operation_vehicle_id = ?", operatingVehicleID).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) OperatingVehicleInfo(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) (<-chan []*model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OperatingVehicleInfoAggregate(ctx context.Context, distinctOn []model.OperatingVehicleInfoSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleInfoOrderBy, where *model.OperatingVehicleInfoBoolExp) (<-chan *model.OperatingVehicleInfoAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) OperatingVehicleInfoByPk(ctx context.Context, operatingVehicleID int64) (<-chan *model.OperatingVehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
