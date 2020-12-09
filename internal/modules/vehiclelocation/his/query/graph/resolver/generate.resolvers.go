package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/vehiclelocation/his/model"
	"VehicleSupervision/internal/modules/vehiclelocation/his/query/graph/generated"
	"VehicleSupervision/internal/modules/vehiclelocation/his/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) VehicleLocationHis(ctx context.Context, distinctOn []model.VehicleLocationHisSelectColumn, limit *int, offset *int, orderBy []*model.VehicleLocationHisOrderBy, where *model.VehicleLocationHisBoolExp) ([]*model1.VehicleLocationHis, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationHis{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.VehicleLocationHis
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) VehicleLocationHisAggregate(ctx context.Context, distinctOn []model.VehicleLocationHisSelectColumn, limit *int, offset *int, orderBy []*model.VehicleLocationHisOrderBy, where *model.VehicleLocationHisBoolExp) (*model.VehicleLocationHisAggregate, error) {
	var rs model.VehicleLocationHisAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.VehicleLocationHis{})
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

func (r *queryResolver) VehicleLocationHisByPk(ctx context.Context, id int64) (*model1.VehicleLocationHis, error) {
	var rs model1.VehicleLocationHis
	tx := db.DB.Model(&model1.VehicleLocationHis{}).First(&rs, id)
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
