package resolver

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/device/graph/model"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/pkg/graphql/util"
	util2 "VehicleSupervision/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteTerminalChange(ctx context.Context, where model.TerminalChangeBoolExp) (*model.TerminalChangeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalChange{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalChange
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
	return &model.TerminalChangeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalChangeByPk(ctx context.Context, Id int64) (*model1.TerminalChange, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalChange
	tx := db.DB.Model(&model1.TerminalChange{})
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

func (r *mutationResolver) InsertTerminalChange(ctx context.Context, objects []*model.TerminalChangeInsertInput) (*model.TerminalChangeMutationResponse, error) {
	rs := make([]*model1.TerminalChange, 0)
	for _, object := range objects {
		v := &model1.TerminalChange{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalChange{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalChangeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalChangeOne(ctx context.Context, object model.TerminalChangeInsertInput) (*model1.TerminalChange, error) {
	rs := &model1.TerminalChange{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalChange{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalChange(ctx context.Context, inc *model.TerminalChangeIncInput, set *model.TerminalChangeSetInput, where model.TerminalChangeBoolExp) (*model.TerminalChangeMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalChange{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalChangeMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalChange
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
	return &model.TerminalChangeMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalChangeByPk(ctx context.Context, inc *model.TerminalChangeIncInput, set *model.TerminalChangeSetInput, Id int64) (*model1.TerminalChange, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalChange{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalChange
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TerminalChange(ctx context.Context, distinctOn []model.TerminalChangeSelectColumn, limit *int, offset *int, orderBy []*model.TerminalChangeOrderBy, where *model.TerminalChangeBoolExp) ([]*model1.TerminalChange, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalChange{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalChange
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalChangeAggregate(ctx context.Context, distinctOn []model.TerminalChangeSelectColumn, limit *int, offset *int, orderBy []*model.TerminalChangeOrderBy, where *model.TerminalChangeBoolExp) (*model.TerminalChangeAggregate, error) {
	var rs model.TerminalChangeAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalChange{})
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

func (r *queryResolver) TerminalChangeByPk(ctx context.Context, Id int64) (*model1.TerminalChange, error) {
	var rs model1.TerminalChange
	tx := db.DB.Model(&model1.TerminalChange{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
