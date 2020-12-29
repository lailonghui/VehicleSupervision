package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/driving/graph/model"
	model1 "VehicleSupervision/internal/modules/driving/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleNightTravelRecord(ctx context.Context, where model.VehicleNightTravelRecordBoolExp) (*model.VehicleNightTravelRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleNightTravelRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleNightTravelRecord
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
	return &model.VehicleNightTravelRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleNightTravelRecordByPk(ctx context.Context, Id int64) (*model1.VehicleNightTravelRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleNightTravelRecord
	tx := db.DB.Model(&model1.VehicleNightTravelRecord{})
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

func (r *mutationResolver) InsertVehicleNightTravelRecord(ctx context.Context, objects []*model.VehicleNightTravelRecordInsertInput) (*model.VehicleNightTravelRecordMutationResponse, error) {
	rs := make([]*model1.VehicleNightTravelRecord, 0)
	for _, object := range objects {
		v := &model1.VehicleNightTravelRecord{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleNightTravelRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleNightTravelRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleNightTravelRecordOne(ctx context.Context, object model.VehicleNightTravelRecordInsertInput) (*model1.VehicleNightTravelRecord, error) {
	rs := &model1.VehicleNightTravelRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleNightTravelRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleNightTravelRecord(ctx context.Context, inc *model.VehicleNightTravelRecordIncInput, set *model.VehicleNightTravelRecordSetInput, where model.VehicleNightTravelRecordBoolExp) (*model.VehicleNightTravelRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleNightTravelRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleNightTravelRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleNightTravelRecord
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
	return &model.VehicleNightTravelRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateVehicleNightTravelRecordByPk(ctx context.Context, inc *model.VehicleNightTravelRecordIncInput, set *model.VehicleNightTravelRecordSetInput, Id int64) (*model1.VehicleNightTravelRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleNightTravelRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleNightTravelRecord
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleNightTravelRecord(ctx context.Context, distinctOn []model.VehicleNightTravelRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleNightTravelRecordOrderBy, where *model.VehicleNightTravelRecordBoolExp) ([]*model1.VehicleNightTravelRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleNightTravelRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleNightTravelRecord
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleNightTravelRecordAggregate(ctx context.Context, distinctOn []model.VehicleNightTravelRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleNightTravelRecordOrderBy, where *model.VehicleNightTravelRecordBoolExp) (*model.VehicleNightTravelRecordAggregate, error) {
	var rs model.VehicleNightTravelRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleNightTravelRecord{})
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

func (r *queryResolver) VehicleNightTravelRecordByPk(ctx context.Context, Id int64) (*model1.VehicleNightTravelRecord, error) {
	var rs model1.VehicleNightTravelRecord
	tx := db.DB.Model(&model1.VehicleNightTravelRecord{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
