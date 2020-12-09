package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/dictionary/dict/model"
	"VehicleSupervision/internal/modules/dictionary/dict/query/graph/generated"
	"VehicleSupervision/internal/modules/dictionary/dict/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
