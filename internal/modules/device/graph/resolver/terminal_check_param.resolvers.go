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

func (r *mutationResolver) DeleteTerminalCheckParam(ctx context.Context, where model.TerminalCheckParamBoolExp) (*model.TerminalCheckParamMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalCheckParam{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalCheckParam
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
	return &model.TerminalCheckParamMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalCheckParamByPk(ctx context.Context, Id int64) (*model1.TerminalCheckParam, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalCheckParam
	tx := db.DB.Model(&model1.TerminalCheckParam{})
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

func (r *mutationResolver) InsertTerminalCheckParam(ctx context.Context, objects []*model.TerminalCheckParamInsertInput) (*model.TerminalCheckParamMutationResponse, error) {
	rs := make([]*model1.TerminalCheckParam, 0)
	for _, object := range objects {
		v := &model1.TerminalCheckParam{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalCheckParam{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalCheckParamMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalCheckParamOne(ctx context.Context, object model.TerminalCheckParamInsertInput) (*model1.TerminalCheckParam, error) {
	rs := &model1.TerminalCheckParam{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalCheckParam{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalCheckParam(ctx context.Context, inc *model.TerminalCheckParamIncInput, set *model.TerminalCheckParamSetInput, where model.TerminalCheckParamBoolExp) (*model.TerminalCheckParamMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalCheckParam{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalCheckParamMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalCheckParam
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
	return &model.TerminalCheckParamMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalCheckParamByPk(ctx context.Context, inc *model.TerminalCheckParamIncInput, set *model.TerminalCheckParamSetInput, Id int64) (*model1.TerminalCheckParam, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalCheckParam{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalCheckParam
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TerminalCheckParam(ctx context.Context, distinctOn []model.TerminalCheckParamSelectColumn, limit *int, offset *int, orderBy []*model.TerminalCheckParamOrderBy, where *model.TerminalCheckParamBoolExp) ([]*model1.TerminalCheckParam, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalCheckParam{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalCheckParam
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalCheckParamAggregate(ctx context.Context, distinctOn []model.TerminalCheckParamSelectColumn, limit *int, offset *int, orderBy []*model.TerminalCheckParamOrderBy, where *model.TerminalCheckParamBoolExp) (*model.TerminalCheckParamAggregate, error) {
	var rs model.TerminalCheckParamAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalCheckParam{})
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

func (r *queryResolver) TerminalCheckParamByPk(ctx context.Context, Id int64) (*model1.TerminalCheckParam, error) {
	var rs model1.TerminalCheckParam
	tx := db.DB.Model(&model1.TerminalCheckParam{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}
