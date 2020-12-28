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

func (r *mutationResolver) DeleteVehicleReserveHistoryRecord(ctx context.Context, where model.VehicleReserveHistoryRecordBoolExp) (*model.VehicleReserveHistoryRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleReserveHistoryRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleReserveHistoryRecord
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
	return &model.VehicleReserveHistoryRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleReserveHistoryRecordByPk(ctx context.Context, Id int64) (*model1.VehicleReserveHistoryRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleReserveHistoryRecord
	tx := db.DB.Model(&model1.VehicleReserveHistoryRecord{})
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

func (r *mutationResolver) InsertVehicleReserveHistoryRecord(ctx context.Context, objects []*model.VehicleReserveHistoryRecordInsertInput) (*model.VehicleReserveHistoryRecordMutationResponse, error) {
	rs := []*model1.VehicleReserveHistoryRecord{}
	for _, object := range objects {
		v := &model1.VehicleReserveHistoryRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleReserveHistoryRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleReserveHistoryRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleReserveHistoryRecordOne(ctx context.Context, object model.VehicleReserveHistoryRecordInsertInput) (*model1.VehicleReserveHistoryRecord, error) {
	rs := &model1.VehicleReserveHistoryRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleReserveHistoryRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleReserveHistoryRecord(ctx context.Context, inc *model.VehicleReserveHistoryRecordIncInput, set *model.VehicleReserveHistoryRecordSetInput, where model.VehicleReserveHistoryRecordBoolExp) (*model.VehicleReserveHistoryRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleReserveHistoryRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleReserveHistoryRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleReserveHistoryRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleReserveHistoryRecordByPk(ctx context.Context, inc *model.VehicleReserveHistoryRecordIncInput, set *model.VehicleReserveHistoryRecordSetInput, Id int64) (*model1.VehicleReserveHistoryRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleReserveHistoryRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleReserveHistoryRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleReserveHistoryRecord(ctx context.Context, distinctOn []model.VehicleReserveHistoryRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleReserveHistoryRecordOrderBy, where *model.VehicleReserveHistoryRecordBoolExp) ([]*model1.VehicleReserveHistoryRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleReserveHistoryRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleReserveHistoryRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleReserveHistoryRecordAggregate(ctx context.Context, distinctOn []model.VehicleReserveHistoryRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleReserveHistoryRecordOrderBy, where *model.VehicleReserveHistoryRecordBoolExp) (*model.VehicleReserveHistoryRecordAggregate, error) {
	var rs model.VehicleReserveHistoryRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleReserveHistoryRecord{})
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

func (r *queryResolver) VehicleReserveHistoryRecordByPk(ctx context.Context, Id int64) (*model1.VehicleReserveHistoryRecord, error) {
	var rs model1.VehicleReserveHistoryRecord
	tx := db.DB.Model(&model1.VehicleReserveHistoryRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
