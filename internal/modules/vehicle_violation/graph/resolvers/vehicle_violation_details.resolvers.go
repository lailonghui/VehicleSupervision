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

func (r *mutationResolver) DeleteVehicleViolationDetails(ctx context.Context, where model.VehicleViolationDetailsBoolExp) (*model.VehicleViolationDetailsMutationResponse, error) {
	var err error
	var rs []*model.VehicleViolationDetails
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationDetails{})
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
	return &model.VehicleViolationDetailsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteVehicleViolationDetailsByPk(ctx context.Context, id int64) (*model.VehicleViolationDetails, error) {
	var rs model.VehicleViolationDetails
	var err error
	tx := db.DB.Model(&model.VehicleViolationDetails{})
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

func (r *mutationResolver) InsertVehicleViolationDetails(ctx context.Context, objects []*model.VehicleViolationDetailsInsertInput, onConflict *model.VehicleViolationDetailsOnConflict) (*model.VehicleViolationDetailsMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleViolationDetailsOne(ctx context.Context, object model.VehicleViolationDetailsInsertInput, onConflict *model.VehicleViolationDetailsOnConflict) (*model.VehicleViolationDetails, error) {
	v := &model.VehicleViolationDetails{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateVehicleViolationDetails(ctx context.Context, inc *model.VehicleViolationDetailsIncInput, set *model.VehicleViolationDetailsSetInput, where model.VehicleViolationDetailsBoolExp) (*model.VehicleViolationDetailsMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationDetails{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.VehicleViolationDetailsMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.VehicleViolationDetails
	tx.Scan(&vehicleList)
	return &model.VehicleViolationDetailsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateVehicleViolationDetailsByPk(ctx context.Context, inc *model.VehicleViolationDetailsIncInput, set *model.VehicleViolationDetailsSetInput, pkColumns model.VehicleViolationDetailsPkColumnsInput) (*model.VehicleViolationDetails, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.VehicleViolationDetails{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.VehicleViolationDetails
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) VehicleViolationDetails(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) ([]*model.VehicleViolationDetails, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationDetails{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.VehicleViolationDetails
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) VehicleViolationDetailsAggregate(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) (*model.VehicleViolationDetailsAggregate, error) {
	var rs model.VehicleViolationDetailsAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationDetails{})
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

func (r *queryResolver) VehicleViolationDetailsByPk(ctx context.Context, id int64) (*model.VehicleViolationDetails, error) {
	var rs model.VehicleViolationDetails
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) VehicleViolationDetails(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) (<-chan []*model.VehicleViolationDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationDetailsAggregate(ctx context.Context, distinctOn []model.VehicleViolationDetailsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationDetailsOrderBy, where *model.VehicleViolationDetailsBoolExp) (<-chan *model.VehicleViolationDetailsAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationDetailsByPk(ctx context.Context, id int64) (<-chan *model.VehicleViolationDetails, error) {
	panic(fmt.Errorf("not implemented"))
}
