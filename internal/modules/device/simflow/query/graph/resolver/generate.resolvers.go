package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/device/simflow/model"
	"VehicleSupervision/internal/modules/device/simflow/query/graph/generated"
	"VehicleSupervision/internal/modules/device/simflow/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) SimCardFlow(ctx context.Context, distinctOn []model.SimCardFlowSelectColumn, limit *int, offset *int, orderBy []*model.SimCardFlowOrderBy, where *model.SimCardFlowBoolExp) ([]*model1.SimCardFlow, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.SimCardFlow{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.SimCardFlow
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) SimCardFlowAggregate(ctx context.Context, distinctOn []model.SimCardFlowSelectColumn, limit *int, offset *int, orderBy []*model.SimCardFlowOrderBy, where *model.SimCardFlowBoolExp) (*model.SimCardFlowAggregate, error) {
	var rs model.SimCardFlowAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.SimCardFlow{})
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

func (r *queryResolver) SimCardFlowByPk(ctx context.Context, id int64) (*model1.SimCardFlow, error) {
	var rs model1.SimCardFlow
	tx := db.DB.Model(&model1.SimCardFlow{}).First(&rs, id)
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
