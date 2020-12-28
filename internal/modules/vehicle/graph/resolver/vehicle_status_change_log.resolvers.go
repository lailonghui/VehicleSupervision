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

func (r *mutationResolver) DeleteVehicleStatusChangeLog(ctx context.Context, where model.VehicleStatusChangeLogBoolExp) (*model.VehicleStatusChangeLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStatusChangeLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleStatusChangeLog
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
	return &model.VehicleStatusChangeLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleStatusChangeLogByPk(ctx context.Context, Id int64) (*model1.VehicleStatusChangeLog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleStatusChangeLog
	tx := db.DB.Model(&model1.VehicleStatusChangeLog{})
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

func (r *mutationResolver) InsertVehicleStatusChangeLog(ctx context.Context, objects []*model.VehicleStatusChangeLogInsertInput) (*model.VehicleStatusChangeLogMutationResponse, error) {
	rs := []*model1.VehicleStatusChangeLog{}
	for _, object := range objects {
		v := &model1.VehicleStatusChangeLog{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleStatusChangeLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleStatusChangeLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleStatusChangeLogOne(ctx context.Context, object model.VehicleStatusChangeLogInsertInput) (*model1.VehicleStatusChangeLog, error) {
	rs := &model1.VehicleStatusChangeLog{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleStatusChangeLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleStatusChangeLog(ctx context.Context, inc *model.VehicleStatusChangeLogIncInput, set *model.VehicleStatusChangeLogSetInput, where model.VehicleStatusChangeLogBoolExp) (*model.VehicleStatusChangeLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStatusChangeLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleStatusChangeLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleStatusChangeLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleStatusChangeLogByPk(ctx context.Context, inc *model.VehicleStatusChangeLogIncInput, set *model.VehicleStatusChangeLogSetInput, Id int64) (*model1.VehicleStatusChangeLog, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleStatusChangeLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleStatusChangeLog
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleStatusChangeLog(ctx context.Context, distinctOn []model.VehicleStatusChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleStatusChangeLogOrderBy, where *model.VehicleStatusChangeLogBoolExp) ([]*model1.VehicleStatusChangeLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStatusChangeLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleStatusChangeLog
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleStatusChangeLogAggregate(ctx context.Context, distinctOn []model.VehicleStatusChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleStatusChangeLogOrderBy, where *model.VehicleStatusChangeLogBoolExp) (*model.VehicleStatusChangeLogAggregate, error) {
	var rs model.VehicleStatusChangeLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleStatusChangeLog{})
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

func (r *queryResolver) VehicleStatusChangeLogByPk(ctx context.Context, Id int64) (*model1.VehicleStatusChangeLog, error) {
	var rs model1.VehicleStatusChangeLog
	tx := db.DB.Model(&model1.VehicleStatusChangeLog{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
