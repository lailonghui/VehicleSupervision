package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	"VehicleSupervision/internal/modules/area/query/graph/generated"
	"VehicleSupervision/internal/modules/area/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) City(ctx context.Context, distinctOn []model.CitySelectColumn, limit *int, offset *int, orderBy []*model.CityOrderBy, where *model.CityBoolExp) ([]*model.City, error) {
	qt := util.NewQueryTranslator(db.DB, &model.City{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.City
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) CityAggregate(ctx context.Context, distinctOn []model.CitySelectColumn, limit *int, offset *int, orderBy []*model.CityOrderBy, where *model.CityBoolExp) (*model.CityAggregate, error) {
	var rs model.CityAggregate

	qt := util.NewQueryTranslator(db.DB, &model.City{})
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

func (r *queryResolver) CityByPk(ctx context.Context, id int64) (*model.City, error) {
	var rs model.City
	tx := db.DB.Model(&model.City{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *queryResolver) District(ctx context.Context, distinctOn []model.DistrictSelectColumn, limit *int, offset *int, orderBy []*model.DistrictOrderBy, where *model.DistrictBoolExp) ([]*model.District, error) {
	qt := util.NewQueryTranslator(db.DB, &model.District{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.District
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DistrictAggregate(ctx context.Context, distinctOn []model.DistrictSelectColumn, limit *int, offset *int, orderBy []*model.DistrictOrderBy, where *model.DistrictBoolExp) (*model.DistrictAggregate, error) {
	var rs model.DistrictAggregate

	qt := util.NewQueryTranslator(db.DB, &model.District{})
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

func (r *queryResolver) DistrictByPk(ctx context.Context, id int64) (*model.District, error) {
	var rs model.District
	tx := db.DB.Model(&model.District{}).First(&rs, id)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}

func (r *queryResolver) Province(ctx context.Context, distinctOn []model.ProvinceSelectColumn, limit *int, offset *int, orderBy []*model.ProvinceOrderBy, where *model.ProvinceBoolExp) ([]*model.Province, error) {
	qt := util.NewQueryTranslator(db.DB, &model.Province{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model.Province
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) ProvinceAggregate(ctx context.Context, distinctOn []model.ProvinceSelectColumn, limit *int, offset *int, orderBy []*model.ProvinceOrderBy, where *model.ProvinceBoolExp) (*model.ProvinceAggregate, error) {
	var rs model.ProvinceAggregate

	qt := util.NewQueryTranslator(db.DB, &model.Province{})
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

func (r *queryResolver) ProvinceByPk(ctx context.Context, id int64) (*model.Province, error) {
	var rs model.Province
	tx := db.DB.Model(&model.Province{}).First(&rs, id)
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
