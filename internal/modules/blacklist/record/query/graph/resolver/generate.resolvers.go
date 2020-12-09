package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/blacklist/record/model"
	"VehicleSupervision/internal/modules/blacklist/record/query/graph/generated"
	"VehicleSupervision/internal/modules/blacklist/record/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) BlacklistOperationRecord(ctx context.Context, distinctOn []model.BlacklistOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.BlacklistOperationRecordOrderBy, where *model.BlacklistOperationRecordBoolExp) ([]*model1.BlacklistOperationRecord, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.BlacklistOperationRecord{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.BlacklistOperationRecord
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) BlacklistOperationRecordAggregate(ctx context.Context, distinctOn []model.BlacklistOperationRecordSelectColumn, limit *int, offset *int, orderBy []*model.BlacklistOperationRecordOrderBy, where *model.BlacklistOperationRecordBoolExp) (*model.BlacklistOperationRecordAggregate, error) {
	var rs model.BlacklistOperationRecordAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.BlacklistOperationRecord{})
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

func (r *queryResolver) BlacklistOperationRecordByPk(ctx context.Context, id int64) (*model1.BlacklistOperationRecord, error) {
	var rs model1.BlacklistOperationRecord
	tx := db.DB.Model(&model1.BlacklistOperationRecord{}).First(&rs, id)
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
