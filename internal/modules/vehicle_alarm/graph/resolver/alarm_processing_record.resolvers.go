package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_alarm/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_alarm/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteAlarmProcessingRecord(ctx context.Context, where model.AlarmProcessingRecordBoolExp) (*model.AlarmProcessingRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.AlarmProcessingRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.AlarmProcessingRecord
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
	return &model.AlarmProcessingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteAlarmProcessingRecordByPk(ctx context.Context, Id int64) (*model1.AlarmProcessingRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.AlarmProcessingRecord
	tx := db.DB.Model(&model1.AlarmProcessingRecord{})
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

func (r *mutationResolver) InsertAlarmProcessingRecord(ctx context.Context, objects []*model.AlarmProcessingRecordInsertInput) (*model.AlarmProcessingRecordMutationResponse, error) {
	rs := []*model1.AlarmProcessingRecord{}
	for _, object := range objects {
		v := &model1.AlarmProcessingRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.AlarmProcessingRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.AlarmProcessingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertAlarmProcessingRecordOne(ctx context.Context, object model.AlarmProcessingRecordInsertInput) (*model1.AlarmProcessingRecord, error) {
	rs := &model1.AlarmProcessingRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.AlarmProcessingRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateAlarmProcessingRecord(ctx context.Context, inc *model.AlarmProcessingRecordIncInput, set *model.AlarmProcessingRecordSetInput, where model.AlarmProcessingRecordBoolExp) (*model.AlarmProcessingRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.AlarmProcessingRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.AlarmProcessingRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.AlarmProcessingRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateAlarmProcessingRecordByPk(ctx context.Context, inc *model.AlarmProcessingRecordIncInput, set *model.AlarmProcessingRecordSetInput, Id int64) (*model1.AlarmProcessingRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.AlarmProcessingRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.AlarmProcessingRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) AlarmProcessingRecord(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) ([]*model1.AlarmProcessingRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.AlarmProcessingRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.AlarmProcessingRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) AlarmProcessingRecordAggregate(ctx context.Context, distinctOn []model.AlarmProcessingRecordSelectColumn, limit *int, offset *int, orderBy []*model.AlarmProcessingRecordOrderBy, where *model.AlarmProcessingRecordBoolExp) (*model.AlarmProcessingRecordAggregate, error) {
	var rs model.AlarmProcessingRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.AlarmProcessingRecord{})
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

func (r *queryResolver) AlarmProcessingRecordByPk(ctx context.Context, Id int64) (*model1.AlarmProcessingRecord, error) {
	var rs model1.AlarmProcessingRecord
	tx := db.DB.Model(&model1.AlarmProcessingRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
