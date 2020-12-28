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

func (r *mutationResolver) DeleteOperatingVehicleExt(ctx context.Context, where model.OperatingVehicleExtBoolExp) (*model.OperatingVehicleExtMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OperatingVehicleExt{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.OperatingVehicleExt
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
	return &model.OperatingVehicleExtMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteOperatingVehicleExtByPk(ctx context.Context, Id int64) (*model1.OperatingVehicleExt, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.OperatingVehicleExt
	tx := db.DB.Model(&model1.OperatingVehicleExt{})
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

func (r *mutationResolver) InsertOperatingVehicleExt(ctx context.Context, objects []*model.OperatingVehicleExtInsertInput) (*model.OperatingVehicleExtMutationResponse, error) {
	rs := []*model1.OperatingVehicleExt{}
	for _, object := range objects {
		v := &model1.OperatingVehicleExt{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.OperatingVehicleExt{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.OperatingVehicleExtMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertOperatingVehicleExtOne(ctx context.Context, object model.OperatingVehicleExtInsertInput) (*model1.OperatingVehicleExt, error) {
	rs := &model1.OperatingVehicleExt{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.OperatingVehicleExt{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateOperatingVehicleExt(ctx context.Context, inc *model.OperatingVehicleExtIncInput, set *model.OperatingVehicleExtSetInput, where model.OperatingVehicleExtBoolExp) (*model.OperatingVehicleExtMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OperatingVehicleExt{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.OperatingVehicleExtMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.OperatingVehicleExtMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateOperatingVehicleExtByPk(ctx context.Context, inc *model.OperatingVehicleExtIncInput, set *model.OperatingVehicleExtSetInput, Id int64) (*model1.OperatingVehicleExt, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.OperatingVehicleExt{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.OperatingVehicleExt
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) OperatingVehicleExt(ctx context.Context, distinctOn []model.OperatingVehicleExtSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleExtOrderBy, where *model.OperatingVehicleExtBoolExp) ([]*model1.OperatingVehicleExt, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.OperatingVehicleExt{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.OperatingVehicleExt
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) OperatingVehicleExtAggregate(ctx context.Context, distinctOn []model.OperatingVehicleExtSelectColumn, limit *int, offset *int, orderBy []*model.OperatingVehicleExtOrderBy, where *model.OperatingVehicleExtBoolExp) (*model.OperatingVehicleExtAggregate, error) {
	var rs model.OperatingVehicleExtAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.OperatingVehicleExt{})
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

func (r *queryResolver) OperatingVehicleExtByPk(ctx context.Context, Id int64) (*model1.OperatingVehicleExt, error) {
	var rs model1.OperatingVehicleExt
	tx := db.DB.Model(&model1.OperatingVehicleExt{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
