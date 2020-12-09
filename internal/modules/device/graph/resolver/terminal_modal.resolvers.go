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

func (r *mutationResolver) DeleteTerminalModal(ctx context.Context, where model.TerminalModalBoolExp) (*model.TerminalModalMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalModal{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalModal
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
	return &model.TerminalModalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalModalByPk(ctx context.Context, id int64) (*model1.TerminalModal, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalModal
	tx := db.DB.Model(&model1.TerminalModal{})
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

func (r *mutationResolver) InsertTerminalModal(ctx context.Context, objects []*model.TerminalModalInsertInput, onConflict *model.TerminalModalOnConflict) (*model.TerminalModalMutationResponse, error) {
	rs := terminalModalInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.TerminalModal{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalModalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalModalOne(ctx context.Context, object model.TerminalModalInsertInput, onConflict *model.TerminalModalOnConflict) (*model1.TerminalModal, error) {
	rs := terminalModalInsertInputConvert(&object)
	tx := db.DB.Model(&model1.TerminalModal{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalModal(ctx context.Context, inc *model.TerminalModalIncInput, set *model.TerminalModalSetInput, where model.TerminalModalBoolExp) (*model.TerminalModalMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalModal{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalModalMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.TerminalModalMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateTerminalModalByPk(ctx context.Context, inc *model.TerminalModalIncInput, set *model.TerminalModalSetInput, pkColumns model.TerminalModalPkColumnsInput) (*model1.TerminalModal, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.TerminalModal{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.TerminalModal
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) TerminalModal(ctx context.Context, distinctOn []model.TerminalModalSelectColumn, limit *int, offset *int, orderBy []*model.TerminalModalOrderBy, where *model.TerminalModalBoolExp) ([]*model1.TerminalModal, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalModal{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalModal
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) TerminalModalAggregate(ctx context.Context, distinctOn []model.TerminalModalSelectColumn, limit *int, offset *int, orderBy []*model.TerminalModalOrderBy, where *model.TerminalModalBoolExp) (*model.TerminalModalAggregate, error) {
	var rs model.TerminalModalAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalModal{})
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

func (r *queryResolver) TerminalModalByPk(ctx context.Context, id int64) (*model1.TerminalModal, error) {
	var rs model1.TerminalModal
	tx := db.DB.Model(&model1.TerminalModal{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
