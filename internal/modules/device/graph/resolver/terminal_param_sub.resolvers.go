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

func (r *mutationResolver) DeleteTerminalParamSub(ctx context.Context, where model.TerminalParamSubBoolExp) (*model.TerminalParamSubMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamSub{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalParamSub
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
	return &model.TerminalParamSubMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalParamSubByPk(ctx context.Context, Id int64) (*model1.TerminalParamSub, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalParamSub
	tx := db.DB.Model(&model1.TerminalParamSub{})
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

func (r *mutationResolver) InsertTerminalParamSub(ctx context.Context, objects []*model.TerminalParamSubInsertInput) (*model.TerminalParamSubMutationResponse, error) {
	rs := make([]*model1.TerminalParamSub, 0)
	for _, object := range objects {
		v := &model1.TerminalParamSub{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalParamSub{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalParamSubMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalParamSubOne(ctx context.Context, object model.TerminalParamSubInsertInput) (*model1.TerminalParamSub, error) {
	rs := &model1.TerminalParamSub{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalParamSub{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalParamSub(ctx context.Context, inc *model.TerminalParamSubIncInput, set *model.TerminalParamSubSetInput, where model.TerminalParamSubBoolExp) (*model.TerminalParamSubMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamSub{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalParamSubMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalParamSub
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
	return &model.TerminalParamSubMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalParamSubByPk(ctx context.Context, inc *model.TerminalParamSubIncInput, set *model.TerminalParamSubSetInput, Id int64) (*model1.TerminalParamSub, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalParamSub{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalParamSub
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TerminalParamSub(ctx context.Context, distinctOn []model.TerminalParamSubSelectColumn, limit *int, offset *int, orderBy []*model.TerminalParamSubOrderBy, where *model.TerminalParamSubBoolExp) ([]*model1.TerminalParamSub, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamSub{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalParamSub
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalParamSubAggregate(ctx context.Context, distinctOn []model.TerminalParamSubSelectColumn, limit *int, offset *int, orderBy []*model.TerminalParamSubOrderBy, where *model.TerminalParamSubBoolExp) (*model.TerminalParamSubAggregate, error) {
	var rs model.TerminalParamSubAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamSub{})
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

func (r *queryResolver) TerminalParamSubByPk(ctx context.Context, Id int64) (*model1.TerminalParamSub, error) {
	var rs model1.TerminalParamSub
	tx := db.DB.Model(&model1.TerminalParamSub{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
