package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/dictionary/category/model"
	"VehicleSupervision/internal/modules/dictionary/category/query/graph/generated"
	"VehicleSupervision/internal/modules/dictionary/category/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) DataDictionaryCategory(ctx context.Context, distinctOn []model.DataDictionaryCategorySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryCategoryOrderBy, where *model.DataDictionaryCategoryBoolExp) ([]*model1.DataDictionaryCategory, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionaryCategory{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DataDictionaryCategory
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DataDictionaryCategoryAggregate(ctx context.Context, distinctOn []model.DataDictionaryCategorySelectColumn, limit *int, offset *int, orderBy []*model.DataDictionaryCategoryOrderBy, where *model.DataDictionaryCategoryBoolExp) (*model.DataDictionaryCategoryAggregate, error) {
	var rs model.DataDictionaryCategoryAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DataDictionaryCategory{})
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

func (r *queryResolver) DataDictionaryCategoryByPk(ctx context.Context, id int64) (*model1.DataDictionaryCategory, error) {
	var rs model1.DataDictionaryCategory
	tx := db.DB.Model(&model1.DataDictionaryCategory{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
