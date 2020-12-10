package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/driving/log/model"
	"VehicleSupervision/internal/modules/driving/log/query/graph/generated"
	"VehicleSupervision/internal/modules/driving/log/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) DrivingLog(ctx context.Context, distinctOn []model.DrivingLogSelectColumn, limit *int, offset *int, orderBy []*model.DrivingLogOrderBy, where *model.DrivingLogBoolExp) ([]*model1.DrivingLog, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLog{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.DrivingLog
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DrivingLogAggregate(ctx context.Context, distinctOn []model.DrivingLogSelectColumn, limit *int, offset *int, orderBy []*model.DrivingLogOrderBy, where *model.DrivingLogBoolExp) (*model.DrivingLogAggregate, error) {
	var rs model.DrivingLogAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.DrivingLog{})
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

func (r *queryResolver) DrivingLogByPk(ctx context.Context, id int64) (*model1.DrivingLog, error) {
	var rs model1.DrivingLog
	tx := db.DB.Model(&model1.DrivingLog{}).First(&rs, id)
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
