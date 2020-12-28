package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/muck_truck_recommend_catalog/graph/model"
	model1 "VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleSaleRecord(ctx context.Context, where model.VehicleSaleRecordBoolExp) (*model.VehicleSaleRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSaleRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleSaleRecord
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
	return &model.VehicleSaleRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleSaleRecordByPk(ctx context.Context, Id int64) (*model1.VehicleSaleRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleSaleRecord
	tx := db.DB.Model(&model1.VehicleSaleRecord{})
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

func (r *mutationResolver) InsertVehicleSaleRecord(ctx context.Context, objects []*model.VehicleSaleRecordInsertInput) (*model.VehicleSaleRecordMutationResponse, error) {
	rs := []*model1.VehicleSaleRecord{}
	for _, object := range objects {
		v := &model1.VehicleSaleRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleSaleRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleSaleRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleSaleRecordOne(ctx context.Context, object model.VehicleSaleRecordInsertInput) (*model1.VehicleSaleRecord, error) {
	rs := &model1.VehicleSaleRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleSaleRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleSaleRecord(ctx context.Context, inc *model.VehicleSaleRecordIncInput, set *model.VehicleSaleRecordSetInput, where model.VehicleSaleRecordBoolExp) (*model.VehicleSaleRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSaleRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleSaleRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleSaleRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleSaleRecordByPk(ctx context.Context, inc *model.VehicleSaleRecordIncInput, set *model.VehicleSaleRecordSetInput, Id int64) (*model1.VehicleSaleRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleSaleRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleSaleRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleSaleRecord(ctx context.Context, distinctOn []model.VehicleSaleRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSaleRecordOrderBy, where *model.VehicleSaleRecordBoolExp) ([]*model1.VehicleSaleRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSaleRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleSaleRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleSaleRecordAggregate(ctx context.Context, distinctOn []model.VehicleSaleRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleSaleRecordOrderBy, where *model.VehicleSaleRecordBoolExp) (*model.VehicleSaleRecordAggregate, error) {
	var rs model.VehicleSaleRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleSaleRecord{})
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

func (r *queryResolver) VehicleSaleRecordByPk(ctx context.Context, Id int64) (*model1.VehicleSaleRecord, error) {
	var rs model1.VehicleSaleRecord
	tx := db.DB.Model(&model1.VehicleSaleRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
