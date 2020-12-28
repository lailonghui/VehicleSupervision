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

func (r *mutationResolver) DeleteVehicleEnterpriseChangeLog(ctx context.Context, where model.VehicleEnterpriseChangeLogBoolExp) (*model.VehicleEnterpriseChangeLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleEnterpriseChangeLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleEnterpriseChangeLog
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
	return &model.VehicleEnterpriseChangeLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleEnterpriseChangeLogByPk(ctx context.Context, Id int64) (*model1.VehicleEnterpriseChangeLog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleEnterpriseChangeLog
	tx := db.DB.Model(&model1.VehicleEnterpriseChangeLog{})
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

func (r *mutationResolver) InsertVehicleEnterpriseChangeLog(ctx context.Context, objects []*model.VehicleEnterpriseChangeLogInsertInput) (*model.VehicleEnterpriseChangeLogMutationResponse, error) {
	rs := []*model1.VehicleEnterpriseChangeLog{}
	for _, object := range objects {
		v := &model1.VehicleEnterpriseChangeLog{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleEnterpriseChangeLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleEnterpriseChangeLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleEnterpriseChangeLogOne(ctx context.Context, object model.VehicleEnterpriseChangeLogInsertInput) (*model1.VehicleEnterpriseChangeLog, error) {
	rs := &model1.VehicleEnterpriseChangeLog{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleEnterpriseChangeLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleEnterpriseChangeLog(ctx context.Context, inc *model.VehicleEnterpriseChangeLogIncInput, set *model.VehicleEnterpriseChangeLogSetInput, where model.VehicleEnterpriseChangeLogBoolExp) (*model.VehicleEnterpriseChangeLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleEnterpriseChangeLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleEnterpriseChangeLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleEnterpriseChangeLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleEnterpriseChangeLogByPk(ctx context.Context, inc *model.VehicleEnterpriseChangeLogIncInput, set *model.VehicleEnterpriseChangeLogSetInput, Id int64) (*model1.VehicleEnterpriseChangeLog, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleEnterpriseChangeLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleEnterpriseChangeLog
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleEnterpriseChangeLog(ctx context.Context, distinctOn []model.VehicleEnterpriseChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleEnterpriseChangeLogOrderBy, where *model.VehicleEnterpriseChangeLogBoolExp) ([]*model1.VehicleEnterpriseChangeLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleEnterpriseChangeLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleEnterpriseChangeLog
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleEnterpriseChangeLogAggregate(ctx context.Context, distinctOn []model.VehicleEnterpriseChangeLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleEnterpriseChangeLogOrderBy, where *model.VehicleEnterpriseChangeLogBoolExp) (*model.VehicleEnterpriseChangeLogAggregate, error) {
	var rs model.VehicleEnterpriseChangeLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleEnterpriseChangeLog{})
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

func (r *queryResolver) VehicleEnterpriseChangeLogByPk(ctx context.Context, Id int64) (*model1.VehicleEnterpriseChangeLog, error) {
	var rs model1.VehicleEnterpriseChangeLog
	tx := db.DB.Model(&model1.VehicleEnterpriseChangeLog{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
