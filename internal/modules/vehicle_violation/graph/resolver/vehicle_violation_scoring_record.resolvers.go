package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_violation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleViolationScoringRecord(ctx context.Context, where model.VehicleViolationScoringRecordBoolExp) (*model.VehicleViolationScoringRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationScoringRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleViolationScoringRecord
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
	return &model.VehicleViolationScoringRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleViolationScoringRecordByPk(ctx context.Context, Id int64) (*model1.VehicleViolationScoringRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleViolationScoringRecord
	tx := db.DB.Model(&model1.VehicleViolationScoringRecord{})
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

func (r *mutationResolver) InsertVehicleViolationScoringRecord(ctx context.Context, objects []*model.VehicleViolationScoringRecordInsertInput) (*model.VehicleViolationScoringRecordMutationResponse, error) {
	rs := make([]*model1.VehicleViolationScoringRecord, 0)
	for _, object := range objects {
		v := &model1.VehicleViolationScoringRecord{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleViolationScoringRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleViolationScoringRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleViolationScoringRecordOne(ctx context.Context, object model.VehicleViolationScoringRecordInsertInput) (*model1.VehicleViolationScoringRecord, error) {
	rs := &model1.VehicleViolationScoringRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleViolationScoringRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleViolationScoringRecord(ctx context.Context, inc *model.VehicleViolationScoringRecordIncInput, set *model.VehicleViolationScoringRecordSetInput, where model.VehicleViolationScoringRecordBoolExp) (*model.VehicleViolationScoringRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationScoringRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleViolationScoringRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleViolationScoringRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleViolationScoringRecordByPk(ctx context.Context, inc *model.VehicleViolationScoringRecordIncInput, set *model.VehicleViolationScoringRecordSetInput, Id int64) (*model1.VehicleViolationScoringRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleViolationScoringRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleViolationScoringRecord
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleViolationScoringRecord(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) ([]*model1.VehicleViolationScoringRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationScoringRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleViolationScoringRecord
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleViolationScoringRecordAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringRecordOrderBy, where *model.VehicleViolationScoringRecordBoolExp) (*model.VehicleViolationScoringRecordAggregate, error) {
	var rs model.VehicleViolationScoringRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationScoringRecord{})
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

func (r *queryResolver) VehicleViolationScoringRecordByPk(ctx context.Context, Id int64) (*model1.VehicleViolationScoringRecord, error) {
	var rs model1.VehicleViolationScoringRecord
	tx := db.DB.Model(&model1.VehicleViolationScoringRecord{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
