package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dynamic_supervision/graph/model"
	model1 "VehicleSupervision/internal/modules/dynamic_supervision/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleAlarmTimesRecord(ctx context.Context, where model.VehicleAlarmTimesRecordBoolExp) (*model.VehicleAlarmTimesRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmTimesRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleAlarmTimesRecord
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
	return &model.VehicleAlarmTimesRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleAlarmTimesRecordByPk(ctx context.Context, Id int64) (*model1.VehicleAlarmTimesRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleAlarmTimesRecord
	tx := db.DB.Model(&model1.VehicleAlarmTimesRecord{})
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

func (r *mutationResolver) InsertVehicleAlarmTimesRecord(ctx context.Context, objects []*model.VehicleAlarmTimesRecordInsertInput) (*model.VehicleAlarmTimesRecordMutationResponse, error) {
	rs := []*model1.VehicleAlarmTimesRecord{}
	for _, object := range objects {
		v := &model1.VehicleAlarmTimesRecord{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleAlarmTimesRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleAlarmTimesRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleAlarmTimesRecordOne(ctx context.Context, object model.VehicleAlarmTimesRecordInsertInput) (*model1.VehicleAlarmTimesRecord, error) {
	rs := &model1.VehicleAlarmTimesRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleAlarmTimesRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleAlarmTimesRecord(ctx context.Context, inc *model.VehicleAlarmTimesRecordIncInput, set *model.VehicleAlarmTimesRecordSetInput, where model.VehicleAlarmTimesRecordBoolExp) (*model.VehicleAlarmTimesRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmTimesRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleAlarmTimesRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleAlarmTimesRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleAlarmTimesRecordByPk(ctx context.Context, inc *model.VehicleAlarmTimesRecordIncInput, set *model.VehicleAlarmTimesRecordSetInput, Id int64) (*model1.VehicleAlarmTimesRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleAlarmTimesRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleAlarmTimesRecord
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleAlarmTimesRecord(ctx context.Context, distinctOn []model.VehicleAlarmTimesRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmTimesRecordOrderBy, where *model.VehicleAlarmTimesRecordBoolExp) ([]*model1.VehicleAlarmTimesRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmTimesRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleAlarmTimesRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleAlarmTimesRecordAggregate(ctx context.Context, distinctOn []model.VehicleAlarmTimesRecordSelectColumn, limit *int, offset *int, orderBy []*model.VehicleAlarmTimesRecordOrderBy, where *model.VehicleAlarmTimesRecordBoolExp) (*model.VehicleAlarmTimesRecordAggregate, error) {
	var rs model.VehicleAlarmTimesRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleAlarmTimesRecord{})
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

func (r *queryResolver) VehicleAlarmTimesRecordByPk(ctx context.Context, Id int64) (*model1.VehicleAlarmTimesRecord, error) {
	var rs model1.VehicleAlarmTimesRecord
	tx := db.DB.Model(&model1.VehicleAlarmTimesRecord{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
