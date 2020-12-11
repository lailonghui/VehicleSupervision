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

//根据选定条件删除六合一平台同步车辆信息
func (r *mutationResolver) DeleteJjVehicle(ctx context.Context, where model.JjVehicleBoolExp) (*model.JjVehicleMutationResponse, error) {
	var err error
	var rs []*model.JjVehicle
	qt := util2.NewQueryTranslator(db.DB, &model.JjVehicle{})
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
	return &model.JjVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

//根据主键删除六合一平台同步车辆信息
func (r *mutationResolver) DeleteJjVehicleByPk(ctx context.Context, id int64) (*model.JjVehicle, error) {
	var rs model.JjVehicle
	var err error
	tx := db.DB.Model(&model.JjVehicle{})
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

//插入一条六合一平台同步车辆表信息
func (r *mutationResolver) InsertJjVehicleOne(ctx context.Context, object model.JjVehicleInsertInput, onConflict *model.JjVehicleOnConflict) (*model.JjVehicle, error) {
	v := &model.JjVehicle{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

//根据选定条件更新六合一平台同步车辆表信息
func (r *mutationResolver) UpdateJjVehicle(ctx context.Context, inc *model.JjVehicleIncInput, set *model.JjVehicleSetInput, where model.JjVehicleBoolExp) (*model.JjVehicleMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.JjVehicle{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.JjVehicleMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.JjVehicle
	tx.Scan(&vehicleList)
	return &model.JjVehicleMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

//根据主键更新六合一平台同步车辆表信息
func (r *mutationResolver) UpdateJjVehicleByPk(ctx context.Context, inc *model.JjVehicleIncInput, set *model.JjVehicleSetInput, pkColumns model.JjVehiclePkColumnsInput) (*model.JjVehicle, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.JjVehicle{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.JjVehicle
	tx = tx.First(&rs)
	return &rs, err
}

//根据选定条件获取六合一平台同步车辆表信息
func (r *queryResolver) JjVehicle(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) ([]*model.JjVehicle, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.JjVehicle{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.JjVehicle
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

//根据选定条件获取六合一平台同步车辆表信息
func (r *queryResolver) JjVehicleAggregate(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (*model.JjVehicleAggregate, error) {
	var rs model.JjVehicleAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.JjVehicle{})
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

//根据主键获取六合一平台同步车辆表信息
func (r *queryResolver) JjVehicleByPk(ctx context.Context, id int64) (*model.JjVehicle, error) {
	var rs model.JjVehicle
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *mutationResolver) InsertJjVehicle(ctx context.Context, objects []*model.JjVehicleInsertInput, onConflict *model.JjVehicleOnConflict) (*model.JjVehicleMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) JjVehicle(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (<-chan []*model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) JjVehicleAggregate(ctx context.Context, distinctOn []model.JjVehicleSelectColumn, limit *int, offset *int, orderBy []*model.JjVehicleOrderBy, where *model.JjVehicleBoolExp) (<-chan *model.JjVehicleAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) JjVehicleByPk(ctx context.Context, id int64) (<-chan *model.JjVehicle, error) {
	panic(fmt.Errorf("not implemented"))
}
