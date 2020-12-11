package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	util2 "VehicleSupervision/pkg/graphql/util"
	"VehicleSupervision/pkg/util"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleDetainDetails(ctx context.Context, where model.VehicleDetainDetailsBoolExp) (*model.VehicleDetainDetailsMutationResponse, error) {
	var err error
	var rs []*model.VehicleDetainDetails
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleDetainDetails{})
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
	return &model.VehicleDetainDetailsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteVehicleDetainDetailsByPk(ctx context.Context, id int64) (*model.VehicleDetainDetails, error) {
	var rs model.VehicleDetainDetails
	var err error
	tx := db.DB.Model(&model.VehicleDetainDetails{})
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

func (r *mutationResolver) InsertVehicleDetainDetails(ctx context.Context, objects []*model.VehicleDetainDetailsInsertInput, onConflict *model.VehicleDetainDetailsOnConflict) (*model.VehicleDetainDetailsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleDetainDetailsOne(ctx context.Context, object model.VehicleDetainDetailsInsertInput, onConflict *model.VehicleDetainDetailsOnConflict) (*model.VehicleDetainDetails, error) {
	v := &model.VehicleDetainDetails{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateVehicleDetainDetails(ctx context.Context, inc *model.VehicleDetainDetailsIncInput, set *model.VehicleDetainDetailsSetInput, where model.VehicleDetainDetailsBoolExp) (*model.VehicleDetainDetailsMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleDetainDetails{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.VehicleDetainDetailsMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.VehicleDetainDetails
	tx.Scan(&vehicleList)
	return &model.VehicleDetainDetailsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateVehicleDetainDetailsByPk(ctx context.Context, inc *model.VehicleDetainDetailsIncInput, set *model.VehicleDetainDetailsSetInput, pkColumns model.VehicleDetainDetailsPkColumnsInput) (*model.VehicleDetainDetails, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.VehicleDetainDetails{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.VehicleDetainDetails
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) VehicleDetainDetails(ctx context.Context, distinctOn []model.VehicleDetainDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDetainDetailsOrderBy, where *model.VehicleDetainDetailsBoolExp) ([]*model.VehicleDetainDetails, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleDetainDetails{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.VehicleDetainDetails
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) VehicleDetainDetailsAggregate(ctx context.Context, distinctOn []model.VehicleDetainDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDetainDetailsOrderBy, where *model.VehicleDetainDetailsBoolExp) (*model.VehicleDetainDetailsAggregate, error) {
	var rs model.VehicleDetainDetailsAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleDetainDetails{})
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

func (r *queryResolver) VehicleDetainDetailsByPk(ctx context.Context, id int64) (*model.VehicleDetainDetails, error) {
	var rs model.VehicleDetainDetails
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) VehicleDetainDetails(ctx context.Context, distinctOn []model.VehicleDetainDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDetainDetailsOrderBy, where *model.VehicleDetainDetailsBoolExp) (<-chan []*model.VehicleDetainDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDetainDetailsAggregate(ctx context.Context, distinctOn []model.VehicleDetainDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleDetainDetailsOrderBy, where *model.VehicleDetainDetailsBoolExp) (<-chan *model.VehicleDetainDetailsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleDetainDetailsByPk(ctx context.Context, id int64) (<-chan *model.VehicleDetainDetails, error) {
	panic(fmt.Errorf("not implemented"))
}
