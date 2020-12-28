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

func (r *mutationResolver) DeleteVehicleExitCatalogLog(ctx context.Context, where model.VehicleExitCatalogLogBoolExp) (*model.VehicleExitCatalogLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleExitCatalogLog{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleExitCatalogLog
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
	return &model.VehicleExitCatalogLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleExitCatalogLogByPk(ctx context.Context, Id int64) (*model1.VehicleExitCatalogLog, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleExitCatalogLog
	tx := db.DB.Model(&model1.VehicleExitCatalogLog{})
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

func (r *mutationResolver) InsertVehicleExitCatalogLog(ctx context.Context, objects []*model.VehicleExitCatalogLogInsertInput) (*model.VehicleExitCatalogLogMutationResponse, error) {
	rs := make([]*model1.VehicleExitCatalogLog, 0)
	for _, object := range objects {
		v := &model1.VehicleExitCatalogLog{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleExitCatalogLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleExitCatalogLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleExitCatalogLogOne(ctx context.Context, object model.VehicleExitCatalogLogInsertInput) (*model1.VehicleExitCatalogLog, error) {
	rs := &model1.VehicleExitCatalogLog{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleExitCatalogLog{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleExitCatalogLog(ctx context.Context, inc *model.VehicleExitCatalogLogIncInput, set *model.VehicleExitCatalogLogSetInput, where model.VehicleExitCatalogLogBoolExp) (*model.VehicleExitCatalogLogMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleExitCatalogLog{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleExitCatalogLogMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleExitCatalogLogMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleExitCatalogLogByPk(ctx context.Context, inc *model.VehicleExitCatalogLogIncInput, set *model.VehicleExitCatalogLogSetInput, Id int64) (*model1.VehicleExitCatalogLog, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleExitCatalogLog{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.VehicleExitCatalogLog
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) VehicleExitCatalogLog(ctx context.Context, distinctOn []model.VehicleExitCatalogLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleExitCatalogLogOrderBy, where *model.VehicleExitCatalogLogBoolExp) ([]*model1.VehicleExitCatalogLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleExitCatalogLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleExitCatalogLog
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) VehicleExitCatalogLogAggregate(ctx context.Context, distinctOn []model.VehicleExitCatalogLogSelectColumn, limit *int, offset *int, orderBy []*model.VehicleExitCatalogLogOrderBy, where *model.VehicleExitCatalogLogBoolExp) (*model.VehicleExitCatalogLogAggregate, error) {
	var rs model.VehicleExitCatalogLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleExitCatalogLog{})
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

func (r *queryResolver) VehicleExitCatalogLogByPk(ctx context.Context, Id int64) (*model1.VehicleExitCatalogLog, error) {
	var rs model1.VehicleExitCatalogLog
	tx := db.DB.Model(&model1.VehicleExitCatalogLog{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
