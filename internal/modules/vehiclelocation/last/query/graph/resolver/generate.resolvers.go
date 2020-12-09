package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/vehiclelocation/last/model"
	"VehicleSupervision/internal/modules/vehiclelocation/last/query/graph/generated"
	"VehicleSupervision/internal/modules/vehiclelocation/last/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) VehicleLocationLast(ctx context.Context, distinctOn []model.VehicleLocationLastSelectColumn, limit *int, offset *int, orderBy []*model.VehicleLocationLastOrderBy, where *model.VehicleLocationLastBoolExp) ([]*model1.VehicleLocationLast, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationLast{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleLocationLast
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleLocationLastAggregate(ctx context.Context, distinctOn []model.VehicleLocationLastSelectColumn, limit *int, offset *int, orderBy []*model.VehicleLocationLastOrderBy, where *model.VehicleLocationLastBoolExp) (*model.VehicleLocationLastAggregate, error) {
	var rs model.VehicleLocationLastAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationLast{})
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

func (r *queryResolver) VehicleLocationLastByPk(ctx context.Context, id int64) (*model1.VehicleLocationLast, error) {
	var rs model1.VehicleLocationLast
	tx := db.DB.Model(&model1.VehicleLocationLast{}).First(&rs, id)
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
