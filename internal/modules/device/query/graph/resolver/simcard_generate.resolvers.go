package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/device/model"
	"VehicleSupervision/internal/modules/device/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) SimCard(ctx context.Context, distinctOn []model.SimCardSelectColumn, limit *int, offset *int, orderBy []*model.SimCardOrderBy, where *model.SimCardBoolExp) ([]*model1.SimCard, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SimCard{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.SimCard
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) SimCardAggregate(ctx context.Context, distinctOn []model.SimCardSelectColumn, limit *int, offset *int, orderBy []*model.SimCardOrderBy, where *model.SimCardBoolExp) (*model.SimCardAggregate, error) {
	var rs model.SimCardAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.SimCard{})
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

func (r *queryResolver) SimCardByPk(ctx context.Context, id int64) (*model1.SimCard, error) {
	var rs model1.SimCard
	tx := db.DB.Model(&model1.SimCard{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
