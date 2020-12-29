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

func (r *mutationResolver) DeleteTerminalParamItem(ctx context.Context, where model.TerminalParamItemBoolExp) (*model.TerminalParamItemMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamItem{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalParamItem
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
	return &model.TerminalParamItemMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalParamItemByPk(ctx context.Context, Id int64) (*model1.TerminalParamItem, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalParamItem
	tx := db.DB.Model(&model1.TerminalParamItem{})
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

func (r *mutationResolver) InsertTerminalParamItem(ctx context.Context, objects []*model.TerminalParamItemInsertInput) (*model.TerminalParamItemMutationResponse, error) {
	rs := make([]*model1.TerminalParamItem, 0)
	for _, object := range objects {
		v := &model1.TerminalParamItem{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalParamItem{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalParamItemMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalParamItemOne(ctx context.Context, object model.TerminalParamItemInsertInput) (*model1.TerminalParamItem, error) {
	rs := &model1.TerminalParamItem{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalParamItem{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalParamItem(ctx context.Context, inc *model.TerminalParamItemIncInput, set *model.TerminalParamItemSetInput, where model.TerminalParamItemBoolExp) (*model.TerminalParamItemMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamItem{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalParamItemMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalParamItem
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
	return &model.TerminalParamItemMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalParamItemByPk(ctx context.Context, inc *model.TerminalParamItemIncInput, set *model.TerminalParamItemSetInput, Id int64) (*model1.TerminalParamItem, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalParamItem{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalParamItem
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TerminalParamItem(ctx context.Context, distinctOn []model.TerminalParamItemSelectColumn, limit *int, offset *int, orderBy []*model.TerminalParamItemOrderBy, where *model.TerminalParamItemBoolExp) ([]*model1.TerminalParamItem, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamItem{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalParamItem
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalParamItemAggregate(ctx context.Context, distinctOn []model.TerminalParamItemSelectColumn, limit *int, offset *int, orderBy []*model.TerminalParamItemOrderBy, where *model.TerminalParamItemBoolExp) (*model.TerminalParamItemAggregate, error) {
	var rs model.TerminalParamItemAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamItem{})
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

func (r *queryResolver) TerminalParamItemByPk(ctx context.Context, Id int64) (*model1.TerminalParamItem, error) {
	var rs model1.TerminalParamItem
	tx := db.DB.Model(&model1.TerminalParamItem{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
