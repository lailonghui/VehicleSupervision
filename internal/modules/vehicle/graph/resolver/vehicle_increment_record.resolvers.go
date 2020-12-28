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

func (r *mutationResolver) DeleteVehicleIncrementRecord(ctx context.Context, where model.VehicleIncrementRecordBoolExp) (*model.VehicleIncrementRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleIncrementRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleIncrementRecord
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
	return &model.VehicleIncrementRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleIncrementRecordByPk(ctx context.Context, Id int64) (*model1.VehicleIncrementRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleIncrementRecord
	tx := db.DB.Model(&model1.VehicleIncrementRecord{})
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

func (r *mutationResolver) InsertVehicleIncrementRecord(ctx context.Context, objects []*model.VehicleIncrementRecordInsertInput) (*model.VehicleIncrementRecordMutationResponse, error) {
	rs := []*model1.VehicleIncrementRecord{}
	for _, object := range objects {
		v := &model1.VehicleIncrementRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleIncrementRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleIncrementRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleIncrementRecordOne(ctx context.Context, object model.VehicleIncrementRecordInsertInput) (*model1.VehicleIncrementRecord, error) {
	rs := &model1.VehicleIncrementRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleIncrementRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleIncrementRecord(ctx context.Context, inc *model.VehicleIncrementRecordIncInput, set *model.VehicleIncrementRecordSetInput, where model.VehicleIncrementRecordBoolExp) (*model.VehicleIncrementRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleIncrementRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleIncrementRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleIncrementRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleIncrementRecordByPk(ctx context.Context, inc *model.VehicleIncrementRecordIncInput, set *model.VehicleIncrementRecordSetInput, Id int64) (*model1.VehicleIncrementRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleIncrementRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleIncrementRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleIncrementRecord(ctx context.Context, distinctOn []model.VehicleIncrementRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleIncrementRecordOrderBy, where *model.VehicleIncrementRecordBoolExp) ([]*model1.VehicleIncrementRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleIncrementRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleIncrementRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleIncrementRecordAggregate(ctx context.Context, distinctOn []model.VehicleIncrementRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleIncrementRecordOrderBy, where *model.VehicleIncrementRecordBoolExp) (*model.VehicleIncrementRecordAggregate, error) {
	var rs model.VehicleIncrementRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleIncrementRecord{})
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

func (r *queryResolver) VehicleIncrementRecordByPk(ctx context.Context, Id int64) (*model1.VehicleIncrementRecord, error) {
	var rs model1.VehicleIncrementRecord
	tx := db.DB.Model(&model1.VehicleIncrementRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
