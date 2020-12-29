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

func (r *mutationResolver) DeleteVehicleStopTravelRecord(ctx context.Context, where model.VehicleStopTravelRecordBoolExp) (*model.VehicleStopTravelRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStopTravelRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleStopTravelRecord
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
	return &model.VehicleStopTravelRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleStopTravelRecordByPk(ctx context.Context, Id int64) (*model1.VehicleStopTravelRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleStopTravelRecord
	tx := db.DB.Model(&model1.VehicleStopTravelRecord{})
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

func (r *mutationResolver) InsertVehicleStopTravelRecord(ctx context.Context, objects []*model.VehicleStopTravelRecordInsertInput) (*model.VehicleStopTravelRecordMutationResponse, error) {
	rs := make([]*model1.VehicleStopTravelRecord, 0)
	for _, object := range objects {
		v := &model1.VehicleStopTravelRecord{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleStopTravelRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleStopTravelRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleStopTravelRecordOne(ctx context.Context, object model.VehicleStopTravelRecordInsertInput) (*model1.VehicleStopTravelRecord, error) {
	rs := &model1.VehicleStopTravelRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleStopTravelRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleStopTravelRecord(ctx context.Context, inc *model.VehicleStopTravelRecordIncInput, set *model.VehicleStopTravelRecordSetInput, where model.VehicleStopTravelRecordBoolExp) (*model.VehicleStopTravelRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStopTravelRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleStopTravelRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleStopTravelRecord
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
	return &model.VehicleStopTravelRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateVehicleStopTravelRecordByPk(ctx context.Context, inc *model.VehicleStopTravelRecordIncInput, set *model.VehicleStopTravelRecordSetInput, Id int64) (*model1.VehicleStopTravelRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleStopTravelRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleStopTravelRecord
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleStopTravelRecord(ctx context.Context, distinctOn []model.VehicleStopTravelRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleStopTravelRecordOrderBy, where *model.VehicleStopTravelRecordBoolExp) ([]*model1.VehicleStopTravelRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStopTravelRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleStopTravelRecord
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleStopTravelRecordAggregate(ctx context.Context, distinctOn []model.VehicleStopTravelRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleStopTravelRecordOrderBy, where *model.VehicleStopTravelRecordBoolExp) (*model.VehicleStopTravelRecordAggregate, error) {
	var rs model.VehicleStopTravelRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStopTravelRecord{})
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

func (r *queryResolver) VehicleStopTravelRecordByPk(ctx context.Context, Id int64) (*model1.VehicleStopTravelRecord, error) {
	var rs model1.VehicleStopTravelRecord
	tx := db.DB.Model(&model1.VehicleStopTravelRecord{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
