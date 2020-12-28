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

func (r *mutationResolver) DeleteVehicleViolationScoringRecord(ctx context.Context, where model.VehicleViolationScoringRecordBoolExp) (*model.VehicleViolationScoringRecordMutationResponse, error) {
	var err error
	var rs []*model.VehicleViolationScoringRecord
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationScoringRecord{})
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
	return &model.VehicleViolationScoringRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, err
}

func (r *mutationResolver) DeleteVehicleViolationScoringRecordByPk(ctx context.Context, id int64, violationScoringID string) (*model.VehicleViolationScoringRecord, error) {
	var rs model.VehicleViolationScoringRecord
	var err error
	tx := db.DB.Model(&model.VehicleViolationScoringRecord{})
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

func (r *mutationResolver) InsertVehicleViolationScoringRecord(ctx context.Context, objects []*model.VehicleViolationScoringRecordInsertInput, onConflict *model.VehicleViolationScoringRecordOnConflict) (*model.VehicleViolationScoringRecordMutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InsertVehicleViolationScoringRecordOne(ctx context.Context, object model.VehicleViolationScoringRecordInsertInput, onConflict *model.VehicleViolationScoringRecordOnConflict) (*model.VehicleViolationScoringRecord, error) {
	v := &model.VehicleViolationScoringRecord{}
	util.StructAssign(v, &object)
	err := db.DB.Create(v).Error
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func (r *mutationResolver) UpdateVehicleViolationScoringRecord(ctx context.Context, inc *model.VehicleViolationScoringRecordIncInput, set *model.VehicleViolationScoringRecordSetInput, where model.VehicleViolationScoringRecordBoolExp) (*model.VehicleViolationScoringRecordMutationResponse, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationScoringRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return &model.VehicleViolationScoringRecordMutationResponse{
			AffectedRows: 0,
		}, err
	}
	var vehicleList []*model.VehicleViolationScoringRecord
	tx.Scan(&vehicleList)
	return &model.VehicleViolationScoringRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    vehicleList,
	}, err
}

func (r *mutationResolver) UpdateVehicleViolationScoringRecordByPk(ctx context.Context, inc *model.VehicleViolationScoringRecordIncInput, set *model.VehicleViolationScoringRecordSetInput, pkColumns model.VehicleViolationScoringRecordPkColumnsInput) (*model.VehicleViolationScoringRecord, error) {
	var err error
	tx := db.DB.Where("id = ? ", pkColumns.ID)
	qt := util2.NewQueryTranslator(tx, &model.VehicleViolationScoringRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if rowAffected := tx.RowsAffected; rowAffected == 0 {
		err = errors.New("0 rows affected or returned")
		return nil, err
	}
	var rs model.VehicleViolationScoringRecord
	tx = tx.First(&rs)
	return &rs, err
}

func (r *queryResolver) VehicleViolationScoringRecord(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) ([]*model.VehicleViolationScoringRecord, error) {
	var err error
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationScoringRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.VehicleViolationScoringRecord
	tx = tx.Find(&rs)
	if err = tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, err
}

func (r *queryResolver) VehicleViolationScoringRecordAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) (*model.VehicleViolationScoringRecordAggregate, error) {
	var rs model.VehicleViolationScoringRecordAggregate
	qt := util2.NewQueryTranslator(db.DB, &model.VehicleViolationScoringRecord{})
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

func (r *queryResolver) VehicleViolationScoringRecordByPk(ctx context.Context, id int64, violationScoringID string) (*model.VehicleViolationScoringRecord, error) {
	var rs model.VehicleViolationScoringRecord
	err := db.DB.Where("id = ?", id).First(&rs).Error
	return &rs, err
}

func (r *subscriptionResolver) VehicleViolationScoringRecord(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) (<-chan []*model.VehicleViolationScoringRecord, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringRecordAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) (<-chan *model.VehicleViolationScoringRecordAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VehicleViolationScoringRecordByPk(ctx context.Context, id int64, violationScoringID string) (<-chan *model.VehicleViolationScoringRecord, error) {
	panic(fmt.Errorf("not implemented"))
}