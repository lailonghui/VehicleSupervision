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

func (r *mutationResolver) DeleteTerminalFactory(ctx context.Context, where model.TerminalFactoryBoolExp) (*model.TerminalFactoryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalFactory{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalFactory
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
	return &model.TerminalFactoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalFactoryByPk(ctx context.Context, id int64) (*model1.TerminalFactory, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalFactory
	tx := db.DB.Model(&model1.TerminalFactory{})
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

func (r *mutationResolver) InsertTerminalFactory(ctx context.Context, objects []*model.TerminalFactoryInsertInput, onConflict *model.TerminalFactoryOnConflict) (*model.TerminalFactoryMutationResponse, error) {
	rs := terminalFactoryInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.TerminalFactory{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalFactoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalFactoryOne(ctx context.Context, object model.TerminalFactoryInsertInput, onConflict *model.TerminalFactoryOnConflict) (*model1.TerminalFactory, error) {
	rs := terminalFactoryInsertInputConvert(&object)
	tx := db.DB.Model(&model1.TerminalFactory{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalFactory(ctx context.Context, inc *model.TerminalFactoryIncInput, set *model.TerminalFactorySetInput, where model.TerminalFactoryBoolExp) (*model.TerminalFactoryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalFactory{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalFactoryMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.TerminalFactoryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateTerminalFactoryByPk(ctx context.Context, inc *model.TerminalFactoryIncInput, set *model.TerminalFactorySetInput, pkColumns model.TerminalFactoryPkColumnsInput) (*model1.TerminalFactory, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.TerminalFactory{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.TerminalFactory
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) TerminalFactory(ctx context.Context, distinctOn []model.TerminalFactorySelectColumn, limit *int, offset *int, orderBy []*model.TerminalFactoryOrderBy, where *model.TerminalFactoryBoolExp) ([]*model1.TerminalFactory, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalFactory{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalFactory
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) TerminalFactoryAggregate(ctx context.Context, distinctOn []model.TerminalFactorySelectColumn, limit *int, offset *int, orderBy []*model.TerminalFactoryOrderBy, where *model.TerminalFactoryBoolExp) (*model.TerminalFactoryAggregate, error) {
	var rs model.TerminalFactoryAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalFactory{})
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

func (r *queryResolver) TerminalFactoryByPk(ctx context.Context, id int64) (*model1.TerminalFactory, error) {
	var rs model1.TerminalFactory
	tx := db.DB.Model(&model1.TerminalFactory{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
