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

func (r *mutationResolver) DeleteTerminalParamConfig(ctx context.Context, where model.TerminalParamConfigBoolExp) (*model.TerminalParamConfigMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamConfig{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalParamConfig
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
	return &model.TerminalParamConfigMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteTerminalParamConfigByPk(ctx context.Context, Id int64) (*model1.TerminalParamConfig, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalParamConfig
	tx := db.DB.Model(&model1.TerminalParamConfig{})
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

func (r *mutationResolver) DeleteTerminalParamConfigByUnionPk(ctx context.Context, unionId string) (*model1.TerminalParamConfig, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.TerminalParamConfig
	tx := db.DB.Model(&model1.TerminalParamConfig{})
	if len(preloads) > 0 {
		// 如果请求的字段不为空，则先查询一遍数据库
		tx = tx.Select(preloads).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)
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

func (r *mutationResolver) InsertTerminalParamConfig(ctx context.Context, objects []*model.TerminalParamConfigInsertInput) (*model.TerminalParamConfigMutationResponse, error) {
	rs := make([]*model1.TerminalParamConfig, 0)
	for _, object := range objects {
		v := &model1.TerminalParamConfig{}
		util2.StructAssign(v, object)
		rs = append(rs, v)
	}
	tx := db.DB.Model(&model1.TerminalParamConfig{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.TerminalParamConfigMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertTerminalParamConfigOne(ctx context.Context, object model.TerminalParamConfigInsertInput) (*model1.TerminalParamConfig, error) {
	rs := &model1.TerminalParamConfig{}
	util2.StructAssign(rs, &object)
	tx := db.DB.Model(&model1.TerminalParamConfig{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateTerminalParamConfig(ctx context.Context, inc *model.TerminalParamConfigIncInput, set *model.TerminalParamConfigSetInput, where model.TerminalParamConfigBoolExp) (*model.TerminalParamConfigMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamConfig{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.TerminalParamConfigMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.TerminalParamConfig
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
	return &model.TerminalParamConfigMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) UpdateTerminalParamConfigByPk(ctx context.Context, inc *model.TerminalParamConfigIncInput, set *model.TerminalParamConfigSetInput, Id int64) (*model1.TerminalParamConfig, error) {
	tx := db.DB.Where("id = ?", Id)
	qt := util.NewQueryTranslator(tx, &model1.TerminalParamConfig{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}
	var rs model1.TerminalParamConfig
	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *mutationResolver) UpdateTerminalParamConfigByUnionPk(ctx context.Context, inc *model.TerminalParamConfigIncInput, set *model.TerminalParamConfigSetInput, unionId string) (*model1.TerminalParamConfig, error) {
	var rs model1.TerminalParamConfig
	tx := db.DB.Where(rs.UnionPrimaryColumnName()+" = ?", unionId)
	qt := util.NewQueryTranslator(tx, &model1.TerminalParamConfig{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		return nil, err
	}

	tx = tx.First(&rs)
	if err := tx.Error; err != nil {
		return &rs, err
	}
	return &rs, nil
}

func (r *queryResolver) TerminalParamConfig(ctx context.Context, distinctOn []model.TerminalParamConfigSelectColumn, limit *int, offset *int, orderBy []*model.TerminalParamConfigOrderBy, where *model.TerminalParamConfigBoolExp) ([]*model1.TerminalParamConfig, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamConfig{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.TerminalParamConfig
	tx = tx.Select(util.GetTopPreloads(ctx)).Find(&rs)
	err := tx.Error
	return rs, err
}

func (r *queryResolver) TerminalParamConfigAggregate(ctx context.Context, distinctOn []model.TerminalParamConfigSelectColumn, limit *int, offset *int, orderBy []*model.TerminalParamConfigOrderBy, where *model.TerminalParamConfigBoolExp) (*model.TerminalParamConfigAggregate, error) {
	var rs model.TerminalParamConfigAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.TerminalParamConfig{})
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

func (r *queryResolver) TerminalParamConfigByPk(ctx context.Context, Id int64) (*model1.TerminalParamConfig, error) {
	var rs model1.TerminalParamConfig
	tx := db.DB.Model(&model1.TerminalParamConfig{}).Select(util.GetTopPreloads(ctx)).First(&rs, Id)
	err := tx.Error
	return &rs, err
}

func (r *queryResolver) TerminalParamConfigByUnionPk(ctx context.Context, unionId string) (*model1.TerminalParamConfig, error) {
	var rs model1.TerminalParamConfig
	tx := db.DB.Model(&model1.TerminalParamConfig{}).Select(util.GetTopPreloads(ctx)).Where(rs.UnionPrimaryColumnName()+" = ?", unionId).First(&rs)

	err := tx.Error
	return &rs, err
}
