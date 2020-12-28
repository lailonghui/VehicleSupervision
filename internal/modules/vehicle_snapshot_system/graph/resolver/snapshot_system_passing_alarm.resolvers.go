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

func (r *mutationResolver) DeleteSnapshotSystemPassingAlarm(ctx context.Context, where model.SnapshotSystemPassingAlarmBoolExp) (*model.SnapshotSystemPassingAlarmMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SnapshotSystemPassingAlarm{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.SnapshotSystemPassingAlarm
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
	return &model.SnapshotSystemPassingAlarmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteSnapshotSystemPassingAlarmByPk(ctx context.Context, Id int64) (*model1.SnapshotSystemPassingAlarm, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.SnapshotSystemPassingAlarm
	tx := db.DB.Model(&model1.SnapshotSystemPassingAlarm{})
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

func (r *mutationResolver) InsertSnapshotSystemPassingAlarm(ctx context.Context, objects []*model.SnapshotSystemPassingAlarmInsertInput) (*model.SnapshotSystemPassingAlarmMutationResponse, error) {
	rs := []*model1.SnapshotSystemPassingAlarm{}
	for _, object := range objects {
		v := &model1.SnapshotSystemPassingAlarm{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.SnapshotSystemPassingAlarm{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.SnapshotSystemPassingAlarmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertSnapshotSystemPassingAlarmOne(ctx context.Context, object model.SnapshotSystemPassingAlarmInsertInput) (*model1.SnapshotSystemPassingAlarm, error) {
	rs := &model1.SnapshotSystemPassingAlarm{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.SnapshotSystemPassingAlarm{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateSnapshotSystemPassingAlarm(ctx context.Context, inc *model.SnapshotSystemPassingAlarmIncInput, set *model.SnapshotSystemPassingAlarmSetInput, where model.SnapshotSystemPassingAlarmBoolExp) (*model.SnapshotSystemPassingAlarmMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SnapshotSystemPassingAlarm{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.SnapshotSystemPassingAlarmMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.SnapshotSystemPassingAlarmMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateSnapshotSystemPassingAlarmByPk(ctx context.Context, inc *model.SnapshotSystemPassingAlarmIncInput, set *model.SnapshotSystemPassingAlarmSetInput, Id int64) (*model1.SnapshotSystemPassingAlarm, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.SnapshotSystemPassingAlarm{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.SnapshotSystemPassingAlarm
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) SnapshotSystemPassingAlarm(ctx context.Context, distinctOn []model.SnapshotSystemPassingAlarmSelectColumn, limit *int, offset *int, orderBy []*model.SnapshotSystemPassingAlarmOrderBy, where *model.SnapshotSystemPassingAlarmBoolExp) ([]*model1.SnapshotSystemPassingAlarm, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SnapshotSystemPassingAlarm{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.SnapshotSystemPassingAlarm
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) SnapshotSystemPassingAlarmAggregate(ctx context.Context, distinctOn []model.SnapshotSystemPassingAlarmSelectColumn, limit *int, offset *int, orderBy []*model.SnapshotSystemPassingAlarmOrderBy, where *model.SnapshotSystemPassingAlarmBoolExp) (*model.SnapshotSystemPassingAlarmAggregate, error) {
	var rs model.SnapshotSystemPassingAlarmAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.SnapshotSystemPassingAlarm{})
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

func (r *queryResolver) SnapshotSystemPassingAlarmByPk(ctx context.Context, Id int64) (*model1.SnapshotSystemPassingAlarm, error) {
	var rs model1.SnapshotSystemPassingAlarm
	tx := db.DB.Model(&model1.SnapshotSystemPassingAlarm{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
