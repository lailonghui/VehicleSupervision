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

func (r *mutationResolver) DeleteVehicleOperationHistory(ctx context.Context, where model.VehicleOperationHistoryBoolExp) (*model.VehicleOperationHistoryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOperationHistory{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleOperationHistory
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
	return &model.VehicleOperationHistoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleOperationHistoryByPk(ctx context.Context, Id int64) (*model1.VehicleOperationHistory, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleOperationHistory
	tx := db.DB.Model(&model1.VehicleOperationHistory{})
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

func (r *mutationResolver) InsertVehicleOperationHistory(ctx context.Context, objects []*model.VehicleOperationHistoryInsertInput) (*model.VehicleOperationHistoryMutationResponse, error) {
	rs := []*model1.VehicleOperationHistory{}
	for _, object := range objects {
		v := &model1.VehicleOperationHistory{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleOperationHistory{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleOperationHistoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleOperationHistoryOne(ctx context.Context, object model.VehicleOperationHistoryInsertInput) (*model1.VehicleOperationHistory, error) {
	rs := &model1.VehicleOperationHistory{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleOperationHistory{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleOperationHistory(ctx context.Context, inc *model.VehicleOperationHistoryIncInput, set *model.VehicleOperationHistorySetInput, where model.VehicleOperationHistoryBoolExp) (*model.VehicleOperationHistoryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOperationHistory{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleOperationHistoryMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleOperationHistoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleOperationHistoryByPk(ctx context.Context, inc *model.VehicleOperationHistoryIncInput, set *model.VehicleOperationHistorySetInput, Id int64) (*model1.VehicleOperationHistory, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleOperationHistory{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleOperationHistory
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleOperationHistory(ctx context.Context, distinctOn []model.VehicleOperationHistorySelectColumn, limit *int, offset *int, orderBy []*model.VehicleOperationHistoryOrderBy, where *model.VehicleOperationHistoryBoolExp) ([]*model1.VehicleOperationHistory, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOperationHistory{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleOperationHistory
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleOperationHistoryAggregate(ctx context.Context, distinctOn []model.VehicleOperationHistorySelectColumn, limit *int, offset *int, orderBy []*model.VehicleOperationHistoryOrderBy, where *model.VehicleOperationHistoryBoolExp) (*model.VehicleOperationHistoryAggregate, error) {
	var rs model.VehicleOperationHistoryAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleOperationHistory{})
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

func (r *queryResolver) VehicleOperationHistoryByPk(ctx context.Context, Id int64) (*model1.VehicleOperationHistory, error) {
	var rs model1.VehicleOperationHistory
	tx := db.DB.Model(&model1.VehicleOperationHistory{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
