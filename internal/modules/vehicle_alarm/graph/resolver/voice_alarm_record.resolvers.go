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

func (r *mutationResolver) DeleteVoiceAlarmRecord(ctx context.Context, where model.VoiceAlarmRecordBoolExp) (*model.VoiceAlarmRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VoiceAlarmRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VoiceAlarmRecord
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
	return &model.VoiceAlarmRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVoiceAlarmRecordByPk(ctx context.Context, Id int64) (*model1.VoiceAlarmRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VoiceAlarmRecord
	tx := db.DB.Model(&model1.VoiceAlarmRecord{})
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

func (r *mutationResolver) InsertVoiceAlarmRecord(ctx context.Context, objects []*model.VoiceAlarmRecordInsertInput) (*model.VoiceAlarmRecordMutationResponse, error) {
	rs := []*model1.VoiceAlarmRecord{}
	for _, object := range objects {
		v := &model1.VoiceAlarmRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VoiceAlarmRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VoiceAlarmRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVoiceAlarmRecordOne(ctx context.Context, object model.VoiceAlarmRecordInsertInput) (*model1.VoiceAlarmRecord, error) {
	rs := &model1.VoiceAlarmRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VoiceAlarmRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVoiceAlarmRecord(ctx context.Context, inc *model.VoiceAlarmRecordIncInput, set *model.VoiceAlarmRecordSetInput, where model.VoiceAlarmRecordBoolExp) (*model.VoiceAlarmRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VoiceAlarmRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VoiceAlarmRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VoiceAlarmRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVoiceAlarmRecordByPk(ctx context.Context, inc *model.VoiceAlarmRecordIncInput, set *model.VoiceAlarmRecordSetInput, Id int64) (*model1.VoiceAlarmRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VoiceAlarmRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VoiceAlarmRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VoiceAlarmRecord(ctx context.Context, distinctOn []model.VoiceAlarmRecordSelectColumn, limit *int, offset *int, orderBy []*model.VoiceAlarmRecordOrderBy, where *model.VoiceAlarmRecordBoolExp) ([]*model1.VoiceAlarmRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VoiceAlarmRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VoiceAlarmRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VoiceAlarmRecordAggregate(ctx context.Context, distinctOn []model.VoiceAlarmRecordSelectColumn, limit *int, offset *int, orderBy []*model.VoiceAlarmRecordOrderBy, where *model.VoiceAlarmRecordBoolExp) (*model.VoiceAlarmRecordAggregate, error) {
	var rs model.VoiceAlarmRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VoiceAlarmRecord{})
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

func (r *queryResolver) VoiceAlarmRecordByPk(ctx context.Context, Id int64) (*model1.VoiceAlarmRecord, error) {
	var rs model1.VoiceAlarmRecord
	tx := db.DB.Model(&model1.VoiceAlarmRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
