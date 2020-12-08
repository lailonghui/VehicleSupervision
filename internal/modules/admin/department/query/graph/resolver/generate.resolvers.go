package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/admin/department/model"
	"VehicleSupervision/internal/modules/admin/department/query/graph/generated"
	"VehicleSupervision/internal/modules/admin/department/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *queryResolver) Department(ctx context.Context, distinctOn []model.DepartmentSelectColumn, limit *int, offset *int, orderBy []*model.DepartmentOrderBy, where *model.DepartmentBoolExp) ([]*model1.Department, error) {
	qt := util.NewQueryTranslator(db.DB, &model1.Department{})
	tx := qt.DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where).
		Finish()
	var rs []*model1.Department
	tx = tx.Find(&rs)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (r *queryResolver) DepartmentAggregate(ctx context.Context, distinctOn []model.DepartmentSelectColumn, limit *int, offset *int, orderBy []*model.DepartmentOrderBy, where *model.DepartmentBoolExp) (*model.DepartmentAggregate, error) {
	var rs model.DepartmentAggregate

	qt := util.NewQueryTranslator(db.DB, &model1.Department{})
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

func (r *queryResolver) DepartmentByPk(ctx context.Context, id int64) (*model1.Department, error) {
	var rs model1.Department
	tx := db.DB.Model(&model1.Department{}).First(&rs, id)
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
