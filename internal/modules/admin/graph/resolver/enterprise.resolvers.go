package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/admin/graph/model"
	model1 "VehicleSupervision/internal/modules/admin/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteEnterprise(ctx context.Context, where model.EnterpriseBoolExp) (*model.EnterpriseMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Enterprise{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.Enterprise
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
	return &model.EnterpriseMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteEnterpriseByPk(ctx context.Context, id int64) (*model1.Enterprise, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.Enterprise
	tx := db.DB.Model(&model1.Enterprise{})
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

func (r *mutationResolver) InsertEnterprise(ctx context.Context, objects []*model.EnterpriseInsertInput, onConflict *model.EnterpriseOnConflict) (*model.EnterpriseMutationResponse, error) {
	rs := r.enterpriseInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.Enterprise{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.EnterpriseMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertEnterpriseOne(ctx context.Context, object model.EnterpriseInsertInput, onConflict *model.EnterpriseOnConflict) (*model1.Enterprise, error) {
	rs := r.enterpriseInsertInputConvert(&object)
	tx := db.DB.Model(&model1.Enterprise{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateEnterprise(ctx context.Context, inc *model.EnterpriseIncInput, set *model.EnterpriseSetInput, where model.EnterpriseBoolExp) (*model.EnterpriseMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Enterprise{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.EnterpriseMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.EnterpriseMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateEnterpriseByPk(ctx context.Context, inc *model.EnterpriseIncInput, set *model.EnterpriseSetInput, pkColumns model.EnterprisePkColumnsInput) (*model1.Enterprise, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.Enterprise{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.Enterprise
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) Enterprise(ctx context.Context, distinctOn []model.EnterpriseSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseOrderBy, where *model.EnterpriseBoolExp) ([]*model1.Enterprise, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Enterprise{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.Enterprise
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) EnterpriseAggregate(ctx context.Context, distinctOn []model.EnterpriseSelectColumn, limit *int, offset *int, orderBy []*model.EnterpriseOrderBy, where *model.EnterpriseBoolExp) (*model.EnterpriseAggregate, error) {
	var rs model.EnterpriseAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.Enterprise{})
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

func (r *queryResolver) EnterpriseByPk(ctx context.Context, id int64) (*model1.Enterprise, error) {
	var rs model1.Enterprise
	tx := db.DB.Model(&model1.Enterprise{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
