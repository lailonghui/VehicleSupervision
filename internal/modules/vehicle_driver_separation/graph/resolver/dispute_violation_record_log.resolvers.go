package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_driver_separation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_driver_separation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDisputeViolationRecordLog(ctx context.Context, where model.DisputeViolationRecordLogBoolExp) (*model.DisputeViolationRecordLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DisputeViolationRecordLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DisputeViolationRecordLog
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
	return &model.DisputeViolationRecordLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDisputeViolationRecordLogByPk(ctx context.Context, Id int64) (*model1.DisputeViolationRecordLog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DisputeViolationRecordLog
	tx := db.DB.Model(&model1.DisputeViolationRecordLog{})
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

func (r *mutationResolver) InsertDisputeViolationRecordLog(ctx context.Context, objects []*model.DisputeViolationRecordLogInsertInput) (*model.DisputeViolationRecordLogMutationResponse, error) {
	rs := make([]*model1.DisputeViolationRecordLog, 0)
	for _, object := range objects {
		v := &model1.DisputeViolationRecordLog{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.DisputeViolationRecordLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DisputeViolationRecordLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDisputeViolationRecordLogOne(ctx context.Context, object model.DisputeViolationRecordLogInsertInput) (*model1.DisputeViolationRecordLog, error) {
	rs := &model1.DisputeViolationRecordLog{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.DisputeViolationRecordLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDisputeViolationRecordLog(ctx context.Context, inc *model.DisputeViolationRecordLogIncInput, set *model.DisputeViolationRecordLogSetInput, where model.DisputeViolationRecordLogBoolExp) (*model.DisputeViolationRecordLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DisputeViolationRecordLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DisputeViolationRecordLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DisputeViolationRecordLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDisputeViolationRecordLogByPk(ctx context.Context, inc *model.DisputeViolationRecordLogIncInput, set *model.DisputeViolationRecordLogSetInput, Id int64) (*model1.DisputeViolationRecordLog, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.DisputeViolationRecordLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.DisputeViolationRecordLog
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) DisputeViolationRecordLog(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) ([]*model1.DisputeViolationRecordLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DisputeViolationRecordLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DisputeViolationRecordLog
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) DisputeViolationRecordLogAggregate(ctx context.Context, distinctOn []model.DisputeViolationRecordLogSelectColumn, limit *int, offset *int, orderBy []*model.DisputeViolationRecordLogOrderBy, where *model.DisputeViolationRecordLogBoolExp) (*model.DisputeViolationRecordLogAggregate, error) {
	var rs model.DisputeViolationRecordLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DisputeViolationRecordLog{})
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

func (r *queryResolver) DisputeViolationRecordLogByPk(ctx context.Context, Id int64) (*model1.DisputeViolationRecordLog, error) {
	var rs model1.DisputeViolationRecordLog
	tx := db.DB.Model(&model1.DisputeViolationRecordLog{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
