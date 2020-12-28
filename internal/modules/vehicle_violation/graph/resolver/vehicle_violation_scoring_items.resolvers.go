package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/vehicle_violation/graph/model"
	model1 "VehicleSupervision/internal/modules/vehicle_violation/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteVehicleViolationScoringItems(ctx context.Context, where model.VehicleViolationScoringItemsBoolExp) (*model.VehicleViolationScoringItemsMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationScoringItems{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.VehicleViolationScoringItems
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
	return &model.VehicleViolationScoringItemsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteVehicleViolationScoringItemsByPk(ctx context.Context, Id int64) (*model1.VehicleViolationScoringItems, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.VehicleViolationScoringItems
	tx := db.DB.Model(&model1.VehicleViolationScoringItems{})
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

func (r *mutationResolver) InsertVehicleViolationScoringItems(ctx context.Context, objects []*model.VehicleViolationScoringItemsInsertInput) (*model.VehicleViolationScoringItemsMutationResponse, error) {
	rs := []*model1.VehicleViolationScoringItems{}
	for _, object := range objects {
		v := &model1.VehicleViolationScoringItems{}
		util2.StructAssign(v, &object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.VehicleViolationScoringItems{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.VehicleViolationScoringItemsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertVehicleViolationScoringItemsOne(ctx context.Context, object model.VehicleViolationScoringItemsInsertInput) (*model1.VehicleViolationScoringItems, error) {
	rs := &model1.VehicleViolationScoringItems{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.VehicleViolationScoringItems{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateVehicleViolationScoringItems(ctx context.Context, inc *model.VehicleViolationScoringItemsIncInput, set *model.VehicleViolationScoringItemsSetInput, where model.VehicleViolationScoringItemsBoolExp) (*model.VehicleViolationScoringItemsMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationScoringItems{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.VehicleViolationScoringItemsMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.VehicleViolationScoringItemsMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateVehicleViolationScoringItemsByPk(ctx context.Context, inc *model.VehicleViolationScoringItemsIncInput, set *model.VehicleViolationScoringItemsSetInput, Id int64) (*model1.VehicleViolationScoringItems, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.VehicleViolationScoringItems{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.VehicleViolationScoringItems
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) VehicleViolationScoringItems(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) ([]*model1.VehicleViolationScoringItems, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationScoringItems{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleViolationScoringItems
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleViolationScoringItemsAggregate(ctx context.Context, distinctOn []model.VehicleViolationScoringItemsSelectColumn, limit *int, offset *int, orderBy []*model.VehicleViolationScoringItemsOrderBy, where *model.VehicleViolationScoringItemsBoolExp) (*model.VehicleViolationScoringItemsAggregate, error) {
	var rs model.VehicleViolationScoringItemsAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleViolationScoringItems{})
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

func (r *queryResolver) VehicleViolationScoringItemsByPk(ctx context.Context, Id int64) (*model1.VehicleViolationScoringItems, error) {
	var rs model1.VehicleViolationScoringItems
	tx := db.DB.Model(&model1.VehicleViolationScoringItems{}).First(&rs, Id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
