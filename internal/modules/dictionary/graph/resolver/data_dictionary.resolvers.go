package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/dictionary/graph/generated"
	"VehicleSupervision/internal/modules/dictionary/graph/model"
	model1 "VehicleSupervision/internal/modules/dictionary/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *mutationResolver) DeleteDataDictionary(ctx context.Context, where model.DataDictionaryBoolExp) (*model.DataDictionaryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionary{})
	tx := qt.Where(where).Finish()
	// 获取请求的字段
	preloads := util.GetPreloadsMustPrefixAndRemovePrefix(ctx, "returning.")
	var rs []*model1.DataDictionary
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
	return &model.DataDictionaryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) DeleteDataDictionaryByPk(ctx context.Context, id int64) (*model1.DataDictionary, error) {
	preloads := util.GetPreloads(ctx)
	var rs model1.DataDictionary
	tx := db.DB.Model(&model1.DataDictionary{})
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

func (r *mutationResolver) InsertDataDictionary(ctx context.Context, objects []*model.DataDictionaryInsertInput, onConflict *model.DataDictionaryOnConflict) (*model.DataDictionaryMutationResponse, error) {
	rs := r.dataDictionaryInsertInputBatchConvert(objects)
	tx := db.DB.Model(&model1.DataDictionary{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model.DataDictionaryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
		Returning:    rs,
	}, nil
}

func (r *mutationResolver) InsertDataDictionaryOne(ctx context.Context, object model.DataDictionaryInsertInput, onConflict *model.DataDictionaryOnConflict) (*model1.DataDictionary, error) {
	rs := r.dataDictionaryInsertInputConvert(&object)
	tx := db.DB.Model(&model1.DataDictionary{}).Create(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *mutationResolver) UpdateDataDictionary(ctx context.Context, inc *model.DataDictionaryIncInput, set *model.DataDictionarySetInput, where model.DataDictionaryBoolExp) (*model.DataDictionaryMutationResponse, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionary{})
	tx := qt.Where(where).Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DataDictionaryMutationResponse{
				AffectedRows: 0,
			}, nil
		}
		return nil, err
	}
	return &model.DataDictionaryMutationResponse{
		AffectedRows: int(tx.RowsAffected),
	}, nil
}

func (r *mutationResolver) UpdateDataDictionaryByPk(ctx context.Context, inc *model.DataDictionaryIncInput, set *model.DataDictionarySetInput, pkColumns model.DataDictionaryPkColumnsInput) (*model1.DataDictionary, error) {
	tx := db.DB.Where("id = ?", pkColumns.ID)
	qt := util.NewQueryTranslator(tx, &model1.DataDictionary{})
	tx = qt.Inc(inc).Set(set).DoUpdate()
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	var rs model1.DataDictionary
	tx = tx.First(&rs)
	return &rs, nil
}

func (r *queryResolver) DataDictionary(ctx context.Context, distinctOn []model.DataDictionarySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryOrderBy, where *model.DataDictionaryBoolExp) ([]*model1.DataDictionary, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionary{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DataDictionary
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DataDictionaryAggregate(ctx context.Context, distinctOn []model.DataDictionarySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryOrderBy, where *model.DataDictionaryBoolExp) (*model.DataDictionaryAggregate, error) {
	var rs model.DataDictionaryAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionary{})
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

func (r *queryResolver) DataDictionaryByPk(ctx context.Context, id int64) (*model1.DataDictionary, error) {
	var rs model1.DataDictionary
	tx := db.DB.Model(&model1.DataDictionary{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
