package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/device/graph/model"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteTerminalTypes(ctx context.Context, where model.TerminalTypesBoolExp) (*model.TerminalTypesMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalType{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalType
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
	return &model.TerminalTypesMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalTypesByPk(ctx context.Context, id int64) (*model1.TerminalType, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalType
	tx := db.DB.Model(&model1.TerminalType{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where("id = ?", id).First(&rs)
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

func (r *mutationResolver) InsertTerminalTypes(ctx context.Context, objects []*model.TerminalTypesInsertInput, onConflict *model.TerminalTypesOnConflict) (*model.TerminalTypesMutationResponse, error) {
	rs := terminalTypeInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.TerminalType{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalTypesMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalTypesOne(ctx context.Context, object model.TerminalTypesInsertInput, onConflict *model.TerminalTypesOnConflict) (*model1.TerminalType, error) {
	rs := terminalTypeInsertInputConvert(&object)
	tx := db.DB.Model(&model1.TerminalType{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalTypes(ctx context.Context, inc *model.TerminalTypesIncInput, set *model.TerminalTypesSetInput, where model.TerminalTypesBoolExp) (*model.TerminalTypesMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalType{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalTypesMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.TerminalTypesMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateTerminalTypesByPk(ctx context.Context, inc *model.TerminalTypesIncInput, set *model.TerminalTypesSetInput, pkColumns model.TerminalTypesPkColumnsInput) (*model1.TerminalType, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.TerminalType{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.TerminalType
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) TerminalTypes(ctx context.Context, distinctOn []model.TerminalTypesSelectColumn, limit *int, offset *int, orderBy []*model.TerminalTypesOrderBy, where *model.TerminalTypesBoolExp) ([]*model1.TerminalType, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalType{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalType
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) TerminalTypesAggregate(ctx context.Context, distinctOn []model.TerminalTypesSelectColumn, limit *int, offset *int, orderBy []*model.TerminalTypesOrderBy, where *model.TerminalTypesBoolExp) (*model.TerminalTypesAggregate, error) {
	var rs model.TerminalTypesAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalType{})
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

func (r *queryResolver) TerminalTypesByPk(ctx context.Context, id int64) (*model1.TerminalType, error) {
	var rs model1.TerminalType
	tx := db.DB.Model(&model1.TerminalType{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
