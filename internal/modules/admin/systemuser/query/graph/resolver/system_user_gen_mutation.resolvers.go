package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VehicleSupervision/internal/db"
	model1 "VehicleSupervision/internal/modules/admin/systemuser/model"
	"VehicleSupervision/internal/modules/admin/systemuser/query/graph/generated"
	"VehicleSupervision/internal/modules/admin/systemuser/query/graph/model"
	"VehicleSupervision/pkg/graphql/util"
	"context"
	"fmt"
)

func (r *queryResolver) SystemUser(ctx context.Context, distinctOn []model.SystemUserSelectColumn, limit *int, offset *int, orderBy []*model.SystemUserOrderBy, where *model.SystemUserBoolExp) ([]*model1.SystemUser, error) {
	tx := db.DB.Model(&model1.SystemUser{})
	qt := util.QueryTranslator{}
	qt.Model(&model1.SystemUser{}).
		DistinctOn(distinctOn).
		Limit(limit).
		Offset(offset).
		OrderBy(orderBy).
		Where(where)
	var rs []*model1.SystemUser
	tx.Find(rs)
	return rs, nil
}

func (r *queryResolver) SystemUserAggregate(ctx context.Context, distinctOn []model.SystemUserSelectColumn, limit *int, offset *int, orderBy []*model.SystemUserOrderBy, where *model.SystemUserBoolExp) (*model.SystemUserAggregate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SystemUserByPk(ctx context.Context, id int64) (*model1.SystemUser, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
