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

func (r *mutationResolver) DeleteTerminal(ctx context.Context, where model.TerminalBoolExp) (*model.TerminalMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Terminal{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.Terminal
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
	return &model.TerminalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalByPk(ctx context.Context, id int64) (*model1.Terminal, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.Terminal
	tx := db.DB.Model(&model1.Terminal{})
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

func (r *mutationResolver) InsertTerminal(ctx context.Context, objects []*model.TerminalInsertInput, onConflict *model.TerminalOnConflict) (*model.TerminalMutationResponse, error) {
	rs := terminalInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.Terminal{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalOne(ctx context.Context, object model.TerminalInsertInput, onConflict *model.TerminalOnConflict) (*model1.Terminal, error) {
	rs := terminalInsertInputConvert(&object)
	tx := db.DB.Model(&model1.Terminal{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminal(ctx context.Context, inc *model.TerminalIncInput, set *model.TerminalSetInput, where model.TerminalBoolExp) (*model.TerminalMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Terminal{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.TerminalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateTerminalByPk(ctx context.Context, inc *model.TerminalIncInput, set *model.TerminalSetInput, pkColumns model.TerminalPkColumnsInput) (*model1.Terminal, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.Terminal{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.Terminal
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) Terminal(ctx context.Context, distinctOn []model.TerminalSelectColumn, limit *int, offset *int, orderBy []*model.TerminalOrderBy, where *model.TerminalBoolExp) ([]*model1.Terminal, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Terminal{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.Terminal
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) TerminalAggregate(ctx context.Context, distinctOn []model.TerminalSelectColumn, limit *int, offset *int, orderBy []*model.TerminalOrderBy, where *model.TerminalBoolExp) (*model.TerminalAggregate, error) {
	var rs model.TerminalAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.Terminal{})
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

func (r *queryResolver) TerminalByPk(ctx context.Context, id int64) (*model1.Terminal, error) {
	var rs model1.Terminal
	tx := db.DB.Model(&model1.Terminal{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
