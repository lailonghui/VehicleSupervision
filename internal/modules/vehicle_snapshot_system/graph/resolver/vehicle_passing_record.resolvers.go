package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_snapshot_system/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_snapshot_system/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehiclePassingRecord(ctx context.Context, where model.VehiclePassingRecordBoolExp) (*model.VehiclePassingRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehiclePassingRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehiclePassingRecord
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
	return &model.VehiclePassingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehiclePassingRecordByPk(ctx context.Context, Id int64) (*model1.VehiclePassingRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehiclePassingRecord
	tx := db.DB.Model(&model1.VehiclePassingRecord{})
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

func (r *mutationResolver) InsertVehiclePassingRecord(ctx context.Context, objects []*model.VehiclePassingRecordInsertInput) (*model.VehiclePassingRecordMutationResponse, error) {
	rs := []*model1.VehiclePassingRecord{}
	for _, object := range objects {
		v := &model1.VehiclePassingRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehiclePassingRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehiclePassingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehiclePassingRecordOne(ctx context.Context, object model.VehiclePassingRecordInsertInput) (*model1.VehiclePassingRecord, error) {
	rs := &model1.VehiclePassingRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehiclePassingRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehiclePassingRecord(ctx context.Context, inc *model.VehiclePassingRecordIncInput, set *model.VehiclePassingRecordSetInput, where model.VehiclePassingRecordBoolExp) (*model.VehiclePassingRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehiclePassingRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehiclePassingRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehiclePassingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehiclePassingRecordByPk(ctx context.Context, inc *model.VehiclePassingRecordIncInput, set *model.VehiclePassingRecordSetInput, Id int64) (*model1.VehiclePassingRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehiclePassingRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehiclePassingRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehiclePassingRecord(ctx context.Context, distinctOn []model.VehiclePassingRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehiclePassingRecordOrderBy, where *model.VehiclePassingRecordBoolExp) ([]*model1.VehiclePassingRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehiclePassingRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehiclePassingRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehiclePassingRecordAggregate(ctx context.Context, distinctOn []model.VehiclePassingRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehiclePassingRecordOrderBy, where *model.VehiclePassingRecordBoolExp) (*model.VehiclePassingRecordAggregate, error) {
	var rs model.VehiclePassingRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehiclePassingRecord{})
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

func (r *queryResolver) VehiclePassingRecordByPk(ctx context.Context, Id int64) (*model1.VehiclePassingRecord, error) {
	var rs model1.VehiclePassingRecord
	tx := db.DB.Model(&model1.VehiclePassingRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
