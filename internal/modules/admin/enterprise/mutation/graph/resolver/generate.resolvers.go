package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/admin/enterprise/model"
	"VehicleSupervision/internal/modules/admin/enterprise/mutation/graph/generated"
	"VehicleSupervision/internal/modules/admin/enterprise/mutation/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"fmt"
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
	var rs = model1.Enterprise{}
	tx := db.DB.Model(&model1.Enterprise{}).Find(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	tx = db.DB.Delete(id)
	return &rs, tx.Error
}

func (r *mutationResolver) InsertEnterprise(ctx context.Context, objects []*model.EnterpriseInsertInput) (*model.EnterpriseMutationResponse, error) {
	rs := r.batchInsertParamConvert(objects)
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

func (r *mutationResolver) InsertEnterpriseOne(ctx context.Context, object model.EnterpriseInsertInput) (*model1.Enterprise, error) {
	rs := r.insertParamConvert(&object)
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

func (r *queryResolver) T(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
