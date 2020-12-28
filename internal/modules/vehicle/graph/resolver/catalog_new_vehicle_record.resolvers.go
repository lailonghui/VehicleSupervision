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

func (r *mutationResolver) DeleteCatalogNewVehicleRecord(ctx context.Context, where model.CatalogNewVehicleRecordBoolExp) (*model.CatalogNewVehicleRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.CatalogNewVehicleRecord{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.CatalogNewVehicleRecord
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
	return &model.CatalogNewVehicleRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteCatalogNewVehicleRecordByPk(ctx context.Context, Id int64) (*model1.CatalogNewVehicleRecord, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.CatalogNewVehicleRecord
	tx := db.DB.Model(&model1.CatalogNewVehicleRecord{})
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

func (r *mutationResolver) InsertCatalogNewVehicleRecord(ctx context.Context, objects []*model.CatalogNewVehicleRecordInsertInput) (*model.CatalogNewVehicleRecordMutationResponse, error) {
	rs := make([]*model1.CatalogNewVehicleRecord, 0)
	for _, object := range objects {
		v := &model1.CatalogNewVehicleRecord{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.CatalogNewVehicleRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.CatalogNewVehicleRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertCatalogNewVehicleRecordOne(ctx context.Context, object model.CatalogNewVehicleRecordInsertInput) (*model1.CatalogNewVehicleRecord, error) {
	rs := &model1.CatalogNewVehicleRecord{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.CatalogNewVehicleRecord{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateCatalogNewVehicleRecord(ctx context.Context, inc *model.CatalogNewVehicleRecordIncInput, set *model.CatalogNewVehicleRecordSetInput, where model.CatalogNewVehicleRecordBoolExp) (*model.CatalogNewVehicleRecordMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.CatalogNewVehicleRecord{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.CatalogNewVehicleRecordMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.CatalogNewVehicleRecordMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateCatalogNewVehicleRecordByPk(ctx context.Context, inc *model.CatalogNewVehicleRecordIncInput, set *model.CatalogNewVehicleRecordSetInput, Id int64) (*model1.CatalogNewVehicleRecord, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.CatalogNewVehicleRecord{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.CatalogNewVehicleRecord
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) CatalogNewVehicleRecord(ctx context.Context, distinctOn []model.CatalogNewVehicleRecordSelectColumn, limit *int, offset *int, orderBy []*model.CatalogNewVehicleRecordOrderBy, where *model.CatalogNewVehicleRecordBoolExp) ([]*model1.CatalogNewVehicleRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.CatalogNewVehicleRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.CatalogNewVehicleRecord
	tx = tx.Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) CatalogNewVehicleRecordAggregate(ctx context.Context, distinctOn []model.CatalogNewVehicleRecordSelectColumn, limit *int, offset *int, orderBy []*model.CatalogNewVehicleRecordOrderBy, where *model.CatalogNewVehicleRecordBoolExp) (*model.CatalogNewVehicleRecordAggregate, error) {
	var rs model.CatalogNewVehicleRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.CatalogNewVehicleRecord{})
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

func (r *queryResolver) CatalogNewVehicleRecordByPk(ctx context.Context, Id int64) (*model1.CatalogNewVehicleRecord, error) {
	var rs model1.CatalogNewVehicleRecord
	tx := db.DB.Model(&model1.CatalogNewVehicleRecord{}).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
